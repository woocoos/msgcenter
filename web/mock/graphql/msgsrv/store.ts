import { IMockStore, Ref } from "@graphql-tools/mock"

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
 * 添加时列表一起添加
 * @param store
 * @param ref
 * @param addData
 * @returns
 */
export const addListTemp = (store: IMockStore, ref: Ref, addData: Ref) => {
  const typeNameEdge = `${addData.$ref.typeName}Edge`,
    edgeKey = `${Math.round(Math.random() * 1000000)}-${addData.$ref.key}`

  store.set(typeNameEdge, edgeKey, 'node', addData)

  const edgesRef = store.get(ref, 'edges') as Ref[]
  edgesRef.push(
    store.get(typeNameEdge, edgeKey) as Ref
  )

  store.set(ref, 'edges', edgesRef)
  store.set(ref, 'totalCount', edgesRef.length)

  return addData;
}

/**
 * 移除时列表一起移除
 * @param store
 * @param ref
 * @param key
 */
export const delListTemp = (store: IMockStore, ref: Ref, key: string) => {
  const edgesRef = store.get(ref, 'edges') as Ref[]
  const updateEdgesRef = edgesRef.filter(itemRef => itemRef.$ref.key.indexOf(key) === -1)
  store.set(ref, 'edges', updateEdgesRef)
  store.set(ref, 'totalCount', updateEdgesRef.length)
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
    store.get('MsgType', 1),
  ]))

  // -------------root-end------------------------

  // MsgChannel
  store.set('MsgChannel', 1, {
    id: 1, name: 'msgChannel1'
  })


  // MsgType
  store.set('MsgType', 1, {
    id: 1, name: 'MsgType1', appID: 1
  })

}
