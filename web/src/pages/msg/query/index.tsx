import { FormatMsgAlert, MsgAlertAlertStatus, MsgAlertWhereInput, UserInfo } from '@/generated/msgsrv/graphql';
import { EnumMsgAlertStatus, getFormatMsgAlertList, getRenderMsgAlert } from '@/services/msgsrv/list';
import { ActionType, PageContainer, ProColumns, ProTable, useToken } from '@ant-design/pro-components';
import { OrgKind } from '@knockout-js/api/ucenter';
import { KeepAlive } from '@knockout-js/layout';
import { OrgSelect } from '@knockout-js/org';
import { getDate } from '@qeelyn-pb/ims-js/esm/utils';
import { Modal, Space, Typography } from 'antd';
import { definePageConfig, useNavigate } from 'ice';
import { useRef, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { TemplateType } from '@/pages/msg/event';

export default () => {
  const { token } = useToken(),
    navigate = useNavigate(),
    { t } = useTranslation(),
    // 表格相关
    proTableRef = useRef<ActionType>(),
    iframeRef = useRef<HTMLIFrameElement>(null),
    [modal, setModal] = useState<{
      show: boolean;
    }>({
      show: false,
    }),
    columns: ProColumns<FormatMsgAlert>[] = [
      {
        title: '消息事件', dataIndex: 'msgEventComments', width: 120, search: false,
      },
      {
        title: '接收方式',
        dataIndex: 'receiverType',
        width: 120,
        search: false,
        render: (text, record) => {
          return record.receiverType || '-';
        },
      },
      {
        title: '标题', dataIndex: 'msgTemplateTitle', width: 120, search: false,
      },
      {
        title: '开始时间',
        dataIndex: 'startsAt',
        width: 120,
        valueType: 'dateRange',
        order: 3,
        fieldProps: {
          format: 'YYYY-MM-DD',
        },
        render(text, record) {
          return (<>
            <div>{getDate(record.startsAt, 'YYYY-MM-DD HH:mm:ss')}</div>
          </>);
        },
      },
      {
        title: t('end_at'), dataIndex: 'endsAt', width: 120, valueType: 'dateTime', search: false,
      },
      {
        title: '接收用户',
        width: 120,
        search: false,
        render: (text, record) => {
          return record.users?.map((item: UserInfo) => {
            return item.name ? item.name : item.email;
          }).join(',') || '-';
        },
      },
      {
        title: '接收通道', dataIndex: 'msgChannelComments', width: 120, search: false,
      },
      {
        title: t('status'),
        dataIndex: 'state',
        width: 120,
        search: false,
        filters: true,
        valueEnum: EnumMsgAlertStatus,
      },
      {
        title: t('operation'),
        dataIndex: 'actions',
        fixed: 'right',
        align: 'center',
        search: false,
        width: 120,
        render: (text, record) => {
          return (
            <Space><a onClick={async () => {
            const result = await getRenderMsgAlert(record.id, record.receiver);
            setModal({ show: true });
            setTimeout(() => {
              if (iframeRef.current?.contentWindow) {
                iframeRef.current.contentWindow.document.write(`<pre>${result}</pre>`);
              } else if (iframeRef.current?.contentDocument) {
                iframeRef.current.contentDocument.write(`<pre>${result}</pre>`);
              }
            }, 200);
          }}
            >查看内容</a>
              <Typography.Link
                disabled={!record.hasMultiMsg}
                onClick={() => {
                  navigate(`/msg/query/more?id=${record.id}`);
                }}
              >
                更多消息
              </Typography.Link>
            </Space>
          );
        },
      },
    ];

  return (<KeepAlive clearAlive>
    <PageContainer
      header={{
        title: '消息查询',
        style: { background: token.colorBgContainer },
        breadcrumb: {
          items: [
            { title: t('msg_center') },
            { title: '消息查询' },
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
          title: '消息查询',
        }}
        scroll={{ x: 'max-content' }}
        columns={columns}
        request={async (params, sort, filter) => {
          const table = { data: [] as FormatMsgAlert[], success: true, total: 0 },
            where: MsgAlertWhereInput = {};
          if (params?.startsAt?.[0] && params?.startsAt?.[1]) {
            where.startsAtGTE = getDate(params.startsAt[0], 'YYYY-MM-DDT00:00:00Z');
            where.startsAtLTE = getDate(params.startsAt[1], 'YYYY-MM-DDT23:59:59Z');
          }
          const result = await getFormatMsgAlertList({
            current: params.current,
            pageSize: params.pageSize,
            where: where,
          });
          if (result?.totalCount) {
            table.data = result.edges?.map(item => item?.node) as FormatMsgAlert[];
            table.total = result.totalCount;
          }
          return table;
        }}

      />
      <Modal
        title="查看内容"
        open={modal.show}
        destroyOnClose
        footer={null}
        width={800}
        onCancel={() => {
          setModal({ show: false });
        }}
      >
        <iframe style={{ width: '100%', height: '60vh', border: '0 none' }} ref={iframeRef} />
      </Modal>
    </PageContainer>
  </KeepAlive>);
};

export const pageConfig = definePageConfig(() => ({
  auth: ['/msg/list'],
}));
