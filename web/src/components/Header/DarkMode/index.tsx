import React from 'react';
import { Switch } from 'antd';
import store from '@/store';
import { useTranslation } from 'react-i18next';

const I18nDropdown: React.FC = () => {
  const { t } = useTranslation(),
    [appState, appDispatcher] = store.useModel('app');

  return (
    <Switch
      style={{ margin: '0 12px' }}
      checkedChildren={t('bright')}
      unCheckedChildren={t('dark')}
      defaultChecked={!appState.darkMode}
      onChange={(checked) => {
        appDispatcher.updateDarkMode(!checked);
      }}
    />
  );
};

export default I18nDropdown;
