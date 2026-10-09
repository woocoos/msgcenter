import { Input, Tag } from "antd";
import { CSSProperties, useEffect, useState } from "react";
import styles from "./multiple.module.css";

export interface InputMultipleProps {
  value?: string;
  decollator: string;
  disabled?: boolean;
  placeholder?: string;
  tagsStyle?: CSSProperties;
  onChange?: (value?: string) => void;
}

export default (props: InputMultipleProps) => {
  const [tags, setTags] = useState<string[]>([]);
  const [value, setValue] = useState<string>();

  useEffect(() => {
    if (props.value) {
      setTags(props.value?.split(props.decollator))
    } else {
      setTags([])
    }
  }, [props.value])

  return <Input
    prefix={<div className={styles.tags} style={props.tagsStyle}>{
      tags.map((item, index) => <Tag
        key={index}
        closable={props.disabled ? false : true}
        onClose={(e) => {
          e.preventDefault();
          props.onChange?.(tags.filter((_tag, idx) => idx != index).join(props.decollator));
        }}
      >{item}</Tag>)
    }</div>}
    value={value}
    placeholder={props.disabled ? '' : props.placeholder}
    onChange={(event) => {
      setValue(event.target.value);
    }} onPressEnter={(event) => {
      const target = event.target as HTMLInputElement;
      if (target.value) {
        tags.push(target.value);
        props.onChange?.(tags.join(props.decollator));
        setValue('');
      }
    }} />
}
