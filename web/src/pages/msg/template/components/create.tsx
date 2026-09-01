import { updateFormat } from '@/util';
import { DrawerForm, ProFormInstance, ProFormRadio, ProFormText, ProFormTextArea } from '@ant-design/pro-components';
import { useEffect, useRef, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { CreateMsgTemplateInput, MsgEvent, MsgTemplate, MsgTemplateFormat, MsgTemplateReceiverType, UpdateMsgTemplateInput } from '@/generated/msgsrv/graphql';
import { EnumMsgTemplateFormat, createMsgTemplate, getMsgTemplateInfo, updateMsgTemplate, getMsgTemplateDefine } from '@/services/msgsrv/template';
import InputMultiple from '@/components/input/multiple';
import { UploadMultiple, UploadTemp, useLeavePrompt } from '@knockout-js/layout';
import { OrgSelect, UserSelect } from '@knockout-js/org';
import { getOrg } from '@knockout-js/api';
import { Org, OrgKind } from '@knockout-js/api/ucenter';
import store from '@/store';
import { Button, Col, Row, Space, Modal, message } from 'antd';
import { TemplateType } from '../../event';

type ProFormData = {
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
  attachmentNames?: string;
  userID?: string;
};

export default (props: {
  open: boolean;
  title?: string;
  type: string | null;
  id?: string | null;
  msgEvent: MsgEvent;
  receiverType: MsgTemplateReceiverType;
  readonly?: boolean;
  onClose: (isSuccess?: boolean, newInfo?: MsgTemplate) => void;
}) => {
  const { t } = useTranslation(),
    formRef = useRef<ProFormInstance>(),
    [userState] = store.useModel('user'),
    [checkLeave, setLeavePromptWhen] = useLeavePrompt(),
    [info, setInfo] = useState<MsgTemplate>(),
    [showCc, setShowCc] = useState(false),
    [showBcc, setShowBcc] = useState(false),
    [saveLoading, setSaveLoading] = useState(false),
    [saveDisabled, setSaveDisabled] = useState(true),
    iframeRef = useRef<HTMLIFrameElement>(null),
    [modal, setModal] = useState<{
      show: boolean;
    }>({
      show: false,
    });

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
        name: '',
        subject: '',
        format: MsgTemplateFormat.Txt,
      }
      if (props.id) {
        const result = await getMsgTemplateInfo(props.id) as MsgTemplate | null;
        if (result?.id) {
          setInfo(result);
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
          initData.attachments = result.attachments || undefined;
          initData.attachmentNames = result.attachments?.join(',') || undefined;
          initData.userID = Number(result.userID) > 0 ? `${result.userID}` : undefined;
        }
      }
      setShowCc(!!initData.cc)
      setShowBcc(!!initData.bcc)
      return initData;
    },
    onValuesChange = () => {
      setSaveDisabled(false);
    },
    onFinish = async (values: ProFormData) => {
      setSaveLoading(true);

      const input: UpdateMsgTemplateInput | CreateMsgTemplateInput = {
        eventID: props.msgEvent.id,
        format: values.format,
        msgTypeID: Number(props.msgEvent.msgTypeID),
        name: values.name,
        receiverType: props.receiverType,
        subject: values.subject,
        body: values.body,
        tpl: values.tpl,
        comments: values.comments,
      }
      if (props.type === TemplateType.customer) {
        input.tenantID = userState.tenantId;
        if (values.userID) {
          input.userID = values.userID;
        }
      }

      if (props.receiverType === MsgTemplateReceiverType.Email) {
        if (props.type === TemplateType.default) {
          input.attachments = values.attachmentNames
            ? values.attachmentNames.split(',').map(s => s.trim()).filter(Boolean)
            : undefined;
        } else {
          input.attachments = values.attachments;
        }
        input.cc = showCc ? values.cc : undefined;
        input.bcc = showBcc ? values.bcc : undefined;
        input.to = values.to;
        input.from = values.from;
      }
      const result = props.id
        ? await updateMsgTemplate(props.id, updateFormat(input, info || {}))
        : await createMsgTemplate(input as CreateMsgTemplateInput);
      if (result?.id) {
        setSaveDisabled(true);
        props.onClose(true, result as MsgTemplate);
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
      disabled={props.readonly}
      submitter={props.readonly ? false : {
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
      request={getRequest}
      onValuesChange={onValuesChange}
      onFinish={onFinish}
      onOpenChange={onOpenChange}
      formRef={formRef}
    >
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
      {props.type === TemplateType.customer ? (
        <ProFormText
          name="userID"
          label={t('user')}
          tooltip={t('msg_temp_user_tip')}
        >
          <UserSelect changeValue="id" />
        </ProFormText>
      ) : null}
      <ProFormText
        name="subject"
        label={t('subject')}
        rules={[
          { required: true, message: `${t('please_enter_subject')}` },
        ]}
      />
      <ProFormText
        x-if={props.receiverType === MsgTemplateReceiverType.Email}
        name="to"
        label={t('msg_temp_to')}
      >
        <InputMultiple disabled={props.readonly} decollator="," placeholder={`${t('please_enter_msg_temp_to')}`} />
      </ProFormText>
      <ProFormText
        x-if={props.receiverType === MsgTemplateReceiverType.Email && showCc}
        name="cc"
        label={t('msg_temp_cc')}
      >
        <InputMultiple disabled={props.readonly} decollator="," placeholder={`${t('please_enter_msg_temp_cc')}`} />
      </ProFormText>
      <ProFormText
        x-if={props.receiverType === MsgTemplateReceiverType.Email && showBcc}
        name="bcc"
        label={t('msg_temp_bcc')}
      >
        <InputMultiple disabled={props.readonly} decollator="," placeholder={`${t('please_enter_msg_temp_bcc')}`} />
      </ProFormText>
      <div x-if={props.receiverType === MsgTemplateReceiverType.Email && !props.readonly}>
        <Space>
          <a onClick={() => {
            formRef.current?.setFieldValue('cc', undefined);
            setShowCc(!showCc);
          }}>{t(showCc ? 'hidd_cc' : 'show_cc')}</a>
          <a onClick={() => {
            formRef.current?.setFieldValue('bcc', undefined);
            setShowBcc(!showBcc);
          }}>{t(showBcc ? 'hidd_bcc' : 'show_bcc')}</a>
        </Space>
      </div>
      <br x-if={props.receiverType === MsgTemplateReceiverType.Email} />
      <ProFormText
        x-if={props.receiverType === MsgTemplateReceiverType.Email}
        name="from"
        label={t('msg_temp_from')}
      >
        <InputMultiple disabled={props.readonly} decollator="," placeholder={`${t('please_enter_msg_temp_from')}`} />
      </ProFormText>
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
      <ProFormText
        x-if={props.type === TemplateType.customer}
        name="tpl"
      >
        <UploadTemp
          accept=".tmpl"
          directory={`${userState.tenantId}/msg/tpl`}
        />
      </ProFormText>
      <ProFormText
        x-if={props.receiverType === MsgTemplateReceiverType.Email && props.type === TemplateType.customer}
        name="attachments"
        label={t('attachments')}
        tooltip={t('attachments_tip')}
      >
        <UploadMultiple
          accept=".doc,.docx,.jpg,.jpeg,.png,.pdf"
          directory={`${userState.tenantId}/msg/att`}
        />
      </ProFormText>
      <Space x-if={props.type === TemplateType.default}>
        <a onClick={async () => {
          let format = formRef.current?.getFieldValue('format') || '', body = formRef.current?.getFieldValue('body') || '';
          if (format == '' || body == '') {
            return
          }
          const result = await getMsgTemplateDefine(format, body)
          setModal({ show: true })
          setTimeout(() => {
            if (iframeRef.current?.contentWindow) {
              iframeRef.current.contentWindow.document.write(`<pre>${result}</pre>`)
            } else if (iframeRef.current?.contentDocument) {
              iframeRef.current.contentDocument.write(`<pre>${result}</pre>`)
            }
          }, 200)
        }}
        >{t('temp_viewer')}</a>
        <a onClick={async () => {
          let format = formRef.current?.getFieldValue('format') || '', body = formRef.current?.getFieldValue('body') || '';
          if (format == '' || body == '') {
            return
          }
          const result = await getMsgTemplateDefine(format, body)
          if (result) {
            const blob = new Blob([result], { type: 'text/plain;charset=utf-8' })
            const url = URL.createObjectURL(blob)
            const a = document.createElement('a')
            a.href = url
            a.download = `${formRef.current?.getFieldValue('name') || 'template'}.tmpl`
            a.click()
            URL.revokeObjectURL(url)
          } else {
            message.warning(t('format_error'))
          }
        }}
        >{t('temp_down')}</a>
      </Space>
      <div style={{ height: '20px' }} />
      <ProFormText
        x-if={props.receiverType === MsgTemplateReceiverType.Email && props.type === TemplateType.default}
        name="attachmentNames"
        label={t('attachments')}
        tooltip={t('attachments_tip')}
      >
        <InputMultiple disabled={props.readonly} decollator="," placeholder={t('please_enter_attachment_names')} />
      </ProFormText>
      <Modal
        title={t('temp_viewer')}
        open={modal.show}
        destroyOnClose
        footer={null}
        width={800}
        onCancel={() => {
          setModal({ show: false })
        }}
      >
        <iframe style={{ width: '100%', height: '60vh', border: '0 none' }} ref={iframeRef}></iframe>
      </Modal>
    </DrawerForm>
  );
};
