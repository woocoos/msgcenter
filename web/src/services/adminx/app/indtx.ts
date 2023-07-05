import { gql } from "@/__generated__/adminx";
import { App, AppOrder, AppWhereInput } from "@/__generated__/adminx/graphql";
import { pagingRequest, queryRequest } from "..";
import { gid } from "@/util";

export const EnumAppKind = {
  web: { text: 'web' },
  native: { text: 'native' },
  server: { text: 'server' },
};

export const cacheApp: Record<string, App> = {}

const queryAppList = gql(/* GraphQL */`query appList($first: Int,$orderBy:AppOrder,$where:AppWhereInput){
  apps(first:$first,orderBy: $orderBy,where: $where){
    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }
    edges{
      cursor,node{
        id,name,code,kind,redirectURI,appKey,appSecret,scopes,tokenValidity,
        refreshTokenValidity,logo,comments,status,createdAt
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
 * 获取应用信息
 * @param params
 * @param filter
 * @param sort
 * @returns
 */
export async function getAppList(
  gather: {
    current?: number;
    pageSize?: number;
    where?: AppWhereInput;
    orderBy?: AppOrder;
  },
) {
  const
    result = await pagingRequest(
      queryAppList, {
      first: gather.pageSize || 20,
      where: gather.where,
      orderBy: gather.orderBy,
    }, gather.current || 1);

  if (result.data?.apps) {
    return result.data.apps;
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
