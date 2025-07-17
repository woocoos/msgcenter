import Auth from "@/components/auth";
import { MsgInternalTo, MsgInternalToWhereInput } from "@/generated/msgsrv/graphql";
import { delMarkMsg, getUserMsgCategory, getUserMsgCategoryNum, getUserMsgInternalList, markMsgRead } from "@/services/msgsrv/internal";
import { DownOutlined } from "@ant-design/icons";
import { ActionType, PageContainer, ProColumns, ProTable, useToken } from "@ant-design/pro-components";
import { Link } from "@ice/runtime";
import { KeepAlive } from "@knockout-js/layout";
import { DictText } from "@knockout-js/org";
import { Badge, Button, Dropdown, Popconfirm, Space, Tabs, message } from "antd";
import { useEffect, useRef, useState } from "react";
import { useTranslation } from "react-i18next";
import styles from "./index.module.css";
import { definePageConfig } from "ice";
import { delDataSource } from "@/util";

type CategoryTag = {
  name: string,
  code: string,
  num: number,
}

export default () => {
  const { token } = useToken(),
    { t } = useTranslation(),
    // 表格相关
    proTableRef = useRef<ActionType>(),
    [selectCategory, setSelectCategory] = useState('all'),
    [selectItems, setSelectItems] = useState('all'),
    [categorys, setCategorys] = useState<CategoryTag[]>([]),
    columns: ProColumns<MsgInternalTo>[] = [
      // 有需要排序配置  sorter: true
      {
        title: t('subject'), dataIndex: 'subject', width: 120,
        renderText(text, record, index, action) {
          return record.readAt ? record.msgInternal.subject : <Badge color="red" text={record.msgInternal.subject} />
        },
      },
      {
        title: t('msg_type_category'), dataIndex: 'category', width: 120,
        renderText(text, record, index, action) {
          return <DictText dictCode="MsgCategory" value={record.msgInternal.category} />
        },
      },
      {
        title: t('created_at'), dataIndex: 'createdAt', width: 120, valueType: "dateTime"
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
            <Link to={`/msg/internal/detail?toid=${record.id}`} >
              {t('detail')}
            </Link>
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
                <a>
                  {t('delete')}
                </a>
              </Popconfirm>
            </Auth>
          </Space>);
        },
      },
    ],
    [dataSource, setDataSource] = useState<MsgInternalTo[]>([]),
    // 选中处理
    [selectedRowKeys, setSelectedRowKeys] = useState<string[]>([]);

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
            code: `${i}`,
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

  return <KeepAlive clearAlive>
    <PageContainer
      header={{
        title: t('station_msg'),
        style: { background: token.colorBgContainer },
        breadcrumb: {
          items: [
            { title: t('msg_center') },
            { title: t('station_msg') },
          ],
        },
      }}
    >
      <DictText dictCode="MsgCategory" />
      <ProTable
        actionRef={proTableRef}
        search={false}
        className={styles.tableTabs}
        rowKey={'id'}
        toolbar={{
          title: <Tabs
            x-if={categorys.length}
            activeKey={selectCategory}
            items={categorys.map(item => (
              {
                key: item.code, label: <span key={item.code}>
                  {item.code === 'all' ? item.name : <DictText dictCode="MsgCategory" value={item.name} />}
                  {`${item.num ? `(${item.num})` : ''} `}
                </span>
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
        scroll={{ x: 'max-content' }}
        columns={columns}
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
          setSelectedRowKeys([]);
          return table;
        }}
        rowSelection={{
          selectedRowKeys: selectedRowKeys,
          onChange: (selectedRowKeys: string[]) => { setSelectedRowKeys(selectedRowKeys); },
          type: 'checkbox',
        }}
      />
    </PageContainer>
  </KeepAlive>
}

export const pageConfig = definePageConfig(() => ({
  auth: ['/msg/internal'],
}));
