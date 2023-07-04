import { getFilesRaw, updateFiles } from "@/services/files";
import { LoadingOutlined, PlusOutlined } from "@ant-design/icons";
import { Upload, message } from "antd"
import { RcFile } from "antd/es/upload";
import { useEffect, useState } from "react";
import { useTranslation } from "react-i18next";


export default (props: {
  directory?: string;
  value?: string;
  maxSize?: number;
  accept?: string;
  onChange?: (value: string) => void;
}) => {
  const { t } = useTranslation(),
    [loading, setLoading] = useState(false),
    [imgsrc, setImgsrc] = useState<string>();

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
    beforeUpload = (file: RcFile) => {
      let isTrue = true;
      const maxSize = props.maxSize || 1024 * 5000

      if (file.size > maxSize) {
        isTrue = false
        message.error(t('file_size_<_{{str}}', { str: formatFileSize(maxSize) }));
      }
      if (isTrue) {
        updateFile(file)
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
        bucket: "adminx-ui",
        file: file,
      })
      if (result) {
        props.onChange?.(result)
      }
      setLoading(false)
    },
    getFile = async () => {
      if (props.value) {
        const result = await getFilesRaw(props.value)
        if (result) {
          setImgsrc(result)
        }
      }
    }

  useEffect(() => {
    getFile()
  }, [props.value])

  return <Upload
    accept={props.accept}
    listType="picture-card"
    showUploadList={false}
    beforeUpload={beforeUpload}
  >
    {
      imgsrc ? <img
        src={imgsrc} alt="avatar" style={{ width: '100%' }}
      /> : <>
        {loading ? <LoadingOutlined /> : <PlusOutlined />}
        <div style={{ marginTop: 8 }}></div>
      </>
    }
  </Upload>
}
