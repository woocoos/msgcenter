import { useEffect, useState } from "react"
import { getDate } from "@/util";
import { FloatButton, notification } from "antd"
import { Message } from "@/generated/msgsrv/graphql";
import store from "@/store";
import { subscription } from "@knockout-js/ice-urql/request";
import { subMessage } from "@/services/msgsrv/internal/ws";
import { setItem } from "@/pkg/localStore";

export interface WsMsgProps {
  /**
   * 最多显示几条
   */
  maxLength?: number;
  /**
   * 选中一项
   * @param data
   * @returns
   */
  onItemClick?: (data: Message) => void;
  /**
   * 有信消息
   * @returns
   */
  onListenerNewMsg?: () => void;
}

const actions = ['internal'];

/**
 * extras{action,actionID}
 */
export default (props: WsMsgProps) => {
  const [unreadNum, setUnreadNum] = useState(0),
    [docHidden, setDocHidden] = useState(document.hidden),
    [api, contextHolder] = notification.useNotification(),
    [showMsg, setShowMsg] = useState(false),
    [wsState, wsDispatch] = store.useModel('ws'),
    maxLength = props.maxLength ?? 3;

  const handleVisibilitychange = () => {
    setDocHidden(document.hidden);
  }

  useEffect(() => {
    if (showMsg) {
      api.destroy();
      const length = wsState.message.length > maxLength ? maxLength : wsState.message.length;
      for (let i = 0; i < length; i++) {
        const item = wsState.message[i], key = item.extras.actionID;
        api.open({
          key,
          duration: null,
          message: item.title,
          description: getDate(item.sendAt, 'YYYY-MM-DD HH:mm:ss'),
          btn: <>
            <a onClick={() => {
              props.onItemClick?.(item);
              wsDispatch.setMessage(wsState.message.filter(item => item.extras.actionID != key));
            }}>查看</a>
          </>,
          onClose: () => {
            wsDispatch.setMessage(wsState.message.filter(item => item.extras.actionID != key));
          },
        })
      }
    } else {
      api.destroy();
    }
    setUnreadNum(wsState.message?.length);
  }, [wsState.message, showMsg])

  useEffect(() => {
    if (!docHidden && !wsState.handshake) {
      wsDispatch.setHandshake(true);
      subscription(subMessage, {}).subscribe(result => {
        if (result.data?.message) {
          const newMsg = result.data?.message as Message,
            newWs = store.getModel('ws');
          // 判断action是否是支持的action
          if (actions.includes(newMsg.extras.action)) {
            // 判断actionID是否已经存在不存在则设置到message中
            if (!newWs[0].message?.find(item => item.extras.actionID == newMsg.extras.actionID)) {
              newWs[1].setMessage([newMsg, ...(newWs[0].message ?? [])]);
              props.onListenerNewMsg?.();
            }
          }
        }
      })
      window.addEventListener('beforeunload', () => {
        setItem('handshake', false);
      }, false);
    }
  }, [wsState.handshake, docHidden])

  useEffect(() => {
    window.addEventListener("visibilitychange", handleVisibilitychange);

    return () => {
      window.removeEventListener("visibilitychange", handleVisibilitychange);
    }
  }, [])

  return <>
    <FloatButton
      tooltip={!showMsg && unreadNum > 0 ? <div>有新消息</div> : undefined}
      type={showMsg ? "primary" : "default"}
      badge={{ count: unreadNum, color: 'red' }}
      onClick={() => {
        setShowMsg(!showMsg);
      }}
    />
    {contextHolder}
  </>
}
