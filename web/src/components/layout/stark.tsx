import i18n from "@/i18n";
import store from "@/store";
import { Outlet, useLocation } from "@ice/runtime"
import { useEffect } from "react";
import { CollectProviders } from "@knockout-js/layout";

const ICE_APP_CODE = process.env.ICE_APP_CODE ?? '';

export default () => {
  const [appState] = store.useModel('app'),
    [userState] = store.useModel('user'),
    location = useLocation();

  useEffect(() => {
    i18n.changeLanguage(appState.locale);
  }, [appState.locale]);

  return <CollectProviders
    locale={appState.locale}
    dark={appState.darkMode}
    pathname={location.pathname}
    appCode={ICE_APP_CODE}
    tenantId={userState.tenantId}
  >
    <Outlet />
  </CollectProviders>;
}
