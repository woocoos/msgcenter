import { gql } from '@/__generated__/adminx';
import { gid } from '@/util';
import { pagingRequest, queryRequest } from '../';
import { App, AppOrder, AppWhereInput } from '@/__generated__/adminx/graphql';

export const EnumAppKind = {
  web: { text: 'web' },
  native: { text: 'native' },
  server: { text: 'server' },
};

export const cacheApp: Record<string, App> = {}

const queryOrgAppList = gql(/* GraphQL */`query orgAppList($gid: GID!,$first: Int,$orderBy:AppOrder,$where:AppWhereInput){
  node(id:$gid){
    ... on Org{
      id
      apps(first:$first,orderBy: $orderBy,where: $where){
        totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }
        edges{
          cursor,node{
            id,name,code,kind,redirectURI,appKey,appSecret,scopes,
            tokenValidity,refreshTokenValidity,logo,comments,status,createdAt
          }
        }
      }
    }
  }
}`);

const queryAppIdList = gql(/* GraphQL */`query appIdList($ids:[GID!]!){
  nodes(ids: $ids){
    ... on App{
      id,code,name
    }
  }
}`)

/**
 * 组织下的应用
 * @param orgId
 * @param params
 * @param filter
 * @param sort
 * @returns
 */
export async function getOrgAppList(
  orgId: string,
  gather: {
    current?: number;
    pageSize?: number;
    where?: AppWhereInput;
    orderBy?: AppOrder;
  }) {
  const
    result = await pagingRequest(
      queryOrgAppList, {
      gid: gid('org', orgId),
      first: gather.pageSize || 20,
      where: gather.where,
      orderBy: gather.orderBy,
    }, gather.current || 1);

  if (result.data?.node?.__typename === 'Org') {
    return result.data.node.apps;
  }
  return null;
}

/**
 * 缓存app值
 * @param ids
 */
export async function updateCacheAppListByIds(ids: (string | number)[]) {
  const cacheIds = Object.keys(cacheApp)
  const newCacheIds = ids.filter(id => !cacheIds.includes(`${id}`))
  if (newCacheIds.length) {
    const result = await queryRequest(queryAppIdList, {
      ids: newCacheIds.map(id => gid('app', id))
    })
    result.data?.nodes?.forEach(item => {
      if (item?.__typename === 'App') {
        cacheApp[item.id] = item as App
      }
    })
  }
}
