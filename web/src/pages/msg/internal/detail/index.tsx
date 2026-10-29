import { MsgInternalTo } from "@/generated/msgsrv/graphql";
import { getMsgInternalToInfo, getUserMsgInternalList, markMsgRead } from "@/services/msgsrv/internal";
import { getDate, customTextToHtml } from "@/util";
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
     const iframeEl = iframeRef.current;
    let cleanup: (() => void) | undefined;
    if (iframeEl && info) {
      const iframeDoc = iframeEl.contentWindow?.document || iframeEl.contentDocument;
      if (iframeDoc) {
        const body = info.msgInternal.body ?? '';
        iframeDoc.open();
        iframeDoc.write(info.msgInternal.format === 'html' ? body : `<pre style="white-space: pre-wrap;word-wrap: break-word;">${body}</pre>`);
        iframeDoc.close();
        // iframe 高度跟随内容自适应，封顶 60vh，超出部分由 iframe 内部滚动，避免弹窗超出屏幕
        const resizeHeight = () => {
          // 先把高度归零再测量：iframe 内 body.scrollHeight 恒 >= iframe 视口高度，直接测量会形成自引用
          const prev = iframeEl.style.height;
          iframeEl.style.height = '0px';
          const contentHeight = iframeDoc.body?.scrollHeight ?? 0;
          iframeEl.style.height = contentHeight > 0
            ? `${Math.min(contentHeight + 16, Math.floor(window.innerHeight * 0.6))}px`
            : prev;
        };
        // 通过 rAF 延迟到下一帧执行，避免 ResizeObserver 循环报错
        let rafId = 0;
        const scheduleResize = () => {
          cancelAnimationFrame(rafId);
          rafId = requestAnimationFrame(resizeHeight);
        };
        resizeHeight();
        // 内容中图片等资源加载完成会改变高度，通过 ResizeObserver 自动修正
        const observer = new ResizeObserver(scheduleResize);
        observer.observe(iframeDoc.body);
        cleanup = () => {
          cancelAnimationFrame(rafId);
          observer.disconnect();
        };
      }
    }
    if (info && !info.readAt) {
      markMsgRead([info.id], true);
      onRead?.();
    }
    return cleanup;
  }, [info])

  return (
    <ProCard loading={loading}>{
      info ? <div>
        <div style={{ textAlign: 'center' }}>
          {/* Typography.Title 内部会注入 children，与 dangerouslySetInnerHTML 冲突，故将 HTML 放到内层元素 */}
          <Typography.Title level={4}>
            <span dangerouslySetInnerHTML={{ __html: customTextToHtml(info.msgInternal.subject) }} />
          </Typography.Title>
          <Typography.Text >{getDate(info.createdAt, 'YYYY-MM-DD HH:mm:ss')}</Typography.Text>
        </div>
        <Divider />
        <iframe ref={iframeRef} className={styles.iframe} />
      </div> : <Empty image={Empty.PRESENTED_IMAGE_SIMPLE} />
    }</ProCard>
  );
};

export default Detail;
