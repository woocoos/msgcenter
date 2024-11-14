# 消息中心-ui

消息中心ui模块

## 快速开始

### .env.local
```
# 端口冲突可自行替换
PORT=3002

# mock
ICE_PROXY_ADMINX=http://127.0.0.1:3002/mock-api-adminx/
ICE_PROXY_AUTH=http://127.0.0.1:3002/mock-api-auth/
ICE_PROXY_MSGSRV=http://127.0.0.1:3002/mock-api-msgsrv/

```
###  运行启动前
```shell
pnpm i --registry=https://registry.npmmirror.com
```

###  运行启动

```shell
# 启动开发环境
pnpm dev
```
