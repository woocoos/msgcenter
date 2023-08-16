import { defineConfig } from '@ice/app';
import request from '@ice/plugin-request';
import store from '@ice/plugin-store';
import auth from '@ice/plugin-auth';
import antd from '@ice/plugin-antd';
import jsxPlus from '@ice/plugin-jsx-plus';
import icestark from '@ice/plugin-icestark';
import urqlPlugin from '@knockout-js/ice-urql';

// The project config, see https://v3.ice.work/docs/guide/basic/config
const minify = process.env.NODE_ENV === 'production' ? 'swc' : false;

export default defineConfig(() => ({
  ssg: false,
  ssr: false,
  minify,
  codeSplitting: 'page',
  devPublicPath: process.env.ICE_DEV_PUBLIC_PATH,
  publicPath: process.env.ICE_BUILD_PUBLIC_PATH,
  compileDependencies: process.env.NODE_ENV === 'development' ? [/@urql\/core/] : true,
  routes: {
    ignoreFiles: [
      '**/components/**',   // 添加此配置忽略components被解析成路由组件
    ],
  },
  plugins: [
    icestark({ type: 'child' }),
    urqlPlugin(),
    request(),
    store(),
    jsxPlus(),
    antd({
      importStyle: false,
    }),
    auth(),
  ],
  proxy: {
    '/api-msgsrv': {
      target: process.env.ICE_PROXY_MSGSRV,
      changeOrigin: true,
      pathRewrite: { '^/api-msgsrv': '' },
    },
    '/api-adminx': {
      target: process.env.ICE_PROXY_ADMINX,
      changeOrigin: true,
      pathRewrite: { '^/api-adminx': '' },
    },
    '/api-files': {
      target: process.env.ICE_PROXY_FILES,
      changeOrigin: true,
      pathRewrite: { '^/api-files': '' },
    },
  },
}));

