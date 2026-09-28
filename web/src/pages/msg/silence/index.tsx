import { ActionType, PageContainer, ProColumns, ProTable, useToken } from '@ant-design/pro-components';
import { Button, Space, Modal, Divider, Typography, Tag } from 'antd';
import { Key, useRef, useState } from 'react';
import { useTranslation } from 'react-i18next';
import Auth from '@/components/auth';
import { getOrgs } from '@knockout-js/api';
import Create from './components/create';
import { MsgSilence, MsgSilenceSilenceState, MsgSilenceWhereInput } from '@/generated/msgsrv/graphql';
import { EnumSilenceMatchType, EnumSilenceStatus, delSilence, getSilenceList } from '@/services/msgsrv/silence';
import { OrgSelect } from '@knockout-js/org';
import { Org, OrgKind } from '@knockout-js/api/ucenter';
import { KeepAlive } from '@knockout-js/layout';
import { definePageConfig } from 'ice';
import { delDataSource, saveDataSource, onMousedown } from '@/util';
import { useResizableProTable, routeBreadcrumb } from '@/util/hook';


const List = () => {
  const { token } = useToken(),
    { t } = useTranslation(),
    // 表格相关
    proTableRef = useRef<ActionType>(),
    [orgs, setOrgs] = useState<Org[]>([]),
    [dataSource, setDataSource] = useState<MsgSilence[]>([]),
    // 选中处理
    [selectedRowKeys, setSelectedRowKeys] = useState<Key[]>([]),
    // 弹出层处理
    [modal, setModal] = useState<{
      open: boolean;
      title: string;
      id: string;
      scene: 'editor' | 'copy';
    }>({
      open: false,
      title: '',
      id: '',
      scene: 'editor'
    }),
    // 可调整列宽的 ProTable
    { columns: finalColumns, components, tableWidth } = useResizableProTable<MsgSilence>(() => ({
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
            const org = orgs.find(item => item.id == `${record.tenantID}`)
            return record.tenantID ? org?.name || record.tenantID : '-';
          },
        },
        { title: t('starts_at'), dataIndex: 'startsAt', valueType: "dateTime", width: 160 },
        { title: t('end_at'), dataIndex: 'endsAt', valueType: 'dateTime', width: 160 },
        {
          title: t('match_msg'), dataIndex: 'matchers', width: 200, search: false, ellipsis: true,
          render: (text, record) => {
            return record.matchers?.map(item => {
              if (item) {
                return `${item.name}${EnumSilenceMatchType[item.type].text}"${item.value}"`;
              }
              return '';
            }).join(',') || '-';
          }
        },
        {
          title: t('status'), dataIndex: 'state', width: 120, align: 'center', search: false,
          filters: true,
          valueEnum: EnumSilenceStatus,
          render(_, record) {
            const enumItem = EnumSilenceStatus[record.state];
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
          width: 180,
          hideInSetting: true,
          render: (text, record) => {
            return (<Space split={<Divider type="vertical" className="ko-divider-gray" />} size={0}>
              <Auth authKey="updateMsgSilence">
                <Typography.Link
                  key="editor"
                  onClick={() => {
                    setModal({
                      open: true, title: `${t('edit')}:${record.id}`, id: record.id, scene: 'editor'
                    });
                  }}
                >
                  {t('edit')}
                </Typography.Link>
              </Auth>
              <Auth authKey="createMsgSilence">
                <Typography.Link
                  key="editor"
                  onClick={() => {
                    setModal({
                      open: true, title: `${t('copy')}:${record.id}`, id: record.id, scene: 'copy'
                    });
                  }}
                >
                  {t('copy')}
                </Typography.Link>
              </Auth>
              <Auth authKey="deleteMsgSilence">
                <Typography.Link key="delete" onClick={() => onDel(record)}>
                  {t('delete')}
                </Typography.Link>
              </Auth>
            </Space>);
          },
        },
      ],
    }), [orgs]);


  const
    onDel = (record: MsgSilence) => {
      Modal.confirm({
        title: t('delete'),
        content: `${t('confirm_delete')}：${record.id}`,
        onOk: async (close) => {
          const result = await delSilence(record.id);
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
            <Auth authKey="createMsgSilence">
              <Button
                key="created"
                type="primary"
                onClick={() => {
                  setModal({ open: true, title: t('create_silence_msg'), id: '', scene: 'editor' });
                }}
              >
                {t('create_silence_msg')}
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
          const table = { data: [] as MsgSilence[], success: true, total: 0 },
            where: MsgSilenceWhereInput = {};
          where.tenantID = params.org?.id;
          where.startsAt = params.startsAt
          where.endsAt = params.endsAt
          where.stateIn = filter.status as MsgSilenceSilenceState[]
          const result = await getSilenceList({
            current: params.current,
            pageSize: params.pageSize,
            where,
          });
          if (result?.totalCount) {
            table.data = result.edges?.map(item => item?.node) as MsgSilence[]
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
          onChange: (rowKeys) => setSelectedRowKeys(rowKeys as Key[]),
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
        isCopy={modal.scene === 'copy'}
        onClose={async (isSuccess, newInfo) => {
          if (isSuccess && newInfo) {
            if (!orgs.find(item => item.id == `${newInfo.tenantID}`)) {
              setOrgs([...orgs, ...(await getOrgs([newInfo.tenantID]))])
            }
            setDataSource(saveDataSource(dataSource, newInfo))
          }
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
  auth: ['/msg/silence'],
}));
