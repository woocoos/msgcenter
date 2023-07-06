import { getFiles, getFilesRaw, updateFiles } from "@/services/files";
import { UploadOutlined } from "@ant-design/icons";
import { Button, Modal, Space, Typography, Upload, message } from "antd"
import { RcFile } from "antd/es/upload";
import { ReactNode, useEffect, useRef, useState } from "react";
import { useTranslation } from "react-i18next";

export default (props: {
  directory?: string;
  value?: string;
  maxSize?: number;
  accept?: string;
  children?: ReactNode;
  onChange?: (value?: string) => void;
  onChangeFile?: (value?: RcFile) => void;
}) => {
  const { t } = useTranslation(),
    iframeRef = useRef<HTMLIFrameElement>(null),
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
        directory = props.directory || 'temp',
        key = `${directory}/${Math.floor(Math.floor(Math.random() * 100000) + Date.now()).toString(16)}.${suffix}`
      setLoading(true)
      const result = await updateFiles({
        key,
        bucket: "adminx-msg",
        file: file,
      })
      if (result) {
        props.onChange?.(result)
        props.onChangeFile?.(file)
      }
      setLoading(false)
    },
    getFile = async () => {
      if (props.value) {
        const result = await getFiles(props.value);
        if (result?.id) {
          setName(result.name)
        }
        const resultRaw = await getFilesRaw(props.value, 'url')
        if (typeof resultRaw === 'string') {
          setUrlSrc(resultRaw)
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
