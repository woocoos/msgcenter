import { setLeavePromptWhen } from '@/components/LeavePrompt';
import { DrawerForm, ProFormText } from '@ant-design/pro-components';
import { useState } from 'react';
import { useTranslation } from 'react-i18next';
import Editor from '@monaco-editor/react';
import { getMsgEventInfoRoute, updateMsgEvent } from '@/services/msgsrv/event';
import { MatchType, Route } from '@/__generated__/msgsrv/graphql';

type RouteStr = {
  matchers?: string[];
  routes: RouteStr[];
  [key: string]: any;
}

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
    parseRoute = (data: RouteStr) => {
      const nData = { ...data } as Route
      if (data.matchers) {
        nData.matchers = data.matchers.map(item => {
          if (item.indexOf('!=') > -1) {
            const itemSplit = item.split('!=')
            return {
              name: itemSplit[0],
              type: MatchType.MatchNotEqual,
              value: itemSplit[1].replaceAll('"', '')
            }
          } else if (item.indexOf('=~') > -1) {
            const itemSplit = item.split('=~')
            return {
              name: itemSplit[0],
              type: MatchType.MatchRegexp,
              value: itemSplit[1].replaceAll('"', '')
            }
          } else if (item.indexOf('!~') > -1) {
            const itemSplit = item.split('!~')
            return {
              name: itemSplit[0],
              type: MatchType.MatchNotRegexp,
              value: itemSplit[1].replaceAll('"', '')
            }
          } else {
            const itemSplit = item.split('=')
            return {
              name: itemSplit[0],
              type: MatchType.MatchEqual,
              value: itemSplit[1].replaceAll('"', '')
            }
          }
        })
      }
      if (data.routes) {
        nData.routes = data.routes.map(route => parseRoute(route))
      }
      return nData;
    },
    getRequest = async () => {
      setSaveLoading(false);
      setSaveDisabled(true);
      const initData: ProFormData = {
        route: '{}'
      }
      const result = await getMsgEventInfoRoute(props.id);
      if (result?.id && result.routeStr) {
        try {
          initData.route = JSON.stringify(parseRoute(JSON.parse(result.routeStr)), null, 4)
        } catch (error) {
        }
      }
      return initData;
    },
    onValuesChange = () => {
      setSaveDisabled(false);
    },
    onFinish = async (values: ProFormData) => {
      setSaveLoading(true);

      const result = await updateMsgEvent(props.id, {
        route: JSON.parse(values.route),
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
      <ProFormText name="route">
        <Editor
          className="adminx-editor"
          height="70vh"
          defaultLanguage="json"
        />
      </ProFormText>
    </DrawerForm>
  );
};
