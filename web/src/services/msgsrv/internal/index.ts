import { gql } from "@/generated/msgsrv";
import { MsgInternalOrder, MsgInternalOrderField, MsgInternalWhereInput, OrderDirection } from "@/generated/msgsrv/graphql";
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

const queryMsgInternalInfo = gql(/* GraphQL */`query msgInternalInfo($gid:GID!){
  node(id: $gid){
    id
    ... on MsgInternal{
      id,tenantID,createdBy,createdAt,subject,body,format,redirect
    }
  }
}`);


const mutationMarkMsgRead = gql(/* GraphQL */`mutation markMsgRead($ids:[ID!]!,$read:Boolean!){
  markMessageReaOrUnRead(ids:$ids,read:$read)
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
  if (result.data?.markMessageReaOrUnRead) {
    return result.data.markMessageReaOrUnRead
  }
  return null
}

/**
 * 删除站内信消息
 * @param msgInternalIds
 * @returns
 */
export async function delMarkMsg(msgInternalIds: string[]) {
  const result = await mutation(mutationDelMarkMsg, {
    ids: msgInternalIds,
  })
  if (result.data?.markMessageDeleted) {
    return result.data.markMessageDeleted
  }
  return null
}
