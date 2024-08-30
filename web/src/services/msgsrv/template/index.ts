import { gql } from "@/generated/msgsrv";
import { CreateMsgTemplateInput, MsgTemplateOrder, MsgTemplateOrderField, MsgTemplateWhereInput, OrderDirection, UpdateMsgTemplateInput } from "@/generated/msgsrv/graphql";
import { gid } from "@knockout-js/api";
import { mutation, paging, query } from '@knockout-js/ice-urql/request'

export const EnumMsgTemplateStatus = {
  active: { text: '活跃', status: 'success' },
  inactive: { text: '失活', status: 'default' },
  processing: { text: '处理中', status: 'warning' },
};

export const EnumMsgTemplateReceiverType = {
  email: { text: 'email' },
  message: { text: 'message' },
  webhook: { text: 'webhook' },
};

export const EnumMsgTemplateFormat = {
  txt: { text: 'txt' },
  html: { text: 'html' },
};

const queryMsgTemplateList = gql(/* GraphQL */`query msgTemplateList($first: Int,$orderBy:MsgTemplateOrder,$where:MsgTemplateWhereInput){
  msgTemplates(first:$first,orderBy: $orderBy,where: $where){
    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }
    edges{
      cursor,node{
        id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,
        receiverType,format,subject,from,to,cc,bcc,body,tpl,attachments
      }
    }
  }
}`);

const queryMsgTemplateInfo = gql(/* GraphQL */`query MsgTemplateInfo($gid:GID!){
  node(id: $gid){
    id
    ... on MsgTemplate{
      id,name,comments,status,createdAt,msgTypeID,msgEventID,tenantID,
      receiverType,format,subject,from,to,cc,bcc,body,tpl,attachments
    }
  }
}`);

const mutationCreateMsgTemplate = gql(/* GraphQL */`mutation createMsgTemplate($input: CreateMsgTemplateInput!){
  createMsgTemplate(input: $input){id}
}`);

const mutationUpdateMsgTemplate = gql(/* GraphQL */`mutation updateMsgTemplate($id:ID!,$input: UpdateMsgTemplateInput!){
  updateMsgTemplate(id:$id,input: $input){id}
}`);

const mutationDelMsgTemplate = gql(/* GraphQL */`mutation delMsgTemplate($id:ID!){
  deleteMsgTemplate(id:$id)
}`);

const mutationEnableMsgTemplate = gql(/* GraphQL */`mutation enableMsgTemplate($id:ID!){
  enableMsgTemplate(id:$id){id}
}`);

const mutationDisableMsgTemplate = gql(/* GraphQL */`mutation disableMsgTemplate($id:ID!){
  disableMsgTemplate(id:$id){id}
}`);



const mutationTestSendEmailTpl = gql(/* GraphQL */`mutation testSendEmailTpl($annotations: MapString, $email: String!, $labels: MapString, $tplID: ID!){
  testSendEmailTpl(annotations: $annotations, email: $email, labels:$labels, tplID: $tplID)
}`);

const mutationTestSendMessageTpl = gql(/* GraphQL */`mutation testSendMessageTpl($annotations: MapString, $userID: ID!, $labels: MapString, $tplID: ID!){
  testSendMessageTpl(annotations: $annotations, userID: $userID, labels:$labels, tplID: $tplID)
}`);


/**
 * 消息事件列表
 * @param gather
 * @returns
 */
export async function getMsgTemplateList(
  gather: {
    current?: number;
    pageSize?: number;
    where?: MsgTemplateWhereInput;
    orderBy?: MsgTemplateOrder;
  }) {
  const result = await paging(
    queryMsgTemplateList, {
    first: gather.pageSize || 20,
    where: gather.where,
    orderBy: gather.orderBy ?? {
      direction: OrderDirection.Desc,
      field: MsgTemplateOrderField.CreatedAt
    },
  }, gather.current || 1);

  if (result.data?.msgTemplates) {
    return result.data.msgTemplates;
  }
  return null;
}

/**
 * 消息事件详情
 * @param msgTemplateId
 * @returns
 */
export async function getMsgTemplateInfo(msgTemplateId: string) {
  const result = await query(queryMsgTemplateInfo, {
    gid: gid('msg_template', msgTemplateId)
  })
  if (result.data?.node?.__typename === 'MsgTemplate') {
    return result.data.node
  }
  return null
}

/**
 * 创建消息事件
 * @param input
 * @returns
 */
export async function createMsgTemplate(input: CreateMsgTemplateInput) {
  const result = await mutation(mutationCreateMsgTemplate, {
    input
  })
  if (result.data?.createMsgTemplate.id) {
    return result.data.createMsgTemplate
  }
  return null
}

/**
 * 更新消息事件
 * @param msgTemplateId
 * @param input
 * @returns
 */
export async function updateMsgTemplate(msgTemplateId: string, input: UpdateMsgTemplateInput) {
  const result = await mutation(mutationUpdateMsgTemplate, {
    id: msgTemplateId,
    input,
  })
  if (result.data?.updateMsgTemplate.id) {
    return result.data.updateMsgTemplate
  }
  return null
}

/**
 * 删除消息事件
 * @param msgTemplateId
 * @returns
 */
export async function delMsgTemplate(msgTemplateId: string) {
  const result = await mutation(mutationDelMsgTemplate, {
    id: msgTemplateId,
  })
  if (result.data?.deleteMsgTemplate) {
    return result.data.deleteMsgTemplate
  }
  return null
}

/**
 * 禁用消息事件
 * @param msgTemplateId
 * @returns
 */
export async function disableMsgTemplate(msgTemplateId: string) {
  const result = await mutation(mutationDisableMsgTemplate, {
    id: msgTemplateId,
  })
  if (result.data?.disableMsgTemplate.id) {
    return result.data.disableMsgTemplate
  }
  return null
}

/**
 * 启用消息事件
 * @param msgTemplateId
 * @returns
 */
export async function enableMsgTemplate(msgTemplateId: string) {
  const result = await mutation(mutationEnableMsgTemplate, {
    id: msgTemplateId,
  })
  if (result.data?.enableMsgTemplate.id) {
    return result.data.enableMsgTemplate
  }
  return null
}

/**
 * 测试邮件
 * @param tplID
 * @param email
 * @param labels
 * @param annotations
 * @returns
 */
export async function testSendEamil(tplID: string, email: string, labels?: Record<string, string>, annotations?: Record<string, string>) {
  const result = await mutation(mutationTestSendEmailTpl, {
    tplID,
    email,
    labels,
    annotations,
  })
  if (result.data?.testSendEmailTpl) {
    return result.data.testSendEmailTpl
  }
  return null
}

/**
 * 测试站内信发送
 * @param tplID
 * @param userID
 * @param labels
 * @param annotations
 * @returns
 */
export async function testSendMessage(tplID: string, userID: string, labels?: Record<string, string>, annotations?: Record<string, string>) {
  const result = await mutation(mutationTestSendMessageTpl, {
    tplID,
    userID,
    labels,
    annotations,
  })
  if (result.data?.testSendMessageTpl) {
    return result.data.testSendMessageTpl
  }
  return null
}
