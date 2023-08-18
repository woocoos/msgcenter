import { getFiles, getFilesRaw, updateFiles } from "@/services/files";
import { UploadOutlined } from "@ant-design/icons";
import { useToken } from "@ant-design/pro-components";
import { Button, Modal, Space, Typography, Upload, message } from "antd"
import { RcFile, UploadFile } from "antd/es/upload";
import { ReactNode, useEffect, useRef, useState } from "react";
import { useTranslation } from "react-i18next";

export default (props: {
  directory?: string;
  value?: string[];
  maxSize?: number;
  accept?: string;
  children?: ReactNode;
  onChange?: (values?: string[]) => void;
}) => {
  const { t } = useTranslation(),
    { token } = useToken(),
    [list, setList] = useState<UploadFile[]>([]);

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
    isAccept = (file: RcFile) => {
      let isTrue = true;
      if (props.accept) {
        isTrue = props.accept.split(',').includes(`.${file.name.split('.').pop()}`)
      }
      return isTrue
    },
    beforeUpload = async (file: RcFile) => {
      let isTrue = true;
      const maxSize = props.maxSize || 1024 * 5000
      if (!isAccept(file)) {
        isTrue = false;
        message.error(t('file_accept_err_{{accept}}', { accept: file.name.split('.').pop() }));
      }
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
      const
        suffix = file.name.split('.').pop(),
        directory = props.directory || 'temp',
        key = `${directory} / ${Math.floor(Math.floor(Math.random() * 100000) + Date.now()).toString(16)}.${suffix}`;

      const result = await updateFiles({
        key,
        bucket: "adminx-msg",
        file: file,
      })
      if (result) {
        const values = props.value || [];
        values.push(result);
        props.onChange?.(values)
      }
    },
    getValuesFile = async () => {
      if (props.value) {
        const oldKye = list.map(item => item.uid),
          addKeys = props.value.filter(key => !oldKye.includes(key)),
          addLength = addKeys.length;
        for (let i = 0; i < addLength; i++) {
          const data: UploadFile = {
            uid: addKeys[i],
            name: '',
          };
          const result = await getFiles(data.uid);
          if (result?.id) {
            data.name = result.name;
            data.linkProps = {
              download: result.name
            };
          }
          const resultRaw = await getFilesRaw(data.uid, 'url')
          if (typeof resultRaw === 'string') {
            data.url = resultRaw
          }
          list.push(data);
        }
        setList([...list])
      }
    }


  useEffect(() => {
    getValuesFile();
  }, [props.value])

  return <Upload.Dragger
    accept={props.accept}
    beforeUpload={beforeUpload}
    onDrop={(e) => {
      const files = e.dataTransfer.files, l = files.length
      for (let i = 0; i < l; i++) {
        beforeUpload(files[i] as RcFile)
      }
    }}
    multiple={true}
    fileList={list}
  >
    <br />
    <div>
      <UploadOutlined style={{ color: token.colorLink, fontSize: 40 }} />
    </div>
    <br />
    <div>
      <Typography.Text type="secondary">
        {t('click_drag_upload')}
      </Typography.Text>
    </div>
    <div>
      <Typography.Text type="secondary">
        {t('upload_file_accept_{{accept}}', { accept: props.accept?.split(',.').join('、').replace('.', '') })}
      </Typography.Text>
    </div>
  </Upload.Dragger>
}
