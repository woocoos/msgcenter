import { gql } from '@/__generated__/adminx';
import { AppActionKind, AppActionMethod, User, gid } from '@knockout-js/api';
import { query } from '@knockout-js/ice-urql/request';

export const cacheUser: Record<string, User> = {}

const queryUserIdList = gql(/* GraphQL */`query userIdList($ids:[GID!]!){
  nodes(ids: $ids){
    ... on User{
      id,displayName
    }
  }
}`)

const queryUserPermissionList = gql(/* GraphQL */`query userPermissionList($where: AppActionWhereInput){
  userPermissions(where: $where){
    id,appID,name,kind,method
  }
}`);

/**
 * 获取用户的权限
 * @param headers
 * @returns
 */
export async function userPermissions(headers?: Record<string, any>) {
  const result = await query(
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
      instanceName: 'ucenter',
      fetchOptions: { headers },
    }
  );

  if (result.data?.userPermissions) {
    return result?.data?.userPermissions;
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
    const result = await query(queryUserIdList, {
      ids: newCacheIds.map(id => gid('user', id))
    }, {
      instanceName: 'ucenter',
    })
    result.data?.nodes?.forEach(item => {
      if (item?.__typename === 'User') {
        cacheUser[item.id] = item as User
      }
    })
  }
}
