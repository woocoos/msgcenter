import { userMenus } from '@/services/adminx/user';
import { formatTreeData } from '@/util';
import {
  ControlOutlined,
} from '@ant-design/icons';
import type { MenuDataItem } from '@ant-design/pro-components';

const asideMenuConfig: MenuDataItem[] = [
  {
    name: '消息中心',
    framework: true,
    icon: <ControlOutlined />,
    children: [
      { name: '消息类型', path: '/msg/type' },
      { name: '消息事件', path: '/msg/event' },
      { name: '消息通道', path: '/msg/channel' },
      { name: '消息订阅', path: '/msg/subscription' },
    ]
  },
];

/**
 * 菜单的处理
 * @returns
 */
export const userMenuList = async () => {
  const list: MenuDataItem[] = [];
  if (process.env.ICE_CORE_MODE === 'development') {
    list.push(...asideMenuConfig);
  } else {
    if (process.env.ICE_APP_CODE) {
      const menus = await userMenus(process.env.ICE_APP_CODE);
      if (menus) {
        const menuList: MenuDataItem[] = [];
        menus.forEach(item => {
          if (item) {
            const data: MenuDataItem = {
              key: item.id,
              id: item.id,
              name: item.name,
              icon: <i className={item.icon || ''} />,
              parentId: item.parentID,
            };
            if (item.route) {
              data.path = item.route;
            }
            menuList.push(data);
          }
        });
        list.push(...formatTreeData(menuList, undefined, { key: 'id' }));
      }
    }
  }
  return list;
};


