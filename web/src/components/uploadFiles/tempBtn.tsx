import { getFiles, getFilesRaw, updateFiles } from "@/services/files";
import store from "@/store";
import { UploadOutlined } from "@ant-design/icons";
import { Button, Modal, Space, Typography, Upload, message } from "antd"
import { RcFile } from "antd/es/upload";
import { ReactNode, useEffect, useRef, useState } from "react";
import { useTranslation } from "react-i18next";

const ICE_APP_CODE = process.env.ICE_APP_CODE ?? '';

export default (props: {
  bucket?: string;
  appCode?: string;
  tid?: string;
  /**
   * 目录格式  xxx/ss
   */
  directory?: string;
  /**
   * 强制使用目录当前缀
   */
  forceDirectory?: boolean;
  value?: string;
  maxSize?: number;
  accept?: string;
  children?: ReactNode;
  onChange?: (value?: string) => void;
  onChangePath?: (path?: string) => void;
}) => {
  const { t } = useTranslation(),
    iframeRef = useRef<HTMLIFrameElement>(null),
    [userState] = store.useModel('user'),
    [loading, setLoading] = useState(false),
    [modal, setModal] = useState<{
      show: boolean,
    }>({
      show: false,
    }),
    [name, setName] = useState<string>(),
    [urlSrc, setUrlSrc] = useState<string>();

  const
    //字节大小处理
    formatFileSize = (fileSize: number) => {
      if (fileSize < 1024) {
        return fileSize + 'B';
      } else if (fileSize < (1024 * 1024)) {
        return (fileSize / 1024).toFixed(2) + 'KB';
      } else if (fileSize < (1024 * 1024 * 1024)) {
        return (fileSize / (1024 * 1024)).toFixed(2) + 'MB';
      } else {
        return (fileSize / (1024 * 1024 * 1024)).toFixed(2) + 'GB';
      }
    },
    // 随机数
    randomId = (len: number) => {
      let str = '';
      for (; str.length < len; str += Math.random().toString(36).substring(2));
      return str.substring(0, len);
    },
    beforeUpload = async (file: RcFile) => {
      let isTrue = true;
      const maxSize = props.maxSize || 1024 * 5000

      if (file.size > maxSize) {
        isTrue = false
        message.error(t('file_size_<_{{str}}', { str: formatFileSize(maxSize) }));
      }
      if (isTrue) {
        await updateFile(file)
      }
      return false
    },
    updateFile = async (file: RcFile) => {
      const suffix = file.name.split('.').pop(),
        bucket = props.bucket ?? 'local',
        tid = props.tid ?? userState.tenantId,
        appCode = props.appCode ?? ICE_APP_CODE,
        keys: string[] = [];

      if (props.forceDirectory && props.directory) {
        keys.push(props.directory)
      } else {
        if (appCode) {
          keys.push(appCode)
        }
        if (tid) {
          keys.push(tid)
        }
        if (props.directory) {
          keys.push(props.directory)
        }
      }

      keys.push(`${randomId(16)}.${suffix}`)

      const key = `/${keys.join('/')}`.replace('//', '/');
      setLoading(true)
      if (bucket === 'local') {
        const result = await updateFiles({
          key,
          bucket,
          file,
        })
        if (result) {
          props.onChange?.(result);
          props.onChangePath?.(key);
        }
      }
      setLoading(false)
    },
    getFile = async () => {
      if (props.value) {
        const bucket = props.bucket ?? 'local';
        const result = await getFiles(props.value);
        if (result?.id) {
          setName(result.name)
        }
        if (bucket === 'local') {
          const resultRaw = await getFilesRaw(props.value, 'url')
          if (typeof resultRaw === 'string') {
            setUrlSrc(resultRaw)
          }
        }
      }
    },
    updateContent = async () => {
      if (props.value) {
        const result = await getFilesRaw(props.value)
        if (result && typeof result != 'string') {
          const r = new FileReader()
          r.readAsText(result, 'utf-8')
          r.onload = () => {
            if (iframeRef.current?.contentWindow) {
              iframeRef.current.contentWindow.document.write(r.result as string)
            } else if (iframeRef.current?.contentDocument) {
              iframeRef.current.contentDocument.write(r.result as string)
            }
          }
        }
      }
    }

  useEffect(() => {
    getFile()
  }, [props.value])

  useEffect(() => {
    if (modal.show) {
      updateContent()
    }
  }, [modal.show])

  return <>
    <Space>
      <Upload
        accept={props.accept}
        showUploadList={false}
        beforeUpload={beforeUpload}
      >
        <Button loading={loading} icon={<UploadOutlined />}>{t('upload_file')}</Button>
      </Upload>
      <Typography.Text type="secondary">
        {t('upload_file_accept_{{accept}}', { accept: props.accept?.split(',.').join('、').replace('.', '') })}
      </Typography.Text>
    </Space>
    {props.value ? <>
      <div style={{ height: "8px" }}></div>
      <Space >
        <a onClick={() => {
          setModal({ show: true })
        }}>{t('temp_viewer')}</a>
        <a target="_blank" href={urlSrc} download={name}>{t('temp_down')}</a>
        <Modal
          title={t('temp_viewer')}
          open={modal.show}
          destroyOnClose
          footer={null}
          width={800}
          onCancel={() => {
            setModal({ show: false })
          }}
        >
          <iframe className="modalIframe" ref={iframeRef}></iframe>
        </Modal>
      </Space>
    </> : <></>}
  </>
}
