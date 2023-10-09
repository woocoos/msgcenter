import { setItem } from '@/pkg/localStore';
import store from '@/store';
import { User } from '@knockout-js/api/ucenter';
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
    avatarFileId: string;
    domains: {
      id: string;
      name: string;
    }[];
  };
}
const ICE_API_AUTH_PREFIX = process.env.ICE_API_AUTH_PREFIX ?? '/api-auth',
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
    }
  }

  return url
}

/**
 * 解析spm信息
 * @returns
 */
export async function parseSpm() {
  const parseData: {
    token?: string;
    refreshToken?: string;
    tenantId?: string;
    user?: User
  } = {}

  const u = new URL(window.location.href), spm = u.searchParams.get('spm');
  parseData.tenantId = u.searchParams.get('tid') ?? undefined;

  if (spm) {
    try {
      // 存放在cookie中避免重复读取
      const ck = `spm=${spm}`;
      if (document.cookie.indexOf(ck) === -1) {
        const result: LoginRes = await request.post(`${ICE_API_AUTH_PREFIX}/spm/auth`, {
          spm,
        });
        if (result?.accessToken) {
          parseData.token = result.accessToken;
          parseData.refreshToken = result.refreshToken;
          if (!parseData.tenantId) {
            parseData.tenantId = result.user?.domains?.[0]?.id
          }
          if (result.user) {
            parseData.user = {
              id: result.user.id,
              displayName: result.user.displayName,
              avatarFileID: result.user.avatarFileId,
            } as User
          }
          setItem('token', parseData.token);
          setItem('refreshToken', parseData.refreshToken);
          setItem('tenantId', parseData.tenantId);
          setItem('user', parseData.user);
        }
        document.cookie = ck
      }
    } catch (error) {
    }
  }
  return parseData
}
