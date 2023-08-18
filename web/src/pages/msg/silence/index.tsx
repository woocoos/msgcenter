import { ActionType, PageContainer, ProColumns, ProTable, useToken } from '@ant-design/pro-components';
import { Button, Space, Modal } from 'antd';
import { useRef, useState } from 'react';
import { useTranslation } from 'react-i18next';
import Auth from '@/components/auth';
import { Org, getOrgs } from '@knockout-js/api';
import Create from './components/create';
import { Silence, SilenceSilenceState, SilenceWhereInput } from '@/generated/msgsrv/graphql';
import { EnumSilenceMatchType, EnumSilenceStatus, delSilence, getSilenceList } from '@/services/msgsrv/silence';
import { OrgSelect } from '@knockout-js/org';
import { OrgKind } from '@knockout-js/api';


export default () => {
  const { token } = useToken(),
    { t } = useTranslation(),
    // 表格相关
    proTableRef = useRef<ActionType>(),
    columns: ProColumns<Silence>[] = [
      // 有需要排序配置  sorter: true
      {
        title: t('org'), dataIndex: 'org', width: 120,
        renderFormItem: () => {
          return <OrgSelect kind={OrgKind.Root} />
        },
        render: (text, record) => {
          const org = orgs.find(item => item.id == `${record.tenantID}`)
          return record.tenantID ? org?.name || record.tenantID : '-';
        },
      },
      { title: t('starts_at'), dataIndex: 'startsAt', valueType: "dateTime", width: 120 },
      { title: t('end_at'), dataIndex: 'endsAt', valueType: 'dateTime', width: 120 },
      {
        title: t('match_msg'), dataIndex: 'matchers', width: 120, search: false,
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
        title: t('status'), dataIndex: 'state', width: 120, search: false,
        filters: true,
        valueEnum: EnumSilenceStatus,
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
            <Auth authKey="updateSilence">
              <a
                key="editor"
                onClick={() => {
                  setModal({
                    open: true, title: `${t('edit')}:${record.id}`, id: record.id, scene: 'editor'
                  });
                }}
              >
                {t('edit')}
              </a>
            </Auth>
            <Auth authKey="createSilence">
              <a
                key="editor"
                onClick={() => {
                  setModal({
                    open: true, title: `${t('copy')}:${record.id}`, id: record.id, scene: 'copy'
                  });
                }}
              >
                {t('copy')}
              </a>
            </Auth>
            <Auth authKey="deleteSilence">
              <a key="delete" onClick={() => onDel(record)}>
                {t('delete')}
              </a>
            </Auth>
          </Space>);
        },
      },
    ],
    [orgs, setOrgs] = useState<Org[]>([]),
    [dataSource, setDataSource] = useState<Silence[]>([]),
    // 选中处理
    [selectedRowKeys, setSelectedRowKeys] = useState<string[]>([]),
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
    });


  const
    onDel = (record: Silence) => {
      Modal.confirm({
        title: t('delete'),
        content: `${t('confirm_delete')}：${record.id}`,
        onOk: async (close) => {
          const result = await delSilence(record.id);
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
        title: t('silence_msg'),
        style: { background: token.colorBgContainer },
        breadcrumb: {
          items: [
            { title: t('msg_center') },
            { title: t('silence_msg') },
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
          title: t('silence_msg_list'),
          actions: [
            <Auth authKey="createSilence">
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
        scroll={{ x: 'max-content' }}
        columns={columns}
        request={async (params, sort, filter) => {
          const table = { data: [] as Silence[], success: true, total: 0 },
            where: SilenceWhereInput = {};
          where.tenantID = params.org?.id;
          where.startsAt = params.startsAt
          where.endsAt = params.endsAt
          where.stateIn = filter.status as SilenceSilenceState[]
          const result = await getSilenceList({
            current: params.current,
            pageSize: params.pageSize,
            where,
          });
          if (result?.totalCount) {
            table.data = result.edges?.map(item => item?.node) as Silence[]
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
        open={modal.open}
        title={modal.title}
        id={modal.id}
        isCopy={modal.scene === 'copy'}
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
