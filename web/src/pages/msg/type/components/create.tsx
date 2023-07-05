import { App } from '@/__generated__/adminx/graphql';
import { MsgType, MsgTypeSimpleStatus } from '@/__generated__/msgsrv/graphql';
import InputApp from '@/components/Adminx/App/input';
import { setLeavePromptWhen } from '@/components/LeavePrompt';
import { cacheApp } from '@/services/adminx/app/indtx';
import { EnumMsgTypeStatus, createMsgType, getMsgTypeInfo, updateMsgType } from '@/services/msgsrv/type';
import { updateFormat } from '@/util';
import { DrawerForm, ProFormSelect, ProFormText, ProFormTextArea, ProFormSwitch } from '@ant-design/pro-components';
import { useState } from 'react';
import { useTranslation } from 'react-i18next';

type ProFormData = {
  app?: App;
  category: string;
  name: string;
  canSubs: boolean;
  canCustom: boolean;
  status: MsgTypeSimpleStatus;
  comments?: string | null;
};

export default (props: {
  open: boolean;
  title?: string;
  id?: string | null;
  onClose: (isSuccess?: boolean) => void;
}) => {
  const { t } = useTranslation(),
    [info, setInfo] = useState<MsgType>(),
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
        category: '',
        name: '',
        canSubs: false,
        canCustom: false,
        status: MsgTypeSimpleStatus.Active
      }
      if (props.id) {
        const result = await getMsgTypeInfo(props.id);
        if (result?.id) {
          setInfo(result as MsgType);
          initData.app = result.appID ? cacheApp[result.appID] : undefined;
          initData.category = result.category;
          initData.name = result.name;
          initData.comments = result.comments;
          initData.canSubs = !!result.canSubs;
          initData.canCustom = !!result.canCustom;
          initData.status = result.status || MsgTypeSimpleStatus.Active;
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
        ? await updateMsgType(props.id, updateFormat({
          category: values.category,
          name: values.name,
          appID: values.app?.id ? Number(values.app.id) : undefined,
          canCustom: values.canCustom,
          canSubs: values.canSubs,
          comments: values.comments,
          status: values.status,
        }, info || {}))
        : await createMsgType({
          category: values.category,
          name: values.name,
          appID: values.app?.id ? Number(values.app.id) : undefined,
          canCustom: values.canCustom,
          canSubs: values.canSubs,
          comments: values.comments,
          status: values.status,
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
        name="app"
        label={t('app')}
        rules={[
          { required: true, message: `${t('please_enter_app')}` },
        ]}>
        <InputApp />
      </ProFormText>
      <ProFormText
        name="category"
        label={t('category')}
        rules={[
          { required: true, message: `${t('please_enter_category')}` },
        ]}
      />
      <ProFormText
        name="name"
        label={t('name')}
        rules={[
          { required: true, message: `${t('please_enter_name')}` },
        ]}
      />
      <ProFormSwitch
        name="canSubs"
        label={t('open_subscription')}
      />
      <ProFormSwitch
        name="canCustom"
        label={t('open_custom')}
      />
      <ProFormSelect
        name="status"
        label={t('status')}
        valueEnum={EnumMsgTypeStatus}
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
