import { Input } from 'antd';
import ModalMsgType from './modal';
import { useState } from 'react';
import { CloseCircleFilled } from '@ant-design/icons';
import { useTranslation } from 'react-i18next';
import { MsgType } from '@/__generated__/msgsrv/graphql';

export default (props: {
  value?: MsgType;
  disabled?: boolean;
  onChange?: (value?: MsgType) => void;
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
        placeholder={`${t('click_search_msg_type')}`}
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
      <ModalMsgType
        open={modal.open}
        title={`${t('search_msg_type')}`}
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
