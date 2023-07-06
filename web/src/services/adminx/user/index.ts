import { queryRequest } from '../';
import { gql } from '@/__generated__/adminx';
import { AppActionKind, AppActionMethod, AppActionWhereInput, User } from '@/__generated__/adminx/graphql';
import { gid } from '@/util';

export const cacheUser: Record<string, User> = {}

const queryUserIdList = gql(/* GraphQL */`query userIdList($ids:[GID!]!){
  nodes(ids: $ids){
    ... on User{
      id,displayName
    }
  }
}`)

const queryCheckPermission = gql(/* GraphQL */`query  checkPermission($permission:String!){
  checkPermission(permission: $permission)
}`);

const queryUserPermissionList = gql(/* GraphQL */`query userPermissionList($where: AppActionWhereInput){
  userPermissions(where: $where){
    id,appID,name,kind,method
  }
}`);

const queryUserMenuList = gql(/* GraphQL */`query userMenuList($appCode:String!){
  userMenus(appCode: $appCode){
    id,parentID,kind,name,comments,displaySort,icon,route
  }
}`);

const queryUserRootOrgList = gql(/* GraphQL */`query userRootOrgs{
  userRootOrgs{
    id,parentID,code,name,displaySort,countryCode,timezone
  }
}`);

/**
 * 检测权限
 * @param permission  'appCode:action'
 * @returns
 */
export async function checkPermission(permission: string) {
  const result = await queryRequest(
    queryCheckPermission, {
    permission,
  });

  if (result.data?.checkPermission) {
    return result?.data?.checkPermission;
  }
  return null;
}


/**
 * 获取用户的权限
 * @param headers
 * @returns
 */
export async function userPermissions(headers?: Record<string, any>) {
  const result = await queryRequest(
    queryUserPermissionList,
    {
      where: {
        hasAppWith: [{ code: process.env.ICE_APP_CODE }],
        or: [
          { kind: AppActionKind.Function },
          { kindNEQ: AppActionKind.Function, method: AppActionMethod.Write }
        ],
      }
    },
    {
      fetchOptions: { headers }
    }
  );

  if (result.data?.userPermissions) {
    return result?.data?.userPermissions;
  }
  return null;
}

/**
 * 获取用户授权的菜单
 * @param appCode
 * @returns
 */
export async function userMenus(appCode: string) {
  const result = await queryRequest(
    queryUserMenuList, {
    appCode,
  });

  if (result.data?.userMenus) {
    return result?.data?.userMenus;
  }
  return null;
}

/**
 * 获取用户root组织
 * @returns
 */
export async function userRootOrgs() {
  const result = await queryRequest(queryUserRootOrgList, {});
  if (result.data?.userRootOrgs) {
    return result?.data?.userRootOrgs;
  }
  return null;
}

/**
 * 缓存user值
 * @param ids
 */
export async function updateCacheUserListByIds(ids: (string | number)[]) {
  const cacheIds = Object.keys(cacheUser)
  const newCacheIds = ids.filter(id => !cacheIds.includes(`${id}`))
  if (newCacheIds.length) {
    const result = await queryRequest(queryUserIdList, {
      ids: newCacheIds.map(id => gid('user', id))
    })
    result.data?.nodes?.forEach(item => {
      if (item?.__typename === 'User') {
        cacheUser[item.id] = item as User
      }
    })
  }
}
