import { MsgInternalTo } from "@/generated/msgsrv/graphql";
import { getMsgInternalToInfo, getUserMsgInternalList, markMsgRead } from "@/services/msgsrv/internal";
import { getDate } from "@/util";
import { ProCard } from "@ant-design/pro-components";
import { useSearchParams } from "@ice/runtime";
import { Divider, Empty, Typography } from "antd";
import { useEffect, useRef, useState } from "react";
import { useTranslation } from "react-i18next";
import styles from "./index.module.css";

const Detail = ({ id: propId, onRead }: { id?: string; onRead?: () => void }) => {
  const { t } = useTranslation(),
    iframeRef = useRef<HTMLIFrameElement>(null),
    [searchParams] = useSearchParams(),
    [loading, setLoading] = useState(false),
    [info, setInfo] = useState<MsgInternalTo>();

  const id = propId || searchParams.get('toid') || searchParams.get('id') || '';

  const requestToData = async (id: string) => {
    setLoading(true)
    const result = await getMsgInternalToInfo(id);
    if (result) {
      setInfo(result as MsgInternalTo);
    }
    setLoading(false)
  }, requestData = async (id: string) => {
    setLoading(true)
    const result = await getUserMsgInternalList({
      pageSize: 2,
      where: {
        hasMsgInternalWith: [{
          id,
        }]
      }
    });
    if (result?.totalCount) {
      setInfo(result.edges?.[0]?.node as MsgInternalTo);
    }
    setLoading(false)
  }

  useEffect(() => {
    if (!id) return;
    // 优先按 toId 查询
    const isToId = propId || searchParams.get('toid');
    if (isToId) {
      requestToData(id);
    } else {
      requestData(id);
    }
  }, [id])

  useEffect(() => {
    if (iframeRef.current && info) {
      const iframeDoc = iframeRef.current.contentWindow?.document || iframeRef.current.contentDocument;
      if (iframeDoc) {
        const body = info.msgInternal.body ?? '';
        iframeDoc.write(info.msgInternal.format === 'html' ? body : `<pre style="white-space: pre-wrap;word-wrap: break-word;">${body}</pre>`);
        iframeRef.current.style.height = `${(iframeDoc.body.scrollHeight ?? 0) + 60}px`
      }
      if (!info.readAt) {
        markMsgRead([info.id], true);
        onRead?.();
      }
    }
  }, [info])

  return (
    <ProCard loading={loading}>{
      info ? <div>
        <div style={{ textAlign: 'center' }}>
          <Typography.Title level={4} >{info.msgInternal.subject}</Typography.Title>
          <Typography.Text >{getDate(info.createdAt, 'YYYY-MM-DD HH:mm:ss')}</Typography.Text>
        </div>
        <Divider />
        <iframe ref={iframeRef} className={styles.iframe} />
      </div> : <Empty image={Empty.PRESENTED_IMAGE_SIMPLE} />
    }</ProCard>
  );
};

export default Detail;
