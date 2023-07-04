import { ActionType, PageContainer, ProColumns, ProTable, useToken } from '@ant-design/pro-components';
import { Button, Space, Modal } from 'antd';
import { useRef, useState } from 'react';
import { TableFilter, TableParams, TableSort } from '@/services/graphql';
import { useTranslation } from 'react-i18next';
import Auth from '@/components/Auth';
import { MsgType, MsgTypeSimpleStatus, MsgTypeWhereInput } from '@/__generated__/msgsrv/graphql';
import { EnumMsgTypeStatus, delMsgType, getMsgTypeList } from '@/services/msgsrv/type';
import InputApp from '@/components/Adminx/App/input';


export default () => {
  const { token } = useToken(),
    { t } = useTranslation(),
    // 表格相关
    proTableRef = useRef<ActionType>(),
    columns: ProColumns<MsgType>[] = [
      // 有需要排序配置  sorter: true
      {
        title: t('app'), dataIndex: 'app', width: 120,
        renderFormItem() {
          return <InputApp />
        },
        render: (text, record) => {
          return record.appID;
        },
      },
      { title: t('category'), dataIndex: 'category', width: 120 },
      { title: t('name'), dataIndex: 'name', width: 120 },
      {
        title: t('open_subscription'),
        dataIndex: 'canSubs',
        width: 120,
        search: false,
        render: (text, record) => {
          return record.canSubs ? t('yes') : t('no');
        },
      },
      {
        title: t('open_custom'),
        dataIndex: 'canCustom',
        width: 120,
        search: false,
        render: (text, record) => {
          return record.canCustom ? t('yes') : t('no');
        },
      },
      {
        title: t('status'), dataIndex: 'status', width: 120, search: false,
        filters: true,
        valueEnum: EnumMsgTypeStatus,
      },
      { title: t('description'), dataIndex: 'comments', width: 120, search: false },
      {
        title: t('operation'),
        dataIndex: 'actions',
        fixed: 'right',
        align: 'center',
        search: false,
        width: 120,
        render: (text, record) => {
          return (<Space>
            <Auth authKey="updateMsgType">
              <a
                key="editor"
                onClick={() => {
                  setModal({
                    open: true, title: `${t('edit')}:${record.name}`, id: record.id
                  });
                }}
              >
                {t('edit')}
              </a>
            </Auth>
            <Auth authKey="deleteMsgType">
              <a key="delete" onClick={() => onDel(record)}>
                {t('delete')}
              </a>
            </Auth>
          </Space>);
        },
      },
    ],
    [dataSource, setDataSource] = useState<MsgType[]>([]),
    // 选中处理
    [selectedRowKeys, setSelectedRowKeys] = useState<string[]>([]),
    // 弹出层处理
    [modal, setModal] = useState<{
      open: boolean;
      title: string;
      id: string;
    }>({
      open: false,
      title: '',
      id: '',
    });


  const
    getRequest = async (params: TableParams, sort: TableSort, filter: TableFilter) => {
      const table = { data: [] as MsgType[], success: true, total: 0 },
        where: MsgTypeWhereInput = {};
      where.appID = params.app?.id;
      where.category = params.category;
      where.nameContains = params.name;
      where.statusIn = filter.status as MsgTypeSimpleStatus[]
      const result = await getMsgTypeList({
        current: params.current,
        pageSize: params.pageSize,
        where,
      });
      if (result?.totalCount) {
        table.data = result.edges?.map(item => item?.node) as MsgType[]
        table.total = result.totalCount;
      }
      setSelectedRowKeys([]);
      setDataSource(table.data);
      return table;
    },
    onDel = (record: MsgType) => {
      Modal.confirm({
        title: t('delete'),
        content: `${t('confirm_delete')}：${record.name}`,
        onOk: async (close) => {
          const result = await delMsgType(record.id);
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
    };


  return (
    <PageContainer
      header={{
        title: t('msg_type'),
        style: { background: token.colorBgContainer },
        breadcrumb: {
          items: [
            { title: t('msg_center') },
            { title: t('msg_type') },
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
          title: t('msg_type_list'),
          actions: [
            <Auth authKey="createMsgType">
              <Button
                key="created"
                type="primary"
                onClick={() => {
                  setModal({ open: true, title: t('create_msg_type'), id: '' });
                }}
              >
                {t('create_msg_type')}
              </Button>
            </Auth>,
          ],
        }}
        scroll={{ x: 'max-content' }}
        columns={columns}
        request={getRequest}
        rowSelection={{
          selectedRowKeys: selectedRowKeys,
          onChange: (selectedRowKeys: string[]) => { setSelectedRowKeys(selectedRowKeys); },
          type: 'checkbox',
        }}
      />
    </PageContainer>
  );
};