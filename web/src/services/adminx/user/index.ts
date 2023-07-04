import { queryRequest } from '../';
import { gql } from '@/__generated__/adminx';
import { AppActionWhereInput } from '@/__generated__/adminx/graphql';

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
    id,parentID,kind,domain,code,name,status,path,displaySort,countryCode,timezone
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
 * @param where
 * @returns
 */
export async function userPermissions(where: AppActionWhereInput, headers?: Record<string, any>) {
  const result = await queryRequest(
    queryUserPermissionList,
    { where },
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

