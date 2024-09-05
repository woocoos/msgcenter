import { MsgType, MsgTypeSimpleStatus } from '@/generated/msgsrv/graphql';
import { EnumMsgTypeStatus, createMsgType, getMsgTypeInfo, updateMsgType } from '@/services/msgsrv/type';
import { updateFormat } from '@/util';
import { DrawerForm, ProFormSelect, ProFormText, ProFormTextArea, ProFormSwitch } from '@ant-design/pro-components';
import { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { useLeavePrompt } from '@knockout-js/layout';
import { DictSelect } from '@knockout-js/org';
import { AppSelect } from '@knockout-js/org';
import { getApp } from '@knockout-js/api';
import { App } from '@knockout-js/api/ucenter';

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
  onClose: (isSuccess?: boolean, newInfo?: MsgType) => void;
}) => {
  const { t } = useTranslation(),
    [info, setInfo] = useState<MsgType>(),
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
          initData.app = result.appID ? await getApp(result.appID) as App : undefined;
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
          appID: values.app?.id ? values.app.id : undefined,
          canCustom: values.canCustom,
          canSubs: values.canSubs,
          comments: values.comments,
          status: values.status,
        }, info || {}))
        : await createMsgType({
          category: values.category,
          name: values.name,
          appID: values.app?.id ? values.app.id : undefined,
          canCustom: values.canCustom,
          canSubs: values.canSubs,
          comments: values.comments,
          status: values.status,
        });
      if (result?.id) {
        setSaveDisabled(true);
        props.onClose(true, result as MsgType);
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
        name="app"
        label={t('app')}
        rules={[
          { required: true, message: `${t('please_enter_app')}` },
        ]}>
        <AppSelect />
      </ProFormText>
      <ProFormText
        name="category"
        label={t('category')}
        rules={[
          { required: true, message: `${t('please_enter_category')}` },
        ]}
      >
        <DictSelect dictCode="MsgCategory" />
      </ProFormText>
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
