import { MatchType, MatcherInput, } from '@/__generated__/msgsrv/graphql';
import { ModalForm, ProFormSelect, ProFormText } from '@ant-design/pro-components';
import { useState } from 'react';
import { useTranslation } from 'react-i18next';
import { EnumSilenceMatchType } from '@/services/msgsrv/silence';



export default (props: {
  open: boolean;
  title?: string;
  data?: MatcherInput;
  onClose: (data?: MatcherInput) => void;
}) => {
  const { t } = useTranslation(),
    [saveLoading, setSaveLoading] = useState(false),
    [saveDisabled, setSaveDisabled] = useState(true);

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
      const initData: MatcherInput = {
        type: MatchType.MatchEqual,
        name: '',
        value: ''
      }
      if (props.data) {
        initData.name = props.data.name;
        initData.type = props.data.type;
        initData.value = props.data.value;
      }
      return initData;
    },
    onValuesChange = () => {
      setSaveDisabled(false);
    },
    onFinish = async (values: MatcherInput) => {
      setSaveLoading(true);

      setSaveDisabled(true);
      props.onClose(values);
      setSaveLoading(false);
      return false;
    };

  return (
    <ModalForm<MatcherInput>
      modalProps={{
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
        name="name"
        label={t('match_name')}
        rules={[
          { required: true, message: `${t('please_enter_match_name')}` },
        ]}>
      </ProFormText>
      <ProFormSelect
        name="type"
        label={t('match_type')}
        rules={[
          { required: true, message: `${t('please_enter_match_type')}` },
        ]}
        valueEnum={EnumSilenceMatchType}
      />

      <ProFormText
        name="value"
        label={t('match_value')}
        rules={[
          { required: true, message: `${t('please_enter_match_value')}` },
        ]}>
      </ProFormText>

    </ModalForm>
  );
};
