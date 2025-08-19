import { Input } from 'antd';
import ModalMsgEvent from './modalMsgEvent';
import { useState } from 'react';
import { CloseCircleFilled } from '@ant-design/icons';
import { useTranslation } from 'react-i18next';
import { MsgEvent } from '@/generated/msgsrv/graphql';

export default (props: {
  value?: MsgEvent;
  disabled?: boolean;
  onChange?: (value?: MsgEvent) => void;
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
        value={props.value?.comments || ''}
        disabled={props.disabled}
        placeholder={`${t('click_search_msg_event')}`}
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
      <ModalMsgEvent
        open={modal.open}
        title={`${t('search_msg_event')}`}
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
