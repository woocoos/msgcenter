import { defineAppConfig, defineDataLoader } from 'ice';
import { defineAuthConfig } from '@ice/plugin-auth/esm/types';
import { defineStoreConfig } from '@ice/plugin-store/esm/types';
import { defineRequestConfig } from '@ice/plugin-request/esm/types';
import { message } from 'antd';
import store from '@/store';
import '@/assets/styles/index.css';
import { getItem, setItem } from '@/pkg/localStore';
import { userPermissions } from './services/adminx/user';
import { browserLanguage } from './util';
import jwtDcode, { JwtPayload } from 'jwt-decode';
import i18n from './i18n';
import { User } from './__generated__/adminx/graphql';
import { defineChildConfig } from '@ice/plugin-icestark/types';

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
  let locale = getItem<string>('locale'),
    token = getItem<string>('token'),
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
    basis: {
      locale,
      token,
      darkMode,
      compactMode,
      tenantId,
      user,
    },
  };
});

// 权限
export const authConfig = defineAuthConfig(async (appData) => {
  const { basis } = appData,
    initialAuth = {};
  // 判断路由权限
  if (basis.token) {
    const ups = await userPermissions({}, {
      Authorization: `Bearer ${basis.token}`,
      'X-Tenant-ID': basis.tenantId,
    });
    if (ups) {
      ups.forEach(item => {
        if (item) {
          initialAuth[item.name] = true;
        }
      });
    }
  } else {
    store.dispatch.basis.logout();
  }
  return {
    initialAuth,
  };
});

// store数据项
export const storeConfig = defineStoreConfig(async (appData) => {
  const { basis } = appData;
  return {
    initialStates: {
      basis,
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
            const basisState = store.getModelState('basis');
            if (!config.headers['Authorization'] && basisState.token) {
              config.headers['Authorization'] = `Bearer ${basisState.token}`;
            }
            if (!config.headers['X-Tenant-ID'] && basisState.tenantId) {
              config.headers['X-Tenant-ID'] = basisState.tenantId;
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
                store.dispatch.basis.logout();
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

