import { createModel } from 'ice';
import { setItem } from '@/pkg/localStore';
import { CurrentLanguages } from '@/i18n';


type ModelState = {
  locale: CurrentLanguages;
  darkMode: boolean;
  compactMode: boolean;
};


export default createModel({
  state: {
    locale: CurrentLanguages.zhCN,
    darkMode: false,
    compactMode: false,
  } as ModelState,
  reducers: {
    updateLocale(prevState: ModelState, payload: CurrentLanguages) {
      setItem('locale', payload);
      prevState.locale = payload;
    },
    updateDarkMode(prevState: ModelState, payload: boolean) {
      setItem('darkMode', payload);
      prevState.darkMode = payload;
    },
  },
  effects: () => ({}),
});
