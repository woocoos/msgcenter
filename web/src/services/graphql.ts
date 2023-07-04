import { SortOrder } from 'antd/lib/table/interface';
import { ReactNode } from 'react';

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
