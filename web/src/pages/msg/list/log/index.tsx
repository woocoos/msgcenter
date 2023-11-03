import { Nlog, NlogReceiverType, NlogWhereInput } from "@/generated/msgsrv/graphql";
import { EnumNlogReceiverType, getMsgAlertLogList } from "@/services/msgsrv/list";
import { ActionType, PageContainer, ProColumns, ProTable, useToken } from "@ant-design/pro-components";
import { Link, useSearchParams } from "@ice/runtime";
import { OrgKind } from "@knockout-js/api/ucenter";
import { OrgSelect } from "@knockout-js/org";
import { useRef } from "react";
import { useTranslation } from "react-i18next";

export default () => {
  const { token } = useToken(),
    { t } = useTranslation(),
    [searchParams] = useSearchParams(),
    // 表格相关
    proTableRef = useRef<ActionType>(),
    columns: ProColumns<Nlog>[] = [
      // 有需要排序配置  sorter: true
      {
        title: t('org'), dataIndex: 'tenant', width: 120, hideInTable: true,
        renderFormItem: () => {
          return <OrgSelect kind={OrgKind.Root} />
        },
      },
      {
        title: t('send_at'), dataIndex: 'sendAt', width: 120, valueType: "dateTime",
      },
      {
        title: t('expires_at'), dataIndex: 'expiresAt', width: 120, valueType: "dateTime",
      },
      { title: t('msg_log_groupKey'), dataIndex: 'groupKey', width: 120, search: false, },
      { title: t('msg_log_receiver'), dataIndex: 'receiver', width: 120, search: false, },
      {
        title: t('msg_log_receiverType'), dataIndex: 'receiverType', width: 120, search: false,
        filters: true,
        valueEnum: EnumNlogReceiverType,
      },
    ];

  return <PageContainer
    header={{
      title: t('msg_log'),
      style: { background: token.colorBgContainer },
      breadcrumb: {
        items: [
          { title: t('msg_center') },
          { title: <Link to="/msg/list">{t('msg_alert')}</Link> },
          { title: t('msg_log') },
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
        title: t('msg_log_list'),
      }}
      scroll={{ x: 'max-content' }}
      columns={columns}
      request={async (params, sort, filter) => {
        const table = { data: [] as Nlog[], success: true, total: 0 },
          msgAlertId = searchParams.get('id') ?? '',
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

    />
  </PageContainer>
}
