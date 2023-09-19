import { gql } from "@/generated/msgsrv";
import { MsgInternalOrder, MsgInternalOrderField, MsgInternalToOrder, MsgInternalToOrderField, MsgInternalToWhereInput, MsgInternalWhereInput, OrderDirection } from "@/generated/msgsrv/graphql";
import { gid } from "@knockout-js/api";
import { mutation, paging, query } from '@knockout-js/ice-urql/request'

const queryMsgInternalList = gql(/* GraphQL */`query msgInternalList($first: Int,$orderBy: MsgInternalOrder,$where:MsgInternalWhereInput){
  msgInternals(first:$first,orderBy: $orderBy,where: $where){
    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }
    edges{
      cursor,node{
        id,tenantID,createdBy,createdAt,subject,body,format,redirect
      }
    }
  }
}`);

const queryUserMsgInternalList = gql(/* GraphQL */`query userMsgInternalList($first: Int,$orderBy: MsgInternalToOrder,$where:MsgInternalToWhereInput){
  userMessages(first:$first,orderBy: $orderBy,where: $where){
    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }
    edges{
      cursor,node{
        id,msgInternalID,createdAt,deleteAt,readAt,userID
        msgInternal{
          id,tenantID,createdBy,createdAt,subject,body,format,redirect,category
        }
      }
    }
  }
}`);

const queryMsgInternalInfo = gql(/* GraphQL */`query msgInternalInfo($gid:GID!){
  node(id: $gid){
    id
    ... on MsgInternal{
      id,tenantID,createdBy,createdAt,subject,body,format,redirect
    }
  }
}`);

const queryUserMsgCategory = gql(/* GraphQL */`query userMsgCategory{
  userSubMsgCategory
}`);

const queryUserMsgCategoryNum = gql(/* GraphQL */`query userMsgCategoryNum($categories:[String!]!){
  userUnreadMessagesFromMsgCategory(categories:$categories)
}`);

const queryMsgInternalToInfo = gql(/* GraphQL */`query msgInternalToInfo($gid:GID!){
  node(id: $gid){
    id
    ... on MsgInternalTo{
      id,msgInternalID,createdAt,deleteAt,readAt,userID
      msgInternal{
        id,tenantID,createdBy,createdAt,subject,body,format,redirect,category
      }
    }
  }
}`);

const mutationMarkMsgRead = gql(/* GraphQL */`mutation markMsgRead($ids:[ID!]!,$read:Boolean!){
  markMessageReadOrUnRead(ids:$ids,read:$read)
}`);

const mutationDelMarkMsg = gql(/* GraphQL */`mutation delMarkMsg($ids:[ID!]!){
  markMessageDeleted(ids:$ids)
}`);

/**
 * 列表
 * @param gather
 * @returns
 */
export async function getMsgInternalList(
  gather: {
    current?: number;
    pageSize?: number;
    where?: MsgInternalWhereInput;
    orderBy?: MsgInternalOrder
    ;
  }) {
  const result = await paging(
    queryMsgInternalList, {
    first: gather.pageSize || 20,
    where: gather.where,
    orderBy: gather.orderBy ?? {
      direction: OrderDirection.Desc,
      field: MsgInternalOrderField.CreatedAt
    },
  }, gather.current || 1);

  if (result.data?.msgInternals) {
    return result.data.msgInternals;
  }
  return null;
}

/**
 * 详情
 * @param msgInternalId
 * @returns
 */
export async function getMsgInternalInfo(msgInternalId: string) {
  const result = await query(queryMsgInternalInfo, {
    gid: gid('msg_internal', msgInternalId)
  })
  if (result.data?.node?.__typename === "MsgInternal") {
    return result.data.node
  }
  return null
}

/**
 * 批量设置已读未读
 * @param msgInternalId
 * @returns
 */
export async function markMsgRead(msgInternalIds: string[], read: boolean) {
  const result = await mutation(mutationMarkMsgRead, {
    ids: msgInternalIds,
    read
  })
  if (result.data?.markMessageReadOrUnRead) {
    return result.data.markMessageReadOrUnRead
  }
  return null
}

/**
 * 删除站内信消息
 * @param msgInternalToIds
 * @returns
 */
export async function delMarkMsg(msgInternalToIds: string[]) {
  const result = await mutation(mutationDelMarkMsg, {
    ids: msgInternalToIds,
  })
  if (result.data?.markMessageDeleted) {
    return result.data.markMessageDeleted
  }
  return null
}

/**
 * 列表
 * @param gather
 * @returns
 */
export async function getUserMsgInternalList(
  gather: {
    current?: number;
    pageSize?: number;
    where?: MsgInternalToWhereInput;
    orderBy?: MsgInternalToOrder;
  }) {
  const result = await paging(
    queryUserMsgInternalList, {
    first: gather.pageSize || 20,
    where: gather.where,
    orderBy: gather.orderBy ?? {
      direction: OrderDirection.Desc,
      field: MsgInternalToOrderField.CreatedAt
    },
  }, gather.current || 1);

  if (result.data?.userMessages) {
    return result.data.userMessages;
  }
  return null;
}

/**
 * 详情
 * @param msgInternalToId
 * @returns
 */
export async function getMsgInternalToInfo(msgInternalToId: string) {
  const result = await query(queryMsgInternalToInfo, {
    gid: gid('msg_internal_to', msgInternalToId)
  })
  if (result.data?.node?.__typename === "MsgInternal") {
    return result.data.node
  }
  return null
}

/**
 * 当前用户的消息分类
 * @returns
 */
export async function getUserMsgCategory() {
  const result = await query(queryUserMsgCategory, {})
  if (result.data?.userSubMsgCategory) {
    return result.data.userSubMsgCategory
  }
  return null
}

/**
 * 当前用户的消息分类未读数量
 * @returns
 */
export async function getUserMsgCategoryNum(categories: string[]) {
  const result = await query(queryUserMsgCategoryNum, {
    categories,
  })
  if (result.data?.userUnreadMessagesFromMsgCategory) {
    return result.data.userUnreadMessagesFromMsgCategory
  }
  return null
}
