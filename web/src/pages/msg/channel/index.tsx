import { ActionType, PageContainer, ProColumns, ProTable, useToken } from '@ant-design/pro-components';
import { Button, Space, Modal } from 'antd';
import { useRef, useState } from 'react';
import { useTranslation } from 'react-i18next';
import Auth from '@/components/auth';
import { MsgChannel, MsgChannelReceiverType, MsgChannelSimpleStatus, MsgChannelWhereInput } from '@/generated/msgsrv/graphql';
import { EnumMsgChannelReceiverType, EnumMsgChannelStatus, delMsgChannel, disableMsgChannel, enableMsgChannel, getMsgChannelList } from '@/services/msgsrv/channel';
import { getOrgs } from '@knockout-js/api';
import Create from './components/create';
import Config from './components/config';
import { OrgSelect } from '@knockout-js/org';
import { OrgKind, Org } from '@knockout-js/api/ucenter';
import { KeepAlive } from '@knockout-js/layout';
import ConfigExample from './components/configExample';
import { definePageConfig } from 'ice';
import { delDataSource, saveDataSource } from '@/util';
import { it } from 'node:test';


export default () => {
  const { token } = useToken(),
    { t } = useTranslation(),
    // 表格相关
    proTableRef = useRef<ActionType>(),
    columns: ProColumns<MsgChannel>[] = [
      // 有需要排序配置  sorter: true
      {
        title: t('org'), dataIndex: 'org', width: 120,
        renderFormItem: () => {
          return <OrgSelect kind={OrgKind.Root} />
        },
        render: (text, record) => {
          const org = orgs.find(item => item.id == record.tenantID)
          return record.tenantID ? org?.name || record.tenantID : '-';
        },
      },
      { title: t('name'), dataIndex: 'name', width: 120 },
      {
        title: t('type'), dataIndex: 'receiverType', width: 120, search: false,
        filters: true, valueEnum: EnumMsgChannelReceiverType
      },

      {
        title: t('status'), dataIndex: 'status', width: 120, search: false,
        filters: true,
        valueEnum: EnumMsgChannelStatus,
      },
      { title: t('description'), dataIndex: 'comments', width: 120, search: false },
      {
        title: t('operation'),
        dataIndex: 'actions',
        fixed: 'right',
        align: 'center',
        search: false,
        width: 190,
        render: (text, record) => {
          return (<Space>
            <Auth authKey="updateMsgChannel">
              <a
                key="editor"
                onClick={() => {
                  setModal({
                    open: true, title: `${t('edit')}:${record.name}`, id: record.id, scene: 'editor'
                  });
                }}
              >
                {t('edit')}
              </a>
            </Auth>
            <Auth authKey="updateMsgChannel">
              <a
                key="config"
                onClick={() => {
                  setModal({
                    open: true, title: `${t('amend_msg_channel_config')}:${record.name}`, id: record.id, scene: 'config'
                  });
                }}
              >
                {t('configuration')}
              </a>
            </Auth>
            {
              record.status === MsgChannelSimpleStatus.Active ? <></> : <Auth authKey="deleteMsgChannel">
                <a key="delete" onClick={() => onDel(record)}>
                  {t('delete')}
                </a>
              </Auth>
            }
            {
              record.status === MsgChannelSimpleStatus.Active ? <Auth authKey="disableMsgChannel">
                <a key="disable" style={{ color: '#ff0000' }} onClick={() => onClickStatus(record)}>
                  {t('disable')}
                </a>
              </Auth> : <Auth authKey="enableMsgChannel">
                <a key="enable" onClick={() => onClickStatus(record)}>
                  {t('enable')}
                </a>
              </Auth>
            }
          </Space>);
        },
      },
    ],
    [orgs, setOrgs] = useState<Org[]>([]),
    [dataSource, setDataSource] = useState<MsgChannel[]>([]),
    // 选中处理
    [selectedRowKeys, setSelectedRowKeys] = useState<string[]>([]),
    // 弹出层处理
    [modal, setModal] = useState<{
      open: boolean;
      title: string;
      id: string;
      scene: 'editor' | 'config' | 'config_example';
    }>({
      open: false,
      title: '',
      id: '',
      scene: 'editor'
    });


  const
    onDel = (record: MsgChannel) => {
      Modal.confirm({
        title: t('delete'),
        content: `${t('confirm_delete')}：${record.name}`,
        onOk: async (close) => {
          const result = await delMsgChannel(record.id);
          if (result === true) {
            setDataSource(delDataSource(dataSource, record.id))
            if (dataSource.length === 0) {
              const pageInfo = { ...proTableRef.current?.pageInfo };
              pageInfo.current = pageInfo.current ? pageInfo.current > 2 ? pageInfo.current - 1 : 1 : 1;
              proTableRef.current?.setPageInfo?.(pageInfo);
              proTableRef.current?.reload();
            }
            close();
          }
        },
      });
    },
    onClickStatus = (record: MsgChannel) => {
      Modal.confirm({
        title: record.status === MsgChannelSimpleStatus.Active ? t('disable') : t('enable'),
        content: `${record.status === MsgChannelSimpleStatus.Active ? t('disable') : t('enable')}：${record.name}`,
        onOk: async (close) => {
          const result = record.status === MsgChannelSimpleStatus.Active ? await disableMsgChannel(record.id) : await enableMsgChannel(record.id);
          if (result?.id) {
            setDataSource(saveDataSource(dataSource, result as MsgChannel))
            close();
          }
        },
      });
    };


  return (
    <KeepAlive clearAlive>
      <PageContainer
        header={{
          title: t('msg_channel'),
          style: { background: token.colorBgContainer },
          breadcrumb: {
            items: [
              { title: t('msg_center') },
              { title: t('msg_channel') },
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
            title: t('msg_channel_list'),
            actions: [
              <Auth authKey="createMsgChannel">
                <Button
                  key="created"
                  type="primary"
                  onClick={() => {
                    setModal({ open: true, title: t('create_msg_channel'), id: '', scene: 'editor' });
                  }}
                >
                  {t('create_msg_channel')}
                </Button>
              </Auth>,
              <Button
                onClick={() => {
                  setModal({ open: true, title: t('msg_event_config_example'), id: '', scene: 'config_example' });
                }}
              >
                {t('msg_event_config_example')}
              </Button>
            ],
          }}
          scroll={{ x: 'max-content' }}
          columns={columns}
          dataSource={dataSource}
          request={async (params, sort, filter) => {
            const table = { data: [] as MsgChannel[], success: true, total: 0 },
              where: MsgChannelWhereInput = {};
            where.tenantID = params.org?.id;
            where.nameContains = params.name;
            where.receiverTypeIn = filter.receiverType as MsgChannelReceiverType[];
            where.statusIn = filter.status as MsgChannelSimpleStatus[]
            const result = await getMsgChannelList({
              current: params.current,
              pageSize: params.pageSize,
              where,
            });
            if (result?.totalCount) {
              table.data = result.edges?.map(item => item?.node) as MsgChannel[]
              setOrgs(await getOrgs(table.data.map(item => item.tenantID || '')))
              table.total = result.totalCount;
            }
            setSelectedRowKeys([]);
            setDataSource(table.data);
            return table;
          }}
          rowSelection={{
            selectedRowKeys: selectedRowKeys,
            onChange: (selectedRowKeys: string[]) => { setSelectedRowKeys(selectedRowKeys); },
            type: 'checkbox',
          }}
        />
        <Create
          x-if={modal.scene === 'editor'}
          open={modal.open}
          title={modal.title}
          id={modal.id}
          onClose={async (isSuccess, newInfo) => {
            if (isSuccess && newInfo) {
              if (!orgs.find(item => item.id === newInfo.tenantID)) {
                setOrgs([...orgs, ...(await getOrgs([newInfo.tenantID]))])
              }
              setDataSource(saveDataSource(dataSource, newInfo))
            }
            setModal({ open: false, title: modal.title, id: '', scene: modal.scene });
          }}
        />
        <Config
          x-if={modal.scene === 'config'}
          open={modal.open}
          title={modal.title}
          id={modal.id}
          onClose={(isSuccess, newInfo) => {
            if (isSuccess && newInfo) {
              setDataSource(saveDataSource(dataSource, newInfo))
            }
            setModal({ open: false, title: modal.title, id: '', scene: modal.scene });
          }}
        />
        <ConfigExample
          x-if={modal.scene === 'config_example'}
          open={modal.open}
          title={modal.title}
          onClose={() => {
            setModal({ open: false, title: modal.title, id: '', scene: modal.scene });
          }}
        />
      </PageContainer>
    </KeepAlive>
  );
};

export const pageConfig = definePageConfig(() => ({
  auth: ['/msg/channel'],
}));
