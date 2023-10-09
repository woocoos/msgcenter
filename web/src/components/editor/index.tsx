import Editor, { loader } from '@monaco-editor/react';

loader.config({ paths: { vs: 'https://qlcdn.oss-cn-shenzhen.aliyuncs.com/cdn/monaco-editor/0.44.0/vs' } })

export default Editor;
