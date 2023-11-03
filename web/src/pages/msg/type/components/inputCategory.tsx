import { getMsgTypeCategoryList } from '@/services/msgsrv/type';
import { CloseCircleFilled, LoadingOutlined } from '@ant-design/icons';
import { AutoComplete, Input } from 'antd';
import { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';

let stimeFn: NodeJS.Timeout | undefined = undefined;

export default (props: {
  value?: string;
  disabled?: boolean;
  onChange?: (value?: string) => void;
}) => {
  const { t } = useTranslation(),
    [loading, setLoading] = useState(false),
    [options, setOptions] = useState<{ value: string }[]>([]);

  const
    onSearch = async (value?: string) => {
      setLoading(true);
      clearTimeout(stimeFn);
      stimeFn = setTimeout(async () => {
        const result = await getMsgTypeCategoryList(value)
        setOptions(result?.map(item => ({
          value: item,
        })) ?? [])
        setLoading(false);
      }, 500);
    }

  useEffect(() => {
    onSearch();
    return () => {
      clearTimeout(stimeFn);
    }
  }, [])

  return <AutoComplete
    value={props.value}
    options={options}
    onSearch={onSearch}
    onChange={props.onChange}
  >
    <Input
      disabled={props.disabled}
      suffix={loading ? <LoadingOutlined /> : (
        !props.disabled && props.value ? <CloseCircleFilled
          style={{ fontSize: '12px', cursor: 'pointer', color: 'rgba(0, 0, 0, 0.25)' }}
          onClick={() => {
            setOptions([])
            props.onChange?.(undefined);
          }}
        /> : <></>
      )}
      placeholder={t('please_enter_msg_type_category') || ''}
    />
  </AutoComplete>;
};
