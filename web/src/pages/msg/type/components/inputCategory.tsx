import { getMsgTypeCategoryList } from '@/services/msgsrv/type';
import { CloseCircleFilled, LoadingOutlined } from '@ant-design/icons';
import { AutoComplete, Input } from 'antd';
import { useState } from 'react';
import { useTranslation } from 'react-i18next';

export default (props: {
  value?: string;
  disabled?: boolean;
  onChange?: (value?: string) => void;
}) => {
  const { t } = useTranslation(),
    [loading, setLoading] = useState(false),
    [options, setOptions] = useState<{ value: string }[]>([]),
    [stFn, setStFn] = useState<NodeJS.Timeout>();

  const
    onSearch = async (value: string) => {
      setLoading(true);
      clearTimeout(stFn);
      const stout = setTimeout(async () => {
        if (value) {
          const result = await getMsgTypeCategoryList(value)
          if (result) {
            setOptions(result.map(item => ({
              value: item,
            })))
          }
        } else {
          setOptions([])
        }
        setLoading(false);
      }, 500);
      setStFn(stout);
    }

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
