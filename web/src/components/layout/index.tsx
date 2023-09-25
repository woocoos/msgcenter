import store from '@/store';
import { useEffect, useRef, useState } from 'react';
import menuList from './menu.json';
import { history } from 'ice';
import { Outlet, useLocation } from '@ice/runtime';
import i18n from '@/i18n';
import { MenuDataItem, useToken } from '@ant-design/pro-components';
import { monitorKeyChange } from '@/pkg/localStore';
import { Layout, useLeavePrompt } from '@knockout-js/layout';
import { logout, urlSpm } from '@/services/auth';
import defaultAvatar from '@/assets/images/default-avatar.png';
import { createFromIconfontCN } from '@ant-design/icons';
import { getFilesRaw } from '@knockout-js/api';
import FloatMsg, { WsMsgViewActions } from '../floatMsg';
import { MsgDropdownRef } from '@knockout-js/layout/esm/components/msg-dropdown';

const ICE_APP_CODE = process.env.ICE_APP_CODE ?? '',
  ICE_WS_MSGSRV = process.env.ICE_WS_MSGSRV ?? '',
  NODE_ENV = process.env.NODE_ENV ?? '',
  IconFont = createFromIconfontCN({
    scriptUrl: "//at.alicdn.com/t/c/font_4214307_8x56lkek9tu.js"
  });

export default () => {
  const [userState, userDispatcher] = store.useModel('user'),
    [appState, appDispatcher] = store.useModel('app'),
    [, wsDispatcher] = store.useModel('ws'),
    msgRef = useRef<MsgDropdownRef>(null),
    [checkLeave] = useLeavePrompt(),
    location = useLocation(),
    { token } = useToken(),
    [avatar, setAvatar] = useState<string>();

  useEffect(() => {
    i18n.changeLanguage(appState.locale);
  }, [appState.locale]);


  useEffect(() => {
    if (userState.user?.avatarFileId) {
      getFilesRaw(userState.user?.avatarFileId, 'url').then(result => {
        if (typeof result === 'string') {
          setAvatar(result);
        }
      })
    }

    monitorKeyChange([
      {
        key: 'tenantId',
        onChange(value) {
          userDispatcher.updateTenantId(value);
        },
      },
      {
        key: 'token',
        onChange(value) {
          userDispatcher.updateToken(value);
        },
      },
      {
        key: 'user',
        onChange(value) {
          userDispatcher.updateUser(value);
        },
      },
      {
        key: 'locale',
        onChange(value) {
          appDispatcher.updateLocale(value);
        },
      },
      {
        key: 'handshake',
        onChange(value) {
          wsDispatcher.setHandshake(value);
        },
      },
      {
        key: 'message',
        onChange(value) {
          wsDispatcher.setMessage(value);
        },
      },
    ]);
  }, []);

  return <Layout
    msgRef={msgRef}
    appCode={ICE_APP_CODE}
    pathname={location.pathname}
    IconFont={IconFont}
    onClickMenuItem={async (item, isOpen) => {
      if (checkLeave()) {
        const url = await urlSpm(item.path ?? '')
        if (isOpen) {
          window.open(url);
        } else {
          history?.push(url);
        }
      }
    }}
    tenantProps={{
      value: userState.tenantId,
      onChange: (value) => {
        userDispatcher.saveTenantId(value);
      },
    }}
    i18nProps={{
      onChange: (value) => {
        appDispatcher.updateLocale(value);
      },
    }}
    avatarProps={{
      avatar: avatar || defaultAvatar,
      name: userState.user?.displayName,
      onLogoutClick: () => {
        if (checkLeave()) {
          logout();
        }
      },
    }}
    msgProps={{
      onItemClick: (data) => {
        window.open(`/msg/internal/detail?toid=${data.id}`);
      },
      onMoreClick: () => {
        window.open(`/msg/internal`);
      },
    }}
    themeSwitchProps={{
      value: appState.darkMode,
      onChange: (value) => {
        appDispatcher.updateDarkMode(value);
      },
    }}
    proLayoutProps={{
      token: {
        sider: {
          colorMenuBackground: appState.darkMode ? 'linear-gradient(#141414, #000000 28%)' : token.colorBgContainer,
        },
      },
      title: 'Msgsrv',
      [NODE_ENV === 'development' ? 'menu' : '']: {
        request: () => {
          const list: MenuDataItem[] = [];
          menuList.forEach(item => {
            const menuItem: MenuDataItem = { name: item.name };
            if (item.icon) {
              menuItem.icon = <IconFont type={item.icon} />
            }
            if (item.children) {
              menuItem.children = item.children
            }
            list.push(menuItem)
          })
          return list
        }
      }
    }}
  >
    <Outlet />
    {ICE_WS_MSGSRV ? <FloatMsg
      onListenerNewMsg={() => {
        msgRef.current?.setShowDot();
      }}
      onItemClick={(data) => {
        if (data.extras.action === WsMsgViewActions.Internal) {
          window.open(`/msg/internal/detail?id=${data.extras.actionID}`);
        }
      }}
    /> : <></>}
  </Layout>
}
