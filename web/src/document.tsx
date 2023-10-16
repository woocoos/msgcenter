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
          <script crossOrigin="" src="https://jsd.onmicrosoft.cn/npm/react@18.2.0/umd/react.production.min.js"></script>
          <script crossOrigin="" src="https://jsd.onmicrosoft.cn/npm/react-dom@18.2.0/umd/react-dom.production.min.js"></script>
          <script crossOrigin="" src="https://jsd.onmicrosoft.cn/npm/i18next@22.5.0/i18next.min.js"></script>
          <script crossOrigin="" src="https://jsd.onmicrosoft.cn/npm/react-i18next@12.3.1/react-i18next.min.js"></script>
          <script crossOrigin="" src="https://jsd.onmicrosoft.cn/npm/js-yaml@4.1.0/dist/js-yaml.min.js"></script>
          <script crossOrigin="" src="https://jsd.onmicrosoft.cn/npm/dayjs@1.11.10/dayjs.min.js"></script>
          <script crossOrigin="" src="https://jsd.onmicrosoft.cn/npm/antd@5.6.3/dist/antd.min.js"></script>
          <script crossOrigin="" src="https://jsd.onmicrosoft.cn/npm/@ant-design/pro-components@2.6.28/dist/pro-components.min.js"></script>
        </> : <></>}
        <Scripts />
      </body>
    </html>
  );
}
