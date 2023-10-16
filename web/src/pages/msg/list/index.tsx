import { MsgAlert, MsgAlertAlertStatus, MsgAlertWhereInput } from "@/generated/msgsrv/graphql";
import { EnumMsgAlertStatus, getMsgAlertList } from "@/services/msgsrv/list";
import { ActionType, PageContainer, ProColumns, ProTable, useToken } from "@ant-design/pro-components";
import { Link } from "@ice/runtime";
import { OrgKind } from "@knockout-js/api/ucenter";
import { KeepAlive } from "@knockout-js/layout";
import { OrgSelect } from "@knockout-js/org";
import { Space } from "antd";
import { definePageConfig } from "ice";
import { useRef } from "react";
import { useTranslation } from "react-i18next";

export default () => {
  const { token } = useToken(),
    { t } = useTranslation(),
    // 表格相关
    proTableRef = useRef<ActionType>(),
    columns: ProColumns<MsgAlert>[] = [
      // 有需要排序配置  sorter: true
      {
        title: t('org'), dataIndex: 'tenant', width: 120, hideInTable: true,
        renderFormItem: () => {
          return <OrgSelect kind={OrgKind.Root} />
        },
      },
      {
        title: t('starts_at'), dataIndex: 'startsAt', width: 120, valueType: "dateTime",
      },
      {
        title: t('end_at'), dataIndex: 'endsAt', width: 120, valueType: "dateTime",
      },
      {
        title: t('msg_alert_labels'), dataIndex: 'labels', width: 120, search: false,
        render(text, record) {
          return record.labels ? mapStringRender(record.labels) : '-';
        },
      },
      {
        title: t('msg_alert_annotations'), dataIndex: 'annotations', width: 120, search: false,
        render(text, record) {
          return record.annotations ? mapStringRender(record.annotations) : '-';
        },
      },
      {
        title: t('msg_alert_timeout'), dataIndex: 'timeout', width: 120, search: false,
        render(text, record) {
          return record.timeout ? t('yes') : t('no');
        },
      },
      {
        title: t('status'), dataIndex: 'state', width: 120, search: false,
        filters: true,
        valueEnum: EnumMsgAlertStatus,
      },
      {
        title: t('operation'),
        dataIndex: 'actions',
        fixed: 'right',
        align: 'center',
        search: false,
        width: 80,
        render: (text, record) => {
          return (<Space>
            <Link to={`/msg/list/log?id=${record.id}`} >
              {t('log')}
            </Link>
          </Space>);
        },
      },
    ];

  const mapStringRender = (mapString: Record<string, string>) => {
    const strAry: string[] = [];
    for (let key in mapString) {
      strAry.push(`${key}="${mapString[key]}"`);
    }
    return strAry.join(',')
  }

  return <KeepAlive clearAlive>
    <PageContainer
      header={{
        title: t('msg_alert'),
        style: { background: token.colorBgContainer },
        breadcrumb: {
          items: [
            { title: t('msg_center') },
            { title: t('msg_alert') },
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
          title: t('msg_alert_list'),
        }}
        scroll={{ x: 'max-content' }}
        columns={columns}
        request={async (params, sort, filter) => {
          const table = { data: [] as MsgAlert[], success: true, total: 0 },
            where: MsgAlertWhereInput = {};
          where.tenantID = params.tenant?.id;
          where.startsAt = params.startsAt;
          where.endsAt = params.endsAt;
          where.stateIn = filter.status as MsgAlertAlertStatus[];
          const result = await getMsgAlertList({
            current: params.current,
            pageSize: params.pageSize,
            where,
          });
          if (result?.totalCount) {
            table.data = result.edges?.map(item => item?.node) as MsgAlert[]
            table.total = result.totalCount;
          }
          return table;
        }}

      />
    </PageContainer>
  </KeepAlive>
}

export const pageConfig = definePageConfig(() => ({
  auth: ['/msg/list'],
}));
