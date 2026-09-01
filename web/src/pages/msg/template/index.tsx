import { ActionType, PageContainer, ProColumns, ProTable, useToken } from '@ant-design/pro-components';
import { Button, Space, Modal, Dropdown } from 'antd';
import { Key, useRef, useState } from 'react';
import { useTranslation } from 'react-i18next';
import Auth from '@/components/auth';
import { Link, useSearchParams } from '@ice/runtime';
import { MsgEvent, MsgTemplate, MsgTemplateFormat, MsgTemplateReceiverType, MsgTemplateSimpleStatus, MsgTemplateWhereInput } from '@/generated/msgsrv/graphql';
import { EnumMsgTemplateFormat, EnumMsgTemplateReceiverType, EnumMsgTemplateStatus, delMsgTemplate, disableMsgTemplate, enableMsgTemplate, getMsgTemplateList } from '@/services/msgsrv/template';
import { getMsgEventInfo } from '@/services/msgsrv/event';
import { DownOutlined } from '@ant-design/icons';
import Create from './components/create';
import { getOrgs } from '@knockout-js/api';
import { Org } from '@knockout-js/api/ucenter';
import Test from './components/test';
import { delDataSource, saveDataSource } from '@/util';
import { TemplateType } from '../event'
import store from '@/store'

export default () => {
  const { token } = useToken(),
    { t } = useTranslation(),
    userState = store.getModelState('user'),
    [searchParams] = useSearchParams(),
    [msgEventInfo, setMsgEventInfo] = useState<MsgEvent>(),
    // 表格相关
    proTableRef = useRef<ActionType>(),
    columns: ProColumns<MsgTemplate>[] = [
      // 有需要排序配置  sorter: true
      // {
      //   title: t('org'), dataIndex: 'org', width: 120,
      //   render: (text, record) => {
      //     const org = orgs.find(item => item.id == record.tenantID)
      //     return record.tenantID ? org?.name || record.tenantID : '';
      //   },
      // },
      {
        title: t('name'), dataIndex: 'name', width: 120,
      },
      { title: t('subject'), dataIndex: 'subject', width: 120 },
      {
        title: '模板范围', dataIndex: 'userID', width: 100, search: false,
        render(text, record) {
          if (record.userID && record.userID != "0") {
            return `用户(${record.userID})`;
          }
          if (record.tenantID && record.tenantID != "0") {
            return '租户';
          }
          return '全局';
        },
      },
      {
        title: t('way_receiving'), dataIndex: 'receiverType', width: 120, search: false,
        filters: true,
        valueEnum: EnumMsgTemplateFormat,
      },
      {
        title: t('status'), dataIndex: 'status', width: 120, search: false,
        filters: true,
        valueEnum: EnumMsgTemplateStatus,
      },
      { title: t('description'), dataIndex: 'comments', width: 120, search: false },
      {
        title: t('operation'),
        dataIndex: 'actions',
        fixed: 'right',
        align: 'center',
        search: false,
        width: 160,
        render: (text, record) => {
          return (<Space>
            {
              record.status === MsgTemplateSimpleStatus.Active ? <a
                key="view"
                onClick={() => {
                  setModal({
                    open: true,
                    title: `${t('view')}:${record.name}`,
                    id: record.id,
                    receiverType: record.receiverType,
                    readonly: true,
                  });
                }}
              >
                {t('view')}
              </a> : <Auth authKey="updateMsgTemplate">
                <a
                  key="editor"
                  onClick={() => {
                    setModal({
                      open: true,
                      title: `${t('edit')}:${record.name}`,
                      id: record.id,
                      receiverType: record.receiverType,
                    });
                  }}
                >
                  {t('edit')}
                </a>
              </Auth>
            }
            {
              record.status === MsgTemplateSimpleStatus.Active ? <></> : <Auth authKey="deleteMsgTemplate">
                <a key="delete" onClick={() => onDel(record)}>
                  {t('delete')}
                </a>
              </Auth>
            }
            {
              record.status === MsgTemplateSimpleStatus.Active ? <Auth authKey="disableMsgTemplate">
                <a key="disable" style={{ color: '#ff0000' }} onClick={() => onClickStatus(record)}>
                  {t('disable')}
                </a>
              </Auth> : <Auth authKey="enableMsgTemplate">
                <a key="enable" onClick={() => onClickStatus(record)}>
                  {t('enable')}
                </a>
              </Auth>
            }
            <Auth authKey={['testSendMessageTpl', 'testSendEmailTpl']} keyAndOr='or'>
              {
                record.status === MsgTemplateSimpleStatus.Active && record.receiverType != MsgTemplateReceiverType.Webhook ? <a
                  onClick={() => {
                    setModal({
                      open: true,
                      title: `${t('test')}:${record.name}`,
                      id: record.id,
                      type: 'test',
                    });
                  }}
                >{t('test')}</a> : <></>
              }
            </Auth>
          </Space>);
        },
      },
    ],
    [orgs, setOrgs] = useState<Org[]>([]),
    [dataSource, setDataSource] = useState<MsgTemplate[]>([]),
    // 选中处理
    [selectedRowKeys, setSelectedRowKeys] = useState<Key[]>([]),
    // 弹出层处理
    [modal, setModal] = useState<{
      open: boolean;
      title: string;
      id: string;
      receiverType?: MsgTemplateReceiverType;
      type?: 'test';
      readonly?: boolean;
    }>({
      open: false,
      title: '',
      id: '',
    });


  const
    getMsgEvent = async () => {
      const msgEventId = searchParams.get('id');
      if (msgEventId) {
        const result = await getMsgEventInfo(msgEventId);
        if (result?.id) {
          setMsgEventInfo(result as MsgEvent)
          return result
        }
      }
      return null;
    },
    onDel = (record: MsgTemplate) => {
      Modal.confirm({
        title: t('delete'),
        content: `${t('confirm_delete')}：${record.name}`,
        onOk: async (close) => {
          const result = await delMsgTemplate(record.id);
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
    onClickStatus = (record: MsgTemplate) => {
      Modal.confirm({
        title: record.status === MsgTemplateSimpleStatus.Active ? t('disable') : t('enable'),
        content: `${record.status === MsgTemplateSimpleStatus.Active ? t('disable') : t('enable')}：${record.name}`,
        onOk: async (close) => {
          const result = record.status === MsgTemplateSimpleStatus.Active ? await disableMsgTemplate(record.id) : await enableMsgTemplate(record.id);
          if (result?.id) {
            setDataSource(saveDataSource(dataSource, result as MsgTemplate))
            close();
          }
        },
      });
    };


  return (
    <PageContainer
      header={{
        title: searchParams.get('type') === TemplateType.customer ? t('temp_customer') : t('temp_default'),
        style: { background: token.colorBgContainer },
        breadcrumb: {
          items: [
            { title: t('msg_center') },
            { title: <Link to={'/msg/event'}>{t('msg_event')}</Link> },
            { title: searchParams.get('type') === TemplateType.customer ? t('temp_customer') : t('temp_default') },
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
          title: `${t('msg_event')}:${msgEventInfo?.name}`,
          actions: [
            <Auth authKey="createMsgTemplate">
              <Dropdown menu={{
                items: Object.keys(EnumMsgTemplateReceiverType).map(item => ({
                  key: item,
                  label: item,
                  onClick: ({ key }) => {
                    setModal({ open: true, title: t('create_msg_template'), id: '', receiverType: key as MsgTemplateReceiverType });
                  }
                }))
              }}>
                <Button
                  key="created"
                  type="primary"
                >
                  {t('create_msg_template')}<DownOutlined />
                </Button>
              </Dropdown>
            </Auth>,
          ],
        }}
        scroll={{ x: 'max-content' }}
        columns={columns}
        dataSource={dataSource}
        request={async (params, sort, filter) => {
          const table = { data: [] as MsgTemplate[], success: true, total: 0 },
            where: MsgTemplateWhereInput = {};
          const msgEvent = msgEventInfo?.id ? msgEventInfo : await getMsgEvent();
          if (msgEvent?.id) {
            where.msgEventID = msgEvent.id
            if (searchParams.get('type') == TemplateType.customer) {
              where.tenantID = userState.tenantId
            } else {
              where.tenantIDIsNil = true
            }

            where.nameContains = params.name;
            where.subjectContains = params.subject;
            where.receiverTypeIn = filter.modes as MsgTemplateReceiverType[]
            where.statusIn = filter.status as MsgTemplateSimpleStatus[]
            const result = await getMsgTemplateList({
              current: params.current,
              pageSize: params.pageSize,
              where,
            });
            if (result?.totalCount) {
              table.data = result.edges?.map(item => item?.node) as MsgTemplate[];
              // setOrgs(await getOrgs(table.data.map(item => item.tenantID)))
              table.total = result.totalCount;
            }
          }
          setSelectedRowKeys([]);
          setDataSource(table.data);
          return table;
        }}
        rowSelection={{
          selectedRowKeys: selectedRowKeys,
          onChange: (selectedRowKeys) => { setSelectedRowKeys(selectedRowKeys); },
          type: 'checkbox',
        }}
      />
      {msgEventInfo && modal.type != 'test' ?
        <Create
          open={modal.open}
          title={modal.title}
          type={searchParams.get('type')}
          id={modal.id}
          readonly={modal.readonly}
          onClose={(isSuccess, newInfo) => {
            if (isSuccess && newInfo) {
              setDataSource(saveDataSource(dataSource, newInfo))
            }
            setModal({ open: false, title: modal.title, id: '', receiverType: modal.receiverType, readonly: false });
          }}
          msgEvent={msgEventInfo}
          receiverType={modal.receiverType || MsgTemplateReceiverType.Email}
        /> : <></>
      }
      {
        modal.type === 'test' ? <Test
          open={modal.open}
          title={modal.title}
          id={modal.id}
          onClose={() => {
            setModal({ open: false, title: modal.title, id: '' });
          }}
        /> : <></>
      }

    </PageContainer>
  );
};
