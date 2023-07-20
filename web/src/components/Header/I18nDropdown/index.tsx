import React, { useState, useEffect } from 'react';
import { Dropdown } from 'antd';
import styles from './index.module.css';
import store from '@/store';
import { CurrentLanguages } from '@/i18n';

const I18nDropdown: React.FC = () => {
  const [appState, appDispatcher] = store.useModel('app');
  const [locale, setLocale] = useState('');

  const
    onMenuClick = (ev) => {
      const mItem = menu.items.find((item) => item.key === ev.key);
      if (mItem) {
        appDispatcher.updateLocale(mItem.key);
        setLocale(mItem.label);
      }
    };

  const menu = {
    items: [
      {
        key: CurrentLanguages.zhCN,
        label: '简体',
        onClick: onMenuClick,
      },
      {
        key: CurrentLanguages.enUS,
        label: 'English',
        onClick: onMenuClick,
      },
    ],
  };

  useEffect(() => {
    const mItem = menu.items.find((item) => item.key === appState.locale);
    if (mItem) {
      setLocale(mItem.label);
    }
  }, [appState.locale]);

  return (
    <Dropdown menu={menu}>
      <span className={styles.action}>
        <span>{locale}</span>
      </span>
    </Dropdown>
  );
};

export default I18nDropdown;
