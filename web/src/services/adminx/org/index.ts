import { gid } from '@/util';
import { pagingRequest, queryRequest } from '../';
import { gql } from '@/__generated__/adminx';
import { OrderDirection, Org, OrgOrder, OrgOrderField, OrgWhereInput } from '@/__generated__/adminx/graphql';

export const EnumOrgStatus = {
  active: { text: '活跃', status: 'success' },
  inactive: { text: '失活', status: 'default' },
  processing: { text: '处理中', status: 'warning' },
};

export const cacheOrg: Record<string, Org> = {}

const queryOrgList = gql(/* GraphQL */`query orgList($first: Int,$orderBy:OrgOrder,$where:OrgWhereInput){
  organizations(first:$first,orderBy: $orderBy,where: $where){
    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }
    edges{
      cursor,node{
        id,createdBy,createdAt,updatedBy,updatedAt,deletedAt,ownerID,parentID,kind,
        domain,code,name,profile,status,path,displaySort,countryCode,timezone,
        owner { id,displayName }
      }
    }
  }
}`);

const queryOrgIdList = gql(/* GraphQL */`query orgIdList($ids:[GID!]!){
  nodes(ids: $ids){
    ... on Org{
      id,code,name
    }
  }
}`)


/**
 * 获取组织信息
 * @param params
 * @param filter
 * @param sort
 * @returns
 */
export async function getOrgList(gather: {
  current?: number;
  pageSize?: number;
  where?: OrgWhereInput;
  orderBy?: OrgOrder;
}) {
  const result = await pagingRequest(queryOrgList, {
    first: gather.pageSize || 20,
    where: gather.where,
    orderBy: gather.orderBy || {
      direction: OrderDirection.Asc,
      field: OrgOrderField.DisplaySort,
    },
  }, gather.current || 1);
  if (result.data?.organizations) {
    return result.data.organizations;
  }
  return null;
}




/**
 * 缓存org值
 * @param ids
 */
export async function updateCacheOrgListByIds(ids: (string | number)[]) {
  const cacheIds = Object.keys(cacheOrg)
  const newCacheIds = ids.filter(id => !cacheIds.includes(`${id}`))
  if (newCacheIds.length) {
    const result = await queryRequest(queryOrgIdList, {
      ids: newCacheIds.map(id => gid('org', id))
    })
    result.data?.nodes?.forEach(item => {
      if (item?.__typename === 'Org') {
        cacheOrg[item.id] = item as Org
      }
    })
  }
}
