import { ActionType, PageContainer, ProTable } from '@ant-design/pro-components';
import { Space, Divider, Typography } from 'antd';
import { useRef, useState } from 'react';
import { useTranslation } from 'react-i18next';
import Auth from '@/components/auth';
import { MsgEvent, MsgEventSimpleStatus, MsgEventWhereInput } from '@/generated/msgsrv/graphql';
import { getMsgEventListWithSubs } from '@/services/msgsrv/event';
import EventSettings from './components/settings';
import { getOrgRoles, getUsers } from '@knockout-js/api';
import { definePageConfig, useSearchParams } from 'ice';
import { useResizableProTable, routeBreadcrumb } from '@/util/hook';

type ProTableColumnsData = {
  id: string;
  name: string;
  comments?: string;
  receiving_user: string;
  receiving_user_group: string;
  exclude_user: string;
  msgEvent?: MsgEvent;
}

const List = () => {
  const { t } = useTranslation(),
    [searchParams] = useSearchParams(),
    msgTypeId = searchParams.get('msgTypeId') || '',
    msgTypeName = decodeURIComponent(searchParams.get('msgTypeName') || ''),
    // 表格相关
    proTableRef = useRef<ActionType>(),
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
    }),
    // 可调整列宽的 ProTable
    { columns: finalColumns, components, tableWidth } = useResizableProTable<ProTableColumnsData>(() => ({
      columns: [
        // ID 列（首位，默认隐藏，可复制）
        {
          title: 'ID',
          dataIndex: 'id',
          width: 100,
          order: -999,
          render(_, record) {
            return (
              <Typography.Text
                copyable={{ text: record.id, tooltips: ['复制 ID', '已复制'] }}
                onClick={(e) => e.stopPropagation()}
              >
                {record.id}
              </Typography.Text>
            );
          },
        },
        { title: t('msg_event_name'), dataIndex: 'name', width: 200 },
        { title: t('description'), dataIndex: 'comments', width: 200, search: false, ellipsis: true },
        { title: t('receiving_user'), dataIndex: 'receiving_user', width: 160, search: false },
        { title: t('receiving_user_group'), dataIndex: 'receiving_user_group', width: 160, search: false },
        { title: t('exclude_user'), dataIndex: 'exclude_user', width: 160, search: false },
        // 占位列
        { search: false, hideInSetting: true },
        // 操作列
        {
          title: t('operation'),
          dataIndex: 'actions',
          fixed: 'right',
          align: 'center',
          search: false,
          width: 100,
          hideInSetting: true,
          render: (text, record) => {
            return record.msgEvent ? <Space split={<Divider type="vertical" className="ko-divider-gray" />} size={0}>
              <Auth authKey={['createMsgSubscriber', 'deleteMsgSubscriber']}>
                <Typography.Link
                  key="settings"
                  onClick={() => {
                    setModal({
                      open: true, title: `${t('settings')}: ${record.name}`, id: record.id, msgEvent: record.msgEvent
                    });
                  }}
                >
                  {t('settings')}
                </Typography.Link>
              </Auth>
            </Space> : <></>;
          },
        },
      ],
    }), []);

  return (
    <>
      <ProTable
        actionRef={proTableRef}
        sticky={{ offsetHeader: 56 }}
        search={{
          className: 'ko-pro-table-search',
          searchText: `${t('query')}`,
          resetText: `${t('reset')}`,
          labelWidth: 96,
        }}
        rowKey={'id'}
        toolbar={{
          title: `${t('event_subscription')}-${t('msg_type')}: ${msgTypeName}`,
        }}
        scroll={{ x: tableWidth }}
        components={components}
        columns={finalColumns}
        columnsState={{ defaultValue: { id: { show: false } } }}
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
    </>
  );
};

export default () => {
  return (
    <PageContainer
      className="ko-page-container"
      header={{
      }}
    >
      <div className="ka-content">
        <List />
      </div>
    </PageContainer>
  );
};

// export const pageConfig = definePageConfig(() => ({
//   auth: ['/msg/subscription/events'],
// }));
