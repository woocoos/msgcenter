import { gql } from "@/__generated__/adminx";
import { App, gid } from "@knockout-js/api";
import { query } from "@knockout-js/ice-urql/request";

export const cacheApp: Record<string, App> = {}

const queryAppIdList = gql(/* GraphQL */`query appIdList($ids:[GID!]!){
  nodes(ids: $ids){
    ... on App{
      id,code,name
    }
  }
}`)

/**
 * 缓存app值
 * @param ids
 */
export async function updateCacheAppListByIds(ids: (string | number)[]) {
  const cacheIds = Object.keys(cacheApp)
  const newCacheIds = ids.filter(id => !cacheIds.includes(`${id}`))
  if (newCacheIds.length) {
    const result = await query(queryAppIdList, {
      ids: newCacheIds.map(id => gid('app', id))
    }, {
      instanceName: 'ucenter',
    })
    result.data?.nodes?.forEach(item => {
      if (item?.__typename === 'App') {
        cacheApp[item.id] = item as App
      }
    })
  }
}
