import Layout from '@/components/Layout';
import LayoutStark from '@/components/Layout/stark';
import { isInIcestark } from '@ice/stark-app';

export default () => {
  return isInIcestark() ? <LayoutStark /> : <Layout />
}
