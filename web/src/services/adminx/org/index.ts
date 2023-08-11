import { gql } from '@/__generated__/adminx';
import { Org, gid } from '@knockout-js/api';
import { query } from '@knockout-js/ice-urql/request';

export const EnumOrgStatus = {
  active: { text: '活跃', status: 'success' },
  inactive: { text: '失活', status: 'default' },
  processing: { text: '处理中', status: 'warning' },
};

export const cacheOrg: Record<string, Org> = {}

const queryOrgIdList = gql(/* GraphQL */`query orgIdList($ids:[GID!]!){
  nodes(ids: $ids){
    ... on Org{
      id,code,name
    }
  }
}`)


/**
 * 缓存org值
 * @param ids
 */
export async function updateCacheOrgListByIds(ids: (string | number)[]) {
  const cacheIds = Object.keys(cacheOrg)
  const newCacheIds = ids.filter(id => !cacheIds.includes(`${id}`))
  if (newCacheIds.length) {
    const result = await query(queryOrgIdList, {
      ids: newCacheIds.map(id => gid('org', id))
    }, {
      instanceName: 'ucenter',
    })
    result.data?.nodes?.forEach(item => {
      if (item?.__typename === 'Org') {
        cacheOrg[item.id] = item as Org
      }
    })
  }
}
