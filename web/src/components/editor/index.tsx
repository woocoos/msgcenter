import Editor, { loader } from '@monaco-editor/react';

loader.config({ paths: { vs: 'https://jsd.onmicrosoft.cn/npm/monaco-editor@0.44.0/min/vs' } })

export default Editor;
