
import { getMsgChannelReceiverInfo, updateMsgChannel } from '@/services/msgsrv/channel';
import { DrawerForm, ProFormText } from '@ant-design/pro-components';
import { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { useLeavePrompt } from '@knockout-js/layout';
import * as yaml from 'js-yaml'
import Editor from '@/components/editor';

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
          setSaveDisabled(true);
        }
      } else {
        setSaveDisabled(true);
      }
    },
    getRequest = async () => {
      setSaveLoading(false);
      setSaveDisabled(true);
      const initData: ProFormData = {
        receiver: yaml.dump({
          name: '',
          emailConfigs: [],
          messageConfig: {
            redirect: '',
            subject: '',
            to: ''
          }
        })
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
            if (receiver.messageConfig) {
              delete receiver.messageConfig.__typename;
            }
          }
          initData.receiver = yaml.dump(receiver)
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
        receiver: yaml.load(values.receiver, { json: true }),
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
        width: 1000,
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
      <ProFormText name="receiver">
        <Editor
          className="adminx-editor"
          height="70vh"
          defaultLanguage="yaml"
        />
      </ProFormText>
    </DrawerForm>
  );
};
