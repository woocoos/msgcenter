import { AnyVariables, Client, DocumentInput, OperationContext, cacheExchange, fetchExchange } from 'urql';
import { urglMapExchange } from '../graphql';

const baseURL = '/api-msgsrv'
const url = `${baseURL}/graphql/query`;
const client = new Client({
  url,
  requestPolicy: 'cache-and-network',
  exchanges: [
    urglMapExchange,
    cacheExchange,
    fetchExchange,
  ],
});

/**
 * 分页请求
 * @param query
 * @param variables
 * @param current
 * @param headers
 * @returns
 */
export const pagingRequest = async <Data = any, Variables extends AnyVariables = AnyVariables>(
  query: DocumentInput<Data, Variables>,
  variables: Variables,
  current: number,
  headers?: Record<string, any>,
) => {
  const context: Partial<OperationContext> = {}
  context.url = `${url}?p=${current}`;
  if (headers) {
    context.fetchOptions = { headers }
  }
  return await client.query(query, variables, context).toPromise()
}

/**
 * query请求
 * @param query
 * @param variables
 * @param headers
 * @returns
 */
export const queryRequest = async <Data = any, Variables extends AnyVariables = AnyVariables>(
  query: DocumentInput<Data, Variables>,
  variables: Variables,
  headers?: Record<string, any>,
) => {
  const context: Partial<OperationContext> = {}
  if (headers) {
    context.fetchOptions = { headers }
  }
  return await client.query(query, variables, context).toPromise()
}

/**
 * mutation请求
 * @param query
 * @param variables
 * @param headers
 * @returns
 */
export const mutationRequest = async <Data = any, Variables extends AnyVariables = AnyVariables>(
  query: DocumentInput<Data, Variables>,
  variables: Variables,
  headers?: Record<string, any>,
) => {
  const context: Partial<OperationContext> = {}
  if (headers) {
    context.fetchOptions = { headers }
  }
  return await client.mutation(query, variables, context).toPromise()
}
