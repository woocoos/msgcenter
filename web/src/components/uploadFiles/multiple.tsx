import { Files, getFiles, getFilesRaw, updateFiles } from "@/services/files";
import store from "@/store";
import { UploadOutlined } from "@ant-design/icons";
import { useToken } from "@ant-design/pro-components";
import { Modal, Spin, Typography, Upload, message } from "antd"
import { RcFile, UploadFile } from "antd/es/upload";
import { ReactNode, useEffect, useState } from "react";
import { useTranslation } from "react-i18next";

const ICE_APP_CODE = process.env.ICE_APP_CODE ?? '';

let files: RcFile[] = [];
let timeoutFn: NodeJS.Timeout | undefined = undefined;

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
  value?: string[];
  maxSize?: number;
  accept?: string;
  children?: ReactNode;
  onChange?: (values?: string[]) => void;
  onChangePath?: (paths?: string[]) => void;
}) => {
  const { t } = useTranslation(),
    { token } = useToken(),
    [userState] = store.useModel('user'),
    [loading, setLoading] = useState(false),
    [list, setList] = useState<UploadFile<Files>[]>([]);

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
        files.push(file);
        clearTimeout(timeoutFn);
        timeoutFn = setTimeout(async () => {
          await updateFile();
        }, 500);
      }
      return false
    },
    updateFile = async () => {
      if (files.length) {
        setLoading(true);
        const values = props.value || [];
        for (let i in files) {
          const file = files[i],
            suffix = file.name.split('.').pop(),
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
          if (bucket === 'local') {
            try {
              const result = await updateFiles({
                key,
                bucket,
                file,
              })
              if (result) {
                values.push(result);
              }
            } catch (error) {

            }
          }
        }
        await getValuesFile(values);
        props.onChange?.(values);
        setLoading(false);
        files = []
      }
    },
    getValuesFile = async (values?: string[]) => {
      if (values && values.length) {
        const bucket = props.bucket ?? 'local';
        const oldKye = list.map(item => item.uid),
          addKeys = values.filter(key => !oldKye.includes(key)),
          addLength = addKeys.length;
        if (addLength) {
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
              data.response = result
            }
            if (bucket === 'local') {
              const resultRaw = await getFilesRaw(data.uid, 'url')
              if (typeof resultRaw === 'string') {
                data.url = resultRaw
              }
            }
            list.push(data);
          }
          setList([...list])
        }
      } else {
        setList([])
      }
    }

  useEffect(() => {
    getValuesFile(props.value);
  }, [props.value]);


  useEffect(() => {
    const paths: string[] = [];
    list.forEach(item => {
      if (item.response) {
        paths.push(item.response.path)
      }
    })
    props.onChangePath?.(paths)
  }, [list]);


  return <Spin spinning={loading} >
    <Upload.Dragger
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
      onRemove={(file) => {
        Modal.confirm({
          title: t('delete'),
          content: `${t('confirm_delete')}:${file.name}`,
          onOk: () => {
            setList(list.filter(item => item.uid != file.uid));
            props.onChange?.(props.value?.filter(item => item != file.uid) || [])
          }
        })
      }}
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
  </Spin>
}
