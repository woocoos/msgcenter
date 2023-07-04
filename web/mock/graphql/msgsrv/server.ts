import { makeExecutableSchema } from '@graphql-tools/schema';
import { addMocksToSchema, createMockStore, mockServer, relayStylePaginationMock } from '@graphql-tools/mock';
import { readFileSync } from "fs";
import { join } from "path";
import * as casual from "casual";
import { initStoreData } from "./store";

const preserveResolvers = true
const typeDefs = readFileSync(join(process.cwd(), 'script', '__generated__', "msgsrv.graphql"), 'utf-8');
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
    Query: {
      msgChannels: relayStylePaginationMock(store),
      msgEvents: relayStylePaginationMock(store),
      msgTemplates: relayStylePaginationMock(store),
      msgTypes: relayStylePaginationMock(store),
      node: (root, args, context, info) => {
        const decoded = Buffer.from(args.id, 'base64').toString()
        const [type, did] = decoded?.split(':', 2)
        const nType = type.split('_').map(t => t.slice(0, 1).toUpperCase() + t.slice(1)).join('')
        return store.get(nType, did)
      }
    },
    Mutation: {
    }
  }
})


export default mockServer(schemaWithMocks, mocks, preserveResolvers)

