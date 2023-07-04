import i18n from 'i18next';
import { initReactI18next } from 'react-i18next';
import enUS from './locales/en-US';
import zhCN from './locales/zh-CN';

// 多语言文件
const resources = {
  'en-US': enUS,
  'zh-CN': zhCN,
};


i18n
  .use(initReactI18next) // passes i18n down to react-i18next
  .init({
    resources,
    lng: navigator.language, // language to use, more information here: https://www.i18next.com/overview/configuration-options#languages-namespaces-resources
    // you can use the i18n.changeLanguage function to change the language manually: https://www.i18next.com/overview/api#changelanguage
    // if you're using a language detector, do not define the lng option
    interpolation: {
      escapeValue: false, // react already safes from xss
    },
  });

export default i18n;
