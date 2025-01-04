import Layout from '@/components/layout';
import LayoutStark from '@/components/layout/stark';
import { isInIcestark } from '@ice/stark-app';
import { ConfigProvider } from 'antd';
import { store as starkStore } from '@ice/stark-data';

export default () => {
  return isInIcestark() ?
    <ConfigProvider
      theme={starkStore.get('config-provider-theme') ?? undefined}
    >
      <LayoutStark />
    </ConfigProvider>
    : <Layout />
}
