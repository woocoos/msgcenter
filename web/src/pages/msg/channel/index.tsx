import { ActionType, PageContainer, ProColumns, ProTable, useToken } from '@ant-design/pro-components';
import { Button, Space, Modal, Divider, Typography, Tag } from 'antd';
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
import { delDataSource, saveDataSource, onMousedown } from '@/util';
import { useResizableProTable, routeBreadcrumb } from '@/util/hook';

const List = () => {
  const { token } = useToken(),
    { t } = useTranslation(),
    // 表格相关
    proTableRef = useRef<ActionType>(),
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
    }),
    // 可调整列宽的 ProTable
    { columns: finalColumns, components, tableWidth } = useResizableProTable<MsgChannel>(() => ({
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
        {
          title: t('org'), dataIndex: 'org', width: 150,
          renderFormItem: () => {
            return <OrgSelect kind={OrgKind.Root} />
          },
          render: (text, record) => {
            const org = orgs.find(item => item.id == record.tenantID)
            return record.tenantID ? org?.name || record.tenantID : '-';
          },
        },
        { title: t('name'), dataIndex: 'name', width: 160 },
        {
          title: t('type'), dataIndex: 'receiverType', width: 120, align: 'center', search: false,
          filters: true, valueEnum: EnumMsgChannelReceiverType
        },
        {
          title: t('status'), dataIndex: 'status', width: 120, align: 'center', search: false,
          filters: true,
          valueEnum: EnumMsgChannelStatus,
          render(_, record) {
            const enumItem = record.status ? EnumMsgChannelStatus[record.status] : undefined;
            return <Tag color={enumItem?.tagColor}>{enumItem?.text ?? '-'}</Tag>;
          },
        },
        { title: t('description'), dataIndex: 'comments', width: 200, search: false, ellipsis: true },
        // 占位列
        { search: false, hideInSetting: true },
        // 操作列
        {
          title: t('operation'),
          dataIndex: 'actions',
          fixed: 'right',
          align: 'center',
          search: false,
          width: 220,
          hideInSetting: true,
          render: (text, record) => {
            return (<Space split={<Divider type="vertical" className="ko-divider-gray" />} size={0}>
              <Auth authKey="updateMsgChannel">
                <Typography.Link
                  key="editor"
                  onClick={() => {
                    setModal({
                      open: true, title: `${t('edit')}:${record.name}`, id: record.id, scene: 'editor'
                    });
                  }}
                >
                  {t('edit')}
                </Typography.Link>
              </Auth>
              <Auth authKey="updateMsgChannel">
                <Typography.Link
                  key="config"
                  onClick={() => {
                    setModal({
                      open: true, title: `${t('amend_msg_channel_config')}:${record.name}`, id: record.id, scene: 'config'
                    });
                  }}
                >
                  {t('configuration')}
                </Typography.Link>
              </Auth>
              {
                record.status === MsgChannelSimpleStatus.Active ? <></> : <Auth authKey="deleteMsgChannel">
                  <Typography.Link key="delete" onClick={() => onDel(record)}>
                    {t('delete')}
                  </Typography.Link>
                </Auth>
              }
              {
                record.status === MsgChannelSimpleStatus.Active ? <Auth authKey="disableMsgChannel">
                  <Typography.Link key="disable" style={{ color: '#ff0000' }} onClick={() => onClickStatus(record)}>
                    {t('disable')}
                  </Typography.Link>
                </Auth> : <Auth authKey="enableMsgChannel">
                  <Typography.Link key="enable" onClick={() => onClickStatus(record)}>
                    {t('enable')}
                  </Typography.Link>
                </Auth>
              }
            </Space>);
          },
        },
      ],
    }), [orgs]);


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
    <>
      <ProTable
        actionRef={proTableRef}
        sticky={dataSource.length > 0 ? { offsetHeader: 56 } : undefined}
        search={{
          className: 'ko-pro-table-search',
          searchText: `${t('query')}`,
          resetText: `${t('reset')}`,
          labelWidth: 70,
        }}
        rowKey={'id'}
        toolbar={{
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
        scroll={{ x: tableWidth }}
        components={components}
        columns={finalColumns}
        columnsState={{ defaultValue: { id: { show: false } } }}
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
        pagination={{ showSizeChanger: true }}
        rowSelection={{
          type: 'radio',
          hideSelectAll: false,
          selectedRowKeys,
          onChange: (rowKeys) => setSelectedRowKeys(rowKeys as string[]),
        }}
        onRow={(record) => ({
          onMouseDown: (e) => {
            onMousedown({
              target: e.target as HTMLElement,
              click: () => {
                setSelectedRowKeys(prev =>
                  prev.includes(record.id) ? [] : [record.id]
                );
              },
            });
          },
        })}
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
    </>
  );
};

export default () => {
  const [breadcrumbNames] = routeBreadcrumb();
  return (
    <PageContainer
      className="ko-page-container"
      header={{
        breadcrumb: {
          items: breadcrumbNames.map(item => ({ title: item })),
        },
      }}
    >
      <KeepAlive clearAlive>
        <List />
      </KeepAlive>
    </PageContainer>
  );
};

export const pageConfig = definePageConfig(() => ({
  auth: ['/msg/channel'],
}));
