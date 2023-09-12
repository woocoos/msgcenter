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
  store.set('Query', 'ROOT', 'apps', listTemp([
    store.get('App', 1),
    store.get('App', 2),
  ]))
  store.set('Query', 'ROOT', 'users', listTemp([
    store.get('User', 1),
    store.get('User', 2),
    store.get('User', 3),
    store.get('User', 4),
    store.get('User', 5),
    store.get('User', 6),
    store.get('User', 7),
  ]))
  store.set('Query', 'ROOT', 'orgGroups', listTemp([
    store.get('OrgRole', 1),
    store.get('OrgRole', 2),
    store.get('OrgRole', 4),
    store.get('OrgRole', 5),
    store.get('OrgRole', 6),
    store.get('OrgRole', 7),
    store.get('OrgRole', 8),
  ]))
  store.set('Query', 'ROOT', 'orgRoles', listTemp([
    store.get('OrgRole', 3),
  ]))
  store.set('Query', 'ROOT', 'orgRoleUsers', listTemp([
    store.get('User', 1),
  ]))
  store.set('Query', 'ROOT', 'appResources', listTemp([
    store.get('AppRes', 1),
  ]))
  store.set('Query', 'ROOT', 'orgAppResources', listTemp([
    store.get('AppRes', 1),
  ]))
  store.set('Query', 'ROOT', 'orgPolicyReferences', listTemp([
    store.get('Permission', 1),
    store.get('Permission', 2),
  ]))
  // -------------root-end------------------------


  // Org
  store.set('Org', 1, {
    id: 1, name: 'woocoo', code: 'woocoo', domain: 'woocoo', parentID: 0, kind: 'root', ownerID: 1, owner: store.get('User', 1),
    users: listTemp([
      store.get('User', 1),
      store.get('User', 2),
      store.get('User', 3),
      store.get('User', 4),
      store.get('User', 5),
      store.get('User', 6),
      store.get('User', 7),
    ]),
    apps: listTemp([
      store.get('App', 1)
    ]),
    permissions: listTemp([
      store.get("AppPolicy", 1)
    ]),
    policies: listTemp([
      store.get("AppPolicy", 1)
    ]),
  })
  store.set('Org', 2, { id: 2, name: 'org2', code: 'org2', domain: '', parentID: 1, kind: 'org' })
  store.set('Org', 3, { id: 3, name: 'org3', code: 'org3', domain: '', parentID: 1, kind: 'org' })
  store.set('Org', 5, {
    id: 5, name: 'org5', parentID: 0, kind: 'root', ownerID: 1, owner: store.get('User', 1),
    users: listTemp([
      store.get('User', 1)
    ])
  })
  store.set('Org', 4, { id: 4, name: 'org2-1', code: 'org2-1', domain: '', parentID: 2, kind: 'org' })

  // OrgRole
  store.set('OrgRole', 1, { id: 1, name: 'org1group1', kind: 'group', orgID: 1 })
  store.set('OrgRole', 2, { id: 2, name: 'org1group2', kind: 'group', orgID: 1 })
  store.set('OrgRole', 3, { id: 3, name: 'org1role3', kind: 'role', orgID: 1 })
  store.set('OrgRole', 4, { id: 4, name: 'org1group4', kind: 'group', orgID: 1 })
  store.set('OrgRole', 5, { id: 5, name: 'org1group5', kind: 'group', orgID: 1 })
  store.set('OrgRole', 6, { id: 6, name: 'org1group6', kind: 'group', orgID: 1 })
  store.set('OrgRole', 7, { id: 7, name: 'org1group7', kind: 'group', orgID: 1 })
  store.set('OrgRole', 8, { id: 8, name: 'org1group8', kind: 'group', orgID: 1 })

  // OrgPolicy
  store.set('OrgPolicy', 1, { id: 1, name: 'org1Policy1' })

  // User
  store.set('User', 1, {
    id: 1, displayName: 'admin', userType: "account", email: "admin@woocoo.com",
    loginProfile: { mfaEnabled: false }
  })
  store.set('User', 2, {
    id: 2, displayName: 'user2', userType: "account", email: "user2@woocoo.com",
  })
  store.set('User', 3, {
    id: 3, displayName: 'user3', email: "user3@woocoo.com",
  })
  store.set('User', 4, {
    id: 4, displayName: 'user4', email: "user4@woocoo.com",
  })
  store.set('User', 5, {
    id: 5, displayName: 'user5', email: "user5@woocoo.com",
  })
  store.set('User', 6, {
    id: 6, displayName: 'user6', email: "user6@woocoo.com",
  })
  store.set('User', 7, {
    id: 7, displayName: 'user7', email: "user7@woocoo.com",
  })

  // app
  store.set('App', 1, {
    id: 1, name: 'app1', code: 'app1',
    actions: listTemp([
      store.get('AppAction', 1),
      store.get('AppAction', 2),
      store.get('AppAction', 3),
    ]),
    policies: [
      store.get("AppPolicy", 1)
    ],
    menus: listTemp([
      store.get("AppMenu", 1),
      store.get("AppMenu", 2),
      store.get("AppMenu", 3),
      store.get("AppMenu", 4),
      store.get("AppMenu", 5),
      store.get("AppMenu", 6),
      store.get("AppMenu", 7),
      store.get("AppMenu", 8),
      store.get("AppMenu", 9),
      store.get("AppMenu", 10),
      store.get("AppMenu", 11),
      store.get("AppMenu", 12),
      store.get("AppMenu", 13),
      store.get("AppMenu", 14),
    ]),
    roles: [
      store.get("AppRole", 1),
      store.get("AppRole", 2),
      store.get("AppRole", 3),
      store.get("AppRole", 4),
    ],
    resources: listTemp([
      store.get('AppRes', 1),
    ]),
    orgs: listTemp([
      store.get("Org", 1),
      store.get("Org", 5),
    ])
  })
  store.set('App', 2, { id: 2, name: 'app2', code: 'app2' })

  // AppAction
  store.set('AppAction', 1, { id: 1, name: 'app1Action1', method: 'read', appID: 1, app: store.get('App', 1) })
  store.set('AppAction', 2, { id: 2, name: 'app1Action2', method: 'write', appID: 1, app: store.get('App', 1) })
  store.set('AppAction', 3, { id: 3, name: 'app1Action3', method: 'list', appID: 1, app: store.get('App', 1) })

  // AppPolicy
  store.set('AppPolicy', 1, {
    id: 1, name: 'app1Policy1', appID: 1, app: store.get('App', 1), rules: [
      {
        effect: 'allow',
        actions: ['*'],
        resources: ['*'],
        conditions: [],
      }
    ]
  })

  // AppMenu
  store.set('AppMenu', 1, { id: 1, name: '工作台', kind: 'menu', route: '/', parentID: 0, appID: 1, app: store.get('App', 1), })
  store.set('AppMenu', 2, { id: 2, name: '个人中心', kind: 'dir', parentID: 0, appID: 1, app: store.get('App', 1), })
  store.set('AppMenu', 3, { id: 3, name: '基本信息', kind: 'menu', route: '/user/info', parentID: 2, appID: 1, app: store.get('App', 1), })
  store.set('AppMenu', 4, { id: 4, name: '安全设置', kind: 'menu', route: '/user/safety', parentID: 2, appID: 1, app: store.get('App', 1), })
  store.set('AppMenu', 5, { id: 5, name: '组织协作', kind: 'dir', parentID: 0, appID: 1, app: store.get('App', 1), })
  store.set('AppMenu', 6, { id: 6, name: '部门管理', kind: 'menu', route: '/org/departments', parentID: 5, appID: 1, app: store.get('App', 1), })
  store.set('AppMenu', 7, { id: 7, name: '用户管理', kind: 'menu', route: '/org/users', parentID: 5, appID: 1, app: store.get('App', 1), })
  store.set('AppMenu', 8, { id: 8, name: '权限策略', kind: 'menu', route: '/org/policys', parentID: 5, appID: 1, app: store.get('App', 1), })
  store.set('AppMenu', 9, { id: 9, name: '用户组', kind: 'menu', route: '/org/groups', parentID: 5, appID: 1, app: store.get('App', 1), })
  store.set('AppMenu', 10, { id: 10, name: '角色', kind: 'menu', route: '/org/roles', parentID: 5, appID: 1, app: store.get('App', 1), })
  store.set('AppMenu', 11, { id: 11, name: '系统设置', kind: 'dir', parentID: 0, appID: 1, app: store.get('App', 1), })
  store.set('AppMenu', 12, { id: 12, name: '组织管理', kind: 'menu', route: '/system/org', parentID: 11, appID: 1, app: store.get('App', 1), })
  store.set('AppMenu', 13, { id: 13, name: '账户管理', kind: 'menu', route: '/system/account', parentID: 11, appID: 1, app: store.get('App', 1), })
  store.set('AppMenu', 14, { id: 14, name: '应用管理', kind: 'menu', route: '/system/app', parentID: 11, appID: 1, app: store.get('App', 1), })

  // AppRole
  store.set('AppRole', 1, { id: 1, name: 'app1Role1', autoGrant: true, editable: true, appID: 1, app: store.get('App', 1) })
  store.set('AppRole', 2, { id: 2, name: 'app1Role2', autoGrant: false, editable: false, appID: 1, app: store.get('App', 1) })
  store.set('AppRole', 3, { id: 3, name: 'app1Role3', autoGrant: true, editable: false, appID: 1, app: store.get('App', 1) })
  store.set('AppRole', 4, { id: 4, name: 'app1Role4', autoGrant: false, editable: true, appID: 1, app: store.get('App', 1) })

  // AppRes
  store.set('AppRes', 1, {
    id: 1, name: 'app1Res1', typeName: 'App', arnPattern: ':tenant_id:App:private/*:', appID: 1, app: store.get('App', 1),
  })

  //  Permission
  store.set('Permission', 1, {
    id: 1, principalKind: 'user',
    orgID: 1, org: store.get('Org', 1),
    userID: 1, user: store.get('User', 1),
    orgPolicyID: 1, orgPolicy: store.get('OrgPolicy', 1)
  })
  store.set('Permission', 2, {
    id: 2, principalKind: 'role',
    orgID: 1, org: store.get('Org', 1),
    roleID: 3, role: store.get('OrgRole', 3),
    orgPolicyID: 1, orgPolicy: store.get('OrgPolicy', 1)
  })

  // OrgUserPreference
  store.set('OrgUserPreference', 1, {
    id: 1, userID: 1, orgID: 1, menuFavorite: null, menuRecent: null, user: store.get("User", 1), org: store.get('Org', 1)
  })
}
