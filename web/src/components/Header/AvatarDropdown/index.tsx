import React, { useCallback } from 'react';
import { LogoutOutlined } from '@ant-design/icons';
import { Dropdown, Avatar } from 'antd';
import styles from './index.module.css';
import store from '@/store';
import { useTranslation } from 'react-i18next';
import { checkLave } from '@/components/LeavePrompt';

interface AvatarDropdownProps {
  name: string;
  avatar: string;
}

const AvatarDropdown: React.FC<AvatarDropdownProps> = ({ name, avatar }) => {
  const { t } = useTranslation(),
    [, basisDispatcher] = store.useModel('basis');

  const onMenuClick = useCallback((event) => {
    const { key } = event;
    if (key === 'logout') {
      checkLave(() => {
        // 即使退出接口异常前端也需要直接退出掉所以不需要同步处理
        basisDispatcher.logout();
      });
    }
  }, []);

  const menu = {
    items: [
      {
        key: 'logout',
        label: t('logout'),
        icon: <LogoutOutlined />,
        onClick: onMenuClick,
        className: styles.menu,
      },
    ],
  };
  return (
    <Dropdown menu={menu}>
      <span className={`${styles.action} ${styles.account}`}>
        <Avatar
          size="small"
          className={styles.avatar}
          src={avatar}
          alt="avatar"
        />
        <span>{name}</span>
      </span>
    </Dropdown>
  );
};

export default AvatarDropdown;
