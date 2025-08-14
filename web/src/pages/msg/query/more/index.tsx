import { FormatMsgAlert, MsgAlertAlertStatus, MsgAlertWhereInput, UserInfo } from '@/generated/msgsrv/graphql';
import { EnumMsgAlertStatus, getFormatMsgAlertMore, getRenderMsgAlert } from '@/services/msgsrv/list';
import { ActionType, PageContainer, ProColumns, ProTable, useToken } from '@ant-design/pro-components';
import { Link, useSearchParams } from '@ice/runtime';
import { KeepAlive } from '@knockout-js/layout';
import { Modal, Space } from 'antd';
import { definePageConfig } from 'ice';
import { useRef, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { TemplateType } from '@/pages/msg/event';

export default () => {
  const { token } = useToken(),
    { t } = useTranslation(),
    // 表格相关
    proTableRef = useRef<ActionType>(),
    [searchParams] = useSearchParams(),
    iframeRef = useRef<HTMLIFrameElement>(null),
    [msgEventComments, setMsgEventComments] = useState<string>(''),
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
        title: t('starts_at'), dataIndex: 'startsAt', width: 120, valueType: 'dateTime', search: false,
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
            </Space>
          );
        },
      },
    ];

  return (<KeepAlive clearAlive>
    <PageContainer
      header={{
        title: '更多消息',
        style: { background: token.colorBgContainer },
        breadcrumb: {
          items: [
            { title: t('msg_center') },
            { title: <Link to={'/msg/query'}>消息查询</Link> },
            { title: '更多消息' },
          ],
        },
      }}
    >
      <ProTable
        actionRef={proTableRef}
        search={false}
        rowKey={'id'}
        toolbar={{
          title: `消息事件：${msgEventComments}`,
        }}
        scroll={{ x: 'max-content' }}
        columns={columns}
        request={async (params, sort, filter) => {
          const table = { data: [] as FormatMsgAlert[], success: true, total: 0 };
          const msgAlertId = searchParams.get('id');
          const result = await getFormatMsgAlertMore(
            msgAlertId,
          );
          if (result && result.length > 0) {
            table.data = result as FormatMsgAlert[];
            table.total = result.length;
            // 消息事件名称
            if (result[0]) {
              setMsgEventComments(result[0].msgEventComments || '');
            }
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
