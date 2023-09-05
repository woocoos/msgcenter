import { MsgEvent, MsgType } from '@/generated/msgsrv/graphql';
import { createMsgEvent, getMsgEventInfo, updateMsgEvent } from '@/services/msgsrv/event';
import { updateFormat } from '@/util';
import { DrawerForm, ProFormCheckbox, ProFormText, ProFormTextArea } from '@ant-design/pro-components';
import { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import InputMsgType from '../../type/components/input';
import { EnumMsgTemplateReceiverType } from '@/services/msgsrv/template';
import { useLeavePrompt } from '@knockout-js/layout';

type ProFormData = {
  msgType?: MsgType;
  name: string;
  modes: string[];
  comments?: string | null;
};

export default (props: {
  open: boolean;
  title?: string;
  id?: string | null;
  onClose: (isSuccess?: boolean) => void;
}) => {
  const { t } = useTranslation(),
    [info, setInfo] = useState<MsgEvent>(),
    [checkLeave, setLeavePromptWhen] = useLeavePrompt(),
    [saveLoading, setSaveLoading] = useState(false),
    [saveDisabled, setSaveDisabled] = useState(true);

  useEffect(() => {
    setLeavePromptWhen(saveDisabled);
  }, [saveDisabled]);

  const
    onOpenChange = (open: boolean) => {
      if (!open) {
        if (checkLeave()) {
          props.onClose?.();
        }
      } else {
        setSaveDisabled(true);
      }
    },
    getRequest = async () => {
      setSaveLoading(false);
      setSaveDisabled(true);
      const initData: ProFormData = {
        name: '',
        modes: [],
      }
      if (props.id) {
        const result = await getMsgEventInfo(props.id);
        if (result?.id) {
          setInfo(result as MsgEvent);
          initData.msgType = result.msgType as MsgType || undefined;
          initData.name = result.name;
          initData.modes = result.modes.split(',');
          initData.comments = result.comments;
        }
      }
      return initData;
    },
    onValuesChange = () => {
      setSaveDisabled(false);
    },
    onFinish = async (values: ProFormData) => {
      setSaveLoading(true);
      const result = props.id
        ? await updateMsgEvent(props.id, updateFormat({
          modes: values.modes.join(','),
          msgTypeID: values.msgType?.id || '',
          name: values.name,
          comments: values.comments,
        }, info || {}))
        : await createMsgEvent({
          modes: values.modes.join(','),
          msgTypeID: values.msgType?.id || '',
          name: values.name,
          comments: values.comments,
        });
      if (result?.id) {
        setSaveDisabled(true);
        props.onClose(true);
      }
      setSaveLoading(false);
      return false;
    };

  return (
    <DrawerForm<ProFormData>
      drawerProps={{
        width: 500,
        destroyOnClose: true,
        maskClosable: false,
      }}
      submitter={{
        searchConfig: {
          submitText: t('submit'),
          resetText: t('cancel'),
        },
        submitButtonProps: {
          loading: saveLoading,
          disabled: saveDisabled,
        },
      }}
      title={props.title}
      open={props?.open}
      onReset={getRequest}
      request={getRequest}
      onValuesChange={onValuesChange}
      onFinish={onFinish}
      onOpenChange={onOpenChange}
    >
      <ProFormText
        name="name"
        label={t('name')}
        rules={[
          { required: true, message: `${t('please_enter_name')}` },
        ]}
      />
      <ProFormText
        name="msgType"
        label={t('msg_type')}
        rules={[
          { required: true, message: `${t('please_enter_msg_type')}` },
        ]}>
        <InputMsgType />
      </ProFormText>
      <ProFormCheckbox.Group
        name="modes"
        label={t('way_receiving')}
        valueEnum={EnumMsgTemplateReceiverType}
        rules={[
          { required: true, message: `${t('please_enter_way_receiving')}` },
        ]}
      />
      <ProFormTextArea
        name="comments"
        label={t('description')}
        placeholder={`${t('please_enter_description')}`}
      />
    </DrawerForm>
  );
};
