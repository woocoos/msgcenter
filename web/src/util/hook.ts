import { useLocation } from 'ice';
import menuJson from '../components/layout/menu.json'
import { useEffect, useState } from 'react';
import { isInIcestark } from '@ice/stark-app';
import { store } from '@ice/stark-data';
import { AppMenu } from '@knockout-js/api/ucenter';
import { ProColumns } from '@ant-design/pro-components';
import { useAntdColumnResize } from 'react-antd-column-resize';
import { Column, resizeDataType } from 'react-antd-column-resize/dist/useAntdColumnResize/types';

const NODE_ENV = process.env.NODE_ENV ?? ''


interface MenuJsonData {
  path?: string;
  name: string;
  icon?: string;
  children?: MenuJsonData[];
}

export const routeBreadcrumb = () => {
  const location = useLocation(),
    [names, setNames] = useState<string[]>([])

  const getRouteMenu = (ary: MenuJsonData[], path: string, idxStr?: string) => {
    let str: string = ''
    for (let i = 0; i < ary.length; i++) {
      const item = ary[i],
        addIdxStr = idxStr ? `${idxStr}-${i}` : `${i}`
      if (str) {
        break;
      }
      if (item.path && path.includes(item.path)) {
        str = addIdxStr
        break;
      } else if (item.children) {
        str = getRouteMenu(item.children, path, addIdxStr)
      }
    }
    return str
  }

  const getStoreData = (menu: AppMenu[], id?: string, list?: AppMenu[]) => {
    if (!list) {
      list = []
    }
    if (id) {
      const parentData = menu.find(item => item.id == id)
      if (parentData) {
        list.unshift(parentData)
        if (parentData.parentID != '0') {
          getStoreData(menu, parentData.parentID, list)
        }
      }
    } else {
      const curData = menu.find(item => (item?.route && location.pathname.includes(item.route)))
      if (curData) {
        list.unshift(curData)
        if (curData.parentID != '0') {
          getStoreData(menu, curData.parentID, list)
        }
      }
    }
    return list
  }

  useEffect(() => {
    const ns: string[] = []
    if (NODE_ENV === 'development') {
      const idxStr = getRouteMenu(menuJson, location.pathname)
      let data = menuJson
      idxStr.split('-').forEach(idx => {
        if (data[idx]) {
          ns.push(data[idx].name)
          data = data[idx].children
        }
      })
    } else if (isInIcestark()) {
      const userMenu = store.get('userMenu') as AppMenu[]
      const curAppMenu = getStoreData(userMenu)
      curAppMenu.forEach(item => {
        ns.push(item.name)
      })
    }
    setNames(ns)
  }, [location])

  return [names]
}

/**
 * ProTable 列宽拖拽 hook
 * 封装 useAntdColumnResize，自动处理 fixed:right 列（如操作列）的 onHeaderCell 剥离
 *
 * @param setup 返回 columns 的工厂函数，与 useAntdColumnResize 的 setup 一致
 * @param deps  列定义中依赖的外部状态（与 useEffect deps 语义一致）
 *
 * @example
 * const { columns: finalColumns, components, tableWidth } = useResizableProTable(() => ({
 *   columns: [{ title: 'ID', dataIndex: 'id', width: 100 }, ...],
 * }), [someState]);
 *
 * <ProTable scroll={{ x: tableWidth }} components={components} columns={finalColumns} />
 */
export const useResizableProTable = <T extends Record<string, any>>(
  setup: () => {
    columns: ProColumns<T>[];
    minWidth?: number;
    maxWidth?: number;
  },
  deps: any[] = [],
) => {
  // resizeDataType<Column>
  const { resizableColumns, components, tableWidth } = useAntdColumnResize(() => {
    const { columns, minWidth, maxWidth } = setup()
    return {
      columns: columns as Column[],
      minWidth: minWidth ?? 100,
      maxWidth
    }
  }, deps);

  // fixed:right 的列（操作列）不需要拖拽列宽，移除其 onHeaderCell 使 ResizableHeaderCell 走普通 <th> 分支
  const finalColumns = resizableColumns.map((col) => {
    if (col.fixed === 'right' || col.fixed === true) {
      const { onHeaderCell, ...rest } = col;
      return rest;
    }
    return col;
  });

  return {
    columns: finalColumns as ProColumns<T>[],
    components,
    tableWidth: tableWidth - 200,
  };
}
