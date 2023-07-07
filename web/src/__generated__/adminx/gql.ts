/* eslint-disable */
import * as types from './graphql';
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';

/**
 * Map of all GraphQL operations in the project.
 *
 * This map has several performance disadvantages:
 * 1. It is not tree-shakeable, so it will include all operations in the project.
 * 2. It is not minifiable, so the string of a GraphQL query will be multiple times inside the bundle.
 * 3. It does not support dead code elimination, so it will add unused operations.
 *
 * Therefore it is highly recommended to use the babel or swc plugin for production.
 */
const documents = {
    "query appList($first: Int,$orderBy:AppOrder,$where:AppWhereInput){\n  apps(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,code,kind,redirectURI,appKey,appSecret,scopes,tokenValidity,\n        refreshTokenValidity,logo,comments,status,createdAt\n      }\n    }\n  }\n}": types.AppListDocument,
    "query appIdList($ids:[GID!]!){\n  nodes(ids: $ids){\n    ... on App{\n      id,code,name\n    }\n  }\n}": types.AppIdListDocument,
    "query globalID($type:String!,$id:ID!){\n  globalID(type:$type,id:$id)\n}": types.GlobalIdDocument,
    "query orgList($first: Int,$orderBy:OrgOrder,$where:OrgWhereInput){\n  organizations(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,createdBy,createdAt,updatedBy,updatedAt,deletedAt,ownerID,parentID,kind,\n        domain,code,name,profile,status,path,displaySort,countryCode,timezone,\n        owner { id,displayName }\n      }\n    }\n  }\n}": types.OrgListDocument,
    "query orgIdList($ids:[GID!]!){\n  nodes(ids: $ids){\n    ... on Org{\n      id,code,name\n    }\n  }\n}": types.OrgIdListDocument,
    "query orgGroupList($first: Int,$orderBy:OrgRoleOrder,$where:OrgRoleWhereInput){\n  orgGroups(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,orgID,kind,name,comments\n      }\n    }\n  }\n}": types.OrgGroupListDocument,
    "query orgRoleIdList($ids:[GID!]!){\n  nodes(ids: $ids){\n    ... on OrgRole{\n      id,orgID,kind,name\n    }\n  }\n}": types.OrgRoleIdListDocument,
    "query orgUserList($gid: GID!,$first: Int,$orderBy:UserOrder,$where:UserWhereInput){\n  node(id:$gid){\n    ... on Org{\n      id,\n      users(first:$first,orderBy: $orderBy,where: $where){\n        totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n        edges{\n          cursor,node{\n            id,displayName,email\n          }\n        }\n      }\n    }\n  }\n}": types.OrgUserListDocument,
    "query userIdList($ids:[GID!]!){\n  nodes(ids: $ids){\n    ... on User{\n      id,displayName\n    }\n  }\n}": types.UserIdListDocument,
    "query  checkPermission($permission:String!){\n  checkPermission(permission: $permission)\n}": types.CheckPermissionDocument,
    "query userPermissionList($where: AppActionWhereInput){\n  userPermissions(where: $where){\n    id,appID,name,kind,method\n  }\n}": types.UserPermissionListDocument,
    "query userMenuList($appCode:String!){\n  userMenus(appCode: $appCode){\n    id,parentID,kind,name,comments,displaySort,icon,route\n  }\n}": types.UserMenuListDocument,
    "query userRootOrgs{\n  userRootOrgs{\n    id,parentID,code,name,displaySort,countryCode,timezone\n  }\n}": types.UserRootOrgsDocument,
};

/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 *
 *
 * @example
 * ```ts
 * const query = gql(`query GetUser($id: ID!) { user(id: $id) { name } }`);
 * ```
 *
 * The query argument is unknown!
 * Please regenerate the types.
 */
export function gql(source: string): unknown;

/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query appList($first: Int,$orderBy:AppOrder,$where:AppWhereInput){\n  apps(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,code,kind,redirectURI,appKey,appSecret,scopes,tokenValidity,\n        refreshTokenValidity,logo,comments,status,createdAt\n      }\n    }\n  }\n}"): (typeof documents)["query appList($first: Int,$orderBy:AppOrder,$where:AppWhereInput){\n  apps(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,code,kind,redirectURI,appKey,appSecret,scopes,tokenValidity,\n        refreshTokenValidity,logo,comments,status,createdAt\n      }\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query appIdList($ids:[GID!]!){\n  nodes(ids: $ids){\n    ... on App{\n      id,code,name\n    }\n  }\n}"): (typeof documents)["query appIdList($ids:[GID!]!){\n  nodes(ids: $ids){\n    ... on App{\n      id,code,name\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query globalID($type:String!,$id:ID!){\n  globalID(type:$type,id:$id)\n}"): (typeof documents)["query globalID($type:String!,$id:ID!){\n  globalID(type:$type,id:$id)\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query orgList($first: Int,$orderBy:OrgOrder,$where:OrgWhereInput){\n  organizations(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,createdBy,createdAt,updatedBy,updatedAt,deletedAt,ownerID,parentID,kind,\n        domain,code,name,profile,status,path,displaySort,countryCode,timezone,\n        owner { id,displayName }\n      }\n    }\n  }\n}"): (typeof documents)["query orgList($first: Int,$orderBy:OrgOrder,$where:OrgWhereInput){\n  organizations(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,createdBy,createdAt,updatedBy,updatedAt,deletedAt,ownerID,parentID,kind,\n        domain,code,name,profile,status,path,displaySort,countryCode,timezone,\n        owner { id,displayName }\n      }\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query orgIdList($ids:[GID!]!){\n  nodes(ids: $ids){\n    ... on Org{\n      id,code,name\n    }\n  }\n}"): (typeof documents)["query orgIdList($ids:[GID!]!){\n  nodes(ids: $ids){\n    ... on Org{\n      id,code,name\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query orgGroupList($first: Int,$orderBy:OrgRoleOrder,$where:OrgRoleWhereInput){\n  orgGroups(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,orgID,kind,name,comments\n      }\n    }\n  }\n}"): (typeof documents)["query orgGroupList($first: Int,$orderBy:OrgRoleOrder,$where:OrgRoleWhereInput){\n  orgGroups(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,orgID,kind,name,comments\n      }\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query orgRoleIdList($ids:[GID!]!){\n  nodes(ids: $ids){\n    ... on OrgRole{\n      id,orgID,kind,name\n    }\n  }\n}"): (typeof documents)["query orgRoleIdList($ids:[GID!]!){\n  nodes(ids: $ids){\n    ... on OrgRole{\n      id,orgID,kind,name\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query orgUserList($gid: GID!,$first: Int,$orderBy:UserOrder,$where:UserWhereInput){\n  node(id:$gid){\n    ... on Org{\n      id,\n      users(first:$first,orderBy: $orderBy,where: $where){\n        totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n        edges{\n          cursor,node{\n            id,displayName,email\n          }\n        }\n      }\n    }\n  }\n}"): (typeof documents)["query orgUserList($gid: GID!,$first: Int,$orderBy:UserOrder,$where:UserWhereInput){\n  node(id:$gid){\n    ... on Org{\n      id,\n      users(first:$first,orderBy: $orderBy,where: $where){\n        totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n        edges{\n          cursor,node{\n            id,displayName,email\n          }\n        }\n      }\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query userIdList($ids:[GID!]!){\n  nodes(ids: $ids){\n    ... on User{\n      id,displayName\n    }\n  }\n}"): (typeof documents)["query userIdList($ids:[GID!]!){\n  nodes(ids: $ids){\n    ... on User{\n      id,displayName\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query  checkPermission($permission:String!){\n  checkPermission(permission: $permission)\n}"): (typeof documents)["query  checkPermission($permission:String!){\n  checkPermission(permission: $permission)\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query userPermissionList($where: AppActionWhereInput){\n  userPermissions(where: $where){\n    id,appID,name,kind,method\n  }\n}"): (typeof documents)["query userPermissionList($where: AppActionWhereInput){\n  userPermissions(where: $where){\n    id,appID,name,kind,method\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query userMenuList($appCode:String!){\n  userMenus(appCode: $appCode){\n    id,parentID,kind,name,comments,displaySort,icon,route\n  }\n}"): (typeof documents)["query userMenuList($appCode:String!){\n  userMenus(appCode: $appCode){\n    id,parentID,kind,name,comments,displaySort,icon,route\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query userRootOrgs{\n  userRootOrgs{\n    id,parentID,code,name,displaySort,countryCode,timezone\n  }\n}"): (typeof documents)["query userRootOrgs{\n  userRootOrgs{\n    id,parentID,code,name,displaySort,countryCode,timezone\n  }\n}"];

export function gql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;