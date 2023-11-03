import { defineAppConfig, defineDataLoader } from 'ice';
import { defineAuthConfig } from '@ice/plugin-auth/esm/types';
import { defineStoreConfig } from '@ice/plugin-store/esm/types';
import { defineRequestConfig } from '@ice/plugin-request/esm/types';
import { defineUrqlConfig, requestInterceptor } from "@knockout-js/ice-urql/types";
import store from '@/store';
import '@/assets/styles/index.css';
import { getItem, removeItem, setItem } from '@/pkg/localStore';
import { browserLanguage, getMenuAppActions } from './util';
import jwtDcode, { JwtPayload } from 'jwt-decode';
import { defineChildConfig } from '@ice/plugin-icestark/types';
import { instanceName, setFilesApi, userPermissions } from '@knockout-js/api';
import { logout } from './services/auth';
import { User } from '@knockout-js/api/ucenter';
import { Message } from './generated/msgsrv/graphql';
import { RequestHeaderAuthorizationMode, getRequestHeaderAuthorization } from '@knockout-js/ice-urql/request';
import { useTranslation } from 'react-i18next';
import { Result, message } from 'antd';
import { parseSpm } from './services/auth/noStore';

const NODE_ENV = process.env.NODE_ENV ?? '',
  ICE_API_MSGSRV = process.env.ICE_API_MSGSRV ?? '',
  ICE_ROUTER_BASENAME = process.env.ICE_ROUTER_BASENAME ?? '/',
  ICE_API_ADMINX = process.env.ICE_API_ADMINX ?? '',
  ICE_APP_CODE = process.env.ICE_APP_CODE ?? '',
  ICE_LOGIN_URL = process.env.ICE_LOGIN_URL ?? '',
  ICE_HTTP_SIGN = process.env.ICE_HTTP_SIGN ?? '',
  ICE_API_AUTH_PREFIX = process.env.ICE_API_AUTH_PREFIX ?? '',
  ICE_WS_MSGSRV = process.env.ICE_WS_MSGSRV ?? '',
  ICE_API_FILES_PREFIX = process.env.ICE_API_FILES_PREFIX ?? '';

export const icestark = defineChildConfig(() => ({
  mount: (data) => {
    // 在微应用挂载前执行
    if (data?.customProps) {
      setItem('locale', data.customProps.app.locale);
      setItem('darkMode', data.customProps.app.darkMode);
      setItem('compactMode', data.customProps.app.compactMode);
      setItem('token', data.customProps.user.token);
      setItem('refreshToken', data.customProps.user.refreshToken);
      setItem('tenantId', data.customProps.user.tenantId);
      setItem('user', data.customProps.user.user);
    }
  },
  unmount: () => {
    // 在微应用卸载后执行
    removeItem('token');
    removeItem('refreshToken');
    removeItem('tenantId');
  },
}));

// App config, see https://v3.ice.work/docs/guide/basic/app
export default defineAppConfig(() => ({
  // Set your configs here.
  app: {
    rootId: 'app',
  },
  router: {
    basename: ICE_ROUTER_BASENAME,
  }
}));

// 用来做初始化数据
export const dataLoader = defineDataLoader(async () => {
  if (NODE_ENV === 'development') {
    // 开发时使用
    setItem('token', process.env.ICE_DEV_TOKEN)
    setItem('tenantId', process.env.ICE_DEV_TID)
    setItem('user', {
      id: 1,
      displayName: 'admin',
    })
  }

  setFilesApi(ICE_API_FILES_PREFIX);

  const signCid = `sign_cid=${ICE_APP_CODE}`;
  if (document.cookie.indexOf(signCid) === -1) {
    removeItem('token');
    removeItem('refreshToken');
  }
  document.cookie = signCid;
  await parseSpm();

  let locale = getItem<string>('locale'),
    darkMode = getItem<string>('darkMode'),
    compactMode = getItem<string>('compactMode'),
    handshake = getItem<boolean>('handshake'),
    message = getItem<Message[]>('message') ?? [],
    token = getItem<string>('token'),
    refreshToken = getItem<string>('refreshToken'),
    tenantId = getItem<string>('tenantId'),
    user = getItem<User>('user');

  if (token) {
    // 增加jwt判断token过期的处理
    try {
      const jwt = jwtDcode<JwtPayload>(token);
      if ((jwt.exp || 0) * 1000 < Date.now()) {
        token = '';
        removeItem('token');
      }
    } catch (err) {
      token = '';
      removeItem('token');
    }
  }
  if (!locale) {
    locale = browserLanguage();
  }

  return {
    app: {
      locale,
      darkMode,
      compactMode,
    },
    user: {
      token,
      refreshToken,
      tenantId,
      user,
    },
    ws: {
      handshake,
      message,
    }
  };
});


// urql
export const urqlConfig = defineUrqlConfig([
  {
    instanceName: 'default',
    url: ICE_API_MSGSRV,
    exchangeOpt: {
      authOpts: {
        store: {
          getState: () => {
            const userState = store.getModelState('user'),
              token = userState.token ? userState.token : getItem<string>('token') as string,
              tenantId = userState.tenantId ? userState.tenantId : getItem<string>('tenantId') as string,
              refreshToken = userState.refreshToken ? userState.refreshToken : getItem<string>('refreshToken') as string;

            return {
              token: token,
              tenantId: tenantId,
              refreshToken: refreshToken,
            }
          },
          setStateToken: (newToken) => {
            store.dispatch.user.updateToken(newToken);
          }
        },
        error: (err, errstr) => {
          if (errstr) {
            message.error(errstr)
          }
          return false;
        },
        beforeRefreshTime: 5 * 60 * 1000,
        headerMode: ICE_HTTP_SIGN === 'ko' ? RequestHeaderAuthorizationMode.KO : undefined,
        login: ICE_LOGIN_URL,
        refreshApi: `${ICE_API_AUTH_PREFIX}/login/refresh-token`
      },
      subOpts: {
        url: ICE_WS_MSGSRV,
        store: {
          getState: () => {
            const userState = store.getModelState('user'),
              token = userState.token ? userState.token : getItem<string>('token') as string,
              tenantId = userState.tenantId ? userState.tenantId : getItem<string>('tenantId') as string;

            return {
              token: token,
              tenantId: tenantId,
            };
          },
        }
      }
    }
  },
  {
    instanceName: instanceName.UCENTER,
    url: ICE_API_ADMINX,
  },
])


// 权限
export const authConfig = defineAuthConfig(async (appData) => {
  const initialAuth = getMenuAppActions(),
    token = appData?.user?.token ? appData.user.token : getItem<string>('token'),
    tenantId = appData?.user?.tenantId ? appData.user.tenantId : getItem<string>('tenantId');
  // 判断路由权限
  if (appData?.user?.token) {
    const result = await userPermissions(ICE_APP_CODE, {
      Authorization: getRequestHeaderAuthorization(token, ICE_HTTP_SIGN === 'ko' ? RequestHeaderAuthorizationMode.KO : undefined),
      'X-Tenant-ID': tenantId,
    });
    if (result) {
      result.forEach(item => {
        if (item) {
          initialAuth[item.name] = true;
        }
      });
    }
  } else {
    await logout();
  }
  return {
    initialAuth,
    NoAuthFallback: () => {
      const { t } = useTranslation()
      return (
        <Result status="403"
          title="403"
          subTitle={t('page_403')} />
      )
    }
  };
});

// store数据项
export const storeConfig = defineStoreConfig(async (appData) => {
  return {
    initialStates: {
      app: appData?.app,
      user: appData?.user,
      ws: appData?.ws,
    },
  };
});


// 请求配置
export const requestConfig = defineRequestConfig({
  interceptors: requestInterceptor({
    store: {
      getState: () => {
        const token = getItem<string>('token') as string,
          tenantId = getItem<string>('tenantId') as string;
        return {
          token: token,
          tenantId: tenantId,
        }
      },
    },
    headerMode: ICE_HTTP_SIGN === 'ko' ? RequestHeaderAuthorizationMode.KO : undefined,
    login: ICE_LOGIN_URL,
    error: (err, str) => {
      if (str) {
        window.antd.message.error(str)
      }
    }
  })
});

