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
            id,createdBy,createdAt,updatedBy,updatedAt,principalName,displayName,
            email,mobile,userType,creationType,registerIP,status,comments
          }
        }
      }
    }
  }
}`);

const queryOrgUserListAndIsOrgRole = gql(/* GraphQL */`query orgUserListAndIsOrgRole($gid: GID!,$orgRoleId:ID!,$first: Int,$orderBy:UserOrder,$where:UserWhereInput){
  node(id:$gid){
    ... on Org{
      id,
      users(first:$first,orderBy: $orderBy,where: $where){
        totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }
        edges{
          cursor,node{
            id,createdBy,createdAt,updatedBy,updatedAt,principalName,displayName,
            email,mobile,userType,creationType,registerIP,status,comments
            isAssignOrgRole(orgRoleID: $orgRoleId)
            isAllowRevokeRole(orgRoleID: $orgRoleId)
          }
        }
      }
    }
  }
}`);

const queryOrgRoleUserList = gql(/* GraphQL */`query orgRoleUserList($roleId: ID!,$first: Int,$orderBy:UserOrder,$where:UserWhereInput){
  orgRoleUsers(roleID:$roleId,first:$first,orderBy: $orderBy,where: $where){
    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }
    edges{
      cursor,node{
        id,createdBy,createdAt,updatedBy,updatedAt,principalName,displayName,
        email,mobile,userType,creationType,registerIP,status,comments
      }
    }
  }
}`);

const queryOrgRoleUserListAndIsOrgRole = gql(/* GraphQL */`query orgRoleUserListAndIsOrgRole($roleId: ID!,$orgRoleId:ID!,$first: Int,$orderBy:UserOrder,$where:UserWhereInput){
  orgRoleUsers(roleID:$roleId,first:$first,orderBy: $orderBy,where: $where){
    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }
    edges{
      cursor,node{
        id,createdBy,createdAt,updatedBy,updatedAt,principalName,displayName,
        email,mobile,userType,creationType,registerIP,status,comments
        isAssignOrgRole(orgRoleID: $orgRoleId)
        isAllowRevokeRole(orgRoleID: $orgRoleId)
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
  isGrant?: {
    orgRoleId?: string;
  },
) {

  const result = isGrant?.orgRoleId ? await pagingRequest(
    queryOrgUserListAndIsOrgRole, {
    orgRoleId: isGrant.orgRoleId,
    gid: gid('org', orgId),
    first: gather.pageSize || 20,
    where: gather.where,
    orderBy: gather.orderBy,
  }, gather.current || 1) : await pagingRequest(
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

/**
 * 组织下的角色用户
 * @param roleId
 * @param params
 * @param filter
 * @param sort
 * @returns
 */
export async function getOrgRoleUserList(
  roleId: string,
  gather: {
    current?: number;
    pageSize?: number;
    where?: UserWhereInput;
    orderBy?: UserOrder;
  },
  isGrant?: {
    orgRoleId?: string;
  },
) {

  const result = isGrant?.orgRoleId ? await pagingRequest(
    queryOrgRoleUserListAndIsOrgRole, {
    roleId,
    orgRoleId: isGrant.orgRoleId,
    first: gather.pageSize || 20,
    where: gather.where,
    orderBy: gather.orderBy,
  }, gather.current || 1) : await pagingRequest(
    queryOrgRoleUserList, {
    roleId,
    first: gather.pageSize || 20,
    where: gather.where,
    orderBy: gather.orderBy,
  }, gather.current || 1);

  if (result.data?.orgRoleUsers) {
    return result.data?.orgRoleUsers;
  }
  return null;
}

