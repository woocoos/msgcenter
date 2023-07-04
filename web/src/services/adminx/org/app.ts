import { gql } from '@/__generated__/adminx';
import { gid } from '@/util';
import { pagingRequest } from '../';
import { AppOrder, AppWhereInput } from '@/__generated__/adminx/graphql';

export const EnumAppKind = {
  web: { text: 'web' },
  native: { text: 'native' },
  server: { text: 'server' },
};

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

