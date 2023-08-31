import { DrawerForm, ProFormText } from '@ant-design/pro-components';
import { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import Editor from '@monaco-editor/react';
import { getMsgEventInfoRoute, updateMsgEvent } from '@/services/msgsrv/event';
import { MatchType, RouteStrType } from '@/generated/msgsrv/graphql';
import * as yaml from 'js-yaml'
import { Typography } from 'antd';
import { useLeavePrompt } from '@knockout-js/layout';

type ProFormData = {
  route: string;
};

export default (props: {
  open: boolean;
  title?: string;
  id: string;
  onClose: (isSuccess?: boolean) => void;
}) => {
  const { t } = useTranslation(),
    [errStr, setErrStr] = useState<string>(),
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
        route: ``
      }
      const result = await getMsgEventInfoRoute(props.id, RouteStrType.Yaml);
      if (result?.id) {
        initData.route = result.routeStr.
          replaceAll(MatchType.MatchEqual, '=').
          replaceAll(MatchType.MatchNotEqual, '!=').
          replaceAll(MatchType.MatchRegexp, '=~').
          replaceAll(MatchType.MatchNotRegexp, '!~');
      }
      return initData;
    },
    onValuesChange = () => {
      setSaveDisabled(false);
      setErrStr(undefined);
    },
    onFinish = async (values: ProFormData) => {
      setSaveLoading(true);
      try {
        const jsonRoute = yaml.load(values.route, { json: true });
        const route = JSON.parse(JSON.stringify(jsonRoute).
          replaceAll('!=', MatchType.MatchNotEqual).
          replaceAll('=~', MatchType.MatchRegexp).
          replaceAll('!~', MatchType.MatchNotRegexp).
          replaceAll('=', MatchType.MatchEqual));
        const result = await updateMsgEvent(props.id, {
          route,
        });
        if (result?.id) {
          setSaveDisabled(true);
          props.onClose(true);
        }
      } catch (error) {
        setErrStr(error.message)
      }
      setSaveLoading(false);
      return false;
    };

  return (
    <DrawerForm<ProFormData>
      drawerProps={{
        width: 700,
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
      <ProFormText name="route" extra={errStr ? <Typography.Text type="danger">{errStr}</Typography.Text> : <></>}>
        <Editor
          className="adminx-editor"
          height="70vh"
          defaultLanguage="yaml"
        />
      </ProFormText>
    </DrawerForm>
  );
};
