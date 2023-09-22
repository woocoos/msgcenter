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
    "query msgChannelReceiverInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgChannel{\n      id,receiver{\n        name,\n        emailConfigs{\n          authIdentity,authPassword,authSecret,authType,authUsername,from,headers,smartHost,to\n        },\n        messageConfig{\n          redirect,subject,to\n        }\n      }\n    }\n  }\n}": types.MsgChannelReceiverInfoDocument,
    "mutation createMsgChannel($input: CreateMsgChannelInput!){\n createMsgChannel(input:$input){id}\n}": types.CreateMsgChannelDocument,
    "mutation updateMsgChannel($id:ID!,$input: UpdateMsgChannelInput!){\n updateMsgChannel(id:$id,input:$input){id}\n}": types.UpdateMsgChannelDocument,
    "mutation delMsgChannel($id:ID!){\n deleteMsgChannel(id:$id)\n}": types.DelMsgChannelDocument,
    "mutation enableMsgChannel($id:ID!){\n enableMsgChannel(id:$id){id}\n}": types.EnableMsgChannelDocument,
    "mutation disableMsgChannel($id:ID!){\n disableMsgChannel(id:$id){id}\n}": types.DisableMsgChannelDocument,
    "query msgEventList($first: Int,$orderBy:MsgEventOrder,$where:MsgEventWhereInput){\n  msgEvents(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,comments,status,createdAt,msgTypeID,modes\n        msgType{\n          id,category,appID,name\n        }\n      }\n    }\n  }\n}": types.MsgEventListDocument,
    "query MsgEventInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgEvent{\n      id,name,comments,status,createdAt,msgTypeID,modes\n      msgType{\n        id,category,appID,name\n      }\n    }\n  }\n}": types.MsgEventInfoDocument,
    "query MsgEventInfoRoute($gid:GID!,$type:RouteStrType!){\n  node(id: $gid){\n    id\n    ... on MsgEvent{\n      id,name,comments,status,createdAt,msgTypeID,modes,routeStr(type:$type)\n    }\n  }\n}\n": types.MsgEventInfoRouteDocument,
    "mutation createMsgEvent($input: CreateMsgEventInput!){\n  createMsgEvent(input: $input){id}\n}": types.CreateMsgEventDocument,
    "mutation updateMsgEvent($id:ID!,$input: UpdateMsgEventInput!){\n  updateMsgEvent(id:$id,input: $input){id}\n}": types.UpdateMsgEventDocument,
    "mutation delMsgEvent($id:ID!){\n  deleteMsgEvent(id:$id)\n}": types.DelMsgEventDocument,
    "mutation enableMsgEvent($id:ID!){\n  enableMsgEvent(id:$id){id}\n}": types.EnableMsgEventDocument,
    "mutation disableMsgEvent($id:ID!){\n  disableMsgEvent(id:$id){id}\n}": types.DisableMsgEventDocument,
    "query msgInternalList($first: Int,$orderBy: MsgInternalOrder,$where:MsgInternalWhereInput){\n  msgInternals(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,tenantID,createdBy,createdAt,subject,body,format,redirect\n      }\n    }\n  }\n}": types.MsgInternalListDocument,
    "query userMsgInternalList($first: Int,$orderBy: MsgInternalToOrder,$where:MsgInternalToWhereInput){\n  userMsgInternalTos(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,msgInternalID,createdAt,deleteAt,readAt,userID\n        msgInternal{\n          id,tenantID,createdBy,createdAt,subject,body,format,redirect,category\n        }\n      }\n    }\n  }\n}": types.UserMsgInternalListDocument,
    "query msgInternalInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgInternal{\n      id,tenantID,createdBy,createdAt,subject,body,format,redirect\n    }\n  }\n}": types.MsgInternalInfoDocument,
    "query userMsgCategory{\n  userSubMsgCategory\n}": types.UserMsgCategoryDocument,
    "query userMsgCategoryNum($categories:[String!]!){\n  userUnreadMsgInternalsFromMsgCategory(categories:$categories)\n}": types.UserMsgCategoryNumDocument,
    "query msgInternalToInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgInternalTo{\n      id,msgInternalID,createdAt,deleteAt,readAt,userID\n      msgInternal{\n        id,tenantID,createdBy,createdAt,subject,body,format,redirect,category\n      }\n    }\n  }\n}": types.MsgInternalToInfoDocument,
    "mutation markMsgRead($ids:[ID!]!,$read:Boolean!){\n  markMsgInternalToReadOrUnRead(ids:$ids,read:$read)\n}": types.MarkMsgReadDocument,
    "mutation delMarkMsg($ids:[ID!]!){\n  markMsgInternalToDeleted(ids:$ids)\n}": types.DelMarkMsgDocument,
    "subscription subMsg{\n  message{\n    content,extras,format,sendAt,title,url\n  }\n}": types.SubMsgDocument,
    "query msgAlertList($first: Int,$orderBy:MsgAlertOrder,$where:MsgAlertWhereInput){\n  msgAlerts(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,startsAt,endsAt,labels,annotations,state,timeout\n      }\n    }\n  }\n}": types.MsgAlertListDocument,
    "query msgAlertLogList($gid:GID!,$first: Int,$orderBy:NlogOrder,$where:NlogWhereInput){\n   node(id: $gid){\n    id\n    ... on MsgAlert{\n      id,\n      nlog(first:$first,orderBy: $orderBy,where: $where){\n        totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n        edges{\n          cursor,node{\n            id,sendAt,expiresAt,groupKey,receiver,receiverType\n          }\n        }\n      }\n    }\n  }\n}": types.MsgAlertLogListDocument,
    "query silenceList($first: Int,$orderBy:SilenceOrder,$where:SilenceWhereInput){\n  silences(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,tenantID,comments,createdAt,startsAt,endsAt,state,\n        matchers{\n          type,name,value\n        }\n\n      }\n    }\n  }\n}": types.SilenceListDocument,
    "query SilenceInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on Silence{\n      id,tenantID,startsAt,endsAt,comments,state\n      matchers{\n        type,name,value\n      }\n    }\n  }\n}": types.SilenceInfoDocument,
    "mutation createSilence($input: CreateSilenceInput!){\n  createSilence(input: $input){id}\n}": types.CreateSilenceDocument,
    "mutation updateSilence($id:ID!,$input: UpdateSilenceInput!){\n  updateSilence(id:$id,input: $input){id}\n}": types.UpdateSilenceDocument,
    "mutation delSilence($id:ID!){\n  deleteSilence(id:$id)\n}": types.DelSilenceDocument,
    "query msgTemplateList($first: Int,$orderBy:MsgTemplateOrder,$where:MsgTemplateWhereInput){\n  msgTemplates(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,\n        receiverType,format,subject,from,to,cc,bcc,body,tpl,attachments\n      }\n    }\n  }\n}": types.MsgTemplateListDocument,
    "query MsgTemplateInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgTemplate{\n      id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,\n      receiverType,format,subject,from,to,cc,bcc,body,tplFileID,attachmentsFileIds,tpl,attachments\n    }\n  }\n}": types.MsgTemplateInfoDocument,
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
    "query msgTypeCategory($keyword:String,$appID:ID){\n  msgTypeCategories(keyword: $keyword,appID:$appID)\n}": types.MsgTypeCategoryDocument,
    "query msgTypeListAndSub($first: Int,$orderBy:MsgTypeOrder,$where:MsgTypeWhereInput){\n  msgTypes(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,category,\n        subscriberUsers{\n          id,tenantID,msgTypeID,userID\n        },\n        subscriberRoles{\n          id,tenantID,msgTypeID,orgRoleID\n        },\n        excludeSubscriberUsers{\n          id,tenantID,msgTypeID,userID\n        }\n      }\n    }\n  }\n}": types.MsgTypeListAndSubDocument,
    "query msgTypeAndSubInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgType{\n      id,name,appID,category,\n      subscriberUsers{\n        id,tenantID,msgTypeID,userID\n      },\n      subscriberRoles{\n        id,tenantID,msgTypeID,orgRoleID\n      },\n      excludeSubscriberUsers{\n        id,tenantID,msgTypeID,userID\n      }\n    }\n  }\n}": types.MsgTypeAndSubInfoDocument,
    "mutation createMsgSubscriber($inputs: [CreateMsgSubscriberInput!]!){\n  createMsgSubscriber(inputs: $inputs){id}\n}": types.CreateMsgSubscriberDocument,
    "mutation deleteMsgSubscriber($ids: [ID!]!){\n  deleteMsgSubscriber(ids: $ids)\n}": types.DeleteMsgSubscriberDocument,
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
export function gql(source: "query msgChannelReceiverInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgChannel{\n      id,receiver{\n        name,\n        emailConfigs{\n          authIdentity,authPassword,authSecret,authType,authUsername,from,headers,smartHost,to\n        },\n        messageConfig{\n          redirect,subject,to\n        }\n      }\n    }\n  }\n}"): (typeof documents)["query msgChannelReceiverInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgChannel{\n      id,receiver{\n        name,\n        emailConfigs{\n          authIdentity,authPassword,authSecret,authType,authUsername,from,headers,smartHost,to\n        },\n        messageConfig{\n          redirect,subject,to\n        }\n      }\n    }\n  }\n}"];
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
export function gql(source: "query MsgEventInfoRoute($gid:GID!,$type:RouteStrType!){\n  node(id: $gid){\n    id\n    ... on MsgEvent{\n      id,name,comments,status,createdAt,msgTypeID,modes,routeStr(type:$type)\n    }\n  }\n}\n"): (typeof documents)["query MsgEventInfoRoute($gid:GID!,$type:RouteStrType!){\n  node(id: $gid){\n    id\n    ... on MsgEvent{\n      id,name,comments,status,createdAt,msgTypeID,modes,routeStr(type:$type)\n    }\n  }\n}\n"];
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
export function gql(source: "query msgInternalList($first: Int,$orderBy: MsgInternalOrder,$where:MsgInternalWhereInput){\n  msgInternals(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,tenantID,createdBy,createdAt,subject,body,format,redirect\n      }\n    }\n  }\n}"): (typeof documents)["query msgInternalList($first: Int,$orderBy: MsgInternalOrder,$where:MsgInternalWhereInput){\n  msgInternals(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,tenantID,createdBy,createdAt,subject,body,format,redirect\n      }\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query userMsgInternalList($first: Int,$orderBy: MsgInternalToOrder,$where:MsgInternalToWhereInput){\n  userMsgInternalTos(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,msgInternalID,createdAt,deleteAt,readAt,userID\n        msgInternal{\n          id,tenantID,createdBy,createdAt,subject,body,format,redirect,category\n        }\n      }\n    }\n  }\n}"): (typeof documents)["query userMsgInternalList($first: Int,$orderBy: MsgInternalToOrder,$where:MsgInternalToWhereInput){\n  userMsgInternalTos(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,msgInternalID,createdAt,deleteAt,readAt,userID\n        msgInternal{\n          id,tenantID,createdBy,createdAt,subject,body,format,redirect,category\n        }\n      }\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query msgInternalInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgInternal{\n      id,tenantID,createdBy,createdAt,subject,body,format,redirect\n    }\n  }\n}"): (typeof documents)["query msgInternalInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgInternal{\n      id,tenantID,createdBy,createdAt,subject,body,format,redirect\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query userMsgCategory{\n  userSubMsgCategory\n}"): (typeof documents)["query userMsgCategory{\n  userSubMsgCategory\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query userMsgCategoryNum($categories:[String!]!){\n  userUnreadMsgInternalsFromMsgCategory(categories:$categories)\n}"): (typeof documents)["query userMsgCategoryNum($categories:[String!]!){\n  userUnreadMsgInternalsFromMsgCategory(categories:$categories)\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query msgInternalToInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgInternalTo{\n      id,msgInternalID,createdAt,deleteAt,readAt,userID\n      msgInternal{\n        id,tenantID,createdBy,createdAt,subject,body,format,redirect,category\n      }\n    }\n  }\n}"): (typeof documents)["query msgInternalToInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgInternalTo{\n      id,msgInternalID,createdAt,deleteAt,readAt,userID\n      msgInternal{\n        id,tenantID,createdBy,createdAt,subject,body,format,redirect,category\n      }\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation markMsgRead($ids:[ID!]!,$read:Boolean!){\n  markMsgInternalToReadOrUnRead(ids:$ids,read:$read)\n}"): (typeof documents)["mutation markMsgRead($ids:[ID!]!,$read:Boolean!){\n  markMsgInternalToReadOrUnRead(ids:$ids,read:$read)\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation delMarkMsg($ids:[ID!]!){\n  markMsgInternalToDeleted(ids:$ids)\n}"): (typeof documents)["mutation delMarkMsg($ids:[ID!]!){\n  markMsgInternalToDeleted(ids:$ids)\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "subscription subMsg{\n  message{\n    content,extras,format,sendAt,title,url\n  }\n}"): (typeof documents)["subscription subMsg{\n  message{\n    content,extras,format,sendAt,title,url\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query msgAlertList($first: Int,$orderBy:MsgAlertOrder,$where:MsgAlertWhereInput){\n  msgAlerts(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,startsAt,endsAt,labels,annotations,state,timeout\n      }\n    }\n  }\n}"): (typeof documents)["query msgAlertList($first: Int,$orderBy:MsgAlertOrder,$where:MsgAlertWhereInput){\n  msgAlerts(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,startsAt,endsAt,labels,annotations,state,timeout\n      }\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query msgAlertLogList($gid:GID!,$first: Int,$orderBy:NlogOrder,$where:NlogWhereInput){\n   node(id: $gid){\n    id\n    ... on MsgAlert{\n      id,\n      nlog(first:$first,orderBy: $orderBy,where: $where){\n        totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n        edges{\n          cursor,node{\n            id,sendAt,expiresAt,groupKey,receiver,receiverType\n          }\n        }\n      }\n    }\n  }\n}"): (typeof documents)["query msgAlertLogList($gid:GID!,$first: Int,$orderBy:NlogOrder,$where:NlogWhereInput){\n   node(id: $gid){\n    id\n    ... on MsgAlert{\n      id,\n      nlog(first:$first,orderBy: $orderBy,where: $where){\n        totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n        edges{\n          cursor,node{\n            id,sendAt,expiresAt,groupKey,receiver,receiverType\n          }\n        }\n      }\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query silenceList($first: Int,$orderBy:SilenceOrder,$where:SilenceWhereInput){\n  silences(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,tenantID,comments,createdAt,startsAt,endsAt,state,\n        matchers{\n          type,name,value\n        }\n\n      }\n    }\n  }\n}"): (typeof documents)["query silenceList($first: Int,$orderBy:SilenceOrder,$where:SilenceWhereInput){\n  silences(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,tenantID,comments,createdAt,startsAt,endsAt,state,\n        matchers{\n          type,name,value\n        }\n\n      }\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query SilenceInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on Silence{\n      id,tenantID,startsAt,endsAt,comments,state\n      matchers{\n        type,name,value\n      }\n    }\n  }\n}"): (typeof documents)["query SilenceInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on Silence{\n      id,tenantID,startsAt,endsAt,comments,state\n      matchers{\n        type,name,value\n      }\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation createSilence($input: CreateSilenceInput!){\n  createSilence(input: $input){id}\n}"): (typeof documents)["mutation createSilence($input: CreateSilenceInput!){\n  createSilence(input: $input){id}\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation updateSilence($id:ID!,$input: UpdateSilenceInput!){\n  updateSilence(id:$id,input: $input){id}\n}"): (typeof documents)["mutation updateSilence($id:ID!,$input: UpdateSilenceInput!){\n  updateSilence(id:$id,input: $input){id}\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation delSilence($id:ID!){\n  deleteSilence(id:$id)\n}"): (typeof documents)["mutation delSilence($id:ID!){\n  deleteSilence(id:$id)\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query msgTemplateList($first: Int,$orderBy:MsgTemplateOrder,$where:MsgTemplateWhereInput){\n  msgTemplates(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,\n        receiverType,format,subject,from,to,cc,bcc,body,tpl,attachments\n      }\n    }\n  }\n}"): (typeof documents)["query msgTemplateList($first: Int,$orderBy:MsgTemplateOrder,$where:MsgTemplateWhereInput){\n  msgTemplates(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,\n        receiverType,format,subject,from,to,cc,bcc,body,tpl,attachments\n      }\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query MsgTemplateInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgTemplate{\n      id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,\n      receiverType,format,subject,from,to,cc,bcc,body,tplFileID,attachmentsFileIds,tpl,attachments\n    }\n  }\n}"): (typeof documents)["query MsgTemplateInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgTemplate{\n      id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,\n      receiverType,format,subject,from,to,cc,bcc,body,tplFileID,attachmentsFileIds,tpl,attachments\n    }\n  }\n}"];
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
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query msgTypeCategory($keyword:String,$appID:ID){\n  msgTypeCategories(keyword: $keyword,appID:$appID)\n}"): (typeof documents)["query msgTypeCategory($keyword:String,$appID:ID){\n  msgTypeCategories(keyword: $keyword,appID:$appID)\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query msgTypeListAndSub($first: Int,$orderBy:MsgTypeOrder,$where:MsgTypeWhereInput){\n  msgTypes(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,category,\n        subscriberUsers{\n          id,tenantID,msgTypeID,userID\n        },\n        subscriberRoles{\n          id,tenantID,msgTypeID,orgRoleID\n        },\n        excludeSubscriberUsers{\n          id,tenantID,msgTypeID,userID\n        }\n      }\n    }\n  }\n}"): (typeof documents)["query msgTypeListAndSub($first: Int,$orderBy:MsgTypeOrder,$where:MsgTypeWhereInput){\n  msgTypes(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,category,\n        subscriberUsers{\n          id,tenantID,msgTypeID,userID\n        },\n        subscriberRoles{\n          id,tenantID,msgTypeID,orgRoleID\n        },\n        excludeSubscriberUsers{\n          id,tenantID,msgTypeID,userID\n        }\n      }\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query msgTypeAndSubInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgType{\n      id,name,appID,category,\n      subscriberUsers{\n        id,tenantID,msgTypeID,userID\n      },\n      subscriberRoles{\n        id,tenantID,msgTypeID,orgRoleID\n      },\n      excludeSubscriberUsers{\n        id,tenantID,msgTypeID,userID\n      }\n    }\n  }\n}"): (typeof documents)["query msgTypeAndSubInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgType{\n      id,name,appID,category,\n      subscriberUsers{\n        id,tenantID,msgTypeID,userID\n      },\n      subscriberRoles{\n        id,tenantID,msgTypeID,orgRoleID\n      },\n      excludeSubscriberUsers{\n        id,tenantID,msgTypeID,userID\n      }\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation createMsgSubscriber($inputs: [CreateMsgSubscriberInput!]!){\n  createMsgSubscriber(inputs: $inputs){id}\n}"): (typeof documents)["mutation createMsgSubscriber($inputs: [CreateMsgSubscriberInput!]!){\n  createMsgSubscriber(inputs: $inputs){id}\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation deleteMsgSubscriber($ids: [ID!]!){\n  deleteMsgSubscriber(ids: $ids)\n}"): (typeof documents)["mutation deleteMsgSubscriber($ids: [ID!]!){\n  deleteMsgSubscriber(ids: $ids)\n}"];

export function gql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;