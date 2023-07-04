import { gql } from "@/__generated__/msgsrv";
import { CreateMsgTypeInput, MsgTypeOrder, MsgTypeWhereInput, UpdateMsgTypeInput } from "@/__generated__/msgsrv/graphql";
import { mutationRequest, pagingRequest, queryRequest } from "..";
import { gid } from "@/util";

export const EnumMsgTypeStatus = {
  active: { text: '活跃', status: 'success' },
  inactive: { text: '失活', status: 'default' },
  processing: { text: '处理中', status: 'warning' },
};

const queryMsgTypeList = gql(/* GraphQL */`query msgTypeList($first: Int,$orderBy:MsgTypeOrder,$where:MsgTypeWhereInput){
  msgTypes(first:$first,orderBy: $orderBy,where: $where){
    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }
    edges{
      cursor,node{
        id,name,comments,appID,status,category,status,canSubs,canCustom,createdAt
      }
    }
  }
}`);


const queryMsgTypeInfo = gql(/* GraphQL */`query msgTypeInfo($gid:GID!){
  node(id: $gid){
    id
    ... on MsgType{
      id,name,comments,appID,status,category,status,canSubs,canCustom,createdAt
    }
  }
}`);

const mutationCreateMsgType = gql(/* GraphQL */`mutation createMsgType($input: CreateMsgTypeInput!){
  createMsgType(input: $input){id}
}`);

const mutationUpdateMsgType = gql(/* GraphQL */`mutation updateMsgType($id:ID!,$input: UpdateMsgTypeInput!){
  updateMsgType(id:$id,input: $input){id}
}`);

const mutationDelMsgType = gql(/* GraphQL */`mutation delMsgType($id:ID!){
  deleteMsgType(id:$id)
}`);

/**
 * 消息类型列表
 * @param gather
 * @returns
 */
export async function getMsgTypeList(
  gather: {
    current?: number;
    pageSize?: number;
    where?: MsgTypeWhereInput;
    orderBy?: MsgTypeOrder;
  }) {
  const result = await pagingRequest(
    queryMsgTypeList, {
    first: gather.pageSize || 20,
    where: gather.where,
    orderBy: gather.orderBy,
  }, gather.current || 1);

  if (result.data?.msgTypes) {
    return result.data.msgTypes;
  }
  return null;
}

/**
 * 消息类型详情
 * @param msgTypeId
 * @returns
 */
export async function getMsgTypeInfo(msgTypeId: string) {
  const result = await queryRequest(queryMsgTypeInfo, {
    gid: gid('msg_type', msgTypeId)
  })
  if (result.data?.node?.__typename === 'MsgType') {
    return result.data.node
  }
  return null
}

/**
 * 创建消息类型
 * @param input
 * @returns
 */
export async function createMsgType(input: CreateMsgTypeInput) {
  const result = await mutationRequest(mutationCreateMsgType, {
    input
  })
  if (result.data?.createMsgType.id) {
    return result.data.createMsgType
  }
  return null
}

/**
 * 更新消息类型
 * @param input
 * @returns
 */
export async function updateMsgType(msgTypeId: string, input: UpdateMsgTypeInput) {
  const result = await mutationRequest(mutationUpdateMsgType, {
    id: msgTypeId,
    input,
  })
  if (result.data?.updateMsgType.id) {
    return result.data.updateMsgType
  }
  return null
}

/**
 * 删除消息类型
 * @param input
 * @returns
 */
export async function delMsgType(msgTypeId: string) {
  const result = await mutationRequest(mutationDelMsgType, {
    id: msgTypeId,
  })
  if (result.data?.deleteMsgType) {
    return result.data.deleteMsgType
  }
  return null
}
