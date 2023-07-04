import React, { useState, useEffect } from 'react';
import { Dropdown } from 'antd';
import styles from './index.module.css';
import store from '@/store';
import { LocalLanguage } from '@/models/basis';

const I18nDropdown: React.FC = () => {
  const [basisState, basisDispatcher] = store.useModel('basis');
  const [locale, setLocale] = useState('');

  const
    onMenuClick = (ev) => {
      const mItem = menu.items.find((item) => item.key === ev.key);
      if (mItem) {
        basisDispatcher.updateLocale(mItem.key as LocalLanguage);
        setLocale(mItem.label);
      }
    };

  const menu = {
    items: [
      {
        key: 'zh-CN',
        label: '简体',
        onClick: onMenuClick,
      },
      {
        key: 'en-US',
        label: 'English',
        onClick: onMenuClick,
      },
    ],
  };

  useEffect(() => {
    const mItem = menu.items.find((item) => item.key === basisState.locale);
    if (mItem) {
      setLocale(mItem.label);
    }
  }, [basisState.locale]);

  return (
    <Dropdown menu={menu}>
      <span className={`${styles.action} ${styles.account}`}>
        <span>{locale}</span>
      </span>
    </Dropdown>
  );
};

export default I18nDropdown;
