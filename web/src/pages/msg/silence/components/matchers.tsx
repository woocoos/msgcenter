import { MatcherInput } from "@/generated/msgsrv/graphql";
import { CloseOutlined, PlusOutlined } from "@ant-design/icons";
import { Button, Popconfirm, Space, Tag } from "antd";
import { useState } from "react";
import { useTranslation } from "react-i18next";
import MatchersForm from "./matchersForm";
import { EnumSilenceMatchType } from "@/services/msgsrv/silence";

export default (props: {
  value?: MatcherInput[];
  onChange?: (value?: MatcherInput[]) => void;
}) => {
  const { t } = useTranslation(),
    [modal, setModal] = useState<{
      open: boolean,
      title: string,
      data?: MatcherInput
    }>({
      open: false,
      title: '',
    });

  const
    toMatcherString = (value: MatcherInput) => {
      return `${value.name}${EnumSilenceMatchType[value.type].text}"${value.value}"`
    },
    onClose = (value: MatcherInput) => {
      const values = props.value?.filter(item => !(item.name == value.name && item.type == value.type && item.value == value.value))
      props.onChange?.(values)
    }

  return <>
    <Space>
      {props.value?.map((item, index) => {
        return <Tag
          key={`tag-${index}`}
          closable
          closeIcon={
            <Popconfirm
              title={t('delete')}
              description={`${t('confirm_delete')}：${toMatcherString(item)}`}
              onConfirm={() => {
                onClose(item);
              }}
            >
              <CloseOutlined />
            </Popconfirm>
          }
          onClose={(e) => {
            e.preventDefault()
          }}
        >
          <a onClick={() => {
            setModal({ open: true, title: `${t(`amend`)}:${toMatcherString(item)}`, data: { ...item } })
          }}>
            {toMatcherString(item)}
          </a>
        </Tag>
      })}

      <Button size="small" type="primary" onClick={() => {
        setModal({ open: true, title: t('add') })
      }}>
        <PlusOutlined />
      </Button>
    </Space >
    <MatchersForm
      open={modal.open}
      title={modal.title}
      data={modal.data}
      onClose={(data) => {
        if (data) {
          const values = [...props.value || []]
          values.push(data)
          props.onChange?.(values)
        }
        setModal({ open: false, title: modal.title })
      }} />
  </>
}
