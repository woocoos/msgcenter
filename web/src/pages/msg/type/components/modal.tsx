
import { Modal } from 'antd';
import { useState } from 'react';
import { useTranslation } from 'react-i18next';
import { ProColumns, ProTable } from '@ant-design/pro-components';
import { TableFilter, TableParams, TableSort } from '@/services/graphql';
import { MsgType, MsgTypeWhereInput } from '@/__generated__/msgsrv/graphql';
import InputApp from '@/components/Adminx/App/input';
import { cacheApp, updateCacheAppListByIds } from '@/services/adminx/app/indtx';
import { getMsgTypeList } from '@/services/msgsrv/type';

export default (props: {
  open: boolean;
  isMultiple?: boolean;
  title: string;
  tableTitle?: string;
  onClose: (selectData?: MsgType[]) => void;
}) => {
  const { t } = useTranslation(),
    columns: ProColumns<MsgType>[] = [
      // 有需要排序配置  sorter: true
      {
        title: t('app'), dataIndex: 'app', width: 120,
        renderFormItem() {
          return <InputApp />
        },
        render: (text, record) => {
          return record.appID ? cacheApp[record.appID]?.name || record.appID : '-';
        },
      },
      { title: t('category'), dataIndex: 'category', width: 120 },
      { title: t('name'), dataIndex: 'name', width: 120 },
      { title: t('description'), dataIndex: 'comments', width: 120, search: false },
    ],
    [dataSource, setDataSource] = useState<MsgType[]>([]),
    // 选中处理
    [selectedRowKeys, setSelectedRowKeys] = useState<string[]>([]);

  const
    getRequest = async (params: TableParams, sort: TableSort, filter: TableFilter) => {
      const table = { data: [] as MsgType[], success: true, total: 0 },
        where: MsgTypeWhereInput = {};
      where.appID = params.app?.id;
      where.category = params.category;
      where.nameContains = params.name;

      const result = await getMsgTypeList({
        current: params.current,
        pageSize: params.pageSize,
        where,
      });
      if (result?.totalCount) {
        table.data = result.edges?.map(item => item?.node) as MsgType[];
        await updateCacheAppListByIds(table.data.map(item => item.appID || ''))
        table.total = result.totalCount;
      }
      setSelectedRowKeys([]);
      setDataSource(table.data);
      return table;
    },
    handleOk = () => {
      props?.onClose(dataSource.filter(item => selectedRowKeys.includes(item.id)));
    },
    handleCancel = () => {
      props?.onClose();
    };

  return (
    <Modal title={props.title} open={props.open} onOk={handleOk} onCancel={handleCancel} width={900}>
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
        request={getRequest}
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
