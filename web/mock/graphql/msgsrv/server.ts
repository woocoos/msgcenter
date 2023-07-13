import { makeExecutableSchema } from '@graphql-tools/schema';
import { Ref, addMocksToSchema, createMockStore, mockServer, relayStylePaginationMock } from '@graphql-tools/mock';
import { readFileSync } from "fs";
import { join } from "path";
import * as casual from "casual";
import { addList, addListTemp, delList, delListTemp, getObject, initStoreData, listTemp } from "./store";
import * as yaml from 'js-yaml'

const preserveResolvers = true

const typeDefs = readFileSync(join(process.cwd(), 'script', '__generated__', "msgsrv.graphql"), 'utf-8');

const schema = makeExecutableSchema({ typeDefs });

const mocks = {
  ID: () => casual.integer(1, 1000000000),
  Time: () => casual.date('YYYY-MM-DDTHH:mm:ss.SSSZZ'),
  Cursor: () => casual._string(),
  GID: () => casual._string(),
  MapString: () => { },
  HostPort: () => casual._string(),
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
    MsgEvent: {
      routeStr: (pRef: Ref, { type }) => {
        const routeRef = store.get(pRef, 'route') as Ref
        const route = getObject(store, routeRef)
        if (route) {
          if (type === 'Json') {
            return JSON.stringify(route);
          } else if (type === 'Yaml') {
            return yaml.dump(route)
          }
        } else {
          if (type === 'Json') {
            return readFileSync(join(process.cwd(), 'mock', 'graphql', 'msgsrv', "route.json"), 'utf-8')
          } else if (type == 'Yaml') {
            return readFileSync(join(process.cwd(), 'mock', 'graphql', 'msgsrv', "route.yaml"), 'utf-8')
          }
        }
      }
    },
    Query: {
      msgChannels: relayStylePaginationMock(store),
      msgEvents: relayStylePaginationMock(store),
      msgTemplates: relayStylePaginationMock(store),
      msgTypes: relayStylePaginationMock(store),
      silences: relayStylePaginationMock(store),
      msgTypeCategories: () => {
        return ['故障消息', '业务消息', '客户交易'];
      },
      nodes: (_, args) => {
        return args.ids.map(gid => {
          const { type, id } = parseGid(gid);
          return store.get(type, id)
        })
      },
      node: (root, args, context, info) => {
        const { type, id } = parseGid(args.id)
        return store.get(type, id)
      }
    },
    Mutation: {
      createSilence(_, { input }) {
        input.id = `${Date.now()}`;
        store.set('Silence', input.id, input)
        return addListTemp(
          store,
          store.get('Query', 'ROOT', 'silences') as Ref,
          store.get('Silence', input.id) as Ref
        );
      },
      updateSilence(_, { id, input }) {
        store.set('Silence', id, input)
        return store.get('Silence', id)
      },
      deleteSilence(_, { id }) {
        delListTemp(
          store,
          store.get('Query', 'ROOT', 'silences') as Ref,
          id,
        )
        return true;
      },
      createMsgSubscriber(_, { inputs }) {
        const ids: string[] = [];
        inputs.forEach((input, index) => {
          input.id = `${Date.now()}-${index}`;
          ids.push(input.id);
          store.set('MsgSubscriber', input.id, input);
          if (input.userID && !input.exclude) {
            addList(
              store,
              store.get('MsgType', input.msgTypeID) as Ref,
              'subscriberUsers',
              store.get('MsgSubscriber', input.id) as Ref,
            )
          } else if (input.userID && input.exclude) {
            addList(
              store,
              store.get('MsgType', input.msgTypeID) as Ref,
              'excludeSubscriberUsers',
              store.get('MsgSubscriber', input.id) as Ref,
            )
          } else if (input.orgRoleID && !input.exclude) {
            addList(
              store,
              store.get('MsgType', input.msgTypeID) as Ref,
              'subscriberRoles',
              store.get('MsgSubscriber', input.id) as Ref,
            )
          }
        })

        return ids.map(id => store.get('MsgSubscriber', id));
      },
      deleteMsgSubscriber(_, { ids }) {
        ids.forEach(id => {
          const msgSub = store.get('MsgSubscriber', id) as Ref,
            msgTypeId = store.get(msgSub, 'msgTypeID') as string,
            exclude = store.get(msgSub, 'exclude'),
            orgRoleId = store.get(msgSub, 'orgRoleID'),
            userId = store.get(msgSub, 'userID');

          if (userId && !exclude) {
            delList(
              store,
              store.get('MsgType', msgTypeId) as Ref,
              'subscriberUsers',
              msgSub.$ref.key,
            )
          } else if (userId && exclude) {
            delList(
              store,
              store.get('MsgType', msgTypeId) as Ref,
              'excludeSubscriberUsers',
              msgSub.$ref.key,
            )
          } else if (orgRoleId && !exclude) {
            delList(
              store,
              store.get('MsgType', msgTypeId) as Ref,
              'subscriberRoles',
              msgSub.$ref.key,
            )
          }
        })
        return true;
      },
      createMsgTemplate(_, { input }) {
        input.id = `${Date.now()}`;
        if (input.eventID) {
          input.msgEventID = input.eventID
          input.event = store.get('MsgEvent', input.eventID)
          delete input.eventID;
        }
        store.set('MsgTemplate', input.id, input)
        return addListTemp(
          store,
          store.get('Query', 'ROOT', 'msgTemplates') as Ref,
          store.get('MsgTemplate', input.id) as Ref
        );
      },
      updateMsgTemplate(_, { id, input }) {
        if (input.eventID) {
          input.msgEventID = input.eventID
          input.event = store.get('MsgEvent', input.eventID)
          delete input.eventID;
        }
        store.set('MsgTemplate', id, input)
        return store.get('MsgTemplate', id)
      },
      deleteMsgTemplate(_, { id }) {
        delListTemp(
          store,
          store.get('Query', 'ROOT', 'msgTemplates') as Ref,
          id,
        )
        return true;
      },
      enableMsgTemplate(_, { id }) {
        store.set('MsgTemplate', id, "status", 'active')
        return store.get('MsgTemplate', id)
      },
      disableMsgTemplate(_, { id }) {
        store.set('MsgTemplate', id, "status", 'inactive')
        return store.get('MsgTemplate', id)
      },
      createMsgEvent(_, { input }) {
        input.id = `${Date.now()}`;
        if (input.msgTypeID) {
          input.msgType = store.get('MsgType', input.msgTypeID)
        }
        store.set('MsgEvent', input.id, input)
        return addListTemp(
          store,
          store.get('Query', 'ROOT', 'msgEvents') as Ref,
          store.get('MsgEvent', input.id) as Ref
        );
      },
      updateMsgEvent(_, { id, input }) {
        if (input.msgTypeID) {
          input.msgType = store.get('MsgType', input.msgTypeID)
        }
        // if (input.route) {
        // yaml
        // input.routeStr = yaml.dump(input.route)
        // json
        // input.routeStr = JSON.stringify(routeStr(input.route));
        // }
        store.set('MsgEvent', id, input)
        return store.get('MsgEvent', id)
      },
      deleteMsgEvent(_, { id }) {
        delListTemp(
          store,
          store.get('Query', 'ROOT', 'msgEvents') as Ref,
          id,
        )
        return true;
      },
      enableMsgEvent(_, { id }) {
        store.set('MsgEvent', id, "status", 'active')
        return store.get('MsgEvent', id)
      },
      disableMsgEvent(_, { id }) {
        store.set('MsgEvent', id, "status", 'inactive')
        return store.get('MsgEvent', id)
      },
      createMsgChannel(_, { input }) {
        input.id = `${Date.now()}`;
        store.set('MsgChannel', input.id, input)
        return addListTemp(
          store,
          store.get('Query', 'ROOT', 'msgChannels') as Ref,
          store.get('MsgChannel', input.id) as Ref
        );
      },
      updateMsgChannel(_, { id, input }) {
        store.set('MsgChannel', id, input)
        return store.get('MsgChannel', id)
      },
      deleteMsgChannel(_, { id }) {
        delListTemp(
          store,
          store.get('Query', 'ROOT', 'msgChannels') as Ref,
          id,
        )
        return true;
      },
      enableMsgChannel(_, { id }) {
        store.set('MsgChannel', id, "status", 'active')
        return store.get('MsgChannel', id)
      },
      disableMsgChannel(_, { id }) {
        store.set('MsgChannel', id, "status", 'inactive')
        return store.get('MsgChannel', id)
      },
      createMsgType(_, { input }) {
        input.id = `${Date.now()}`;
        store.set('MsgType', input.id, input)
        return addListTemp(
          store,
          store.get('Query', 'ROOT', 'msgTypes') as Ref,
          store.get('MsgType', input.id) as Ref
        );
      },
      updateMsgType(_, { id, input }) {
        store.set('MsgType', id, input)
        return store.get('MsgType', id)
      },
      deleteMsgType(_, { id }) {
        delListTemp(
          store,
          store.get('Query', 'ROOT', 'msgTypes') as Ref,
          id,
        )
        return true;
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
