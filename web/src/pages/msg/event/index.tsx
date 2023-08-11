import { ActionType, PageContainer, ProColumns, ProTable, useToken } from '@ant-design/pro-components';
import { Button, Space, Modal } from 'antd';
import { useRef, useState } from 'react';
import { useTranslation } from 'react-i18next';
import Auth from '@/components/Auth';
import { MsgEvent, MsgEventSimpleStatus, MsgEventWhereInput } from '@/__generated__/msgsrv/graphql';
import { EnumMsgEventStatus, delMsgEvent, disableMsgEvent, enableMsgEvent, getMsgEventList } from '@/services/msgsrv/event';
import Create from './components/create';
import { Link } from '@ice/runtime';
import Config from './components/config';


export default () => {
  const { token } = useToken(),
    { t } = useTranslation(),
    // 表格相关
    proTableRef = useRef<ActionType>(),
    columns: ProColumns<MsgEvent>[] = [
      // 有需要排序配置  sorter: true
      {
        title: t('msg_type_category'), dataIndex: 'msgTypeCategory', width: 120,
        render(text, record) {
          return record.msgType.category
        },
      },
      {
        title: t('msg_type_name'), dataIndex: 'msgTypeName', width: 120,
        render(text, record) {
          return record.msgType.name
        },
      },
      { title: t('msg_event_name'), dataIndex: 'name', width: 120 },
      {
        title: t('way_receiving'), dataIndex: 'modes', width: 120, search: false,
        render(text, record) {
          return record.modes.split(',').join('、')
        },
      },
      {
        title: t('status'), dataIndex: 'status', width: 120, search: false,
        filters: true,
        valueEnum: EnumMsgEventStatus,
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
            <Auth authKey="updateMsgEvent">
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
            <Link
              key="template"
              to={`/msg/template?id=${record.id}`}
            >
              {t('template')}
            </Link>
            <Auth authKey="updateMsgEvent">
              <a
                key="config"
                onClick={() => {
                  setModal({
                    open: true, title: `${t('amend_msg_event_config')}:${record.name}`, id: record.id, scene: 'config'
                  });
                }}
              >
                {t('configuration')}
              </a>
            </Auth>
            <Auth authKey="deleteMsgEvent">
              <a key="delete" onClick={() => onDel(record)}>
                {t('delete')}
              </a>
            </Auth>
            {
              record.status === MsgEventSimpleStatus.Active ? <Auth authKey="disableMsgEvent">
                <a key="disable" style={{ color: '#ff0000' }} onClick={() => onClickStatus(record)}>
                  {t('disable')}
                </a>
              </Auth> : <Auth authKey="enableMsgEvent">
                <a key="enable" onClick={() => onClickStatus(record)}>
                  {t('enable')}
                </a>
              </Auth>
            }
          </Space>);
        },
      },
    ],
    [dataSource, setDataSource] = useState<MsgEvent[]>([]),
    // 选中处理
    [selectedRowKeys, setSelectedRowKeys] = useState<string[]>([]),
    // 弹出层处理
    [modal, setModal] = useState<{
      open: boolean;
      title: string;
      id: string;
      scene: 'editor' | 'config';
    }>({
      open: false,
      title: '',
      id: '',
      scene: 'editor'
    });


  const
    onDel = (record: MsgEvent) => {
      Modal.confirm({
        title: t('delete'),
        content: `${t('confirm_delete')}：${record.name}`,
        onOk: async (close) => {
          const result = await delMsgEvent(record.id);
          if (result === true) {
            if (dataSource.length === 1) {
              const pageInfo = { ...proTableRef.current?.pageInfo };
              pageInfo.current = pageInfo.current ? pageInfo.current > 2 ? pageInfo.current - 1 : 1 : 1;
              proTableRef.current?.setPageInfo?.(pageInfo);
            }
            proTableRef.current?.reload();
            close();
          }
        },
      });
    },
    onClickStatus = (record: MsgEvent) => {
      Modal.confirm({
        title: record.status === MsgEventSimpleStatus.Active ? t('disable') : t('enable'),
        content: `${record.status === MsgEventSimpleStatus.Active ? t('disable') : t('enable')}：${record.name}`,
        onOk: async (close) => {
          const result = record.status === MsgEventSimpleStatus.Active ? await disableMsgEvent(record.id) : await enableMsgEvent(record.id);
          if (result?.id) {
            proTableRef.current?.reload();
            close();
          }
        },
      });
    };


  return (
    <PageContainer
      header={{
        title: t('msg_event'),
        style: { background: token.colorBgContainer },
        breadcrumb: {
          items: [
            { title: t('msg_center') },
            { title: t('msg_event') },
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
          title: t('msg_event_list'),
          actions: [
            <Auth authKey="createMsgEvent">
              <Button
                key="created"
                type="primary"
                onClick={() => {
                  setModal({ open: true, title: t('create_msg_event'), id: '', scene: 'editor' });
                }}
              >
                {t('create_msg_event')}
              </Button>
            </Auth>,
          ],
        }}
        scroll={{ x: 'max-content' }}
        columns={columns}
        request={async (params, sort, filter) => {
          const table = { data: [] as MsgEvent[], success: true, total: 0 },
            where: MsgEventWhereInput = {};
          where.nameContains = params.name;
          where.hasMsgTypeWith = [{
            nameContains: params.msgTypeName,
            categoryContains: params.msgTypeCategory,
          }];
          where.statusIn = filter.status as MsgEventSimpleStatus[]
          const result = await getMsgEventList({
            current: params.current,
            pageSize: params.pageSize,
            where,
          });
          if (result?.totalCount) {
            table.data = result.edges?.map(item => item?.node) as MsgEvent[]
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
        onClose={(isSuccess) => {
          if (isSuccess) {
            proTableRef.current?.reload();
          }
          setModal({ open: false, title: modal.title, id: '', scene: modal.scene });
        }}
      />
      <Config
        x-if={modal.scene === 'config'}
        open={modal.open}
        title={modal.title}
        id={modal.id}
        onClose={(isSuccess) => {
          if (isSuccess) {
            proTableRef.current?.reload();
          }
          setModal({ open: false, title: modal.title, id: '', scene: modal.scene });
        }}
      />
    </PageContainer>
  );
};
