import { getItem, setItem } from '@/pkg/localStore';
import { User } from '@knockout-js/api/ucenter';
import { request } from 'ice';
import { LoginRes } from '.';
import { getI18n } from 'react-i18next';
import { randomId } from '@/util';

const ICE_API_AUTH_PREFIX = process.env.ICE_API_AUTH_PREFIX ?? '/api-auth'
const ICE_API_I18N_PREFIX = process.env.ICE_API_I18N_PREFIX ?? ''

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

/**
 * 多语言文件获取
 */
export const initFillI18n = async () => {
  if (ICE_API_I18N_PREFIX) {
    const i18n = getI18n()
    try {
      const file = await request.get(`${ICE_API_I18N_PREFIX}/${i18n.language}.json?t=${randomId(5)}`)
      if (typeof file === 'object') {
        i18n.addResources(i18n.language, 'translation', file)
      }
    } catch (error) {
      console.error(`${i18n?.language ?? 'i18n'}读取失败！`)
    }
  }
}