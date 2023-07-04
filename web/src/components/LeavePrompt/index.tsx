import { history, useLocation } from 'ice';
import { ReactNode, useEffect } from 'react';
import i18n from '@/i18n';


const pathName = {
  when: true,
};

/**
 * 拦截离开检查确认离开后回调
 * @param callback 确认离开
 */
export const checkLave = (callback: () => void) => {
  if (pathName.when) {
    callback();
  } else {
    if (confirm(`${i18n.t('leave_prompt_tip')}`)) {
      pathName.when = true;
      callback();
    }
  }
};

/**
 * 设置是否拦截离开
 * @param when true:不拦截，false:拦截
 */
export const setLeavePromptWhen = (when: boolean) => {
  pathName.when = when;
};

/**
 * 一般在layout引入，主要检测浏览器刷新
 * TODO：浏览器的前进和回退无法拦截
 */
export default () => {
  const location = useLocation();

  useEffect(() => {
    setLeavePromptWhen(true);
  }, [location.pathname]);

  useEffect(() => {
    window.addEventListener('beforeunload', (event) => {
      if (pathName.when === true) {
        return;
      }
      event.preventDefault();
      event.returnValue = i18n.t('leave_prompt_tip');
      return i18n.t('leave_prompt_tip');
    });
  }, []);

  return <></>;
};

/**
 * 使用在Link 标签上的处理
 * @param props
 * @returns
 */
export const Link = (props: {
  to: string;
  children: ReactNode;
}) => {
  return (<a onClick={() => {
    checkLave(() => {
      pathName.when = true;
      history?.push(props.to);
    });
  }}
  >
    {props.children}
  </a>);
};
