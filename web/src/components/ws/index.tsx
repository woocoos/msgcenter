import { MsgInternal } from "@/generated/msgsrv/graphql";
import { getMsgInternalList, markMsgRead } from "@/services/msgsrv/internal";
import { getDate } from "@/util";
import { FloatButton, notification } from "antd"
import { useEffect, useState } from "react"

export default () => {
  const [unreadNum, setUnreadNum] = useState(0),
    [dataSource, setDataSource] = useState<MsgInternal[]>([]),
    [api, contextHolder] = notification.useNotification(),
    [showMsg, setShowMsg] = useState(false);

  const
    requestUnreadMsg = async () => {
      const result = await getMsgInternalList({
        current: 1,
        pageSize: 1000,
        where: {}
      }), list: MsgInternal[] = [];

      result?.edges?.forEach(item => {
        if (item?.node) {
          list.push(item.node)
        }
      })

      setDataSource(list);
      setUnreadNum(result?.totalCount ?? 0);
    },
    setRead = async (id: string) => {
      await markMsgRead([id], true);
      setDataSource(dataSource.filter(item => item.id != id));
      setUnreadNum(unreadNum - 1);
    }

  useEffect(() => {
    requestUnreadMsg();
  }, [])

  useEffect(() => {
    if (showMsg) {
      api.destroy();
      const maxLength = dataSource.length > 3 ? 3 : dataSource.length;
      for (let i = 0; i < maxLength; i++) {
        const item = dataSource[i];
        api.open({
          key: item.id,
          duration: null,
          message: item.subject,
          description: <div>
            {item.body}
            <div>{getDate(item.createdAt, 'YYYY-MM-DD HH:mm:ss')}</div>
          </div>,
          onClose: () => {
            setRead(item.id);
          },
        })
      }
    } else {
      api.destroy();
    }
  }, [dataSource, showMsg])


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
