import store from '@/store';
import { SortOrder } from 'antd/lib/table/interface';
import { ReactNode } from 'react';
import { makeOperation, mapExchange } from 'urql';
import i18n from '@/i18n';
import { message } from 'antd';
import { refreshToken } from './auth';
import { goLogin } from '@/util';

export type TreeDataState<T> = {
  key: string;
  title: string | ReactNode;
  children?: TreeDataState<T>[];
  parentId: string | number;
  node?: T;
};

export interface TableParams {
  pageSize?: number;
  current?: number;
  keyword?: string;
  [field: string]: any;
}

export type TableFilter = Record<string, (string | number)[] | null>;

export type TableSort = Record<string, SortOrder>;

export type TreeEditorAction = 'editor' | 'peer' | 'child';

export const urglMapExchange = mapExchange({
  onOperation(operation) {
    const userState = store.getModelState('user'), headers: Record<string, any> = {};
    if (operation.context.fetchOptions?.['headers']?.['Authorization']) {
      headers['Authorization'] = operation.context.fetchOptions?.['headers']?.['Authorization'];
    } else if (userState.token) {
      headers['Authorization'] = `Bearer ${userState.token}`;
    }
    if (operation.context.fetchOptions?.['headers']?.['X-Tenant-ID']) {
      headers['X-Tenant-ID'] = operation.context.fetchOptions?.['headers']?.['X-Tenant-ID'];
    } else if (userState.tenantId) {
      headers['X-Tenant-ID'] = userState.tenantId;
    }

    return makeOperation(operation.kind, operation, {
      fetchOptions: {
        headers,
      },
    });
  },
  onResult(result) {
    if (result.data) {
      refreshToken()
    }
    return result;
  },
  onError: (error) => {
    let msg = '';
    switch (error.response.status) {
      case 401:
        store.dispatch.user.logout();
        goLogin();
        msg = i18n.t('401');
        break;
      case 403:
        msg = i18n.t('403');
        break;
      case 404:
        msg = i18n.t('404');
        break;
      default:
        msg = error.toString();
    }
    if (msg) {
      message.error(msg);
    }
  },
});
