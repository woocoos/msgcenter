
import { Modal } from 'antd';
import { useState } from 'react';
import { useTranslation } from 'react-i18next';
import { ProColumns, ProTable } from '@ant-design/pro-components';
import { MsgType, MsgTypeWhereInput } from '@/generated/msgsrv/graphql';
import { getMsgTypeList } from '@/services/msgsrv/type';
import { AppSelect } from '@knockout-js/org';
import { App, getApps } from '@knockout-js/api';

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
          return <AppSelect />
        },
        render: (text, record) => {
          const app = apps.find(item => item.id == record.appID)
          return record.appID ? app?.name || record.appID : '-';
        },
      },
      { title: t('category'), dataIndex: 'category', width: 120 },
      { title: t('name'), dataIndex: 'name', width: 120 },
      { title: t('description'), dataIndex: 'comments', width: 120, search: false },
    ],
    [apps, setApps] = useState<App[]>([]),
    [dataSource, setDataSource] = useState<MsgType[]>([]),
    // 选中处理
    [selectedRowKeys, setSelectedRowKeys] = useState<string[]>([]);

  return (
    <Modal
      title={props.title}
      open={props.open}
      width={900}
      onOk={() => {
        props?.onClose(dataSource.filter(item => selectedRowKeys.includes(item.id)));
      }}
      onCancel={() => {
        props?.onClose();
      }}
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
        request={async (params) => {
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
            setApps(await getApps(table.data.map(item => item.appID || '')));
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
