import { Meta, Title, Links, Main, Scripts } from 'ice';

const ICE_STATIC_CDN = process.env.ICE_STATIC_CDN ?? '';

export default function Document() {
  return (
    <html>
      <head>
        <meta charSet="utf-8" />
        <meta name="description" content="ice.js 3 antd pro scaffold" />
        <link rel="icon" href="/favicon.ico" type="image/x-icon" />
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" />
        <Meta />
        <Title />
        <Links />
      </head>
      <body>
        <Main />
        {ICE_STATIC_CDN ? <>
          <script crossOrigin="" src="https://qlcdn.oss-cn-shenzhen.aliyuncs.com/cdn/react/18.2.0/index.min.js"></script>
          <script crossOrigin="" src="https://qlcdn.oss-cn-shenzhen.aliyuncs.com/cdn/react-dom/18.2.0/index.min.js"></script>
          <script crossOrigin="" src="https://qlcdn.oss-cn-shenzhen.aliyuncs.com/cdn/i18next/22.5.0/index.min.js"></script>
          <script crossOrigin="" src="https://qlcdn.oss-cn-shenzhen.aliyuncs.com/cdn/react-i18next/12.3.1/index.min.js"></script>
          <script crossOrigin="" src="https://qlcdn.oss-cn-shenzhen.aliyuncs.com/cdn/%40monaco-editor/loader/1.4.0/index.min.js"></script>
          <script crossOrigin="" src="https://qlcdn.oss-cn-shenzhen.aliyuncs.com/cdn/jwt-decode/3.1.2/index.min.js"></script>
          <script crossOrigin="" src="https://qlcdn.oss-cn-shenzhen.aliyuncs.com/cdn/js-yaml/4.1.0/index.min.js "></script>
          <script crossOrigin="" src="https://qlcdn.oss-cn-shenzhen.aliyuncs.com/cdn/dayjs/1.11.8/index.min.js "></script>
          <script crossOrigin="" src="https://qlcdn.oss-cn-shenzhen.aliyuncs.com/cdn/antd/5.6.3/index.min.js"></script>
          <script crossOrigin="" src="https://qlcdn.oss-cn-shenzhen.aliyuncs.com/cdn/%40ant-design/pro-components/2.6.28/index.min.js"></script>
        </> : <></>}
        <Scripts />
      </body>
    </html>
  );
}
