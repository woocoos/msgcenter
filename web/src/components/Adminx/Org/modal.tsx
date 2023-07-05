
import { Modal } from 'antd';
import { useState } from 'react';
import { useTranslation } from 'react-i18next';
import { ProColumns, ProTable } from '@ant-design/pro-components';
import { TableFilter, TableParams, TableSort } from '@/services/graphql';
import { Org, OrgWhereInput } from '@/__generated__/adminx/graphql';
import { getOrgList } from '@/services/adminx/org';

export default (props: {
  open: boolean;
  isMultiple?: boolean;
  title: string;
  tableTitle?: string;
  onClose: (selectData?: Org[]) => void;
}) => {
  const { t } = useTranslation(),
    columns: ProColumns<Org>[] = [
      // 有需要排序配置  sorter: true
      {
        title: t('name'),
        dataIndex: 'name',
        width: 120,
        search: {
          transform: (value) => ({ nameContains: value || undefined }),
        },
      },
      {
        title: t('domain'),
        dataIndex: 'domain',
        width: 120,
        search: {
          transform: (value) => ({ codeContains: value || undefined }),
        },
      },
      {
        title: t('manage_account'),
        dataIndex: 'owner',
        width: 120,
        search: false,
        render: (text, record) => {
          return <div>{record?.owner?.displayName || '-'}</div>;
        },
      },
      { title: t('description'), dataIndex: 'profile', width: 160, search: false },
    ],
    [dataSource, setDataSource] = useState<Org[]>([]),
    // 选中处理
    [selectedRowKeys, setSelectedRowKeys] = useState<string[]>([]);

  const
    getRequest = async (params: TableParams, sort: TableSort, filter: TableFilter) => {
      const table = { data: [] as Org[], success: true, total: 0 },
        where: OrgWhereInput = {};
      where.nameContains = params.nameContains;
      where.codeContains = params.codeContains;

      const result = await getOrgList({
        current: params.current,
        pageSize: params.pageSize,
        where,
      });
      if (result?.totalCount) {
        table.data = result.edges?.map(item => item?.node) as Org[];
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
