import { gql } from '@/generated/msgsrv';

export const subMessage = gql(/* GraphQL */`subscription subMsg{
  message{
    content,extras,format,sendAt,title,url
  }
}`);


