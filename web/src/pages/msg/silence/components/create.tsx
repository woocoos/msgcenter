import { MatcherInput, Silence } from '@/__generated__/msgsrv/graphql';
import { setLeavePromptWhen } from '@/components/LeavePrompt';
import { dateRangeTurnDuration, durationTurnEndDate, getDate, updateFormat } from '@/util';
import { DrawerForm, ProFormDateTimeRangePicker, ProFormInstance, ProFormText, ProFormTextArea } from '@ant-design/pro-components';
import { useRef, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Org } from '@/__generated__/adminx/graphql';
import { createSilence, getSilenceInfo, updateSilence } from '@/services/msgsrv/silence';
import { cacheOrg } from '@/services/adminx/org';
import InputOrg from '@/components/Adminx/Org/input';
import Matchers from './matchers';
import { Col, Row } from 'antd';

type ProFormData = {
  org?: Org;
  rangeAt?: [string, string];
  duration?: string;
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
    formRef = useRef<ProFormInstance<ProFormData>>(),
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
          if (result.startsAt && result.endsAt) {
            initData.rangeAt = [getDate(result.startsAt, 'YYYY-MM-DD HH:mm:ss') as string, getDate(result.endsAt, 'YYYY-MM-DD HH:mm:ss') as string];
          } else {
            initData.rangeAt = undefined
          }
          initData.duration = initData.rangeAt
            ? dateRangeTurnDuration(initData.rangeAt) || undefined
            : undefined;
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
        startsAt: getDate(values.rangeAt?.[0], 'YYYY-MM-DDTHH:mm:ssZ'),
        endsAt: getDate(values.rangeAt?.[1], 'YYYY-MM-DDTHH:mm:ssZ'),
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
        width: 580,
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
      formRef={formRef}
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
      <Row gutter={20}>
        <Col span={16}>
          <ProFormDateTimeRangePicker
            name="rangeAt"
            label={t('effective_time')}
            fieldProps={{
              style: { width: '100%' },
              format: 'YYYY-MM-DD HH:mm:ss',
              onChange: (values) => {
                if (values) {
                  formRef.current?.setFieldValue('duration', dateRangeTurnDuration([values[0], values[1]]));
                } else {
                  formRef.current?.setFieldValue('duration', null);
                }
              }
            }}
            rules={[
              { required: true, message: `${t('please_enter_effective_time')}` },
            ]}
          />
        </Col>
        <Col span={8}>
          <ProFormText
            name="duration"
            label={t('duration')}
            rules={[
              { required: true, message: `${t('please_enter_duration')}` },
            ]}
            fieldProps={{
              onBlur: () => {
                const startAt = formRef.current?.getFieldValue('rangeAt');
                if (startAt?.[0]) {
                  const endDate = durationTurnEndDate(startAt[0], formRef.current?.getFieldValue('duration'), 'YYYY-MM-DD HH:mm:ss');
                  if (endDate) {
                    formRef.current?.setFieldValue('rangeAt', [startAt[0], endDate]);
                  }
                }
              }
            }}
          />
        </Col>
      </Row>
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
