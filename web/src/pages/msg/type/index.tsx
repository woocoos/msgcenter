import { ActionType, PageContainer, ProColumns, ProTable, useToken } from '@ant-design/pro-components';
import { Button, Space, Modal, Divider, Typography, Tag } from 'antd';
import { useRef, useState } from 'react';
import { useTranslation } from 'react-i18next';
import Auth from '@/components/auth';
import { MsgType, MsgTypeSimpleStatus, MsgTypeWhereInput } from '@/generated/msgsrv/graphql';
import { EnumMsgTypeStatus, delMsgType, getMsgTypeList } from '@/services/msgsrv/type';
import Create from './components/create';
import { AppSelect } from '@knockout-js/org';
import { getApps } from '@knockout-js/api';
import { App } from '@knockout-js/api/ucenter';
import { KeepAlive } from '@knockout-js/layout';
import { DictSelect, DictText } from '@knockout-js/org';
import { definePageConfig } from 'ice';
import { delDataSource, saveDataSource, onMousedown } from '@/util';
import { useResizableProTable, routeBreadcrumb } from '@/util/hook';

const List = () => {
  const { token } = useToken(),
    { t } = useTranslation(),
    // 表格相关
    proTableRef = useRef<ActionType>(),
    [apps, setApps] = useState<App[]>([]),
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
    }),
    // 可调整列宽的 ProTable
    { columns: finalColumns, components, tableWidth } = useResizableProTable<MsgType>(() => ({
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
          title: t('app'), dataIndex: 'app', width: 160,
          renderFormItem() {
            return <AppSelect />
          },
          render: (text, record) => {
            const app = apps.find(item => item.id == record.appID)
            return record.appID ? app?.name || record.appID : '-';
          },
        },
        {
          title: t('category'), dataIndex: 'category', width: 140,
          renderText(text, record, index, action) {
            return <DictText dictCode="MsgCategory" value={record.category} />
          },
          renderFormItem() {
            return <DictSelect dictCode="MsgCategory" placeholder={t('please_enter_category')} />
          },
        },
        { title: t('name'), dataIndex: 'name', width: 160 },
        {
          title: t('open_subscription'),
          dataIndex: 'canSubs',
          width: 120,
          search: false,
          align: 'center',
          render: (text, record) => {
            return record.canSubs ? t('yes') : t('no');
          },
        },
        {
          title: t('open_custom'),
          dataIndex: 'canCustom',
          width: 120,
          search: false,
          align: 'center',
          render: (text, record) => {
            return record.canCustom ? t('yes') : t('no');
          },
        },
        {
          title: t('status'), dataIndex: 'status', width: 120, align: 'center', search: false,
          filters: true,
          valueEnum: EnumMsgTypeStatus,
          render(_, record) {
            const enumItem = record.status ? EnumMsgTypeStatus[record.status] : undefined;
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
          width: 120,
          hideInSetting: true,
          render: (text, record) => {
            return (<Space split={<Divider type="vertical" className="ko-divider-gray" />} size={0}>
              <Auth authKey="updateMsgType">
                <Typography.Link
                  key="editor"
                  onClick={() => {
                    setModal({
                      open: true, title: `${t('edit')}:${record.name}`, id: record.id
                    });
                  }}
                >
                  {t('edit')}
                </Typography.Link>
              </Auth>
              <Auth authKey="deleteMsgType">
                <Typography.Link key="delete" onClick={() => onDel(record)}>
                  {t('delete')}
                </Typography.Link>
              </Auth>
            </Space>);
          },
        },
      ],
    }), [apps]);


  const
    onDel = (record: MsgType) => {
      Modal.confirm({
        title: t('delete'),
        content: `${t('confirm_delete')}：${record.name}`,
        onOk: async (close) => {
          const result = await delMsgType(record.id);
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
        scroll={{ x: tableWidth }}
        components={components}
        columns={finalColumns}
        columnsState={{ defaultValue: { id: { show: false } } }}
        dataSource={dataSource}
        request={async (params, sort, filter) => {
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
            setApps(await getApps(table.data.map(item => item.appID || '')))
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
        open={modal.open}
        title={modal.title}
        id={modal.id}
        onClose={async (isSuccess, newInfo) => {
          if (isSuccess && newInfo) {
            if (newInfo.appID && !apps.find(item => item.id === newInfo.appID)) {
              setApps([...apps, ...(await getApps([newInfo.appID]))])
            }
            setDataSource(saveDataSource(dataSource, newInfo))
          }
          setModal({ open: false, title: modal.title, id: '' });
        }} />
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
  auth: ['/msg/type'],
}));
