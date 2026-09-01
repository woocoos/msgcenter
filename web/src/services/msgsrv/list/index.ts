import { gql } from '@/generated/msgsrv';
import {
  MsgAlertOrder,
  MsgAlertOrderField,
  MsgAlertWhereInput,
  MsgTemplateReceiverType,
  NlogOrder,
  NlogOrderField,
  NlogWhereInput,
  OrderDirection,
} from '@/generated/msgsrv/graphql';
import { gid } from '@knockout-js/api';
import { mutation, paging, query } from '@knockout-js/ice-urql/request';

export const EnumMsgAlertStatus = {
  firing: { text: 'firing', status: 'success' },
  none: { text: 'none', status: 'default' },
  resolved: { text: 'resolved', status: 'processing' },
};

export const EnumNlogReceiverType = {
  email: { text: 'email' },
  message: { text: 'message' },
  webhook: { text: 'webhook' },
  umeng: { text: 'umeng' },
};

const queryMsgAlertList = gql(/* GraphQL */`query msgAlertList($first: Int,$orderBy:MsgAlertOrder,$where:MsgAlertWhereInput){
  msgAlerts(first:$first,orderBy: $orderBy,where: $where){
    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }
    edges{
      cursor,node{
        id,startsAt,endsAt,labels,annotations,state,timeout
      }
    }
  }
}`);

const queryFormatMsgAlertList = gql(/* GraphQL */`query formatMsgAlerts($first: Int,$alertName:String,$userID:String,$receiverType:MsgTemplateReceiverType,$orderBy:MsgAlertOrder,$where:MsgAlertWhereInput){
  formatMsgAlerts(first:$first,alertName: $alertName,userID: $userID,receiverType: $receiverType,orderBy: $orderBy,where: $where){
    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }
    edges{
      cursor,node{
        id,startsAt,endsAt,tenantID,state,msgEventComments,msgChannelComments,msgTemplateTitle,receiverType,users{name,email},state,receiver,hasMultiMsg,
      }
    }
  }
}`);

const queryFormatMsgAlertMoreList = gql(/* GraphQL */`query formatMsgAlertMore($msgAlertID:ID!){
  formatMsgAlertMore(msgAlertID: $msgAlertID){
    id,startsAt,endsAt,tenantID,state,msgEventComments,msgChannelComments,msgTemplateTitle,receiverType,users{name,email},state,receiver,hasMultiMsg,
  }
}`);

const queryRenderMsgAlert = gql(/* GraphQL */`query renderMsgAlert($msgAlertID:ID!,$receiver:String!){
  renderMsgAlert(msgAlertID: $msgAlertID,receiver:$receiver)
}`);

const queryMsgAlertLogList = gql(/* GraphQL */`query msgAlertLogList($gid:GID!,$first: Int,$orderBy:NlogOrder,$where:NlogWhereInput){
   node(id: $gid){
    id
    ... on MsgAlert{
      id,
      nlog(first:$first,orderBy: $orderBy,where: $where){
        totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }
        edges{
          cursor,node{
            id,sendAt,expiresAt,groupKey,receiver,receiverType
          }
        }
      }
    }
  }
}`);


/**
 * 消息列表
 * @param gather
 * @returns
 */
export async function getMsgAlertList(
  gather: {
    current?: number;
    pageSize?: number;
    where?: MsgAlertWhereInput;
    orderBy?: MsgAlertOrder;
  }) {
  const result = await paging(
    queryMsgAlertList, {
    first: gather.pageSize || 20,
    where: gather.where,
    orderBy: gather.orderBy ?? {
      direction: OrderDirection.Desc,
      field: MsgAlertOrderField.CreatedAt,
    },
  }, gather.current || 1);

  if (result.data?.msgAlerts) {
    return result.data.msgAlerts;
  }
  return null;
}

/**
 * 消息查询
 * @param gather
 * @returns
 */
export async function getFormatMsgAlertList(
  gather: {
    current?: number;
    pageSize?: number;
    alertName?: string;
    userID?: string;
    receiverType?: MsgTemplateReceiverType;
    orderBy?: MsgAlertOrder;
    where?: MsgAlertWhereInput;
  }) {
  let variables = {
    first: gather.pageSize || 20,
    orderBy: gather.orderBy ?? {
      direction: OrderDirection.Desc,
      field: MsgAlertOrderField.CreatedAt,
    },
    where: gather.where,
  };
  if (gather?.alertName) {
    variables['alertName'] = gather?.alertName;
  }
  if (gather?.userID) {
    variables['userID'] = gather?.userID;
  }
  if (gather?.receiverType) {
    variables['receiverType'] = gather?.receiverType;
  }
  const result = await paging(
    queryFormatMsgAlertList, variables, gather.current || 1);

  if (result.data?.formatMsgAlerts) {
    return result.data.formatMsgAlerts;
  }
  return null;
}

/**
 * 消息查询
 * @param msgAlertID
 * @returns
 */
export async function getFormatMsgAlertMore(msgAlertID: string | null) {
  const result = await query(queryFormatMsgAlertMoreList, {
      msgAlertID,
    });
  if (result.data?.formatMsgAlertMore) {
    return result.data.formatMsgAlertMore;
  }
  return null;
}

/**
 * 消息渲染结果
 * @param msgAlertID
 * @param receiver
 * @returns
 */
export async function getRenderMsgAlert(msgAlertID: string, receiver: string | undefined) {
  const result = await query(queryRenderMsgAlert, {
    msgAlertID, receiver,
  });
  if (result.data?.renderMsgAlert) {
    return result.data.renderMsgAlert;
  }
  return null;
}

/**
 * 消息日志列表
 * @param msgAlertId
 * @param gather
 * @returns
 */
export async function getMsgAlertLogList(
  msgAlertId: string,
  gather: {
    current?: number;
    pageSize?: number;
    where?: NlogWhereInput;
    orderBy?: NlogOrder;
  }) {
  const result = await paging(
    queryMsgAlertLogList, {
    gid: gid('MsgAlert', msgAlertId),
    first: gather.pageSize || 20,
    where: gather.where,
    orderBy: gather.orderBy ?? {
      direction: OrderDirection.Desc,
      field: NlogOrderField.CreatedAt,
    },
  }, gather.current || 1);

  if (result.data?.node?.__typename === 'MsgAlert') {
    return result.data.node.nlog;
  }
  return null;
}
