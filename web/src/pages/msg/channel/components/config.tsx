import { setLeavePromptWhen } from '@/components/LeavePrompt';
import { getMsgChannelReceiverInfo, updateMsgChannel } from '@/services/msgsrv/channel';
import { DrawerForm, ProFormText } from '@ant-design/pro-components';
import { useState } from 'react';
import { useTranslation } from 'react-i18next';
import Editor from '@monaco-editor/react';

type ProFormData = {
  receiver: string;
};

export default (props: {
  open: boolean;
  title?: string;
  id: string;
  onClose: (isSuccess?: boolean) => void;
}) => {
  const { t } = useTranslation(),
    [saveLoading, setSaveLoading] = useState(false),
    [saveDisabled, setSaveDisabled] = useState(true);

  setLeavePromptWhen(saveDisabled);

  const
    onOpenChange = (open: boolean) => {
      if (!open) {
        props.onClose?.();
      }
      setSaveDisabled(true);
    },
    getRequest = async () => {
      setSaveLoading(false);
      setSaveDisabled(true);
      const initData: ProFormData = {
        receiver: JSON.stringify({
          name: '',
          emailConfigs: []
        }, null, 4)
      }
      const result = await getMsgChannelReceiverInfo(props.id);
      if (result?.id) {
        if (result.receiver) {
          const receiver = result.receiver
          if (receiver?.__typename) {
            delete receiver.__typename
            receiver?.emailConfigs?.forEach(item => {
              if (item?.__typename) {
                delete item.__typename
              }
            })
          }
          initData.receiver = JSON.stringify(receiver, null, 4)
        }
      }
      return initData;
    },
    onValuesChange = () => {
      setSaveDisabled(false);
    },
    onFinish = async (values: ProFormData) => {
      setSaveLoading(true);

      const result = await updateMsgChannel(props.id, {
        receiver: JSON.parse(values.receiver),
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
      <ProFormText name="receiver">
        <Editor
          className="adminx-editor"
          height="70vh"
          defaultLanguage="json"
        />
      </ProFormText>
    </DrawerForm>
  );
};
