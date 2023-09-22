import { DeleteOutlined } from "@ant-design/icons";
import { Button, Col, Input, Row } from "antd";
import { useEffect, useState } from "react";
import { useTranslation } from "react-i18next";

export interface InputStringRecordProps {
  value?: Record<string, string>;
  titles?: ['key', 'value'];
  keys?: string[];
  disabled?: boolean;
  onChange?: (value?: Record<string, string>) => void;
}

export default (props: InputStringRecordProps) => {
  const { t } = useTranslation(),
    [list, setList] = useState<{ key: string, value: string }[]>([]);

  const updateValue = (data: { key: string, value: string }[]) => {
    if (data.length) {
      const result: Record<string, string> = {};
      data.forEach(item => {
        if (item.key) {
          result[item.key] = item.value;
        }
      })
      props.onChange?.(result);
    } else {
      props.onChange?.();
    }
  }

  useEffect(() => {
    if (props.keys && !props.value) {
      setList(props.keys.map(key => ({
        key,
        value: '',
      })))
    }
  }, [props.keys])

  return <>
    <Row gutter={20}>
      <Col span={11}>
        {props.titles?.[0] ?? 'key'}
      </Col>
      <Col span={11}>
        {props.titles?.[1] ?? 'value'}
      </Col>
    </Row>
    <div style={{ height: '10px' }} />
    {
      list.map((item, index) => <Row gutter={20} style={{ marginBottom: "10px" }} key={index}>
        <Col span={11}>
          <Input
            value={item.key}
            disabled={props.disabled}
            placeholder={`${t('please_enter')}`}
            onChange={(e) => {
              const newList = [...list];
              newList[index].key = e.target.value
              setList(newList);
              updateValue(newList);
            }}
          />
        </Col>
        <Col span={11}>
          <Input
            value={item.value}
            placeholder={`${t('please_enter')}`}
            disabled={props.disabled}
            onChange={(e) => {
              const newList = [...list];
              newList[index].value = e.target.value;
              setList(newList);
              updateValue(newList);
            }}
          />
        </Col>
        <Col span={2}>
          {props.disabled ? <></> : <DeleteOutlined
            onClick={() => {
              const newList = list.filter((_item, idx) => index != idx)
              setList(newList);
              updateValue(newList);
            }}
          />
          }
        </Col>
      </Row>)
    }
    {
      !props.disabled ? <Row>
        <Col span={22}>
          <Button
            block
            type="dashed"
            onClick={() => {
              setList([...list, { key: '', value: '' }]);
            }}
          >+ {t('add')}</Button>
        </Col>
      </Row> : <></>
    }
  </>
}
