import { MatcherInput, Silence } from '@/__generated__/msgsrv/graphql';
import { setLeavePromptWhen } from '@/components/LeavePrompt';
import { updateFormat } from '@/util';
import { DrawerForm, ProFormDateTimeRangePicker, ProFormText, ProFormTextArea } from '@ant-design/pro-components';
import { useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Org } from '@/__generated__/adminx/graphql';
import { createSilence, getSilenceInfo, updateSilence } from '@/services/msgsrv/silence';
import { cacheOrg } from '@/services/adminx/org';
import InputOrg from '@/components/Adminx/Org/input';
import Matchers from './matchers';

type ProFormData = {
  org?: Org;
  rangeAt?: [string, string];
  matchers: MatcherInput[];
  comments?: string;
};

export default (props: {
  open: boolean;
  title?: string;
  id?: string | null;
  isCopy?: boolean;
  onClose: (isSuccess?: boolean) => void;
}) => {
  const { t } = useTranslation(),
    [info, setInfo] = useState<Silence>(),
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
        matchers: []
      }
      if (props.id) {
        const result = await getSilenceInfo(props.id);
        if (result?.id) {
          setInfo(result as Silence);
          result.matchers?.forEach(item => {
            if (item) {
              initData.matchers.push({
                name: item?.name,
                type: item?.type,
                value: item?.value,
              })
            }
          });
          initData.org = cacheOrg[result.tenantID];
          initData.rangeAt = [result.startsAt, result.endsAt];
          initData.comments = result.comments || undefined;
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
        tenantID: values.org?.id ? Number(values.org.id) : 0,
        startsAt: values.rangeAt?.[0],
        endsAt: values.rangeAt?.[1],
        comments: values.comments,
        matchers: values.matchers,
      };

      const result = props.id && !props.isCopy
        ? await updateSilence(props.id, updateFormat(input, info || {}))
        : await createSilence(input);
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
        <InputOrg />
      </ProFormText>
      <ProFormDateTimeRangePicker
        name="rangeAt"
        label={t('effective_time')}
        fieldProps={{
          style: { width: '100%' }
        }}
        rules={[
          { required: true, message: `${t('please_enter_effective_time')}` },
        ]}
      />
      <ProFormText
        name="matchers"
        label={t('match_msg')}
        rules={[
          { required: true, message: `${t('please_enter_match_msg')}` },
        ]}>
        <Matchers />
      </ProFormText>
      <ProFormTextArea
        name="comments"
        label={t('description')}
        placeholder={`${t('please_enter_description')}`}
      />
    </DrawerForm>
  );
};
