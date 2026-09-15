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
 * Learn more about it here: https://the-guild.dev/graphql/codegen/plugins/presets/preset-client#reducing-bundle-size
 */
type Documents = {
    "query msgChannelList($first: Int,$orderBy:MsgChannelOrder,$where:MsgChannelWhereInput){\n  msgChannels(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,receiverType,tenantID,comments,status,status,createdAt\n      }\n    }\n  }\n}": typeof types.MsgChannelListDocument,
    "query msgChannelInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgChannel{\n      id,name,receiverType,tenantID,comments,status,status,createdAt\n    }\n  }\n}": typeof types.MsgChannelInfoDocument,
    "query msgChannelReceiverInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgChannel{\n      id,name,receiverType,tenantID,comments,status,status,createdAt,\n      receiver{\n        name,\n        emailConfigs{\n          authIdentity,authPassword,authSecret,authType,authUsername,from,headers,smartHost,to\n        },\n        messageConfig{\n          redirect,subject,to\n        },\n        webhookConfigs{\n          sendResolved,url,urlFile,maxAlerts,timeout,headers,subject,body\n        },\n        umengConfigs{\n          sendResolved,apiURL,apps,productionMode\n        }\n      }\n    }\n  }\n}": typeof types.MsgChannelReceiverInfoDocument,
    "mutation createMsgChannel($input: CreateMsgChannelInput!){\n createMsgChannel(input:$input){id,name,receiverType,tenantID,comments,status,status,createdAt}\n}": typeof types.CreateMsgChannelDocument,
    "mutation updateMsgChannel($id:ID!,$input: UpdateMsgChannelInput!){\n updateMsgChannel(id:$id,input:$input){id,name,receiverType,tenantID,comments,status,status,createdAt}\n}": typeof types.UpdateMsgChannelDocument,
    "mutation delMsgChannel($id:ID!){\n deleteMsgChannel(id:$id)\n}": typeof types.DelMsgChannelDocument,
    "mutation enableMsgChannel($id:ID!){\n enableMsgChannel(id:$id){id,name,receiverType,tenantID,comments,status,status,createdAt}\n}": typeof types.EnableMsgChannelDocument,
    "mutation disableMsgChannel($id:ID!){\n disableMsgChannel(id:$id){id,name,receiverType,tenantID,comments,status,status,createdAt}\n}": typeof types.DisableMsgChannelDocument,
    "query msgEventList($first: Int,$orderBy:MsgEventOrder,$where:MsgEventWhereInput){\n  msgEvents(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,comments,status,createdAt,msgTypeID,modes,canSubs\n        msgType{\n          id,category,appID,name\n        }\n      }\n    }\n  }\n}": typeof types.MsgEventListDocument,
    "query MsgEventInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgEvent{\n      id,name,comments,status,createdAt,msgTypeID,modes,canSubs\n      msgType{\n        id,category,appID,name\n      }\n    }\n  }\n}": typeof types.MsgEventInfoDocument,
    "query MsgEventInfoRoute($gid:GID!,$type:RouteStrType!){\n  node(id: $gid){\n    id\n    ... on MsgEvent{\n      id,name,comments,status,createdAt,msgTypeID,modes,canSubs,routeStr(type:$type)\n      msgType{\n        id,category,appID,name\n      }\n    }\n  }\n}\n": typeof types.MsgEventInfoRouteDocument,
    "mutation createMsgEvent($input: CreateMsgEventInput!){\n  createMsgEvent(input: $input){\n    id,name,comments,status,createdAt,msgTypeID,modes,canSubs\n    msgType{\n      id,category,appID,name\n    }\n  }\n}": typeof types.CreateMsgEventDocument,
    "mutation updateMsgEvent($id:ID!,$input: UpdateMsgEventInput!){\n  updateMsgEvent(id:$id,input: $input){\n    id,name,comments,status,createdAt,msgTypeID,modes,canSubs\n    msgType{\n      id,category,appID,name\n    }\n  }\n}": typeof types.UpdateMsgEventDocument,
    "mutation delMsgEvent($id:ID!){\n  deleteMsgEvent(id:$id)\n}": typeof types.DelMsgEventDocument,
    "mutation enableMsgEvent($id:ID!){\n  enableMsgEvent(id:$id){\n    id,name,comments,status,createdAt,msgTypeID,modes\n    msgType{\n      id,category,appID,name\n    }\n  }\n}": typeof types.EnableMsgEventDocument,
    "mutation disableMsgEvent($id:ID!){\n  disableMsgEvent(id:$id){\n    id,name,comments,status,createdAt,msgTypeID,modes\n    msgType{\n      id,category,appID,name\n    }\n  }\n}": typeof types.DisableMsgEventDocument,
    "query msgEventListWithSubs($first: Int,$orderBy:MsgEventOrder,$where:MsgEventWhereInput){\n  msgEvents(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,comments,status,createdAt,msgTypeID,modes,canSubs\n        msgType{\n          id,category,appID,name\n        }\n        subscriberUsers{\n          id\n          userID\n          orgRoleID\n          exclude\n        }\n        subscriberRoles{\n          id\n          userID\n          orgRoleID\n          exclude\n        }\n        excludeSubscriberUsers{\n          id\n          userID\n          orgRoleID\n          exclude\n        }\n      }\n    }\n  }\n}": typeof types.MsgEventListWithSubsDocument,
    "query msgEventWithSubs($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgEvent{\n      id,name,comments,status,createdAt,msgTypeID,modes,canSubs\n      msgType{\n        id,category,appID,name\n      }\n      subscriberUsers{\n        id\n        userID\n        orgRoleID\n        exclude\n      }\n      subscriberRoles{\n        id\n        userID\n        orgRoleID\n        exclude\n      }\n      excludeSubscriberUsers{\n        id\n        userID\n        orgRoleID\n        exclude\n      }\n    }\n  }\n}": typeof types.MsgEventWithSubsDocument,
    "query msgInternalList($first: Int,$orderBy: MsgInternalOrder,$where:MsgInternalWhereInput){\n  msgInternals(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,tenantID,createdBy,createdAt,subject,body,format,redirect\n      }\n    }\n  }\n}": typeof types.MsgInternalListDocument,
    "query userMsgInternalList($first: Int,$orderBy: MsgInternalToOrder,$where:MsgInternalToWhereInput){\n  userMsgInternalTos(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,msgInternalID,createdAt,deleteAt,readAt,userID\n        msgInternal{\n          id,tenantID,createdBy,createdAt,subject,body,format,redirect,category\n        }\n      }\n    }\n  }\n}": typeof types.UserMsgInternalListDocument,
    "query msgInternalInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgInternal{\n      id,tenantID,createdBy,createdAt,subject,body,format,redirect,category\n    }\n  }\n}": typeof types.MsgInternalInfoDocument,
    "query userMsgCategory{\n  userSubMsgCategory\n}": typeof types.UserMsgCategoryDocument,
    "query userMsgCategoryNum($categories:[String!]!){\n  userUnreadMsgInternalsFromMsgCategory(categories:$categories)\n}": typeof types.UserMsgCategoryNumDocument,
    "query msgInternalToInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgInternalTo{\n      id,msgInternalID,createdAt,deleteAt,readAt,userID\n      msgInternal{\n        id,tenantID,createdBy,createdAt,subject,body,format,redirect,category\n      }\n    }\n  }\n}": typeof types.MsgInternalToInfoDocument,
    "mutation markMsgRead($ids:[ID!]!,$read:Boolean!){\n  markMsgInternalToReadOrUnRead(ids:$ids,read:$read)\n}": typeof types.MarkMsgReadDocument,
    "mutation delMarkMsg($ids:[ID!]!){\n  markMsgInternalToDeleted(ids:$ids)\n}": typeof types.DelMarkMsgDocument,
    "subscription subMsg{\n  message{\n    content,extras,format,sendAt,title,url\n  }\n}": typeof types.SubMsgDocument,
    "query msgAlertList($first: Int,$orderBy:MsgAlertOrder,$where:MsgAlertWhereInput){\n  msgAlerts(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,startsAt,endsAt,labels,annotations,state,timeout\n      }\n    }\n  }\n}": typeof types.MsgAlertListDocument,
    "query formatMsgAlerts($first: Int,$alertName:String,$userID:String,$receiverType:MsgTemplateReceiverType,$orderBy:MsgAlertOrder,$where:MsgAlertWhereInput){\n  formatMsgAlerts(first:$first,alertName: $alertName,userID: $userID,receiverType: $receiverType,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,startsAt,endsAt,tenantID,state,msgEventComments,msgChannelComments,msgTemplateTitle,receiverType,users{name,email},state,receiver,hasMultiMsg,\n      }\n    }\n  }\n}": typeof types.FormatMsgAlertsDocument,
    "query formatMsgAlertMore($msgAlertID:ID!){\n  formatMsgAlertMore(msgAlertID: $msgAlertID){\n    id,startsAt,endsAt,tenantID,state,msgEventComments,msgChannelComments,msgTemplateTitle,receiverType,users{name,email},state,receiver,hasMultiMsg,\n  }\n}": typeof types.FormatMsgAlertMoreDocument,
    "query renderMsgAlert($msgAlertID:ID!,$receiver:String!){\n  renderMsgAlert(msgAlertID: $msgAlertID,receiver:$receiver)\n}": typeof types.RenderMsgAlertDocument,
    "query msgAlertLogList($gid:GID!,$first: Int,$orderBy:NlogOrder,$where:NlogWhereInput){\n   node(id: $gid){\n    id\n    ... on MsgAlert{\n      id,\n      nlog(first:$first,orderBy: $orderBy,where: $where){\n        totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n        edges{\n          cursor,node{\n            id,sendAt,expiresAt,groupKey,receiver,receiverType\n          }\n        }\n      }\n    }\n  }\n}": typeof types.MsgAlertLogListDocument,
    "query silenceList($first: Int,$orderBy:MsgSilenceOrder,$where:MsgSilenceWhereInput){\n  silences(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,tenantID,comments,createdAt,startsAt,endsAt,state,\n        matchers{\n          type,name,value\n        }\n\n      }\n    }\n  }\n}": typeof types.SilenceListDocument,
    "query SilenceInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgSilence{\n      id,tenantID,startsAt,endsAt,comments,state\n      matchers{\n        type,name,value\n      }\n    }\n  }\n}": typeof types.SilenceInfoDocument,
    "mutation createMsgSilence($input: CreateMsgSilenceInput!){\n  createMsgSilence(input: $input){\n    id,tenantID,comments,createdAt,startsAt,endsAt,state,\n    matchers{\n      type,name,value\n    }\n  }\n}": typeof types.CreateMsgSilenceDocument,
    "mutation updateMsgSilence($id:ID!,$input: UpdateMsgSilenceInput!){\n  updateMsgSilence(id:$id,input: $input){\n    id,tenantID,comments,createdAt,startsAt,endsAt,state,\n    matchers{\n      type,name,value\n    }\n  }\n}": typeof types.UpdateMsgSilenceDocument,
    "mutation delMsgSilence($id:ID!){\n  deleteMsgSilence(id:$id)\n}": typeof types.DelMsgSilenceDocument,
    "query msgTemplateList($first: Int,$orderBy:MsgTemplateOrder,$where:MsgTemplateWhereInput){\n  msgTemplates(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,userID,\n        receiverType,format,subject,from,to,cc,bcc,body,tpl,attachments\n      }\n    }\n  }\n}": typeof types.MsgTemplateListDocument,
    "query MsgTemplateInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgTemplate{\n      id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,userID,\n      receiverType,format,subject,from,to,cc,bcc,body,tpl,attachments\n    }\n  }\n}": typeof types.MsgTemplateInfoDocument,
    "query msgTemplateDefineByName($format:MsgTemplateFormat!,$body:String!){\n  msgTemplateDefineByName(format: $format,body: $body)\n}": typeof types.MsgTemplateDefineByNameDocument,
    "mutation createMsgTemplate($input: CreateMsgTemplateInput!){\n  createMsgTemplate(input: $input){\n    id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,userID,\n    receiverType,format,subject,from,to,cc,bcc,body,tpl,attachments\n  }\n}": typeof types.CreateMsgTemplateDocument,
    "mutation updateMsgTemplate($id:ID!,$input: UpdateMsgTemplateInput!){\n  updateMsgTemplate(id:$id,input: $input){\n    id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,userID,\n    receiverType,format,subject,from,to,cc,bcc,body,tpl,attachments\n  }\n}": typeof types.UpdateMsgTemplateDocument,
    "mutation delMsgTemplate($id:ID!){\n  deleteMsgTemplate(id:$id)\n}": typeof types.DelMsgTemplateDocument,
    "mutation enableMsgTemplate($id:ID!){\n  enableMsgTemplate(id:$id){\n    id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,userID,\n    receiverType,format,subject,from,to,cc,bcc,body,tpl,attachments\n  }\n}": typeof types.EnableMsgTemplateDocument,
    "mutation disableMsgTemplate($id:ID!){\n  disableMsgTemplate(id:$id){\n    id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,userID,\n    receiverType,format,subject,from,to,cc,bcc,body,tpl,attachments\n  }\n}": typeof types.DisableMsgTemplateDocument,
    "mutation testSendEmailTpl($annotations: MapString, $email: String!, $labels: MapString, $tplID: ID!){\n  testSendEmailTpl(annotations: $annotations, email: $email, labels:$labels, tplID: $tplID)\n}": typeof types.TestSendEmailTplDocument,
    "mutation testSendMessageTpl($annotations: MapString, $userID: ID!, $labels: MapString, $tplID: ID!){\n  testSendMessageTpl(annotations: $annotations, userID: $userID, labels:$labels, tplID: $tplID)\n}": typeof types.TestSendMessageTplDocument,
    "mutation refreshTemplateParams{\n  refreshTemplateParams\n}": typeof types.RefreshTemplateParamsDocument,
    "query msgTypeList($first: Int,$orderBy:MsgTypeOrder,$where:MsgTypeWhereInput){\n  msgTypes(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,comments,appID,status,category,canSubs,canCustom,createdAt\n      }\n    }\n  }\n}": typeof types.MsgTypeListDocument,
    "query msgTypeInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgType{\n      id,name,comments,appID,status,category,canSubs,canCustom,createdAt\n    }\n  }\n}": typeof types.MsgTypeInfoDocument,
    "mutation createMsgType($input: CreateMsgTypeInput!){\n  createMsgType(input: $input){\n    id,name,comments,appID,status,category,canSubs,canCustom,createdAt\n  }\n}": typeof types.CreateMsgTypeDocument,
    "mutation updateMsgType($id:ID!,$input: UpdateMsgTypeInput!){\n  updateMsgType(id:$id,input: $input){\n    id,name,comments,appID,status,category,canSubs,canCustom,createdAt\n  }\n}": typeof types.UpdateMsgTypeDocument,
    "mutation delMsgType($id:ID!){\n  deleteMsgType(id:$id)\n}": typeof types.DelMsgTypeDocument,
    "query msgTypeCategory($keyword:String,$appID:ID){\n  msgTypeCategories(keyword: $keyword,appID:$appID)\n}": typeof types.MsgTypeCategoryDocument,
    "query msgTypeListAndSub($first: Int,$orderBy:MsgTypeOrder,$where:MsgTypeWhereInput){\n  msgTypes(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,comments,appID,status,category,canSubs,canCustom,createdAt\n        subscriberUsers{\n          id,tenantID,msgTypeID,userID\n        },\n        subscriberRoles{\n          id,tenantID,msgTypeID,orgRoleID\n        },\n        excludeSubscriberUsers{\n          id,tenantID,msgTypeID,userID\n        }\n      }\n    }\n  }\n}": typeof types.MsgTypeListAndSubDocument,
    "query msgTypeAndSubInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgType{\n      id,name,comments,appID,status,category,canSubs,canCustom,createdAt\n      subscriberUsers{\n        id,tenantID,msgTypeID,userID\n      },\n      subscriberRoles{\n        id,tenantID,msgTypeID,orgRoleID\n      },\n      excludeSubscriberUsers{\n        id,tenantID,msgTypeID,userID\n      }\n    }\n  }\n}": typeof types.MsgTypeAndSubInfoDocument,
    "mutation createMsgSubscriber($inputs: [CreateMsgSubscriberInput!]!){\n  createMsgSubscriber(inputs: $inputs){\n    id,tenantID,msgTypeID,userID\n  }\n}": typeof types.CreateMsgSubscriberDocument,
    "mutation deleteMsgSubscriber($ids: [ID!]!){\n  deleteMsgSubscriber(ids: $ids)\n}": typeof types.DeleteMsgSubscriberDocument,
};
const documents: Documents = {
    "query msgChannelList($first: Int,$orderBy:MsgChannelOrder,$where:MsgChannelWhereInput){\n  msgChannels(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,receiverType,tenantID,comments,status,status,createdAt\n      }\n    }\n  }\n}": types.MsgChannelListDocument,
    "query msgChannelInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgChannel{\n      id,name,receiverType,tenantID,comments,status,status,createdAt\n    }\n  }\n}": types.MsgChannelInfoDocument,
    "query msgChannelReceiverInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgChannel{\n      id,name,receiverType,tenantID,comments,status,status,createdAt,\n      receiver{\n        name,\n        emailConfigs{\n          authIdentity,authPassword,authSecret,authType,authUsername,from,headers,smartHost,to\n        },\n        messageConfig{\n          redirect,subject,to\n        },\n        webhookConfigs{\n          sendResolved,url,urlFile,maxAlerts,timeout,headers,subject,body\n        },\n        umengConfigs{\n          sendResolved,apiURL,apps,productionMode\n        }\n      }\n    }\n  }\n}": types.MsgChannelReceiverInfoDocument,
    "mutation createMsgChannel($input: CreateMsgChannelInput!){\n createMsgChannel(input:$input){id,name,receiverType,tenantID,comments,status,status,createdAt}\n}": types.CreateMsgChannelDocument,
    "mutation updateMsgChannel($id:ID!,$input: UpdateMsgChannelInput!){\n updateMsgChannel(id:$id,input:$input){id,name,receiverType,tenantID,comments,status,status,createdAt}\n}": types.UpdateMsgChannelDocument,
    "mutation delMsgChannel($id:ID!){\n deleteMsgChannel(id:$id)\n}": types.DelMsgChannelDocument,
    "mutation enableMsgChannel($id:ID!){\n enableMsgChannel(id:$id){id,name,receiverType,tenantID,comments,status,status,createdAt}\n}": types.EnableMsgChannelDocument,
    "mutation disableMsgChannel($id:ID!){\n disableMsgChannel(id:$id){id,name,receiverType,tenantID,comments,status,status,createdAt}\n}": types.DisableMsgChannelDocument,
    "query msgEventList($first: Int,$orderBy:MsgEventOrder,$where:MsgEventWhereInput){\n  msgEvents(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,comments,status,createdAt,msgTypeID,modes,canSubs\n        msgType{\n          id,category,appID,name\n        }\n      }\n    }\n  }\n}": types.MsgEventListDocument,
    "query MsgEventInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgEvent{\n      id,name,comments,status,createdAt,msgTypeID,modes,canSubs\n      msgType{\n        id,category,appID,name\n      }\n    }\n  }\n}": types.MsgEventInfoDocument,
    "query MsgEventInfoRoute($gid:GID!,$type:RouteStrType!){\n  node(id: $gid){\n    id\n    ... on MsgEvent{\n      id,name,comments,status,createdAt,msgTypeID,modes,canSubs,routeStr(type:$type)\n      msgType{\n        id,category,appID,name\n      }\n    }\n  }\n}\n": types.MsgEventInfoRouteDocument,
    "mutation createMsgEvent($input: CreateMsgEventInput!){\n  createMsgEvent(input: $input){\n    id,name,comments,status,createdAt,msgTypeID,modes,canSubs\n    msgType{\n      id,category,appID,name\n    }\n  }\n}": types.CreateMsgEventDocument,
    "mutation updateMsgEvent($id:ID!,$input: UpdateMsgEventInput!){\n  updateMsgEvent(id:$id,input: $input){\n    id,name,comments,status,createdAt,msgTypeID,modes,canSubs\n    msgType{\n      id,category,appID,name\n    }\n  }\n}": types.UpdateMsgEventDocument,
    "mutation delMsgEvent($id:ID!){\n  deleteMsgEvent(id:$id)\n}": types.DelMsgEventDocument,
    "mutation enableMsgEvent($id:ID!){\n  enableMsgEvent(id:$id){\n    id,name,comments,status,createdAt,msgTypeID,modes\n    msgType{\n      id,category,appID,name\n    }\n  }\n}": types.EnableMsgEventDocument,
    "mutation disableMsgEvent($id:ID!){\n  disableMsgEvent(id:$id){\n    id,name,comments,status,createdAt,msgTypeID,modes\n    msgType{\n      id,category,appID,name\n    }\n  }\n}": types.DisableMsgEventDocument,
    "query msgEventListWithSubs($first: Int,$orderBy:MsgEventOrder,$where:MsgEventWhereInput){\n  msgEvents(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,comments,status,createdAt,msgTypeID,modes,canSubs\n        msgType{\n          id,category,appID,name\n        }\n        subscriberUsers{\n          id\n          userID\n          orgRoleID\n          exclude\n        }\n        subscriberRoles{\n          id\n          userID\n          orgRoleID\n          exclude\n        }\n        excludeSubscriberUsers{\n          id\n          userID\n          orgRoleID\n          exclude\n        }\n      }\n    }\n  }\n}": types.MsgEventListWithSubsDocument,
    "query msgEventWithSubs($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgEvent{\n      id,name,comments,status,createdAt,msgTypeID,modes,canSubs\n      msgType{\n        id,category,appID,name\n      }\n      subscriberUsers{\n        id\n        userID\n        orgRoleID\n        exclude\n      }\n      subscriberRoles{\n        id\n        userID\n        orgRoleID\n        exclude\n      }\n      excludeSubscriberUsers{\n        id\n        userID\n        orgRoleID\n        exclude\n      }\n    }\n  }\n}": types.MsgEventWithSubsDocument,
    "query msgInternalList($first: Int,$orderBy: MsgInternalOrder,$where:MsgInternalWhereInput){\n  msgInternals(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,tenantID,createdBy,createdAt,subject,body,format,redirect\n      }\n    }\n  }\n}": types.MsgInternalListDocument,
    "query userMsgInternalList($first: Int,$orderBy: MsgInternalToOrder,$where:MsgInternalToWhereInput){\n  userMsgInternalTos(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,msgInternalID,createdAt,deleteAt,readAt,userID\n        msgInternal{\n          id,tenantID,createdBy,createdAt,subject,body,format,redirect,category\n        }\n      }\n    }\n  }\n}": types.UserMsgInternalListDocument,
    "query msgInternalInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgInternal{\n      id,tenantID,createdBy,createdAt,subject,body,format,redirect,category\n    }\n  }\n}": types.MsgInternalInfoDocument,
    "query userMsgCategory{\n  userSubMsgCategory\n}": types.UserMsgCategoryDocument,
    "query userMsgCategoryNum($categories:[String!]!){\n  userUnreadMsgInternalsFromMsgCategory(categories:$categories)\n}": types.UserMsgCategoryNumDocument,
    "query msgInternalToInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgInternalTo{\n      id,msgInternalID,createdAt,deleteAt,readAt,userID\n      msgInternal{\n        id,tenantID,createdBy,createdAt,subject,body,format,redirect,category\n      }\n    }\n  }\n}": types.MsgInternalToInfoDocument,
    "mutation markMsgRead($ids:[ID!]!,$read:Boolean!){\n  markMsgInternalToReadOrUnRead(ids:$ids,read:$read)\n}": types.MarkMsgReadDocument,
    "mutation delMarkMsg($ids:[ID!]!){\n  markMsgInternalToDeleted(ids:$ids)\n}": types.DelMarkMsgDocument,
    "subscription subMsg{\n  message{\n    content,extras,format,sendAt,title,url\n  }\n}": types.SubMsgDocument,
    "query msgAlertList($first: Int,$orderBy:MsgAlertOrder,$where:MsgAlertWhereInput){\n  msgAlerts(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,startsAt,endsAt,labels,annotations,state,timeout\n      }\n    }\n  }\n}": types.MsgAlertListDocument,
    "query formatMsgAlerts($first: Int,$alertName:String,$userID:String,$receiverType:MsgTemplateReceiverType,$orderBy:MsgAlertOrder,$where:MsgAlertWhereInput){\n  formatMsgAlerts(first:$first,alertName: $alertName,userID: $userID,receiverType: $receiverType,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,startsAt,endsAt,tenantID,state,msgEventComments,msgChannelComments,msgTemplateTitle,receiverType,users{name,email},state,receiver,hasMultiMsg,\n      }\n    }\n  }\n}": types.FormatMsgAlertsDocument,
    "query formatMsgAlertMore($msgAlertID:ID!){\n  formatMsgAlertMore(msgAlertID: $msgAlertID){\n    id,startsAt,endsAt,tenantID,state,msgEventComments,msgChannelComments,msgTemplateTitle,receiverType,users{name,email},state,receiver,hasMultiMsg,\n  }\n}": types.FormatMsgAlertMoreDocument,
    "query renderMsgAlert($msgAlertID:ID!,$receiver:String!){\n  renderMsgAlert(msgAlertID: $msgAlertID,receiver:$receiver)\n}": types.RenderMsgAlertDocument,
    "query msgAlertLogList($gid:GID!,$first: Int,$orderBy:NlogOrder,$where:NlogWhereInput){\n   node(id: $gid){\n    id\n    ... on MsgAlert{\n      id,\n      nlog(first:$first,orderBy: $orderBy,where: $where){\n        totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n        edges{\n          cursor,node{\n            id,sendAt,expiresAt,groupKey,receiver,receiverType\n          }\n        }\n      }\n    }\n  }\n}": types.MsgAlertLogListDocument,
    "query silenceList($first: Int,$orderBy:MsgSilenceOrder,$where:MsgSilenceWhereInput){\n  silences(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,tenantID,comments,createdAt,startsAt,endsAt,state,\n        matchers{\n          type,name,value\n        }\n\n      }\n    }\n  }\n}": types.SilenceListDocument,
    "query SilenceInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgSilence{\n      id,tenantID,startsAt,endsAt,comments,state\n      matchers{\n        type,name,value\n      }\n    }\n  }\n}": types.SilenceInfoDocument,
    "mutation createMsgSilence($input: CreateMsgSilenceInput!){\n  createMsgSilence(input: $input){\n    id,tenantID,comments,createdAt,startsAt,endsAt,state,\n    matchers{\n      type,name,value\n    }\n  }\n}": types.CreateMsgSilenceDocument,
    "mutation updateMsgSilence($id:ID!,$input: UpdateMsgSilenceInput!){\n  updateMsgSilence(id:$id,input: $input){\n    id,tenantID,comments,createdAt,startsAt,endsAt,state,\n    matchers{\n      type,name,value\n    }\n  }\n}": types.UpdateMsgSilenceDocument,
    "mutation delMsgSilence($id:ID!){\n  deleteMsgSilence(id:$id)\n}": types.DelMsgSilenceDocument,
    "query msgTemplateList($first: Int,$orderBy:MsgTemplateOrder,$where:MsgTemplateWhereInput){\n  msgTemplates(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,userID,\n        receiverType,format,subject,from,to,cc,bcc,body,tpl,attachments\n      }\n    }\n  }\n}": types.MsgTemplateListDocument,
    "query MsgTemplateInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgTemplate{\n      id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,userID,\n      receiverType,format,subject,from,to,cc,bcc,body,tpl,attachments\n    }\n  }\n}": types.MsgTemplateInfoDocument,
    "query msgTemplateDefineByName($format:MsgTemplateFormat!,$body:String!){\n  msgTemplateDefineByName(format: $format,body: $body)\n}": types.MsgTemplateDefineByNameDocument,
    "mutation createMsgTemplate($input: CreateMsgTemplateInput!){\n  createMsgTemplate(input: $input){\n    id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,userID,\n    receiverType,format,subject,from,to,cc,bcc,body,tpl,attachments\n  }\n}": types.CreateMsgTemplateDocument,
    "mutation updateMsgTemplate($id:ID!,$input: UpdateMsgTemplateInput!){\n  updateMsgTemplate(id:$id,input: $input){\n    id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,userID,\n    receiverType,format,subject,from,to,cc,bcc,body,tpl,attachments\n  }\n}": types.UpdateMsgTemplateDocument,
    "mutation delMsgTemplate($id:ID!){\n  deleteMsgTemplate(id:$id)\n}": types.DelMsgTemplateDocument,
    "mutation enableMsgTemplate($id:ID!){\n  enableMsgTemplate(id:$id){\n    id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,userID,\n    receiverType,format,subject,from,to,cc,bcc,body,tpl,attachments\n  }\n}": types.EnableMsgTemplateDocument,
    "mutation disableMsgTemplate($id:ID!){\n  disableMsgTemplate(id:$id){\n    id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,userID,\n    receiverType,format,subject,from,to,cc,bcc,body,tpl,attachments\n  }\n}": types.DisableMsgTemplateDocument,
    "mutation testSendEmailTpl($annotations: MapString, $email: String!, $labels: MapString, $tplID: ID!){\n  testSendEmailTpl(annotations: $annotations, email: $email, labels:$labels, tplID: $tplID)\n}": types.TestSendEmailTplDocument,
    "mutation testSendMessageTpl($annotations: MapString, $userID: ID!, $labels: MapString, $tplID: ID!){\n  testSendMessageTpl(annotations: $annotations, userID: $userID, labels:$labels, tplID: $tplID)\n}": types.TestSendMessageTplDocument,
    "mutation refreshTemplateParams{\n  refreshTemplateParams\n}": types.RefreshTemplateParamsDocument,
    "query msgTypeList($first: Int,$orderBy:MsgTypeOrder,$where:MsgTypeWhereInput){\n  msgTypes(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,comments,appID,status,category,canSubs,canCustom,createdAt\n      }\n    }\n  }\n}": types.MsgTypeListDocument,
    "query msgTypeInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgType{\n      id,name,comments,appID,status,category,canSubs,canCustom,createdAt\n    }\n  }\n}": types.MsgTypeInfoDocument,
    "mutation createMsgType($input: CreateMsgTypeInput!){\n  createMsgType(input: $input){\n    id,name,comments,appID,status,category,canSubs,canCustom,createdAt\n  }\n}": types.CreateMsgTypeDocument,
    "mutation updateMsgType($id:ID!,$input: UpdateMsgTypeInput!){\n  updateMsgType(id:$id,input: $input){\n    id,name,comments,appID,status,category,canSubs,canCustom,createdAt\n  }\n}": types.UpdateMsgTypeDocument,
    "mutation delMsgType($id:ID!){\n  deleteMsgType(id:$id)\n}": types.DelMsgTypeDocument,
    "query msgTypeCategory($keyword:String,$appID:ID){\n  msgTypeCategories(keyword: $keyword,appID:$appID)\n}": types.MsgTypeCategoryDocument,
    "query msgTypeListAndSub($first: Int,$orderBy:MsgTypeOrder,$where:MsgTypeWhereInput){\n  msgTypes(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,comments,appID,status,category,canSubs,canCustom,createdAt\n        subscriberUsers{\n          id,tenantID,msgTypeID,userID\n        },\n        subscriberRoles{\n          id,tenantID,msgTypeID,orgRoleID\n        },\n        excludeSubscriberUsers{\n          id,tenantID,msgTypeID,userID\n        }\n      }\n    }\n  }\n}": types.MsgTypeListAndSubDocument,
    "query msgTypeAndSubInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgType{\n      id,name,comments,appID,status,category,canSubs,canCustom,createdAt\n      subscriberUsers{\n        id,tenantID,msgTypeID,userID\n      },\n      subscriberRoles{\n        id,tenantID,msgTypeID,orgRoleID\n      },\n      excludeSubscriberUsers{\n        id,tenantID,msgTypeID,userID\n      }\n    }\n  }\n}": types.MsgTypeAndSubInfoDocument,
    "mutation createMsgSubscriber($inputs: [CreateMsgSubscriberInput!]!){\n  createMsgSubscriber(inputs: $inputs){\n    id,tenantID,msgTypeID,userID\n  }\n}": types.CreateMsgSubscriberDocument,
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
export function gql(source: "query msgChannelReceiverInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgChannel{\n      id,name,receiverType,tenantID,comments,status,status,createdAt,\n      receiver{\n        name,\n        emailConfigs{\n          authIdentity,authPassword,authSecret,authType,authUsername,from,headers,smartHost,to\n        },\n        messageConfig{\n          redirect,subject,to\n        },\n        webhookConfigs{\n          sendResolved,url,urlFile,maxAlerts,timeout,headers,subject,body\n        },\n        umengConfigs{\n          sendResolved,apiURL,apps,productionMode\n        }\n      }\n    }\n  }\n}"): (typeof documents)["query msgChannelReceiverInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgChannel{\n      id,name,receiverType,tenantID,comments,status,status,createdAt,\n      receiver{\n        name,\n        emailConfigs{\n          authIdentity,authPassword,authSecret,authType,authUsername,from,headers,smartHost,to\n        },\n        messageConfig{\n          redirect,subject,to\n        },\n        webhookConfigs{\n          sendResolved,url,urlFile,maxAlerts,timeout,headers,subject,body\n        },\n        umengConfigs{\n          sendResolved,apiURL,apps,productionMode\n        }\n      }\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation createMsgChannel($input: CreateMsgChannelInput!){\n createMsgChannel(input:$input){id,name,receiverType,tenantID,comments,status,status,createdAt}\n}"): (typeof documents)["mutation createMsgChannel($input: CreateMsgChannelInput!){\n createMsgChannel(input:$input){id,name,receiverType,tenantID,comments,status,status,createdAt}\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation updateMsgChannel($id:ID!,$input: UpdateMsgChannelInput!){\n updateMsgChannel(id:$id,input:$input){id,name,receiverType,tenantID,comments,status,status,createdAt}\n}"): (typeof documents)["mutation updateMsgChannel($id:ID!,$input: UpdateMsgChannelInput!){\n updateMsgChannel(id:$id,input:$input){id,name,receiverType,tenantID,comments,status,status,createdAt}\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation delMsgChannel($id:ID!){\n deleteMsgChannel(id:$id)\n}"): (typeof documents)["mutation delMsgChannel($id:ID!){\n deleteMsgChannel(id:$id)\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation enableMsgChannel($id:ID!){\n enableMsgChannel(id:$id){id,name,receiverType,tenantID,comments,status,status,createdAt}\n}"): (typeof documents)["mutation enableMsgChannel($id:ID!){\n enableMsgChannel(id:$id){id,name,receiverType,tenantID,comments,status,status,createdAt}\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation disableMsgChannel($id:ID!){\n disableMsgChannel(id:$id){id,name,receiverType,tenantID,comments,status,status,createdAt}\n}"): (typeof documents)["mutation disableMsgChannel($id:ID!){\n disableMsgChannel(id:$id){id,name,receiverType,tenantID,comments,status,status,createdAt}\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query msgEventList($first: Int,$orderBy:MsgEventOrder,$where:MsgEventWhereInput){\n  msgEvents(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,comments,status,createdAt,msgTypeID,modes,canSubs\n        msgType{\n          id,category,appID,name\n        }\n      }\n    }\n  }\n}"): (typeof documents)["query msgEventList($first: Int,$orderBy:MsgEventOrder,$where:MsgEventWhereInput){\n  msgEvents(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,comments,status,createdAt,msgTypeID,modes,canSubs\n        msgType{\n          id,category,appID,name\n        }\n      }\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query MsgEventInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgEvent{\n      id,name,comments,status,createdAt,msgTypeID,modes,canSubs\n      msgType{\n        id,category,appID,name\n      }\n    }\n  }\n}"): (typeof documents)["query MsgEventInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgEvent{\n      id,name,comments,status,createdAt,msgTypeID,modes,canSubs\n      msgType{\n        id,category,appID,name\n      }\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query MsgEventInfoRoute($gid:GID!,$type:RouteStrType!){\n  node(id: $gid){\n    id\n    ... on MsgEvent{\n      id,name,comments,status,createdAt,msgTypeID,modes,canSubs,routeStr(type:$type)\n      msgType{\n        id,category,appID,name\n      }\n    }\n  }\n}\n"): (typeof documents)["query MsgEventInfoRoute($gid:GID!,$type:RouteStrType!){\n  node(id: $gid){\n    id\n    ... on MsgEvent{\n      id,name,comments,status,createdAt,msgTypeID,modes,canSubs,routeStr(type:$type)\n      msgType{\n        id,category,appID,name\n      }\n    }\n  }\n}\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation createMsgEvent($input: CreateMsgEventInput!){\n  createMsgEvent(input: $input){\n    id,name,comments,status,createdAt,msgTypeID,modes,canSubs\n    msgType{\n      id,category,appID,name\n    }\n  }\n}"): (typeof documents)["mutation createMsgEvent($input: CreateMsgEventInput!){\n  createMsgEvent(input: $input){\n    id,name,comments,status,createdAt,msgTypeID,modes,canSubs\n    msgType{\n      id,category,appID,name\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation updateMsgEvent($id:ID!,$input: UpdateMsgEventInput!){\n  updateMsgEvent(id:$id,input: $input){\n    id,name,comments,status,createdAt,msgTypeID,modes,canSubs\n    msgType{\n      id,category,appID,name\n    }\n  }\n}"): (typeof documents)["mutation updateMsgEvent($id:ID!,$input: UpdateMsgEventInput!){\n  updateMsgEvent(id:$id,input: $input){\n    id,name,comments,status,createdAt,msgTypeID,modes,canSubs\n    msgType{\n      id,category,appID,name\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation delMsgEvent($id:ID!){\n  deleteMsgEvent(id:$id)\n}"): (typeof documents)["mutation delMsgEvent($id:ID!){\n  deleteMsgEvent(id:$id)\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation enableMsgEvent($id:ID!){\n  enableMsgEvent(id:$id){\n    id,name,comments,status,createdAt,msgTypeID,modes\n    msgType{\n      id,category,appID,name\n    }\n  }\n}"): (typeof documents)["mutation enableMsgEvent($id:ID!){\n  enableMsgEvent(id:$id){\n    id,name,comments,status,createdAt,msgTypeID,modes\n    msgType{\n      id,category,appID,name\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation disableMsgEvent($id:ID!){\n  disableMsgEvent(id:$id){\n    id,name,comments,status,createdAt,msgTypeID,modes\n    msgType{\n      id,category,appID,name\n    }\n  }\n}"): (typeof documents)["mutation disableMsgEvent($id:ID!){\n  disableMsgEvent(id:$id){\n    id,name,comments,status,createdAt,msgTypeID,modes\n    msgType{\n      id,category,appID,name\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query msgEventListWithSubs($first: Int,$orderBy:MsgEventOrder,$where:MsgEventWhereInput){\n  msgEvents(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,comments,status,createdAt,msgTypeID,modes,canSubs\n        msgType{\n          id,category,appID,name\n        }\n        subscriberUsers{\n          id\n          userID\n          orgRoleID\n          exclude\n        }\n        subscriberRoles{\n          id\n          userID\n          orgRoleID\n          exclude\n        }\n        excludeSubscriberUsers{\n          id\n          userID\n          orgRoleID\n          exclude\n        }\n      }\n    }\n  }\n}"): (typeof documents)["query msgEventListWithSubs($first: Int,$orderBy:MsgEventOrder,$where:MsgEventWhereInput){\n  msgEvents(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,comments,status,createdAt,msgTypeID,modes,canSubs\n        msgType{\n          id,category,appID,name\n        }\n        subscriberUsers{\n          id\n          userID\n          orgRoleID\n          exclude\n        }\n        subscriberRoles{\n          id\n          userID\n          orgRoleID\n          exclude\n        }\n        excludeSubscriberUsers{\n          id\n          userID\n          orgRoleID\n          exclude\n        }\n      }\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query msgEventWithSubs($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgEvent{\n      id,name,comments,status,createdAt,msgTypeID,modes,canSubs\n      msgType{\n        id,category,appID,name\n      }\n      subscriberUsers{\n        id\n        userID\n        orgRoleID\n        exclude\n      }\n      subscriberRoles{\n        id\n        userID\n        orgRoleID\n        exclude\n      }\n      excludeSubscriberUsers{\n        id\n        userID\n        orgRoleID\n        exclude\n      }\n    }\n  }\n}"): (typeof documents)["query msgEventWithSubs($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgEvent{\n      id,name,comments,status,createdAt,msgTypeID,modes,canSubs\n      msgType{\n        id,category,appID,name\n      }\n      subscriberUsers{\n        id\n        userID\n        orgRoleID\n        exclude\n      }\n      subscriberRoles{\n        id\n        userID\n        orgRoleID\n        exclude\n      }\n      excludeSubscriberUsers{\n        id\n        userID\n        orgRoleID\n        exclude\n      }\n    }\n  }\n}"];
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
export function gql(source: "query msgInternalInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgInternal{\n      id,tenantID,createdBy,createdAt,subject,body,format,redirect,category\n    }\n  }\n}"): (typeof documents)["query msgInternalInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgInternal{\n      id,tenantID,createdBy,createdAt,subject,body,format,redirect,category\n    }\n  }\n}"];
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
export function gql(source: "query formatMsgAlerts($first: Int,$alertName:String,$userID:String,$receiverType:MsgTemplateReceiverType,$orderBy:MsgAlertOrder,$where:MsgAlertWhereInput){\n  formatMsgAlerts(first:$first,alertName: $alertName,userID: $userID,receiverType: $receiverType,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,startsAt,endsAt,tenantID,state,msgEventComments,msgChannelComments,msgTemplateTitle,receiverType,users{name,email},state,receiver,hasMultiMsg,\n      }\n    }\n  }\n}"): (typeof documents)["query formatMsgAlerts($first: Int,$alertName:String,$userID:String,$receiverType:MsgTemplateReceiverType,$orderBy:MsgAlertOrder,$where:MsgAlertWhereInput){\n  formatMsgAlerts(first:$first,alertName: $alertName,userID: $userID,receiverType: $receiverType,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,startsAt,endsAt,tenantID,state,msgEventComments,msgChannelComments,msgTemplateTitle,receiverType,users{name,email},state,receiver,hasMultiMsg,\n      }\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query formatMsgAlertMore($msgAlertID:ID!){\n  formatMsgAlertMore(msgAlertID: $msgAlertID){\n    id,startsAt,endsAt,tenantID,state,msgEventComments,msgChannelComments,msgTemplateTitle,receiverType,users{name,email},state,receiver,hasMultiMsg,\n  }\n}"): (typeof documents)["query formatMsgAlertMore($msgAlertID:ID!){\n  formatMsgAlertMore(msgAlertID: $msgAlertID){\n    id,startsAt,endsAt,tenantID,state,msgEventComments,msgChannelComments,msgTemplateTitle,receiverType,users{name,email},state,receiver,hasMultiMsg,\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query renderMsgAlert($msgAlertID:ID!,$receiver:String!){\n  renderMsgAlert(msgAlertID: $msgAlertID,receiver:$receiver)\n}"): (typeof documents)["query renderMsgAlert($msgAlertID:ID!,$receiver:String!){\n  renderMsgAlert(msgAlertID: $msgAlertID,receiver:$receiver)\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query msgAlertLogList($gid:GID!,$first: Int,$orderBy:NlogOrder,$where:NlogWhereInput){\n   node(id: $gid){\n    id\n    ... on MsgAlert{\n      id,\n      nlog(first:$first,orderBy: $orderBy,where: $where){\n        totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n        edges{\n          cursor,node{\n            id,sendAt,expiresAt,groupKey,receiver,receiverType\n          }\n        }\n      }\n    }\n  }\n}"): (typeof documents)["query msgAlertLogList($gid:GID!,$first: Int,$orderBy:NlogOrder,$where:NlogWhereInput){\n   node(id: $gid){\n    id\n    ... on MsgAlert{\n      id,\n      nlog(first:$first,orderBy: $orderBy,where: $where){\n        totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n        edges{\n          cursor,node{\n            id,sendAt,expiresAt,groupKey,receiver,receiverType\n          }\n        }\n      }\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query silenceList($first: Int,$orderBy:MsgSilenceOrder,$where:MsgSilenceWhereInput){\n  silences(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,tenantID,comments,createdAt,startsAt,endsAt,state,\n        matchers{\n          type,name,value\n        }\n\n      }\n    }\n  }\n}"): (typeof documents)["query silenceList($first: Int,$orderBy:MsgSilenceOrder,$where:MsgSilenceWhereInput){\n  silences(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,tenantID,comments,createdAt,startsAt,endsAt,state,\n        matchers{\n          type,name,value\n        }\n\n      }\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query SilenceInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgSilence{\n      id,tenantID,startsAt,endsAt,comments,state\n      matchers{\n        type,name,value\n      }\n    }\n  }\n}"): (typeof documents)["query SilenceInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgSilence{\n      id,tenantID,startsAt,endsAt,comments,state\n      matchers{\n        type,name,value\n      }\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation createMsgSilence($input: CreateMsgSilenceInput!){\n  createMsgSilence(input: $input){\n    id,tenantID,comments,createdAt,startsAt,endsAt,state,\n    matchers{\n      type,name,value\n    }\n  }\n}"): (typeof documents)["mutation createMsgSilence($input: CreateMsgSilenceInput!){\n  createMsgSilence(input: $input){\n    id,tenantID,comments,createdAt,startsAt,endsAt,state,\n    matchers{\n      type,name,value\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation updateMsgSilence($id:ID!,$input: UpdateMsgSilenceInput!){\n  updateMsgSilence(id:$id,input: $input){\n    id,tenantID,comments,createdAt,startsAt,endsAt,state,\n    matchers{\n      type,name,value\n    }\n  }\n}"): (typeof documents)["mutation updateMsgSilence($id:ID!,$input: UpdateMsgSilenceInput!){\n  updateMsgSilence(id:$id,input: $input){\n    id,tenantID,comments,createdAt,startsAt,endsAt,state,\n    matchers{\n      type,name,value\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation delMsgSilence($id:ID!){\n  deleteMsgSilence(id:$id)\n}"): (typeof documents)["mutation delMsgSilence($id:ID!){\n  deleteMsgSilence(id:$id)\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query msgTemplateList($first: Int,$orderBy:MsgTemplateOrder,$where:MsgTemplateWhereInput){\n  msgTemplates(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,userID,\n        receiverType,format,subject,from,to,cc,bcc,body,tpl,attachments\n      }\n    }\n  }\n}"): (typeof documents)["query msgTemplateList($first: Int,$orderBy:MsgTemplateOrder,$where:MsgTemplateWhereInput){\n  msgTemplates(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,userID,\n        receiverType,format,subject,from,to,cc,bcc,body,tpl,attachments\n      }\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query MsgTemplateInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgTemplate{\n      id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,userID,\n      receiverType,format,subject,from,to,cc,bcc,body,tpl,attachments\n    }\n  }\n}"): (typeof documents)["query MsgTemplateInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgTemplate{\n      id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,userID,\n      receiverType,format,subject,from,to,cc,bcc,body,tpl,attachments\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query msgTemplateDefineByName($format:MsgTemplateFormat!,$body:String!){\n  msgTemplateDefineByName(format: $format,body: $body)\n}"): (typeof documents)["query msgTemplateDefineByName($format:MsgTemplateFormat!,$body:String!){\n  msgTemplateDefineByName(format: $format,body: $body)\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation createMsgTemplate($input: CreateMsgTemplateInput!){\n  createMsgTemplate(input: $input){\n    id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,userID,\n    receiverType,format,subject,from,to,cc,bcc,body,tpl,attachments\n  }\n}"): (typeof documents)["mutation createMsgTemplate($input: CreateMsgTemplateInput!){\n  createMsgTemplate(input: $input){\n    id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,userID,\n    receiverType,format,subject,from,to,cc,bcc,body,tpl,attachments\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation updateMsgTemplate($id:ID!,$input: UpdateMsgTemplateInput!){\n  updateMsgTemplate(id:$id,input: $input){\n    id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,userID,\n    receiverType,format,subject,from,to,cc,bcc,body,tpl,attachments\n  }\n}"): (typeof documents)["mutation updateMsgTemplate($id:ID!,$input: UpdateMsgTemplateInput!){\n  updateMsgTemplate(id:$id,input: $input){\n    id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,userID,\n    receiverType,format,subject,from,to,cc,bcc,body,tpl,attachments\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation delMsgTemplate($id:ID!){\n  deleteMsgTemplate(id:$id)\n}"): (typeof documents)["mutation delMsgTemplate($id:ID!){\n  deleteMsgTemplate(id:$id)\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation enableMsgTemplate($id:ID!){\n  enableMsgTemplate(id:$id){\n    id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,userID,\n    receiverType,format,subject,from,to,cc,bcc,body,tpl,attachments\n  }\n}"): (typeof documents)["mutation enableMsgTemplate($id:ID!){\n  enableMsgTemplate(id:$id){\n    id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,userID,\n    receiverType,format,subject,from,to,cc,bcc,body,tpl,attachments\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation disableMsgTemplate($id:ID!){\n  disableMsgTemplate(id:$id){\n    id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,userID,\n    receiverType,format,subject,from,to,cc,bcc,body,tpl,attachments\n  }\n}"): (typeof documents)["mutation disableMsgTemplate($id:ID!){\n  disableMsgTemplate(id:$id){\n    id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,userID,\n    receiverType,format,subject,from,to,cc,bcc,body,tpl,attachments\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation testSendEmailTpl($annotations: MapString, $email: String!, $labels: MapString, $tplID: ID!){\n  testSendEmailTpl(annotations: $annotations, email: $email, labels:$labels, tplID: $tplID)\n}"): (typeof documents)["mutation testSendEmailTpl($annotations: MapString, $email: String!, $labels: MapString, $tplID: ID!){\n  testSendEmailTpl(annotations: $annotations, email: $email, labels:$labels, tplID: $tplID)\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation testSendMessageTpl($annotations: MapString, $userID: ID!, $labels: MapString, $tplID: ID!){\n  testSendMessageTpl(annotations: $annotations, userID: $userID, labels:$labels, tplID: $tplID)\n}"): (typeof documents)["mutation testSendMessageTpl($annotations: MapString, $userID: ID!, $labels: MapString, $tplID: ID!){\n  testSendMessageTpl(annotations: $annotations, userID: $userID, labels:$labels, tplID: $tplID)\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation refreshTemplateParams{\n  refreshTemplateParams\n}"): (typeof documents)["mutation refreshTemplateParams{\n  refreshTemplateParams\n}"];
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
export function gql(source: "mutation createMsgType($input: CreateMsgTypeInput!){\n  createMsgType(input: $input){\n    id,name,comments,appID,status,category,canSubs,canCustom,createdAt\n  }\n}"): (typeof documents)["mutation createMsgType($input: CreateMsgTypeInput!){\n  createMsgType(input: $input){\n    id,name,comments,appID,status,category,canSubs,canCustom,createdAt\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation updateMsgType($id:ID!,$input: UpdateMsgTypeInput!){\n  updateMsgType(id:$id,input: $input){\n    id,name,comments,appID,status,category,canSubs,canCustom,createdAt\n  }\n}"): (typeof documents)["mutation updateMsgType($id:ID!,$input: UpdateMsgTypeInput!){\n  updateMsgType(id:$id,input: $input){\n    id,name,comments,appID,status,category,canSubs,canCustom,createdAt\n  }\n}"];
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
export function gql(source: "query msgTypeListAndSub($first: Int,$orderBy:MsgTypeOrder,$where:MsgTypeWhereInput){\n  msgTypes(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,comments,appID,status,category,canSubs,canCustom,createdAt\n        subscriberUsers{\n          id,tenantID,msgTypeID,userID\n        },\n        subscriberRoles{\n          id,tenantID,msgTypeID,orgRoleID\n        },\n        excludeSubscriberUsers{\n          id,tenantID,msgTypeID,userID\n        }\n      }\n    }\n  }\n}"): (typeof documents)["query msgTypeListAndSub($first: Int,$orderBy:MsgTypeOrder,$where:MsgTypeWhereInput){\n  msgTypes(first:$first,orderBy: $orderBy,where: $where){\n    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }\n    edges{\n      cursor,node{\n        id,name,comments,appID,status,category,canSubs,canCustom,createdAt\n        subscriberUsers{\n          id,tenantID,msgTypeID,userID\n        },\n        subscriberRoles{\n          id,tenantID,msgTypeID,orgRoleID\n        },\n        excludeSubscriberUsers{\n          id,tenantID,msgTypeID,userID\n        }\n      }\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "query msgTypeAndSubInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgType{\n      id,name,comments,appID,status,category,canSubs,canCustom,createdAt\n      subscriberUsers{\n        id,tenantID,msgTypeID,userID\n      },\n      subscriberRoles{\n        id,tenantID,msgTypeID,orgRoleID\n      },\n      excludeSubscriberUsers{\n        id,tenantID,msgTypeID,userID\n      }\n    }\n  }\n}"): (typeof documents)["query msgTypeAndSubInfo($gid:GID!){\n  node(id: $gid){\n    id\n    ... on MsgType{\n      id,name,comments,appID,status,category,canSubs,canCustom,createdAt\n      subscriberUsers{\n        id,tenantID,msgTypeID,userID\n      },\n      subscriberRoles{\n        id,tenantID,msgTypeID,orgRoleID\n      },\n      excludeSubscriberUsers{\n        id,tenantID,msgTypeID,userID\n      }\n    }\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation createMsgSubscriber($inputs: [CreateMsgSubscriberInput!]!){\n  createMsgSubscriber(inputs: $inputs){\n    id,tenantID,msgTypeID,userID\n  }\n}"): (typeof documents)["mutation createMsgSubscriber($inputs: [CreateMsgSubscriberInput!]!){\n  createMsgSubscriber(inputs: $inputs){\n    id,tenantID,msgTypeID,userID\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "mutation deleteMsgSubscriber($ids: [ID!]!){\n  deleteMsgSubscriber(ids: $ids)\n}"): (typeof documents)["mutation deleteMsgSubscriber($ids: [ID!]!){\n  deleteMsgSubscriber(ids: $ids)\n}"];

export function gql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;