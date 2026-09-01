import { ActionType, PageContainer, ProColumns, ProTable, useToken } from '@ant-design/pro-components';
import { Space } from 'antd';
import { useRef, useState } from 'react';
import { useTranslation } from 'react-i18next';
import Auth from '@/components/auth';
import { MsgEvent, MsgEventSimpleStatus, MsgEventWhereInput } from '@/generated/msgsrv/graphql';
import { getMsgEventListWithSubs } from '@/services/msgsrv/event';
import EventSettings from './components/settings';
import { getOrgRoles, getUsers } from '@knockout-js/api';
import { KeepAlive } from '@knockout-js/layout';
import { definePageConfig, useSearchParams } from 'ice';

type ProTableColumnsData = {
  id: string;
  name: string;
  comments?: string;
  receiving_user: string;
  receiving_user_group: string;
  exclude_user: string;
  msgEvent?: MsgEvent;
}

export default () => {
  const { token } = useToken(),
    { t } = useTranslation(),
    [searchParams] = useSearchParams(),
    msgTypeId = searchParams.get('msgTypeId') || '',
    msgTypeName = decodeURIComponent(searchParams.get('msgTypeName') || ''),
    // 表格相关
    proTableRef = useRef<ActionType>(),
    columns: ProColumns<ProTableColumnsData>[] = [
      { title: t('msg_event_name'), dataIndex: 'name', width: 200 },
      { title: t('description'), dataIndex: 'comments', width: 200, search: false },
      { title: t('receiving_user'), dataIndex: 'receiving_user', width: 120, search: false },
      { title: t('receiving_user_group'), dataIndex: 'receiving_user_group', width: 120, search: false },
      { title: t('exclude_user'), dataIndex: 'exclude_user', width: 120, search: false },
      {
        title: t('operation'),
        dataIndex: 'actions',
        fixed: 'right',
        align: 'center',
        search: false,
        width: 80,
        render: (text, record) => {
          return record.msgEvent ? <Space>
            <Auth authKey={['createMsgSubscriber', 'deleteMsgSubscriber']}>
              <a
                key="settings"
                onClick={() => {
                  setModal({
                    open: true, title: `${t('settings')}: ${record.name}`, id: record.id, msgEvent: record.msgEvent
                  });
                }}
              >
                {t('settings')}
              </a>
            </Auth>
          </Space> : <></>;
        },
      },
    ],
    // 弹出层处理
    [modal, setModal] = useState<{
      open: boolean;
      title: string;
      id: string;
      msgEvent?: MsgEvent;
    }>({
      open: false,
      title: '',
      id: '',
    });

  return (
    <KeepAlive clearAlive>
      <PageContainer
        header={{
          title: `${t('event_subscription')}: ${msgTypeName}`,
          style: { background: token.colorBgContainer },
          breadcrumb: {
            items: [
              { title: t('msg_center') },
              { title: t('msg_subscription'), href: '/msg/subscription' },
              { title: t('event_subscription') },
            ],
          },
        }}
      >
        <ProTable
          actionRef={proTableRef}
          search={{
            searchText: `${t('query')}`,
            resetText: `${t('reset')}`,
            labelWidth: 'auto',
          }}
          rowKey={'id'}
          toolbar={{
            title: t('event_subscription_list'),
          }}
          scroll={{ x: 'max-content' }}
          columns={columns}
          request={async (params) => {
            const table = { data: [] as ProTableColumnsData[], success: true, total: 0 },
              where: MsgEventWhereInput = {
                canSubs: true,
                status: MsgEventSimpleStatus.Active,
                hasMsgTypeWith: [{ id: msgTypeId }],
              };
            where.nameContains = params.name;

            const result = await getMsgEventListWithSubs({
              current: params.current,
              pageSize: 999,
              where,
            });

            if (result?.totalCount) {
              const msgEventList = result.edges?.map(item => item?.node),
                userIds: string[] = [],
                userGroupIds: string[] = [],
                data: ProTableColumnsData[] = [];

              msgEventList?.forEach(item => {
                if (item) {
                  item.subscriberUsers?.forEach(su => {
                    if (su.userID) {
                      userIds.push(su.userID);
                    }
                  });
                  item.excludeSubscriberUsers?.forEach(su => {
                    if (su.userID) {
                      userIds.push(su.userID);
                    }
                  });
                  item.subscriberRoles?.forEach(sr => {
                    if (sr.orgRoleID) {
                      userGroupIds.push(`${sr.orgRoleID}`);
                    }
                  });
                }
              });

              const users = await getUsers(userIds);
              const userGroups = await getOrgRoles(userGroupIds);

              msgEventList?.forEach(me => {
                if (me) {
                  data.push({
                    id: me.id,
                    name: me.name,
                    comments: me.comments || '',
                    receiving_user: me.subscriberUsers?.map(su => {
                      const user = users.find(u => u.id == su.userID);
                      return su.userID ? user?.displayName : '';
                    }).filter(su => !!su).join('、') || '',
                    receiving_user_group: me.subscriberRoles?.map(sr => {
                      const userGroup = userGroups.find(ug => ug.id == sr.orgRoleID);
                      return sr.orgRoleID ? userGroup?.name : '';
                    }).filter(sr => !!sr).join('、') || '',
                    exclude_user: me.excludeSubscriberUsers?.map(su => {
                      const user = users.find(u => u.id == su.userID);
                      return su.userID ? user?.displayName : '';
                    }).filter(su => !!su).join('、') || '',
                    msgEvent: me as MsgEvent,
                  });
                }
              });

              table.data = data;
              table.total = data.length;
            }
            return table;
          }}
          pagination={false}
        />
        <EventSettings
          open={modal.open}
          title={modal.title}
          id={modal.id}
          msgEvent={modal.msgEvent}
          onClose={(isSuccess) => {
            if (isSuccess) {
              proTableRef.current?.reload();
            }
            setModal({ open: false, title: modal.title, id: '' });
          }}
        />
      </PageContainer>
    </KeepAlive>
  );
};

// export const pageConfig = definePageConfig(() => ({
//   auth: ['/msg/subscription/events'],
// }));
