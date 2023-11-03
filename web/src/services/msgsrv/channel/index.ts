import { gql } from "@/generated/msgsrv";
import { mutation, paging, query } from '@knockout-js/ice-urql/request'
import { CreateMsgChannelInput, MsgChannelOrder, MsgChannelOrderField, MsgChannelWhereInput, UpdateMsgChannelInput, OrderDirection } from "@/generated/msgsrv/graphql";
import { gid } from "@knockout-js/api";


export const EnumMsgChannelReceiverType = {
  email: { text: 'email' },
  message: { text: 'message' },
  webhook: { text: 'webhook' },
};

export const EnumMsgChannelStatus = {
  active: { text: '活跃', status: 'success' },
  inactive: { text: '失活', status: 'default' },
  processing: { text: '处理中', status: 'warning' },
};

const queryMsgChannelList = gql(/* GraphQL */`query msgChannelList($first: Int,$orderBy:MsgChannelOrder,$where:MsgChannelWhereInput){
  msgChannels(first:$first,orderBy: $orderBy,where: $where){
    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }
    edges{
      cursor,node{
        id,name,receiverType,tenantID,comments,status,status,createdAt
      }
    }
  }
}`);


const queryMsgChannelInfo = gql(/* GraphQL */`query msgChannelInfo($gid:GID!){
  node(id: $gid){
    id
    ... on MsgChannel{
      id,name,receiverType,tenantID,comments,status,status,createdAt
    }
  }
}`);

const queryMsgChannelReceiverInfo = gql(/* GraphQL */`query msgChannelReceiverInfo($gid:GID!){
  node(id: $gid){
    id
    ... on MsgChannel{
      id,receiver{
        name,
        emailConfigs{
          authIdentity,authPassword,authSecret,authType,authUsername,from,headers,smartHost,to
        },
        messageConfig{
          redirect,subject,to
        }
      }
    }
  }
}`);

const mutationCreateMsgChannel = gql(/* GraphQL */`mutation createMsgChannel($input: CreateMsgChannelInput!){
 createMsgChannel(input:$input){id}
}`);

const mutationUpdateMsgChannel = gql(/* GraphQL */`mutation updateMsgChannel($id:ID!,$input: UpdateMsgChannelInput!){
 updateMsgChannel(id:$id,input:$input){id}
}`);

const mutationDelMsgChannel = gql(/* GraphQL */`mutation delMsgChannel($id:ID!){
 deleteMsgChannel(id:$id)
}`);

const mutationEnableMsgChannel = gql(/* GraphQL */`mutation enableMsgChannel($id:ID!){
 enableMsgChannel(id:$id){id}
}`);

const mutationDisableMsgChannel = gql(/* GraphQL */`mutation disableMsgChannel($id:ID!){
 disableMsgChannel(id:$id){id}
}`);



/**
 * 消息通道列表
 * @param gather
 * @returns
 */
export async function getMsgChannelList(
  gather: {
    current?: number;
    pageSize?: number;
    where?: MsgChannelWhereInput;
    orderBy?: MsgChannelOrder;
  }) {
  const result = await paging(
    queryMsgChannelList, {
    first: gather.pageSize || 20,
    where: gather.where,
    orderBy: gather.orderBy ?? {
      direction: OrderDirection.Desc,
      field: MsgChannelOrderField.CreatedAt
    },
  }, gather.current || 1);

  if (result.data?.msgChannels) {
    return result.data.msgChannels;
  }
  return null;
}

/**
 * 消息通道详情
 * @param msgChannelId
 * @returns
 */
export async function getMsgChannelInfo(msgChannelId: string) {
  const result = await query(queryMsgChannelInfo, {
    gid: gid('msg_channel', msgChannelId)
  })
  if (result.data?.node?.__typename === 'MsgChannel') {
    return result.data.node
  }
  return null
}

/**
 * 消息通道详情
 * @param msgChannelId
 * @returns
 */
export async function getMsgChannelReceiverInfo(msgChannelId: string) {
  const result = await query(queryMsgChannelReceiverInfo, {
    gid: gid('msg_channel', msgChannelId)
  })
  if (result.data?.node?.__typename === 'MsgChannel') {
    return result.data.node
  }
  return null
}

/**
 * 创建消息通道
 * @param input
 * @returns
 */
export async function createMsgChannel(input: CreateMsgChannelInput) {
  const result = await mutation(mutationCreateMsgChannel, {
    input
  })
  if (result.data?.createMsgChannel.id) {
    return result.data.createMsgChannel
  }
  return null
}

/**
 * 更新消息通道
 * @param msgChannelId
 * @param input
 * @returns
 */
export async function updateMsgChannel(msgChannelId: string, input: UpdateMsgChannelInput) {
  const result = await mutation(mutationUpdateMsgChannel, {
    id: msgChannelId,
    input,
  })
  if (result.data?.updateMsgChannel.id) {
    return result.data.updateMsgChannel
  }
  return null
}

/**
 * 禁用消息通道
 * @param msgChannelId
 * @returns
 */
export async function disableMsgChannel(msgChannelId: string) {
  const result = await mutation(mutationDisableMsgChannel, {
    id: msgChannelId,
  })
  if (result.data?.disableMsgChannel.id) {
    return result.data.disableMsgChannel
  }
  return null
}

/**
 * 启用消息通道
 * @param msgChannelId
 * @returns
 */
export async function enableMsgChannel(msgChannelId: string) {
  const result = await mutation(mutationEnableMsgChannel, {
    id: msgChannelId,
  })
  if (result.data?.enableMsgChannel.id) {
    return result.data.enableMsgChannel
  }
  return null
}

/**
 * 启用消息通道
 * @param msgChannelId
 * @returns
 */
export async function delMsgChannel(msgChannelId: string) {
  const result = await mutation(mutationDelMsgChannel, {
    id: msgChannelId,
  })
  if (result.data?.deleteMsgChannel) {
    return result.data.deleteMsgChannel
  }
  return null
}
