import { makeExecutableSchema } from '@graphql-tools/schema';
import { Ref, addMocksToSchema, createMockStore, mockServer, relayStylePaginationMock } from '@graphql-tools/mock';
import { readFileSync } from "fs";
import { join } from "path";
import * as casual from "casual";
import { addListTemp, delListTemp, initStoreData, listTemp } from "./store";

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
    Query: {
      msgChannels: relayStylePaginationMock(store),
      msgEvents: relayStylePaginationMock(store),
      msgTemplates: relayStylePaginationMock(store),
      msgTypes: relayStylePaginationMock(store),
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
