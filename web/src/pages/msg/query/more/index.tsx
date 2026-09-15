import { FormatMsgAlert, MsgAlertAlertStatus, MsgAlertWhereInput, UserInfo } from '@/generated/msgsrv/graphql';
import { EnumMsgAlertStatus, getFormatMsgAlertMore, getRenderMsgAlert } from '@/services/msgsrv/list';
import { ActionType, PageContainer, ProTable } from '@ant-design/pro-components';
import { Link, useSearchParams } from '@ice/runtime';
import { KeepAlive } from '@knockout-js/layout';
import { Modal, Space, Typography, Divider } from 'antd';
import { useRef, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { useResizableProTable, routeBreadcrumb } from '@/util/hook';

const List = () => {
  const { t } = useTranslation(),
    [searchParams] = useSearchParams(),
    // 表格相关
    proTableRef = useRef<ActionType>(),
    iframeRef = useRef<HTMLIFrameElement>(null),
    [msgEventComments, setMsgEventComments] = useState<string>(''),
    [modal, setModal] = useState<{
      show: boolean;
    }>({
      show: false,
    }),
    // 可调整列宽的 ProTable
    { columns: finalColumns, components, tableWidth } = useResizableProTable<FormatMsgAlert>(() => ({
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
          title: t('msg_event'), dataIndex: 'msgEventComments', width: 200, search: false, ellipsis: true,
        },
        {
          title: t('receiving_type'),
          dataIndex: 'receiverType',
          width: 120,
          align: 'center',
          search: false,
          render: (text, record) => {
            return record.receiverType || '-';
          },
        },
        {
          title: t('subject'), dataIndex: 'msgTemplateTitle', width: 200, search: false, ellipsis: true,
        },
        {
          title: t('starts_at'), dataIndex: 'startsAt', width: 160, valueType: 'dateTime', search: false,
        },
        {
          title: t('end_at'), dataIndex: 'endsAt', width: 160, valueType: 'dateTime', search: false,
        },
        {
          title: t('receiving_user'),
          width: 160,
          search: false,
          ellipsis: true,
          render: (text, record) => {
            return record.users?.map((item) => {
              return item?.name ? item.name : item?.email;
            }).join(',') || '-';
          },
        },
        {
          title: t('receive_channel'), dataIndex: 'msgChannelComments', width: 200, search: false, ellipsis: true,
        },
        {
          title: t('status'),
          dataIndex: 'state',
          width: 120,
          align: 'center',
          search: false,
          filters: true,
          valueEnum: EnumMsgAlertStatus,
        },
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
            return (
              <Space split={<Divider type="vertical" className="ko-divider-gray" />} size={0}>
                <Typography.Link onClick={async () => {
                  const result = await getRenderMsgAlert(record.id, record.receiver);
                  setModal({ show: true });
                  setTimeout(() => {
                    if (iframeRef.current?.contentWindow) {
                      iframeRef.current.contentWindow.document.write(`<pre>${result}</pre>`);
                    } else if (iframeRef.current?.contentDocument) {
                      iframeRef.current.contentDocument.write(`<pre>${result}</pre>`);
                    }
                  }, 200);
                }}>
                  {t('view_content')}
                </Typography.Link>
              </Space>
            );
          },
        },
      ],
    }), []);

  return (
    <>
      <ProTable
        actionRef={proTableRef}
        sticky={{ offsetHeader: 56 }}
        search={false}
        rowKey={'id'}
        toolbar={{
          title: `${t('msg_event')}：${msgEventComments}`,
        }}
        scroll={{ x: tableWidth }}
        components={components}
        columns={finalColumns}
        columnsState={{ defaultValue: { id: { show: false } } }}
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
        title={t('view_content')}
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
    </>
  );
};

export default () => {
  return (
    <PageContainer
      className="ko-page-container"
    >
      <div className="ka-content">
        <KeepAlive clearAlive>
          <List />
        </KeepAlive>
      </div>
    </PageContainer>
  );
};
