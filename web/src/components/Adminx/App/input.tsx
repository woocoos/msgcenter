import { Input } from 'antd';
import ModalApp from './modal';
import { useState } from 'react';
import { CloseCircleFilled } from '@ant-design/icons';
import { useTranslation } from 'react-i18next';
import { App } from '@/__generated__/adminx/graphql';

export default (props: {
  value?: App;
  disabled?: boolean;
  onChange?: (value?: App) => void;
}) => {
  const { t } = useTranslation(),
    [modal, setModal] = useState<{
      open: boolean;
    }>({
      open: false,
    });

  return (
    <>
      <Input.Search
        value={props.value?.name || ''}
        disabled={props.disabled}
        placeholder={`${t('click_search_app')}`}
        suffix={props.value && props.disabled != true ? <CloseCircleFilled
          style={{ fontSize: '12px', cursor: 'pointer', color: 'rgba(0, 0, 0, 0.25)' }}
          onClick={() => {
            props.onChange?.(undefined);
          }}
        /> : <span />}
        onSearch={() => {
          modal.open = true;
          setModal({ ...modal });
        }}
      />
      <ModalApp
        open={modal.open}
        title={`${t('search_app')}`}
        onClose={(selectData) => {
          if (selectData?.length) {
            props.onChange?.(selectData[0]);
          }
          modal.open = false;
          setModal({ ...modal });
        }}
      />
    </>
  );
};
