import { ActionType, PageContainer, ProColumns, ProTable, useToken } from '@ant-design/pro-components';
import { Space } from 'antd';
import { useRef, useState } from 'react';
import { useTranslation } from 'react-i18next';
import Auth from '@/components/auth';
import { MsgType, MsgTypeSimpleStatus, MsgTypeWhereInput } from '@/generated/msgsrv/graphql';
import { getMsgTypeListAndSub } from '@/services/msgsrv/type';
import Settings from './components/settings';
import { getOrgRoles, getUsers } from '@knockout-js/api';
import { KeepAlive } from '@knockout-js/layout';
import { DictSelect, DictText } from '@knockout-js/org';

type ProTableColumnsData = {
  id: string;
  name: string;
  receiving_user: string;
  receiving_user_group: string;
  exclude_user: string;
  children?: ProTableColumnsData[];
  msgType?: MsgType;
}

export default () => {
  const { token } = useToken(),
    { t } = useTranslation(),
    // 表格相关
    proTableRef = useRef<ActionType>(),
    columns: ProColumns<ProTableColumnsData>[] = [
      // 有需要排序配置  sorter: true
      {
        title: t('msg_type'), dataIndex: 'name', width: 200,
        renderText(text, record, index, action) {
          return record.id.indexOf('category-') === 0 ? <DictText dictCode="MsgCategory" value={record.name} /> : record.name
        },
      },
      {
        title: t('msg_type_category'), dataIndex: 'msgTypeCategory', hideInTable: true,
        renderFormItem() {
          return <DictSelect dictCode="MsgCategory" placeholder={t('please_enter_category')} />
        },
      },
      { title: t('receiving_user'), dataIndex: 'receiving_user', width: 120, search: false },
      { title: t('receiving_user_group'), dataIndex: 'receiving_user_group', width: 120, search: false },
      { title: t('exclude_user'), dataIndex: 'exclude_user', width: 120, search: false },
      {
        title: t('operation'),
        dataIndex: 'actions',
        fixed: 'right',
        align: 'center',
        search: false,
        width: 60,
        render: (text, record) => {
          return record.msgType ? <Space>
            <Auth authKey={['createMsgSubscriber', 'deleteMsgSubscriber']}>
              <a
                key="settings"
                onClick={() => {
                  setModal({
                    open: true, title: t('settings'), id: record.id, msgType: record.msgType
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
    [expandedRowKeys, setExpandedRowKeys] = useState<string[]>([]),
    // 弹出层处理
    [modal, setModal] = useState<{
      open: boolean;
      title: string;
      id: string;
      msgType?: MsgType;
    }>({
      open: false,
      title: '',
      id: '',
    });

  return (
    <KeepAlive clearAlive>
      <PageContainer
        header={{
          title: t('msg_subscription'),
          style: { background: token.colorBgContainer },
          breadcrumb: {
            items: [
              { title: t('msg_center') },
              { title: t('msg_subscription') },
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
            title: t('msg_subscription_list'),
          }}
          scroll={{ x: 'max-content' }}
          columns={columns}
          request={async (params) => {
            const table = { data: [] as ProTableColumnsData[], success: true, total: 0 },
              where: MsgTypeWhereInput = {
                canSubs: true,
                status: MsgTypeSimpleStatus.Active,
              };
            where.nameContains = params.name;
            where.categoryContains = params.msgTypeCategory;
            const result = await getMsgTypeListAndSub({
              current: params.current,
              pageSize: 999,
              where,
            });
            if (result?.totalCount) {
              const msgTypeList = result.edges?.map(item => item?.node),
                userIds: string[] = [],
                userGroupIds: string[] = [],
                data: ProTableColumnsData[] = [];

              msgTypeList?.forEach(item => {
                if (item) {
                  item.subscriberUsers.forEach(su => {
                    if (su.userID) {
                      userIds.push(su.userID)
                    }
                  })
                  item.excludeSubscriberUsers.forEach(su => {
                    if (su.userID) {
                      userIds.push(su.userID)
                    }
                  })
                  item.subscriberRoles.forEach(sr => {
                    if (sr.orgRoleID) {
                      userGroupIds.push(`${sr.orgRoleID}`)
                    }
                  })
                }
              })

              const users = await getUsers(userIds);
              const userGroups = await getOrgRoles(userGroupIds);

              msgTypeList?.forEach(mt => {
                if (mt) {
                  const dataItem = data.find(item => item.name == mt.category),
                    addData = {
                      id: mt.id,
                      name: mt.name,
                      receiving_user: mt.subscriberUsers.map(su => {
                        const user = users.find(u => u.id == su.userID);
                        return su.userID ? user?.displayName : ''
                      }).filter(su => !!su).join('、'),
                      receiving_user_group: mt.subscriberRoles.map(sr => {
                        const userGroup = userGroups.find(ug => ug.id == sr.orgRoleID);
                        return sr.orgRoleID ? userGroup?.name : ''
                      }).filter(sr => !!sr).join('、'),
                      exclude_user: mt.excludeSubscriberUsers.map(su => {
                        const user = users.find(u => u.id == su.userID);
                        return su.userID ? user?.displayName : ''
                      }).filter(su => !!su).join('、'),
                      msgType: mt as MsgType,
                    }
                  if (dataItem) {
                    dataItem.children?.push(addData)
                  } else {
                    data.push({
                      id: `category-${mt.category}`,
                      name: mt.category,
                      receiving_user: '',
                      receiving_user_group: '',
                      exclude_user: '',
                      children: [addData]
                    })
                  }
                }
              })
              table.data = data
              table.total = data.length;

            }
            setExpandedRowKeys(table.data.map(item => item.id))
            return table;
          }}
          pagination={false}
          expandable={{
            expandedRowKeys,
          }}
        />
        <Settings
          open={modal.open}
          title={modal.title}
          id={modal.id}
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
