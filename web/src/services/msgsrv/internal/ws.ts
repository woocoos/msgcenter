import { gql } from "@/generated/msgsrv";
import { subscription } from '@knockout-js/ice-urql/request'

const querySubMsg = gql(/* GraphQL */`subscription subMsg{
  message{
    action,payload,key,topic,sendAt
  }
}`);

/**
 * 订阅消息
 * @returns
 */
export const subMsg = async () => {
  const result = await subscription(querySubMsg, {});

  if (result.data?.message) {
    return result.data.message;
  }

  return null;
}
