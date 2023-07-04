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
    "query msgTypeList($first: Int,$orderBy:MsgTypeOrder,$where:MsgTypeWhereInput){\n  msgTypes(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,comments,appID,status,category,status,canSubs,canCustom,createdAt\n      }\n    }\n  }\n}": types.MsgTypeListDocument,
    "query msgTypeInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgType{\n      id,name,comments,appID,status,category,status,canSubs,canCustom,createdAt\n    }\n  }\n}": types.MsgTypeInfoDocument,
    "mutation createMsgType($input: CreateMsgTypeInput!){\n  createMsgType(input: $input){id}\n}": types.CreateMsgTypeDocument,
    "mutation updateMsgType($id:ID!,$input: UpdateMsgTypeInput!){\n  updateMsgType(id:$id,input: $input){id}\n}": types.UpdateMsgTypeDocument,
    "mutation delMsgType($id:ID!){\n  deleteMsgType(id:$id)\n}": types.DelMsgTypeDocument,
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
export function gql(source: "query msgTypeList($first: Int,$orderBy:MsgTypeOrder,$where:MsgTypeWhereInput){\n  msgTypes(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,comments,appID,status,category,status,canSubs,canCustom,createdAt\n      }\n    }\n  }\n}"): (typeof documents)["query msgTypeList($first: Int,$orderBy:MsgTypeOrder,$where:MsgTypeWhereInput){\n  msgTypes(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,comments,appID,status,category,status,canSubs,canCustom,createdAt\n      }\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query msgTypeInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgType{\n      id,name,comments,appID,status,category,status,canSubs,canCustom,createdAt\n    }\n  }\n}"): (typeof documents)["query msgTypeInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgType{\n      id,name,comments,appID,status,category,status,canSubs,canCustom,createdAt\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation createMsgType($input: CreateMsgTypeInput!){\n  createMsgType(input: $input){id}\n}"): (typeof documents)["mutation createMsgType($input: CreateMsgTypeInput!){\n  createMsgType(input: $input){id}\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation updateMsgType($id:ID!,$input: UpdateMsgTypeInput!){\n  updateMsgType(id:$id,input: $input){id}\n}"): (typeof documents)["mutation updateMsgType($id:ID!,$input: UpdateMsgTypeInput!){\n  updateMsgType(id:$id,input: $input){id}\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation delMsgType($id:ID!){\n  deleteMsgType(id:$id)\n}"): (typeof documents)["mutation delMsgType($id:ID!){\n  deleteMsgType(id:$id)\n}"];

export function gql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;