import i18n from 'i18next';
import { initReactI18next } from 'react-i18next';
import enUS from './locales/en-US';
import zhCN from './locales/zh-CN';

export enum CurrentLanguages {
  'zhCN' = 'zh-CN',
  'enUS' = 'en-US'
};

// 多语言文件
const resources = {
  [CurrentLanguages.enUS]: enUS,
  [CurrentLanguages.zhCN]: zhCN,
};


i18n
  .use(initReactI18next)
  .init({
    resources,
    lng: CurrentLanguages.zhCN,
    interpolation: {
      escapeValue: false,
    },
  });

export default i18n;
