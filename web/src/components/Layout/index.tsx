import store from '@/store';
import { useEffect, useState } from 'react';
import { userMenuList } from './menuConfig';
import AvatarDropdown from '@/components/Header/AvatarDropdown';
import I18nDropdown from '@/components/Header/I18nDropdown';
import DarkMode from '@/components/Header/DarkMode';
import styles from './layout.module.css';
import logo from '@/assets/images/woocoo.png';
import defaultAvatar from '@/assets/images/default-avatar.png';
import { Outlet } from '@ice/runtime';
import i18n, { CurrentLanguages } from '@/i18n';
import { ProLayout, ProConfigProvider, useToken } from '@ant-design/pro-components';
import LeavePrompt, { Link } from '@/components/LeavePrompt';
import { AliveScope } from 'react-activation';
import TenantDropdown from '@/components/Header/TenantDropdown';
import { monitorKeyChange } from '@/pkg/localStore';
import { getFilesRaw } from '@/services/files';
import zhCN from 'antd/locale/zh_CN';
import enUS from 'antd/locale/en_US';
import { Locale } from 'antd/es/locale';
import { ConfigProvider } from 'antd';
import { EnterOutlined, NodeExpandOutlined } from '@ant-design/icons';

export default () => {
  const [userState, userDispatcher] = store.useModel('user'),
    [appState, appDispatcher] = store.useModel('app'),
    [locale, setLocale] = useState<Locale>(),
    { token } = useToken(),
    [avatar, setAvatar] = useState<string>();

  useEffect(() => {
    if (userState.user?.avatarFileId) {
      getFilesRaw(userState.user?.avatarFileId, 'url').then(result => {
        if (typeof result === 'string') {
          setAvatar(result);
        }
      })
    }
  }, [userState.user]);

  useEffect(() => {
    i18n.changeLanguage(appState.locale);
    switch (appState.locale) {
      case CurrentLanguages.zhCN:
        setLocale(zhCN)
        break;
      case CurrentLanguages.enUS:
        setLocale(enUS)
        break;
      default:
        setLocale(zhCN)
        break;
    }
  }, [appState.locale]);


  useEffect(() => {
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
    ]);
  }, []);


  return <ProConfigProvider dark={appState.darkMode} >
    <ConfigProvider locale={locale}>
      <LeavePrompt />
      <ProLayout
        token={{
          sider: {
            colorMenuBackground: appState.darkMode ? 'linear-gradient(#141414, #000000 28%)' : token.colorBgContainer,
          },
        }}
        className={styles.layout}
        menu={{
          locale: true,
          request: userMenuList,
        }}
        fixSiderbar
        locale={appState.locale}
        logo={<img src={logo} alt="logo" />}
        title="Adminx"
        location={{
          pathname: location.pathname,
        }}
        layout="mix"
        actionsRender={() => [
          <I18nDropdown />,
          <TenantDropdown />,
          <AvatarDropdown
            avatar={avatar || defaultAvatar}
            name={userState.user?.displayName || ''}
          />,
          <DarkMode />,
        ]}
        menuItemRender={(item, defaultDom) => (item.path ? <>
          <Link to={item.path}>{defaultDom}</Link>
          <NodeExpandOutlined className={styles.menuIconPopup} onClick={() => {
            window.open(item.path)
          }} />
        </> : defaultDom)}
      >
        <AliveScope>
          <Outlet />
        </AliveScope>
      </ProLayout>
    </ConfigProvider>
  </ProConfigProvider>
}
