import { ActionType, PageContainer, ProColumns, ProTable, useToken } from '@ant-design/pro-components';
import {Button, Space, Modal, message, Divider, Typography} from 'antd';
import { useRef, useState } from 'react';
import { useTranslation } from 'react-i18next';
import Auth from '@/components/auth';
import { MsgEvent, MsgEventSimpleStatus, MsgEventWhereInput } from '@/generated/msgsrv/graphql';
import { EnumMsgEventStatus, delMsgEvent, disableMsgEvent, enableMsgEvent, getMsgEventList } from '@/services/msgsrv/event';
import { refreshTemplateParams } from '@/services/msgsrv/template';
import Create from './components/create';
import { Link } from '@ice/runtime';
import Config from './components/config';
import ConfigExample from './components/configExample';
import { KeepAlive } from '@knockout-js/layout';
import { DictSelect, DictText } from '@knockout-js/org';
import { definePageConfig } from 'ice';
import { delDataSource, saveDataSource, onMousedown } from '@/util';
import { useResizableProTable, routeBreadcrumb } from '@/util/hook';

export const TemplateType = {
  customer: 'customer',
  default: 'default',
};

const List = () => {
  const { token } = useToken(),
    { t } = useTranslation(),
    // 表格相关
    proTableRef = useRef<ActionType>(),
    [dataSource, setDataSource] = useState<MsgEvent[]>([]),
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
    [loadingTempParams,setLoadingTempParams] = useState(false),
    // 可调整列宽的 ProTable
    { columns: finalColumns, components, tableWidth } = useResizableProTable<MsgEvent>(() => ({
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
          title: t('msg_type_category'), dataIndex: 'msgTypeCategory', width: 140,
          renderFormItem() {
            return <DictSelect dictCode="MsgCategory" placeholder={t('please_enter_category')} />
          },
          render(text, record) {
            return <DictText dictCode="MsgCategory" value={record.msgType?.category} />
          },
        },
        {
          title: t('msg_type_name'), dataIndex: 'msgTypeName', width: 160,
          render(text, record) {
            return record.msgType?.name ?? ''
          },
        },
        { title: t('msg_event_name'), dataIndex: 'name', width: 160 },
        {
          title: t('way_receiving'), dataIndex: 'modes', width: 140, search: false,
          render(text, record) {
            return record.modes?.split(',')?.join('、')
          },
        },
        {
          title: t('open_subscription'), dataIndex: 'canSubs', width: 120, search: false,
          align: 'center',
          valueEnum: {
            true: { text: t('yes'), status: 'Success' },
            false: { text: t('no'), status: 'Default' },
          },
        },
        {
          title: t('status'), dataIndex: 'status', width: 120, align: 'center', search: false,
          filters: true,
          valueEnum: EnumMsgEventStatus,
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
          width: 340,
          hideInSetting: true,
          render: (text, record) => {
            return (<Space split={<Divider type="vertical" className="ko-divider-gray" />} size={0}>
              <Auth authKey="updateMsgEvent">
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
              <Link
                key="template_customer"
                to={`/msg/template?id=${record.id}&type=${TemplateType.customer}`}
              >
                {t('temp_customer')}
              </Link>
              <Link
                key="template_default"
                to={`/msg/template?id=${record.id}&type=${TemplateType.default}`}
              >
                {t('temp_default')}
              </Link>
              <Auth authKey="updateMsgEvent">
                <Typography.Link
                  key="config"
                  onClick={() => {
                    setModal({
                      open: true, title: `${t('amend_msg_event_config')}:${record.name}`, id: record.id, scene: 'config'
                    });
                  }}
                >
                  {t('configuration')}
                </Typography.Link>
              </Auth>
              {
                record.status === MsgEventSimpleStatus.Active ? <></> : <Auth authKey="deleteMsgEvent">
                  <Typography.Link key="delete" onClick={() => onDel(record)}>
                    {t('delete')}
                  </Typography.Link>
                </Auth>
              }
              {
                record.status === MsgEventSimpleStatus.Active ? <Auth authKey="disableMsgEvent">
                  <Typography.Link key="disable" style={{ color: '#ff0000' }} onClick={() => onClickStatus(record)}>
                    {t('disable')}
                  </Typography.Link>
                </Auth> : <Auth authKey="enableMsgEvent">
                  <Typography.Link key="enable" onClick={() => onClickStatus(record)}>
                    {t('enable')}
                  </Typography.Link>
                </Auth>
              }
            </Space>);
          },
        },
      ],
    }), []);

  const
    onDel = (record: MsgEvent) => {
      Modal.confirm({
        title: t('delete'),
        content: `${t('confirm_delete')}：${record.name}`,
        onOk: async (close) => {
          const result = await delMsgEvent(record.id);
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
    onClickStatus = (record: MsgEvent) => {
      Modal.confirm({
        title: record.status === MsgEventSimpleStatus.Active ? t('disable') : t('enable'),
        content: `${record.status === MsgEventSimpleStatus.Active ? t('disable') : t('enable')}：${record.name}`,
        onOk: async (close) => {
          const result = record.status === MsgEventSimpleStatus.Active ? await disableMsgEvent(record.id) : await enableMsgEvent(record.id);
          if (result?.id) {
            setDataSource(saveDataSource(dataSource, result as MsgEvent))
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
          labelWidth: 96,
        }}
        rowKey={'id'}
        toolbar={{
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
        request={async (params, sort, filter) => {
          const table = { data: [] as MsgEvent[], success: true, total: 0 },
            where: MsgEventWhereInput = {};
          where.nameContains = params.name;
          if (params.msgTypeName || params.msgTypeCategory) {
            where.hasMsgTypeWith = [{
              nameContains: params.msgTypeName,
              categoryContains: params.msgTypeCategory,
            }];
          }
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
        onClose={(isSuccess, newInfo) => {
          if (isSuccess && newInfo) {
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
  const { t } = useTranslation(),
    [breadcrumbNames] = routeBreadcrumb();
  const [loadingTempParams,setLoadingTempParams] = useState(false);
  return (
    <PageContainer
      className="ko-page-container"
      header={{
        breadcrumb: {
          items: breadcrumbNames.map(item => ({ title: item })),
        },
        extra: <Auth authKey={'refreshTemplateParams'}>
          <Button size="middle" loading={loadingTempParams} onClick={ async () => {
            setLoadingTempParams(true )
            let result = await refreshTemplateParams();
            if (result){
              message.success(t('submit_success'));
            }
            setLoadingTempParams(false)
          }}>{t('template_params_refresh')}</Button>
        </Auth>,
      }}
    >
      <KeepAlive clearAlive>
        <List />
      </KeepAlive>
    </PageContainer>
  );
};


export const pageConfig = definePageConfig(() => ({
  auth: ['/msg/event'],
}));
