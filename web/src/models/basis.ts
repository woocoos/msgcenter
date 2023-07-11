import { createModel } from 'ice';
import { LoginRes } from '@/services/basis';
import { setItem, removeItem } from '@/pkg/localStore';
import { User } from '@/__generated__/adminx/graphql';

type BasisUserState = {
  id: string;
  displayName: string;
  avatarFileId?: string;
};

type BasisModelState = {
  locale: LocalLanguage;
  token: string;
  refreshToken: string;
  tenantId: string;
  darkMode: boolean;
  compactMode: boolean;
  user: BasisUserState | null;
};

export type LocalLanguage = 'zh-CN' | 'en-US';

export default createModel({
  state: {
    locale: 'zh-CN',
    token: '',
    refreshToken: '',
    tenantId: '',
    user: null,
    darkMode: false,
    compactMode: false,
  } as BasisModelState,
  reducers: {
    updateLocale(prevState: BasisModelState, payload: LocalLanguage) {
      setItem('locale', payload);
      prevState.locale = payload;
    },
    updateToken(prevState: BasisModelState, payload: string) {
      if (payload) {
        setItem('token', payload);
      } else {
        removeItem('token');
      }
      prevState.token = payload;
    },
    updateRefreshToken(prevState: BasisModelState, payload: string) {
      if (payload) {
        setItem('refreshToken', payload);
      } else {
        removeItem('refreshToken');
      }
      prevState.refreshToken = payload;
    },
    updateTenantId(prevState: BasisModelState, payload: string) {
      if (payload) {
        setItem('tenantId', payload);
      } else {
        removeItem('tenantId');
      }
      prevState.tenantId = payload;
    },
    updateUser(prevState: BasisModelState, payload: BasisUserState | null) {
      if (payload) {
        setItem('user', payload);
      } else {
        removeItem('user');
      }
      prevState.user = payload;
    },
    updateDarkMode(prevState: BasisModelState, payload: boolean) {
      setItem('darkMode', payload);
      prevState.darkMode = payload;
    },
  },
  effects: () => ({
    /**
     * 登录
     * @param payload
     */
    async login(payload: LoginRes, rootState: any) {
      if (payload.accessToken) {
        this.updateToken(payload.accessToken);
        if (payload.user) {
          this.saveUser({
            id: payload.user.id,
            displayName: payload.user.displayName,
            avatarFileID: payload.user?.avatarFileId || '',
          } as User);
          if (payload.user.domains?.length) {
            if (!payload.user.domains.find(item => item.id == rootState.basis.tenantId)) {
              this.updateTenantId(payload.user.domains[0].id);
            }
          } else {
            this.updateTenantId('');
          }
        }
        this.updateRefreshToken(payload.refreshToken || '');
      }
    },
    /**
     * 退出
     * @param isHistory
     */
    async logout() {
      this.updateToken('');
      this.updateUser(null);

      if (!location.pathname.split('/').includes('login')) {
        location.href = `/login?redirect=${encodeURIComponent(location.href)}`
      }
    },
    /**
     * 更新用户信息
     * @param user
     */
    async saveUser(user: User) {
      this.updateUser({
        id: user.id,
        displayName: user.displayName,
        avatarFileId: user.avatarFileID || undefined,
      });
    },
    /**
     * 更新租户id
     * @param key
     */
    async saveTenantId(tenantId: string) {
      this.updateTenantId(tenantId);
    },
  }),
});
