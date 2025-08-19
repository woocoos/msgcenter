// import { EnumAppKind, EnumAppStatus, getAppList } from '@/services/adminx/app';
import { Modal } from 'antd';
import { useState } from 'react';
import { useTranslation } from 'react-i18next';
import { ProColumns, ProTable } from '@ant-design/pro-components';
import store from '@/store';
import { MsgEvent, MsgEventWhereInput } from '@/generated/msgsrv/graphql';
import { EnumMsgEventStatus, getMsgEventList } from '@/services/msgsrv/event';
import { DictSelect, DictText } from '@knockout-js/org';

export default (props: {
  open: boolean;
  isMultiple?: boolean;
  title: string;
  tableTitle?: string;
  onClose: (selectData?: MsgEvent[]) => void;
}) => {
  const { t } = useTranslation(),
    columns: ProColumns<MsgEvent>[] = [
      // 有需要排序配置
      {
        title: t('msg_type_category'),
        dataIndex: 'msgTypeCategory',
        width: 120,
        order: 5,
        renderFormItem() {
          return <DictSelect dictCode="MsgCategory" placeholder={t('please_enter_category')} />
        },
        render(text, record) {
          return <DictText dictCode="MsgCategory" value={record.msgType?.category} />
        },
      },
      { title: t('msg_event_name'), dataIndex: 'name', width: 120 },
      {
        title: t('way_receiving'),
        dataIndex: 'modes',
        width: 120,
        search: false,
        render(text, record) {
          return record.modes?.split(',')?.join('、');
        },
      },
      { title: t('description'), dataIndex: 'comments', width: 120, order: 4 },
    ],
    [dataSource, setDataSource] = useState<MsgEvent[]>([]),
    // 选中处理
    [selectedRowKeys, setSelectedRowKeys] = useState<string[]>([]);

  return (
    <Modal
      title={props.title}
      open={props.open}
      onOk={() => {
        props?.onClose(dataSource.filter(item => selectedRowKeys.includes(item.id)));
      }}
      onCancel={() => {
        props?.onClose();
      }}
      width={900}
    >
      <ProTable
        rowKey={'id'}
        size="small"
        search={{
          searchText: `${t('query')}`,
          resetText: `${t('reset')}`,
          labelWidth: 'auto',
        }}
        options={false}
        scroll={{ x: 'max-content', y: 300 }}
        columns={columns}
        request={async (params, sort, filter) => {
          const table = { data: [] as MsgEvent[], success: true, total: 0 },
            where: MsgEventWhereInput = {};
          where.nameContains = params.name;
          if (params.msgTypeName || params.msgTypeCategory) {
            where.hasMsgTypeWith = [{
              nameContains: params.msgTypeName,
              categoryContains: params.msgTypeCategory,
            }];
          }
          if (params.comments) {
            where.commentsContains = params.comments;
          }
            const result = await getMsgEventList({
              current: params.current,
              pageSize: params.pageSize,
              where,
            });
            if (result?.totalCount && result.edges) {
              for (const item of result.edges) {
                if (item?.node) {
                  table.data.push(item.node as MsgEvent);
                }
              }
              table.total = result.totalCount;
            }
          setSelectedRowKeys([]);
          setDataSource(table.data);
          return table;
        }}
        pagination={{ showSizeChanger: true }}
        rowSelection={{
          selectedRowKeys: selectedRowKeys,
          onChange: (selectedRowKeys: string[]) => { setSelectedRowKeys(selectedRowKeys); },
          type: props.isMultiple ? 'checkbox' : 'radio',
        }}
        onRow={(record) => {
          return {
            onClick: () => {
              if (props.isMultiple) {
                if (selectedRowKeys.includes(record.id)) {
                  setSelectedRowKeys(selectedRowKeys.filter(id => id != record.id));
                } else {
                  selectedRowKeys.push(record.id);
                  setSelectedRowKeys([...selectedRowKeys]);
                }
              } else {
                setSelectedRowKeys([record.id]);
              }
            },
          };
        }}
      />
    </Modal>
  );
};
