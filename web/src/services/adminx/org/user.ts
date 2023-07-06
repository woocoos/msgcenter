import { gql } from '@/__generated__/adminx';
import { gid } from '@/util';
import { pagingRequest } from '../';
import { UserOrder, UserWhereInput } from '@/__generated__/adminx/graphql';


const queryOrgUserList = gql(/* GraphQL */`query orgUserList($gid: GID!,$first: Int,$orderBy:UserOrder,$where:UserWhereInput){
  node(id:$gid){
    ... on Org{
      id,
      users(first:$first,orderBy: $orderBy,where: $where){
        totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }
        edges{
          cursor,node{
            id,displayName,
          }
        }
      }
    }
  }
}`);


/**
 * 组织下的用户信息
 * @param orgId
 * @returns
 */
export async function getOrgUserList(
  orgId: string,
  gather: {
    current?: number;
    pageSize?: number;
    where?: UserWhereInput;
    orderBy?: UserOrder;
  },
) {
  const result = await pagingRequest(
    queryOrgUserList, {
    gid: gid('org', orgId),
    first: gather.pageSize || 20,
    where: gather.where,
    orderBy: gather.orderBy,
  }, gather.current || 1);

  if (result.data?.node?.__typename === 'Org') {
    return result.data.node.users;
  }
  return null;
}
