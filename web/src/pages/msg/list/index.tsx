import { MsgAlert, MsgAlertAlertStatus, MsgAlertWhereInput } from "@/generated/msgsrv/graphql";
import { EnumMsgAlertStatus, getMsgAlertList } from "@/services/msgsrv/list";
import { ActionType, PageContainer, ProTable } from "@ant-design/pro-components";
import { Link } from "@ice/runtime";
import { OrgKind } from "@knockout-js/api/ucenter";
import { KeepAlive } from "@knockout-js/layout";
import { OrgSelect } from "@knockout-js/org";
import { Divider, Space, Typography } from "antd";
import { definePageConfig } from "ice";
import { useRef, useState } from "react";
import { useTranslation } from "react-i18next";
import { useResizableProTable, routeBreadcrumb } from "@/util/hook";

const List = () => {
  const { t } = useTranslation(),
    // 表格相关
    proTableRef = useRef<ActionType>(),
    // 可调整列宽的 ProTable
    { columns: finalColumns, components, tableWidth } = useResizableProTable<MsgAlert>(() => ({
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
          title: t('starts_at'), dataIndex: 'startsAt', width: 160, valueType: "dateTime",
        },
        {
          title: t('end_at'), dataIndex: 'endsAt', width: 160, valueType: "dateTime",
        },
        {
          title: t('msg_alert_labels'), dataIndex: 'labels', width: 200, ellipsis: true, search: false,
          render(text, record) {
            return record.labels ? mapStringRender(record.labels) : '-';
          },
        },
        {
          title: t('msg_alert_annotations'), dataIndex: 'annotations', width: 200, ellipsis: true, search: false,
          render(text, record) {
            return record.annotations ? mapStringRender(record.annotations) : '-';
          },
        },
        {
          title: t('msg_alert_timeout'), dataIndex: 'timeout', width: 120, align: 'center', search: false,
          render(text, record) {
            return record.timeout ? t('yes') : t('no');
          },
        },
        {
          title: t('status'), dataIndex: 'state', width: 120, align: 'center', search: false,
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
          width: 100,
          hideInSetting: true,
          render: (text, record) => {
            return (<Space split={<Divider type="vertical" className="ko-divider-gray" />} size={0}>
              <Link to={`/msg/list/log?id=${record.id}`} >
                {t('log')}
              </Link>
            </Space>);
          },
        },
      ],
    }), []);

  const mapStringRender = (mapString: Record<string, string>) => {
    const strAry: string[] = [];
    for (let key in mapString) {
      strAry.push(`${key}="${mapString[key]}"`);
    }
    return strAry.map(str => <div title={str}>{str}</div>)
  }

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
        pagination={{ showSizeChanger: true }}
      />
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
  auth: ['/msg/list'],
}));
