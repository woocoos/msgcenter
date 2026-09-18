import { Nlog, NlogReceiverType, NlogWhereInput } from "@/generated/msgsrv/graphql";
import { EnumNlogReceiverType, getMsgAlertLogList } from "@/services/msgsrv/list";
import { ActionType, PageContainer, ProTable } from "@ant-design/pro-components";
import { Link, useSearchParams } from "@ice/runtime";
import { OrgKind } from "@knockout-js/api/ucenter";
import { OrgSelect } from "@knockout-js/org";
import { Divider, Space, Typography } from "antd";
import { useRef } from "react";
import { useTranslation } from "react-i18next";
import { useResizableProTable, routeBreadcrumb } from "@/util/hook";

const List = () => {
  const { t } = useTranslation(),
    [searchParams] = useSearchParams(),
    msgAlertId = searchParams.get('id'),
    // 表格相关
    proTableRef = useRef<ActionType>(),
    // 可调整列宽的 ProTable
    { columns: finalColumns, components, tableWidth } = useResizableProTable<Nlog>(() => ({
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
          title: t('org'), dataIndex: 'tenant', width: 120, hideInTable: true,
          renderFormItem: () => {
            return <OrgSelect kind={OrgKind.Root} />
          },
        },
        {
          title: t('send_at'), dataIndex: 'sendAt', width: 160, valueType: "dateTime",
        },
        {
          title: t('expires_at'), dataIndex: 'expiresAt', width: 160, valueType: "dateTime",
        },
        { title: t('msg_log_groupKey'), dataIndex: 'groupKey', width: 200, search: false, ellipsis: true },
        { title: t('msg_log_receiver'), dataIndex: 'receiver', width: 200, search: false, ellipsis: true },
        {
          title: t('msg_log_receiverType'), dataIndex: 'receiverType', width: 120, align: 'center', search: false,
          filters: true,
          valueEnum: EnumNlogReceiverType,
        },
        // 占位列
        { search: false, hideInSetting: true },
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
        toolbar={{
          title: `${t('msg_log')}-ID：${msgAlertId}`
        }}
        scroll={{ x: tableWidth }}
        components={components}
        columns={finalColumns}
        columnsState={{ defaultValue: { id: { show: false } } }}
        request={async (params, sort, filter) => {
          const table = { data: [] as Nlog[], success: true, total: 0 },
            where: NlogWhereInput = {};
          where.tenantID = params.tenant?.id;
          where.sendAt = params.sendAt;
          where.expiresAt = params.expiresAt;
          where.receiverTypeIn = filter.receiverType as NlogReceiverType[];
          if (msgAlertId) {
            const result = await getMsgAlertLogList(msgAlertId, {
              current: params.current,
              pageSize: params.pageSize,
              where,
            });
            if (result?.totalCount) {
              table.data = result.edges?.map(item => item?.node) as Nlog[]
              table.total = result.totalCount;
            }
          }
          return table;
        }}
        pagination={{ showSizeChanger: true }}
      />
    </>
  );
};

export default () => {
  return (
    <PageContainer
      className="ko-page-container"
    >
      <div className="ka-content">
        <List />
      </div>
    </PageContainer>
  );
};
