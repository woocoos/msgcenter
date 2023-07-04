import { gid } from '@/util';
import { pagingRequest, queryRequest } from '../';
import { gql } from '@/__generated__/adminx';
import { OrderDirection, Org, OrgKind, OrgOrder, OrgOrderField, OrgWhereInput } from '@/__generated__/adminx/graphql';

export const EnumOrgStatus = {
  active: { text: '活跃', status: 'success' },
  inactive: { text: '失活', status: 'default' },
  processing: { text: '处理中', status: 'warning' },
},
  EnumOrgKind = {
    root: { text: '组织' },
    org: { text: '部门' },
  };

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

const queryOrgInfo = gql(/* GraphQL */`query orgInfo($gid:GID!){
  node(id: $gid){
    ... on Org{
      id,createdBy,createdAt,updatedBy,updatedAt,deletedAt,ownerID,parentID,kind,
      domain,code,name,profile,status,path,displaySort,countryCode,timezone,
      owner { id,displayName }
    }
  }
}`);


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
 * 通过path获取整个组织树结构
 * @param orgId
 * @returns
 */
export async function getOrgPathList(orgId: string, kind: OrgKind) {
  const topOrg = await getOrgInfo(orgId),
    orgList: Org[] = [];
  if (topOrg?.id) {
    orgList.push(topOrg as Org);
    const result = await getOrgList({
      pageSize: 9999,
      where: {
        pathHasPrefix: `${topOrg.path}/`,
        kind: kind,
      },
    });
    if (result?.totalCount) {
      orgList.push(...(result.edges?.map(item => item?.node) as Org[] || []));
    }
  }
  return orgList;
}


/**
 * 获取组织信息
 * @param orgId
 * @returns
 */
export async function getOrgInfo(orgId: string) {
  const
    result = await queryRequest(queryOrgInfo, {
      gid: gid('org', orgId),
    });
  if (result.data?.node?.__typename === 'Org') {
    return result.data.node;
  }
  return null;
}
