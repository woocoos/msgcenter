import { setLeavePromptWhen } from '@/components/LeavePrompt';
import { updateFormat } from '@/util';
import { DrawerForm, ProFormInstance, ProFormRadio, ProFormText, ProFormTextArea } from '@ant-design/pro-components';
import { useRef, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Org } from '@/__generated__/adminx/graphql';
import { MsgEvent, MsgTemplate, MsgTemplateFormat, MsgTemplateReceiverType } from '@/__generated__/msgsrv/graphql';
import { EnumMsgTemplateFormat, createMsgTemplate, getMsgTemplateInfo, updateMsgTemplate } from '@/services/msgsrv/template';
import InputOrg from '@/components/Adminx/Org/input';
import TempBtnUpload from '@/components/UploadFiles/tempBtn';
import { cacheOrg } from '@/services/adminx/org';
import Multiple from '@/components/UploadFiles/multiple';

type ProFormData = {
  org?: Org;
  name: string;
  comments?: string;
  subject: string;
  from?: string;
  to?: string;
  cc?: string;
  bcc?: string;
  format: MsgTemplateFormat;
  body?: string;
  tpl?: string;
  attachments?: string[];
};

export default (props: {
  open: boolean;
  title?: string;
  id?: string | null;
  msgEvent: MsgEvent;
  receiverType: MsgTemplateReceiverType
  onClose: (isSuccess?: boolean) => void;
}) => {
  const { t } = useTranslation(),
    formRef = useRef<ProFormInstance>(),
    [info, setInfo] = useState<MsgTemplate>(),
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
        name: '',
        subject: '',
        format: MsgTemplateFormat.Txt,
      }
      if (props.id) {
        const result = await getMsgTemplateInfo(props.id);
        if (result?.id) {
          setInfo(result as MsgTemplate);
          initData.org = cacheOrg[result.tenantID];
          initData.name = result.name;
          initData.subject = result.subject || '';
          initData.format = result.format;
          initData.comments = result.comments || undefined;
          initData.from = result.from || undefined;
          initData.to = result.to || undefined;
          initData.cc = result.cc || undefined;
          initData.bcc = result.bcc || undefined;
          initData.body = result.body || undefined;
          initData.tpl = result.tpl || undefined;
          initData.attachments = result.attachments?.split(',') || undefined;
        }
      }
      return initData;
    },
    onValuesChange = () => {
      setSaveDisabled(false);
    },
    onFinish = async (values: ProFormData) => {
      setSaveLoading(true);
      const input = {
        eventID: props.msgEvent.id,
        format: values.format,
        msgTypeID: Number(props.msgEvent.msgTypeID),
        name: values.name,
        receiverType: props.receiverType,
        tenantID: values.org?.id ? Number(values.org.id) : 0,
        subject: values.subject,
        from: values.from,
        to: values.to,
        cc: values.cc,
        bcc: values.bcc,
        body: values.body,
        tpl: values.tpl,
        attachments: values.attachments ? values.attachments.join(',') : undefined,
        comments: values.comments,
      }

      const result = props.id
        ? await updateMsgTemplate(props.id, updateFormat(input, info || {}))
        : await createMsgTemplate(input);
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
      formRef={formRef}
    >
      <ProFormText
        name="org"
        label={t('org')}
        rules={[
          { required: true, message: `${t('please_enter_org')}` },
        ]}>
        <InputOrg />
      </ProFormText>
      <ProFormText
        name="name"
        label={t('name')}
        rules={[
          { required: true, message: `${t('please_enter_name')}` },
        ]}
      />
      <ProFormTextArea
        name="comments"
        label={t('description')}
        placeholder={`${t('please_enter_description')}`}
      />
      <ProFormText
        name="subject"
        label={t('subject')}
        rules={[
          { required: true, message: `${t('please_enter_subject')}` },
        ]}
      />
      <ProFormText
        x-if={props.receiverType === MsgTemplateReceiverType.Email}
        name="cc"
        label={t('msg_temp_cc')}
      />
      <ProFormText
        x-if={props.receiverType === MsgTemplateReceiverType.Email}
        name="bcc"
        label={t('msg_temp_bcc')}
      />
      <ProFormText
        x-if={props.receiverType === MsgTemplateReceiverType.Email}
        name="to"
        label={t('msg_temp_to')}
      />
      <ProFormText
        x-if={props.receiverType === MsgTemplateReceiverType.Email}
        name="from"
        label={t('msg_temp_from')}
      />
      <ProFormRadio.Group
        name="format"
        label={t('msg_temp_format')}
        tooltip={t('msg_temp_format_tip')}
        valueEnum={EnumMsgTemplateFormat}
        rules={[
          { required: true, message: `${t('please_enter_msg_temp_format')}` },
        ]}
      />
      <ProFormTextArea
        name="body"
        fieldProps={{
          rows: 6,
        }}
      />
      <ProFormText name="tpl">
        <TempBtnUpload accept=".html,.txt" />
      </ProFormText>
      <ProFormText
        x-if={props.receiverType === MsgTemplateReceiverType.Email}
        name="attachments"
        label={t('attachments')}
        tooltip={t('attachments_tip')}
      >
        <Multiple accept=".doc,.docx,.jpg,.jpeg,.png,.pdf" />
      </ProFormText>
    </DrawerForm>
  );
};
