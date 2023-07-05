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
    "query msgChannelList($first: Int,$orderBy:MsgChannelOrder,$where:MsgChannelWhereInput){\n  msgChannels(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,receiverType,tenantID,comments,status,status,createdAt\n      }\n    }\n  }\n}": types.MsgChannelListDocument,
    "query msgChannelInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgChannel{\n      id,name,receiverType,tenantID,comments,status,status,createdAt\n    }\n  }\n}": types.MsgChannelInfoDocument,
    "query msgChannelReceiverInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgChannel{\n      id,receiver{\n        name,\n        emailConfigs{\n          authIdentity,authPassword,authSecret,authType,authUsername,from,headers,smartHost,to\n        }\n      }\n    }\n  }\n}": types.MsgChannelReceiverInfoDocument,
    "mutation createMsgChannel($input: CreateMsgChannelInput!){\n createMsgChannel(input:$input){id}\n}": types.CreateMsgChannelDocument,
    "mutation updateMsgChannel($id:ID!,$input: UpdateMsgChannelInput!){\n updateMsgChannel(id:$id,input:$input){id}\n}": types.UpdateMsgChannelDocument,
    "mutation delMsgChannel($id:ID!){\n deleteMsgChannel(id:$id)\n}": types.DelMsgChannelDocument,
    "mutation enableMsgChannel($id:ID!){\n enableMsgChannel(id:$id){id}\n}": types.EnableMsgChannelDocument,
    "mutation disableMsgChannel($id:ID!){\n disableMsgChannel(id:$id){id}\n}": types.DisableMsgChannelDocument,
    "query msgEventList($first: Int,$orderBy:MsgEventOrder,$where:MsgEventWhereInput){\n  msgEvents(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,comments,status,createdAt,msgTypeID,modes\n        msgType{\n          id,category,appID,name\n        }\n      }\n    }\n  }\n}": types.MsgEventListDocument,
    "query MsgEventInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgEvent{\n      id,name,comments,status,createdAt,msgTypeID,modes\n      msgType{\n        id,category,appID,name\n      }\n    }\n  }\n}": types.MsgEventInfoDocument,
    "mutation createMsgEvent($input: CreateMsgEventInput!){\n  createMsgEvent(input: $input){id}\n}": types.CreateMsgEventDocument,
    "mutation updateMsgEvent($id:ID!,$input: UpdateMsgEventInput!){\n  updateMsgEvent(id:$id,input: $input){id}\n}": types.UpdateMsgEventDocument,
    "mutation delMsgEvent($id:ID!){\n  deleteMsgEvent(id:$id)\n}": types.DelMsgEventDocument,
    "mutation enableMsgEvent($id:ID!){\n  enableMsgEvent(id:$id){id}\n}": types.EnableMsgEventDocument,
    "mutation disableMsgEvent($id:ID!){\n  disableMsgEvent(id:$id){id}\n}": types.DisableMsgEventDocument,
    "query msgTemplateList($first: Int,$orderBy:MsgTemplateOrder,$where:MsgTemplateWhereInput){\n  msgTemplates(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,\n        receiverType,format,subject,from,to,cc,bcc,body,tpl,attachments\n      }\n    }\n  }\n}": types.MsgTemplateListDocument,
    "query MsgTemplateInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgTemplate{\n      id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,\n      receiverType,format,subject,from,to,cc,bcc,body,tpl,attachments\n    }\n  }\n}": types.MsgTemplateInfoDocument,
    "mutation createMsgTemplate($input: CreateMsgTemplateInput!){\n  createMsgTemplate(input: $input){id}\n}": types.CreateMsgTemplateDocument,
    "mutation updateMsgTemplate($id:ID!,$input: UpdateMsgTemplateInput!){\n  updateMsgTemplate(id:$id,input: $input){id}\n}": types.UpdateMsgTemplateDocument,
    "mutation delMsgTemplate($id:ID!){\n  deleteMsgTemplate(id:$id)\n}": types.DelMsgTemplateDocument,
    "mutation enableMsgTemplate($id:ID!){\n  enableMsgTemplate(id:$id){id}\n}": types.EnableMsgTemplateDocument,
    "mutation disableMsgTemplate($id:ID!){\n  disableMsgTemplate(id:$id){id}\n}": types.DisableMsgTemplateDocument,
    "query msgTypeList($first: Int,$orderBy:MsgTypeOrder,$where:MsgTypeWhereInput){\n  msgTypes(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,comments,appID,status,category,canSubs,canCustom,createdAt\n      }\n    }\n  }\n}": types.MsgTypeListDocument,
    "query msgTypeInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgType{\n      id,name,comments,appID,status,category,canSubs,canCustom,createdAt\n    }\n  }\n}": types.MsgTypeInfoDocument,
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
export function gql(source: "query msgChannelList($first: Int,$orderBy:MsgChannelOrder,$where:MsgChannelWhereInput){\n  msgChannels(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,receiverType,tenantID,comments,status,status,createdAt\n      }\n    }\n  }\n}"): (typeof documents)["query msgChannelList($first: Int,$orderBy:MsgChannelOrder,$where:MsgChannelWhereInput){\n  msgChannels(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,receiverType,tenantID,comments,status,status,createdAt\n      }\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query msgChannelInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgChannel{\n      id,name,receiverType,tenantID,comments,status,status,createdAt\n    }\n  }\n}"): (typeof documents)["query msgChannelInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgChannel{\n      id,name,receiverType,tenantID,comments,status,status,createdAt\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query msgChannelReceiverInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgChannel{\n      id,receiver{\n        name,\n        emailConfigs{\n          authIdentity,authPassword,authSecret,authType,authUsername,from,headers,smartHost,to\n        }\n      }\n    }\n  }\n}"): (typeof documents)["query msgChannelReceiverInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgChannel{\n      id,receiver{\n        name,\n        emailConfigs{\n          authIdentity,authPassword,authSecret,authType,authUsername,from,headers,smartHost,to\n        }\n      }\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation createMsgChannel($input: CreateMsgChannelInput!){\n createMsgChannel(input:$input){id}\n}"): (typeof documents)["mutation createMsgChannel($input: CreateMsgChannelInput!){\n createMsgChannel(input:$input){id}\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation updateMsgChannel($id:ID!,$input: UpdateMsgChannelInput!){\n updateMsgChannel(id:$id,input:$input){id}\n}"): (typeof documents)["mutation updateMsgChannel($id:ID!,$input: UpdateMsgChannelInput!){\n updateMsgChannel(id:$id,input:$input){id}\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation delMsgChannel($id:ID!){\n deleteMsgChannel(id:$id)\n}"): (typeof documents)["mutation delMsgChannel($id:ID!){\n deleteMsgChannel(id:$id)\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation enableMsgChannel($id:ID!){\n enableMsgChannel(id:$id){id}\n}"): (typeof documents)["mutation enableMsgChannel($id:ID!){\n enableMsgChannel(id:$id){id}\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation disableMsgChannel($id:ID!){\n disableMsgChannel(id:$id){id}\n}"): (typeof documents)["mutation disableMsgChannel($id:ID!){\n disableMsgChannel(id:$id){id}\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query msgEventList($first: Int,$orderBy:MsgEventOrder,$where:MsgEventWhereInput){\n  msgEvents(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,comments,status,createdAt,msgTypeID,modes\n        msgType{\n          id,category,appID,name\n        }\n      }\n    }\n  }\n}"): (typeof documents)["query msgEventList($first: Int,$orderBy:MsgEventOrder,$where:MsgEventWhereInput){\n  msgEvents(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,comments,status,createdAt,msgTypeID,modes\n        msgType{\n          id,category,appID,name\n        }\n      }\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query MsgEventInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgEvent{\n      id,name,comments,status,createdAt,msgTypeID,modes\n      msgType{\n        id,category,appID,name\n      }\n    }\n  }\n}"): (typeof documents)["query MsgEventInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgEvent{\n      id,name,comments,status,createdAt,msgTypeID,modes\n      msgType{\n        id,category,appID,name\n      }\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation createMsgEvent($input: CreateMsgEventInput!){\n  createMsgEvent(input: $input){id}\n}"): (typeof documents)["mutation createMsgEvent($input: CreateMsgEventInput!){\n  createMsgEvent(input: $input){id}\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation updateMsgEvent($id:ID!,$input: UpdateMsgEventInput!){\n  updateMsgEvent(id:$id,input: $input){id}\n}"): (typeof documents)["mutation updateMsgEvent($id:ID!,$input: UpdateMsgEventInput!){\n  updateMsgEvent(id:$id,input: $input){id}\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation delMsgEvent($id:ID!){\n  deleteMsgEvent(id:$id)\n}"): (typeof documents)["mutation delMsgEvent($id:ID!){\n  deleteMsgEvent(id:$id)\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation enableMsgEvent($id:ID!){\n  enableMsgEvent(id:$id){id}\n}"): (typeof documents)["mutation enableMsgEvent($id:ID!){\n  enableMsgEvent(id:$id){id}\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation disableMsgEvent($id:ID!){\n  disableMsgEvent(id:$id){id}\n}"): (typeof documents)["mutation disableMsgEvent($id:ID!){\n  disableMsgEvent(id:$id){id}\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query msgTemplateList($first: Int,$orderBy:MsgTemplateOrder,$where:MsgTemplateWhereInput){\n  msgTemplates(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,\n        receiverType,format,subject,from,to,cc,bcc,body,tpl,attachments\n      }\n    }\n  }\n}"): (typeof documents)["query msgTemplateList($first: Int,$orderBy:MsgTemplateOrder,$where:MsgTemplateWhereInput){\n  msgTemplates(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,\n        receiverType,format,subject,from,to,cc,bcc,body,tpl,attachments\n      }\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query MsgTemplateInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgTemplate{\n      id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,\n      receiverType,format,subject,from,to,cc,bcc,body,tpl,attachments\n    }\n  }\n}"): (typeof documents)["query MsgTemplateInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgTemplate{\n      id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,\n      receiverType,format,subject,from,to,cc,bcc,body,tpl,attachments\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation createMsgTemplate($input: CreateMsgTemplateInput!){\n  createMsgTemplate(input: $input){id}\n}"): (typeof documents)["mutation createMsgTemplate($input: CreateMsgTemplateInput!){\n  createMsgTemplate(input: $input){id}\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation updateMsgTemplate($id:ID!,$input: UpdateMsgTemplateInput!){\n  updateMsgTemplate(id:$id,input: $input){id}\n}"): (typeof documents)["mutation updateMsgTemplate($id:ID!,$input: UpdateMsgTemplateInput!){\n  updateMsgTemplate(id:$id,input: $input){id}\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation delMsgTemplate($id:ID!){\n  deleteMsgTemplate(id:$id)\n}"): (typeof documents)["mutation delMsgTemplate($id:ID!){\n  deleteMsgTemplate(id:$id)\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation enableMsgTemplate($id:ID!){\n  enableMsgTemplate(id:$id){id}\n}"): (typeof documents)["mutation enableMsgTemplate($id:ID!){\n  enableMsgTemplate(id:$id){id}\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation disableMsgTemplate($id:ID!){\n  disableMsgTemplate(id:$id){id}\n}"): (typeof documents)["mutation disableMsgTemplate($id:ID!){\n  disableMsgTemplate(id:$id){id}\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query msgTypeList($first: Int,$orderBy:MsgTypeOrder,$where:MsgTypeWhereInput){\n  msgTypes(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,comments,appID,status,category,canSubs,canCustom,createdAt\n      }\n    }\n  }\n}"): (typeof documents)["query msgTypeList($first: Int,$orderBy:MsgTypeOrder,$where:MsgTypeWhereInput){\n  msgTypes(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,comments,appID,status,category,canSubs,canCustom,createdAt\n      }\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query msgTypeInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgType{\n      id,name,comments,appID,status,category,canSubs,canCustom,createdAt\n    }\n  }\n}"): (typeof documents)["query msgTypeInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgType{\n      id,name,comments,appID,status,category,canSubs,canCustom,createdAt\n    }\n  }\n}"];
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