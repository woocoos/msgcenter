import { DrawerForm, ProFormInstance, ProFormText } from '@ant-design/pro-components';
import { useEffect, useRef, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { MsgTemplate, MsgTemplateReceiverType } from '@/generated/msgsrv/graphql';
import { getMsgTemplateInfo, testSendEamil, testSendMessage } from '@/services/msgsrv/template';
import { useLeavePrompt } from '@knockout-js/layout';
import StringRecord from '@/components/input/stringRecord';
import { UserSelect } from '@knockout-js/org';
import { parseGoTempKey } from '@/util';
import { getFileRaw, parseStorageData } from '@knockout-js/api';

type ProFormData = {
  userID?: string;
  email?: string;
  labels?: Record<string, string>;
  annotations?: Record<string, string>;
}

export default (props: {
  open: boolean;
  title?: string;
  id: string;
  onClose: (isSuccess?: boolean) => void;
}) => {
  const { t } = useTranslation(),
    formRef = useRef<ProFormInstance>(),
    [checkLeave, setLeavePromptWhen] = useLeavePrompt(),
    [info, setInfo] = useState<MsgTemplate>(),
    [tplParams, setTplParams] = useState<string[]>([]),
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
      const result = await getMsgTemplateInfo(props.id), keys: string[] = [];
      if (result?.id) {
        if (result.tpl) {
          const tplData = await parseStorageData(result.tpl)
          if (tplData?.path) {
            const fileRaw = await getFileRaw(tplData.path);
            if (fileRaw) {
              const c = await fileRaw.Body?.transformToString('utf-8')
              const r = new FileReader()
              r.readAsText(new Blob([c ?? '']), 'utf-8')
              r.onload = () => {
                const pgkeys = parseGoTempKey(r.result as string);
                if (pgkeys) {
                  setTplParams(pgkeys)
                }
              }
            }
          }
        } else if (result.body) {
          const pgkeys = parseGoTempKey(result.body);
          if (pgkeys) {
            setTplParams(pgkeys)
          }
        }
        setInfo(result as MsgTemplate);
      }
      return {};
    },
    onValuesChange = () => {
      setSaveDisabled(false);
    },
    onFinish = async (values: ProFormData) => {
      if (info) {
        setSaveLoading(true);
        let result: boolean | null = null;
        if (info.receiverType === MsgTemplateReceiverType.Email) {
          result = await testSendEamil(info.id, values.email ?? '', values.labels, values.annotations);
        } else if (info.receiverType === MsgTemplateReceiverType.Message) {
          result = await testSendMessage(info.id, values.userID ?? '', values.labels, values.annotations);
        }
        if (result) {
          setSaveDisabled(true);
          props.onClose(true);
        }
        setSaveLoading(false);
      }
      return false;
    };

  return (
    <DrawerForm<ProFormData>
      drawerProps={{
        width: 800,
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
      formRef={formRef}
    >
      {info?.receiverType === MsgTemplateReceiverType.Email ? <ProFormText
        name="email"
        label={t('receive_mail')}
      /> : <></>}
      {info?.receiverType === MsgTemplateReceiverType.Message ? <ProFormText
        name="userID"
        label={t('receiver')}
      >
        <UserSelect changeValue="id" />
      </ProFormText> : <></>}
      <ProFormText name="labels" label={`labels ${t('argument')}`} tooltip={`${t('default_added')}：receiver、alertname、tenant、skipSub、timestamp`}>
        <StringRecord />
      </ProFormText>
      <ProFormText name="annotations" label={`annotations${t('argument')}`}>
        <StringRecord keys={tplParams} />
      </ProFormText>
    </DrawerForm>
  );
};
