import store from '@/store';
import { useLocation, useSearchParams } from '@ice/runtime';
import { ReactNode, useEffect } from 'react';
import KeepAlive, { useAliveController } from 'react-activation';

export default (props: {
  children: ReactNode;
  clearAlive?: boolean;
}) => {
  const [userState] = store.useModel('user'),
    [searchParams] = useSearchParams(),
    location = useLocation(),
    id = searchParams.get('id') || userState.tenantId,
    cacheKey = btoa(location.pathname),
    { dropScope, getCachingNodes } = useAliveController();

  useEffect(() => {
    if (props.clearAlive) {
      getCachingNodes().forEach(item => {
        if (item.name && cacheKey != item.name) {
          dropScope(item.name);
        }
      });
    }
  }, [props.clearAlive]);

  return (<KeepAlive when autoFreeze={false} cacheKey={cacheKey} name={cacheKey} id={id}>
    {props.children}
  </KeepAlive>);
};
