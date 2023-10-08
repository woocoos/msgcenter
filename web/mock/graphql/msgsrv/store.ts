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
 * 普通列表添加
 * @param store
 * @param ref
 * @param field
 * @param addData
 */
export const addList = (store: IMockStore, ref: Ref, field: string, addData: Ref) => {
  const refs = store.get(ref, field) as Ref[]
  refs.push(addData);
  store.set(ref, field, refs);
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
 * 普通列表移除
 * @param store
 * @param ref
 * @param field
 * @param key
 */
export const delList = (store: IMockStore, ref: Ref, field: string, key: string) => {
  const refs = store.get(ref, field) as Ref[]
  const updateRefs = refs.filter(itemRef => itemRef.$ref.key != key)
  store.set(ref, field, updateRefs);
}

/**
 * 获取完整的对象
 * @param store
 * @param ref
 */
export const getObject = (store: IMockStore, ref: Ref) => {
  const data = store['store'][ref.$ref.typeName][ref.$ref.key],
    keys = Object.keys(data);
  if (keys.length) {
    const result = {};
    keys.forEach(key => {
      const keyData: (Ref | number | string | boolean)[] | Ref | number | boolean | string = data[key];
      if (keyData) {
        if (Array.isArray(keyData)) {
          result[key] = keyData.map(item => {
            if (item) {
              if (typeof item === 'object') {
                return getObject(store, store.get(item.$ref.typeName, item.$ref.key) as Ref)
              } else {
                return item;
              }
            } else {
              return null;
            }
          })
        } else if (typeof keyData === 'object') {
          result[key] = getObject(store, store.get(keyData.$ref.typeName, keyData.$ref.key) as Ref)
        } else {
          result[key] = keyData
        }
      } else {
        result[key] = null;
      }
    })
    return result;
  }
  return null
}




/**
 * store内的基础数据
 */
export const initStoreData = (store: IMockStore) => {
  // -------------root------------------------
  store.set('Query', 'ROOT', 'msgChannels', listTemp([
    store.get('MsgChannel', 1),
  ]))
  store.set('Query', 'ROOT', 'msgEvents', listTemp([
    store.get('MsgEvent', 1),
  ]))
  store.set('Query', 'ROOT', 'msgTemplates', listTemp([
    store.get('MsgTemplate', 1),
    store.get('MsgTemplate', 2),
    store.get('MsgTemplate', 3),
  ]))
  store.set('Query', 'ROOT', 'msgTypes', listTemp([
    store.get('MsgType', 1),
    store.get('MsgType', 2),
    store.get('MsgType', 3),
  ]))
  store.set('Query', 'ROOT', 'silences', listTemp([
  ]))
  store.set('Query', 'ROOT', 'msgAlerts', listTemp([
    store.get('MsgAlert', 1),
    store.get('MsgAlert', 2),
    store.get('MsgAlert', 3),
  ]))
  store.set('Query', 'ROOT', 'msgInternals', listTemp([
    store.get('MsgInternal', 1),
    store.get('MsgInternal', 2),
    store.get('MsgInternal', 3),
    store.get('MsgInternal', 4),
    store.get('MsgInternal', 5),
    store.get('MsgInternal', 6),
    store.get('MsgInternal', 7),
    store.get('MsgInternal', 8),
    store.get('MsgInternal', 9),
    store.get('MsgInternal', 10),
  ]))
  store.set('Query', 'ROOT', 'userMsgInternalTos', listTemp([
    store.get('MsgInternalTo', 1),
    store.get('MsgInternalTo', 2),
    store.get('MsgInternalTo', 3),
    store.get('MsgInternalTo', 4),
    store.get('MsgInternalTo', 5),
    store.get('MsgInternalTo', 6),
    store.get('MsgInternalTo', 7),
    store.get('MsgInternalTo', 8),
    store.get('MsgInternalTo', 9),
    store.get('MsgInternalTo', 10),
  ]))


  // -------------root-end------------------------

  // MsgChannel
  store.set('MsgChannel', 1, {
    id: 1, name: 'msgChannel1'
  })

  // MsgType
  store.set('MsgType', 1, {
    id: 1, name: 'MsgType1', appID: 1, category: 'malfunction',
    subscriberUsers: [
      store.get('MsgSubscriber', 1)
    ],
    subscriberRoles: [
      store.get('MsgSubscriber', 3)
    ],
    excludeSubscriberUsers: [
      store.get('MsgSubscriber', 2)
    ],
  })
  store.set('MsgType', 2, {
    id: 2, name: 'MsgType2', appID: 1, category: 'malfunction'
  })
  store.set('MsgType', 3, {
    id: 3, name: 'MsgType2', appID: 1, category: 'business'
  })

  // MsgEvent
  store.set('MsgEvent', 1, {
    id: 1, name: 'MsgEvent1', modes: 'email,webhook', msgType: store.get('MsgType', 1), msgTypeID: 1,
  })

  // msgTemplate
  store.set('MsgTemplate', 1, {
    id: 1, name: 'msgTemplate1', msgEventID: '1', receiverType: 'email', tenantID: '1', subject: '邮件', format: 'txt', body: "{{.name}}{{.text}}", status: 'active',
    tpl: "/xxx/test.tmpl", tplFileID: "tmpl", attachments: ['/xxx/test.png', '/xxx/test.jpg'], attachmentsFileIds: ['png', 'jpg']
  })
  store.set('MsgTemplate', 2, {
    id: 2, name: 'msgTemplate2', msgEventID: '1', receiverType: 'message', tenantID: '1', subject: '邮件', format: 'txt', body: "{{.name}}{{.text}}", status: 'active',
    tpl: "/xxx/test.tmpl", tplFileID: "tmpl", attachments: [], attachmentsFileIds: []
  })
  store.set('MsgTemplate', 3, {
    id: 3, name: 'msgTemplate3', msgEventID: '1', receiverType: 'email', tenantID: '1', subject: '邮件', format: 'html', body: "<div>{{.name}}{{.text}}<div>", status: 'active',
    tpl: "/xxx/test.tmpl", tplFileID: "tmpl", attachments: ['/xxx/test.png', '/xxx/test.jpg'], attachmentsFileIds: ['png', 'jpg']
  })

  // MsgSubscriber
  store.set('MsgSubscriber', 1, {
    id: 1, tenantID: 1, msgTypeID: 1, userID: 1, exclude: false,
  })
  store.set('MsgSubscriber', 2, {
    id: 2, tenantID: 1, msgTypeID: 1, userID: 2, exclude: true,
  })
  store.set('MsgSubscriber', 3, {
    id: 3, tenantID: 1, msgTypeID: 1, orgRoleID: 1, exclude: false,
  })

  // MsgAlert
  store.set('MsgAlert', 1, {
    id: 1, tenantID: 1, state: "none", nlog: listTemp([
      store.get('Nlog', 1),
      store.get('Nlog', 2),
    ])
  })
  store.set('MsgAlert', 2, {
    id: 2, tenantID: 1, state: "firing"
  })
  store.set('MsgAlert', 3, {
    id: 3, tenantID: 1, state: "resolved"
  })

  // Nlog
  store.set('Nlog', 1, {
    id: 1, tenantID: 1, groupKey: "groupKey", receiver: "email"
  })
  store.set('Nlog', 2, {
    id: 2, tenantID: 1, groupKey: "groupKey2", receiver: "email"
  })
  store.set('Nlog', 3, {
    id: 3, tenantID: 1, groupKey: "groupKey3", receiver: "email"
  })

  // MsgInternal
  for (let i = 1; i <= 10; i++) {
    store.set('MsgInternal', i, {
      id: i, tenantID: 1, subject: `主题${i}`, body: `内容${i}`, format: 'text', category: i % 2 ? "business" : "malfunction"
    })
  }

  // MsgInternalTo
  for (let i = 1; i <= 10; i++) {
    store.set('MsgInternalTo', i, {
      id: i, tenantID: 1, userID: 1, msgInternal: store.get('MsgInternal', i), msgInternalID: i
    })
  }
}
