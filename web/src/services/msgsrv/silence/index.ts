import { gql } from "@/generated/msgsrv";
import { CreateSilenceInput, OrderDirection, SilenceOrder, SilenceOrderField, SilenceWhereInput, UpdateSilenceInput } from "@/generated/msgsrv/graphql";
import { gid } from "@knockout-js/api";
import { mutation, paging, query } from '@knockout-js/ice-urql/request'

export const EnumSilenceStatus = {
  active: { text: 'active', status: 'success' },
  expired: { text: 'expired', status: 'default' },
  pending: { text: 'pending', status: 'warning' },
};

export const EnumSilenceMatchType = {
  MatchEqual: { text: '=', },
  MatchNotEqual: { text: '!=', },
  MatchRegexp: { text: '=~', },
  MatchNotRegexp: { text: '!~', },
};

const querySilenceList = gql(/* GraphQL */`query silenceList($first: Int,$orderBy:SilenceOrder,$where:SilenceWhereInput){
  silences(first:$first,orderBy: $orderBy,where: $where){
    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }
    edges{
      cursor,node{
        id,tenantID,comments,createdAt,startsAt,endsAt,state,
        matchers{
          type,name,value
        }

      }
    }
  }
}`);

const querySilenceInfo = gql(/* GraphQL */`query SilenceInfo($gid:GID!){
  node(id: $gid){
    id
    ... on Silence{
      id,tenantID,startsAt,endsAt,comments,state
      matchers{
        type,name,value
      }
    }
  }
}`);


const mutationCreateSilence = gql(/* GraphQL */`mutation createSilence($input: CreateSilenceInput!){
  createSilence(input: $input){
    id,tenantID,comments,createdAt,startsAt,endsAt,state,
    matchers{
      type,name,value
    }
  }
}`);

const mutationUpdateSilence = gql(/* GraphQL */`mutation updateSilence($id:ID!,$input: UpdateSilenceInput!){
  updateSilence(id:$id,input: $input){
    id,tenantID,comments,createdAt,startsAt,endsAt,state,
    matchers{
      type,name,value
    }
  }
}`);

const mutationDelSilence = gql(/* GraphQL */`mutation delSilence($id:ID!){
  deleteSilence(id:$id)
}`);



/**
 * 消息事件列表
 * @param gather
 * @returns
 */
export async function getSilenceList(
  gather: {
    current?: number;
    pageSize?: number;
    where?: SilenceWhereInput;
    orderBy?: SilenceOrder;
  }) {
  const result = await paging(
    querySilenceList, {
    first: gather.pageSize || 20,
    where: gather.where,
    orderBy: gather.orderBy ?? {
      direction: OrderDirection.Desc,
      field: SilenceOrderField.CreatedAt
    },
  }, gather.current || 1);

  if (result.data?.silences) {
    return result.data.silences;
  }
  return null;
}


/**
 * 详情
 * @param silenceId
 * @returns
 */
export async function getSilenceInfo(silenceId: string) {
  const result = await query(querySilenceInfo, {
    gid: gid('msg_silence', silenceId)
  })
  if (result.data?.node?.__typename === "Silence") {
    return result.data.node
  }
  return null
}


/**
 * 创建
 * @param input
 * @returns
 */
export async function createSilence(input: CreateSilenceInput) {
  const result = await mutation(mutationCreateSilence, {
    input
  })
  if (result.data?.createSilence.id) {
    return result.data.createSilence
  }
  return null
}

/**
 * 更新
 * @param silenceId
 * @param input
 * @returns
 */
export async function updateSilence(silenceId: string, input: UpdateSilenceInput) {
  const result = await mutation(mutationUpdateSilence, {
    id: silenceId,
    input,
  })
  if (result.data?.updateSilence.id) {
    return result.data.updateSilence
  }
  return null
}

/**
 * 删除
 * @param silenceId
 * @returns
 */
export async function delSilence(silenceId: string) {
  const result = await mutation(mutationDelSilence, {
    id: silenceId,
  })
  if (result.data?.deleteSilence) {
    return result.data.deleteSilence
  }
  return null
}
