import { makeExecutableSchema } from '@graphql-tools/schema';
import { addMocksToSchema, createMockStore, mockServer, relayStylePaginationMock } from '@graphql-tools/mock';
import { readFileSync } from "fs";
import { join } from "path";
import * as casual from "casual";
import { initStoreData, listTemp } from "./store";

const preserveResolvers = true
const typeDefs = readFileSync(join(process.cwd(), 'script', 'generated', "adminx.graphql"), 'utf-8');
const schema = makeExecutableSchema({ typeDefs });
const mocks = {
  ID: () => casual.integer(1, 1000000000),
  Time: () => casual.date('YYYY-MM-DDTHH:mm:ss.SSSZZ'),
  Cursor: () => casual._string(),
  GID: () => casual._string(),
  Query: {},
  Mutation: {},
}

const store = createMockStore({ schema, mocks })

initStoreData(store)

const schemaWithMocks = addMocksToSchema({
  schema,
  store,
  preserveResolvers,
  resolvers: {
    App: {
      menus: relayStylePaginationMock(store),
      actions: relayStylePaginationMock(store),
      resources: relayStylePaginationMock(store),
      orgs: relayStylePaginationMock(store),
    },
    Org: {
      users: relayStylePaginationMock(store),
      permissions: relayStylePaginationMock(store),
      policies: relayStylePaginationMock(store),
      apps: relayStylePaginationMock(store),
    },
    User: {
      permissions: relayStylePaginationMock(store),
    },
    Query: {
      fileIdentitiesForOrg: () => [
        store.get('OrgFileIdentity', 1),
      ],
      appAccess: () => true,
      apps: relayStylePaginationMock(store),
      organizations: (_, { where }) => {
        if (where.kind === 'org') {
          return listTemp([
            store.get('Org', 2),
            store.get('Org', 3),
            store.get('Org', 4),
          ])
        } else {
          return listTemp([
            store.get('Org', 1),
            store.get('Org', 5),
          ])
        }
      },
      users: relayStylePaginationMock(store),
      orgGroups: relayStylePaginationMock(store),
      orgRoleUsers: relayStylePaginationMock(store),
      orgRoles: relayStylePaginationMock(store),
      appRoleAssignedToOrgs: () => [
        store.get('Org', 1),
      ],
      appPolicyAssignedToOrgs: () => [
        store.get('Org', 1),
      ],
      orgUserPreference: () => store.get('OrgUserPreference', 1),
      orgPolicyReferences: relayStylePaginationMock(store),
      appResources: relayStylePaginationMock(store),
      orgAppResources: relayStylePaginationMock(store),
      userGroups: relayStylePaginationMock(store),
      userExtendGroupPolicies: relayStylePaginationMock(store),
      userMenus: () => [
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
      ],
      userPermissions: () => [
        store.get('AppAction', 1),
        // store.get('AppAction', 2),
        // store.get('AppAction', 3),
        // store.get('AppAction', 4),
        // store.get('AppAction', 5),
        // store.get('AppAction', 6),
        // store.get('AppAction', 7),
      ],
      checkPermission: (_, { permission }) => {
        // permission => appCode:action
        return true;
      },
      orgAppActions: () => [
        store.get('AppAction', 1),
        store.get('AppAction', 2),
        store.get('AppAction', 3),
      ],
      userRootOrgs: () => [
        store.get('Org', 1),
      ],
      userApps: () => [
        store.get('App', 1),
      ],
      orgRecycleUsers: relayStylePaginationMock(store),
      appDictByRefCode: (_, { refCodes }) => {
        return [
          store.get('AppDict', 1),
        ]
      },
      appDictItemByRefCode: (_, { refCodes }) => {
        return [
          store.get('AppDictItem', 1),
          store.get('AppDictItem', 2),
        ]
      },
      globalID: (_, { type, id }) => btoa(`${type}:${id}`),
      nodes: (_, args) => {
        return args.ids.map(gid => {
          const { type, id } = parseGid(gid);
          return store.get(type, id)
        })
      },
      node: (root, args, context, info) => {
        const { type, id } = parseGid(args.id);
        return store.get(type, id)
      }
    },
    Mutation: {
      updateUser: (_, { userID, input }) => {
        store.set('User', userID, input)
        return store.get('User', userID)
      },
      saveOrgUserPreference: (_, { input }) => {
        if (input.menuFavorite) {
          store.set("OrgUserPreference", 1, 'menuFavorite', input.menuFavorite)
        } else if (input.menuRecent) {
          store.set("OrgUserPreference", 1, 'menuRecent', input.menuRecent)
        }
        return { id: 1 };
      }
    }
  }
})

/**
 * 解析gid
 * @param gid
 * @returns
 */
function parseGid(gid: string) {
  const decoded = Buffer.from(gid, 'base64').toString()
  const [type, did] = decoded?.split(':', 2)
  const nType = type.split('_').map(t => t.slice(0, 1).toUpperCase() + t.slice(1)).join('')
  return {
    type: nType,
    id: did,
  }
}

export default mockServer(schemaWithMocks, mocks, preserveResolvers)
