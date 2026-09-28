import { gql } from '@/generated/adminx';
import { UserOrder, UserWhereInput } from '@/generated/adminx/graphql';
import { gid, instanceName } from '@knockout-js/api';
import { paging } from '@knockout-js/ice-urql/request';

const orgUserListQuery = gql(/* GraphQL */`query orgUserList($gid: GID!,$first: Int,$orderBy:UserOrder,$where:UserWhereInput){
  node(id:$gid){
    ... on Org{
      id,
      users(first:$first,orderBy: $orderBy,where: $where){
        totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }
        edges{
          cursor,node{
            id,displayName,addresses{email}
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
  const result = await paging(
    orgUserListQuery, {
    gid: gid('Org', orgId),
    first: gather.pageSize || 20,
    where: gather.where,
    orderBy: gather.orderBy,
  }, gather.current || 1, {
    instanceName: instanceName.UCENTER,
  });

  if (result.data?.node?.__typename === 'Org') {
    return result.data.node.users;
  }
  return null;
}
