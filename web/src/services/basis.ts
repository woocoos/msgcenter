import { request } from 'ice';

export interface LoginRes {
  accessToken?: string;
  expiresIn?: number;
  refreshToken?: string;
  stateToken?: string;
  callbackUrl?: string;
  user?: {
    id: string;
    displayName: string;
    domains: {
      id: string;
      name: string;
    }[];
  };
}

const baseURL = "/api-auth"

/**
 * 退出登录
 * @returns
 */
export async function logout() {
  return await request.post(`${baseURL}/logout`);
}
