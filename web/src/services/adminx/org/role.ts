import { gid } from '@/util';
import { gql } from '@/__generated__/adminx';
import { pagingRequest, queryRequest } from '../';
import { OrgRoleOrder, OrgRoleWhereInput } from '@/__generated__/adminx/graphql';

const queryOrgGroupList = gql(/* GraphQL */`query orgGroupList($first: Int,$orderBy:OrgRoleOrder,$where:OrgRoleWhereInput){
  orgGroups(first:$first,orderBy: $orderBy,where: $where){
    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }
    edges{
      cursor,node{
        id,createdBy,createdAt,updatedBy,updatedAt,orgID,kind,name,comments,isAppRole
      }
    }
  }
}`);

const queryOrgGroupListAndIsGrant = gql(/* GraphQL */`query orgGroupListAndIsGrant($userId: ID!,$first: Int,$orderBy:OrgRoleOrder,$where:OrgRoleWhereInput){
  orgGroups(first:$first,orderBy: $orderBy,where: $where){
    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }
    edges{
      cursor,node{
        id,createdBy,createdAt,updatedBy,updatedAt,orgID,kind,name,comments,isAppRole
        isGrantUser(userID: $userId)
      }
    }
  }
}`);

const queryUserGroupList = gql(/* GraphQL */`query userGroupList($userId: ID!,$first: Int,$orderBy:OrgRoleOrder,$where:OrgRoleWhereInput){
  userGroups(userID:$userId,first:$first,orderBy: $orderBy,where: $where){
    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }
    edges{
      cursor,node{
        id,createdBy,createdAt,updatedBy,updatedAt,orgID,kind,name,comments,isAppRole
      }
    }
  }
}`);

const queryOrgRoleList = gql(/* GraphQL */`query orgRoleList($first: Int,$orderBy:OrgRoleOrder,$where:OrgRoleWhereInput){
  orgRoles(first:$first,orderBy: $orderBy,where: $where){
    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }
    edges{
      cursor,node{
        id,createdBy,createdAt,updatedBy,updatedAt,orgID,kind,name,comments,isAppRole
      }
    }
  }
}`);

const queryOrgRoleListAndIsGrant = gql(/* GraphQL */`query orgRoleListAndIsGrant($userId:ID!,$first: Int,$orderBy:OrgRoleOrder,$where:OrgRoleWhereInput){
  orgRoles(first:$first,orderBy: $orderBy,where: $where){
    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }
    edges{
      cursor,node{
        id,createdBy,createdAt,updatedBy,updatedAt,orgID,kind,name,comments,isAppRole
        isGrantUser(userID: $userId)
      }
    }
  }
}`);

const queryOrgRoleInfo = gql(/* GraphQL */`query orgRoleInfo($gid:GID!){
  node(id:$gid){
    ... on OrgRole{
      id,createdBy,createdAt,updatedBy,updatedAt,orgID,kind,name,comments,isAppRole
    }
  }
}`);


/**
 * 获取组织用户组
 * @param params
 * @param filter
 * @param sort
 * @returns
 */
export async function getOrgGroupList(
  gather: {
    current?: number;
    pageSize?: number;
    where?: OrgRoleWhereInput;
    orderBy?: OrgRoleOrder;
  },
  isGrant?: {
    userId?: string;
  }) {

  const result = isGrant?.userId ? await pagingRequest(
    queryOrgGroupListAndIsGrant, {
    userId: isGrant.userId,
    first: gather.pageSize || 20,
    where: gather.where,
    orderBy: gather.orderBy,
  }, gather.current || 1) : await pagingRequest(
    queryOrgGroupList, {
    first: gather.pageSize || 20,
    where: gather.where,
    orderBy: gather.orderBy,
  }, gather.current || 1);

  if (result.data?.orgGroups) {
    return result.data?.orgGroups;
  }
  return null;
}

/**
 * 获取用户加入的用户组
 * @param userId
 * @param params
 * @param filter
 * @param sort
 * @returns
 */
export async function getUserJoinGroupList(
  userId: string,
  gather: {
    current?: number;
    pageSize?: number;
    where?: OrgRoleWhereInput;
    orderBy?: OrgRoleOrder;
  }) {
  const
    result = await pagingRequest(
      queryUserGroupList, {
      userId: userId,
      first: gather.pageSize || 20,
      where: gather.where,
      orderBy: gather.orderBy,
    }, gather.current || 1);

  if (result.data?.userGroups) {
    return result.data.userGroups;
  }
  return null;
}

/**
 * 获取组织用户组
 * @param params
 * @param filter
 * @param sort
 * @returns
 */
export async function getOrgRoleList(
  gather: {
    current?: number;
    pageSize?: number;
    where?: OrgRoleWhereInput;
    orderBy?: OrgRoleOrder;
  },
  isGrant?: {
    userId?: string;
  }) {
  const
    result = isGrant?.userId ? await pagingRequest(
      queryOrgRoleListAndIsGrant, {
      userId: isGrant.userId,
      first: gather.pageSize || 20,
      where: gather.where,
      orderBy: gather.orderBy,
    }, gather.current || 1) : await pagingRequest(
      queryOrgRoleList, {
      first: gather.pageSize || 20,
      where: gather.where,
      orderBy: gather.orderBy,
    }, gather.current || 1);

  if (result.data?.orgRoles) {
    return result.data.orgRoles;
  }
  return null;
}


/**
 * 获取用户信息
 * @param orgRoleId
 * @returns
 */
export async function getOrgRoleInfo(orgRoleId: string) {
  const
    result = await queryRequest(
      queryOrgRoleInfo, {
      gid: gid('org_role', orgRoleId),
    });

  if (result.data?.node?.__typename === 'OrgRole') {
    return result.data.node;
  }
  return null;
}

