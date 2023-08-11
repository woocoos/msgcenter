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
      { name: '静默消息', path: '/msg/silence' },
    ]
  },
];

/**
 * 菜单的处理
 * @returns
 */
export const userMenuList = async () => {
  return asideMenuConfig;
};


