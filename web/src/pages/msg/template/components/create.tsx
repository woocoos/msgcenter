import { updateFormat } from '@/util';
import { DrawerForm, ProFormInstance, ProFormRadio, ProFormText, ProFormTextArea } from '@ant-design/pro-components';
import { useEffect, useRef, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { MsgEvent, MsgTemplate, MsgTemplateFormat, MsgTemplateReceiverType } from '@/generated/msgsrv/graphql';
import { EnumMsgTemplateFormat, createMsgTemplate, getMsgTemplateInfo, updateMsgTemplate } from '@/services/msgsrv/template';
import InputMultiple from '@/components/input/multiple';
import { UploadMultiple, UploadTemp, useLeavePrompt } from '@knockout-js/layout';
import { OrgSelect } from '@knockout-js/org';
import { Org, OrgKind, getOrg } from '@knockout-js/api';
import store from '@/store';
import { Col, Row } from 'antd';

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
  tplFileID?: string;
  attachmentsFileIds?: string[];
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
    [userState] = store.useModel('user'),
    [checkLeave, setLeavePromptWhen] = useLeavePrompt(),
    [info, setInfo] = useState<MsgTemplate>(),
    [tpl, setTpl] = useState<string>(),
    [attachments, setAttachments] = useState<string[]>(),
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
        subject: '',
        format: MsgTemplateFormat.Txt,
      }
      if (props.id) {
        const result = await getMsgTemplateInfo(props.id);
        if (result?.id) {
          setInfo(result as MsgTemplate);
          initData.org = await getOrg(result.tenantID) as Org;
          initData.name = result.name;
          initData.subject = result.subject || '';
          initData.format = result.format;
          initData.comments = result.comments || undefined;
          initData.from = result.from || undefined;
          initData.to = result.to || undefined;
          initData.cc = result.cc || undefined;
          initData.bcc = result.bcc || undefined;
          initData.body = result.body || undefined;
          initData.tplFileID = result.tplFileID || undefined;
          initData.attachmentsFileIds = result.attachmentsFileIds || undefined;
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
        tenantID: values.org?.id ? values.org.id : "",
        subject: values.subject,
        from: values.from,
        to: values.to,
        cc: values.cc,
        bcc: values.bcc,
        body: values.body,
        tpl: tpl,
        tplFileID: values.tplFileID,
        attachments: attachments,
        attachmentsFileIds: values.attachmentsFileIds ?? undefined,
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
      formRef={formRef}
    >
      <Row gutter={20}>
        <Col span={8}>
          <ProFormText
            name="org"
            label={t('org')}
            rules={[
              { required: true, message: `${t('please_enter_org')}` },
            ]}>
            <OrgSelect kind={OrgKind.Root} />
          </ProFormText>
        </Col>
        <Col span={8}>
          <ProFormText
            name="name"
            label={t('name')}
            rules={[
              { required: true, message: `${t('please_enter_name')}` },
            ]}
          />
        </Col>
      </Row>
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
      >
        <InputMultiple decollator=";" />
      </ProFormText>
      <ProFormText
        x-if={props.receiverType === MsgTemplateReceiverType.Email}
        name="bcc"
        label={t('msg_temp_bcc')}
      >
        <InputMultiple decollator=";" />
      </ProFormText>
      <ProFormText
        x-if={props.receiverType === MsgTemplateReceiverType.Email}
        name="to"
        label={t('msg_temp_to')}
      >
        <InputMultiple decollator=";" />
      </ProFormText>
      <ProFormText
        x-if={props.receiverType === MsgTemplateReceiverType.Email}
        name="from"
        label={t('msg_temp_from')}
      >
        <InputMultiple decollator=";" />
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
      <ProFormText name="tplFileID">
        <UploadTemp
          accept=".tmpl"
          forceDirectory
          directory={`/msg/tpl/temp/${userState.tenantId}`}
          onChangePath={setTpl}
        />
      </ProFormText>
      <ProFormText
        x-if={props.receiverType === MsgTemplateReceiverType.Email}
        name="attachmentsFileIds"
        label={t('attachments')}
        tooltip={t('attachments_tip')}
      >
        <UploadMultiple
          accept=".doc,.docx,.jpg,.jpeg,.png,.pdf"
          forceDirectory
          directory={`/msg/att/${userState.tenantId}`}
          onChangePath={setAttachments}
        />
      </ProFormText>
    </DrawerForm>
  );
};
