import store from '@/store';
import { request } from 'ice';
import jwtDcode, { JwtPayload } from 'jwt-decode';

export interface LoginRes {
  accessToken?: string;
  expiresIn?: number;
  refreshToken?: string;
  stateToken?: string;
  callbackUrl?: string;
  user?: {
    id: string;
    displayName: string;
    avatar: string;
    domains: {
      id: string;
      name: string;
    }[];
  };
}

export type AppDeployConfig = {
  title: string;
  appCode: string;
  entry: string;
  forceTenantId: boolean;
};

const ICE_API_AUTH_PREFIX = process.env.ICE_API_AUTH_PREFIX ?? '/api-auth',
  ICE_APP_DEPLOY_CONFIG = process.env.ICE_APP_DEPLOY_CONFIG ?? '',
  ICE_LOGIN_URL = process.env.ICE_LOGIN_URL ?? '/login';

/**
 * 退出登录
 * @returns
 */
export async function logout() {
  const userState = store.getModelState('user');
  if (userState.token) {
    try {
      request.post(`${ICE_API_AUTH_PREFIX}/logout`);
    } catch (error) { }
  }
  const userDispatcher = store.getModelDispatchers('user')
  userDispatcher.logout();
  if (ICE_LOGIN_URL.toLowerCase().startsWith("http")) {
    const url = new URL(ICE_LOGIN_URL);
    if (location.pathname !== url.pathname || location.host != url.host) {
      location.href = `${ICE_LOGIN_URL}?redirect=${encodeURIComponent(location.href)}`
    }
  } else {
    if (location.pathname !== ICE_LOGIN_URL) {
      location.href = `${ICE_LOGIN_URL}?redirect=${encodeURIComponent(location.href)}`
    }
  }
}


const appDeployConfig: AppDeployConfig[] = [];

/**
 * 获取应用部署配置文件
 * @returns
 */
export async function getAppDeployConfig() {
  if (appDeployConfig.length) {
    return appDeployConfig;
  }
  if (ICE_APP_DEPLOY_CONFIG) {
    try {
      const result = await request.get(`${ICE_APP_DEPLOY_CONFIG}?t=${Date.now()}`) as AppDeployConfig[];
      appDeployConfig.push(...result);
      return appDeployConfig;
    } catch (error) {
    }
  }
  return null;
}


let refreshTokenFn: NodeJS.Timeout;

export function refreshToken() {
  clearTimeout(refreshTokenFn);
  refreshTokenFn = setTimeout(async () => {
    const userState = store.getModelState('user');
    if (userState.token && userState.refreshToken) {
      const jwt = jwtDcode<JwtPayload>(userState.token);
      if ((jwt.exp || 0) * 1000 - Date.now() < 30 * 60 * 1000) {
        // 小于30分钟的时候需要刷新token
        try {

          const tr = await request.post(`${ICE_API_AUTH_PREFIX}/login/refresh-token`, {
            refreshToken: userState.refreshToken,
          });
          if (tr.accessToken) {
            store.dispatch.user.updateToken(tr.accessToken);
          }
        } catch (error) {
        }
      }
    }
  }, 2000);
}


/**
 * 处理url是否需要创建spm
 * @returns
 */
export async function urlSpm(url: string, tenantId?: string) {
  if (url.toLowerCase().startsWith("http")) {
    const u = new URL(url);
    if (u.origin != location.origin) {
      try {
        const result = await request.post(`${ICE_API_AUTH_PREFIX}/spm/create`), userState = store.getModelState("user");
        if (typeof result === 'string') {
          u.searchParams.set('spm', result)
          if (tenantId || userState.tenantId) {
            u.searchParams.set('tid', tenantId || userState.tenantId)
          }
        }
      } catch (error) {
      }
      return u.href
    } else {
      return u.href.replace(u.origin, '')
    }
  }

  return url
}
