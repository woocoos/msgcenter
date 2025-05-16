import Editor, { loader } from '@monaco-editor/react';

const ICE_MONACO_CDN_HOST = process.env.ICE_MONACO_CDN_HOST

if (ICE_MONACO_CDN_HOST) {
  loader.config({ paths: { vs: ICE_MONACO_CDN_HOST } })
}

export default Editor;
