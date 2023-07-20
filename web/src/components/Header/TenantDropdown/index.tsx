import { Dropdown, MenuProps, Modal } from 'antd';
import { useTranslation } from 'react-i18next';
import { checkLave } from '@/components/LeavePrompt';
import store from '@/store';
import { useEffect, useState } from 'react';
import { userRootOrgs } from '@/services/adminx/user';
import { Org } from '@/__generated__/adminx/graphql';
import styles from './index.module.css';

export default () => {
  const { t } = useTranslation(),
    [userState, userDispatcher] = store.useModel('user'),
    [orgInfo, setOrgInfo] = useState<Org>(),
    [menu, setMenu] = useState<MenuProps>();

  const
    getRequest = async () => {
      const result = await userRootOrgs();
      if (result) {
        setOrgInfo(result.find(item => item?.id == userState.tenantId) as Org);
        setMenu({
          items: result.filter(item => item?.id != userState.tenantId).map(item => {
            return {
              key: item?.id || '',
              label: item?.name || '',
              onClick: onMenuClick,
            };
          }),
        });
      }
    },
    onMenuClick = (info) => {
      const { key } = info;
      checkLave(() => {
        Modal.confirm({
          title: t('tenant_switch_reminder'),
          content: t('tenant_switch_context'),
          onOk: () => {
            userDispatcher.saveTenantId(key);
            location.reload();
          },
        });
      });
    };


  useEffect(() => {
    if (document.hidden) {
      const tipStr = t('tenant_switch_title');
      document.title = tipStr;
      document.body.innerHTML = `<div style="width:370px;margin:40px auto 0 auto;">${tipStr}</div>`;
      window.close();
    }
  }, [userState.tenantId]);

  useEffect(() => {
    getRequest();
  }, []);

  return orgInfo ? <Dropdown menu={menu} disabled={menu?.items?.length === 0}>
    <span className={styles.action}>
      <span>{orgInfo.name}</span>
    </span>
  </Dropdown> : <></>;
};
