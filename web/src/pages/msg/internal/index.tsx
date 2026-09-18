import Auth from "@/components/auth";
import { MsgInternalTo, MsgInternalToWhereInput } from "@/generated/msgsrv/graphql";
import { delMarkMsg, getUserMsgCategory, getUserMsgCategoryNum, getUserMsgInternalList, markMsgRead } from "@/services/msgsrv/internal";
import { DownOutlined, LeftOutlined, RightOutlined } from "@ant-design/icons";
import { ActionType, PageContainer, ProTable } from "@ant-design/pro-components";
import { KeepAlive, Modal } from "@knockout-js/layout";
import { DictText } from "@knockout-js/org";
import { Badge, Button, Dropdown, Popconfirm, Space, Tabs, Typography, Divider, message } from "antd";
import { useEffect, useRef, useState } from "react";
import { useTranslation } from "react-i18next";
import styles from "./index.module.css";
import { definePageConfig } from "ice";
import { delDataSource } from "@/util";
import { useResizableProTable, routeBreadcrumb } from "@/util/hook";
import Detail from "@/pages/msg/internal/detail";

type CategoryTag = {
  name: string,
  code: string,
  num: number,
}

const List = () => {
  const { t } = useTranslation(),
    // 表格相关
    proTableRef = useRef<ActionType>(),
    [selectCategory, setSelectCategory] = useState('all'),
    [selectItems, setSelectItems] = useState('all'),
    [categorys, setCategorys] = useState<CategoryTag[]>([]),
    [dataSource, setDataSource] = useState<MsgInternalTo[]>([]),
    // 选中处理
    [selectedRowKeys, setSelectedRowKeys] = useState<string[]>([]),
    // 详情弹窗
    [detailModal, setDetailModal] = useState<{ show: boolean; id: string; key: number }>({
      show: false, id: '', key: 0,
    }),
    // 可调整列宽的 ProTable
    { columns: finalColumns, components, tableWidth } = useResizableProTable<MsgInternalTo>(() => ({
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
          title: t('subject'), dataIndex: 'subject', width: 260,
          renderText(text, record, index, action) {
            return record.readAt ? record.msgInternal.subject : <Badge color="red" text={record.msgInternal.subject} />
          },
        },
        {
          title: t('msg_type_category'), dataIndex: 'category', width: 120, align: 'center',
          renderText(text, record, index, action) {
            return <DictText dictCode="MsgCategory" value={record.msgInternal.category} />
          },
        },
        {
          title: t('created_at'), dataIndex: 'createdAt', width: 160, valueType: "dateTime"
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
            return (<Space split={<Divider type="vertical" className="ko-divider-gray" />} size={0}>
              <Typography.Link
                onClick={() => {
                  setDetailModal({ show: true, id: record.id, key: Date.now() });
                }}
              >
                {t('detail')}
              </Typography.Link>
              <Auth authKey="markMsgInternalToDeleted">
                <Popconfirm
                  title={t('delete')}
                  description={`${t('confirm_delete')}：${record.msgInternal.subject}`}
                  onConfirm={async () => {
                    const result = await delMarkMsg([record.id]);
                    if (result) {
                      setDataSource(delDataSource(dataSource, record.id))
                      message.success(t('submit_success'));
                    }
                  }}
                >
                  <Typography.Link>
                    {t('delete')}
                  </Typography.Link>
                </Popconfirm>
              </Auth>
            </Space>);
          },
        },
      ],
    }), []);

  const requestCategory = async () => {
    let allNum = 0;
    const cList: CategoryTag[] = [];
    const cResult = await getUserMsgCategory();

    if (Array.isArray(cResult)) {
      const numResult = await getUserMsgCategoryNum(cResult);
      if (Array.isArray(numResult)) {
        for (let i = 0; i < cResult.length; i++) {
          const curNum = numResult[i] ?? 0;
          cList.push({
            name: cResult[i],
            code: cResult[i],
            num: curNum,
          })
          allNum += curNum;
        }
      }
    }
    setCategorys([
      {
        name: t('all_msg'),
        code: 'all',
        num: allNum,
      },
      ...cList,
    ])
  }

  useEffect(() => {
    requestCategory();
  }, [])

  return (
    <>
      <DictText dictCode="MsgCategory" />
      <ProTable
        actionRef={proTableRef}
        sticky={dataSource.length > 0 ? { offsetHeader: 56 } : undefined}
        search={false}
        className={styles.tableTabs}
        rowKey={'id'}
        toolbar={{
          title: <Tabs
            x-if={categorys.length}
            activeKey={selectCategory}
            items={categorys.map(item => (
              {
                key: item.code,
                label: (
                  <span key={item.code}>
                    {item.code === 'all' ? item.name : <DictText dictCode="MsgCategory" value={item.name} />}
                    {`${item.num ? `(${item.num})` : ''} `}
                  </span>
                )
              }
            ))}
            onChange={(activeKey) => {
              proTableRef.current?.reload(true);
              setSelectCategory(activeKey);
            }}>
          </Tabs>,
          actions: [
            <Dropdown menu={{
              items: [
                {
                  key: 'all', label: <span onClick={() => {
                    proTableRef.current?.reload(true);
                    setSelectItems('all');
                  }}>{t('all_msg')}</span>
                },
                {
                  key: 'read', label: <span onClick={() => {
                    proTableRef.current?.reload(true);
                    setSelectItems('read');
                  }}>{t('read_msg')}</span>
                },
                {
                  key: 'unread', label: <span onClick={() => {
                    proTableRef.current?.reload(true);
                    setSelectItems('unread');
                  }}>{t('unread_msg')}</span>
                },
              ]
            }}>
              <Space>
                {selectItems === 'read' ? t('read_msg') : selectItems === 'unread' ? t('unread_msg') : t('all_msg')}
                <DownOutlined />
              </Space>
            </Dropdown>,
            <Auth authKey="markMsgInternalToReadOrUnRead">
              <Button type="primary" onClick={async () => {
                if (selectedRowKeys.length) {
                  const result = await markMsgRead(selectedRowKeys, true);
                  if (result) {
                    proTableRef.current?.reload();
                    await requestCategory();
                    message.success(t('submit_success'));
                  }
                } else {
                  message.warning(t('please_select_data'))
                }
              }}>{t('mark_read')}</Button>
            </Auth>
          ],
        }}
        scroll={{ x: tableWidth }}
        components={components}
        columns={finalColumns}
        columnsState={{ defaultValue: { id: { show: false } } }}
        dataSource={dataSource}
        request={async (params, sort, filter) => {
          const table = { data: [] as MsgInternalTo[], success: true, total: 0 },
            where: MsgInternalToWhereInput = {};
          if (selectCategory != 'all') {
            where.hasMsgInternalWith = [{
              category: selectCategory
            }]
          }
          if (selectItems === 'unread') {
            where.readAtIsNil = true;
          } else if (selectItems === 'read') {
            where.readAtNotNil = true;
          }
          const result = await getUserMsgInternalList({
            current: params.current,
            pageSize: params.pageSize,
            where,
          });
          if (result?.totalCount) {
            result.edges?.forEach(item => {
              if (item?.node) {
                table.data.push(item.node as MsgInternalTo)
              }
            })
            table.total = result.totalCount;
          }
          setDataSource(table.data)
          setSelectedRowKeys([]);
          return table;
        }}
        rowSelection={{
          selectedRowKeys: selectedRowKeys,
          onChange: (selectedRowKeys) => { setSelectedRowKeys(selectedRowKeys as string[]); },
          type: 'checkbox',
        }}
      />
      <Modal
        title={t('station_msg_detail')}
        open={detailModal.show}
        destroyOnHidden
        width={1100}
        onCancel={() => {
          setDetailModal({ show: false, id: '', key: 0 });
        }}
        footer={
          <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
            <Button
              icon={<LeftOutlined />}
              disabled={!detailModal.id || dataSource.findIndex(d => d.id === detailModal.id) <= 0}
              onClick={() => {
                const currentIndex = dataSource.findIndex(d => d.id === detailModal.id);
                if (currentIndex > 0) {
                  const prevItem = dataSource[currentIndex - 1];
                  if (prevItem) {
                    setDetailModal({ show: true, id: prevItem.id, key: Date.now() });
                  }
                }
              }}
            >
              {t('previous')}
            </Button>
            <Button
              disabled={!detailModal.id || dataSource.findIndex(d => d.id === detailModal.id) >= dataSource.length - 1}
              onClick={() => {
                const currentIndex = dataSource.findIndex(d => d.id === detailModal.id);
                if (currentIndex < dataSource.length - 1) {
                  const nextItem = dataSource[currentIndex + 1];
                  if (nextItem) {
                    setDetailModal({ show: true, id: nextItem.id, key: Date.now() });
                  }
                }
              }}
            >
              {t('next')}
              <RightOutlined />
            </Button>
          </div>
        }
      >
        {detailModal.id ? (
          <Detail
            key={detailModal.key}
            id={detailModal.id}
            onRead={() => {
              proTableRef.current?.reload();
              requestCategory();
            }}
          />
        ) : null}
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
  auth: ['/msg/internal'],
}));
