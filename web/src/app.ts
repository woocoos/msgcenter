import { defineAppConfig, defineDataLoader } from 'ice';
import { defineAuthConfig } from '@ice/plugin-auth/esm/types';
import { defineStoreConfig } from '@ice/plugin-store/esm/types';
import { defineRequestConfig } from '@ice/plugin-request/esm/types';
import { defineUrqlConfig } from "@knockout-js/ice-urql/esm/types";
import { message } from 'antd';
import store from '@/store';
import '@/assets/styles/index.css';
import { getItem, removeItem, setItem } from '@/pkg/localStore';
import { userPermissions } from './services/adminx/user';
import { browserLanguage, goLogin } from './util';
import jwtDcode, { JwtPayload } from 'jwt-decode';
import i18n from './i18n';
import { defineChildConfig } from '@ice/plugin-icestark/types';
import { isInIcestark } from '@ice/stark-app';
import { User } from '@knockout-js/api';

export const icestark = defineChildConfig(() => ({
  mount: () => {
    // 在微应用挂载前执行
  },
  unmount: () => {
    // 在微应用卸载后执行
  },
}));

if (process.env.ICE_CORE_MODE === 'development') {
  // 无登录项目增加前端缓存内容 方便开发和展示
  setItem('token', process.env.ICE_TOKEN)
  setItem('tenantId', process.env.ICE_TENANT_ID)
  setItem('user', {
    id: 1,
    displayName: 'admin',
  })
}

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
    const sign = 'sign_cid=adminx'
    if (document.cookie.indexOf(sign) === -1) {
      removeItem('token')
      removeItem('refreshToken')
    }
    document.cookie = sign
  }
  let locale = getItem<string>('locale'),
    token = getItem<string>('token'),
    refreshToken = getItem<string>('refreshToken'),
    darkMode = getItem<string>('darkMode'),
    compactMode = getItem<string>('compactMode'),
    tenantId = getItem<string>('tenantId'),
    user = getItem<User>('user');

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
    url: '/api-msgsrv/graphql/query',
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
        login: process.env.ICE_LOGIN_URL,
        refreshApi: "/api-auth/login/refresh-token"
      }
    }
  },
  {
    instanceName: 'ucenter',
    url: '/api-adminx/graphql/query',
  },
])


// 权限
export const authConfig = defineAuthConfig(async (appData) => {
  const { user } = appData,
    initialAuth = {};
  // 判断路由权限
  if (user.token) {
    const result = await userPermissions({
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
    store.dispatch.user.logout();
    goLogin();
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
export const requestConfig = defineRequestConfig(() => {
  return [
    {
      interceptors: {
        request: {
          onConfig(config: any) {
            const userState = store.getModelState('user');
            if (!config.headers['Authorization'] && userState.token) {
              config.headers['Authorization'] = `Bearer ${userState.token}`;
            }
            if (!config.headers['X-Tenant-ID'] && userState.tenantId) {
              config.headers['X-Tenant-ID'] = userState.tenantId;
            }

            return config;
          },
        },
        response: {
          onConfig(response) {
            if (response.status === 200 && response.data.errors) {
              // 提取第一个异常来展示
              if (response.data.errors?.[0]?.message) {
                message.error(response.data.errors?.[0]?.message);
              }
            }
            return response;
          },
          onError: (error) => {
            const errRes = error.response as any;
            let msg = '';
            if (errRes?.data?.errors?.[0]?.message) {
              msg = errRes?.data?.errors?.[0]?.message;
            }
            switch (errRes.status) {
              case 401:
                store.dispatch.user.logout();
                goLogin();
                if (!msg) {
                  msg = i18n.t('401');
                }
                break;
              case 403:
                if (!msg) {
                  msg = i18n.t('403');
                }
                break;
              case 404:
                if (!msg) {
                  msg = i18n.t('404');
                }
                break;
              case 500:
                if (!msg) {
                  msg = i18n.t('500');
                }
                break;
              default:
                if (!msg) {
                  msg = errRes.statusText;
                }
            }
            if (msg) {
              message.error(msg);
            }
            // 请求出错：服务端返回错误状态码
            return error;
          },
        },
      },
    },
  ];
});

