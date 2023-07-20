import i18n, { CurrentLanguages } from "@/i18n";
import store from "@/store";
import { ProConfigProvider } from "@ant-design/pro-components";
import { Outlet } from "@ice/runtime"
import { Locale } from "antd/es/locale";
import { useEffect, useState } from "react";
import { AliveScope } from "react-activation";
import zhCN from 'antd/locale/zh_CN';
import enUS from 'antd/locale/en_US';
import { ConfigProvider } from "antd";

export default () => {
  const [appState] = store.useModel('app'),
    [locale, setLocale] = useState<Locale>();

  useEffect(() => {
    i18n.changeLanguage(appState.locale);
    switch (appState.locale) {
      case CurrentLanguages.zhCN:
        setLocale(zhCN)
        break;
      case CurrentLanguages.enUS:
        setLocale(enUS)
        break;
      default:
        setLocale(zhCN)
        break;
    }
  }, [appState.locale]);

  return <ProConfigProvider dark={appState.darkMode} >
    <ConfigProvider locale={locale}>
      <AliveScope>
        <Outlet />
      </AliveScope>
    </ConfigProvider>
  </ProConfigProvider>
}
