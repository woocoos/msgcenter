import { MsgChannel, MsgChannelReceiverType } from '@/__generated__/msgsrv/graphql';
import { cacheOrg } from '@/services/adminx/org';
import { EnumMsgChannelReceiverType, createMsgChannel, getMsgChannelInfo, updateMsgChannel } from '@/services/msgsrv/channel';
import { updateFormat } from '@/util';
import { DrawerForm, ProFormSelect, ProFormText, ProFormTextArea } from '@ant-design/pro-components';
import { Org, OrgKind } from '@knockout-js/api';
import { useLeavePrompt } from '@knockout-js/layout';
import { OrgSelect } from '@knockout-js/org';
import { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';

type ProFormData = {
  org?: Org;
  name: string;
  receiverType?: MsgChannelReceiverType;
  comments?: string | null;
};

export default (props: {
  open: boolean;
  title?: string;
  id?: string | null;
  onClose: (isSuccess?: boolean) => void;
}) => {
  const { t } = useTranslation(),
    [info, setInfo] = useState<MsgChannel>(),
    [, setLeavePromptWhen] = useLeavePrompt(),
    [saveLoading, setSaveLoading] = useState(false),
    [saveDisabled, setSaveDisabled] = useState(true);

  useEffect(() => {
    setLeavePromptWhen(saveDisabled);
  }, [saveDisabled]);

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
      }
      if (props.id) {
        const result = await getMsgChannelInfo(props.id);
        if (result?.id) {
          setInfo(result as MsgChannel);
          initData.org = result.tenantID ? cacheOrg[result.tenantID] : undefined;
          initData.name = result.name;
          initData.comments = result.comments;
          initData.receiverType = result.receiverType;
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
        ? await updateMsgChannel(props.id, updateFormat({
          name: values.name,
          tenantID: values.org?.id ? values.org?.id : '',
          receiverType: values.receiverType || MsgChannelReceiverType.Email,
          comments: values.comments,
        }, info || {}))
        : await createMsgChannel({
          name: values.name,
          tenantID: values.org?.id ? values.org?.id : '',
          receiverType: values.receiverType || MsgChannelReceiverType.Email,
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
        name="org"
        label={t('org')}
        rules={[
          { required: true, message: `${t('please_enter_org')}` },
        ]}>
        <OrgSelect kind={OrgKind.Root} />
      </ProFormText>
      <ProFormText
        name="name"
        label={t('name')}
        rules={[
          { required: true, message: `${t('please_enter_name')}` },
        ]}
      />
      <ProFormSelect
        name="receiverType"
        label={t('type')}
        valueEnum={EnumMsgChannelReceiverType}
        rules={[
          { required: true, message: `${t('please_enter_status')}` },
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
