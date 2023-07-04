import { IMockStore } from "@graphql-tools/mock"

/**
 * 展示列表的模板
 * @param list
 * @returns
 */
export const listTemp = (list: any[]) => {
  return {
    edges: list.map(item => {
      return { node: item }
    }),
    pageInfo: {
      hasNextPage: false,
      hasPreviousPage: false,
    },
    totalCount: list.length,
  }
}


/**
 * store内的基础数据
 */
export const initStoreData = (store: IMockStore) => {
  // -------------root------------------------
  store.set('Query', 'ROOT', 'msgChannels', listTemp([

  ]))
  store.set('Query', 'ROOT', 'msgEvents', listTemp([

  ]))
  store.set('Query', 'ROOT', 'msgTemplates', listTemp([

  ]))
  store.set('Query', 'ROOT', 'msgTypes', listTemp([

  ]))

  // -------------root-end------------------------

  // MsgChannel
  store.set('MsgChannel', 1, {
    id: 1, name: 'msgChannel1'
  })

}
