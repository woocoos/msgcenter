import type { Request, Response } from '@ice/app';
import bodyParser from 'body-parser'
import adminxServer from "./graphql/adminx/server";
import msgsrvServer from "./graphql/msgsrv/server";

/**
 * 文档
 * https://the-guild.dev/graphql/tools/docs/api/modules/mock_src
 */
export default {
  'POST /mock-api-adminx/graphql/query': (request: Request, response: Response) => {
    bodyParser.json()(request, response, async () => {
      const { query, variables } = request.body;
      const result = await adminxServer.query(query as string, variables as any)
      response.send(result);
    })
  },
  'POST /mock-api-msgsrv/graphql/query': (request: Request, response: Response) => {
    bodyParser.json()(request, response, async () => {
      const { query, variables } = request.body;
      const result = await msgsrvServer.query(query as string, variables as any)
      response.send(result);
    })
  },
}
