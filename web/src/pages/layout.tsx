import Layout from '@/components/Layout';
import LayoutChild from '@/components/Layout/child';
import { isInIcestark } from '@ice/stark-app';

export default () => {
  return isInIcestark() ? <LayoutChild /> : <Layout />
}
