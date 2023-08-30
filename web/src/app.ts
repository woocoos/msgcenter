import { defineAppConfig, defineDataLoader } from 'ice';
import { defineAuthConfig } from '@ice/plugin-auth/esm/types';
import { defineStoreConfig } from '@ice/plugin-store/esm/types';
import { defineRequestConfig } from '@ice/plugin-request/esm/types';
import { defineUrqlConfig, requestInterceptor } from "@knockout-js/ice-urql/types";
import store from '@/store';
import '@/assets/styles/index.css';
import { getItem, removeItem, setItem } from '@/pkg/localStore';
import { browserLanguage } from './util';
import jwtDcode, { JwtPayload } from 'jwt-decode';
import { defineChildConfig } from '@ice/plugin-icestark/types';
import { isInIcestark } from '@ice/stark-app';
import { User, userPermissions } from '@knockout-js/api';
import { logout, parseSpm } from './services/auth';

const ICE_API_MSGSRV = process.env.ICE_API_MSGSRV ?? '',
  ICE_API_ADMINX = process.env.ICE_API_ADMINX ?? '',
  NODE_ENV = process.env.NODE_ENV ?? '',
  ICE_DEV_TOKEN = process.env.ICE_DEV_TOKEN ?? '',
  ICE_DEV_TID = process.env.ICE_DEV_TID ?? '',
  ICE_APP_CODE = process.env.ICE_APP_CODE ?? '',
  ICE_LOGIN_URL = process.env.ICE_LOGIN_URL ?? '',
  ICE_API_AUTH_PREFIX = process.env.ICE_API_AUTH_PREFIX ?? '';

if (NODE_ENV === 'development') {
  // 无登录项目增加前端缓存内容 方便开发和展示
  setItem('token', ICE_DEV_TOKEN)
  setItem('tenantId', ICE_DEV_TID)
  setItem('user', {
    id: 1,
    displayName: 'admin',
  })
}

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
}));

// 用来做初始化数据
export const dataLoader = defineDataLoader(async () => {
  if (!isInIcestark()) {
    const signCid = `sign_cid=${ICE_APP_CODE}`
    if (document.cookie.indexOf(signCid) === -1) {
      removeItem('token')
      removeItem('refreshToken')
    }
    document.cookie = signCid
  }
  const spmData = await parseSpm();
  let locale = getItem<string>('locale'),
    token = spmData.token ?? getItem<string>('token'),
    refreshToken = spmData.refreshToken ?? getItem<string>('refreshToken'),
    tenantId = spmData.tenantId ?? getItem<string>('tenantId'),
    darkMode = getItem<string>('darkMode'),
    compactMode = getItem<string>('compactMode'),
    user = spmData.user ?? getItem<User>('user');

  if (token) {
    // 增加jwt判断token过期的处理
    try {
      const jwt = jwtDcode<JwtPayload>(token);
      if ((jwt.exp || 0) * 1000 < Date.now()) {
        token = '';
      }
    } catch (err) {
      token = '';
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
            const { token, tenantId, refreshToken } = store.getModelState('user')
            return {
              token, tenantId, refreshToken
            }
          },
          setStateToken: (newToken) => {
            store.dispatch.user.updateToken(newToken)
          }
        },
        login: ICE_LOGIN_URL,
        refreshApi: `${ICE_API_AUTH_PREFIX}/login/refresh-token`
      }
    }
  },
  {
    instanceName: 'ucenter',
    url: ICE_API_ADMINX,
  },
])


// 权限
export const authConfig = defineAuthConfig(async (appData) => {
  const { user } = appData,
    initialAuth = {};
  // 判断路由权限
  if (user.token) {
    const result = await userPermissions(ICE_APP_CODE, {
      Authorization: `Bearer ${user.token}`,
      'X-Tenant-ID': user.tenantId,
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
  };
});

// store数据项
export const storeConfig = defineStoreConfig(async (appData) => {
  const { user, app } = appData;
  return {
    initialStates: {
      user,
      app,
    },
  };
});


// 请求配置
export const requestConfig = defineRequestConfig({
  interceptors: requestInterceptor({
    store: {
      getState: () => {
        const { token, tenantId } = store.getModelState('user')
        return {
          token, tenantId
        }
      },
    },
    login: ICE_LOGIN_URL,
  })
});

