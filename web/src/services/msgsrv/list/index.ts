import { gql } from "@/generated/msgsrv";
import { MsgAlertOrder, MsgAlertOrderField, MsgAlertWhereInput, NlogOrder, NlogOrderField, NlogWhereInput, OrderDirection } from "@/generated/msgsrv/graphql";
import { gid } from "@knockout-js/api";
import { mutation, paging, query } from '@knockout-js/ice-urql/request'

export const EnumMsgAlertStatus = {
  firing: { text: 'firing', status: 'success' },
  none: { text: 'none', status: 'default' },
  resolved: { text: 'resolved', status: 'processing' },
};

export const EnumNlogReceiverType = {
  email: { text: 'email' },
  message: { text: 'message' },
  webhook: { text: 'webhook' },
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
      field: MsgAlertOrderField.CreatedAt
    },
  }, gather.current || 1);

  if (result.data?.msgAlerts) {
    return result.data.msgAlerts;
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
    gid: gid('msg_alert', msgAlertId),
    first: gather.pageSize || 20,
    where: gather.where,
    orderBy: gather.orderBy ?? {
      direction: OrderDirection.Desc,
      field: NlogOrderField.CreatedAt
    },
  }, gather.current || 1);

  if (result.data?.node?.__typename === 'MsgAlert') {
    return result.data.node.nlog;
  }
  return null;
}
