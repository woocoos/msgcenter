import Editor, { loader } from '@monaco-editor/react';

loader.config({ paths: { vs: 'https://files.qeelyn.com/cdn/monaco-editor/0.52.2/min/vs' } })

export default Editor;
