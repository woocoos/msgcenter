import { request } from 'ice';

type Files = {
  id: string;
  name: string;
  size: number;
  createdAt: string
}

const baseURL = '/api-files'

/**
 * 上传
 * @param data
 * @returns
 */
export async function updateFiles(data: {
  key: string;
  bucket: string;
  file: File;
}) {
  const result = await request.post(`${baseURL}/files`, data, {
    headers: {
      "Content-Type": "multipart/form-data"
    }
  });
  if (typeof result === 'string') {
    return result
  }
  return null
}


/**
 * 获取数据
 * @param fileId
 * @returns
 */
export async function getFiles(fileId: string) {
  const result = await request.get(`${baseURL}/files/${fileId}`)
  if (result?.id) {
    return result as Files;
  }
  return null;
}

/**
 * 删除数据
 * @param fileId
 * @returns
 */
export async function delFiles(fileId: string) {
  const result = await request.delete(`${baseURL}/files/${fileId}`)
  return result;
}



/**
 * 获取数据
 * @param fileId
 * @returns
 */
export async function getFilesRaw(fileId: string) {
  const result = await request.get(`${baseURL}/files/${fileId}/raw`, {
    responseType: "blob",
  })
  if (typeof result === 'object' && result.constructor.name === 'Blob') {
    return URL.createObjectURL(result);
  }
  return null;
}



/**
 * 帮助数组快速获取数据
 * @param arr
 * @param key
 * @returns
 */
export async function formatArrayFilesRaw<T>(arr: T[], key: string, defaultValue?: any) {
  return await Promise.all(arr.map(async (item) => {
    if (item[key]) {
      item[key] = await getFilesRaw(item[key])
    } else {
      item[key] = defaultValue
    }
    return item
  }));
}


