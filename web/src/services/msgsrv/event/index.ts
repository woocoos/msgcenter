import { gql } from "@/__generated__/msgsrv";
import { CreateMsgEventInput, MsgEventOrder, MsgEventWhereInput, UpdateMsgEventInput } from "@/__generated__/msgsrv/graphql";
import { mutationRequest, pagingRequest, queryRequest } from "..";
import { gid } from "@/util";

export const EnumMsgEventStatus = {
  active: { text: '活跃', status: 'success' },
  inactive: { text: '失活', status: 'default' },
  processing: { text: '处理中', status: 'warning' },
};

const queryMsgEventList = gql(/* GraphQL */`query msgEventList($first: Int,$orderBy:MsgEventOrder,$where:MsgEventWhereInput){
  msgEvents(first:$first,orderBy: $orderBy,where: $where){
    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }
    edges{
      cursor,node{
        id,name,comments,status,createdAt,msgTypeID,modes
        msgType{
          id,category,appID,name
        }
      }
    }
  }
}`);


const queryMsgEventInfo = gql(/* GraphQL */`query MsgEventInfo($gid:GID!){
  node(id: $gid){
    id
    ... on MsgEvent{
      id,name,comments,status,createdAt,msgTypeID,modes
      msgType{
        id,category,appID,name
      }
    }
  }
}`);

const mutationCreateMsgEvent = gql(/* GraphQL */`mutation createMsgEvent($input: CreateMsgEventInput!){
  createMsgEvent(input: $input){id}
}`);

const mutationUpdateMsgEvent = gql(/* GraphQL */`mutation updateMsgEvent($id:ID!,$input: UpdateMsgEventInput!){
  updateMsgEvent(id:$id,input: $input){id}
}`);

const mutationDelMsgEvent = gql(/* GraphQL */`mutation delMsgEvent($id:ID!){
  deleteMsgEvent(id:$id)
}`);

const mutationEnableMsgEvent = gql(/* GraphQL */`mutation enableMsgEvent($id:ID!){
  enableMsgEvent(id:$id){id}
}`);

const mutationDisableMsgEvent = gql(/* GraphQL */`mutation disableMsgEvent($id:ID!){
  disableMsgEvent(id:$id){id}
}`);


/**
 * 消息事件列表
 * @param gather
 * @returns
 */
export async function getMsgEventList(
  gather: {
    current?: number;
    pageSize?: number;
    where?: MsgEventWhereInput;
    orderBy?: MsgEventOrder;
  }) {
  const result = await pagingRequest(
    queryMsgEventList, {
    first: gather.pageSize || 20,
    where: gather.where,
    orderBy: gather.orderBy,
  }, gather.current || 1);

  if (result.data?.msgEvents) {
    return result.data.msgEvents;
  }
  return null;
}

/**
 * 消息事件详情
 * @param msgEventId
 * @returns
 */
export async function getMsgEventInfo(msgEventId: string) {
  const result = await queryRequest(queryMsgEventInfo, {
    gid: gid('msg_event', msgEventId)
  })
  if (result.data?.node?.__typename === 'MsgEvent') {
    return result.data.node
  }
  return null
}

/**
 * 创建消息事件
 * @param input
 * @returns
 */
export async function createMsgEvent(input: CreateMsgEventInput) {
  const result = await mutationRequest(mutationCreateMsgEvent, {
    input
  })
  if (result.data?.createMsgEvent.id) {
    return result.data.createMsgEvent
  }
  return null
}

/**
 * 更新消息事件
 * @param msgEventId
 * @param input
 * @returns
 */
export async function updateMsgEvent(msgEventId: string, input: UpdateMsgEventInput) {
  const result = await mutationRequest(mutationUpdateMsgEvent, {
    id: msgEventId,
    input,
  })
  if (result.data?.updateMsgEvent.id) {
    return result.data.updateMsgEvent
  }
  return null
}

/**
 * 删除消息事件
 * @param msgEventId
 * @returns
 */
export async function delMsgEvent(msgEventId: string) {
  const result = await mutationRequest(mutationDelMsgEvent, {
    id: msgEventId,
  })
  if (result.data?.deleteMsgEvent) {
    return result.data.deleteMsgEvent
  }
  return null
}

/**
 * 禁用消息事件
 * @param msgEventId
 * @returns
 */
export async function disableMsgEvent(msgEventId: string) {
  const result = await mutationRequest(mutationDisableMsgEvent, {
    id: msgEventId,
  })
  if (result.data?.disableMsgEvent.id) {
    return result.data.disableMsgEvent
  }
  return null
}

/**
 * 启用消息事件
 * @param msgEventId
 * @returns
 */
export async function enableMsgEvent(msgEventId: string) {
  const result = await mutationRequest(mutationEnableMsgEvent, {
    id: msgEventId,
  })
  if (result.data?.enableMsgEvent.id) {
    return result.data.enableMsgEvent
  }
  return null
}
