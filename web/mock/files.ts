
import type { Request, Response } from '@ice/app';
import { readFileSync } from "fs";
import { join } from "path";


export default {
  'POST /api-files/files': (request: Request, response: Response) => {
    response.send("123");
  },
  'GET /api-files/files/:fileId': (request: Request, response: Response) => {
    response.send({
      id: '123',
      name: 'hello word',
      size: 5000,
      createdAt: new Date()
    });
  },
  'GET /api-files/files/:fileId/raw': (request: Request, response: Response) => {
    const file = readFileSync(join(process.cwd(), 'src', 'assets', 'images', "woocoo.png"))
    response.setHeader('contetn-type', 'image/png')
    response.send(file)
  },
  'DELETE /api-files/files/:fileId': (request: Request, response: Response) => {
    response.send("")
  },
}
