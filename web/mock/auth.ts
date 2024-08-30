import type { Request, Response } from '@ice/app';

export default {
  'POST /mock-api-auth/oss/sts': (request: Request, response: Response) => {
    // 本地的  access_key_id  secret_access_key 需要手动修改
    response.send({
      access_key_id: 'oTbKaIMXjCnx3HzrH5Qo',
      secret_access_key: 'XfvZSw6a954U77hZ3rrSr6jmDKEPFhDNHOJJbW4B',
      expiration: Date.now() + 1000 * 60 * 60 * 24,
      session_token: undefined,
    });
  },
}
