import { FormatMsgAlert, MsgAlertWhereInput } from '@/generated/msgsrv/graphql';
import {
  EnumMsgAlertStatus,
  EnumNlogReceiverType,
  getFormatMsgAlertList,
  getRenderMsgAlert,
} from '@/services/msgsrv/list';
import { ActionType, PageContainer, ProTable } from '@ant-design/pro-components';
import { KeepAlive, Modal } from '@knockout-js/layout';
import { UserSelect } from '@knockout-js/org';
import { Space, Typography, Divider } from 'antd';
import { definePageConfig, useNavigate } from 'ice';
import { useRef, useState } from 'react';
import { useTranslation } from 'react-i18next';
import InputMsgEvent from '@/pages/msg/query/components/inputMsgEvent';
import { getDate } from '@/util';
import { useResizableProTable, routeBreadcrumb } from '@/util/hook';

const List = () => {
  const navigate = useNavigate(),
    { t } = useTranslation(),
    // 表格相关
    proTableRef = useRef<ActionType>(),
    iframeRef = useRef<HTMLIFrameElement>(null),
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
          title: t('msg_event'),
          dataIndex: 'alertName',
          width: 120,
          order: 4,
          renderFormItem: () => {
            return <InputMsgEvent />;
          },
          render: (text, record) => {
            return record.msgEventComments || '-';
          },
        },
        {
          title: t('receiving_type'),
          dataIndex: 'receiverType',
          width: 120,
          valueType: 'select',
          valueEnum: EnumNlogReceiverType,
          order: 3,
          render: (text, record) => {
            return record.receiverType || '-';
          },
        },
        {
          title: t('subject'), dataIndex: 'msgTemplateTitle', width: 200, search: false,
        },
        {
          title: t('send_at'),
          dataIndex: 'startsAt',
          width: 160,
          valueType: 'dateRange',
          order: 1,
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
          title: t('end_at'), dataIndex: 'endsAt', width: 160, valueType: 'dateTime', search: false,
        },
        {
          title: t('receiving_user'),
          dataIndex: 'user',
          width: 160,
          order: 2,
          renderFormItem: () => {
            return <UserSelect />;
          },
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
          width: 160,
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
                <Typography.Link
                  disabled={!record.hasMultiMsg}
                  onClick={() => {
                    navigate(`/msg/query/more?id=${record.id}`);
                  }}
                >
                  {t('more_message')}
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
        search={{
          className: 'ko-pro-table-search',
          searchText: `${t('query')}`,
          resetText: `${t('reset')}`,
          labelWidth: 70,
        }}
        rowKey={'id'}
        toolbar={{}}
        scroll={{ x: tableWidth }}
        components={components}
        columns={finalColumns}
        columnsState={{ defaultValue: { id: { show: false } } }}
        request={async (params, sort, filter) => {
          const table = { data: [] as FormatMsgAlert[], success: true, total: 0 },
            where: MsgAlertWhereInput = {};
          if (params?.startsAt?.[0] && params?.startsAt?.[1]) {
            where.startsAtGTE = getDate(params.startsAt[0], 'YYYY-MM-DDT00:00:00Z');
            where.startsAtLTE = getDate(params.startsAt[1], 'YYYY-MM-DDT23:59:59Z');
          }
          let alertName = '';
          if (params?.alertName?.name) {
            alertName = params?.alertName?.name;
          }
          const result = await getFormatMsgAlertList({
            current: params.current,
            pageSize: params.pageSize,
            receiverType: params.receiverType,
            alertName: alertName,
            userID: params?.user?.id,
            where: where,
          });
          if (result?.totalCount) {
            table.data = result.edges?.map(item => item?.node) as FormatMsgAlert[];
            table.total = result.totalCount;
          }
          return table;
        }}
        pagination={{ showSizeChanger: true }}
      />
      <Modal
        title={t('view_content')}
        open={modal.show}
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
  auth: ['/msg/query'],
}));
