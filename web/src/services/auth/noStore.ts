import { getItem, setItem } from '@/pkg/localStore';
import { User } from '@knockout-js/api/ucenter';
import { request } from 'ice';
import { LoginRes } from '.';

const ICE_API_AUTH_PREFIX = process.env.ICE_API_AUTH_PREFIX ?? '/api-auth'

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
  parseData.tenantId = u.searchParams.get('tid') ?? getItem<string>('tenantId') ?? '';

  if (spm) {
    try {
      // 存放在cookie中避免重复读取
      const result: LoginRes = await request.post(`${ICE_API_AUTH_PREFIX}/spm/auth`, {
        spm,
      });
      if (result?.accessToken) {
        parseData.token = result.accessToken;
        parseData.refreshToken = result.refreshToken;
        if (!result.user?.domains?.find(item => item.id == parseData.tenantId)) {
          parseData.tenantId = `${result.user?.domains?.[0]?.id}`
        }
        if (result.user) {
          parseData.user = {
            id: result.user.id,
            displayName: result.user.displayName,
            avatar: result.user.avatar,
          } as User
        }
        setItem('token', parseData.token);
        setItem('refreshToken', parseData.refreshToken);
        setItem('tenantId', parseData.tenantId);
        setItem('user', parseData.user);
      }

    } catch (error) {
      console.error('parseSpm', error)
    }
    u.searchParams.delete('spm')
    u.searchParams.delete('tid')
    location.replace(u)
  }
}
