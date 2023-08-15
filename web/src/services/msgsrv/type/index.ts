import { gql } from "@/generated/msgsrv";
import { CreateMsgSubscriberInput, CreateMsgTypeInput, MsgTypeOrder, MsgTypeWhereInput, UpdateMsgTypeInput } from "@/generated/msgsrv/graphql";
import { gid } from "@knockout-js/api";
import { mutation, paging, query } from '@knockout-js/ice-urql/request'

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
        id,name,comments,appID,status,category,canSubs,canCustom,createdAt
      }
    }
  }
}`);


const queryMsgTypeInfo = gql(/* GraphQL */`query msgTypeInfo($gid:GID!){
  node(id: $gid){
    id
    ... on MsgType{
      id,name,comments,appID,status,category,canSubs,canCustom,createdAt
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

const queryMsgTypeCategory = gql(/* GraphQL */`query msgTypeCategory($keyword:String,$appID:ID){
  msgTypeCategories(keyword: $keyword,appID:$appID)
}`);

const queryMsgTypeListAndSub = gql(/* GraphQL */`query msgTypeListAndSub($first: Int,$orderBy:MsgTypeOrder,$where:MsgTypeWhereInput){
  msgTypes(first:$first,orderBy: $orderBy,where: $where){
    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }
    edges{
      cursor,node{
        id,name,category,
        subscriberUsers{
          id,tenantID,msgTypeID,userID
        },
        subscriberRoles{
          id,tenantID,msgTypeID,orgRoleID
        },
        excludeSubscriberUsers{
          id,tenantID,msgTypeID,userID
        }
      }
    }
  }
}`);

const queryMsgTypeAndSubInfo = gql(/* GraphQL */`query msgTypeAndSubInfo($gid:GID!){
  node(id: $gid){
    id
    ... on MsgType{
      id,name,appID,category,
      subscriberUsers{
        id,tenantID,msgTypeID,userID
      },
      subscriberRoles{
        id,tenantID,msgTypeID,orgRoleID
      },
      excludeSubscriberUsers{
        id,tenantID,msgTypeID,userID
      }
    }
  }
}`);

const mutationCreateSub = gql(/* GraphQL */`mutation createMsgSubscriber($inputs: [CreateMsgSubscriberInput!]!){
  createMsgSubscriber(inputs: $inputs){id}
}`);

const mutationDelSub = gql(/* GraphQL */`mutation deleteMsgSubscriber($ids: [ID!]!){
  deleteMsgSubscriber(ids: $ids)
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
  const result = await paging(
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
  const result = await query(queryMsgTypeInfo, {
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
  const result = await mutation(mutationCreateMsgType, {
    input
  })
  if (result.data?.createMsgType.id) {
    return result.data.createMsgType
  }
  return null
}

/**
 * 更新消息类型
 * @param msgTypeId
 * @param input
 * @returns
 */
export async function updateMsgType(msgTypeId: string, input: UpdateMsgTypeInput) {
  const result = await mutation(mutationUpdateMsgType, {
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
 * @param msgTypeId
 * @returns
 */
export async function delMsgType(msgTypeId: string) {
  const result = await mutation(mutationDelMsgType, {
    id: msgTypeId,
  })
  if (result.data?.deleteMsgType) {
    return result.data.deleteMsgType
  }
  return null
}



/**
 * 获取消息类型分类列表
 * @param keyword
 * @param appID
 * @returns
 */
export async function getMsgTypeCategoryList(keyword?: string, appId?: string) {
  const result = await query(queryMsgTypeCategory, {
    keyword,
    appID: appId,
  })
  if (result.data?.msgTypeCategories) {
    return result.data.msgTypeCategories
  }
  return null
}


/**
 * 消息类型列表包含订阅信息
 * @param gather
 * @returns
 */
export async function getMsgTypeListAndSub(
  gather: {
    current?: number;
    pageSize?: number;
    where?: MsgTypeWhereInput;
    orderBy?: MsgTypeOrder;
  }) {
  const result = await paging(
    queryMsgTypeListAndSub, {
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
export async function getMsgTypeAndSubInfo(msgTypeId: string) {
  const result = await query(queryMsgTypeAndSubInfo, {
    gid: gid('msg_type', msgTypeId)
  })
  if (result.data?.node?.__typename === 'MsgType') {
    return result.data.node
  }
  return null
}

/**
 * 创建订阅
 * @param inputs
 * @returns
 */
export async function createSub(inputs: CreateMsgSubscriberInput[]) {
  const result = await mutation(mutationCreateSub, {
    inputs,
  })
  if (result.data?.createMsgSubscriber) {
    return result.data.createMsgSubscriber
  }
  return null
}

/**
 * 移除订阅
 * @param ids
 * @returns
 */
export async function delSub(ids: string[]) {
  const result = await mutation(mutationDelSub, {
    ids,
  })
  if (result.data?.deleteMsgSubscriber) {
    return result.data.deleteMsgSubscriber
  }
  return null
}
