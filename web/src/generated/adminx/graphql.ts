/* eslint-disable */
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
  /**
   * Define a Relay Cursor type:
   * https://relay.dev/graphql/connections.htm#sec-Cursor
   */
  Cursor: { input: any; output: any; }
  /** An object with a Global ID,for using in Noder interface. */
  GID: { input: any; output: any; }
  /** The builtin Time type */
  Time: { input: any; output: any; }
};

export type App = Node & {
  __typename?: 'App';
  actions: AppActionConnection;
  /** 应用ID */
  appKey?: Maybe<Scalars['String']['output']>;
  /** 应用密钥 */
  appSecret?: Maybe<Scalars['String']['output']>;
  /** 用于标识应用资源的唯一代码,尽量简短 */
  code: Scalars['String']['output'];
  /** 备注 */
  comments?: Maybe<Scalars['String']['output']>;
  createdAt: Scalars['Time']['output'];
  createdBy: Scalars['Int']['output'];
  dicts: AppDictConnection;
  id: Scalars['ID']['output'];
  /** 应用类型 */
  kind: AppKind;
  /** 应用图标地址 */
  logo?: Maybe<Scalars['String']['output']>;
  menus: AppMenuConnection;
  /** 名称 */
  name: Scalars['String']['output'];
  orgs: OrgConnection;
  /** 策略 */
  policies?: Maybe<Array<AppPolicy>>;
  /** 策略视图 */
  policyViews?: Maybe<Array<AppPolicyView>>;
  /** 回调地址 */
  redirectURI?: Maybe<Scalars['String']['output']>;
  /** refresh_token有效期 */
  refreshTokenValidity?: Maybe<Scalars['Int']['output']>;
  resources: AppResConnection;
  /** 角色 */
  roles?: Maybe<Array<AppRole>>;
  /** 权限范围 */
  scopes?: Maybe<Scalars['String']['output']>;
  /** 状态 */
  status?: Maybe<AppSimpleStatus>;
  /** token有效期 */
  tokenValidity?: Maybe<Scalars['Int']['output']>;
  updatedAt?: Maybe<Scalars['Time']['output']>;
  updatedBy?: Maybe<Scalars['Int']['output']>;
};


export type AppActionsArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<AppActionOrder>;
  where?: InputMaybe<AppActionWhereInput>;
};


export type AppDictsArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<AppDictOrder>;
  where?: InputMaybe<AppDictWhereInput>;
};


export type AppMenusArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<AppMenuOrder>;
  where?: InputMaybe<AppMenuWhereInput>;
};


export type AppOrgsArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<OrgOrder>;
  where?: InputMaybe<OrgWhereInput>;
};


export type AppResourcesArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<AppResOrder>;
  where?: InputMaybe<AppResWhereInput>;
};

export type AppAction = Node & {
  __typename?: 'AppAction';
  app?: Maybe<App>;
  /** 所属应用 */
  appID?: Maybe<Scalars['ID']['output']>;
  /** 备注 */
  comments?: Maybe<Scalars['String']['output']>;
  createdAt: Scalars['Time']['output'];
  createdBy: Scalars['Int']['output'];
  id: Scalars['ID']['output'];
  /** restful,graphql,rpc,function */
  kind: AppActionKind;
  /** 被引用的菜单项 */
  menus?: Maybe<Array<AppMenu>>;
  /** 操作方法:读,写,列表 */
  method: AppActionMethod;
  /** 名称 */
  name: Scalars['String']['output'];
  updatedAt?: Maybe<Scalars['Time']['output']>;
  updatedBy?: Maybe<Scalars['Int']['output']>;
};

/** A connection to a list of items. */
export type AppActionConnection = {
  __typename?: 'AppActionConnection';
  /** A list of edges. */
  edges?: Maybe<Array<Maybe<AppActionEdge>>>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** Identifies the total count of items in the connection. */
  totalCount: Scalars['Int']['output'];
};

/** An edge in a connection. */
export type AppActionEdge = {
  __typename?: 'AppActionEdge';
  /** A cursor for use in pagination. */
  cursor: Scalars['Cursor']['output'];
  /** The item at the end of the edge. */
  node?: Maybe<AppAction>;
};

/** AppActionKind is enum for the field kind */
export enum AppActionKind {
  Function = 'function',
  Graphql = 'graphql',
  Restful = 'restful',
  Route = 'route',
  Rpc = 'rpc'
}

/** AppActionMethod is enum for the field method */
export enum AppActionMethod {
  List = 'list',
  Read = 'read',
  Write = 'write'
}

/** Ordering options for AppAction connections */
export type AppActionOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order AppActions. */
  field: AppActionOrderField;
};

/** Properties by which AppAction connections can be ordered. */
export enum AppActionOrderField {
  CreatedAt = 'createdAt'
}

/**
 * AppActionWhereInput is used for filtering AppAction objects.
 * Input was generated by ent.
 */
export type AppActionWhereInput = {
  and?: InputMaybe<Array<AppActionWhereInput>>;
  /** app_id field predicates */
  appID?: InputMaybe<Scalars['ID']['input']>;
  appIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  appIDIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  appIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  appIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  appIDNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** created_at field predicates */
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  createdAtGT?: InputMaybe<Scalars['Time']['input']>;
  createdAtGTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  createdAtLT?: InputMaybe<Scalars['Time']['input']>;
  createdAtLTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  createdAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  /** created_by field predicates */
  createdBy?: InputMaybe<Scalars['Int']['input']>;
  createdByGT?: InputMaybe<Scalars['Int']['input']>;
  createdByGTE?: InputMaybe<Scalars['Int']['input']>;
  createdByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  createdByLT?: InputMaybe<Scalars['Int']['input']>;
  createdByLTE?: InputMaybe<Scalars['Int']['input']>;
  createdByNEQ?: InputMaybe<Scalars['Int']['input']>;
  createdByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  /** app edge predicates */
  hasApp?: InputMaybe<Scalars['Boolean']['input']>;
  hasAppWith?: InputMaybe<Array<AppWhereInput>>;
  /** menus edge predicates */
  hasMenus?: InputMaybe<Scalars['Boolean']['input']>;
  hasMenusWith?: InputMaybe<Array<AppMenuWhereInput>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** kind field predicates */
  kind?: InputMaybe<AppActionKind>;
  kindIn?: InputMaybe<Array<AppActionKind>>;
  kindNEQ?: InputMaybe<AppActionKind>;
  kindNotIn?: InputMaybe<Array<AppActionKind>>;
  /** method field predicates */
  method?: InputMaybe<AppActionMethod>;
  methodIn?: InputMaybe<Array<AppActionMethod>>;
  methodNEQ?: InputMaybe<AppActionMethod>;
  methodNotIn?: InputMaybe<Array<AppActionMethod>>;
  /** name field predicates */
  name?: InputMaybe<Scalars['String']['input']>;
  nameContains?: InputMaybe<Scalars['String']['input']>;
  nameContainsFold?: InputMaybe<Scalars['String']['input']>;
  nameEqualFold?: InputMaybe<Scalars['String']['input']>;
  nameGT?: InputMaybe<Scalars['String']['input']>;
  nameGTE?: InputMaybe<Scalars['String']['input']>;
  nameHasPrefix?: InputMaybe<Scalars['String']['input']>;
  nameHasSuffix?: InputMaybe<Scalars['String']['input']>;
  nameIn?: InputMaybe<Array<Scalars['String']['input']>>;
  nameLT?: InputMaybe<Scalars['String']['input']>;
  nameLTE?: InputMaybe<Scalars['String']['input']>;
  nameNEQ?: InputMaybe<Scalars['String']['input']>;
  nameNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  not?: InputMaybe<AppActionWhereInput>;
  or?: InputMaybe<Array<AppActionWhereInput>>;
  /** updated_at field predicates */
  updatedAt?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedAtLT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtLTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_by field predicates */
  updatedBy?: InputMaybe<Scalars['Int']['input']>;
  updatedByGT?: InputMaybe<Scalars['Int']['input']>;
  updatedByGTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedByLT?: InputMaybe<Scalars['Int']['input']>;
  updatedByLTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByNEQ?: InputMaybe<Scalars['Int']['input']>;
  updatedByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByNotNil?: InputMaybe<Scalars['Boolean']['input']>;
};

/** A connection to a list of items. */
export type AppConnection = {
  __typename?: 'AppConnection';
  /** A list of edges. */
  edges?: Maybe<Array<Maybe<AppEdge>>>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** Identifies the total count of items in the connection. */
  totalCount: Scalars['Int']['output'];
};

export type AppDict = Node & {
  __typename?: 'AppDict';
  app?: Maybe<App>;
  /** 所属应用 */
  appID?: Maybe<Scalars['ID']['output']>;
  /** 用于标识应用资源的唯一代码,尽量简短 */
  code: Scalars['String']['output'];
  /** 备注 */
  comments?: Maybe<Scalars['String']['output']>;
  createdAt: Scalars['Time']['output'];
  createdBy: Scalars['Int']['output'];
  id: Scalars['ID']['output'];
  items?: Maybe<Array<AppDictItem>>;
  /** 名称 */
  name: Scalars['String']['output'];
  updatedAt?: Maybe<Scalars['Time']['output']>;
  updatedBy?: Maybe<Scalars['Int']['output']>;
};

/** A connection to a list of items. */
export type AppDictConnection = {
  __typename?: 'AppDictConnection';
  /** A list of edges. */
  edges?: Maybe<Array<Maybe<AppDictEdge>>>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** Identifies the total count of items in the connection. */
  totalCount: Scalars['Int']['output'];
};

/** An edge in a connection. */
export type AppDictEdge = {
  __typename?: 'AppDictEdge';
  /** A cursor for use in pagination. */
  cursor: Scalars['Cursor']['output'];
  /** The item at the end of the edge. */
  node?: Maybe<AppDict>;
};

export type AppDictItem = Node & {
  __typename?: 'AppDictItem';
  /** 字典值唯一编码,生效后不可修改. */
  code: Scalars['String']['output'];
  /** 备注 */
  comments?: Maybe<Scalars['String']['output']>;
  createdAt: Scalars['Time']['output'];
  createdBy: Scalars['Int']['output'];
  dict?: Maybe<AppDict>;
  /** 所属字典 */
  dictID?: Maybe<Scalars['ID']['output']>;
  displaySort?: Maybe<Scalars['Int']['output']>;
  id: Scalars['ID']['output'];
  /** 名称 */
  name: Scalars['String']['output'];
  org?: Maybe<Org>;
  /** 租户ID,空为全局字典 */
  orgID?: Maybe<Scalars['ID']['output']>;
  /** 关联代码,由app_code和dict_code组成 */
  refCode: Scalars['String']['output'];
  /** 状态 */
  status?: Maybe<AppDictItemSimpleStatus>;
  updatedAt?: Maybe<Scalars['Time']['output']>;
  updatedBy?: Maybe<Scalars['Int']['output']>;
};

/** Ordering options for AppDictItem connections */
export type AppDictItemOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order AppDictItems. */
  field: AppDictItemOrderField;
};

/** Properties by which AppDictItem connections can be ordered. */
export enum AppDictItemOrderField {
  CreatedAt = 'createdAt',
  DisplaySort = 'displaySort'
}

/** AppDictItemSimpleStatus is enum for the field status */
export enum AppDictItemSimpleStatus {
  Active = 'active',
  Disabled = 'disabled',
  Inactive = 'inactive',
  Processing = 'processing'
}

/**
 * AppDictItemWhereInput is used for filtering AppDictItem objects.
 * Input was generated by ent.
 */
export type AppDictItemWhereInput = {
  and?: InputMaybe<Array<AppDictItemWhereInput>>;
  /** code field predicates */
  code?: InputMaybe<Scalars['String']['input']>;
  codeContains?: InputMaybe<Scalars['String']['input']>;
  codeContainsFold?: InputMaybe<Scalars['String']['input']>;
  codeEqualFold?: InputMaybe<Scalars['String']['input']>;
  codeGT?: InputMaybe<Scalars['String']['input']>;
  codeGTE?: InputMaybe<Scalars['String']['input']>;
  codeHasPrefix?: InputMaybe<Scalars['String']['input']>;
  codeHasSuffix?: InputMaybe<Scalars['String']['input']>;
  codeIn?: InputMaybe<Array<Scalars['String']['input']>>;
  codeLT?: InputMaybe<Scalars['String']['input']>;
  codeLTE?: InputMaybe<Scalars['String']['input']>;
  codeNEQ?: InputMaybe<Scalars['String']['input']>;
  codeNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  /** created_at field predicates */
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  createdAtGT?: InputMaybe<Scalars['Time']['input']>;
  createdAtGTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  createdAtLT?: InputMaybe<Scalars['Time']['input']>;
  createdAtLTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  createdAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  /** created_by field predicates */
  createdBy?: InputMaybe<Scalars['Int']['input']>;
  createdByGT?: InputMaybe<Scalars['Int']['input']>;
  createdByGTE?: InputMaybe<Scalars['Int']['input']>;
  createdByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  createdByLT?: InputMaybe<Scalars['Int']['input']>;
  createdByLTE?: InputMaybe<Scalars['Int']['input']>;
  createdByNEQ?: InputMaybe<Scalars['Int']['input']>;
  createdByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  /** dict_id field predicates */
  dictID?: InputMaybe<Scalars['ID']['input']>;
  dictIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  dictIDIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  dictIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  dictIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  dictIDNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** dict edge predicates */
  hasDict?: InputMaybe<Scalars['Boolean']['input']>;
  hasDictWith?: InputMaybe<Array<AppDictWhereInput>>;
  /** org edge predicates */
  hasOrg?: InputMaybe<Scalars['Boolean']['input']>;
  hasOrgWith?: InputMaybe<Array<OrgWhereInput>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** name field predicates */
  name?: InputMaybe<Scalars['String']['input']>;
  nameContains?: InputMaybe<Scalars['String']['input']>;
  nameContainsFold?: InputMaybe<Scalars['String']['input']>;
  nameEqualFold?: InputMaybe<Scalars['String']['input']>;
  nameGT?: InputMaybe<Scalars['String']['input']>;
  nameGTE?: InputMaybe<Scalars['String']['input']>;
  nameHasPrefix?: InputMaybe<Scalars['String']['input']>;
  nameHasSuffix?: InputMaybe<Scalars['String']['input']>;
  nameIn?: InputMaybe<Array<Scalars['String']['input']>>;
  nameLT?: InputMaybe<Scalars['String']['input']>;
  nameLTE?: InputMaybe<Scalars['String']['input']>;
  nameNEQ?: InputMaybe<Scalars['String']['input']>;
  nameNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  not?: InputMaybe<AppDictItemWhereInput>;
  or?: InputMaybe<Array<AppDictItemWhereInput>>;
  /** org_id field predicates */
  orgID?: InputMaybe<Scalars['ID']['input']>;
  orgIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  orgIDIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  orgIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  orgIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  orgIDNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** status field predicates */
  status?: InputMaybe<AppDictItemSimpleStatus>;
  statusIn?: InputMaybe<Array<AppDictItemSimpleStatus>>;
  statusIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  statusNEQ?: InputMaybe<AppDictItemSimpleStatus>;
  statusNotIn?: InputMaybe<Array<AppDictItemSimpleStatus>>;
  statusNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_at field predicates */
  updatedAt?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedAtLT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtLTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_by field predicates */
  updatedBy?: InputMaybe<Scalars['Int']['input']>;
  updatedByGT?: InputMaybe<Scalars['Int']['input']>;
  updatedByGTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedByLT?: InputMaybe<Scalars['Int']['input']>;
  updatedByLTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByNEQ?: InputMaybe<Scalars['Int']['input']>;
  updatedByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByNotNil?: InputMaybe<Scalars['Boolean']['input']>;
};

/** Ordering options for AppDict connections */
export type AppDictOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order AppDicts. */
  field: AppDictOrderField;
};

/** Properties by which AppDict connections can be ordered. */
export enum AppDictOrderField {
  CreatedAt = 'createdAt'
}

/**
 * AppDictWhereInput is used for filtering AppDict objects.
 * Input was generated by ent.
 */
export type AppDictWhereInput = {
  and?: InputMaybe<Array<AppDictWhereInput>>;
  /** app_id field predicates */
  appID?: InputMaybe<Scalars['ID']['input']>;
  appIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  appIDIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  appIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  appIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  appIDNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** code field predicates */
  code?: InputMaybe<Scalars['String']['input']>;
  codeContains?: InputMaybe<Scalars['String']['input']>;
  codeContainsFold?: InputMaybe<Scalars['String']['input']>;
  codeEqualFold?: InputMaybe<Scalars['String']['input']>;
  codeGT?: InputMaybe<Scalars['String']['input']>;
  codeGTE?: InputMaybe<Scalars['String']['input']>;
  codeHasPrefix?: InputMaybe<Scalars['String']['input']>;
  codeHasSuffix?: InputMaybe<Scalars['String']['input']>;
  codeIn?: InputMaybe<Array<Scalars['String']['input']>>;
  codeLT?: InputMaybe<Scalars['String']['input']>;
  codeLTE?: InputMaybe<Scalars['String']['input']>;
  codeNEQ?: InputMaybe<Scalars['String']['input']>;
  codeNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  /** created_at field predicates */
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  createdAtGT?: InputMaybe<Scalars['Time']['input']>;
  createdAtGTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  createdAtLT?: InputMaybe<Scalars['Time']['input']>;
  createdAtLTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  createdAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  /** created_by field predicates */
  createdBy?: InputMaybe<Scalars['Int']['input']>;
  createdByGT?: InputMaybe<Scalars['Int']['input']>;
  createdByGTE?: InputMaybe<Scalars['Int']['input']>;
  createdByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  createdByLT?: InputMaybe<Scalars['Int']['input']>;
  createdByLTE?: InputMaybe<Scalars['Int']['input']>;
  createdByNEQ?: InputMaybe<Scalars['Int']['input']>;
  createdByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  /** app edge predicates */
  hasApp?: InputMaybe<Scalars['Boolean']['input']>;
  hasAppWith?: InputMaybe<Array<AppWhereInput>>;
  /** items edge predicates */
  hasItems?: InputMaybe<Scalars['Boolean']['input']>;
  hasItemsWith?: InputMaybe<Array<AppDictItemWhereInput>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** name field predicates */
  name?: InputMaybe<Scalars['String']['input']>;
  nameContains?: InputMaybe<Scalars['String']['input']>;
  nameContainsFold?: InputMaybe<Scalars['String']['input']>;
  nameEqualFold?: InputMaybe<Scalars['String']['input']>;
  nameGT?: InputMaybe<Scalars['String']['input']>;
  nameGTE?: InputMaybe<Scalars['String']['input']>;
  nameHasPrefix?: InputMaybe<Scalars['String']['input']>;
  nameHasSuffix?: InputMaybe<Scalars['String']['input']>;
  nameIn?: InputMaybe<Array<Scalars['String']['input']>>;
  nameLT?: InputMaybe<Scalars['String']['input']>;
  nameLTE?: InputMaybe<Scalars['String']['input']>;
  nameNEQ?: InputMaybe<Scalars['String']['input']>;
  nameNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  not?: InputMaybe<AppDictWhereInput>;
  or?: InputMaybe<Array<AppDictWhereInput>>;
  /** updated_at field predicates */
  updatedAt?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedAtLT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtLTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_by field predicates */
  updatedBy?: InputMaybe<Scalars['Int']['input']>;
  updatedByGT?: InputMaybe<Scalars['Int']['input']>;
  updatedByGTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedByLT?: InputMaybe<Scalars['Int']['input']>;
  updatedByLTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByNEQ?: InputMaybe<Scalars['Int']['input']>;
  updatedByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByNotNil?: InputMaybe<Scalars['Boolean']['input']>;
};

/** An edge in a connection. */
export type AppEdge = {
  __typename?: 'AppEdge';
  /** A cursor for use in pagination. */
  cursor: Scalars['Cursor']['output'];
  /** The item at the end of the edge. */
  node?: Maybe<App>;
};

/** AppKind is enum for the field kind */
export enum AppKind {
  Native = 'native',
  Server = 'server',
  Web = 'web'
}

export type AppMenu = Node & {
  __typename?: 'AppMenu';
  /** 需要权限控制时对应的权限 */
  action?: Maybe<AppAction>;
  /** 操作ID */
  actionID?: Maybe<Scalars['ID']['output']>;
  app?: Maybe<App>;
  /** 所属应用 */
  appID?: Maybe<Scalars['ID']['output']>;
  children?: Maybe<Array<AppMenu>>;
  /** 备注 */
  comments?: Maybe<Scalars['String']['output']>;
  createdAt: Scalars['Time']['output'];
  createdBy: Scalars['Int']['output'];
  displaySort?: Maybe<Scalars['Int']['output']>;
  /** 菜单图标 */
  icon?: Maybe<Scalars['String']['output']>;
  id: Scalars['ID']['output'];
  /** 目录,菜单项 */
  kind: AppMenuKind;
  /** 菜单名称 */
  name: Scalars['String']['output'];
  parent: AppMenu;
  /** 父级ID */
  parentID: Scalars['ID']['output'];
  /** 菜单路由 */
  route?: Maybe<Scalars['String']['output']>;
  /** 状态 */
  status?: Maybe<AppMenuSimpleStatus>;
  updatedAt?: Maybe<Scalars['Time']['output']>;
  updatedBy?: Maybe<Scalars['Int']['output']>;
};

/** A connection to a list of items. */
export type AppMenuConnection = {
  __typename?: 'AppMenuConnection';
  /** A list of edges. */
  edges?: Maybe<Array<Maybe<AppMenuEdge>>>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** Identifies the total count of items in the connection. */
  totalCount: Scalars['Int']['output'];
};

/** An edge in a connection. */
export type AppMenuEdge = {
  __typename?: 'AppMenuEdge';
  /** A cursor for use in pagination. */
  cursor: Scalars['Cursor']['output'];
  /** The item at the end of the edge. */
  node?: Maybe<AppMenu>;
};

/** AppMenuKind is enum for the field kind */
export enum AppMenuKind {
  Dir = 'dir',
  Menu = 'menu'
}

/** Ordering options for AppMenu connections */
export type AppMenuOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order AppMenus. */
  field: AppMenuOrderField;
};

/** Properties by which AppMenu connections can be ordered. */
export enum AppMenuOrderField {
  CreatedAt = 'createdAt',
  DisplaySort = 'displaySort'
}

/** AppMenuSimpleStatus is enum for the field status */
export enum AppMenuSimpleStatus {
  Active = 'active',
  Disabled = 'disabled',
  Inactive = 'inactive',
  Processing = 'processing'
}

/**
 * AppMenuWhereInput is used for filtering AppMenu objects.
 * Input was generated by ent.
 */
export type AppMenuWhereInput = {
  and?: InputMaybe<Array<AppMenuWhereInput>>;
  /** app_id field predicates */
  appID?: InputMaybe<Scalars['ID']['input']>;
  appIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  appIDIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  appIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  appIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  appIDNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** created_at field predicates */
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  createdAtGT?: InputMaybe<Scalars['Time']['input']>;
  createdAtGTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  createdAtLT?: InputMaybe<Scalars['Time']['input']>;
  createdAtLTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  createdAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  /** created_by field predicates */
  createdBy?: InputMaybe<Scalars['Int']['input']>;
  createdByGT?: InputMaybe<Scalars['Int']['input']>;
  createdByGTE?: InputMaybe<Scalars['Int']['input']>;
  createdByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  createdByLT?: InputMaybe<Scalars['Int']['input']>;
  createdByLTE?: InputMaybe<Scalars['Int']['input']>;
  createdByNEQ?: InputMaybe<Scalars['Int']['input']>;
  createdByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  /** action edge predicates */
  hasAction?: InputMaybe<Scalars['Boolean']['input']>;
  hasActionWith?: InputMaybe<Array<AppActionWhereInput>>;
  /** app edge predicates */
  hasApp?: InputMaybe<Scalars['Boolean']['input']>;
  hasAppWith?: InputMaybe<Array<AppWhereInput>>;
  /** children edge predicates */
  hasChildren?: InputMaybe<Scalars['Boolean']['input']>;
  hasChildrenWith?: InputMaybe<Array<AppMenuWhereInput>>;
  /** parent edge predicates */
  hasParent?: InputMaybe<Scalars['Boolean']['input']>;
  hasParentWith?: InputMaybe<Array<AppMenuWhereInput>>;
  /** icon field predicates */
  icon?: InputMaybe<Scalars['String']['input']>;
  iconContains?: InputMaybe<Scalars['String']['input']>;
  iconContainsFold?: InputMaybe<Scalars['String']['input']>;
  iconEqualFold?: InputMaybe<Scalars['String']['input']>;
  iconGT?: InputMaybe<Scalars['String']['input']>;
  iconGTE?: InputMaybe<Scalars['String']['input']>;
  iconHasPrefix?: InputMaybe<Scalars['String']['input']>;
  iconHasSuffix?: InputMaybe<Scalars['String']['input']>;
  iconIn?: InputMaybe<Array<Scalars['String']['input']>>;
  iconIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  iconLT?: InputMaybe<Scalars['String']['input']>;
  iconLTE?: InputMaybe<Scalars['String']['input']>;
  iconNEQ?: InputMaybe<Scalars['String']['input']>;
  iconNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  iconNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** kind field predicates */
  kind?: InputMaybe<AppMenuKind>;
  kindIn?: InputMaybe<Array<AppMenuKind>>;
  kindNEQ?: InputMaybe<AppMenuKind>;
  kindNotIn?: InputMaybe<Array<AppMenuKind>>;
  /** name field predicates */
  name?: InputMaybe<Scalars['String']['input']>;
  nameContains?: InputMaybe<Scalars['String']['input']>;
  nameContainsFold?: InputMaybe<Scalars['String']['input']>;
  nameEqualFold?: InputMaybe<Scalars['String']['input']>;
  nameGT?: InputMaybe<Scalars['String']['input']>;
  nameGTE?: InputMaybe<Scalars['String']['input']>;
  nameHasPrefix?: InputMaybe<Scalars['String']['input']>;
  nameHasSuffix?: InputMaybe<Scalars['String']['input']>;
  nameIn?: InputMaybe<Array<Scalars['String']['input']>>;
  nameLT?: InputMaybe<Scalars['String']['input']>;
  nameLTE?: InputMaybe<Scalars['String']['input']>;
  nameNEQ?: InputMaybe<Scalars['String']['input']>;
  nameNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  not?: InputMaybe<AppMenuWhereInput>;
  or?: InputMaybe<Array<AppMenuWhereInput>>;
  /** parent_id field predicates */
  parentID?: InputMaybe<Scalars['ID']['input']>;
  parentIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  parentIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  parentIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** route field predicates */
  route?: InputMaybe<Scalars['String']['input']>;
  routeContains?: InputMaybe<Scalars['String']['input']>;
  routeContainsFold?: InputMaybe<Scalars['String']['input']>;
  routeEqualFold?: InputMaybe<Scalars['String']['input']>;
  routeGT?: InputMaybe<Scalars['String']['input']>;
  routeGTE?: InputMaybe<Scalars['String']['input']>;
  routeHasPrefix?: InputMaybe<Scalars['String']['input']>;
  routeHasSuffix?: InputMaybe<Scalars['String']['input']>;
  routeIn?: InputMaybe<Array<Scalars['String']['input']>>;
  routeIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  routeLT?: InputMaybe<Scalars['String']['input']>;
  routeLTE?: InputMaybe<Scalars['String']['input']>;
  routeNEQ?: InputMaybe<Scalars['String']['input']>;
  routeNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  routeNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** status field predicates */
  status?: InputMaybe<AppMenuSimpleStatus>;
  statusIn?: InputMaybe<Array<AppMenuSimpleStatus>>;
  statusIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  statusNEQ?: InputMaybe<AppMenuSimpleStatus>;
  statusNotIn?: InputMaybe<Array<AppMenuSimpleStatus>>;
  statusNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_at field predicates */
  updatedAt?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedAtLT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtLTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_by field predicates */
  updatedBy?: InputMaybe<Scalars['Int']['input']>;
  updatedByGT?: InputMaybe<Scalars['Int']['input']>;
  updatedByGTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedByLT?: InputMaybe<Scalars['Int']['input']>;
  updatedByLTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByNEQ?: InputMaybe<Scalars['Int']['input']>;
  updatedByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByNotNil?: InputMaybe<Scalars['Boolean']['input']>;
};

/** Ordering options for App connections */
export type AppOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order Apps. */
  field: AppOrderField;
};

/** Properties by which App connections can be ordered. */
export enum AppOrderField {
  CreatedAt = 'createdAt'
}

export type AppPolicy = Node & {
  __typename?: 'AppPolicy';
  app?: Maybe<App>;
  /** 所属应用 */
  appID?: Maybe<Scalars['ID']['output']>;
  /** 标识是否自动授予到账户 */
  autoGrant: Scalars['Boolean']['output'];
  /** 描述 */
  comments?: Maybe<Scalars['String']['output']>;
  createdAt: Scalars['Time']['output'];
  createdBy: Scalars['Int']['output'];
  id: Scalars['ID']['output'];
  /** 是否授权role */
  isGrantAppRole: Scalars['Boolean']['output'];
  /** 分类：app-应用策略、view-策略视图 */
  kind: AppPolicyKind;
  /** 策略名称 */
  name: Scalars['String']['output'];
  /** 策略授权的组织策略 */
  orgPolicies?: Maybe<Array<OrgPolicy>>;
  /** 策略视图 */
  policyViews?: Maybe<Array<AppPolicyView>>;
  roles?: Maybe<Array<AppRole>>;
  /** 策略规则 */
  rules: Array<Maybe<PolicyRule>>;
  /** 状态 */
  status?: Maybe<AppPolicySimpleStatus>;
  updatedAt?: Maybe<Scalars['Time']['output']>;
  updatedBy?: Maybe<Scalars['Int']['output']>;
};


export type AppPolicyIsGrantAppRoleArgs = {
  appRoleID: Scalars['ID']['input'];
};

/** A connection to a list of items. */
export type AppPolicyConnection = {
  __typename?: 'AppPolicyConnection';
  /** A list of edges. */
  edges?: Maybe<Array<Maybe<AppPolicyEdge>>>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** Identifies the total count of items in the connection. */
  totalCount: Scalars['Int']['output'];
};

/** An edge in a connection. */
export type AppPolicyEdge = {
  __typename?: 'AppPolicyEdge';
  /** A cursor for use in pagination. */
  cursor: Scalars['Cursor']['output'];
  /** The item at the end of the edge. */
  node?: Maybe<AppPolicy>;
};

/** AppPolicyKind is enum for the field kind */
export enum AppPolicyKind {
  App = 'app',
  View = 'view'
}

/** Ordering options for AppPolicy connections */
export type AppPolicyOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order AppPolicies. */
  field: AppPolicyOrderField;
};

/** Properties by which AppPolicy connections can be ordered. */
export enum AppPolicyOrderField {
  CreatedAt = 'createdAt'
}

/** AppPolicySimpleStatus is enum for the field status */
export enum AppPolicySimpleStatus {
  Active = 'active',
  Disabled = 'disabled',
  Inactive = 'inactive',
  Processing = 'processing'
}

export type AppPolicyView = Node & {
  __typename?: 'AppPolicyView';
  app?: Maybe<App>;
  /** 所属应用 */
  appID?: Maybe<Scalars['ID']['output']>;
  appPolicy?: Maybe<AppPolicy>;
  children?: Maybe<Array<AppPolicyView>>;
  /** 描述 */
  comments?: Maybe<Scalars['String']['output']>;
  createdAt: Scalars['Time']['output'];
  createdBy: Scalars['Int']['output'];
  displaySort?: Maybe<Scalars['Int']['output']>;
  id: Scalars['ID']['output'];
  /** 分类：dir-目录、policy-权限策略 */
  kind: AppPolicyViewKind;
  /** 名称 */
  name: Scalars['String']['output'];
  parent: AppPolicyView;
  /** 父级ID,0为顶级 */
  parentID: Scalars['ID']['output'];
  /** 路径编码 */
  path?: Maybe<Scalars['String']['output']>;
  /** 关联的应用策略 */
  policyID?: Maybe<Scalars['ID']['output']>;
  updatedAt?: Maybe<Scalars['Time']['output']>;
  updatedBy?: Maybe<Scalars['Int']['output']>;
};

/** A connection to a list of items. */
export type AppPolicyViewConnection = {
  __typename?: 'AppPolicyViewConnection';
  /** A list of edges. */
  edges?: Maybe<Array<Maybe<AppPolicyViewEdge>>>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** Identifies the total count of items in the connection. */
  totalCount: Scalars['Int']['output'];
};

/** An edge in a connection. */
export type AppPolicyViewEdge = {
  __typename?: 'AppPolicyViewEdge';
  /** A cursor for use in pagination. */
  cursor: Scalars['Cursor']['output'];
  /** The item at the end of the edge. */
  node?: Maybe<AppPolicyView>;
};

/** AppPolicyViewKind is enum for the field kind */
export enum AppPolicyViewKind {
  Dir = 'dir',
  Policy = 'policy'
}

/** Ordering options for AppPolicyView connections */
export type AppPolicyViewOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order AppPolicyViews. */
  field: AppPolicyViewOrderField;
};

/** Properties by which AppPolicyView connections can be ordered. */
export enum AppPolicyViewOrderField {
  CreatedAt = 'createdAt',
  DisplaySort = 'displaySort'
}

export type AppPolicyViewOrgPolicy = {
  __typename?: 'AppPolicyViewOrgPolicy';
  appPolicyView?: Maybe<AppPolicyView>;
  orgPolicy?: Maybe<OrgPolicy>;
};

/**
 * AppPolicyViewWhereInput is used for filtering AppPolicyView objects.
 * Input was generated by ent.
 */
export type AppPolicyViewWhereInput = {
  and?: InputMaybe<Array<AppPolicyViewWhereInput>>;
  /** app_id field predicates */
  appID?: InputMaybe<Scalars['ID']['input']>;
  appIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  appIDIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  appIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  appIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  appIDNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** created_at field predicates */
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  createdAtGT?: InputMaybe<Scalars['Time']['input']>;
  createdAtGTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  createdAtLT?: InputMaybe<Scalars['Time']['input']>;
  createdAtLTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  createdAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  /** created_by field predicates */
  createdBy?: InputMaybe<Scalars['Int']['input']>;
  createdByGT?: InputMaybe<Scalars['Int']['input']>;
  createdByGTE?: InputMaybe<Scalars['Int']['input']>;
  createdByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  createdByLT?: InputMaybe<Scalars['Int']['input']>;
  createdByLTE?: InputMaybe<Scalars['Int']['input']>;
  createdByNEQ?: InputMaybe<Scalars['Int']['input']>;
  createdByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  /** app edge predicates */
  hasApp?: InputMaybe<Scalars['Boolean']['input']>;
  /** app_policy edge predicates */
  hasAppPolicy?: InputMaybe<Scalars['Boolean']['input']>;
  hasAppPolicyWith?: InputMaybe<Array<AppPolicyWhereInput>>;
  hasAppWith?: InputMaybe<Array<AppWhereInput>>;
  /** children edge predicates */
  hasChildren?: InputMaybe<Scalars['Boolean']['input']>;
  hasChildrenWith?: InputMaybe<Array<AppPolicyViewWhereInput>>;
  /** parent edge predicates */
  hasParent?: InputMaybe<Scalars['Boolean']['input']>;
  hasParentWith?: InputMaybe<Array<AppPolicyViewWhereInput>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** kind field predicates */
  kind?: InputMaybe<AppPolicyViewKind>;
  kindIn?: InputMaybe<Array<AppPolicyViewKind>>;
  kindNEQ?: InputMaybe<AppPolicyViewKind>;
  kindNotIn?: InputMaybe<Array<AppPolicyViewKind>>;
  /** name field predicates */
  name?: InputMaybe<Scalars['String']['input']>;
  nameContains?: InputMaybe<Scalars['String']['input']>;
  nameContainsFold?: InputMaybe<Scalars['String']['input']>;
  nameEqualFold?: InputMaybe<Scalars['String']['input']>;
  nameGT?: InputMaybe<Scalars['String']['input']>;
  nameGTE?: InputMaybe<Scalars['String']['input']>;
  nameHasPrefix?: InputMaybe<Scalars['String']['input']>;
  nameHasSuffix?: InputMaybe<Scalars['String']['input']>;
  nameIn?: InputMaybe<Array<Scalars['String']['input']>>;
  nameLT?: InputMaybe<Scalars['String']['input']>;
  nameLTE?: InputMaybe<Scalars['String']['input']>;
  nameNEQ?: InputMaybe<Scalars['String']['input']>;
  nameNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  not?: InputMaybe<AppPolicyViewWhereInput>;
  or?: InputMaybe<Array<AppPolicyViewWhereInput>>;
  /** parent_id field predicates */
  parentID?: InputMaybe<Scalars['ID']['input']>;
  parentIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  parentIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  parentIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** path field predicates */
  path?: InputMaybe<Scalars['String']['input']>;
  pathContains?: InputMaybe<Scalars['String']['input']>;
  pathContainsFold?: InputMaybe<Scalars['String']['input']>;
  pathEqualFold?: InputMaybe<Scalars['String']['input']>;
  pathGT?: InputMaybe<Scalars['String']['input']>;
  pathGTE?: InputMaybe<Scalars['String']['input']>;
  pathHasPrefix?: InputMaybe<Scalars['String']['input']>;
  pathHasSuffix?: InputMaybe<Scalars['String']['input']>;
  pathIn?: InputMaybe<Array<Scalars['String']['input']>>;
  pathIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  pathLT?: InputMaybe<Scalars['String']['input']>;
  pathLTE?: InputMaybe<Scalars['String']['input']>;
  pathNEQ?: InputMaybe<Scalars['String']['input']>;
  pathNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  pathNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** policy_id field predicates */
  policyID?: InputMaybe<Scalars['ID']['input']>;
  policyIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  policyIDIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  policyIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  policyIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  policyIDNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_at field predicates */
  updatedAt?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedAtLT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtLTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_by field predicates */
  updatedBy?: InputMaybe<Scalars['Int']['input']>;
  updatedByGT?: InputMaybe<Scalars['Int']['input']>;
  updatedByGTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedByLT?: InputMaybe<Scalars['Int']['input']>;
  updatedByLTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByNEQ?: InputMaybe<Scalars['Int']['input']>;
  updatedByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByNotNil?: InputMaybe<Scalars['Boolean']['input']>;
};

/**
 * AppPolicyWhereInput is used for filtering AppPolicy objects.
 * Input was generated by ent.
 */
export type AppPolicyWhereInput = {
  and?: InputMaybe<Array<AppPolicyWhereInput>>;
  /** app_id field predicates */
  appID?: InputMaybe<Scalars['ID']['input']>;
  appIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  appIDIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  appIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  appIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  appIDNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** auto_grant field predicates */
  autoGrant?: InputMaybe<Scalars['Boolean']['input']>;
  autoGrantNEQ?: InputMaybe<Scalars['Boolean']['input']>;
  /** comments field predicates */
  comments?: InputMaybe<Scalars['String']['input']>;
  commentsContains?: InputMaybe<Scalars['String']['input']>;
  commentsContainsFold?: InputMaybe<Scalars['String']['input']>;
  commentsEqualFold?: InputMaybe<Scalars['String']['input']>;
  commentsGT?: InputMaybe<Scalars['String']['input']>;
  commentsGTE?: InputMaybe<Scalars['String']['input']>;
  commentsHasPrefix?: InputMaybe<Scalars['String']['input']>;
  commentsHasSuffix?: InputMaybe<Scalars['String']['input']>;
  commentsIn?: InputMaybe<Array<Scalars['String']['input']>>;
  commentsIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  commentsLT?: InputMaybe<Scalars['String']['input']>;
  commentsLTE?: InputMaybe<Scalars['String']['input']>;
  commentsNEQ?: InputMaybe<Scalars['String']['input']>;
  commentsNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  commentsNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** created_at field predicates */
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  createdAtGT?: InputMaybe<Scalars['Time']['input']>;
  createdAtGTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  createdAtLT?: InputMaybe<Scalars['Time']['input']>;
  createdAtLTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  createdAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  /** created_by field predicates */
  createdBy?: InputMaybe<Scalars['Int']['input']>;
  createdByGT?: InputMaybe<Scalars['Int']['input']>;
  createdByGTE?: InputMaybe<Scalars['Int']['input']>;
  createdByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  createdByLT?: InputMaybe<Scalars['Int']['input']>;
  createdByLTE?: InputMaybe<Scalars['Int']['input']>;
  createdByNEQ?: InputMaybe<Scalars['Int']['input']>;
  createdByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  /** app edge predicates */
  hasApp?: InputMaybe<Scalars['Boolean']['input']>;
  /** app_role_policy edge predicates */
  hasAppRolePolicy?: InputMaybe<Scalars['Boolean']['input']>;
  hasAppRolePolicyWith?: InputMaybe<Array<AppRolePolicyWhereInput>>;
  hasAppWith?: InputMaybe<Array<AppWhereInput>>;
  /** org_policies edge predicates */
  hasOrgPolicies?: InputMaybe<Scalars['Boolean']['input']>;
  hasOrgPoliciesWith?: InputMaybe<Array<OrgPolicyWhereInput>>;
  /** policy_views edge predicates */
  hasPolicyViews?: InputMaybe<Scalars['Boolean']['input']>;
  hasPolicyViewsWith?: InputMaybe<Array<AppPolicyViewWhereInput>>;
  /** roles edge predicates */
  hasRoles?: InputMaybe<Scalars['Boolean']['input']>;
  hasRolesWith?: InputMaybe<Array<AppRoleWhereInput>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** kind field predicates */
  kind?: InputMaybe<AppPolicyKind>;
  kindIn?: InputMaybe<Array<AppPolicyKind>>;
  kindNEQ?: InputMaybe<AppPolicyKind>;
  kindNotIn?: InputMaybe<Array<AppPolicyKind>>;
  /** name field predicates */
  name?: InputMaybe<Scalars['String']['input']>;
  nameContains?: InputMaybe<Scalars['String']['input']>;
  nameContainsFold?: InputMaybe<Scalars['String']['input']>;
  nameEqualFold?: InputMaybe<Scalars['String']['input']>;
  nameGT?: InputMaybe<Scalars['String']['input']>;
  nameGTE?: InputMaybe<Scalars['String']['input']>;
  nameHasPrefix?: InputMaybe<Scalars['String']['input']>;
  nameHasSuffix?: InputMaybe<Scalars['String']['input']>;
  nameIn?: InputMaybe<Array<Scalars['String']['input']>>;
  nameLT?: InputMaybe<Scalars['String']['input']>;
  nameLTE?: InputMaybe<Scalars['String']['input']>;
  nameNEQ?: InputMaybe<Scalars['String']['input']>;
  nameNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  not?: InputMaybe<AppPolicyWhereInput>;
  or?: InputMaybe<Array<AppPolicyWhereInput>>;
  /** status field predicates */
  status?: InputMaybe<AppPolicySimpleStatus>;
  statusIn?: InputMaybe<Array<AppPolicySimpleStatus>>;
  statusIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  statusNEQ?: InputMaybe<AppPolicySimpleStatus>;
  statusNotIn?: InputMaybe<Array<AppPolicySimpleStatus>>;
  statusNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_at field predicates */
  updatedAt?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedAtLT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtLTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_by field predicates */
  updatedBy?: InputMaybe<Scalars['Int']['input']>;
  updatedByGT?: InputMaybe<Scalars['Int']['input']>;
  updatedByGTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedByLT?: InputMaybe<Scalars['Int']['input']>;
  updatedByLTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByNEQ?: InputMaybe<Scalars['Int']['input']>;
  updatedByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByNotNil?: InputMaybe<Scalars['Boolean']['input']>;
};

export type AppRes = Node & {
  __typename?: 'AppRes';
  app?: Maybe<App>;
  /** 所属应用 */
  appID?: Maybe<Scalars['ID']['output']>;
  /** 应用资源表达式 */
  arnPattern: Scalars['String']['output'];
  createdAt: Scalars['Time']['output'];
  createdBy: Scalars['Int']['output'];
  id: Scalars['ID']['output'];
  /** 资源名称 */
  name: Scalars['String']['output'];
  /** 资源类型名称,如数据库表名 */
  typeName: Scalars['String']['output'];
  updatedAt?: Maybe<Scalars['Time']['output']>;
  updatedBy?: Maybe<Scalars['Int']['output']>;
};

/** A connection to a list of items. */
export type AppResConnection = {
  __typename?: 'AppResConnection';
  /** A list of edges. */
  edges?: Maybe<Array<Maybe<AppResEdge>>>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** Identifies the total count of items in the connection. */
  totalCount: Scalars['Int']['output'];
};

/** An edge in a connection. */
export type AppResEdge = {
  __typename?: 'AppResEdge';
  /** A cursor for use in pagination. */
  cursor: Scalars['Cursor']['output'];
  /** The item at the end of the edge. */
  node?: Maybe<AppRes>;
};

/** Ordering options for AppRes connections */
export type AppResOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order AppResSlice. */
  field: AppResOrderField;
};

/** Properties by which AppRes connections can be ordered. */
export enum AppResOrderField {
  CreatedAt = 'createdAt'
}

/**
 * AppResWhereInput is used for filtering AppRes objects.
 * Input was generated by ent.
 */
export type AppResWhereInput = {
  and?: InputMaybe<Array<AppResWhereInput>>;
  /** app_id field predicates */
  appID?: InputMaybe<Scalars['ID']['input']>;
  appIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  appIDIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  appIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  appIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  appIDNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** arn_pattern field predicates */
  arnPattern?: InputMaybe<Scalars['String']['input']>;
  arnPatternContains?: InputMaybe<Scalars['String']['input']>;
  arnPatternContainsFold?: InputMaybe<Scalars['String']['input']>;
  arnPatternEqualFold?: InputMaybe<Scalars['String']['input']>;
  arnPatternGT?: InputMaybe<Scalars['String']['input']>;
  arnPatternGTE?: InputMaybe<Scalars['String']['input']>;
  arnPatternHasPrefix?: InputMaybe<Scalars['String']['input']>;
  arnPatternHasSuffix?: InputMaybe<Scalars['String']['input']>;
  arnPatternIn?: InputMaybe<Array<Scalars['String']['input']>>;
  arnPatternLT?: InputMaybe<Scalars['String']['input']>;
  arnPatternLTE?: InputMaybe<Scalars['String']['input']>;
  arnPatternNEQ?: InputMaybe<Scalars['String']['input']>;
  arnPatternNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  /** created_at field predicates */
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  createdAtGT?: InputMaybe<Scalars['Time']['input']>;
  createdAtGTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  createdAtLT?: InputMaybe<Scalars['Time']['input']>;
  createdAtLTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  createdAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  /** created_by field predicates */
  createdBy?: InputMaybe<Scalars['Int']['input']>;
  createdByGT?: InputMaybe<Scalars['Int']['input']>;
  createdByGTE?: InputMaybe<Scalars['Int']['input']>;
  createdByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  createdByLT?: InputMaybe<Scalars['Int']['input']>;
  createdByLTE?: InputMaybe<Scalars['Int']['input']>;
  createdByNEQ?: InputMaybe<Scalars['Int']['input']>;
  createdByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  /** app edge predicates */
  hasApp?: InputMaybe<Scalars['Boolean']['input']>;
  hasAppWith?: InputMaybe<Array<AppWhereInput>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** name field predicates */
  name?: InputMaybe<Scalars['String']['input']>;
  nameContains?: InputMaybe<Scalars['String']['input']>;
  nameContainsFold?: InputMaybe<Scalars['String']['input']>;
  nameEqualFold?: InputMaybe<Scalars['String']['input']>;
  nameGT?: InputMaybe<Scalars['String']['input']>;
  nameGTE?: InputMaybe<Scalars['String']['input']>;
  nameHasPrefix?: InputMaybe<Scalars['String']['input']>;
  nameHasSuffix?: InputMaybe<Scalars['String']['input']>;
  nameIn?: InputMaybe<Array<Scalars['String']['input']>>;
  nameLT?: InputMaybe<Scalars['String']['input']>;
  nameLTE?: InputMaybe<Scalars['String']['input']>;
  nameNEQ?: InputMaybe<Scalars['String']['input']>;
  nameNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  not?: InputMaybe<AppResWhereInput>;
  or?: InputMaybe<Array<AppResWhereInput>>;
  /** type_name field predicates */
  typeName?: InputMaybe<Scalars['String']['input']>;
  typeNameContains?: InputMaybe<Scalars['String']['input']>;
  typeNameContainsFold?: InputMaybe<Scalars['String']['input']>;
  typeNameEqualFold?: InputMaybe<Scalars['String']['input']>;
  typeNameGT?: InputMaybe<Scalars['String']['input']>;
  typeNameGTE?: InputMaybe<Scalars['String']['input']>;
  typeNameHasPrefix?: InputMaybe<Scalars['String']['input']>;
  typeNameHasSuffix?: InputMaybe<Scalars['String']['input']>;
  typeNameIn?: InputMaybe<Array<Scalars['String']['input']>>;
  typeNameLT?: InputMaybe<Scalars['String']['input']>;
  typeNameLTE?: InputMaybe<Scalars['String']['input']>;
  typeNameNEQ?: InputMaybe<Scalars['String']['input']>;
  typeNameNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  /** updated_at field predicates */
  updatedAt?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedAtLT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtLTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_by field predicates */
  updatedBy?: InputMaybe<Scalars['Int']['input']>;
  updatedByGT?: InputMaybe<Scalars['Int']['input']>;
  updatedByGTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedByLT?: InputMaybe<Scalars['Int']['input']>;
  updatedByLTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByNEQ?: InputMaybe<Scalars['Int']['input']>;
  updatedByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByNotNil?: InputMaybe<Scalars['Boolean']['input']>;
};

export type AppRole = Node & {
  __typename?: 'AppRole';
  app?: Maybe<App>;
  /** 所属应用 */
  appID?: Maybe<Scalars['ID']['output']>;
  /** 标识是否自动授予到账户 */
  autoGrant: Scalars['Boolean']['output'];
  /** 备注 */
  comments?: Maybe<Scalars['String']['output']>;
  createdAt: Scalars['Time']['output'];
  createdBy: Scalars['Int']['output'];
  /** 授权后是否可编辑 */
  editable: Scalars['Boolean']['output'];
  id: Scalars['ID']['output'];
  /** 角色名称 */
  name: Scalars['String']['output'];
  /** 权限授权策略 */
  policies?: Maybe<Array<AppPolicy>>;
  updatedAt?: Maybe<Scalars['Time']['output']>;
  updatedBy?: Maybe<Scalars['Int']['output']>;
};

/** Ordering options for AppRole connections */
export type AppRoleOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order AppRoles. */
  field: AppRoleOrderField;
};

/** Properties by which AppRole connections can be ordered. */
export enum AppRoleOrderField {
  CreatedAt = 'createdAt'
}

/** Ordering options for AppRolePolicy connections */
export type AppRolePolicyOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order AppRolePolicies. */
  field: AppRolePolicyOrderField;
};

/** Properties by which AppRolePolicy connections can be ordered. */
export enum AppRolePolicyOrderField {
  CreatedAt = 'createdAt'
}

/**
 * AppRolePolicyWhereInput is used for filtering AppRolePolicy objects.
 * Input was generated by ent.
 */
export type AppRolePolicyWhereInput = {
  and?: InputMaybe<Array<AppRolePolicyWhereInput>>;
  /** app_id field predicates */
  appID?: InputMaybe<Scalars['Int']['input']>;
  appIDGT?: InputMaybe<Scalars['Int']['input']>;
  appIDGTE?: InputMaybe<Scalars['Int']['input']>;
  appIDIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  appIDLT?: InputMaybe<Scalars['Int']['input']>;
  appIDLTE?: InputMaybe<Scalars['Int']['input']>;
  appIDNEQ?: InputMaybe<Scalars['Int']['input']>;
  appIDNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  /** created_at field predicates */
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  createdAtGT?: InputMaybe<Scalars['Time']['input']>;
  createdAtGTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  createdAtLT?: InputMaybe<Scalars['Time']['input']>;
  createdAtLTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  createdAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  /** created_by field predicates */
  createdBy?: InputMaybe<Scalars['Int']['input']>;
  createdByGT?: InputMaybe<Scalars['Int']['input']>;
  createdByGTE?: InputMaybe<Scalars['Int']['input']>;
  createdByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  createdByLT?: InputMaybe<Scalars['Int']['input']>;
  createdByLTE?: InputMaybe<Scalars['Int']['input']>;
  createdByNEQ?: InputMaybe<Scalars['Int']['input']>;
  createdByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  not?: InputMaybe<AppRolePolicyWhereInput>;
  or?: InputMaybe<Array<AppRolePolicyWhereInput>>;
  /** updated_at field predicates */
  updatedAt?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedAtLT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtLTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_by field predicates */
  updatedBy?: InputMaybe<Scalars['Int']['input']>;
  updatedByGT?: InputMaybe<Scalars['Int']['input']>;
  updatedByGTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedByLT?: InputMaybe<Scalars['Int']['input']>;
  updatedByLTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByNEQ?: InputMaybe<Scalars['Int']['input']>;
  updatedByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByNotNil?: InputMaybe<Scalars['Boolean']['input']>;
};

/**
 * AppRoleWhereInput is used for filtering AppRole objects.
 * Input was generated by ent.
 */
export type AppRoleWhereInput = {
  and?: InputMaybe<Array<AppRoleWhereInput>>;
  /** app_id field predicates */
  appID?: InputMaybe<Scalars['ID']['input']>;
  appIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  appIDIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  appIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  appIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  appIDNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** auto_grant field predicates */
  autoGrant?: InputMaybe<Scalars['Boolean']['input']>;
  autoGrantNEQ?: InputMaybe<Scalars['Boolean']['input']>;
  /** created_at field predicates */
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  createdAtGT?: InputMaybe<Scalars['Time']['input']>;
  createdAtGTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  createdAtLT?: InputMaybe<Scalars['Time']['input']>;
  createdAtLTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  createdAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  /** created_by field predicates */
  createdBy?: InputMaybe<Scalars['Int']['input']>;
  createdByGT?: InputMaybe<Scalars['Int']['input']>;
  createdByGTE?: InputMaybe<Scalars['Int']['input']>;
  createdByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  createdByLT?: InputMaybe<Scalars['Int']['input']>;
  createdByLTE?: InputMaybe<Scalars['Int']['input']>;
  createdByNEQ?: InputMaybe<Scalars['Int']['input']>;
  createdByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  /** editable field predicates */
  editable?: InputMaybe<Scalars['Boolean']['input']>;
  editableNEQ?: InputMaybe<Scalars['Boolean']['input']>;
  /** app edge predicates */
  hasApp?: InputMaybe<Scalars['Boolean']['input']>;
  /** app_role_policy edge predicates */
  hasAppRolePolicy?: InputMaybe<Scalars['Boolean']['input']>;
  hasAppRolePolicyWith?: InputMaybe<Array<AppRolePolicyWhereInput>>;
  hasAppWith?: InputMaybe<Array<AppWhereInput>>;
  /** policies edge predicates */
  hasPolicies?: InputMaybe<Scalars['Boolean']['input']>;
  hasPoliciesWith?: InputMaybe<Array<AppPolicyWhereInput>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** name field predicates */
  name?: InputMaybe<Scalars['String']['input']>;
  nameContains?: InputMaybe<Scalars['String']['input']>;
  nameContainsFold?: InputMaybe<Scalars['String']['input']>;
  nameEqualFold?: InputMaybe<Scalars['String']['input']>;
  nameGT?: InputMaybe<Scalars['String']['input']>;
  nameGTE?: InputMaybe<Scalars['String']['input']>;
  nameHasPrefix?: InputMaybe<Scalars['String']['input']>;
  nameHasSuffix?: InputMaybe<Scalars['String']['input']>;
  nameIn?: InputMaybe<Array<Scalars['String']['input']>>;
  nameLT?: InputMaybe<Scalars['String']['input']>;
  nameLTE?: InputMaybe<Scalars['String']['input']>;
  nameNEQ?: InputMaybe<Scalars['String']['input']>;
  nameNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  not?: InputMaybe<AppRoleWhereInput>;
  or?: InputMaybe<Array<AppRoleWhereInput>>;
  /** updated_at field predicates */
  updatedAt?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedAtLT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtLTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_by field predicates */
  updatedBy?: InputMaybe<Scalars['Int']['input']>;
  updatedByGT?: InputMaybe<Scalars['Int']['input']>;
  updatedByGTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedByLT?: InputMaybe<Scalars['Int']['input']>;
  updatedByLTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByNEQ?: InputMaybe<Scalars['Int']['input']>;
  updatedByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByNotNil?: InputMaybe<Scalars['Boolean']['input']>;
};

/** AppSimpleStatus is enum for the field status */
export enum AppSimpleStatus {
  Active = 'active',
  Disabled = 'disabled',
  Inactive = 'inactive',
  Processing = 'processing'
}

/**
 * AppWhereInput is used for filtering App objects.
 * Input was generated by ent.
 */
export type AppWhereInput = {
  and?: InputMaybe<Array<AppWhereInput>>;
  /** app_key field predicates */
  appKey?: InputMaybe<Scalars['String']['input']>;
  appKeyContains?: InputMaybe<Scalars['String']['input']>;
  appKeyContainsFold?: InputMaybe<Scalars['String']['input']>;
  appKeyEqualFold?: InputMaybe<Scalars['String']['input']>;
  appKeyGT?: InputMaybe<Scalars['String']['input']>;
  appKeyGTE?: InputMaybe<Scalars['String']['input']>;
  appKeyHasPrefix?: InputMaybe<Scalars['String']['input']>;
  appKeyHasSuffix?: InputMaybe<Scalars['String']['input']>;
  appKeyIn?: InputMaybe<Array<Scalars['String']['input']>>;
  appKeyIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  appKeyLT?: InputMaybe<Scalars['String']['input']>;
  appKeyLTE?: InputMaybe<Scalars['String']['input']>;
  appKeyNEQ?: InputMaybe<Scalars['String']['input']>;
  appKeyNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  appKeyNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** app_secret field predicates */
  appSecret?: InputMaybe<Scalars['String']['input']>;
  appSecretContains?: InputMaybe<Scalars['String']['input']>;
  appSecretContainsFold?: InputMaybe<Scalars['String']['input']>;
  appSecretEqualFold?: InputMaybe<Scalars['String']['input']>;
  appSecretGT?: InputMaybe<Scalars['String']['input']>;
  appSecretGTE?: InputMaybe<Scalars['String']['input']>;
  appSecretHasPrefix?: InputMaybe<Scalars['String']['input']>;
  appSecretHasSuffix?: InputMaybe<Scalars['String']['input']>;
  appSecretIn?: InputMaybe<Array<Scalars['String']['input']>>;
  appSecretIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  appSecretLT?: InputMaybe<Scalars['String']['input']>;
  appSecretLTE?: InputMaybe<Scalars['String']['input']>;
  appSecretNEQ?: InputMaybe<Scalars['String']['input']>;
  appSecretNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  appSecretNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** code field predicates */
  code?: InputMaybe<Scalars['String']['input']>;
  codeContains?: InputMaybe<Scalars['String']['input']>;
  codeContainsFold?: InputMaybe<Scalars['String']['input']>;
  codeEqualFold?: InputMaybe<Scalars['String']['input']>;
  codeGT?: InputMaybe<Scalars['String']['input']>;
  codeGTE?: InputMaybe<Scalars['String']['input']>;
  codeHasPrefix?: InputMaybe<Scalars['String']['input']>;
  codeHasSuffix?: InputMaybe<Scalars['String']['input']>;
  codeIn?: InputMaybe<Array<Scalars['String']['input']>>;
  codeLT?: InputMaybe<Scalars['String']['input']>;
  codeLTE?: InputMaybe<Scalars['String']['input']>;
  codeNEQ?: InputMaybe<Scalars['String']['input']>;
  codeNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  /** comments field predicates */
  comments?: InputMaybe<Scalars['String']['input']>;
  commentsContains?: InputMaybe<Scalars['String']['input']>;
  commentsContainsFold?: InputMaybe<Scalars['String']['input']>;
  commentsEqualFold?: InputMaybe<Scalars['String']['input']>;
  commentsGT?: InputMaybe<Scalars['String']['input']>;
  commentsGTE?: InputMaybe<Scalars['String']['input']>;
  commentsHasPrefix?: InputMaybe<Scalars['String']['input']>;
  commentsHasSuffix?: InputMaybe<Scalars['String']['input']>;
  commentsIn?: InputMaybe<Array<Scalars['String']['input']>>;
  commentsIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  commentsLT?: InputMaybe<Scalars['String']['input']>;
  commentsLTE?: InputMaybe<Scalars['String']['input']>;
  commentsNEQ?: InputMaybe<Scalars['String']['input']>;
  commentsNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  commentsNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** created_at field predicates */
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  createdAtGT?: InputMaybe<Scalars['Time']['input']>;
  createdAtGTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  createdAtLT?: InputMaybe<Scalars['Time']['input']>;
  createdAtLTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  createdAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  /** created_by field predicates */
  createdBy?: InputMaybe<Scalars['Int']['input']>;
  createdByGT?: InputMaybe<Scalars['Int']['input']>;
  createdByGTE?: InputMaybe<Scalars['Int']['input']>;
  createdByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  createdByLT?: InputMaybe<Scalars['Int']['input']>;
  createdByLTE?: InputMaybe<Scalars['Int']['input']>;
  createdByNEQ?: InputMaybe<Scalars['Int']['input']>;
  createdByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  /** actions edge predicates */
  hasActions?: InputMaybe<Scalars['Boolean']['input']>;
  hasActionsWith?: InputMaybe<Array<AppActionWhereInput>>;
  /** dicts edge predicates */
  hasDicts?: InputMaybe<Scalars['Boolean']['input']>;
  hasDictsWith?: InputMaybe<Array<AppDictWhereInput>>;
  /** menus edge predicates */
  hasMenus?: InputMaybe<Scalars['Boolean']['input']>;
  hasMenusWith?: InputMaybe<Array<AppMenuWhereInput>>;
  /** orgs edge predicates */
  hasOrgs?: InputMaybe<Scalars['Boolean']['input']>;
  hasOrgsWith?: InputMaybe<Array<OrgWhereInput>>;
  /** policies edge predicates */
  hasPolicies?: InputMaybe<Scalars['Boolean']['input']>;
  hasPoliciesWith?: InputMaybe<Array<AppPolicyWhereInput>>;
  /** policy_views edge predicates */
  hasPolicyViews?: InputMaybe<Scalars['Boolean']['input']>;
  hasPolicyViewsWith?: InputMaybe<Array<AppPolicyViewWhereInput>>;
  /** resources edge predicates */
  hasResources?: InputMaybe<Scalars['Boolean']['input']>;
  hasResourcesWith?: InputMaybe<Array<AppResWhereInput>>;
  /** roles edge predicates */
  hasRoles?: InputMaybe<Scalars['Boolean']['input']>;
  hasRolesWith?: InputMaybe<Array<AppRoleWhereInput>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** kind field predicates */
  kind?: InputMaybe<AppKind>;
  kindIn?: InputMaybe<Array<AppKind>>;
  kindNEQ?: InputMaybe<AppKind>;
  kindNotIn?: InputMaybe<Array<AppKind>>;
  /** name field predicates */
  name?: InputMaybe<Scalars['String']['input']>;
  nameContains?: InputMaybe<Scalars['String']['input']>;
  nameContainsFold?: InputMaybe<Scalars['String']['input']>;
  nameEqualFold?: InputMaybe<Scalars['String']['input']>;
  nameGT?: InputMaybe<Scalars['String']['input']>;
  nameGTE?: InputMaybe<Scalars['String']['input']>;
  nameHasPrefix?: InputMaybe<Scalars['String']['input']>;
  nameHasSuffix?: InputMaybe<Scalars['String']['input']>;
  nameIn?: InputMaybe<Array<Scalars['String']['input']>>;
  nameLT?: InputMaybe<Scalars['String']['input']>;
  nameLTE?: InputMaybe<Scalars['String']['input']>;
  nameNEQ?: InputMaybe<Scalars['String']['input']>;
  nameNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  not?: InputMaybe<AppWhereInput>;
  or?: InputMaybe<Array<AppWhereInput>>;
  /** redirect_uri field predicates */
  redirectURI?: InputMaybe<Scalars['String']['input']>;
  redirectURIContains?: InputMaybe<Scalars['String']['input']>;
  redirectURIContainsFold?: InputMaybe<Scalars['String']['input']>;
  redirectURIEqualFold?: InputMaybe<Scalars['String']['input']>;
  redirectURIGT?: InputMaybe<Scalars['String']['input']>;
  redirectURIGTE?: InputMaybe<Scalars['String']['input']>;
  redirectURIHasPrefix?: InputMaybe<Scalars['String']['input']>;
  redirectURIHasSuffix?: InputMaybe<Scalars['String']['input']>;
  redirectURIIn?: InputMaybe<Array<Scalars['String']['input']>>;
  redirectURIIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  redirectURILT?: InputMaybe<Scalars['String']['input']>;
  redirectURILTE?: InputMaybe<Scalars['String']['input']>;
  redirectURINEQ?: InputMaybe<Scalars['String']['input']>;
  redirectURINotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  redirectURINotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** refresh_token_validity field predicates */
  refreshTokenValidity?: InputMaybe<Scalars['Int']['input']>;
  refreshTokenValidityGT?: InputMaybe<Scalars['Int']['input']>;
  refreshTokenValidityGTE?: InputMaybe<Scalars['Int']['input']>;
  refreshTokenValidityIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  refreshTokenValidityIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  refreshTokenValidityLT?: InputMaybe<Scalars['Int']['input']>;
  refreshTokenValidityLTE?: InputMaybe<Scalars['Int']['input']>;
  refreshTokenValidityNEQ?: InputMaybe<Scalars['Int']['input']>;
  refreshTokenValidityNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  refreshTokenValidityNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** scopes field predicates */
  scopes?: InputMaybe<Scalars['String']['input']>;
  scopesContains?: InputMaybe<Scalars['String']['input']>;
  scopesContainsFold?: InputMaybe<Scalars['String']['input']>;
  scopesEqualFold?: InputMaybe<Scalars['String']['input']>;
  scopesGT?: InputMaybe<Scalars['String']['input']>;
  scopesGTE?: InputMaybe<Scalars['String']['input']>;
  scopesHasPrefix?: InputMaybe<Scalars['String']['input']>;
  scopesHasSuffix?: InputMaybe<Scalars['String']['input']>;
  scopesIn?: InputMaybe<Array<Scalars['String']['input']>>;
  scopesIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  scopesLT?: InputMaybe<Scalars['String']['input']>;
  scopesLTE?: InputMaybe<Scalars['String']['input']>;
  scopesNEQ?: InputMaybe<Scalars['String']['input']>;
  scopesNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  scopesNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** status field predicates */
  status?: InputMaybe<AppSimpleStatus>;
  statusIn?: InputMaybe<Array<AppSimpleStatus>>;
  statusIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  statusNEQ?: InputMaybe<AppSimpleStatus>;
  statusNotIn?: InputMaybe<Array<AppSimpleStatus>>;
  statusNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** token_validity field predicates */
  tokenValidity?: InputMaybe<Scalars['Int']['input']>;
  tokenValidityGT?: InputMaybe<Scalars['Int']['input']>;
  tokenValidityGTE?: InputMaybe<Scalars['Int']['input']>;
  tokenValidityIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  tokenValidityIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  tokenValidityLT?: InputMaybe<Scalars['Int']['input']>;
  tokenValidityLTE?: InputMaybe<Scalars['Int']['input']>;
  tokenValidityNEQ?: InputMaybe<Scalars['Int']['input']>;
  tokenValidityNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  tokenValidityNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_at field predicates */
  updatedAt?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedAtLT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtLTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_by field predicates */
  updatedBy?: InputMaybe<Scalars['Int']['input']>;
  updatedByGT?: InputMaybe<Scalars['Int']['input']>;
  updatedByGTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedByLT?: InputMaybe<Scalars['Int']['input']>;
  updatedByLTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByNEQ?: InputMaybe<Scalars['Int']['input']>;
  updatedByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByNotNil?: InputMaybe<Scalars['Boolean']['input']>;
};

export type AssignRoleUserInput = {
  /** 生效结束时间 */
  endAt?: InputMaybe<Scalars['Time']['input']>;
  /** 授权类型为角色或用户组的ID */
  orgRoleID: Scalars['ID']['input'];
  /** 生效开始时间 */
  startAt?: InputMaybe<Scalars['Time']['input']>;
  userID: Scalars['ID']['input'];
};

export type Country = Node & {
  __typename?: 'Country';
  /** 国家代码 */
  code: Scalars['String']['output'];
  createdAt: Scalars['Time']['output'];
  createdBy: Scalars['Int']['output'];
  displaySort?: Maybe<Scalars['Int']['output']>;
  id: Scalars['ID']['output'];
  /** 国家中文名称 */
  name?: Maybe<Scalars['String']['output']>;
  /** 国家英文名称 */
  nameEn?: Maybe<Scalars['String']['output']>;
  /** 地区信息 */
  regions?: Maybe<Array<Region>>;
  /** 状态 */
  status?: Maybe<CountrySimpleStatus>;
  updatedAt?: Maybe<Scalars['Time']['output']>;
  updatedBy?: Maybe<Scalars['Int']['output']>;
};

/** A connection to a list of items. */
export type CountryConnection = {
  __typename?: 'CountryConnection';
  /** A list of edges. */
  edges?: Maybe<Array<Maybe<CountryEdge>>>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** Identifies the total count of items in the connection. */
  totalCount: Scalars['Int']['output'];
};

/** An edge in a connection. */
export type CountryEdge = {
  __typename?: 'CountryEdge';
  /** A cursor for use in pagination. */
  cursor: Scalars['Cursor']['output'];
  /** The item at the end of the edge. */
  node?: Maybe<Country>;
};

/** Ordering options for Country connections */
export type CountryOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order Countries. */
  field: CountryOrderField;
};

/** Properties by which Country connections can be ordered. */
export enum CountryOrderField {
  CreatedAt = 'createdAt',
  DisplaySort = 'displaySort'
}

/** CountrySimpleStatus is enum for the field status */
export enum CountrySimpleStatus {
  Active = 'active',
  Disabled = 'disabled',
  Inactive = 'inactive',
  Processing = 'processing'
}

/**
 * CountryWhereInput is used for filtering Country objects.
 * Input was generated by ent.
 */
export type CountryWhereInput = {
  and?: InputMaybe<Array<CountryWhereInput>>;
  /** code field predicates */
  code?: InputMaybe<Scalars['String']['input']>;
  codeContains?: InputMaybe<Scalars['String']['input']>;
  codeContainsFold?: InputMaybe<Scalars['String']['input']>;
  codeEqualFold?: InputMaybe<Scalars['String']['input']>;
  codeGT?: InputMaybe<Scalars['String']['input']>;
  codeGTE?: InputMaybe<Scalars['String']['input']>;
  codeHasPrefix?: InputMaybe<Scalars['String']['input']>;
  codeHasSuffix?: InputMaybe<Scalars['String']['input']>;
  codeIn?: InputMaybe<Array<Scalars['String']['input']>>;
  codeLT?: InputMaybe<Scalars['String']['input']>;
  codeLTE?: InputMaybe<Scalars['String']['input']>;
  codeNEQ?: InputMaybe<Scalars['String']['input']>;
  codeNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  /** created_at field predicates */
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  createdAtGT?: InputMaybe<Scalars['Time']['input']>;
  createdAtGTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  createdAtLT?: InputMaybe<Scalars['Time']['input']>;
  createdAtLTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  createdAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  /** created_by field predicates */
  createdBy?: InputMaybe<Scalars['Int']['input']>;
  createdByGT?: InputMaybe<Scalars['Int']['input']>;
  createdByGTE?: InputMaybe<Scalars['Int']['input']>;
  createdByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  createdByLT?: InputMaybe<Scalars['Int']['input']>;
  createdByLTE?: InputMaybe<Scalars['Int']['input']>;
  createdByNEQ?: InputMaybe<Scalars['Int']['input']>;
  createdByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  /** regions edge predicates */
  hasRegions?: InputMaybe<Scalars['Boolean']['input']>;
  hasRegionsWith?: InputMaybe<Array<RegionWhereInput>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** name field predicates */
  name?: InputMaybe<Scalars['String']['input']>;
  nameContains?: InputMaybe<Scalars['String']['input']>;
  nameContainsFold?: InputMaybe<Scalars['String']['input']>;
  /** name_en field predicates */
  nameEn?: InputMaybe<Scalars['String']['input']>;
  nameEnContains?: InputMaybe<Scalars['String']['input']>;
  nameEnContainsFold?: InputMaybe<Scalars['String']['input']>;
  nameEnEqualFold?: InputMaybe<Scalars['String']['input']>;
  nameEnGT?: InputMaybe<Scalars['String']['input']>;
  nameEnGTE?: InputMaybe<Scalars['String']['input']>;
  nameEnHasPrefix?: InputMaybe<Scalars['String']['input']>;
  nameEnHasSuffix?: InputMaybe<Scalars['String']['input']>;
  nameEnIn?: InputMaybe<Array<Scalars['String']['input']>>;
  nameEnIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  nameEnLT?: InputMaybe<Scalars['String']['input']>;
  nameEnLTE?: InputMaybe<Scalars['String']['input']>;
  nameEnNEQ?: InputMaybe<Scalars['String']['input']>;
  nameEnNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  nameEnNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  nameEqualFold?: InputMaybe<Scalars['String']['input']>;
  nameGT?: InputMaybe<Scalars['String']['input']>;
  nameGTE?: InputMaybe<Scalars['String']['input']>;
  nameHasPrefix?: InputMaybe<Scalars['String']['input']>;
  nameHasSuffix?: InputMaybe<Scalars['String']['input']>;
  nameIn?: InputMaybe<Array<Scalars['String']['input']>>;
  nameIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  nameLT?: InputMaybe<Scalars['String']['input']>;
  nameLTE?: InputMaybe<Scalars['String']['input']>;
  nameNEQ?: InputMaybe<Scalars['String']['input']>;
  nameNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  nameNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  not?: InputMaybe<CountryWhereInput>;
  or?: InputMaybe<Array<CountryWhereInput>>;
  /** status field predicates */
  status?: InputMaybe<CountrySimpleStatus>;
  statusIn?: InputMaybe<Array<CountrySimpleStatus>>;
  statusIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  statusNEQ?: InputMaybe<CountrySimpleStatus>;
  statusNotIn?: InputMaybe<Array<CountrySimpleStatus>>;
  statusNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_at field predicates */
  updatedAt?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedAtLT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtLTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_by field predicates */
  updatedBy?: InputMaybe<Scalars['Int']['input']>;
  updatedByGT?: InputMaybe<Scalars['Int']['input']>;
  updatedByGTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedByLT?: InputMaybe<Scalars['Int']['input']>;
  updatedByLTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByNEQ?: InputMaybe<Scalars['Int']['input']>;
  updatedByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByNotNil?: InputMaybe<Scalars['Boolean']['input']>;
};

/**
 * CreateAppActionInput is used for create AppAction object.
 * Input was generated by ent.
 */
export type CreateAppActionInput = {
  appID?: InputMaybe<Scalars['ID']['input']>;
  /** 备注 */
  comments?: InputMaybe<Scalars['String']['input']>;
  /** restful,graphql,rpc,function */
  kind: AppActionKind;
  menuIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** 操作方法:读,写,列表 */
  method: AppActionMethod;
  /** 名称 */
  name: Scalars['String']['input'];
};

/**
 * CreateAppDictInput is used for create AppDict object.
 * Input was generated by ent.
 */
export type CreateAppDictInput = {
  appID?: InputMaybe<Scalars['ID']['input']>;
  /** 用于标识应用资源的唯一代码,尽量简短 */
  code: Scalars['String']['input'];
  /** 备注 */
  comments?: InputMaybe<Scalars['String']['input']>;
  itemIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** 名称 */
  name: Scalars['String']['input'];
};

/**
 * CreateAppDictItemInput is used for create AppDictItem object.
 * Input was generated by ent.
 */
export type CreateAppDictItemInput = {
  /** 字典值唯一编码,生效后不可修改. */
  code: Scalars['String']['input'];
  /** 备注 */
  comments?: InputMaybe<Scalars['String']['input']>;
  dictID?: InputMaybe<Scalars['ID']['input']>;
  /** 名称 */
  name: Scalars['String']['input'];
  orgID?: InputMaybe<Scalars['ID']['input']>;
  /** 状态 */
  status?: InputMaybe<AppDictItemSimpleStatus>;
};

/**
 * CreateAppInput is used for create App object.
 * Input was generated by ent.
 */
export type CreateAppInput = {
  actionIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** 应用ID */
  appKey?: InputMaybe<Scalars['String']['input']>;
  /** 应用密钥 */
  appSecret?: InputMaybe<Scalars['String']['input']>;
  /** 用于标识应用资源的唯一代码,尽量简短 */
  code: Scalars['String']['input'];
  /** 备注 */
  comments?: InputMaybe<Scalars['String']['input']>;
  dictIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** 应用类型 */
  kind: AppKind;
  /** 应用图标地址 */
  logo?: InputMaybe<Scalars['String']['input']>;
  menuIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** 名称 */
  name: Scalars['String']['input'];
  policyIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  policyViewIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** 回调地址 */
  redirectURI?: InputMaybe<Scalars['String']['input']>;
  /** refresh_token有效期 */
  refreshTokenValidity?: InputMaybe<Scalars['Int']['input']>;
  resourceIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  roleIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** 权限范围 */
  scopes?: InputMaybe<Scalars['String']['input']>;
  /** 状态 */
  status?: InputMaybe<AppSimpleStatus>;
  /** token有效期 */
  tokenValidity?: InputMaybe<Scalars['Int']['input']>;
};

/**
 * CreateAppMenuInput is used for create AppMenu object.
 * Input was generated by ent.
 */
export type CreateAppMenuInput = {
  actionID?: InputMaybe<Scalars['ID']['input']>;
  appID?: InputMaybe<Scalars['ID']['input']>;
  childIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** 备注 */
  comments?: InputMaybe<Scalars['String']['input']>;
  /** 菜单图标 */
  icon?: InputMaybe<Scalars['String']['input']>;
  /** 目录,菜单项 */
  kind: AppMenuKind;
  /** 菜单名称 */
  name: Scalars['String']['input'];
  parentID: Scalars['ID']['input'];
  /** 菜单路由 */
  route?: InputMaybe<Scalars['String']['input']>;
  /** 状态 */
  status?: InputMaybe<AppMenuSimpleStatus>;
};

/**
 * CreateAppPolicyInput is used for create AppPolicy object.
 * Input was generated by ent.
 */
export type CreateAppPolicyInput = {
  appID?: InputMaybe<Scalars['ID']['input']>;
  /** 标识是否自动授予到账户 */
  autoGrant?: InputMaybe<Scalars['Boolean']['input']>;
  /** 描述 */
  comments?: InputMaybe<Scalars['String']['input']>;
  /** 分类：app-应用策略、view-策略视图 */
  kind?: InputMaybe<AppPolicyKind>;
  /** 策略名称 */
  name: Scalars['String']['input'];
  orgPolicyIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  policyViewIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  roleIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** 策略规则 */
  rules: Array<InputMaybe<PolicyRuleInput>>;
  /** 状态 */
  status?: InputMaybe<AppPolicySimpleStatus>;
};

/**
 * CreateAppPolicyViewInput is used for create AppPolicyView object.
 * Input was generated by ent.
 */
export type CreateAppPolicyViewInput = {
  appID?: InputMaybe<Scalars['ID']['input']>;
  appPolicyID?: InputMaybe<Scalars['ID']['input']>;
  childIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** 描述 */
  comments?: InputMaybe<Scalars['String']['input']>;
  /** 分类：dir-目录、policy-权限策略 */
  kind: AppPolicyViewKind;
  /** 名称 */
  name: Scalars['String']['input'];
  parentID: Scalars['ID']['input'];
};

/**
 * CreateAppResInput is used for create AppRes object.
 * Input was generated by ent.
 */
export type CreateAppResInput = {
  appID?: InputMaybe<Scalars['ID']['input']>;
  /** 应用资源表达式 */
  arnPattern: Scalars['String']['input'];
  /** 资源名称 */
  name: Scalars['String']['input'];
  /** 资源类型名称,如数据库表名 */
  typeName: Scalars['String']['input'];
};

/**
 * CreateAppRoleInput is used for create AppRole object.
 * Input was generated by ent.
 */
export type CreateAppRoleInput = {
  appID?: InputMaybe<Scalars['ID']['input']>;
  /** 标识是否自动授予到账户 */
  autoGrant?: InputMaybe<Scalars['Boolean']['input']>;
  /** 备注 */
  comments?: InputMaybe<Scalars['String']['input']>;
  /** 授权后是否可编辑 */
  editable?: InputMaybe<Scalars['Boolean']['input']>;
  /** 角色名称 */
  name: Scalars['String']['input'];
};

/**
 * CreateCountryInput is used for create Country object.
 * Input was generated by ent.
 */
export type CreateCountryInput = {
  /** 国家代码 */
  code: Scalars['String']['input'];
  /** 国家中文名称 */
  name?: InputMaybe<Scalars['String']['input']>;
  /** 国家英文名称 */
  nameEn?: InputMaybe<Scalars['String']['input']>;
  regionIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** 状态 */
  status?: InputMaybe<CountrySimpleStatus>;
};

/**
 * CreateCurrencyInput is used for create Currency object.
 * Input was generated by ent.
 */
export type CreateCurrencyInput = {
  /** 代码 */
  code: Scalars['String']['input'];
  /** 中文名称 */
  name: Scalars['String']['input'];
  /** 货币符号 */
  sign?: InputMaybe<Scalars['String']['input']>;
  /** 状态 */
  status?: InputMaybe<CurrencySimpleStatus>;
};

/**
 * CreateFileIdentityInput is used for create FileIdentity object.
 * Input was generated by ent.
 */
export type CreateFileIdentityInput = {
  /** accesskey id */
  accessKeyID: Scalars['String']['input'];
  /** accesskey secret */
  accessKeySecret: Scalars['String']['input'];
  /** 备注 */
  comments?: InputMaybe<Scalars['String']['input']>;
  /** STS令牌的有效期，默认3600s */
  durationSeconds?: InputMaybe<Scalars['Int']['input']>;
  orgID: Scalars['ID']['input'];
  /** 指定返回的STS令牌的权限的策略 */
  policy?: InputMaybe<Scalars['String']['input']>;
  /** 角色的资源名称(ARN)，用于STS */
  roleArn: Scalars['String']['input'];
  sourceID: Scalars['ID']['input'];
};

/**
 * CreateFileSourceInput is used for create FileSource object.
 * Input was generated by ent.
 */
export type CreateFileSourceInput = {
  /** 文件存储空间 */
  bucket: Scalars['String']['input'];
  /** 文件存储空间地址，用于匹配url */
  bucketURL: Scalars['String']['input'];
  /** 备注 */
  comments?: InputMaybe<Scalars['String']['input']>;
  /** 对外服务的访问域名 */
  endpoint: Scalars['String']['input'];
  /** 是否禁止修改endpoint，如果是自定义域名设为true */
  endpointImmutable?: InputMaybe<Scalars['Boolean']['input']>;
  identityIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** 文件来源 */
  kind: FileSourceKind;
  /** 地域，数据存储的物理位置 */
  region: Scalars['String']['input'];
  /** sts服务的访问域名 */
  stsEndpoint: Scalars['String']['input'];
};

/**
 * CreateOauthClientInput is used for create OauthClient object.
 * Input was generated by ent.
 */
export type CreateOauthClientInput = {
  /** 授权类型 */
  grantTypes: OauthClientGrantTypes;
  /** 名称 */
  name: Scalars['String']['input'];
  userID: Scalars['ID']['input'];
};

/**
 * CreateOrgInput is used for create Org object.
 * Input was generated by ent.
 */
export type CreateOrgInput = {
  appIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  childIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** 国家或地区2字码 */
  countryCode?: InputMaybe<Scalars['String']['input']>;
  /** 自定义域名 */
  customDomain?: InputMaybe<Array<Scalars['String']['input']>>;
  /** 默认域名 */
  domain?: InputMaybe<Scalars['String']['input']>;
  fileIdentityIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** 组织本位币 */
  localCurrency?: InputMaybe<Scalars['String']['input']>;
  /** 组织图标 */
  logo?: InputMaybe<OrgLogoInput>;
  /** 组织名称 */
  name: Scalars['String']['input'];
  ownerID?: InputMaybe<Scalars['ID']['input']>;
  parentID: Scalars['ID']['input'];
  permissionIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  policyIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** 简介 */
  profile?: InputMaybe<Scalars['String']['input']>;
  rolesAndGroupIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** 状态 */
  status?: InputMaybe<OrgSimpleStatus>;
  /** 时区 */
  timezone?: InputMaybe<Scalars['String']['input']>;
  userIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  userPasswordPolicyID?: InputMaybe<Scalars['ID']['input']>;
};

/**
 * CreateOrgPolicyInput is used for create OrgPolicy object.
 * Input was generated by ent.
 */
export type CreateOrgPolicyInput = {
  appID?: InputMaybe<Scalars['ID']['input']>;
  appPolicyID?: InputMaybe<Scalars['ID']['input']>;
  /** 描述 */
  comments?: InputMaybe<Scalars['String']['input']>;
  /** 策略名称 */
  name: Scalars['String']['input'];
  orgID?: InputMaybe<Scalars['ID']['input']>;
  permissionIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** 策略规则 */
  rules: Array<InputMaybe<PolicyRuleInput>>;
};

/**
 * CreateOrgRoleInput is used for create OrgRole object.
 * Input was generated by ent.
 */
export type CreateOrgRoleInput = {
  /** 备注 */
  comments?: InputMaybe<Scalars['String']['input']>;
  /** 类型,group:组,role:角色 */
  kind: OrgRoleKind;
  /** 名称 */
  name: Scalars['String']['input'];
  orgID?: InputMaybe<Scalars['ID']['input']>;
};

/**
 * CreateOrgUserInput is used for create OrgUser object.
 * Input was generated by ent.
 */
export type CreateOrgUserInput = {
  /** 在组织内的显示名称 */
  displayName: Scalars['String']['input'];
  /** 加入时间 */
  joinedAt?: InputMaybe<Scalars['Time']['input']>;
  orgID: Scalars['ID']['input'];
  userID: Scalars['ID']['input'];
  /** 用户类型，区分内部及外部用户 */
  userType?: InputMaybe<OrgUserUserType>;
};

/**
 * CreateOrgUserPreferenceInput is used for create OrgUserPreference object.
 * Input was generated by ent.
 */
export type CreateOrgUserPreferenceInput = {
  /** 用户收藏菜单 */
  menuFavorite?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** 用户最近访问菜单 */
  menuRecent?: InputMaybe<Array<Scalars['ID']['input']>>;
};

/**
 * CreatePermissionInput is used for create Permission object.
 * Input was generated by ent.
 */
export type CreatePermissionInput = {
  /** 生效结束时间 */
  endAt?: InputMaybe<Scalars['Time']['input']>;
  orgID: Scalars['ID']['input'];
  orgPolicyID: Scalars['ID']['input'];
  /** 授权类型:角色,用户 */
  principalKind: PermissionPrincipalKind;
  roleID?: InputMaybe<Scalars['ID']['input']>;
  /** 生效开始时间 */
  startAt?: InputMaybe<Scalars['Time']['input']>;
  userID?: InputMaybe<Scalars['ID']['input']>;
};

/**
 * CreateQuotaInput is used for create Quota object.
 * Input was generated by ent.
 */
export type CreateQuotaInput = {
  /** 过期时间 */
  endAt?: InputMaybe<Scalars['Time']['input']>;
  /** 限制值 */
  limit: Scalars['Int']['input'];
  quotaItemID: Scalars['ID']['input'];
  quotaOrgID?: InputMaybe<Scalars['ID']['input']>;
  quotaUserID?: InputMaybe<Scalars['ID']['input']>;
  /** 生效时间 */
  startAt?: InputMaybe<Scalars['Time']['input']>;
};

/**
 * CreateQuotaItemInput is used for create QuotaItem object.
 * Input was generated by ent.
 */
export type CreateQuotaItemInput = {
  /** 是否启用 */
  active?: InputMaybe<Scalars['Boolean']['input']>;
  /** 配额项代码,如: users,orgs */
  code: Scalars['String']['input'];
  /** 默认限制值 */
  defaultLimit?: InputMaybe<Scalars['Int']['input']>;
  /** 描述 */
  description?: InputMaybe<Scalars['String']['input']>;
  /** 配额项名称 */
  name: Scalars['String']['input'];
  quotumIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** 资源类型 */
  resourceType: QuotaItemResourceType;
  /** 单位,如: 个,MB,GB */
  unit?: InputMaybe<Scalars['String']['input']>;
};

/**
 * CreateRegionInput is used for create Region object.
 * Input was generated by ent.
 */
export type CreateRegionInput = {
  childIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  countryID?: InputMaybe<Scalars['ID']['input']>;
  /** 地区中文名称 */
  name?: InputMaybe<Scalars['String']['input']>;
  /** 地区英文名称 */
  nameEn?: InputMaybe<Scalars['String']['input']>;
  parentID?: InputMaybe<Scalars['ID']['input']>;
  /** 编码 */
  shortCode?: InputMaybe<Scalars['String']['input']>;
  /** 状态 */
  status?: InputMaybe<RegionSimpleStatus>;
  /** 邮政编码 */
  zipCode?: InputMaybe<Scalars['String']['input']>;
};

/**
 * CreateUserAddrInput is used for create UserAddr object.
 * Input was generated by ent.
 */
export type CreateUserAddrInput = {
  /** 详细地址 */
  addr?: InputMaybe<Scalars['String']['input']>;
  /** 邮箱 */
  email?: InputMaybe<Scalars['String']['input']>;
  /** 传真 */
  fax?: InputMaybe<Scalars['String']['input']>;
  /** 是否默认地址，类型为delivery时使用 */
  isDefault?: InputMaybe<Scalars['Boolean']['input']>;
  /** 手机 */
  mobile?: InputMaybe<Scalars['String']['input']>;
  /** 联系人名称 */
  name?: InputMaybe<Scalars['String']['input']>;
  regionID?: InputMaybe<Scalars['ID']['input']>;
  /** 电话 */
  tel?: InputMaybe<Scalars['String']['input']>;
  userID?: InputMaybe<Scalars['ID']['input']>;
  /** 邮编 */
  zipCode?: InputMaybe<Scalars['String']['input']>;
};

/**
 * CreateUserDeviceInput is used for create UserDevice object.
 * Input was generated by ent.
 */
export type CreateUserDeviceInput = {
  /** app版本 */
  appVersion?: InputMaybe<Scalars['String']['input']>;
  /** 备注 */
  comments?: InputMaybe<Scalars['String']['input']>;
  /** 设备型号 */
  deviceModel?: InputMaybe<Scalars['String']['input']>;
  /** 设备名称 */
  deviceName?: InputMaybe<Scalars['String']['input']>;
  /** 设备唯一ID */
  deviceUID: Scalars['String']['input'];
  /** 状态,可用或不可用及其他待确认状态 */
  status?: InputMaybe<UserDeviceSimpleStatus>;
  /** 系统名称 */
  systemName?: InputMaybe<Scalars['String']['input']>;
  /** 系统版本 */
  systemVersion?: InputMaybe<Scalars['String']['input']>;
  userID?: InputMaybe<Scalars['ID']['input']>;
};

/**
 * CreateUserIdentityInput is used for create UserIdentity object.
 * Input was generated by ent.
 */
export type CreateUserIdentityInput = {
  /** 用户名、邮箱、手机、unionid、qq */
  code?: InputMaybe<Scalars['String']['input']>;
  /** 扩展标识码,比如微信的openID */
  codeExtend?: InputMaybe<Scalars['String']['input']>;
  /** 身份标识类型 手机、邮箱、用户名、微信、qq */
  kind: UserIdentityKind;
  /** 状态,部分登陆方式需要验证通过才可启用 */
  status?: InputMaybe<UserIdentitySimpleStatus>;
  userID?: InputMaybe<Scalars['ID']['input']>;
};

/**
 * CreateUserInput is used for create User object.
 * Input was generated by ent.
 */
export type CreateUserInput = {
  addressIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** 头像地址 */
  avatar?: InputMaybe<Scalars['String']['input']>;
  citizenshipID?: InputMaybe<Scalars['ID']['input']>;
  /** 备注 */
  comments?: InputMaybe<Scalars['String']['input']>;
  /** 地址信息 */
  contact?: InputMaybe<CreateUserAddrInput>;
  deviceIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** 显示名 */
  displayName: Scalars['String']['input'];
  /** 名字 */
  firstName?: InputMaybe<Scalars['String']['input']>;
  /** 性别 */
  gender?: InputMaybe<UserGender>;
  identityIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** 语言 */
  lang?: InputMaybe<Scalars['String']['input']>;
  /** 姓氏 */
  lastName?: InputMaybe<Scalars['String']['input']>;
  loginProfile?: InputMaybe<CreateUserLoginProfileInput>;
  loginProfileID?: InputMaybe<Scalars['ID']['input']>;
  /** 中间名 */
  middleName?: InputMaybe<Scalars['String']['input']>;
  oauthClientIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** 如指定密码则填入,否则由系统自动生成密码 */
  password?: InputMaybe<CreateUserPasswordInput>;
  passwordIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** 登陆名称 */
  principalName: Scalars['String']['input'];
  /** 状态 */
  status?: InputMaybe<UserSimpleStatus>;
};

/**
 * CreateUserLoginProfileInput is used for create UserLoginProfile object.
 * Input was generated by ent.
 */
export type CreateUserLoginProfileInput = {
  /** 是否允许使用密码登陆控制台 */
  canLogin?: InputMaybe<Scalars['Boolean']['input']>;
  /** 下次登陆时需要重置密码 */
  passwordReset?: InputMaybe<Scalars['Boolean']['input']>;
  /** 设置密码:keep-保持不变,customer-客户自行设置,auto-自动生成 */
  setKind: UserLoginProfileSetKind;
  userID?: InputMaybe<Scalars['ID']['input']>;
  /** 是否开启设备认证 */
  verifyDevice: Scalars['Boolean']['input'];
};

/**
 * CreateUserPasswordInput is used for create UserPassword object.
 * Input was generated by ent.
 */
export type CreateUserPasswordInput = {
  /** 密码 */
  password?: InputMaybe<Scalars['String']['input']>;
  /** 场景: login 普通登陆 */
  scene: UserPasswordScene;
  /** 生效状态,默认生效 */
  status?: InputMaybe<UserPasswordSimpleStatus>;
  userID?: InputMaybe<Scalars['ID']['input']>;
};

/**
 * CreateUserPasswordPolicyInput is used for create UserPasswordPolicy object.
 * Input was generated by ent.
 */
export type CreateUserPasswordPolicyInput = {
  /** 是否允许包含用户名 */
  allowIncludeUserName?: InputMaybe<Scalars['Boolean']['input']>;
  /** 密码错误多少次出现验证码，最大5次，0代表不出现验证码 */
  captchaTimes?: InputMaybe<Scalars['Int']['input']>;
  /** 最少包含的不同字符数，最多8个，0代表不限制 */
  includeChar?: InputMaybe<Scalars['Int']['input']>;
  /** 必须包含的元素，异或：1-小写字母，2-大写字母，4-数字，8-符号 */
  includeElement?: InputMaybe<Scalars['Int']['input']>;
  /** 有效天数，最大1095天，0代表不过期 */
  invalidDay?: InputMaybe<Scalars['Int']['input']>;
  /** 过期后是否限制登录 */
  invalidLoginLimit?: InputMaybe<Scalars['Boolean']['input']>;
  /** 密码最短长度，长度应在6-32位之间 */
  length?: InputMaybe<Scalars['Int']['input']>;
  orgID?: InputMaybe<Scalars['ID']['input']>;
  /** 一小时内密码错误最多尝试次数，最大32次，0代表不限次数 */
  retry?: InputMaybe<Scalars['Int']['input']>;
};

export type Currency = Node & {
  __typename?: 'Currency';
  /** 代码 */
  code: Scalars['String']['output'];
  createdAt: Scalars['Time']['output'];
  createdBy: Scalars['Int']['output'];
  id: Scalars['ID']['output'];
  /** 中文名称 */
  name: Scalars['String']['output'];
  /** 货币符号 */
  sign?: Maybe<Scalars['String']['output']>;
  /** 状态 */
  status: CurrencySimpleStatus;
  updatedAt?: Maybe<Scalars['Time']['output']>;
  updatedBy?: Maybe<Scalars['Int']['output']>;
};

/** A connection to a list of items. */
export type CurrencyConnection = {
  __typename?: 'CurrencyConnection';
  /** A list of edges. */
  edges?: Maybe<Array<Maybe<CurrencyEdge>>>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** Identifies the total count of items in the connection. */
  totalCount: Scalars['Int']['output'];
};

/** An edge in a connection. */
export type CurrencyEdge = {
  __typename?: 'CurrencyEdge';
  /** A cursor for use in pagination. */
  cursor: Scalars['Cursor']['output'];
  /** The item at the end of the edge. */
  node?: Maybe<Currency>;
};

/** Ordering options for Currency connections */
export type CurrencyOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order Currencies. */
  field: CurrencyOrderField;
};

/** Properties by which Currency connections can be ordered. */
export enum CurrencyOrderField {
  CreatedAt = 'createdAt'
}

/** CurrencySimpleStatus is enum for the field status */
export enum CurrencySimpleStatus {
  Active = 'active',
  Disabled = 'disabled',
  Inactive = 'inactive',
  Processing = 'processing'
}

/**
 * CurrencyWhereInput is used for filtering Currency objects.
 * Input was generated by ent.
 */
export type CurrencyWhereInput = {
  and?: InputMaybe<Array<CurrencyWhereInput>>;
  /** code field predicates */
  code?: InputMaybe<Scalars['String']['input']>;
  codeContains?: InputMaybe<Scalars['String']['input']>;
  codeContainsFold?: InputMaybe<Scalars['String']['input']>;
  codeEqualFold?: InputMaybe<Scalars['String']['input']>;
  codeGT?: InputMaybe<Scalars['String']['input']>;
  codeGTE?: InputMaybe<Scalars['String']['input']>;
  codeHasPrefix?: InputMaybe<Scalars['String']['input']>;
  codeHasSuffix?: InputMaybe<Scalars['String']['input']>;
  codeIn?: InputMaybe<Array<Scalars['String']['input']>>;
  codeLT?: InputMaybe<Scalars['String']['input']>;
  codeLTE?: InputMaybe<Scalars['String']['input']>;
  codeNEQ?: InputMaybe<Scalars['String']['input']>;
  codeNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  /** created_at field predicates */
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  createdAtGT?: InputMaybe<Scalars['Time']['input']>;
  createdAtGTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  createdAtLT?: InputMaybe<Scalars['Time']['input']>;
  createdAtLTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  createdAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  /** created_by field predicates */
  createdBy?: InputMaybe<Scalars['Int']['input']>;
  createdByGT?: InputMaybe<Scalars['Int']['input']>;
  createdByGTE?: InputMaybe<Scalars['Int']['input']>;
  createdByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  createdByLT?: InputMaybe<Scalars['Int']['input']>;
  createdByLTE?: InputMaybe<Scalars['Int']['input']>;
  createdByNEQ?: InputMaybe<Scalars['Int']['input']>;
  createdByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** name field predicates */
  name?: InputMaybe<Scalars['String']['input']>;
  nameContains?: InputMaybe<Scalars['String']['input']>;
  nameContainsFold?: InputMaybe<Scalars['String']['input']>;
  nameEqualFold?: InputMaybe<Scalars['String']['input']>;
  nameGT?: InputMaybe<Scalars['String']['input']>;
  nameGTE?: InputMaybe<Scalars['String']['input']>;
  nameHasPrefix?: InputMaybe<Scalars['String']['input']>;
  nameHasSuffix?: InputMaybe<Scalars['String']['input']>;
  nameIn?: InputMaybe<Array<Scalars['String']['input']>>;
  nameLT?: InputMaybe<Scalars['String']['input']>;
  nameLTE?: InputMaybe<Scalars['String']['input']>;
  nameNEQ?: InputMaybe<Scalars['String']['input']>;
  nameNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  not?: InputMaybe<CurrencyWhereInput>;
  or?: InputMaybe<Array<CurrencyWhereInput>>;
  /** sign field predicates */
  sign?: InputMaybe<Scalars['String']['input']>;
  signContains?: InputMaybe<Scalars['String']['input']>;
  signContainsFold?: InputMaybe<Scalars['String']['input']>;
  signEqualFold?: InputMaybe<Scalars['String']['input']>;
  signGT?: InputMaybe<Scalars['String']['input']>;
  signGTE?: InputMaybe<Scalars['String']['input']>;
  signHasPrefix?: InputMaybe<Scalars['String']['input']>;
  signHasSuffix?: InputMaybe<Scalars['String']['input']>;
  signIn?: InputMaybe<Array<Scalars['String']['input']>>;
  signIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  signLT?: InputMaybe<Scalars['String']['input']>;
  signLTE?: InputMaybe<Scalars['String']['input']>;
  signNEQ?: InputMaybe<Scalars['String']['input']>;
  signNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  signNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** status field predicates */
  status?: InputMaybe<CurrencySimpleStatus>;
  statusIn?: InputMaybe<Array<CurrencySimpleStatus>>;
  statusNEQ?: InputMaybe<CurrencySimpleStatus>;
  statusNotIn?: InputMaybe<Array<CurrencySimpleStatus>>;
  /** updated_at field predicates */
  updatedAt?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedAtLT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtLTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_by field predicates */
  updatedBy?: InputMaybe<Scalars['Int']['input']>;
  updatedByGT?: InputMaybe<Scalars['Int']['input']>;
  updatedByGTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedByLT?: InputMaybe<Scalars['Int']['input']>;
  updatedByLTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByNEQ?: InputMaybe<Scalars['Int']['input']>;
  updatedByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByNotNil?: InputMaybe<Scalars['Boolean']['input']>;
};

export type EnableDirectoryInput = {
  /** 域名 */
  domain: Scalars['String']['input'];
  name: Scalars['String']['input'];
};

export type FileIdentity = Node & {
  __typename?: 'FileIdentity';
  /** accesskey id */
  accessKeyID: Scalars['String']['output'];
  /** 备注 */
  comments?: Maybe<Scalars['String']['output']>;
  createdAt: Scalars['Time']['output'];
  createdBy: Scalars['Int']['output'];
  /** STS令牌的有效期，默认3600s */
  durationSeconds?: Maybe<Scalars['Int']['output']>;
  /** 文件来源ID */
  fileSourceID: Scalars['ID']['output'];
  id: Scalars['ID']['output'];
  /** 租户默认的凭证 */
  isDefault: Scalars['Boolean']['output'];
  org: Org;
  /** 指定返回的STS令牌的权限的策略 */
  policy?: Maybe<Scalars['String']['output']>;
  /** 角色的资源名称(ARN)，用于STS */
  roleArn: Scalars['String']['output'];
  source: FileSource;
  /** 租户ID */
  tenantID: Scalars['ID']['output'];
  updatedAt?: Maybe<Scalars['Time']['output']>;
  updatedBy?: Maybe<Scalars['Int']['output']>;
};

/** A connection to a list of items. */
export type FileIdentityConnection = {
  __typename?: 'FileIdentityConnection';
  /** A list of edges. */
  edges?: Maybe<Array<Maybe<FileIdentityEdge>>>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** Identifies the total count of items in the connection. */
  totalCount: Scalars['Int']['output'];
};

/** An edge in a connection. */
export type FileIdentityEdge = {
  __typename?: 'FileIdentityEdge';
  /** A cursor for use in pagination. */
  cursor: Scalars['Cursor']['output'];
  /** The item at the end of the edge. */
  node?: Maybe<FileIdentity>;
};

/** 内部调用fileIdentity */
export type FileIdentityForApp = Node & {
  __typename?: 'FileIdentityForApp';
  accessKeyID: Scalars['String']['output'];
  accessKeySecret: Scalars['String']['output'];
  /** STS令牌的有效期，默认3600s */
  durationSeconds?: Maybe<Scalars['Int']['output']>;
  id: Scalars['ID']['output'];
  /** 租户默认的凭证 */
  isDefault: Scalars['Boolean']['output'];
  /** 指定返回的STS令牌的权限的策略 */
  policy?: Maybe<Scalars['String']['output']>;
  /** 角色的资源名称(ARN)，用于STS */
  roleArn: Scalars['String']['output'];
  source: FileSource;
  tenantID: Scalars['ID']['output'];
};

/** Ordering options for FileIdentity connections */
export type FileIdentityOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order FileIdentities. */
  field: FileIdentityOrderField;
};

/** Properties by which FileIdentity connections can be ordered. */
export enum FileIdentityOrderField {
  CreatedAt = 'createdAt'
}

/**
 * FileIdentityWhereInput is used for filtering FileIdentity objects.
 * Input was generated by ent.
 */
export type FileIdentityWhereInput = {
  /** access_key_id field predicates */
  accessKeyID?: InputMaybe<Scalars['String']['input']>;
  accessKeyIDContains?: InputMaybe<Scalars['String']['input']>;
  accessKeyIDContainsFold?: InputMaybe<Scalars['String']['input']>;
  accessKeyIDEqualFold?: InputMaybe<Scalars['String']['input']>;
  accessKeyIDGT?: InputMaybe<Scalars['String']['input']>;
  accessKeyIDGTE?: InputMaybe<Scalars['String']['input']>;
  accessKeyIDHasPrefix?: InputMaybe<Scalars['String']['input']>;
  accessKeyIDHasSuffix?: InputMaybe<Scalars['String']['input']>;
  accessKeyIDIn?: InputMaybe<Array<Scalars['String']['input']>>;
  accessKeyIDLT?: InputMaybe<Scalars['String']['input']>;
  accessKeyIDLTE?: InputMaybe<Scalars['String']['input']>;
  accessKeyIDNEQ?: InputMaybe<Scalars['String']['input']>;
  accessKeyIDNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  and?: InputMaybe<Array<FileIdentityWhereInput>>;
  /** created_at field predicates */
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  createdAtGT?: InputMaybe<Scalars['Time']['input']>;
  createdAtGTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  createdAtLT?: InputMaybe<Scalars['Time']['input']>;
  createdAtLTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  createdAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  /** created_by field predicates */
  createdBy?: InputMaybe<Scalars['Int']['input']>;
  createdByGT?: InputMaybe<Scalars['Int']['input']>;
  createdByGTE?: InputMaybe<Scalars['Int']['input']>;
  createdByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  createdByLT?: InputMaybe<Scalars['Int']['input']>;
  createdByLTE?: InputMaybe<Scalars['Int']['input']>;
  createdByNEQ?: InputMaybe<Scalars['Int']['input']>;
  createdByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  /** duration_seconds field predicates */
  durationSeconds?: InputMaybe<Scalars['Int']['input']>;
  durationSecondsGT?: InputMaybe<Scalars['Int']['input']>;
  durationSecondsGTE?: InputMaybe<Scalars['Int']['input']>;
  durationSecondsIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  durationSecondsIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  durationSecondsLT?: InputMaybe<Scalars['Int']['input']>;
  durationSecondsLTE?: InputMaybe<Scalars['Int']['input']>;
  durationSecondsNEQ?: InputMaybe<Scalars['Int']['input']>;
  durationSecondsNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  durationSecondsNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** file_source_id field predicates */
  fileSourceID?: InputMaybe<Scalars['ID']['input']>;
  fileSourceIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  fileSourceIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  fileSourceIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** org edge predicates */
  hasOrg?: InputMaybe<Scalars['Boolean']['input']>;
  hasOrgWith?: InputMaybe<Array<OrgWhereInput>>;
  /** source edge predicates */
  hasSource?: InputMaybe<Scalars['Boolean']['input']>;
  hasSourceWith?: InputMaybe<Array<FileSourceWhereInput>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** is_default field predicates */
  isDefault?: InputMaybe<Scalars['Boolean']['input']>;
  isDefaultNEQ?: InputMaybe<Scalars['Boolean']['input']>;
  not?: InputMaybe<FileIdentityWhereInput>;
  or?: InputMaybe<Array<FileIdentityWhereInput>>;
  /** policy field predicates */
  policy?: InputMaybe<Scalars['String']['input']>;
  policyContains?: InputMaybe<Scalars['String']['input']>;
  policyContainsFold?: InputMaybe<Scalars['String']['input']>;
  policyEqualFold?: InputMaybe<Scalars['String']['input']>;
  policyGT?: InputMaybe<Scalars['String']['input']>;
  policyGTE?: InputMaybe<Scalars['String']['input']>;
  policyHasPrefix?: InputMaybe<Scalars['String']['input']>;
  policyHasSuffix?: InputMaybe<Scalars['String']['input']>;
  policyIn?: InputMaybe<Array<Scalars['String']['input']>>;
  policyIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  policyLT?: InputMaybe<Scalars['String']['input']>;
  policyLTE?: InputMaybe<Scalars['String']['input']>;
  policyNEQ?: InputMaybe<Scalars['String']['input']>;
  policyNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  policyNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** role_arn field predicates */
  roleArn?: InputMaybe<Scalars['String']['input']>;
  roleArnContains?: InputMaybe<Scalars['String']['input']>;
  roleArnContainsFold?: InputMaybe<Scalars['String']['input']>;
  roleArnEqualFold?: InputMaybe<Scalars['String']['input']>;
  roleArnGT?: InputMaybe<Scalars['String']['input']>;
  roleArnGTE?: InputMaybe<Scalars['String']['input']>;
  roleArnHasPrefix?: InputMaybe<Scalars['String']['input']>;
  roleArnHasSuffix?: InputMaybe<Scalars['String']['input']>;
  roleArnIn?: InputMaybe<Array<Scalars['String']['input']>>;
  roleArnLT?: InputMaybe<Scalars['String']['input']>;
  roleArnLTE?: InputMaybe<Scalars['String']['input']>;
  roleArnNEQ?: InputMaybe<Scalars['String']['input']>;
  roleArnNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  /** tenant_id field predicates */
  tenantID?: InputMaybe<Scalars['ID']['input']>;
  tenantIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  tenantIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  tenantIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** updated_at field predicates */
  updatedAt?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedAtLT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtLTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_by field predicates */
  updatedBy?: InputMaybe<Scalars['Int']['input']>;
  updatedByGT?: InputMaybe<Scalars['Int']['input']>;
  updatedByGTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedByLT?: InputMaybe<Scalars['Int']['input']>;
  updatedByLTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByNEQ?: InputMaybe<Scalars['Int']['input']>;
  updatedByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByNotNil?: InputMaybe<Scalars['Boolean']['input']>;
};

export type FileSource = Node & {
  __typename?: 'FileSource';
  /** 文件存储空间 */
  bucket: Scalars['String']['output'];
  /** 文件存储空间地址，用于匹配url */
  bucketURL: Scalars['String']['output'];
  /** 备注 */
  comments?: Maybe<Scalars['String']['output']>;
  createdAt: Scalars['Time']['output'];
  createdBy: Scalars['Int']['output'];
  /** 对外服务的访问域名 */
  endpoint: Scalars['String']['output'];
  /** 是否禁止修改endpoint，如果是自定义域名设为true */
  endpointImmutable: Scalars['Boolean']['output'];
  id: Scalars['ID']['output'];
  /** 文件来源 */
  kind: FileSourceKind;
  /** 地域，数据存储的物理位置 */
  region: Scalars['String']['output'];
  /** sts服务的访问域名 */
  stsEndpoint: Scalars['String']['output'];
  updatedAt?: Maybe<Scalars['Time']['output']>;
  updatedBy?: Maybe<Scalars['Int']['output']>;
};

/** A connection to a list of items. */
export type FileSourceConnection = {
  __typename?: 'FileSourceConnection';
  /** A list of edges. */
  edges?: Maybe<Array<Maybe<FileSourceEdge>>>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** Identifies the total count of items in the connection. */
  totalCount: Scalars['Int']['output'];
};

/** An edge in a connection. */
export type FileSourceEdge = {
  __typename?: 'FileSourceEdge';
  /** A cursor for use in pagination. */
  cursor: Scalars['Cursor']['output'];
  /** The item at the end of the edge. */
  node?: Maybe<FileSource>;
};

/** FileSourceKind is enum for the field kind */
export enum FileSourceKind {
  AliOss = 'aliOSS',
  AwsS3 = 'awsS3',
  Minio = 'minio'
}

/** Ordering options for FileSource connections */
export type FileSourceOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order FileSources. */
  field: FileSourceOrderField;
};

/** Properties by which FileSource connections can be ordered. */
export enum FileSourceOrderField {
  CreatedAt = 'createdAt'
}

/**
 * FileSourceWhereInput is used for filtering FileSource objects.
 * Input was generated by ent.
 */
export type FileSourceWhereInput = {
  and?: InputMaybe<Array<FileSourceWhereInput>>;
  /** bucket field predicates */
  bucket?: InputMaybe<Scalars['String']['input']>;
  bucketContains?: InputMaybe<Scalars['String']['input']>;
  bucketContainsFold?: InputMaybe<Scalars['String']['input']>;
  bucketEqualFold?: InputMaybe<Scalars['String']['input']>;
  bucketGT?: InputMaybe<Scalars['String']['input']>;
  bucketGTE?: InputMaybe<Scalars['String']['input']>;
  bucketHasPrefix?: InputMaybe<Scalars['String']['input']>;
  bucketHasSuffix?: InputMaybe<Scalars['String']['input']>;
  bucketIn?: InputMaybe<Array<Scalars['String']['input']>>;
  bucketLT?: InputMaybe<Scalars['String']['input']>;
  bucketLTE?: InputMaybe<Scalars['String']['input']>;
  bucketNEQ?: InputMaybe<Scalars['String']['input']>;
  bucketNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  /** bucket_url field predicates */
  bucketURL?: InputMaybe<Scalars['String']['input']>;
  bucketURLContains?: InputMaybe<Scalars['String']['input']>;
  bucketURLContainsFold?: InputMaybe<Scalars['String']['input']>;
  bucketURLEqualFold?: InputMaybe<Scalars['String']['input']>;
  bucketURLGT?: InputMaybe<Scalars['String']['input']>;
  bucketURLGTE?: InputMaybe<Scalars['String']['input']>;
  bucketURLHasPrefix?: InputMaybe<Scalars['String']['input']>;
  bucketURLHasSuffix?: InputMaybe<Scalars['String']['input']>;
  bucketURLIn?: InputMaybe<Array<Scalars['String']['input']>>;
  bucketURLLT?: InputMaybe<Scalars['String']['input']>;
  bucketURLLTE?: InputMaybe<Scalars['String']['input']>;
  bucketURLNEQ?: InputMaybe<Scalars['String']['input']>;
  bucketURLNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  /** created_at field predicates */
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  createdAtGT?: InputMaybe<Scalars['Time']['input']>;
  createdAtGTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  createdAtLT?: InputMaybe<Scalars['Time']['input']>;
  createdAtLTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  createdAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  /** created_by field predicates */
  createdBy?: InputMaybe<Scalars['Int']['input']>;
  createdByGT?: InputMaybe<Scalars['Int']['input']>;
  createdByGTE?: InputMaybe<Scalars['Int']['input']>;
  createdByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  createdByLT?: InputMaybe<Scalars['Int']['input']>;
  createdByLTE?: InputMaybe<Scalars['Int']['input']>;
  createdByNEQ?: InputMaybe<Scalars['Int']['input']>;
  createdByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  /** endpoint field predicates */
  endpoint?: InputMaybe<Scalars['String']['input']>;
  endpointContains?: InputMaybe<Scalars['String']['input']>;
  endpointContainsFold?: InputMaybe<Scalars['String']['input']>;
  endpointEqualFold?: InputMaybe<Scalars['String']['input']>;
  endpointGT?: InputMaybe<Scalars['String']['input']>;
  endpointGTE?: InputMaybe<Scalars['String']['input']>;
  endpointHasPrefix?: InputMaybe<Scalars['String']['input']>;
  endpointHasSuffix?: InputMaybe<Scalars['String']['input']>;
  /** endpoint_immutable field predicates */
  endpointImmutable?: InputMaybe<Scalars['Boolean']['input']>;
  endpointImmutableNEQ?: InputMaybe<Scalars['Boolean']['input']>;
  endpointIn?: InputMaybe<Array<Scalars['String']['input']>>;
  endpointLT?: InputMaybe<Scalars['String']['input']>;
  endpointLTE?: InputMaybe<Scalars['String']['input']>;
  endpointNEQ?: InputMaybe<Scalars['String']['input']>;
  endpointNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  /** identities edge predicates */
  hasIdentities?: InputMaybe<Scalars['Boolean']['input']>;
  hasIdentitiesWith?: InputMaybe<Array<FileIdentityWhereInput>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** kind field predicates */
  kind?: InputMaybe<FileSourceKind>;
  kindIn?: InputMaybe<Array<FileSourceKind>>;
  kindNEQ?: InputMaybe<FileSourceKind>;
  kindNotIn?: InputMaybe<Array<FileSourceKind>>;
  not?: InputMaybe<FileSourceWhereInput>;
  or?: InputMaybe<Array<FileSourceWhereInput>>;
  /** region field predicates */
  region?: InputMaybe<Scalars['String']['input']>;
  regionContains?: InputMaybe<Scalars['String']['input']>;
  regionContainsFold?: InputMaybe<Scalars['String']['input']>;
  regionEqualFold?: InputMaybe<Scalars['String']['input']>;
  regionGT?: InputMaybe<Scalars['String']['input']>;
  regionGTE?: InputMaybe<Scalars['String']['input']>;
  regionHasPrefix?: InputMaybe<Scalars['String']['input']>;
  regionHasSuffix?: InputMaybe<Scalars['String']['input']>;
  regionIn?: InputMaybe<Array<Scalars['String']['input']>>;
  regionLT?: InputMaybe<Scalars['String']['input']>;
  regionLTE?: InputMaybe<Scalars['String']['input']>;
  regionNEQ?: InputMaybe<Scalars['String']['input']>;
  regionNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  /** sts_endpoint field predicates */
  stsEndpoint?: InputMaybe<Scalars['String']['input']>;
  stsEndpointContains?: InputMaybe<Scalars['String']['input']>;
  stsEndpointContainsFold?: InputMaybe<Scalars['String']['input']>;
  stsEndpointEqualFold?: InputMaybe<Scalars['String']['input']>;
  stsEndpointGT?: InputMaybe<Scalars['String']['input']>;
  stsEndpointGTE?: InputMaybe<Scalars['String']['input']>;
  stsEndpointHasPrefix?: InputMaybe<Scalars['String']['input']>;
  stsEndpointHasSuffix?: InputMaybe<Scalars['String']['input']>;
  stsEndpointIn?: InputMaybe<Array<Scalars['String']['input']>>;
  stsEndpointLT?: InputMaybe<Scalars['String']['input']>;
  stsEndpointLTE?: InputMaybe<Scalars['String']['input']>;
  stsEndpointNEQ?: InputMaybe<Scalars['String']['input']>;
  stsEndpointNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  /** updated_at field predicates */
  updatedAt?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedAtLT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtLTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_by field predicates */
  updatedBy?: InputMaybe<Scalars['Int']['input']>;
  updatedByGT?: InputMaybe<Scalars['Int']['input']>;
  updatedByGTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedByLT?: InputMaybe<Scalars['Int']['input']>;
  updatedByLTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByNEQ?: InputMaybe<Scalars['Int']['input']>;
  updatedByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByNotNil?: InputMaybe<Scalars['Boolean']['input']>;
};

export type GrantInput = {
  orgScope: Scalars['ID']['input'];
  policyID: Scalars['ID']['input'];
  principal: Scalars['GID']['input'];
};

/** 列表操作类型 */
export enum ListAction {
  /** 下移 */
  Down = 'down',
  /** 上移 */
  Up = 'up'
}

export type Mfa = {
  __typename?: 'Mfa';
  account: Scalars['String']['output'];
  secret: Scalars['String']['output'];
};

export type Mutation = {
  __typename?: 'Mutation';
  /** 将用户分配到组织下 */
  allotOrganizationUser: Scalars['Boolean']['output'];
  /** 角色添加策略 */
  assignAppRolePolicy: Scalars['Boolean']['output'];
  /** 应用角色添加策略视图权限 */
  assignAppRolePolicyView: Scalars['Boolean']['output'];
  /** 组织角色添加策略视图权限 */
  assignOrgRolePolicyView: Scalars['Boolean']['output'];
  /** 组织用户添加策略视图权限 */
  assignOrgUserPolicyView: Scalars['Boolean']['output'];
  /** 分配应用,将自动分配应用下的所有资源 */
  assignOrganizationApp: Scalars['Boolean']['output'];
  /** 分配应用策略到组织 */
  assignOrganizationAppPolicy: Scalars['Boolean']['output'];
  /** 分配应用角色到组织 */
  assignOrganizationAppRole: Scalars['Boolean']['output'];
  /** 分配组用户 */
  assignRoleUser: Scalars['Boolean']['output'];
  /** 自动授权应用，系统开户、创建web交易用户使用 */
  autoGrantApp: Scalars['Boolean']['output'];
  /** 绑定用户凭证(管理端使用) */
  bindUserIdentity?: Maybe<UserIdentity>;
  /** 更改组织用户类型 */
  changeOrgUserType: Scalars['Boolean']['output'];
  /** 用户修改自己的登录密码 */
  changePassword: Scalars['Boolean']['output'];
  /** 创建应用 */
  createApp?: Maybe<App>;
  /** 创建应用操作 */
  createAppActions: Array<Maybe<AppAction>>;
  /** 创建应用数据字典 */
  createAppDict?: Maybe<AppDict>;
  /** 创建应用数据字典项 */
  createAppDictItem?: Maybe<AppDictItem>;
  /** 创建应用菜单 */
  createAppMenus: Array<Maybe<AppMenu>>;
  /** 创建应用策略模板 */
  createAppPolicy?: Maybe<AppPolicy>;
  /** 创建策略视图 */
  createAppPolicyView: AppPolicyView;
  /** 创建应用角色 */
  createAppRole?: Maybe<AppRole>;
  /** 创建国家 */
  createCountry?: Maybe<Country>;
  /** 创建货币 */
  createCurrency?: Maybe<Currency>;
  /** 创建文件凭证 */
  createFileIdentity: FileIdentity;
  /** 创建文件来源 */
  createFileSource: FileSource;
  /** 创建用户 AccessKey */
  createOauthClient: OauthClient;
  /** 创建组织目录 */
  createOrganization?: Maybe<Org>;
  /** 创建组织成员(管理账户) */
  createOrganizationAccount?: Maybe<User>;
  /** 创建组织策略 */
  createOrganizationPolicy?: Maybe<OrgPolicy>;
  /** 创建组织用户 */
  createOrganizationUser?: Maybe<User>;
  /** 设置配额 */
  createQuota: Quota;
  /** 创建配额项 */
  createQuotaItem: QuotaItem;
  /** 创建地区 */
  createRegion?: Maybe<Region>;
  /** 创建角色或组 */
  createRole?: Maybe<OrgRole>;
  /** 创建组织根节点(管理端使用) */
  createRoot?: Maybe<Org>;
  /** 设置密码策略 */
  createUserPasswordPolicy?: Maybe<UserPasswordPolicy>;
  /** 删除应用 */
  deleteApp: Scalars['Boolean']['output'];
  /** 删除应用操作 */
  deleteAppAction: Scalars['Boolean']['output'];
  /** 删除应用数据字典 */
  deleteAppDict: Scalars['Boolean']['output'];
  /** 删除应用数据字典项 */
  deleteAppDictItem: Scalars['Boolean']['output'];
  /** 删除应用菜单 */
  deleteAppMenu: Scalars['Boolean']['output'];
  /** 删除应用策略模板 */
  deleteAppPolicy: Scalars['Boolean']['output'];
  /** 删除策略视图 */
  deleteAppPolicyView: Scalars['Boolean']['output'];
  /** 删除应用角色 */
  deleteAppRole: Scalars['Boolean']['output'];
  /** 删除国家 */
  deleteCountry: Scalars['Boolean']['output'];
  /** 删除货币 */
  deleteCurrency: Scalars['Boolean']['output'];
  /** 删除文件凭证 */
  deleteFileIdentity: Scalars['Boolean']['output'];
  /** 删除文件来源 */
  deleteFileSource: Scalars['Boolean']['output'];
  /** 删除OauthClient */
  deleteOauthClient: Scalars['Boolean']['output'];
  /** 删除组织目录 */
  deleteOrganization: Scalars['Boolean']['output'];
  /** 删除组织策略 */
  deleteOrganizationPolicy: Scalars['Boolean']['output'];
  /** 删除配额 */
  deleteQuota: Scalars['Boolean']['output'];
  /** 删除配额项 */
  deleteQuotaItem: Scalars['Boolean']['output'];
  /** 删除地区 */
  deleteRegion: Scalars['Boolean']['output'];
  /** 删除角色或组 */
  deleteRole: Scalars['Boolean']['output'];
  /** 删除用户 */
  deleteUser: Scalars['Boolean']['output'];
  /** 删除用户设备 */
  deleteUserDevice: Scalars['Boolean']['output'];
  /** 删除用户凭证 */
  deleteUserIdentity: Scalars['Boolean']['output'];
  /** 删除密码策略 */
  deleteUserPasswordPolicy: Scalars['Boolean']['output'];
  /** 禁用MFA */
  disableMFA: Scalars['Boolean']['output'];
  /** 禁用OauthClient */
  disableOauthClient: OauthClient;
  /** 启用目录管理,返回根节点组织信息 */
  enableDirectory?: Maybe<Org>;
  /** 启用MFA */
  enableMFA: Mfa;
  /** 启用OauthClient */
  enableOauthClient: OauthClient;
  /** 开启用户设备验证 */
  enableVerifyUserDevice: Scalars['Boolean']['output'];
  /** 授权 */
  grant?: Maybe<Permission>;
  /** 移动节点 */
  moveAppDictItem: Scalars['Boolean']['output'];
  /** 应用菜单位置调整，targetLocation: child, up, down */
  moveAppMenu: Scalars['Boolean']['output'];
  /** 移动策略视图 */
  moveAppPolicyView: Scalars['Boolean']['output'];
  /** 国家位置调整，action: child, up, down */
  moveCountry: Scalars['Boolean']['output'];
  /** 组织位置调整，action: child, up, down */
  moveOrganization: Scalars['Boolean']['output'];
  /** 地区位置调整，action: child, up, down */
  moveRegion: Scalars['Boolean']['output'];
  /** 恢复用户 */
  recoverOrgUser: User;
  /** 从组织目录中移除用户 */
  removeOrganizationUser: Scalars['Boolean']['output'];
  /** 重置用户密码并发送邮件 */
  resetUserPasswordByEmail: Scalars['Boolean']['output'];
  /** 取消授权 */
  revoke: Scalars['Boolean']['output'];
  /** 角色移除策略 */
  revokeAppRolePolicy: Scalars['Boolean']['output'];
  /** 取消分配应用 */
  revokeOrganizationApp: Scalars['Boolean']['output'];
  /** 取消分配到组织应用策略 */
  revokeOrganizationAppPolicy: Scalars['Boolean']['output'];
  /** 取消分配到组织应用角色 */
  revokeOrganizationAppRole: Scalars['Boolean']['output'];
  /** 取消分配组用户 */
  revokeRoleUser: Scalars['Boolean']['output'];
  /** 保存组织用户偏好 */
  saveOrgUserPreference: OrgUserPreference;
  /** 发送MFA至用户邮箱 */
  sendMFAToUserByEmail: Scalars['Boolean']['output'];
  /** 设置默认凭证 */
  setDefaultFileIdentity: Scalars['Boolean']['output'];
  /** 同步角色权限策略到组织 */
  syncAppRoleToOrg: Scalars['Boolean']['output'];
  /** 更新应用 */
  updateApp?: Maybe<App>;
  /** 更新应用操作 */
  updateAppAction?: Maybe<AppAction>;
  /** 更新应用数据字典 */
  updateAppDict?: Maybe<AppDict>;
  /** 更新应用数据字典项 */
  updateAppDictItem?: Maybe<AppDictItem>;
  /** 更新应用菜单 */
  updateAppMenu?: Maybe<AppMenu>;
  /** 更新应用策略模板 */
  updateAppPolicy?: Maybe<AppPolicy>;
  /** 更新策略视图 */
  updateAppPolicyView: AppPolicyView;
  /** 修改资源名称 */
  updateAppRes?: Maybe<AppRes>;
  /** 更新应用角色 */
  updateAppRole?: Maybe<AppRole>;
  /** 更新国家 */
  updateCountry?: Maybe<Country>;
  /** 更新货币 */
  updateCurrency?: Maybe<Currency>;
  /** 更新文件凭证 */
  updateFileIdentity: FileIdentity;
  /** 更新文件来源 */
  updateFileSource: FileSource;
  /** 用户登陆配置 */
  updateLoginProfile?: Maybe<UserLoginProfile>;
  /** 更新组织目录 */
  updateOrganization?: Maybe<Org>;
  /** 更新组织策略 */
  updateOrganizationPolicy?: Maybe<OrgPolicy>;
  /** 更新授权 */
  updatePermission?: Maybe<Permission>;
  /** 更新配额 */
  updateQuota: Quota;
  /** 更新配额项 */
  updateQuotaItem: QuotaItem;
  /** 更新地区 */
  updateRegion?: Maybe<Region>;
  /** 更新角色或组 */
  updateRole?: Maybe<OrgRole>;
  /** 更新用户 */
  updateUser?: Maybe<User>;
  /** 更新用户设备 */
  updateUserDevice: UserDevice;
  /** 更新密码策略 */
  updateUserPasswordPolicy?: Maybe<UserPasswordPolicy>;
};


export type MutationAllotOrganizationUserArgs = {
  input: CreateOrgUserInput;
};


export type MutationAssignAppRolePolicyArgs = {
  appID: Scalars['ID']['input'];
  policyIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  roleID: Scalars['ID']['input'];
};


export type MutationAssignAppRolePolicyViewArgs = {
  addAppPolicyIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  appID: Scalars['ID']['input'];
  rmAppPolicyIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  roleID: Scalars['ID']['input'];
};


export type MutationAssignOrgRolePolicyViewArgs = {
  addOrgPolicyIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  orgID: Scalars['ID']['input'];
  rmOrgPolicyIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  roleID: Scalars['ID']['input'];
};


export type MutationAssignOrgUserPolicyViewArgs = {
  addOrgPolicyIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  orgID: Scalars['ID']['input'];
  rmOrgPolicyIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  userID: Scalars['ID']['input'];
};


export type MutationAssignOrganizationAppArgs = {
  appID: Scalars['ID']['input'];
  orgID: Scalars['ID']['input'];
};


export type MutationAssignOrganizationAppPolicyArgs = {
  appPolicyID: Scalars['ID']['input'];
  orgID: Scalars['ID']['input'];
};


export type MutationAssignOrganizationAppRoleArgs = {
  appRoleID: Scalars['ID']['input'];
  orgID: Scalars['ID']['input'];
};


export type MutationAssignRoleUserArgs = {
  input: AssignRoleUserInput;
};


export type MutationAutoGrantAppArgs = {
  appCode: Scalars['String']['input'];
  orgID: Scalars['ID']['input'];
  userID: Scalars['ID']['input'];
};


export type MutationBindUserIdentityArgs = {
  input: CreateUserIdentityInput;
};


export type MutationChangeOrgUserTypeArgs = {
  orgID: Scalars['ID']['input'];
  userID: Scalars['ID']['input'];
  userType: OrgUserUserType;
};


export type MutationChangePasswordArgs = {
  newPwd: Scalars['String']['input'];
  oldPwd: Scalars['String']['input'];
};


export type MutationCreateAppArgs = {
  input: CreateAppInput;
};


export type MutationCreateAppActionsArgs = {
  appID: Scalars['ID']['input'];
  input?: InputMaybe<Array<CreateAppActionInput>>;
};


export type MutationCreateAppDictArgs = {
  appID: Scalars['ID']['input'];
  input: CreateAppDictInput;
};


export type MutationCreateAppDictItemArgs = {
  dictID: Scalars['ID']['input'];
  input: CreateAppDictItemInput;
};


export type MutationCreateAppMenusArgs = {
  appID: Scalars['ID']['input'];
  input?: InputMaybe<Array<CreateAppMenuInput>>;
};


export type MutationCreateAppPolicyArgs = {
  appID: Scalars['ID']['input'];
  appPolicyViewID?: InputMaybe<Scalars['ID']['input']>;
  input: CreateAppPolicyInput;
};


export type MutationCreateAppPolicyViewArgs = {
  input: CreateAppPolicyViewInput;
};


export type MutationCreateAppRoleArgs = {
  appID: Scalars['ID']['input'];
  input: CreateAppRoleInput;
};


export type MutationCreateCountryArgs = {
  input: CreateCountryInput;
};


export type MutationCreateCurrencyArgs = {
  input: CreateCurrencyInput;
};


export type MutationCreateFileIdentityArgs = {
  input: CreateFileIdentityInput;
};


export type MutationCreateFileSourceArgs = {
  input: CreateFileSourceInput;
};


export type MutationCreateOauthClientArgs = {
  input: CreateOauthClientInput;
};


export type MutationCreateOrganizationArgs = {
  input: CreateOrgInput;
};


export type MutationCreateOrganizationAccountArgs = {
  input: CreateUserInput;
  rootOrgID: Scalars['ID']['input'];
};


export type MutationCreateOrganizationPolicyArgs = {
  input: CreateOrgPolicyInput;
};


export type MutationCreateOrganizationUserArgs = {
  input: CreateUserInput;
  orgUserType?: InputMaybe<OrgUserUserType>;
  rootOrgID: Scalars['ID']['input'];
};


export type MutationCreateQuotaArgs = {
  input: CreateQuotaInput;
};


export type MutationCreateQuotaItemArgs = {
  input: CreateQuotaItemInput;
};


export type MutationCreateRegionArgs = {
  input: CreateRegionInput;
};


export type MutationCreateRoleArgs = {
  input: CreateOrgRoleInput;
};


export type MutationCreateRootArgs = {
  input: CreateOrgInput;
};


export type MutationCreateUserPasswordPolicyArgs = {
  input: CreateUserPasswordPolicyInput;
  orgID: Scalars['ID']['input'];
};


export type MutationDeleteAppArgs = {
  appID: Scalars['ID']['input'];
};


export type MutationDeleteAppActionArgs = {
  actionID: Scalars['ID']['input'];
};


export type MutationDeleteAppDictArgs = {
  dictID: Scalars['ID']['input'];
};


export type MutationDeleteAppDictItemArgs = {
  itemID: Scalars['ID']['input'];
};


export type MutationDeleteAppMenuArgs = {
  menuID: Scalars['ID']['input'];
};


export type MutationDeleteAppPolicyArgs = {
  policyID: Scalars['ID']['input'];
};


export type MutationDeleteAppPolicyViewArgs = {
  appPolicyViewID: Scalars['ID']['input'];
};


export type MutationDeleteAppRoleArgs = {
  roleID: Scalars['ID']['input'];
};


export type MutationDeleteCountryArgs = {
  countryID: Scalars['ID']['input'];
};


export type MutationDeleteCurrencyArgs = {
  currencyID: Scalars['ID']['input'];
};


export type MutationDeleteFileIdentityArgs = {
  id: Scalars['ID']['input'];
};


export type MutationDeleteFileSourceArgs = {
  fsID: Scalars['ID']['input'];
};


export type MutationDeleteOauthClientArgs = {
  id: Scalars['ID']['input'];
};


export type MutationDeleteOrganizationArgs = {
  orgID: Scalars['ID']['input'];
};


export type MutationDeleteOrganizationPolicyArgs = {
  orgPolicyID: Scalars['ID']['input'];
};


export type MutationDeleteQuotaArgs = {
  id: Scalars['ID']['input'];
};


export type MutationDeleteQuotaItemArgs = {
  id: Scalars['ID']['input'];
};


export type MutationDeleteRegionArgs = {
  regionID: Scalars['ID']['input'];
};


export type MutationDeleteRoleArgs = {
  roleID: Scalars['ID']['input'];
};


export type MutationDeleteUserArgs = {
  userID: Scalars['ID']['input'];
};


export type MutationDeleteUserDeviceArgs = {
  deviceID: Scalars['ID']['input'];
  userID: Scalars['ID']['input'];
};


export type MutationDeleteUserIdentityArgs = {
  id: Scalars['ID']['input'];
};


export type MutationDisableMfaArgs = {
  userID: Scalars['ID']['input'];
};


export type MutationDisableOauthClientArgs = {
  id: Scalars['ID']['input'];
};


export type MutationEnableDirectoryArgs = {
  input: EnableDirectoryInput;
};


export type MutationEnableMfaArgs = {
  userID: Scalars['ID']['input'];
};


export type MutationEnableOauthClientArgs = {
  id: Scalars['ID']['input'];
};


export type MutationEnableVerifyUserDeviceArgs = {
  deviceInfoInput: CreateUserDeviceInput;
  enable: Scalars['Boolean']['input'];
  userID: Scalars['ID']['input'];
};


export type MutationGrantArgs = {
  input: CreatePermissionInput;
};


export type MutationMoveAppDictItemArgs = {
  action: TreeAction;
  sourceID: Scalars['ID']['input'];
  targetID: Scalars['ID']['input'];
};


export type MutationMoveAppMenuArgs = {
  action: TreeAction;
  sourceID: Scalars['ID']['input'];
  targetID: Scalars['ID']['input'];
};


export type MutationMoveAppPolicyViewArgs = {
  action: TreeAction;
  sourceID: Scalars['ID']['input'];
  targetID: Scalars['ID']['input'];
};


export type MutationMoveCountryArgs = {
  action: ListAction;
  sourceID: Scalars['ID']['input'];
  targetId: Scalars['ID']['input'];
};


export type MutationMoveOrganizationArgs = {
  action: TreeAction;
  sourceID: Scalars['ID']['input'];
  targetId: Scalars['ID']['input'];
};


export type MutationMoveRegionArgs = {
  action: TreeAction;
  sourceID: Scalars['ID']['input'];
  targetId: Scalars['ID']['input'];
};


export type MutationRecoverOrgUserArgs = {
  contact?: InputMaybe<UpdateUserAddrInput>;
  pwdInput?: InputMaybe<CreateUserPasswordInput>;
  pwdKind: UserLoginProfileSetKind;
  userID: Scalars['ID']['input'];
  userInput: UpdateUserInput;
};


export type MutationRemoveOrganizationUserArgs = {
  orgID: Scalars['ID']['input'];
  userID: Scalars['ID']['input'];
};


export type MutationResetUserPasswordByEmailArgs = {
  userId: Scalars['ID']['input'];
};


export type MutationRevokeArgs = {
  orgID: Scalars['ID']['input'];
  permissionID: Scalars['ID']['input'];
};


export type MutationRevokeAppRolePolicyArgs = {
  appID: Scalars['ID']['input'];
  policyIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  roleID: Scalars['ID']['input'];
};


export type MutationRevokeOrganizationAppArgs = {
  appID: Scalars['ID']['input'];
  orgID: Scalars['ID']['input'];
};


export type MutationRevokeOrganizationAppPolicyArgs = {
  appPolicyID: Scalars['ID']['input'];
  orgID: Scalars['ID']['input'];
};


export type MutationRevokeOrganizationAppRoleArgs = {
  appRoleID: Scalars['ID']['input'];
  orgID: Scalars['ID']['input'];
};


export type MutationRevokeRoleUserArgs = {
  roleID: Scalars['ID']['input'];
  userID: Scalars['ID']['input'];
};


export type MutationSaveOrgUserPreferenceArgs = {
  input: OrgUserPreferenceInput;
};


export type MutationSendMfaToUserByEmailArgs = {
  userID: Scalars['ID']['input'];
};


export type MutationSetDefaultFileIdentityArgs = {
  identityID: Scalars['ID']['input'];
  orgID: Scalars['ID']['input'];
};


export type MutationSyncAppRoleToOrgArgs = {
  appRoleID: Scalars['ID']['input'];
  orgID: Scalars['ID']['input'];
};


export type MutationUpdateAppArgs = {
  appID: Scalars['ID']['input'];
  input: UpdateAppInput;
};


export type MutationUpdateAppActionArgs = {
  actionID: Scalars['ID']['input'];
  input: UpdateAppActionInput;
};


export type MutationUpdateAppDictArgs = {
  dictID: Scalars['ID']['input'];
  input: UpdateAppDictInput;
};


export type MutationUpdateAppDictItemArgs = {
  input: UpdateAppDictItemInput;
  itemID: Scalars['ID']['input'];
};


export type MutationUpdateAppMenuArgs = {
  input: UpdateAppMenuInput;
  menuID: Scalars['ID']['input'];
};


export type MutationUpdateAppPolicyArgs = {
  input: UpdateAppPolicyInput;
  policyID: Scalars['ID']['input'];
};


export type MutationUpdateAppPolicyViewArgs = {
  appPolicyViewID: Scalars['ID']['input'];
  input: UpdateAppPolicyViewInput;
};


export type MutationUpdateAppResArgs = {
  appResID: Scalars['ID']['input'];
  input: UpdateAppResInput;
};


export type MutationUpdateAppRoleArgs = {
  input: UpdateAppRoleInput;
  roleID: Scalars['ID']['input'];
};


export type MutationUpdateCountryArgs = {
  countryID: Scalars['ID']['input'];
  input: UpdateCountryInput;
};


export type MutationUpdateCurrencyArgs = {
  currencyID: Scalars['ID']['input'];
  input: UpdateCurrencyInput;
};


export type MutationUpdateFileIdentityArgs = {
  id: Scalars['ID']['input'];
  input: UpdateFileIdentityInput;
};


export type MutationUpdateFileSourceArgs = {
  fsID: Scalars['ID']['input'];
  input: UpdateFileSourceInput;
};


export type MutationUpdateLoginProfileArgs = {
  input: UpdateUserLoginProfileInput;
  userID: Scalars['ID']['input'];
};


export type MutationUpdateOrganizationArgs = {
  input: UpdateOrgInput;
  orgID: Scalars['ID']['input'];
};


export type MutationUpdateOrganizationPolicyArgs = {
  input: UpdateOrgPolicyInput;
  orgPolicyID: Scalars['ID']['input'];
};


export type MutationUpdatePermissionArgs = {
  input: UpdatePermissionInput;
  permissionID: Scalars['ID']['input'];
};


export type MutationUpdateQuotaArgs = {
  id: Scalars['ID']['input'];
  input: UpdateQuotaInput;
};


export type MutationUpdateQuotaItemArgs = {
  id: Scalars['ID']['input'];
  input: UpdateQuotaItemInput;
};


export type MutationUpdateRegionArgs = {
  input: UpdateRegionInput;
  regionID: Scalars['ID']['input'];
};


export type MutationUpdateRoleArgs = {
  input: UpdateOrgRoleInput;
  roleID: Scalars['ID']['input'];
};


export type MutationUpdateUserArgs = {
  contact?: InputMaybe<UpdateUserAddrInput>;
  input: UpdateUserInput;
  userID: Scalars['ID']['input'];
};


export type MutationUpdateUserDeviceArgs = {
  deviceID: Scalars['ID']['input'];
  input: UpdateUserDeviceInput;
};


export type MutationUpdateUserPasswordPolicyArgs = {
  input: UpdateUserPasswordPolicyInput;
  orgID: Scalars['ID']['input'];
};

/**
 * An object with an ID.
 * Follows the [Relay Global Object Identification Specification](https://relay.dev/graphql/objectidentification.htm)
 */
export type Node = {
  /** The id of the object. */
  id: Scalars['ID']['output'];
};

export type OauthClient = Node & {
  __typename?: 'OauthClient';
  /** id */
  clientID: Scalars['String']['output'];
  /** 密钥 */
  clientSecret: Scalars['String']['output'];
  createdAt: Scalars['Time']['output'];
  createdBy: Scalars['Int']['output'];
  /** 授权类型 */
  grantTypes: OauthClientGrantTypes;
  id: Scalars['ID']['output'];
  /** 最后认证时间 */
  lastAuthAt?: Maybe<Scalars['Time']['output']>;
  /** 名称 */
  name: Scalars['String']['output'];
  /** 状态 */
  status: OauthClientSimpleStatus;
  updatedAt?: Maybe<Scalars['Time']['output']>;
  updatedBy?: Maybe<Scalars['Int']['output']>;
  user: User;
  /** 关联用户id */
  userID: Scalars['ID']['output'];
};

/** OauthClientGrantTypes is enum for the field grant_types */
export enum OauthClientGrantTypes {
  ClientCredentials = 'client_credentials'
}

/** Ordering options for OauthClient connections */
export type OauthClientOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order OauthClients. */
  field: OauthClientOrderField;
};

/** Properties by which OauthClient connections can be ordered. */
export enum OauthClientOrderField {
  CreatedAt = 'createdAt'
}

/** OauthClientSimpleStatus is enum for the field status */
export enum OauthClientSimpleStatus {
  Active = 'active',
  Disabled = 'disabled',
  Inactive = 'inactive',
  Processing = 'processing'
}

/**
 * OauthClientWhereInput is used for filtering OauthClient objects.
 * Input was generated by ent.
 */
export type OauthClientWhereInput = {
  and?: InputMaybe<Array<OauthClientWhereInput>>;
  /** created_at field predicates */
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  createdAtGT?: InputMaybe<Scalars['Time']['input']>;
  createdAtGTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  createdAtLT?: InputMaybe<Scalars['Time']['input']>;
  createdAtLTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  createdAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  /** created_by field predicates */
  createdBy?: InputMaybe<Scalars['Int']['input']>;
  createdByGT?: InputMaybe<Scalars['Int']['input']>;
  createdByGTE?: InputMaybe<Scalars['Int']['input']>;
  createdByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  createdByLT?: InputMaybe<Scalars['Int']['input']>;
  createdByLTE?: InputMaybe<Scalars['Int']['input']>;
  createdByNEQ?: InputMaybe<Scalars['Int']['input']>;
  createdByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  /** grant_types field predicates */
  grantTypes?: InputMaybe<OauthClientGrantTypes>;
  grantTypesIn?: InputMaybe<Array<OauthClientGrantTypes>>;
  grantTypesNEQ?: InputMaybe<OauthClientGrantTypes>;
  grantTypesNotIn?: InputMaybe<Array<OauthClientGrantTypes>>;
  /** user edge predicates */
  hasUser?: InputMaybe<Scalars['Boolean']['input']>;
  hasUserWith?: InputMaybe<Array<UserWhereInput>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** name field predicates */
  name?: InputMaybe<Scalars['String']['input']>;
  nameContains?: InputMaybe<Scalars['String']['input']>;
  nameContainsFold?: InputMaybe<Scalars['String']['input']>;
  nameEqualFold?: InputMaybe<Scalars['String']['input']>;
  nameGT?: InputMaybe<Scalars['String']['input']>;
  nameGTE?: InputMaybe<Scalars['String']['input']>;
  nameHasPrefix?: InputMaybe<Scalars['String']['input']>;
  nameHasSuffix?: InputMaybe<Scalars['String']['input']>;
  nameIn?: InputMaybe<Array<Scalars['String']['input']>>;
  nameLT?: InputMaybe<Scalars['String']['input']>;
  nameLTE?: InputMaybe<Scalars['String']['input']>;
  nameNEQ?: InputMaybe<Scalars['String']['input']>;
  nameNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  not?: InputMaybe<OauthClientWhereInput>;
  or?: InputMaybe<Array<OauthClientWhereInput>>;
  /** status field predicates */
  status?: InputMaybe<OauthClientSimpleStatus>;
  statusIn?: InputMaybe<Array<OauthClientSimpleStatus>>;
  statusNEQ?: InputMaybe<OauthClientSimpleStatus>;
  statusNotIn?: InputMaybe<Array<OauthClientSimpleStatus>>;
  /** updated_at field predicates */
  updatedAt?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedAtLT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtLTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_by field predicates */
  updatedBy?: InputMaybe<Scalars['Int']['input']>;
  updatedByGT?: InputMaybe<Scalars['Int']['input']>;
  updatedByGTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedByLT?: InputMaybe<Scalars['Int']['input']>;
  updatedByLTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByNEQ?: InputMaybe<Scalars['Int']['input']>;
  updatedByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** user_id field predicates */
  userID?: InputMaybe<Scalars['ID']['input']>;
  userIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  userIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  userIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
};

/** Possible directions in which to order a list of items when provided an `orderBy` argument. */
export enum OrderDirection {
  /** Specifies an ascending order for a given `orderBy` argument. */
  Asc = 'ASC',
  /** Specifies a descending order for a given `orderBy` argument. */
  Desc = 'DESC'
}

export type Org = Node & {
  __typename?: 'Org';
  /** 获取顶级组织 */
  TopOrg?: Maybe<Org>;
  /** 组织的domain，如果没有设置，则会取最近的父组织的domain */
  actualDomain: Scalars['String']['output'];
  apps: AppConnection;
  children?: Maybe<Array<Org>>;
  /** 系统代码 */
  code?: Maybe<Scalars['String']['output']>;
  /** 国家或地区2字码 */
  countryCode?: Maybe<Scalars['String']['output']>;
  createdAt: Scalars['Time']['output'];
  createdBy: Scalars['Int']['output'];
  /** 自定义域名 */
  customDomain?: Maybe<Array<Scalars['String']['output']>>;
  deletedAt?: Maybe<Scalars['Time']['output']>;
  displaySort?: Maybe<Scalars['Int']['output']>;
  /** 默认域名 */
  domain?: Maybe<Scalars['String']['output']>;
  /** 组织下文件凭证 */
  fileIdentities?: Maybe<Array<FileIdentity>>;
  id: Scalars['ID']['output'];
  /** 是否允许解除应用策略 */
  isAllowRevokeAppPolicy: Scalars['Boolean']['output'];
  /** 分类: 根节点,组织节点 */
  kind: OrgKind;
  /** 组织本位币 */
  localCurrency?: Maybe<Scalars['String']['output']>;
  /** 组织图标 */
  logo?: Maybe<OrgLogo>;
  /** 组织名称 */
  name: Scalars['String']['output'];
  /** 组织下登录策略 */
  orgQuota?: Maybe<Array<Quota>>;
  /** 管理账户 */
  owner?: Maybe<User>;
  /** 管理账户ID,如果设置则该组织将升级为根组织 */
  ownerID?: Maybe<Scalars['ID']['output']>;
  parent: Org;
  /** 父级ID,0为根组织. */
  parentID: Scalars['ID']['output'];
  /** 路径编码 */
  path?: Maybe<Scalars['String']['output']>;
  permissions: PermissionConnection;
  policies: OrgPolicyConnection;
  /** 简介 */
  profile?: Maybe<Scalars['String']['output']>;
  /** 状态 */
  status?: Maybe<OrgSimpleStatus>;
  /** 时区 */
  timezone?: Maybe<Scalars['String']['output']>;
  updatedAt?: Maybe<Scalars['Time']['output']>;
  updatedBy?: Maybe<Scalars['Int']['output']>;
  /** 组织下密码策略 */
  userPasswordPolicy?: Maybe<UserPasswordPolicy>;
  users: UserConnection;
};


export type OrgAppsArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<AppOrder>;
  where?: InputMaybe<AppWhereInput>;
};


export type OrgIsAllowRevokeAppPolicyArgs = {
  appPolicyID: Scalars['ID']['input'];
};


export type OrgPermissionsArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<PermissionOrder>;
  where?: InputMaybe<PermissionWhereInput>;
};


export type OrgPoliciesArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<OrgPolicyOrder>;
  where?: InputMaybe<OrgPolicyWhereInput>;
};


export type OrgUsersArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<UserOrder>;
  where?: InputMaybe<UserWhereInput>;
};

/** A connection to a list of items. */
export type OrgConnection = {
  __typename?: 'OrgConnection';
  /** A list of edges. */
  edges?: Maybe<Array<Maybe<OrgEdge>>>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** Identifies the total count of items in the connection. */
  totalCount: Scalars['Int']['output'];
};

/** An edge in a connection. */
export type OrgEdge = {
  __typename?: 'OrgEdge';
  /** A cursor for use in pagination. */
  cursor: Scalars['Cursor']['output'];
  /** The item at the end of the edge. */
  node?: Maybe<Org>;
};

/** 业务调用的fileIdentity */
export type OrgFileIdentity = {
  __typename?: 'OrgFileIdentity';
  /** 备注 */
  comments?: Maybe<Scalars['String']['output']>;
  createdAt: Scalars['Time']['output'];
  createdBy: Scalars['Int']['output'];
  /** 文件来源ID */
  fileSourceID: Scalars['ID']['output'];
  id: Scalars['ID']['output'];
  /** 租户默认的凭证 */
  isDefault: Scalars['Boolean']['output'];
  source: FileSource;
  /** 组织ID */
  tenantID: Scalars['ID']['output'];
  updatedAt?: Maybe<Scalars['Time']['output']>;
  updatedBy?: Maybe<Scalars['Int']['output']>;
};

/** OrgKind is enum for the field kind */
export enum OrgKind {
  Org = 'org',
  Root = 'root'
}

export type OrgLogo = {
  __typename?: 'OrgLogo';
  favicon?: Maybe<Scalars['String']['output']>;
  logo?: Maybe<Scalars['String']['output']>;
  thumbLogo?: Maybe<Scalars['String']['output']>;
};

export type OrgLogoInput = {
  favicon?: InputMaybe<Scalars['String']['input']>;
  logo?: InputMaybe<Scalars['String']['input']>;
  thumbLogo?: InputMaybe<Scalars['String']['input']>;
};

/** Ordering options for Org connections */
export type OrgOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order Orgs. */
  field: OrgOrderField;
};

/** Properties by which Org connections can be ordered. */
export enum OrgOrderField {
  CreatedAt = 'createdAt',
  DisplaySort = 'displaySort'
}

export type OrgPolicy = Node & {
  __typename?: 'OrgPolicy';
  app?: Maybe<App>;
  appPolicy?: Maybe<AppPolicy>;
  /** 所属应用策略,如果是自定义应用策略,则为空 */
  appPolicyID?: Maybe<Scalars['ID']['output']>;
  /** 描述 */
  comments?: Maybe<Scalars['String']['output']>;
  createdAt: Scalars['Time']['output'];
  createdBy: Scalars['Int']['output'];
  id: Scalars['ID']['output'];
  /** 是否授权role */
  isGrantRole: Scalars['Boolean']['output'];
  /** 是否授权user */
  isGrantUser: Scalars['Boolean']['output'];
  /** 策略名称 */
  name: Scalars['String']['output'];
  org?: Maybe<Org>;
  /** 租户ID */
  orgID?: Maybe<Scalars['ID']['output']>;
  permissions?: Maybe<Array<Permission>>;
  /** 策略规则 */
  rules: Array<Maybe<PolicyRule>>;
  updatedAt?: Maybe<Scalars['Time']['output']>;
  updatedBy?: Maybe<Scalars['Int']['output']>;
};


export type OrgPolicyIsGrantRoleArgs = {
  roleID: Scalars['ID']['input'];
};


export type OrgPolicyIsGrantUserArgs = {
  userID: Scalars['ID']['input'];
};

/** A connection to a list of items. */
export type OrgPolicyConnection = {
  __typename?: 'OrgPolicyConnection';
  /** A list of edges. */
  edges?: Maybe<Array<Maybe<OrgPolicyEdge>>>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** Identifies the total count of items in the connection. */
  totalCount: Scalars['Int']['output'];
};

/** An edge in a connection. */
export type OrgPolicyEdge = {
  __typename?: 'OrgPolicyEdge';
  /** A cursor for use in pagination. */
  cursor: Scalars['Cursor']['output'];
  /** The item at the end of the edge. */
  node?: Maybe<OrgPolicy>;
};

/** Ordering options for OrgPolicy connections */
export type OrgPolicyOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order OrgPolicies. */
  field: OrgPolicyOrderField;
};

/** Properties by which OrgPolicy connections can be ordered. */
export enum OrgPolicyOrderField {
  CreatedAt = 'createdAt'
}

/**
 * OrgPolicyWhereInput is used for filtering OrgPolicy objects.
 * Input was generated by ent.
 */
export type OrgPolicyWhereInput = {
  and?: InputMaybe<Array<OrgPolicyWhereInput>>;
  /** app_policy_id field predicates */
  appPolicyID?: InputMaybe<Scalars['ID']['input']>;
  appPolicyIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  appPolicyIDIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  appPolicyIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  appPolicyIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  appPolicyIDNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** comments field predicates */
  comments?: InputMaybe<Scalars['String']['input']>;
  commentsContains?: InputMaybe<Scalars['String']['input']>;
  commentsContainsFold?: InputMaybe<Scalars['String']['input']>;
  commentsEqualFold?: InputMaybe<Scalars['String']['input']>;
  commentsGT?: InputMaybe<Scalars['String']['input']>;
  commentsGTE?: InputMaybe<Scalars['String']['input']>;
  commentsHasPrefix?: InputMaybe<Scalars['String']['input']>;
  commentsHasSuffix?: InputMaybe<Scalars['String']['input']>;
  commentsIn?: InputMaybe<Array<Scalars['String']['input']>>;
  commentsIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  commentsLT?: InputMaybe<Scalars['String']['input']>;
  commentsLTE?: InputMaybe<Scalars['String']['input']>;
  commentsNEQ?: InputMaybe<Scalars['String']['input']>;
  commentsNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  commentsNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** created_at field predicates */
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  createdAtGT?: InputMaybe<Scalars['Time']['input']>;
  createdAtGTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  createdAtLT?: InputMaybe<Scalars['Time']['input']>;
  createdAtLTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  createdAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  /** created_by field predicates */
  createdBy?: InputMaybe<Scalars['Int']['input']>;
  createdByGT?: InputMaybe<Scalars['Int']['input']>;
  createdByGTE?: InputMaybe<Scalars['Int']['input']>;
  createdByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  createdByLT?: InputMaybe<Scalars['Int']['input']>;
  createdByLTE?: InputMaybe<Scalars['Int']['input']>;
  createdByNEQ?: InputMaybe<Scalars['Int']['input']>;
  createdByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  /** app edge predicates */
  hasApp?: InputMaybe<Scalars['Boolean']['input']>;
  /** app_policy edge predicates */
  hasAppPolicy?: InputMaybe<Scalars['Boolean']['input']>;
  hasAppPolicyWith?: InputMaybe<Array<AppPolicyWhereInput>>;
  hasAppWith?: InputMaybe<Array<AppWhereInput>>;
  /** org edge predicates */
  hasOrg?: InputMaybe<Scalars['Boolean']['input']>;
  hasOrgWith?: InputMaybe<Array<OrgWhereInput>>;
  /** permissions edge predicates */
  hasPermissions?: InputMaybe<Scalars['Boolean']['input']>;
  hasPermissionsWith?: InputMaybe<Array<PermissionWhereInput>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** name field predicates */
  name?: InputMaybe<Scalars['String']['input']>;
  nameContains?: InputMaybe<Scalars['String']['input']>;
  nameContainsFold?: InputMaybe<Scalars['String']['input']>;
  nameEqualFold?: InputMaybe<Scalars['String']['input']>;
  nameGT?: InputMaybe<Scalars['String']['input']>;
  nameGTE?: InputMaybe<Scalars['String']['input']>;
  nameHasPrefix?: InputMaybe<Scalars['String']['input']>;
  nameHasSuffix?: InputMaybe<Scalars['String']['input']>;
  nameIn?: InputMaybe<Array<Scalars['String']['input']>>;
  nameLT?: InputMaybe<Scalars['String']['input']>;
  nameLTE?: InputMaybe<Scalars['String']['input']>;
  nameNEQ?: InputMaybe<Scalars['String']['input']>;
  nameNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  not?: InputMaybe<OrgPolicyWhereInput>;
  or?: InputMaybe<Array<OrgPolicyWhereInput>>;
  /** org_id field predicates */
  orgID?: InputMaybe<Scalars['ID']['input']>;
  orgIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  orgIDIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  orgIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  orgIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  orgIDNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_at field predicates */
  updatedAt?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedAtLT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtLTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_by field predicates */
  updatedBy?: InputMaybe<Scalars['Int']['input']>;
  updatedByGT?: InputMaybe<Scalars['Int']['input']>;
  updatedByGTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedByLT?: InputMaybe<Scalars['Int']['input']>;
  updatedByLTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByNEQ?: InputMaybe<Scalars['Int']['input']>;
  updatedByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByNotNil?: InputMaybe<Scalars['Boolean']['input']>;
};

export type OrgRole = Node & {
  __typename?: 'OrgRole';
  /** 备注 */
  comments?: Maybe<Scalars['String']['output']>;
  createdAt: Scalars['Time']['output'];
  createdBy: Scalars['Int']['output'];
  id: Scalars['ID']['output'];
  /** 是否系统角色 */
  isAppRole: Scalars['Boolean']['output'];
  /** 是否分配给user */
  isGrantUser: Scalars['Boolean']['output'];
  /** 类型,group:组,role:角色 */
  kind: OrgRoleKind;
  /** 名称 */
  name: Scalars['String']['output'];
  /** 租户ID */
  orgID?: Maybe<Scalars['ID']['output']>;
  updatedAt?: Maybe<Scalars['Time']['output']>;
  updatedBy?: Maybe<Scalars['Int']['output']>;
};


export type OrgRoleIsGrantUserArgs = {
  userID: Scalars['ID']['input'];
};

/** A connection to a list of items. */
export type OrgRoleConnection = {
  __typename?: 'OrgRoleConnection';
  /** A list of edges. */
  edges?: Maybe<Array<Maybe<OrgRoleEdge>>>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** Identifies the total count of items in the connection. */
  totalCount: Scalars['Int']['output'];
};

/** An edge in a connection. */
export type OrgRoleEdge = {
  __typename?: 'OrgRoleEdge';
  /** A cursor for use in pagination. */
  cursor: Scalars['Cursor']['output'];
  /** The item at the end of the edge. */
  node?: Maybe<OrgRole>;
};

/** OrgRoleKind is enum for the field kind */
export enum OrgRoleKind {
  Group = 'group',
  Role = 'role'
}

/** Ordering options for OrgRole connections */
export type OrgRoleOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order OrgRoles. */
  field: OrgRoleOrderField;
};

/** Properties by which OrgRole connections can be ordered. */
export enum OrgRoleOrderField {
  CreatedAt = 'createdAt'
}

/** Ordering options for OrgRoleUser connections */
export type OrgRoleUserOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order OrgRoleUsers. */
  field: OrgRoleUserOrderField;
};

/** Properties by which OrgRoleUser connections can be ordered. */
export enum OrgRoleUserOrderField {
  CreatedAt = 'createdAt'
}

/**
 * OrgRoleUserWhereInput is used for filtering OrgRoleUser objects.
 * Input was generated by ent.
 */
export type OrgRoleUserWhereInput = {
  and?: InputMaybe<Array<OrgRoleUserWhereInput>>;
  /** created_at field predicates */
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  createdAtGT?: InputMaybe<Scalars['Time']['input']>;
  createdAtGTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  createdAtLT?: InputMaybe<Scalars['Time']['input']>;
  createdAtLTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  createdAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  /** created_by field predicates */
  createdBy?: InputMaybe<Scalars['Int']['input']>;
  createdByGT?: InputMaybe<Scalars['Int']['input']>;
  createdByGTE?: InputMaybe<Scalars['Int']['input']>;
  createdByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  createdByLT?: InputMaybe<Scalars['Int']['input']>;
  createdByLTE?: InputMaybe<Scalars['Int']['input']>;
  createdByNEQ?: InputMaybe<Scalars['Int']['input']>;
  createdByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  not?: InputMaybe<OrgRoleUserWhereInput>;
  or?: InputMaybe<Array<OrgRoleUserWhereInput>>;
  /** updated_at field predicates */
  updatedAt?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedAtLT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtLTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_by field predicates */
  updatedBy?: InputMaybe<Scalars['Int']['input']>;
  updatedByGT?: InputMaybe<Scalars['Int']['input']>;
  updatedByGTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedByLT?: InputMaybe<Scalars['Int']['input']>;
  updatedByLTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByNEQ?: InputMaybe<Scalars['Int']['input']>;
  updatedByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByNotNil?: InputMaybe<Scalars['Boolean']['input']>;
};

/**
 * OrgRoleWhereInput is used for filtering OrgRole objects.
 * Input was generated by ent.
 */
export type OrgRoleWhereInput = {
  and?: InputMaybe<Array<OrgRoleWhereInput>>;
  /** comments field predicates */
  comments?: InputMaybe<Scalars['String']['input']>;
  commentsContains?: InputMaybe<Scalars['String']['input']>;
  commentsContainsFold?: InputMaybe<Scalars['String']['input']>;
  commentsEqualFold?: InputMaybe<Scalars['String']['input']>;
  commentsGT?: InputMaybe<Scalars['String']['input']>;
  commentsGTE?: InputMaybe<Scalars['String']['input']>;
  commentsHasPrefix?: InputMaybe<Scalars['String']['input']>;
  commentsHasSuffix?: InputMaybe<Scalars['String']['input']>;
  commentsIn?: InputMaybe<Array<Scalars['String']['input']>>;
  commentsIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  commentsLT?: InputMaybe<Scalars['String']['input']>;
  commentsLTE?: InputMaybe<Scalars['String']['input']>;
  commentsNEQ?: InputMaybe<Scalars['String']['input']>;
  commentsNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  commentsNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** created_at field predicates */
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  createdAtGT?: InputMaybe<Scalars['Time']['input']>;
  createdAtGTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  createdAtLT?: InputMaybe<Scalars['Time']['input']>;
  createdAtLTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  createdAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  /** created_by field predicates */
  createdBy?: InputMaybe<Scalars['Int']['input']>;
  createdByGT?: InputMaybe<Scalars['Int']['input']>;
  createdByGTE?: InputMaybe<Scalars['Int']['input']>;
  createdByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  createdByLT?: InputMaybe<Scalars['Int']['input']>;
  createdByLTE?: InputMaybe<Scalars['Int']['input']>;
  createdByNEQ?: InputMaybe<Scalars['Int']['input']>;
  createdByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  /** org edge predicates */
  hasOrg?: InputMaybe<Scalars['Boolean']['input']>;
  hasOrgWith?: InputMaybe<Array<OrgWhereInput>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** kind field predicates */
  kind?: InputMaybe<OrgRoleKind>;
  kindIn?: InputMaybe<Array<OrgRoleKind>>;
  kindNEQ?: InputMaybe<OrgRoleKind>;
  kindNotIn?: InputMaybe<Array<OrgRoleKind>>;
  /** name field predicates */
  name?: InputMaybe<Scalars['String']['input']>;
  nameContains?: InputMaybe<Scalars['String']['input']>;
  nameContainsFold?: InputMaybe<Scalars['String']['input']>;
  nameEqualFold?: InputMaybe<Scalars['String']['input']>;
  nameGT?: InputMaybe<Scalars['String']['input']>;
  nameGTE?: InputMaybe<Scalars['String']['input']>;
  nameHasPrefix?: InputMaybe<Scalars['String']['input']>;
  nameHasSuffix?: InputMaybe<Scalars['String']['input']>;
  nameIn?: InputMaybe<Array<Scalars['String']['input']>>;
  nameLT?: InputMaybe<Scalars['String']['input']>;
  nameLTE?: InputMaybe<Scalars['String']['input']>;
  nameNEQ?: InputMaybe<Scalars['String']['input']>;
  nameNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  not?: InputMaybe<OrgRoleWhereInput>;
  or?: InputMaybe<Array<OrgRoleWhereInput>>;
  /** org_id field predicates */
  orgID?: InputMaybe<Scalars['ID']['input']>;
  orgIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  orgIDIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  orgIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  orgIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  orgIDNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_at field predicates */
  updatedAt?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedAtLT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtLTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_by field predicates */
  updatedBy?: InputMaybe<Scalars['Int']['input']>;
  updatedByGT?: InputMaybe<Scalars['Int']['input']>;
  updatedByGTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedByLT?: InputMaybe<Scalars['Int']['input']>;
  updatedByLTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByNEQ?: InputMaybe<Scalars['Int']['input']>;
  updatedByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByNotNil?: InputMaybe<Scalars['Boolean']['input']>;
};

/** OrgSimpleStatus is enum for the field status */
export enum OrgSimpleStatus {
  Active = 'active',
  Disabled = 'disabled',
  Inactive = 'inactive',
  Processing = 'processing'
}

/** Ordering options for OrgUser connections */
export type OrgUserOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order OrgUsers. */
  field: OrgUserOrderField;
};

/** Properties by which OrgUser connections can be ordered. */
export enum OrgUserOrderField {
  CreatedAt = 'createdAt'
}

export type OrgUserPreference = Node & {
  __typename?: 'OrgUserPreference';
  createdAt: Scalars['Time']['output'];
  createdBy: Scalars['Int']['output'];
  id: Scalars['ID']['output'];
  /** 用户收藏菜单 */
  menuFavorite?: Maybe<Array<Scalars['ID']['output']>>;
  /** 用户最近访问菜单 */
  menuRecent?: Maybe<Array<Scalars['ID']['output']>>;
  org: Org;
  /** 租户ID */
  orgID: Scalars['ID']['output'];
  updatedAt?: Maybe<Scalars['Time']['output']>;
  updatedBy?: Maybe<Scalars['Int']['output']>;
  user: User;
  /** 用户id */
  userID: Scalars['ID']['output'];
};

/** A connection to a list of items. */
export type OrgUserPreferenceConnection = {
  __typename?: 'OrgUserPreferenceConnection';
  /** A list of edges. */
  edges?: Maybe<Array<Maybe<OrgUserPreferenceEdge>>>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** Identifies the total count of items in the connection. */
  totalCount: Scalars['Int']['output'];
};

/** An edge in a connection. */
export type OrgUserPreferenceEdge = {
  __typename?: 'OrgUserPreferenceEdge';
  /** A cursor for use in pagination. */
  cursor: Scalars['Cursor']['output'];
  /** The item at the end of the edge. */
  node?: Maybe<OrgUserPreference>;
};

export type OrgUserPreferenceInput = {
  /** 用户收藏菜单 */
  menuFavorite?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** 用户最近访问菜单 */
  menuRecent?: InputMaybe<Array<Scalars['ID']['input']>>;
};

/** Ordering options for OrgUserPreference connections */
export type OrgUserPreferenceOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order OrgUserPreferences. */
  field: OrgUserPreferenceOrderField;
};

/** Properties by which OrgUserPreference connections can be ordered. */
export enum OrgUserPreferenceOrderField {
  CreatedAt = 'createdAt'
}

/**
 * OrgUserPreferenceWhereInput is used for filtering OrgUserPreference objects.
 * Input was generated by ent.
 */
export type OrgUserPreferenceWhereInput = {
  and?: InputMaybe<Array<OrgUserPreferenceWhereInput>>;
  /** created_at field predicates */
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  createdAtGT?: InputMaybe<Scalars['Time']['input']>;
  createdAtGTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  createdAtLT?: InputMaybe<Scalars['Time']['input']>;
  createdAtLTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  createdAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  /** created_by field predicates */
  createdBy?: InputMaybe<Scalars['Int']['input']>;
  createdByGT?: InputMaybe<Scalars['Int']['input']>;
  createdByGTE?: InputMaybe<Scalars['Int']['input']>;
  createdByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  createdByLT?: InputMaybe<Scalars['Int']['input']>;
  createdByLTE?: InputMaybe<Scalars['Int']['input']>;
  createdByNEQ?: InputMaybe<Scalars['Int']['input']>;
  createdByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  /** org edge predicates */
  hasOrg?: InputMaybe<Scalars['Boolean']['input']>;
  hasOrgWith?: InputMaybe<Array<OrgWhereInput>>;
  /** user edge predicates */
  hasUser?: InputMaybe<Scalars['Boolean']['input']>;
  hasUserWith?: InputMaybe<Array<UserWhereInput>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  not?: InputMaybe<OrgUserPreferenceWhereInput>;
  or?: InputMaybe<Array<OrgUserPreferenceWhereInput>>;
  /** org_id field predicates */
  orgID?: InputMaybe<Scalars['ID']['input']>;
  orgIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  orgIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  orgIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** updated_at field predicates */
  updatedAt?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedAtLT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtLTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_by field predicates */
  updatedBy?: InputMaybe<Scalars['Int']['input']>;
  updatedByGT?: InputMaybe<Scalars['Int']['input']>;
  updatedByGTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedByLT?: InputMaybe<Scalars['Int']['input']>;
  updatedByLTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByNEQ?: InputMaybe<Scalars['Int']['input']>;
  updatedByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** user_id field predicates */
  userID?: InputMaybe<Scalars['ID']['input']>;
  userIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  userIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  userIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
};

/** OrgUserUserType is enum for the field user_type */
export enum OrgUserUserType {
  External = 'external',
  Internal = 'internal'
}

/**
 * OrgUserWhereInput is used for filtering OrgUser objects.
 * Input was generated by ent.
 */
export type OrgUserWhereInput = {
  and?: InputMaybe<Array<OrgUserWhereInput>>;
  /** created_at field predicates */
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  createdAtGT?: InputMaybe<Scalars['Time']['input']>;
  createdAtGTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  createdAtLT?: InputMaybe<Scalars['Time']['input']>;
  createdAtLTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  createdAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  /** created_by field predicates */
  createdBy?: InputMaybe<Scalars['Int']['input']>;
  createdByGT?: InputMaybe<Scalars['Int']['input']>;
  createdByGTE?: InputMaybe<Scalars['Int']['input']>;
  createdByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  createdByLT?: InputMaybe<Scalars['Int']['input']>;
  createdByLTE?: InputMaybe<Scalars['Int']['input']>;
  createdByNEQ?: InputMaybe<Scalars['Int']['input']>;
  createdByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  /** display_name field predicates */
  displayName?: InputMaybe<Scalars['String']['input']>;
  displayNameContains?: InputMaybe<Scalars['String']['input']>;
  displayNameContainsFold?: InputMaybe<Scalars['String']['input']>;
  displayNameEqualFold?: InputMaybe<Scalars['String']['input']>;
  displayNameGT?: InputMaybe<Scalars['String']['input']>;
  displayNameGTE?: InputMaybe<Scalars['String']['input']>;
  displayNameHasPrefix?: InputMaybe<Scalars['String']['input']>;
  displayNameHasSuffix?: InputMaybe<Scalars['String']['input']>;
  displayNameIn?: InputMaybe<Array<Scalars['String']['input']>>;
  displayNameLT?: InputMaybe<Scalars['String']['input']>;
  displayNameLTE?: InputMaybe<Scalars['String']['input']>;
  displayNameNEQ?: InputMaybe<Scalars['String']['input']>;
  displayNameNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** joined_at field predicates */
  joinedAt?: InputMaybe<Scalars['Time']['input']>;
  joinedAtGT?: InputMaybe<Scalars['Time']['input']>;
  joinedAtGTE?: InputMaybe<Scalars['Time']['input']>;
  joinedAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  joinedAtLT?: InputMaybe<Scalars['Time']['input']>;
  joinedAtLTE?: InputMaybe<Scalars['Time']['input']>;
  joinedAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  joinedAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  not?: InputMaybe<OrgUserWhereInput>;
  or?: InputMaybe<Array<OrgUserWhereInput>>;
  /** updated_at field predicates */
  updatedAt?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedAtLT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtLTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_by field predicates */
  updatedBy?: InputMaybe<Scalars['Int']['input']>;
  updatedByGT?: InputMaybe<Scalars['Int']['input']>;
  updatedByGTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedByLT?: InputMaybe<Scalars['Int']['input']>;
  updatedByLTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByNEQ?: InputMaybe<Scalars['Int']['input']>;
  updatedByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** user_type field predicates */
  userType?: InputMaybe<OrgUserUserType>;
  userTypeIn?: InputMaybe<Array<OrgUserUserType>>;
  userTypeNEQ?: InputMaybe<OrgUserUserType>;
  userTypeNotIn?: InputMaybe<Array<OrgUserUserType>>;
};

/**
 * OrgWhereInput is used for filtering Org objects.
 * Input was generated by ent.
 */
export type OrgWhereInput = {
  and?: InputMaybe<Array<OrgWhereInput>>;
  /** code field predicates */
  code?: InputMaybe<Scalars['String']['input']>;
  codeContains?: InputMaybe<Scalars['String']['input']>;
  codeContainsFold?: InputMaybe<Scalars['String']['input']>;
  codeEqualFold?: InputMaybe<Scalars['String']['input']>;
  codeGT?: InputMaybe<Scalars['String']['input']>;
  codeGTE?: InputMaybe<Scalars['String']['input']>;
  codeHasPrefix?: InputMaybe<Scalars['String']['input']>;
  codeHasSuffix?: InputMaybe<Scalars['String']['input']>;
  codeIn?: InputMaybe<Array<Scalars['String']['input']>>;
  codeIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  codeLT?: InputMaybe<Scalars['String']['input']>;
  codeLTE?: InputMaybe<Scalars['String']['input']>;
  codeNEQ?: InputMaybe<Scalars['String']['input']>;
  codeNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  codeNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** country_code field predicates */
  countryCode?: InputMaybe<Scalars['String']['input']>;
  countryCodeContains?: InputMaybe<Scalars['String']['input']>;
  countryCodeContainsFold?: InputMaybe<Scalars['String']['input']>;
  countryCodeEqualFold?: InputMaybe<Scalars['String']['input']>;
  countryCodeGT?: InputMaybe<Scalars['String']['input']>;
  countryCodeGTE?: InputMaybe<Scalars['String']['input']>;
  countryCodeHasPrefix?: InputMaybe<Scalars['String']['input']>;
  countryCodeHasSuffix?: InputMaybe<Scalars['String']['input']>;
  countryCodeIn?: InputMaybe<Array<Scalars['String']['input']>>;
  countryCodeIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  countryCodeLT?: InputMaybe<Scalars['String']['input']>;
  countryCodeLTE?: InputMaybe<Scalars['String']['input']>;
  countryCodeNEQ?: InputMaybe<Scalars['String']['input']>;
  countryCodeNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  countryCodeNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** created_at field predicates */
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  createdAtGT?: InputMaybe<Scalars['Time']['input']>;
  createdAtGTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  createdAtLT?: InputMaybe<Scalars['Time']['input']>;
  createdAtLTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  createdAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  /** created_by field predicates */
  createdBy?: InputMaybe<Scalars['Int']['input']>;
  createdByGT?: InputMaybe<Scalars['Int']['input']>;
  createdByGTE?: InputMaybe<Scalars['Int']['input']>;
  createdByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  createdByLT?: InputMaybe<Scalars['Int']['input']>;
  createdByLTE?: InputMaybe<Scalars['Int']['input']>;
  createdByNEQ?: InputMaybe<Scalars['Int']['input']>;
  createdByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  /** deleted_at field predicates */
  deletedAt?: InputMaybe<Scalars['Time']['input']>;
  deletedAtGT?: InputMaybe<Scalars['Time']['input']>;
  deletedAtGTE?: InputMaybe<Scalars['Time']['input']>;
  deletedAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  deletedAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  deletedAtLT?: InputMaybe<Scalars['Time']['input']>;
  deletedAtLTE?: InputMaybe<Scalars['Time']['input']>;
  deletedAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  deletedAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  deletedAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** domain field predicates */
  domain?: InputMaybe<Scalars['String']['input']>;
  domainContains?: InputMaybe<Scalars['String']['input']>;
  domainContainsFold?: InputMaybe<Scalars['String']['input']>;
  domainEqualFold?: InputMaybe<Scalars['String']['input']>;
  domainGT?: InputMaybe<Scalars['String']['input']>;
  domainGTE?: InputMaybe<Scalars['String']['input']>;
  domainHasPrefix?: InputMaybe<Scalars['String']['input']>;
  domainHasSuffix?: InputMaybe<Scalars['String']['input']>;
  domainIn?: InputMaybe<Array<Scalars['String']['input']>>;
  domainIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  domainLT?: InputMaybe<Scalars['String']['input']>;
  domainLTE?: InputMaybe<Scalars['String']['input']>;
  domainNEQ?: InputMaybe<Scalars['String']['input']>;
  domainNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  domainNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** apps edge predicates */
  hasApps?: InputMaybe<Scalars['Boolean']['input']>;
  hasAppsWith?: InputMaybe<Array<AppWhereInput>>;
  /** children edge predicates */
  hasChildren?: InputMaybe<Scalars['Boolean']['input']>;
  hasChildrenWith?: InputMaybe<Array<OrgWhereInput>>;
  /** file_identities edge predicates */
  hasFileIdentities?: InputMaybe<Scalars['Boolean']['input']>;
  hasFileIdentitiesWith?: InputMaybe<Array<FileIdentityWhereInput>>;
  /** org_quota edge predicates */
  hasOrgQuota?: InputMaybe<Scalars['Boolean']['input']>;
  hasOrgQuotaWith?: InputMaybe<Array<QuotaWhereInput>>;
  /** org_user edge predicates */
  hasOrgUser?: InputMaybe<Scalars['Boolean']['input']>;
  hasOrgUserWith?: InputMaybe<Array<OrgUserWhereInput>>;
  /** owner edge predicates */
  hasOwner?: InputMaybe<Scalars['Boolean']['input']>;
  hasOwnerWith?: InputMaybe<Array<UserWhereInput>>;
  /** parent edge predicates */
  hasParent?: InputMaybe<Scalars['Boolean']['input']>;
  hasParentWith?: InputMaybe<Array<OrgWhereInput>>;
  /** permissions edge predicates */
  hasPermissions?: InputMaybe<Scalars['Boolean']['input']>;
  hasPermissionsWith?: InputMaybe<Array<PermissionWhereInput>>;
  /** policies edge predicates */
  hasPolicies?: InputMaybe<Scalars['Boolean']['input']>;
  hasPoliciesWith?: InputMaybe<Array<OrgPolicyWhereInput>>;
  /** roles_and_groups edge predicates */
  hasRolesAndGroups?: InputMaybe<Scalars['Boolean']['input']>;
  hasRolesAndGroupsWith?: InputMaybe<Array<OrgRoleWhereInput>>;
  /** user_password_policy edge predicates */
  hasUserPasswordPolicy?: InputMaybe<Scalars['Boolean']['input']>;
  hasUserPasswordPolicyWith?: InputMaybe<Array<UserPasswordPolicyWhereInput>>;
  /** users edge predicates */
  hasUsers?: InputMaybe<Scalars['Boolean']['input']>;
  hasUsersWith?: InputMaybe<Array<UserWhereInput>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** kind field predicates */
  kind?: InputMaybe<OrgKind>;
  kindIn?: InputMaybe<Array<OrgKind>>;
  kindNEQ?: InputMaybe<OrgKind>;
  kindNotIn?: InputMaybe<Array<OrgKind>>;
  /** local_currency field predicates */
  localCurrency?: InputMaybe<Scalars['String']['input']>;
  localCurrencyContains?: InputMaybe<Scalars['String']['input']>;
  localCurrencyContainsFold?: InputMaybe<Scalars['String']['input']>;
  localCurrencyEqualFold?: InputMaybe<Scalars['String']['input']>;
  localCurrencyGT?: InputMaybe<Scalars['String']['input']>;
  localCurrencyGTE?: InputMaybe<Scalars['String']['input']>;
  localCurrencyHasPrefix?: InputMaybe<Scalars['String']['input']>;
  localCurrencyHasSuffix?: InputMaybe<Scalars['String']['input']>;
  localCurrencyIn?: InputMaybe<Array<Scalars['String']['input']>>;
  localCurrencyIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  localCurrencyLT?: InputMaybe<Scalars['String']['input']>;
  localCurrencyLTE?: InputMaybe<Scalars['String']['input']>;
  localCurrencyNEQ?: InputMaybe<Scalars['String']['input']>;
  localCurrencyNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  localCurrencyNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** name field predicates */
  name?: InputMaybe<Scalars['String']['input']>;
  nameContains?: InputMaybe<Scalars['String']['input']>;
  nameContainsFold?: InputMaybe<Scalars['String']['input']>;
  nameEqualFold?: InputMaybe<Scalars['String']['input']>;
  nameGT?: InputMaybe<Scalars['String']['input']>;
  nameGTE?: InputMaybe<Scalars['String']['input']>;
  nameHasPrefix?: InputMaybe<Scalars['String']['input']>;
  nameHasSuffix?: InputMaybe<Scalars['String']['input']>;
  nameIn?: InputMaybe<Array<Scalars['String']['input']>>;
  nameLT?: InputMaybe<Scalars['String']['input']>;
  nameLTE?: InputMaybe<Scalars['String']['input']>;
  nameNEQ?: InputMaybe<Scalars['String']['input']>;
  nameNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  not?: InputMaybe<OrgWhereInput>;
  or?: InputMaybe<Array<OrgWhereInput>>;
  /** owner_id field predicates */
  ownerID?: InputMaybe<Scalars['ID']['input']>;
  ownerIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  ownerIDIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  ownerIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  ownerIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  ownerIDNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** parent_id field predicates */
  parentID?: InputMaybe<Scalars['ID']['input']>;
  parentIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  parentIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  parentIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** path field predicates */
  path?: InputMaybe<Scalars['String']['input']>;
  pathContains?: InputMaybe<Scalars['String']['input']>;
  pathContainsFold?: InputMaybe<Scalars['String']['input']>;
  pathEqualFold?: InputMaybe<Scalars['String']['input']>;
  pathGT?: InputMaybe<Scalars['String']['input']>;
  pathGTE?: InputMaybe<Scalars['String']['input']>;
  pathHasPrefix?: InputMaybe<Scalars['String']['input']>;
  pathHasSuffix?: InputMaybe<Scalars['String']['input']>;
  pathIn?: InputMaybe<Array<Scalars['String']['input']>>;
  pathIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  pathLT?: InputMaybe<Scalars['String']['input']>;
  pathLTE?: InputMaybe<Scalars['String']['input']>;
  pathNEQ?: InputMaybe<Scalars['String']['input']>;
  pathNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  pathNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** status field predicates */
  status?: InputMaybe<OrgSimpleStatus>;
  statusIn?: InputMaybe<Array<OrgSimpleStatus>>;
  statusIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  statusNEQ?: InputMaybe<OrgSimpleStatus>;
  statusNotIn?: InputMaybe<Array<OrgSimpleStatus>>;
  statusNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** timezone field predicates */
  timezone?: InputMaybe<Scalars['String']['input']>;
  timezoneContains?: InputMaybe<Scalars['String']['input']>;
  timezoneContainsFold?: InputMaybe<Scalars['String']['input']>;
  timezoneEqualFold?: InputMaybe<Scalars['String']['input']>;
  timezoneGT?: InputMaybe<Scalars['String']['input']>;
  timezoneGTE?: InputMaybe<Scalars['String']['input']>;
  timezoneHasPrefix?: InputMaybe<Scalars['String']['input']>;
  timezoneHasSuffix?: InputMaybe<Scalars['String']['input']>;
  timezoneIn?: InputMaybe<Array<Scalars['String']['input']>>;
  timezoneIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  timezoneLT?: InputMaybe<Scalars['String']['input']>;
  timezoneLTE?: InputMaybe<Scalars['String']['input']>;
  timezoneNEQ?: InputMaybe<Scalars['String']['input']>;
  timezoneNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  timezoneNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_at field predicates */
  updatedAt?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedAtLT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtLTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_by field predicates */
  updatedBy?: InputMaybe<Scalars['Int']['input']>;
  updatedByGT?: InputMaybe<Scalars['Int']['input']>;
  updatedByGTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedByLT?: InputMaybe<Scalars['Int']['input']>;
  updatedByLTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByNEQ?: InputMaybe<Scalars['Int']['input']>;
  updatedByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByNotNil?: InputMaybe<Scalars['Boolean']['input']>;
};

/**
 * Information about pagination in a connection.
 * https://relay.dev/graphql/connections.htm#sec-undefined.PageInfo
 */
export type PageInfo = {
  __typename?: 'PageInfo';
  /** When paginating forwards, the cursor to continue. */
  endCursor?: Maybe<Scalars['Cursor']['output']>;
  /** When paginating forwards, are there more items? */
  hasNextPage: Scalars['Boolean']['output'];
  /** When paginating backwards, are there more items? */
  hasPreviousPage: Scalars['Boolean']['output'];
  /** When paginating backwards, the cursor to continue. */
  startCursor?: Maybe<Scalars['Cursor']['output']>;
};

export type Permission = Node & {
  __typename?: 'Permission';
  createdAt: Scalars['Time']['output'];
  createdBy: Scalars['Int']['output'];
  /** 生效结束时间 */
  endAt?: Maybe<Scalars['Time']['output']>;
  id: Scalars['ID']['output'];
  /** 是否允许撤销：根用户授权及系统角色授权不允许撤销 */
  isAllowRevoke: Scalars['Boolean']['output'];
  org: Org;
  /** 授权的租户 */
  orgID: Scalars['ID']['output'];
  orgPolicy: OrgPolicy;
  /** 策略 */
  orgPolicyID: Scalars['ID']['output'];
  /** 授权类型:角色,用户 */
  principalKind: PermissionPrincipalKind;
  role?: Maybe<OrgRole>;
  /** 授权类型为角色或用户组的ID */
  roleID?: Maybe<Scalars['ID']['output']>;
  /** 生效开始时间 */
  startAt?: Maybe<Scalars['Time']['output']>;
  /** 状态 */
  status?: Maybe<PermissionSimpleStatus>;
  updatedAt?: Maybe<Scalars['Time']['output']>;
  updatedBy?: Maybe<Scalars['Int']['output']>;
  user?: Maybe<User>;
  /** 授权类型为用户的ID */
  userID?: Maybe<Scalars['ID']['output']>;
};

/** A connection to a list of items. */
export type PermissionConnection = {
  __typename?: 'PermissionConnection';
  /** A list of edges. */
  edges?: Maybe<Array<Maybe<PermissionEdge>>>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** Identifies the total count of items in the connection. */
  totalCount: Scalars['Int']['output'];
};

/** An edge in a connection. */
export type PermissionEdge = {
  __typename?: 'PermissionEdge';
  /** A cursor for use in pagination. */
  cursor: Scalars['Cursor']['output'];
  /** The item at the end of the edge. */
  node?: Maybe<Permission>;
};

/** Ordering options for Permission connections */
export type PermissionOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order Permissions. */
  field: PermissionOrderField;
};

/** Properties by which Permission connections can be ordered. */
export enum PermissionOrderField {
  CreatedAt = 'createdAt'
}

/** PermissionPrincipalKind is enum for the field principal_kind */
export enum PermissionPrincipalKind {
  Role = 'role',
  User = 'user'
}

/** PermissionSimpleStatus is enum for the field status */
export enum PermissionSimpleStatus {
  Active = 'active',
  Disabled = 'disabled',
  Inactive = 'inactive',
  Processing = 'processing'
}

/**
 * PermissionWhereInput is used for filtering Permission objects.
 * Input was generated by ent.
 */
export type PermissionWhereInput = {
  and?: InputMaybe<Array<PermissionWhereInput>>;
  /** created_at field predicates */
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  createdAtGT?: InputMaybe<Scalars['Time']['input']>;
  createdAtGTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  createdAtLT?: InputMaybe<Scalars['Time']['input']>;
  createdAtLTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  createdAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  /** created_by field predicates */
  createdBy?: InputMaybe<Scalars['Int']['input']>;
  createdByGT?: InputMaybe<Scalars['Int']['input']>;
  createdByGTE?: InputMaybe<Scalars['Int']['input']>;
  createdByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  createdByLT?: InputMaybe<Scalars['Int']['input']>;
  createdByLTE?: InputMaybe<Scalars['Int']['input']>;
  createdByNEQ?: InputMaybe<Scalars['Int']['input']>;
  createdByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  /** end_at field predicates */
  endAt?: InputMaybe<Scalars['Time']['input']>;
  endAtGT?: InputMaybe<Scalars['Time']['input']>;
  endAtGTE?: InputMaybe<Scalars['Time']['input']>;
  endAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  endAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  endAtLT?: InputMaybe<Scalars['Time']['input']>;
  endAtLTE?: InputMaybe<Scalars['Time']['input']>;
  endAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  endAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  endAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** org edge predicates */
  hasOrg?: InputMaybe<Scalars['Boolean']['input']>;
  /** org_policy edge predicates */
  hasOrgPolicy?: InputMaybe<Scalars['Boolean']['input']>;
  hasOrgPolicyWith?: InputMaybe<Array<OrgPolicyWhereInput>>;
  hasOrgWith?: InputMaybe<Array<OrgWhereInput>>;
  /** role edge predicates */
  hasRole?: InputMaybe<Scalars['Boolean']['input']>;
  hasRoleWith?: InputMaybe<Array<OrgRoleWhereInput>>;
  /** user edge predicates */
  hasUser?: InputMaybe<Scalars['Boolean']['input']>;
  hasUserWith?: InputMaybe<Array<UserWhereInput>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  not?: InputMaybe<PermissionWhereInput>;
  or?: InputMaybe<Array<PermissionWhereInput>>;
  /** org_id field predicates */
  orgID?: InputMaybe<Scalars['ID']['input']>;
  orgIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  orgIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  orgIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** org_policy_id field predicates */
  orgPolicyID?: InputMaybe<Scalars['ID']['input']>;
  orgPolicyIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  orgPolicyIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  orgPolicyIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** principal_kind field predicates */
  principalKind?: InputMaybe<PermissionPrincipalKind>;
  principalKindIn?: InputMaybe<Array<PermissionPrincipalKind>>;
  principalKindNEQ?: InputMaybe<PermissionPrincipalKind>;
  principalKindNotIn?: InputMaybe<Array<PermissionPrincipalKind>>;
  /** role_id field predicates */
  roleID?: InputMaybe<Scalars['ID']['input']>;
  roleIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  roleIDIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  roleIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  roleIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  roleIDNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** start_at field predicates */
  startAt?: InputMaybe<Scalars['Time']['input']>;
  startAtGT?: InputMaybe<Scalars['Time']['input']>;
  startAtGTE?: InputMaybe<Scalars['Time']['input']>;
  startAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  startAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  startAtLT?: InputMaybe<Scalars['Time']['input']>;
  startAtLTE?: InputMaybe<Scalars['Time']['input']>;
  startAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  startAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  startAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** status field predicates */
  status?: InputMaybe<PermissionSimpleStatus>;
  statusIn?: InputMaybe<Array<PermissionSimpleStatus>>;
  statusIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  statusNEQ?: InputMaybe<PermissionSimpleStatus>;
  statusNotIn?: InputMaybe<Array<PermissionSimpleStatus>>;
  statusNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_at field predicates */
  updatedAt?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedAtLT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtLTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_by field predicates */
  updatedBy?: InputMaybe<Scalars['Int']['input']>;
  updatedByGT?: InputMaybe<Scalars['Int']['input']>;
  updatedByGTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedByLT?: InputMaybe<Scalars['Int']['input']>;
  updatedByLTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByNEQ?: InputMaybe<Scalars['Int']['input']>;
  updatedByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** user_id field predicates */
  userID?: InputMaybe<Scalars['ID']['input']>;
  userIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  userIDIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  userIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  userIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  userIDNotNil?: InputMaybe<Scalars['Boolean']['input']>;
};

export enum PolicyEffect {
  Allow = 'allow',
  Deny = 'deny'
}

export type PolicyRule = {
  __typename?: 'PolicyRule';
  actions?: Maybe<Array<Scalars['String']['output']>>;
  conditions?: Maybe<Array<Scalars['String']['output']>>;
  effect: PolicyEffect;
  resources?: Maybe<Array<Scalars['String']['output']>>;
};

export type PolicyRuleInput = {
  actions?: InputMaybe<Array<Scalars['String']['input']>>;
  conditions?: InputMaybe<Array<Scalars['String']['input']>>;
  effect: PolicyEffect;
  resources?: InputMaybe<Array<Scalars['String']['input']>>;
};

export type Query = {
  __typename?: 'Query';
  /** 检测应用登录授权 */
  appAccess: Scalars['Boolean']['output'];
  /** 检测应用登录授权,针对获取token处理 */
  appAccessForToken: Scalars['Boolean']['output'];
  /** 根据ref_code获取数据字典,用于批量获取 */
  appDictByRefCode: Array<AppDict>;
  /** 根据ref_code获取数据字典值 */
  appDictItemByRefCode: Array<AppDictItem>;
  /** 数据字典查询 */
  appDicts: AppDictConnection;
  /** 应用策略授权的组织列表 */
  appPolicyAssignedToOrgs: Array<Org>;
  /** 应用策略视图 */
  appPolicyView: Array<AppPolicyView>;
  /** 应用策略视图角色选中项 */
  appPolicyViewRoleAssigned: Array<AppPolicyView>;
  /** 获取应用资源模板 */
  appResources: AppResConnection;
  /** 应用角色授权的组织列表 */
  appRoleAssignedToOrgs: Array<Org>;
  /** 公开应用查询 */
  apps: AppConnection;
  /** 检测权限 */
  checkPermission: Scalars['Boolean']['output'];
  /**
   * 检测权限 ko-proxy使用
   * @deprecated 未来将弃用.
   */
  checkPermissionByJwt: Scalars['Boolean']['output'];
  /** 国家查询 */
  countries: CountryConnection;
  /** 货币查询 */
  currencies: CurrencyConnection;
  /** 文件凭证 */
  fileIdentities: FileIdentityConnection;
  /** 获取文件凭证 */
  fileIdentitiesForApp: Array<FileIdentityForApp>;
  /** 获取凭证AccessKeySecret */
  fileIdentityAccessKeySecret: Scalars['String']['output'];
  /** 文件来源 */
  fileSources: FileSourceConnection;
  /** 获取全局ID,开发用途 */
  globalID?: Maybe<Scalars['GID']['output']>;
  /** Fetches an object given its ID. */
  node?: Maybe<Node>;
  /** Lookup nodes by a list of IDs. */
  nodes: Array<Maybe<Node>>;
  /** 组织策略可授权的appActions */
  orgAppActions: Array<AppAction>;
  /** 获取组织应用资源模板 */
  orgAppResources: AppResConnection;
  /** 获取登录组织的文件凭证，登录组织未取到，则往上级组织获取 */
  orgFileIdentities: Array<FileIdentity>;
  /** 用户组 */
  orgGroups: OrgRoleConnection;
  /** 权限策略引用列表 */
  orgPolicyReferences: PermissionConnection;
  /** 登录用户策略视图 */
  orgPolicyView: Array<AppPolicyViewOrgPolicy>;
  /** 组织策略视图角色/用户组选中项，返回OrgPolicyIDs */
  orgPolicyViewRoleAssigned: Array<Scalars['ID']['output']>;
  /** 组织策略视图用户选择项，返回OrgPolicyIDs */
  orgPolicyViewUserAssigned: Array<Scalars['ID']['output']>;
  /** 组织策略视图，用户授权的角色/用户组选中项，返回OrgPolicyIDs */
  orgPolicyViewUserRoleAssigned: Array<Scalars['ID']['output']>;
  /** 组织回收站列表 */
  orgRecycleUsers: UserConnection;
  /** 用户组组成员 */
  orgRoleUsers: UserConnection;
  /** 角色 */
  orgRoles: OrgRoleConnection;
  /** 获取组织用户偏好 */
  orgUserPreference?: Maybe<OrgUserPreference>;
  organizations: OrgConnection;
  /** 获取父组织用户，如果是部门取所在组织用户，如果是组织取登录组织用户 */
  parentOrgUsers: UserConnection;
  /** 配额定义 */
  quotaItems: QuotaItemConnection;
  /** 配额管理 */
  quotas: QuotaConnection;
  /** 地区查询 */
  regions: RegionConnection;
  /** 用户授权的应用列表 */
  userApps: Array<App>;
  /** 用户继承用户组的权限策略 */
  userExtendGroupPolicies: PermissionConnection;
  /** 用户继承角色的权限策略 */
  userExtendRolePolicies: PermissionConnection;
  /** 用户加入的用户组 */
  userGroups: OrgRoleConnection;
  /** 成员列表 */
  userMembers: UserConnection;
  /** 用户菜单 */
  userMenus: Array<AppMenu>;
  /** 查看用户mfa信息 */
  userMfaInfo: UserMfaInfo;
  /** 用户在登录组织授权的组织角色 */
  userOrgRoles: OrgRoleConnection;
  /** 获取租户密码策略 */
  userPasswordPolicy?: Maybe<UserPasswordPolicy>;
  /** 获取用户所有权限 */
  userPermissions: Array<AppAction>;
  /** 用户加入的角色 */
  userRoles: OrgRoleConnection;
  /** 用户加入的root组织 */
  userRootOrgs: Array<Org>;
  users: UserConnection;
  /** 获取登录用户的信息 */
  viewer: User;
};


export type QueryAppAccessArgs = {
  appCode: Scalars['String']['input'];
};


export type QueryAppAccessForTokenArgs = {
  appCode: Scalars['String']['input'];
  clientID: Scalars['String']['input'];
  clientSecret: Scalars['String']['input'];
};


export type QueryAppDictByRefCodeArgs = {
  refCodes: Array<Scalars['String']['input']>;
};


export type QueryAppDictItemByRefCodeArgs = {
  refCode: Scalars['String']['input'];
};


export type QueryAppDictsArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<AppDictOrder>;
  where?: InputMaybe<AppDictWhereInput>;
};


export type QueryAppPolicyAssignedToOrgsArgs = {
  policyID: Scalars['ID']['input'];
  where?: InputMaybe<OrgWhereInput>;
};


export type QueryAppPolicyViewArgs = {
  appCode: Scalars['String']['input'];
};


export type QueryAppPolicyViewRoleAssignedArgs = {
  appRoleID: Scalars['ID']['input'];
};


export type QueryAppResourcesArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  appID: Scalars['ID']['input'];
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<AppResOrder>;
  where?: InputMaybe<AppResWhereInput>;
};


export type QueryAppRoleAssignedToOrgsArgs = {
  roleID: Scalars['ID']['input'];
  where?: InputMaybe<OrgWhereInput>;
};


export type QueryAppsArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<AppOrder>;
  where?: InputMaybe<AppWhereInput>;
};


export type QueryCheckPermissionArgs = {
  permission: Scalars['String']['input'];
};


export type QueryCheckPermissionByJwtArgs = {
  action: Scalars['String']['input'];
  appCode: Scalars['String']['input'];
  jwtStr: Scalars['String']['input'];
  orgID: Scalars['ID']['input'];
};


export type QueryCountriesArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<CountryOrder>;
  where?: InputMaybe<CountryWhereInput>;
};


export type QueryCurrenciesArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<CurrencyOrder>;
  where?: InputMaybe<CurrencyWhereInput>;
};


export type QueryFileIdentitiesArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<FileIdentityOrder>;
  where?: InputMaybe<FileIdentityWhereInput>;
};


export type QueryFileIdentitiesForAppArgs = {
  where?: InputMaybe<FileIdentityWhereInput>;
};


export type QueryFileIdentityAccessKeySecretArgs = {
  id: Scalars['ID']['input'];
};


export type QueryFileSourcesArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<FileSourceOrder>;
  where?: InputMaybe<FileSourceWhereInput>;
};


export type QueryGlobalIdArgs = {
  id: Scalars['ID']['input'];
  type: Scalars['String']['input'];
};


export type QueryNodeArgs = {
  id: Scalars['GID']['input'];
};


export type QueryNodesArgs = {
  ids: Array<Scalars['GID']['input']>;
};


export type QueryOrgAppActionsArgs = {
  appCode: Scalars['String']['input'];
  orgID: Scalars['ID']['input'];
};


export type QueryOrgAppResourcesArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  appID: Scalars['ID']['input'];
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<AppResOrder>;
  where?: InputMaybe<AppResWhereInput>;
};


export type QueryOrgGroupsArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<OrgRoleOrder>;
  where?: InputMaybe<OrgRoleWhereInput>;
};


export type QueryOrgPolicyReferencesArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<PermissionOrder>;
  policyID: Scalars['ID']['input'];
  where?: InputMaybe<PermissionWhereInput>;
};


export type QueryOrgPolicyViewArgs = {
  appCode: Scalars['String']['input'];
  orgID?: InputMaybe<Scalars['ID']['input']>;
};


export type QueryOrgPolicyViewRoleAssignedArgs = {
  appCode: Scalars['String']['input'];
  orgID?: InputMaybe<Scalars['ID']['input']>;
  orgRoleID: Scalars['ID']['input'];
};


export type QueryOrgPolicyViewUserAssignedArgs = {
  appCode: Scalars['String']['input'];
  orgID?: InputMaybe<Scalars['ID']['input']>;
  userID: Scalars['ID']['input'];
};


export type QueryOrgPolicyViewUserRoleAssignedArgs = {
  appCode: Scalars['String']['input'];
  orgID: Scalars['ID']['input'];
  userID: Scalars['ID']['input'];
};


export type QueryOrgRecycleUsersArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<UserOrder>;
  where?: InputMaybe<UserWhereInput>;
};


export type QueryOrgRoleUsersArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<UserOrder>;
  roleID: Scalars['ID']['input'];
  where?: InputMaybe<UserWhereInput>;
};


export type QueryOrgRolesArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<OrgRoleOrder>;
  where?: InputMaybe<OrgRoleWhereInput>;
};


export type QueryOrganizationsArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<OrgOrder>;
  where?: InputMaybe<OrgWhereInput>;
};


export type QueryParentOrgUsersArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<UserOrder>;
  orgID: Scalars['ID']['input'];
  where?: InputMaybe<UserWhereInput>;
};


export type QueryQuotaItemsArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<QuotaItemOrder>;
  where?: InputMaybe<QuotaItemWhereInput>;
};


export type QueryQuotasArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<QuotaOrder>;
  where?: InputMaybe<QuotaWhereInput>;
};


export type QueryRegionsArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<RegionOrder>;
  where?: InputMaybe<RegionWhereInput>;
};


export type QueryUserExtendGroupPoliciesArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<PermissionOrder>;
  orgID?: InputMaybe<Scalars['ID']['input']>;
  userID: Scalars['ID']['input'];
  where?: InputMaybe<PermissionWhereInput>;
};


export type QueryUserExtendRolePoliciesArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<PermissionOrder>;
  orgID?: InputMaybe<Scalars['ID']['input']>;
  userID: Scalars['ID']['input'];
  where?: InputMaybe<PermissionWhereInput>;
};


export type QueryUserGroupsArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<OrgRoleOrder>;
  orgID?: InputMaybe<Scalars['ID']['input']>;
  userID: Scalars['ID']['input'];
  where?: InputMaybe<OrgRoleWhereInput>;
};


export type QueryUserMembersArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<UserOrder>;
  where?: InputMaybe<UserWhereInput>;
};


export type QueryUserMenusArgs = {
  appCode: Scalars['String']['input'];
};


export type QueryUserMfaInfoArgs = {
  orgID: Scalars['ID']['input'];
  userID: Scalars['ID']['input'];
};


export type QueryUserOrgRolesArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<OrgRoleOrder>;
  where?: InputMaybe<OrgRoleWhereInput>;
};


export type QueryUserPermissionsArgs = {
  where?: InputMaybe<AppActionWhereInput>;
};


export type QueryUserRolesArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<OrgRoleOrder>;
  orgID?: InputMaybe<Scalars['ID']['input']>;
  userID: Scalars['ID']['input'];
  where?: InputMaybe<OrgRoleWhereInput>;
};


export type QueryUsersArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<UserOrder>;
  where?: InputMaybe<UserWhereInput>;
};

export type Quota = Node & {
  __typename?: 'Quota';
  createdAt: Scalars['Time']['output'];
  createdBy: Scalars['Int']['output'];
  /** 过期时间 */
  endAt?: Maybe<Scalars['Time']['output']>;
  id: Scalars['ID']['output'];
  /** 限制值 */
  limit: Scalars['Int']['output'];
  /** 配额定义 */
  quotaItem: QuotaItem;
  /** 配额项ID */
  quotaItemID: Scalars['ID']['output'];
  /** 配额关联租户 */
  quotaOrg?: Maybe<Org>;
  /** 配额关联用户 */
  quotaUser?: Maybe<User>;
  /** 生效时间 */
  startAt?: Maybe<Scalars['Time']['output']>;
  /** 租户ID,来源于root的组织ID. */
  tenantID?: Maybe<Scalars['ID']['output']>;
  updatedAt?: Maybe<Scalars['Time']['output']>;
  updatedBy?: Maybe<Scalars['Int']['output']>;
  /** 已使用值 */
  used: Scalars['Int']['output'];
  /** 来源于用户ID */
  userID?: Maybe<Scalars['ID']['output']>;
};

/** A connection to a list of items. */
export type QuotaConnection = {
  __typename?: 'QuotaConnection';
  /** A list of edges. */
  edges?: Maybe<Array<Maybe<QuotaEdge>>>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** Identifies the total count of items in the connection. */
  totalCount: Scalars['Int']['output'];
};

/** An edge in a connection. */
export type QuotaEdge = {
  __typename?: 'QuotaEdge';
  /** A cursor for use in pagination. */
  cursor: Scalars['Cursor']['output'];
  /** The item at the end of the edge. */
  node?: Maybe<Quota>;
};

export type QuotaItem = Node & {
  __typename?: 'QuotaItem';
  /** 是否启用 */
  active: Scalars['Boolean']['output'];
  /** 配额项代码,如: users,orgs */
  code: Scalars['String']['output'];
  createdAt: Scalars['Time']['output'];
  createdBy: Scalars['Int']['output'];
  /** 默认限制值 */
  defaultLimit?: Maybe<Scalars['Int']['output']>;
  /** 描述 */
  description?: Maybe<Scalars['String']['output']>;
  id: Scalars['ID']['output'];
  /** 配额项名称 */
  name: Scalars['String']['output'];
  quota: QuotaConnection;
  /** 资源类型 */
  resourceType: QuotaItemResourceType;
  /** 单位,如: 个,MB,GB */
  unit?: Maybe<Scalars['String']['output']>;
  updatedAt?: Maybe<Scalars['Time']['output']>;
  updatedBy?: Maybe<Scalars['Int']['output']>;
};


export type QuotaItemQuotaArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<QuotaOrder>;
  where?: InputMaybe<QuotaWhereInput>;
};

/** A connection to a list of items. */
export type QuotaItemConnection = {
  __typename?: 'QuotaItemConnection';
  /** A list of edges. */
  edges?: Maybe<Array<Maybe<QuotaItemEdge>>>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** Identifies the total count of items in the connection. */
  totalCount: Scalars['Int']['output'];
};

/** An edge in a connection. */
export type QuotaItemEdge = {
  __typename?: 'QuotaItemEdge';
  /** A cursor for use in pagination. */
  cursor: Scalars['Cursor']['output'];
  /** The item at the end of the edge. */
  node?: Maybe<QuotaItem>;
};

/** Ordering options for QuotaItem connections */
export type QuotaItemOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order QuotaItems. */
  field: QuotaItemOrderField;
};

/** Properties by which QuotaItem connections can be ordered. */
export enum QuotaItemOrderField {
  CreatedAt = 'createdAt'
}

/** QuotaItemResourceType is enum for the field resource_type */
export enum QuotaItemResourceType {
  Network = 'network',
  Number = 'number',
  Storage = 'storage'
}

/**
 * QuotaItemWhereInput is used for filtering QuotaItem objects.
 * Input was generated by ent.
 */
export type QuotaItemWhereInput = {
  /** active field predicates */
  active?: InputMaybe<Scalars['Boolean']['input']>;
  activeNEQ?: InputMaybe<Scalars['Boolean']['input']>;
  and?: InputMaybe<Array<QuotaItemWhereInput>>;
  /** code field predicates */
  code?: InputMaybe<Scalars['String']['input']>;
  codeContains?: InputMaybe<Scalars['String']['input']>;
  codeContainsFold?: InputMaybe<Scalars['String']['input']>;
  codeEqualFold?: InputMaybe<Scalars['String']['input']>;
  codeGT?: InputMaybe<Scalars['String']['input']>;
  codeGTE?: InputMaybe<Scalars['String']['input']>;
  codeHasPrefix?: InputMaybe<Scalars['String']['input']>;
  codeHasSuffix?: InputMaybe<Scalars['String']['input']>;
  codeIn?: InputMaybe<Array<Scalars['String']['input']>>;
  codeLT?: InputMaybe<Scalars['String']['input']>;
  codeLTE?: InputMaybe<Scalars['String']['input']>;
  codeNEQ?: InputMaybe<Scalars['String']['input']>;
  codeNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  /** created_at field predicates */
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  createdAtGT?: InputMaybe<Scalars['Time']['input']>;
  createdAtGTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  createdAtLT?: InputMaybe<Scalars['Time']['input']>;
  createdAtLTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  createdAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  /** created_by field predicates */
  createdBy?: InputMaybe<Scalars['Int']['input']>;
  createdByGT?: InputMaybe<Scalars['Int']['input']>;
  createdByGTE?: InputMaybe<Scalars['Int']['input']>;
  createdByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  createdByLT?: InputMaybe<Scalars['Int']['input']>;
  createdByLTE?: InputMaybe<Scalars['Int']['input']>;
  createdByNEQ?: InputMaybe<Scalars['Int']['input']>;
  createdByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  /** default_limit field predicates */
  defaultLimit?: InputMaybe<Scalars['Int']['input']>;
  defaultLimitGT?: InputMaybe<Scalars['Int']['input']>;
  defaultLimitGTE?: InputMaybe<Scalars['Int']['input']>;
  defaultLimitIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  defaultLimitIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  defaultLimitLT?: InputMaybe<Scalars['Int']['input']>;
  defaultLimitLTE?: InputMaybe<Scalars['Int']['input']>;
  defaultLimitNEQ?: InputMaybe<Scalars['Int']['input']>;
  defaultLimitNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  defaultLimitNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** quota edge predicates */
  hasQuota?: InputMaybe<Scalars['Boolean']['input']>;
  hasQuotaWith?: InputMaybe<Array<QuotaWhereInput>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** name field predicates */
  name?: InputMaybe<Scalars['String']['input']>;
  nameContains?: InputMaybe<Scalars['String']['input']>;
  nameContainsFold?: InputMaybe<Scalars['String']['input']>;
  nameEqualFold?: InputMaybe<Scalars['String']['input']>;
  nameGT?: InputMaybe<Scalars['String']['input']>;
  nameGTE?: InputMaybe<Scalars['String']['input']>;
  nameHasPrefix?: InputMaybe<Scalars['String']['input']>;
  nameHasSuffix?: InputMaybe<Scalars['String']['input']>;
  nameIn?: InputMaybe<Array<Scalars['String']['input']>>;
  nameLT?: InputMaybe<Scalars['String']['input']>;
  nameLTE?: InputMaybe<Scalars['String']['input']>;
  nameNEQ?: InputMaybe<Scalars['String']['input']>;
  nameNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  not?: InputMaybe<QuotaItemWhereInput>;
  or?: InputMaybe<Array<QuotaItemWhereInput>>;
  /** resource_type field predicates */
  resourceType?: InputMaybe<QuotaItemResourceType>;
  resourceTypeIn?: InputMaybe<Array<QuotaItemResourceType>>;
  resourceTypeNEQ?: InputMaybe<QuotaItemResourceType>;
  resourceTypeNotIn?: InputMaybe<Array<QuotaItemResourceType>>;
  /** updated_at field predicates */
  updatedAt?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedAtLT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtLTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_by field predicates */
  updatedBy?: InputMaybe<Scalars['Int']['input']>;
  updatedByGT?: InputMaybe<Scalars['Int']['input']>;
  updatedByGTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedByLT?: InputMaybe<Scalars['Int']['input']>;
  updatedByLTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByNEQ?: InputMaybe<Scalars['Int']['input']>;
  updatedByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByNotNil?: InputMaybe<Scalars['Boolean']['input']>;
};

/** Ordering options for Quota connections */
export type QuotaOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order QuotaSlice. */
  field: QuotaOrderField;
};

/** Properties by which Quota connections can be ordered. */
export enum QuotaOrderField {
  CreatedAt = 'createdAt'
}

/**
 * QuotaWhereInput is used for filtering Quota objects.
 * Input was generated by ent.
 */
export type QuotaWhereInput = {
  and?: InputMaybe<Array<QuotaWhereInput>>;
  /** created_at field predicates */
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  createdAtGT?: InputMaybe<Scalars['Time']['input']>;
  createdAtGTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  createdAtLT?: InputMaybe<Scalars['Time']['input']>;
  createdAtLTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  createdAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  /** created_by field predicates */
  createdBy?: InputMaybe<Scalars['Int']['input']>;
  createdByGT?: InputMaybe<Scalars['Int']['input']>;
  createdByGTE?: InputMaybe<Scalars['Int']['input']>;
  createdByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  createdByLT?: InputMaybe<Scalars['Int']['input']>;
  createdByLTE?: InputMaybe<Scalars['Int']['input']>;
  createdByNEQ?: InputMaybe<Scalars['Int']['input']>;
  createdByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  /** end_at field predicates */
  endAt?: InputMaybe<Scalars['Time']['input']>;
  endAtGT?: InputMaybe<Scalars['Time']['input']>;
  endAtGTE?: InputMaybe<Scalars['Time']['input']>;
  endAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  endAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  endAtLT?: InputMaybe<Scalars['Time']['input']>;
  endAtLTE?: InputMaybe<Scalars['Time']['input']>;
  endAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  endAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  endAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** quota_item edge predicates */
  hasQuotaItem?: InputMaybe<Scalars['Boolean']['input']>;
  hasQuotaItemWith?: InputMaybe<Array<QuotaItemWhereInput>>;
  /** quota_org edge predicates */
  hasQuotaOrg?: InputMaybe<Scalars['Boolean']['input']>;
  hasQuotaOrgWith?: InputMaybe<Array<OrgWhereInput>>;
  /** quota_user edge predicates */
  hasQuotaUser?: InputMaybe<Scalars['Boolean']['input']>;
  hasQuotaUserWith?: InputMaybe<Array<UserWhereInput>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  not?: InputMaybe<QuotaWhereInput>;
  or?: InputMaybe<Array<QuotaWhereInput>>;
  /** quota_item_id field predicates */
  quotaItemID?: InputMaybe<Scalars['ID']['input']>;
  quotaItemIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  quotaItemIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  quotaItemIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** start_at field predicates */
  startAt?: InputMaybe<Scalars['Time']['input']>;
  startAtGT?: InputMaybe<Scalars['Time']['input']>;
  startAtGTE?: InputMaybe<Scalars['Time']['input']>;
  startAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  startAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  startAtLT?: InputMaybe<Scalars['Time']['input']>;
  startAtLTE?: InputMaybe<Scalars['Time']['input']>;
  startAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  startAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  startAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** tenant_id field predicates */
  tenantID?: InputMaybe<Scalars['ID']['input']>;
  tenantIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  tenantIDIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  tenantIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  tenantIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  tenantIDNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_at field predicates */
  updatedAt?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedAtLT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtLTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_by field predicates */
  updatedBy?: InputMaybe<Scalars['Int']['input']>;
  updatedByGT?: InputMaybe<Scalars['Int']['input']>;
  updatedByGTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedByLT?: InputMaybe<Scalars['Int']['input']>;
  updatedByLTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByNEQ?: InputMaybe<Scalars['Int']['input']>;
  updatedByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** user_id field predicates */
  userID?: InputMaybe<Scalars['ID']['input']>;
  userIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  userIDIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  userIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  userIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  userIDNotNil?: InputMaybe<Scalars['Boolean']['input']>;
};

export type Region = Node & {
  __typename?: 'Region';
  children?: Maybe<Array<Region>>;
  country?: Maybe<Country>;
  /** 国家id */
  countryID?: Maybe<Scalars['ID']['output']>;
  createdAt: Scalars['Time']['output'];
  createdBy: Scalars['Int']['output'];
  displaySort?: Maybe<Scalars['Int']['output']>;
  id: Scalars['ID']['output'];
  /** 地区中文名称 */
  name?: Maybe<Scalars['String']['output']>;
  /** 地区英文名称 */
  nameEn?: Maybe<Scalars['String']['output']>;
  parent?: Maybe<Region>;
  /** 父id，0为顶级 */
  parentID?: Maybe<Scalars['ID']['output']>;
  /** 编码 */
  shortCode?: Maybe<Scalars['String']['output']>;
  /** 状态 */
  status?: Maybe<RegionSimpleStatus>;
  updatedAt?: Maybe<Scalars['Time']['output']>;
  updatedBy?: Maybe<Scalars['Int']['output']>;
  /** 邮政编码 */
  zipCode?: Maybe<Scalars['String']['output']>;
};

/** A connection to a list of items. */
export type RegionConnection = {
  __typename?: 'RegionConnection';
  /** A list of edges. */
  edges?: Maybe<Array<Maybe<RegionEdge>>>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** Identifies the total count of items in the connection. */
  totalCount: Scalars['Int']['output'];
};

/** An edge in a connection. */
export type RegionEdge = {
  __typename?: 'RegionEdge';
  /** A cursor for use in pagination. */
  cursor: Scalars['Cursor']['output'];
  /** The item at the end of the edge. */
  node?: Maybe<Region>;
};

/** Ordering options for Region connections */
export type RegionOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order Regions. */
  field: RegionOrderField;
};

/** Properties by which Region connections can be ordered. */
export enum RegionOrderField {
  CreatedAt = 'createdAt',
  DisplaySort = 'displaySort'
}

/** RegionSimpleStatus is enum for the field status */
export enum RegionSimpleStatus {
  Active = 'active',
  Disabled = 'disabled',
  Inactive = 'inactive',
  Processing = 'processing'
}

/**
 * RegionWhereInput is used for filtering Region objects.
 * Input was generated by ent.
 */
export type RegionWhereInput = {
  and?: InputMaybe<Array<RegionWhereInput>>;
  /** country_id field predicates */
  countryID?: InputMaybe<Scalars['ID']['input']>;
  countryIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  countryIDIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  countryIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  countryIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  countryIDNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** created_at field predicates */
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  createdAtGT?: InputMaybe<Scalars['Time']['input']>;
  createdAtGTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  createdAtLT?: InputMaybe<Scalars['Time']['input']>;
  createdAtLTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  createdAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  /** created_by field predicates */
  createdBy?: InputMaybe<Scalars['Int']['input']>;
  createdByGT?: InputMaybe<Scalars['Int']['input']>;
  createdByGTE?: InputMaybe<Scalars['Int']['input']>;
  createdByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  createdByLT?: InputMaybe<Scalars['Int']['input']>;
  createdByLTE?: InputMaybe<Scalars['Int']['input']>;
  createdByNEQ?: InputMaybe<Scalars['Int']['input']>;
  createdByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  /** children edge predicates */
  hasChildren?: InputMaybe<Scalars['Boolean']['input']>;
  hasChildrenWith?: InputMaybe<Array<RegionWhereInput>>;
  /** country edge predicates */
  hasCountry?: InputMaybe<Scalars['Boolean']['input']>;
  hasCountryWith?: InputMaybe<Array<CountryWhereInput>>;
  /** parent edge predicates */
  hasParent?: InputMaybe<Scalars['Boolean']['input']>;
  hasParentWith?: InputMaybe<Array<RegionWhereInput>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** name field predicates */
  name?: InputMaybe<Scalars['String']['input']>;
  nameContains?: InputMaybe<Scalars['String']['input']>;
  nameContainsFold?: InputMaybe<Scalars['String']['input']>;
  /** name_en field predicates */
  nameEn?: InputMaybe<Scalars['String']['input']>;
  nameEnContains?: InputMaybe<Scalars['String']['input']>;
  nameEnContainsFold?: InputMaybe<Scalars['String']['input']>;
  nameEnEqualFold?: InputMaybe<Scalars['String']['input']>;
  nameEnGT?: InputMaybe<Scalars['String']['input']>;
  nameEnGTE?: InputMaybe<Scalars['String']['input']>;
  nameEnHasPrefix?: InputMaybe<Scalars['String']['input']>;
  nameEnHasSuffix?: InputMaybe<Scalars['String']['input']>;
  nameEnIn?: InputMaybe<Array<Scalars['String']['input']>>;
  nameEnIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  nameEnLT?: InputMaybe<Scalars['String']['input']>;
  nameEnLTE?: InputMaybe<Scalars['String']['input']>;
  nameEnNEQ?: InputMaybe<Scalars['String']['input']>;
  nameEnNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  nameEnNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  nameEqualFold?: InputMaybe<Scalars['String']['input']>;
  nameGT?: InputMaybe<Scalars['String']['input']>;
  nameGTE?: InputMaybe<Scalars['String']['input']>;
  nameHasPrefix?: InputMaybe<Scalars['String']['input']>;
  nameHasSuffix?: InputMaybe<Scalars['String']['input']>;
  nameIn?: InputMaybe<Array<Scalars['String']['input']>>;
  nameIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  nameLT?: InputMaybe<Scalars['String']['input']>;
  nameLTE?: InputMaybe<Scalars['String']['input']>;
  nameNEQ?: InputMaybe<Scalars['String']['input']>;
  nameNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  nameNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  not?: InputMaybe<RegionWhereInput>;
  or?: InputMaybe<Array<RegionWhereInput>>;
  /** parent_id field predicates */
  parentID?: InputMaybe<Scalars['ID']['input']>;
  parentIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  parentIDIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  parentIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  parentIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  parentIDNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** short_code field predicates */
  shortCode?: InputMaybe<Scalars['String']['input']>;
  shortCodeContains?: InputMaybe<Scalars['String']['input']>;
  shortCodeContainsFold?: InputMaybe<Scalars['String']['input']>;
  shortCodeEqualFold?: InputMaybe<Scalars['String']['input']>;
  shortCodeGT?: InputMaybe<Scalars['String']['input']>;
  shortCodeGTE?: InputMaybe<Scalars['String']['input']>;
  shortCodeHasPrefix?: InputMaybe<Scalars['String']['input']>;
  shortCodeHasSuffix?: InputMaybe<Scalars['String']['input']>;
  shortCodeIn?: InputMaybe<Array<Scalars['String']['input']>>;
  shortCodeIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  shortCodeLT?: InputMaybe<Scalars['String']['input']>;
  shortCodeLTE?: InputMaybe<Scalars['String']['input']>;
  shortCodeNEQ?: InputMaybe<Scalars['String']['input']>;
  shortCodeNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  shortCodeNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** status field predicates */
  status?: InputMaybe<RegionSimpleStatus>;
  statusIn?: InputMaybe<Array<RegionSimpleStatus>>;
  statusIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  statusNEQ?: InputMaybe<RegionSimpleStatus>;
  statusNotIn?: InputMaybe<Array<RegionSimpleStatus>>;
  statusNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_at field predicates */
  updatedAt?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedAtLT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtLTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_by field predicates */
  updatedBy?: InputMaybe<Scalars['Int']['input']>;
  updatedByGT?: InputMaybe<Scalars['Int']['input']>;
  updatedByGTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedByLT?: InputMaybe<Scalars['Int']['input']>;
  updatedByLTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByNEQ?: InputMaybe<Scalars['Int']['input']>;
  updatedByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** zip_code field predicates */
  zipCode?: InputMaybe<Scalars['String']['input']>;
  zipCodeContains?: InputMaybe<Scalars['String']['input']>;
  zipCodeContainsFold?: InputMaybe<Scalars['String']['input']>;
  zipCodeEqualFold?: InputMaybe<Scalars['String']['input']>;
  zipCodeGT?: InputMaybe<Scalars['String']['input']>;
  zipCodeGTE?: InputMaybe<Scalars['String']['input']>;
  zipCodeHasPrefix?: InputMaybe<Scalars['String']['input']>;
  zipCodeHasSuffix?: InputMaybe<Scalars['String']['input']>;
  zipCodeIn?: InputMaybe<Array<Scalars['String']['input']>>;
  zipCodeIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  zipCodeLT?: InputMaybe<Scalars['String']['input']>;
  zipCodeLTE?: InputMaybe<Scalars['String']['input']>;
  zipCodeNEQ?: InputMaybe<Scalars['String']['input']>;
  zipCodeNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  zipCodeNotNil?: InputMaybe<Scalars['Boolean']['input']>;
};

/** 树操作类型 */
export enum TreeAction {
  /** 作为子节点 */
  Child = 'child',
  /** 下移 */
  Down = 'down',
  /** 上移 */
  Up = 'up'
}

/**
 * UpdateAppActionInput is used for update AppAction object.
 * Input was generated by ent.
 */
export type UpdateAppActionInput = {
  addMenuIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  clearComments?: InputMaybe<Scalars['Boolean']['input']>;
  clearMenus?: InputMaybe<Scalars['Boolean']['input']>;
  /** 备注 */
  comments?: InputMaybe<Scalars['String']['input']>;
  /** restful,graphql,rpc,function */
  kind?: InputMaybe<AppActionKind>;
  /** 操作方法:读,写,列表 */
  method?: InputMaybe<AppActionMethod>;
  /** 名称 */
  name?: InputMaybe<Scalars['String']['input']>;
  removeMenuIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
};

/**
 * UpdateAppDictInput is used for update AppDict object.
 * Input was generated by ent.
 */
export type UpdateAppDictInput = {
  addItemIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  clearComments?: InputMaybe<Scalars['Boolean']['input']>;
  clearItems?: InputMaybe<Scalars['Boolean']['input']>;
  /** 备注 */
  comments?: InputMaybe<Scalars['String']['input']>;
  /** 名称 */
  name?: InputMaybe<Scalars['String']['input']>;
  removeItemIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
};

/**
 * UpdateAppDictItemInput is used for update AppDictItem object.
 * Input was generated by ent.
 */
export type UpdateAppDictItemInput = {
  clearComments?: InputMaybe<Scalars['Boolean']['input']>;
  clearStatus?: InputMaybe<Scalars['Boolean']['input']>;
  /** 备注 */
  comments?: InputMaybe<Scalars['String']['input']>;
  /** 名称 */
  name?: InputMaybe<Scalars['String']['input']>;
  /** 状态 */
  status?: InputMaybe<AppDictItemSimpleStatus>;
};

/**
 * UpdateAppInput is used for update App object.
 * Input was generated by ent.
 */
export type UpdateAppInput = {
  addActionIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  addDictIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  addMenuIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  addPolicyIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  addPolicyViewIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  addResourceIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  addRoleIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** 应用ID */
  appKey?: InputMaybe<Scalars['String']['input']>;
  /** 应用密钥 */
  appSecret?: InputMaybe<Scalars['String']['input']>;
  clearActions?: InputMaybe<Scalars['Boolean']['input']>;
  clearAppKey?: InputMaybe<Scalars['Boolean']['input']>;
  clearAppSecret?: InputMaybe<Scalars['Boolean']['input']>;
  clearComments?: InputMaybe<Scalars['Boolean']['input']>;
  clearDicts?: InputMaybe<Scalars['Boolean']['input']>;
  clearLogo?: InputMaybe<Scalars['Boolean']['input']>;
  clearMenus?: InputMaybe<Scalars['Boolean']['input']>;
  clearPolicies?: InputMaybe<Scalars['Boolean']['input']>;
  clearPolicyViews?: InputMaybe<Scalars['Boolean']['input']>;
  clearRedirectURI?: InputMaybe<Scalars['Boolean']['input']>;
  clearRefreshTokenValidity?: InputMaybe<Scalars['Boolean']['input']>;
  clearResources?: InputMaybe<Scalars['Boolean']['input']>;
  clearRoles?: InputMaybe<Scalars['Boolean']['input']>;
  clearScopes?: InputMaybe<Scalars['Boolean']['input']>;
  clearStatus?: InputMaybe<Scalars['Boolean']['input']>;
  clearTokenValidity?: InputMaybe<Scalars['Boolean']['input']>;
  /** 备注 */
  comments?: InputMaybe<Scalars['String']['input']>;
  /** 应用类型 */
  kind?: InputMaybe<AppKind>;
  /** 应用图标地址 */
  logo?: InputMaybe<Scalars['String']['input']>;
  /** 名称 */
  name?: InputMaybe<Scalars['String']['input']>;
  /** 回调地址 */
  redirectURI?: InputMaybe<Scalars['String']['input']>;
  /** refresh_token有效期 */
  refreshTokenValidity?: InputMaybe<Scalars['Int']['input']>;
  removeActionIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  removeDictIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  removeMenuIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  removePolicyIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  removePolicyViewIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  removeResourceIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  removeRoleIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** 权限范围 */
  scopes?: InputMaybe<Scalars['String']['input']>;
  /** 状态 */
  status?: InputMaybe<AppSimpleStatus>;
  /** token有效期 */
  tokenValidity?: InputMaybe<Scalars['Int']['input']>;
};

/**
 * UpdateAppMenuInput is used for update AppMenu object.
 * Input was generated by ent.
 */
export type UpdateAppMenuInput = {
  actionID?: InputMaybe<Scalars['ID']['input']>;
  addChildIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  clearAction?: InputMaybe<Scalars['Boolean']['input']>;
  clearChildren?: InputMaybe<Scalars['Boolean']['input']>;
  clearComments?: InputMaybe<Scalars['Boolean']['input']>;
  clearIcon?: InputMaybe<Scalars['Boolean']['input']>;
  clearRoute?: InputMaybe<Scalars['Boolean']['input']>;
  clearStatus?: InputMaybe<Scalars['Boolean']['input']>;
  /** 备注 */
  comments?: InputMaybe<Scalars['String']['input']>;
  /** 菜单图标 */
  icon?: InputMaybe<Scalars['String']['input']>;
  /** 目录,菜单项 */
  kind?: InputMaybe<AppMenuKind>;
  /** 菜单名称 */
  name?: InputMaybe<Scalars['String']['input']>;
  parentID?: InputMaybe<Scalars['ID']['input']>;
  removeChildIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** 菜单路由 */
  route?: InputMaybe<Scalars['String']['input']>;
  /** 状态 */
  status?: InputMaybe<AppMenuSimpleStatus>;
};

/**
 * UpdateAppPolicyInput is used for update AppPolicy object.
 * Input was generated by ent.
 */
export type UpdateAppPolicyInput = {
  addOrgPolicyIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  addPolicyViewIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  addRoleIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  appendRules?: InputMaybe<Array<InputMaybe<PolicyRuleInput>>>;
  /** 标识是否自动授予到账户 */
  autoGrant?: InputMaybe<Scalars['Boolean']['input']>;
  clearComments?: InputMaybe<Scalars['Boolean']['input']>;
  clearOrgPolicies?: InputMaybe<Scalars['Boolean']['input']>;
  clearPolicyViews?: InputMaybe<Scalars['Boolean']['input']>;
  clearRoles?: InputMaybe<Scalars['Boolean']['input']>;
  clearStatus?: InputMaybe<Scalars['Boolean']['input']>;
  /** 描述 */
  comments?: InputMaybe<Scalars['String']['input']>;
  /** 分类：app-应用策略、view-策略视图 */
  kind?: InputMaybe<AppPolicyKind>;
  /** 策略名称 */
  name?: InputMaybe<Scalars['String']['input']>;
  removeOrgPolicyIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  removePolicyViewIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  removeRoleIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** 策略规则 */
  rules?: InputMaybe<Array<InputMaybe<PolicyRuleInput>>>;
  /** 状态 */
  status?: InputMaybe<AppPolicySimpleStatus>;
};

/**
 * UpdateAppPolicyViewInput is used for update AppPolicyView object.
 * Input was generated by ent.
 */
export type UpdateAppPolicyViewInput = {
  addChildIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  appPolicyID?: InputMaybe<Scalars['ID']['input']>;
  clearAppPolicy?: InputMaybe<Scalars['Boolean']['input']>;
  clearChildren?: InputMaybe<Scalars['Boolean']['input']>;
  clearComments?: InputMaybe<Scalars['Boolean']['input']>;
  /** 描述 */
  comments?: InputMaybe<Scalars['String']['input']>;
  /** 分类：dir-目录、policy-权限策略 */
  kind?: InputMaybe<AppPolicyViewKind>;
  /** 名称 */
  name?: InputMaybe<Scalars['String']['input']>;
  parentID?: InputMaybe<Scalars['ID']['input']>;
  removeChildIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
};

/**
 * UpdateAppResInput is used for update AppRes object.
 * Input was generated by ent.
 */
export type UpdateAppResInput = {
  /** 资源名称 */
  name?: InputMaybe<Scalars['String']['input']>;
};

/**
 * UpdateAppRoleInput is used for update AppRole object.
 * Input was generated by ent.
 */
export type UpdateAppRoleInput = {
  /** 标识是否自动授予到账户 */
  autoGrant?: InputMaybe<Scalars['Boolean']['input']>;
  clearComments?: InputMaybe<Scalars['Boolean']['input']>;
  /** 备注 */
  comments?: InputMaybe<Scalars['String']['input']>;
  /** 授权后是否可编辑 */
  editable?: InputMaybe<Scalars['Boolean']['input']>;
  /** 角色名称 */
  name?: InputMaybe<Scalars['String']['input']>;
};

/**
 * UpdateCountryInput is used for update Country object.
 * Input was generated by ent.
 */
export type UpdateCountryInput = {
  addRegionIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  clearName?: InputMaybe<Scalars['Boolean']['input']>;
  clearNameEn?: InputMaybe<Scalars['Boolean']['input']>;
  clearRegions?: InputMaybe<Scalars['Boolean']['input']>;
  clearStatus?: InputMaybe<Scalars['Boolean']['input']>;
  /** 国家代码 */
  code?: InputMaybe<Scalars['String']['input']>;
  /** 国家中文名称 */
  name?: InputMaybe<Scalars['String']['input']>;
  /** 国家英文名称 */
  nameEn?: InputMaybe<Scalars['String']['input']>;
  removeRegionIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** 状态 */
  status?: InputMaybe<CountrySimpleStatus>;
};

/**
 * UpdateCurrencyInput is used for update Currency object.
 * Input was generated by ent.
 */
export type UpdateCurrencyInput = {
  clearSign?: InputMaybe<Scalars['Boolean']['input']>;
  /** 代码 */
  code?: InputMaybe<Scalars['String']['input']>;
  /** 中文名称 */
  name?: InputMaybe<Scalars['String']['input']>;
  /** 货币符号 */
  sign?: InputMaybe<Scalars['String']['input']>;
  /** 状态 */
  status?: InputMaybe<CurrencySimpleStatus>;
};

/**
 * UpdateFileIdentityInput is used for update FileIdentity object.
 * Input was generated by ent.
 */
export type UpdateFileIdentityInput = {
  /** accesskey id */
  accessKeyID?: InputMaybe<Scalars['String']['input']>;
  /** accesskey secret */
  accessKeySecret?: InputMaybe<Scalars['String']['input']>;
  clearComments?: InputMaybe<Scalars['Boolean']['input']>;
  clearDurationSeconds?: InputMaybe<Scalars['Boolean']['input']>;
  clearPolicy?: InputMaybe<Scalars['Boolean']['input']>;
  /** 备注 */
  comments?: InputMaybe<Scalars['String']['input']>;
  /** STS令牌的有效期，默认3600s */
  durationSeconds?: InputMaybe<Scalars['Int']['input']>;
  /** 指定返回的STS令牌的权限的策略 */
  policy?: InputMaybe<Scalars['String']['input']>;
  /** 角色的资源名称(ARN)，用于STS */
  roleArn?: InputMaybe<Scalars['String']['input']>;
  sourceID?: InputMaybe<Scalars['ID']['input']>;
};

/**
 * UpdateFileSourceInput is used for update FileSource object.
 * Input was generated by ent.
 */
export type UpdateFileSourceInput = {
  addIdentityIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** 文件存储空间 */
  bucket?: InputMaybe<Scalars['String']['input']>;
  /** 文件存储空间地址，用于匹配url */
  bucketURL?: InputMaybe<Scalars['String']['input']>;
  clearComments?: InputMaybe<Scalars['Boolean']['input']>;
  clearIdentities?: InputMaybe<Scalars['Boolean']['input']>;
  /** 备注 */
  comments?: InputMaybe<Scalars['String']['input']>;
  /** 对外服务的访问域名 */
  endpoint?: InputMaybe<Scalars['String']['input']>;
  /** 是否禁止修改endpoint，如果是自定义域名设为true */
  endpointImmutable?: InputMaybe<Scalars['Boolean']['input']>;
  /** 文件来源 */
  kind?: InputMaybe<FileSourceKind>;
  /** 地域，数据存储的物理位置 */
  region?: InputMaybe<Scalars['String']['input']>;
  removeIdentityIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** sts服务的访问域名 */
  stsEndpoint?: InputMaybe<Scalars['String']['input']>;
};

/**
 * UpdateOauthClientInput is used for update OauthClient object.
 * Input was generated by ent.
 */
export type UpdateOauthClientInput = {
  /** 授权类型 */
  grantTypes?: InputMaybe<OauthClientGrantTypes>;
  /** 名称 */
  name?: InputMaybe<Scalars['String']['input']>;
  userID?: InputMaybe<Scalars['ID']['input']>;
};

/**
 * UpdateOrgInput is used for update Org object.
 * Input was generated by ent.
 */
export type UpdateOrgInput = {
  addAppIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  addChildIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  addFileIdentityIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  addPermissionIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  addPolicyIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  addRolesAndGroupIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  addUserIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  appendCustomDomain?: InputMaybe<Array<Scalars['String']['input']>>;
  clearApps?: InputMaybe<Scalars['Boolean']['input']>;
  clearChildren?: InputMaybe<Scalars['Boolean']['input']>;
  clearCountryCode?: InputMaybe<Scalars['Boolean']['input']>;
  clearCustomDomain?: InputMaybe<Scalars['Boolean']['input']>;
  clearDomain?: InputMaybe<Scalars['Boolean']['input']>;
  clearFileIdentities?: InputMaybe<Scalars['Boolean']['input']>;
  clearLocalCurrency?: InputMaybe<Scalars['Boolean']['input']>;
  clearLogo?: InputMaybe<Scalars['Boolean']['input']>;
  clearOwner?: InputMaybe<Scalars['Boolean']['input']>;
  clearPermissions?: InputMaybe<Scalars['Boolean']['input']>;
  clearPolicies?: InputMaybe<Scalars['Boolean']['input']>;
  clearProfile?: InputMaybe<Scalars['Boolean']['input']>;
  clearRolesAndGroups?: InputMaybe<Scalars['Boolean']['input']>;
  clearStatus?: InputMaybe<Scalars['Boolean']['input']>;
  clearTimezone?: InputMaybe<Scalars['Boolean']['input']>;
  clearUserPasswordPolicy?: InputMaybe<Scalars['Boolean']['input']>;
  clearUsers?: InputMaybe<Scalars['Boolean']['input']>;
  /** 国家或地区2字码 */
  countryCode?: InputMaybe<Scalars['String']['input']>;
  /** 自定义域名 */
  customDomain?: InputMaybe<Array<Scalars['String']['input']>>;
  /** 默认域名 */
  domain?: InputMaybe<Scalars['String']['input']>;
  /** 组织本位币 */
  localCurrency?: InputMaybe<Scalars['String']['input']>;
  /** 组织图标 */
  logo?: InputMaybe<OrgLogoInput>;
  /** 组织名称 */
  name?: InputMaybe<Scalars['String']['input']>;
  ownerID?: InputMaybe<Scalars['ID']['input']>;
  parentID?: InputMaybe<Scalars['ID']['input']>;
  /** 简介 */
  profile?: InputMaybe<Scalars['String']['input']>;
  removeAppIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  removeChildIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  removeFileIdentityIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  removePermissionIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  removePolicyIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  removeRolesAndGroupIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  removeUserIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** 状态 */
  status?: InputMaybe<OrgSimpleStatus>;
  /** 时区 */
  timezone?: InputMaybe<Scalars['String']['input']>;
  userPasswordPolicyID?: InputMaybe<Scalars['ID']['input']>;
};

/**
 * UpdateOrgPolicyInput is used for update OrgPolicy object.
 * Input was generated by ent.
 */
export type UpdateOrgPolicyInput = {
  addPermissionIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  appID?: InputMaybe<Scalars['ID']['input']>;
  appPolicyID?: InputMaybe<Scalars['ID']['input']>;
  appendRules?: InputMaybe<Array<InputMaybe<PolicyRuleInput>>>;
  clearApp?: InputMaybe<Scalars['Boolean']['input']>;
  clearAppPolicy?: InputMaybe<Scalars['Boolean']['input']>;
  clearComments?: InputMaybe<Scalars['Boolean']['input']>;
  clearPermissions?: InputMaybe<Scalars['Boolean']['input']>;
  /** 描述 */
  comments?: InputMaybe<Scalars['String']['input']>;
  /** 策略名称 */
  name?: InputMaybe<Scalars['String']['input']>;
  removePermissionIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** 策略规则 */
  rules?: InputMaybe<Array<InputMaybe<PolicyRuleInput>>>;
};

/**
 * UpdateOrgRoleInput is used for update OrgRole object.
 * Input was generated by ent.
 */
export type UpdateOrgRoleInput = {
  clearComments?: InputMaybe<Scalars['Boolean']['input']>;
  /** 备注 */
  comments?: InputMaybe<Scalars['String']['input']>;
  /** 类型,group:组,role:角色 */
  kind?: InputMaybe<OrgRoleKind>;
  /** 名称 */
  name?: InputMaybe<Scalars['String']['input']>;
};

/**
 * UpdateOrgUserInput is used for update OrgUser object.
 * Input was generated by ent.
 */
export type UpdateOrgUserInput = {
  /** 在组织内的显示名称 */
  displayName?: InputMaybe<Scalars['String']['input']>;
  /** 加入时间 */
  joinedAt?: InputMaybe<Scalars['Time']['input']>;
  orgID?: InputMaybe<Scalars['ID']['input']>;
  userID?: InputMaybe<Scalars['ID']['input']>;
  /** 用户类型，区分内部及外部用户 */
  userType?: InputMaybe<OrgUserUserType>;
};

/**
 * UpdateOrgUserPreferenceInput is used for update OrgUserPreference object.
 * Input was generated by ent.
 */
export type UpdateOrgUserPreferenceInput = {
  appendMenuFavorite?: InputMaybe<Array<Scalars['ID']['input']>>;
  appendMenuRecent?: InputMaybe<Array<Scalars['ID']['input']>>;
  clearMenuFavorite?: InputMaybe<Scalars['Boolean']['input']>;
  clearMenuRecent?: InputMaybe<Scalars['Boolean']['input']>;
  /** 用户收藏菜单 */
  menuFavorite?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** 用户最近访问菜单 */
  menuRecent?: InputMaybe<Array<Scalars['ID']['input']>>;
};

/**
 * UpdatePermissionInput is used for update Permission object.
 * Input was generated by ent.
 */
export type UpdatePermissionInput = {
  clearEndAt?: InputMaybe<Scalars['Boolean']['input']>;
  clearStartAt?: InputMaybe<Scalars['Boolean']['input']>;
  clearStatus?: InputMaybe<Scalars['Boolean']['input']>;
  /** 生效结束时间 */
  endAt?: InputMaybe<Scalars['Time']['input']>;
  /** 生效开始时间 */
  startAt?: InputMaybe<Scalars['Time']['input']>;
  /** 状态 */
  status?: InputMaybe<PermissionSimpleStatus>;
};

/**
 * UpdateQuotaInput is used for update Quota object.
 * Input was generated by ent.
 */
export type UpdateQuotaInput = {
  clearEndAt?: InputMaybe<Scalars['Boolean']['input']>;
  clearQuotaOrg?: InputMaybe<Scalars['Boolean']['input']>;
  clearQuotaUser?: InputMaybe<Scalars['Boolean']['input']>;
  clearStartAt?: InputMaybe<Scalars['Boolean']['input']>;
  /** 过期时间 */
  endAt?: InputMaybe<Scalars['Time']['input']>;
  /** 限制值 */
  limit?: InputMaybe<Scalars['Int']['input']>;
  quotaItemID?: InputMaybe<Scalars['ID']['input']>;
  quotaOrgID?: InputMaybe<Scalars['ID']['input']>;
  quotaUserID?: InputMaybe<Scalars['ID']['input']>;
  /** 生效时间 */
  startAt?: InputMaybe<Scalars['Time']['input']>;
};

/**
 * UpdateQuotaItemInput is used for update QuotaItem object.
 * Input was generated by ent.
 */
export type UpdateQuotaItemInput = {
  /** 是否启用 */
  active?: InputMaybe<Scalars['Boolean']['input']>;
  addQuotumIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  clearDefaultLimit?: InputMaybe<Scalars['Boolean']['input']>;
  clearDescription?: InputMaybe<Scalars['Boolean']['input']>;
  clearQuota?: InputMaybe<Scalars['Boolean']['input']>;
  clearUnit?: InputMaybe<Scalars['Boolean']['input']>;
  /** 配额项代码,如: users,orgs */
  code?: InputMaybe<Scalars['String']['input']>;
  /** 默认限制值 */
  defaultLimit?: InputMaybe<Scalars['Int']['input']>;
  /** 描述 */
  description?: InputMaybe<Scalars['String']['input']>;
  /** 配额项名称 */
  name?: InputMaybe<Scalars['String']['input']>;
  removeQuotumIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** 资源类型 */
  resourceType?: InputMaybe<QuotaItemResourceType>;
  /** 单位,如: 个,MB,GB */
  unit?: InputMaybe<Scalars['String']['input']>;
};

/**
 * UpdateRegionInput is used for update Region object.
 * Input was generated by ent.
 */
export type UpdateRegionInput = {
  addChildIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  clearChildren?: InputMaybe<Scalars['Boolean']['input']>;
  clearCountry?: InputMaybe<Scalars['Boolean']['input']>;
  clearName?: InputMaybe<Scalars['Boolean']['input']>;
  clearNameEn?: InputMaybe<Scalars['Boolean']['input']>;
  clearParent?: InputMaybe<Scalars['Boolean']['input']>;
  clearShortCode?: InputMaybe<Scalars['Boolean']['input']>;
  clearStatus?: InputMaybe<Scalars['Boolean']['input']>;
  clearZipCode?: InputMaybe<Scalars['Boolean']['input']>;
  countryID?: InputMaybe<Scalars['ID']['input']>;
  /** 地区中文名称 */
  name?: InputMaybe<Scalars['String']['input']>;
  /** 地区英文名称 */
  nameEn?: InputMaybe<Scalars['String']['input']>;
  parentID?: InputMaybe<Scalars['ID']['input']>;
  removeChildIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** 编码 */
  shortCode?: InputMaybe<Scalars['String']['input']>;
  /** 状态 */
  status?: InputMaybe<RegionSimpleStatus>;
  /** 邮政编码 */
  zipCode?: InputMaybe<Scalars['String']['input']>;
};

/**
 * UpdateUserAddrInput is used for update UserAddr object.
 * Input was generated by ent.
 */
export type UpdateUserAddrInput = {
  /** 详细地址 */
  addr?: InputMaybe<Scalars['String']['input']>;
  clearAddr?: InputMaybe<Scalars['Boolean']['input']>;
  clearEmail?: InputMaybe<Scalars['Boolean']['input']>;
  clearFax?: InputMaybe<Scalars['Boolean']['input']>;
  clearMobile?: InputMaybe<Scalars['Boolean']['input']>;
  clearName?: InputMaybe<Scalars['Boolean']['input']>;
  clearRegion?: InputMaybe<Scalars['Boolean']['input']>;
  clearTel?: InputMaybe<Scalars['Boolean']['input']>;
  clearZipCode?: InputMaybe<Scalars['Boolean']['input']>;
  /** 邮箱 */
  email?: InputMaybe<Scalars['String']['input']>;
  /** 传真 */
  fax?: InputMaybe<Scalars['String']['input']>;
  /** 是否默认地址，类型为delivery时使用 */
  isDefault?: InputMaybe<Scalars['Boolean']['input']>;
  /** 手机 */
  mobile?: InputMaybe<Scalars['String']['input']>;
  /** 联系人名称 */
  name?: InputMaybe<Scalars['String']['input']>;
  regionID?: InputMaybe<Scalars['ID']['input']>;
  /** 电话 */
  tel?: InputMaybe<Scalars['String']['input']>;
  /** 邮编 */
  zipCode?: InputMaybe<Scalars['String']['input']>;
};

/**
 * UpdateUserDeviceInput is used for update UserDevice object.
 * Input was generated by ent.
 */
export type UpdateUserDeviceInput = {
  /** app版本 */
  appVersion?: InputMaybe<Scalars['String']['input']>;
  clearAppVersion?: InputMaybe<Scalars['Boolean']['input']>;
  clearComments?: InputMaybe<Scalars['Boolean']['input']>;
  clearDeviceModel?: InputMaybe<Scalars['Boolean']['input']>;
  clearDeviceName?: InputMaybe<Scalars['Boolean']['input']>;
  clearStatus?: InputMaybe<Scalars['Boolean']['input']>;
  clearSystemName?: InputMaybe<Scalars['Boolean']['input']>;
  clearSystemVersion?: InputMaybe<Scalars['Boolean']['input']>;
  /** 备注 */
  comments?: InputMaybe<Scalars['String']['input']>;
  /** 设备型号 */
  deviceModel?: InputMaybe<Scalars['String']['input']>;
  /** 设备名称 */
  deviceName?: InputMaybe<Scalars['String']['input']>;
  /** 设备唯一ID */
  deviceUID?: InputMaybe<Scalars['String']['input']>;
  /** 状态,可用或不可用及其他待确认状态 */
  status?: InputMaybe<UserDeviceSimpleStatus>;
  /** 系统名称 */
  systemName?: InputMaybe<Scalars['String']['input']>;
  /** 系统版本 */
  systemVersion?: InputMaybe<Scalars['String']['input']>;
};

/**
 * UpdateUserIdentityInput is used for update UserIdentity object.
 * Input was generated by ent.
 */
export type UpdateUserIdentityInput = {
  clearCode?: InputMaybe<Scalars['Boolean']['input']>;
  clearCodeExtend?: InputMaybe<Scalars['Boolean']['input']>;
  clearStatus?: InputMaybe<Scalars['Boolean']['input']>;
  /** 用户名、邮箱、手机、unionid、qq */
  code?: InputMaybe<Scalars['String']['input']>;
  /** 扩展标识码,比如微信的openID */
  codeExtend?: InputMaybe<Scalars['String']['input']>;
  /** 身份标识类型 手机、邮箱、用户名、微信、qq */
  kind?: InputMaybe<UserIdentityKind>;
  /** 状态,部分登陆方式需要验证通过才可启用 */
  status?: InputMaybe<UserIdentitySimpleStatus>;
};

/**
 * UpdateUserInput is used for update User object.
 * Input was generated by ent.
 */
export type UpdateUserInput = {
  /** 头像地址 */
  avatar?: InputMaybe<Scalars['String']['input']>;
  citizenshipID?: InputMaybe<Scalars['ID']['input']>;
  clearAvatar?: InputMaybe<Scalars['Boolean']['input']>;
  clearCitizenship?: InputMaybe<Scalars['Boolean']['input']>;
  clearComments?: InputMaybe<Scalars['Boolean']['input']>;
  clearFirstName?: InputMaybe<Scalars['Boolean']['input']>;
  clearLang?: InputMaybe<Scalars['Boolean']['input']>;
  clearLastName?: InputMaybe<Scalars['Boolean']['input']>;
  clearMiddleName?: InputMaybe<Scalars['Boolean']['input']>;
  /** 备注 */
  comments?: InputMaybe<Scalars['String']['input']>;
  /** 显示名 */
  displayName?: InputMaybe<Scalars['String']['input']>;
  /** 名字 */
  firstName?: InputMaybe<Scalars['String']['input']>;
  /** 性别 */
  gender?: InputMaybe<UserGender>;
  /** 语言 */
  lang?: InputMaybe<Scalars['String']['input']>;
  /** 姓氏 */
  lastName?: InputMaybe<Scalars['String']['input']>;
  /** 中间名 */
  middleName?: InputMaybe<Scalars['String']['input']>;
  /** 登陆名称 */
  principalName?: InputMaybe<Scalars['String']['input']>;
};

/**
 * UpdateUserLoginProfileInput is used for update UserLoginProfile object.
 * Input was generated by ent.
 */
export type UpdateUserLoginProfileInput = {
  /** 是否允许使用密码登陆控制台 */
  canLogin?: InputMaybe<Scalars['Boolean']['input']>;
  clearCanLogin?: InputMaybe<Scalars['Boolean']['input']>;
  clearPasswordReset?: InputMaybe<Scalars['Boolean']['input']>;
  /** 下次登陆时需要重置密码 */
  passwordReset?: InputMaybe<Scalars['Boolean']['input']>;
  /** 设置密码:keep-保持不变,customer-客户自行设置,auto-自动生成 */
  setKind?: InputMaybe<UserLoginProfileSetKind>;
  /** 是否开启设备认证 */
  verifyDevice?: InputMaybe<Scalars['Boolean']['input']>;
};

/**
 * UpdateUserPasswordInput is used for update UserPassword object.
 * Input was generated by ent.
 */
export type UpdateUserPasswordInput = {
  clearPassword?: InputMaybe<Scalars['Boolean']['input']>;
  clearStatus?: InputMaybe<Scalars['Boolean']['input']>;
  /** 密码 */
  password?: InputMaybe<Scalars['String']['input']>;
  /** 场景: login 普通登陆 */
  scene?: InputMaybe<UserPasswordScene>;
  /** 生效状态,默认生效 */
  status?: InputMaybe<UserPasswordSimpleStatus>;
};

/**
 * UpdateUserPasswordPolicyInput is used for update UserPasswordPolicy object.
 * Input was generated by ent.
 */
export type UpdateUserPasswordPolicyInput = {
  /** 是否允许包含用户名 */
  allowIncludeUserName?: InputMaybe<Scalars['Boolean']['input']>;
  /** 密码错误多少次出现验证码，最大5次，0代表不出现验证码 */
  captchaTimes?: InputMaybe<Scalars['Int']['input']>;
  clearAllowIncludeUserName?: InputMaybe<Scalars['Boolean']['input']>;
  clearCaptchaTimes?: InputMaybe<Scalars['Boolean']['input']>;
  clearIncludeChar?: InputMaybe<Scalars['Boolean']['input']>;
  clearIncludeElement?: InputMaybe<Scalars['Boolean']['input']>;
  clearInvalidDay?: InputMaybe<Scalars['Boolean']['input']>;
  clearInvalidLoginLimit?: InputMaybe<Scalars['Boolean']['input']>;
  clearLength?: InputMaybe<Scalars['Boolean']['input']>;
  clearRetry?: InputMaybe<Scalars['Boolean']['input']>;
  /** 最少包含的不同字符数，最多8个，0代表不限制 */
  includeChar?: InputMaybe<Scalars['Int']['input']>;
  /** 必须包含的元素，异或：1-小写字母，2-大写字母，4-数字，8-符号 */
  includeElement?: InputMaybe<Scalars['Int']['input']>;
  /** 有效天数，最大1095天，0代表不过期 */
  invalidDay?: InputMaybe<Scalars['Int']['input']>;
  /** 过期后是否限制登录 */
  invalidLoginLimit?: InputMaybe<Scalars['Boolean']['input']>;
  /** 密码最短长度，长度应在6-32位之间 */
  length?: InputMaybe<Scalars['Int']['input']>;
  /** 一小时内密码错误最多尝试次数，最大32次，0代表不限次数 */
  retry?: InputMaybe<Scalars['Int']['input']>;
};

export type User = Node & {
  __typename?: 'User';
  /** 用户联系信息 */
  addresses?: Maybe<Array<UserAddr>>;
  /** 头像地址 */
  avatar?: Maybe<Scalars['String']['output']>;
  /** 国籍信息 */
  citizenship?: Maybe<Country>;
  /** 国籍 */
  citizenshipID?: Maybe<Scalars['ID']['output']>;
  /** 备注 */
  comments?: Maybe<Scalars['String']['output']>;
  /** 地址信息 */
  contact?: Maybe<UserAddr>;
  createdAt: Scalars['Time']['output'];
  createdBy: Scalars['Int']['output'];
  /** 创建类型,邀请，注册,手工创建 */
  creationType: UserCreationType;
  deletedAt?: Maybe<Scalars['Time']['output']>;
  devices: UserDeviceConnection;
  /** 显示名 */
  displayName: Scalars['String']['output'];
  /** 名字 */
  firstName?: Maybe<Scalars['String']['output']>;
  /** 性别 */
  gender: UserGender;
  id: Scalars['ID']['output'];
  /** 用户身份标识 */
  identities?: Maybe<Array<UserIdentity>>;
  /** 是否允许解除角色授权 */
  isAllowRevokeRole: Scalars['Boolean']['output'];
  /** 是否分配role */
  isAssignOrgRole: Scalars['Boolean']['output'];
  /** 语言 */
  lang?: Maybe<Scalars['String']['output']>;
  /** 姓氏 */
  lastName?: Maybe<Scalars['String']['output']>;
  /** 登陆设置 */
  loginProfile?: Maybe<UserLoginProfile>;
  /** 中间名 */
  middleName?: Maybe<Scalars['String']['output']>;
  /** 用户AccessKey */
  oauthClients?: Maybe<Array<OauthClient>>;
  /** 组织用户类型 */
  orgUserType: OrgUserUserType;
  permissions: PermissionConnection;
  /** 登陆名称 */
  principalName: Scalars['String']['output'];
  /** 注册时IP */
  registerIP: Scalars['String']['output'];
  /** 状态 */
  status?: Maybe<UserSimpleStatus>;
  updatedAt?: Maybe<Scalars['Time']['output']>;
  updatedBy?: Maybe<Scalars['Int']['output']>;
  /** 用户配额 */
  userQuota?: Maybe<Array<Quota>>;
  /** 用户类型 */
  userType: UserUserType;
};


export type UserDevicesArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<UserDeviceOrder>;
  where?: InputMaybe<UserDeviceWhereInput>;
};


export type UserIsAllowRevokeRoleArgs = {
  orgRoleID: Scalars['ID']['input'];
};


export type UserIsAssignOrgRoleArgs = {
  orgRoleID: Scalars['ID']['input'];
};


export type UserOrgUserTypeArgs = {
  orgID: Scalars['ID']['input'];
};


export type UserPermissionsArgs = {
  after?: InputMaybe<Scalars['Cursor']['input']>;
  before?: InputMaybe<Scalars['Cursor']['input']>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<PermissionOrder>;
  where?: InputMaybe<PermissionWhereInput>;
};

export type UserAddr = Node & {
  __typename?: 'UserAddr';
  /** 详细地址 */
  addr?: Maybe<Scalars['String']['output']>;
  /** 地址类型，contact：基本信息，delivery：收货地址 */
  addrType: UserAddrAddrType;
  createdAt: Scalars['Time']['output'];
  createdBy: Scalars['Int']['output'];
  /** 邮箱 */
  email?: Maybe<Scalars['String']['output']>;
  /** 传真 */
  fax?: Maybe<Scalars['String']['output']>;
  id: Scalars['ID']['output'];
  /** 是否默认地址，类型为delivery时使用 */
  isDefault: Scalars['Boolean']['output'];
  /** 手机 */
  mobile?: Maybe<Scalars['String']['output']>;
  /** 联系人名称 */
  name?: Maybe<Scalars['String']['output']>;
  /** 地区信息 */
  region?: Maybe<Region>;
  /** 地址地区：市 */
  regionID?: Maybe<Scalars['ID']['output']>;
  /** 电话 */
  tel?: Maybe<Scalars['String']['output']>;
  updatedAt?: Maybe<Scalars['Time']['output']>;
  updatedBy?: Maybe<Scalars['Int']['output']>;
  user?: Maybe<User>;
  userID?: Maybe<Scalars['ID']['output']>;
  /** 邮编 */
  zipCode?: Maybe<Scalars['String']['output']>;
};

/** UserAddrAddrType is enum for the field addr_type */
export enum UserAddrAddrType {
  Contact = 'contact',
  Delivery = 'delivery'
}

/** Ordering options for UserAddr connections */
export type UserAddrOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order UserAddrs. */
  field: UserAddrOrderField;
};

/** Properties by which UserAddr connections can be ordered. */
export enum UserAddrOrderField {
  CreatedAt = 'createdAt'
}

/**
 * UserAddrWhereInput is used for filtering UserAddr objects.
 * Input was generated by ent.
 */
export type UserAddrWhereInput = {
  /** addr field predicates */
  addr?: InputMaybe<Scalars['String']['input']>;
  addrContains?: InputMaybe<Scalars['String']['input']>;
  addrContainsFold?: InputMaybe<Scalars['String']['input']>;
  addrEqualFold?: InputMaybe<Scalars['String']['input']>;
  addrGT?: InputMaybe<Scalars['String']['input']>;
  addrGTE?: InputMaybe<Scalars['String']['input']>;
  addrHasPrefix?: InputMaybe<Scalars['String']['input']>;
  addrHasSuffix?: InputMaybe<Scalars['String']['input']>;
  addrIn?: InputMaybe<Array<Scalars['String']['input']>>;
  addrIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  addrLT?: InputMaybe<Scalars['String']['input']>;
  addrLTE?: InputMaybe<Scalars['String']['input']>;
  addrNEQ?: InputMaybe<Scalars['String']['input']>;
  addrNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  addrNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** addr_type field predicates */
  addrType?: InputMaybe<UserAddrAddrType>;
  addrTypeIn?: InputMaybe<Array<UserAddrAddrType>>;
  addrTypeNEQ?: InputMaybe<UserAddrAddrType>;
  addrTypeNotIn?: InputMaybe<Array<UserAddrAddrType>>;
  and?: InputMaybe<Array<UserAddrWhereInput>>;
  /** created_at field predicates */
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  createdAtGT?: InputMaybe<Scalars['Time']['input']>;
  createdAtGTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  createdAtLT?: InputMaybe<Scalars['Time']['input']>;
  createdAtLTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  createdAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  /** created_by field predicates */
  createdBy?: InputMaybe<Scalars['Int']['input']>;
  createdByGT?: InputMaybe<Scalars['Int']['input']>;
  createdByGTE?: InputMaybe<Scalars['Int']['input']>;
  createdByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  createdByLT?: InputMaybe<Scalars['Int']['input']>;
  createdByLTE?: InputMaybe<Scalars['Int']['input']>;
  createdByNEQ?: InputMaybe<Scalars['Int']['input']>;
  createdByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  /** email field predicates */
  email?: InputMaybe<Scalars['String']['input']>;
  emailContains?: InputMaybe<Scalars['String']['input']>;
  emailContainsFold?: InputMaybe<Scalars['String']['input']>;
  emailEqualFold?: InputMaybe<Scalars['String']['input']>;
  emailGT?: InputMaybe<Scalars['String']['input']>;
  emailGTE?: InputMaybe<Scalars['String']['input']>;
  emailHasPrefix?: InputMaybe<Scalars['String']['input']>;
  emailHasSuffix?: InputMaybe<Scalars['String']['input']>;
  emailIn?: InputMaybe<Array<Scalars['String']['input']>>;
  emailIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  emailLT?: InputMaybe<Scalars['String']['input']>;
  emailLTE?: InputMaybe<Scalars['String']['input']>;
  emailNEQ?: InputMaybe<Scalars['String']['input']>;
  emailNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  emailNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** fax field predicates */
  fax?: InputMaybe<Scalars['String']['input']>;
  faxContains?: InputMaybe<Scalars['String']['input']>;
  faxContainsFold?: InputMaybe<Scalars['String']['input']>;
  faxEqualFold?: InputMaybe<Scalars['String']['input']>;
  faxGT?: InputMaybe<Scalars['String']['input']>;
  faxGTE?: InputMaybe<Scalars['String']['input']>;
  faxHasPrefix?: InputMaybe<Scalars['String']['input']>;
  faxHasSuffix?: InputMaybe<Scalars['String']['input']>;
  faxIn?: InputMaybe<Array<Scalars['String']['input']>>;
  faxIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  faxLT?: InputMaybe<Scalars['String']['input']>;
  faxLTE?: InputMaybe<Scalars['String']['input']>;
  faxNEQ?: InputMaybe<Scalars['String']['input']>;
  faxNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  faxNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** region edge predicates */
  hasRegion?: InputMaybe<Scalars['Boolean']['input']>;
  hasRegionWith?: InputMaybe<Array<RegionWhereInput>>;
  /** user edge predicates */
  hasUser?: InputMaybe<Scalars['Boolean']['input']>;
  hasUserWith?: InputMaybe<Array<UserWhereInput>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** is_default field predicates */
  isDefault?: InputMaybe<Scalars['Boolean']['input']>;
  isDefaultNEQ?: InputMaybe<Scalars['Boolean']['input']>;
  /** mobile field predicates */
  mobile?: InputMaybe<Scalars['String']['input']>;
  mobileContains?: InputMaybe<Scalars['String']['input']>;
  mobileContainsFold?: InputMaybe<Scalars['String']['input']>;
  mobileEqualFold?: InputMaybe<Scalars['String']['input']>;
  mobileGT?: InputMaybe<Scalars['String']['input']>;
  mobileGTE?: InputMaybe<Scalars['String']['input']>;
  mobileHasPrefix?: InputMaybe<Scalars['String']['input']>;
  mobileHasSuffix?: InputMaybe<Scalars['String']['input']>;
  mobileIn?: InputMaybe<Array<Scalars['String']['input']>>;
  mobileIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  mobileLT?: InputMaybe<Scalars['String']['input']>;
  mobileLTE?: InputMaybe<Scalars['String']['input']>;
  mobileNEQ?: InputMaybe<Scalars['String']['input']>;
  mobileNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  mobileNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** name field predicates */
  name?: InputMaybe<Scalars['String']['input']>;
  nameContains?: InputMaybe<Scalars['String']['input']>;
  nameContainsFold?: InputMaybe<Scalars['String']['input']>;
  nameEqualFold?: InputMaybe<Scalars['String']['input']>;
  nameGT?: InputMaybe<Scalars['String']['input']>;
  nameGTE?: InputMaybe<Scalars['String']['input']>;
  nameHasPrefix?: InputMaybe<Scalars['String']['input']>;
  nameHasSuffix?: InputMaybe<Scalars['String']['input']>;
  nameIn?: InputMaybe<Array<Scalars['String']['input']>>;
  nameIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  nameLT?: InputMaybe<Scalars['String']['input']>;
  nameLTE?: InputMaybe<Scalars['String']['input']>;
  nameNEQ?: InputMaybe<Scalars['String']['input']>;
  nameNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  nameNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  not?: InputMaybe<UserAddrWhereInput>;
  or?: InputMaybe<Array<UserAddrWhereInput>>;
  /** region_id field predicates */
  regionID?: InputMaybe<Scalars['ID']['input']>;
  regionIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  regionIDIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  regionIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  regionIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  regionIDNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** tel field predicates */
  tel?: InputMaybe<Scalars['String']['input']>;
  telContains?: InputMaybe<Scalars['String']['input']>;
  telContainsFold?: InputMaybe<Scalars['String']['input']>;
  telEqualFold?: InputMaybe<Scalars['String']['input']>;
  telGT?: InputMaybe<Scalars['String']['input']>;
  telGTE?: InputMaybe<Scalars['String']['input']>;
  telHasPrefix?: InputMaybe<Scalars['String']['input']>;
  telHasSuffix?: InputMaybe<Scalars['String']['input']>;
  telIn?: InputMaybe<Array<Scalars['String']['input']>>;
  telIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  telLT?: InputMaybe<Scalars['String']['input']>;
  telLTE?: InputMaybe<Scalars['String']['input']>;
  telNEQ?: InputMaybe<Scalars['String']['input']>;
  telNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  telNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_at field predicates */
  updatedAt?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedAtLT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtLTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_by field predicates */
  updatedBy?: InputMaybe<Scalars['Int']['input']>;
  updatedByGT?: InputMaybe<Scalars['Int']['input']>;
  updatedByGTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedByLT?: InputMaybe<Scalars['Int']['input']>;
  updatedByLTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByNEQ?: InputMaybe<Scalars['Int']['input']>;
  updatedByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** user_id field predicates */
  userID?: InputMaybe<Scalars['ID']['input']>;
  userIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  userIDIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  userIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  userIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  userIDNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** zip_code field predicates */
  zipCode?: InputMaybe<Scalars['String']['input']>;
  zipCodeContains?: InputMaybe<Scalars['String']['input']>;
  zipCodeContainsFold?: InputMaybe<Scalars['String']['input']>;
  zipCodeEqualFold?: InputMaybe<Scalars['String']['input']>;
  zipCodeGT?: InputMaybe<Scalars['String']['input']>;
  zipCodeGTE?: InputMaybe<Scalars['String']['input']>;
  zipCodeHasPrefix?: InputMaybe<Scalars['String']['input']>;
  zipCodeHasSuffix?: InputMaybe<Scalars['String']['input']>;
  zipCodeIn?: InputMaybe<Array<Scalars['String']['input']>>;
  zipCodeIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  zipCodeLT?: InputMaybe<Scalars['String']['input']>;
  zipCodeLTE?: InputMaybe<Scalars['String']['input']>;
  zipCodeNEQ?: InputMaybe<Scalars['String']['input']>;
  zipCodeNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  zipCodeNotNil?: InputMaybe<Scalars['Boolean']['input']>;
};

/** A connection to a list of items. */
export type UserConnection = {
  __typename?: 'UserConnection';
  /** A list of edges. */
  edges?: Maybe<Array<Maybe<UserEdge>>>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** Identifies the total count of items in the connection. */
  totalCount: Scalars['Int']['output'];
};

/** UserCreationType is enum for the field creation_type */
export enum UserCreationType {
  Invitation = 'invitation',
  Manual = 'manual',
  Register = 'register'
}

export type UserDevice = Node & {
  __typename?: 'UserDevice';
  /** app版本 */
  appVersion?: Maybe<Scalars['String']['output']>;
  /** 备注 */
  comments?: Maybe<Scalars['String']['output']>;
  createdAt: Scalars['Time']['output'];
  createdBy: Scalars['Int']['output'];
  /** 设备型号 */
  deviceModel?: Maybe<Scalars['String']['output']>;
  /** 设备名称 */
  deviceName?: Maybe<Scalars['String']['output']>;
  /** 设备唯一ID */
  deviceUID: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  /** 状态,可用或不可用及其他待确认状态 */
  status?: Maybe<UserDeviceSimpleStatus>;
  /** 系统名称 */
  systemName?: Maybe<Scalars['String']['output']>;
  /** 系统版本 */
  systemVersion?: Maybe<Scalars['String']['output']>;
  updatedAt?: Maybe<Scalars['Time']['output']>;
  updatedBy?: Maybe<Scalars['Int']['output']>;
  user?: Maybe<User>;
  userID?: Maybe<Scalars['ID']['output']>;
};

/** A connection to a list of items. */
export type UserDeviceConnection = {
  __typename?: 'UserDeviceConnection';
  /** A list of edges. */
  edges?: Maybe<Array<Maybe<UserDeviceEdge>>>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** Identifies the total count of items in the connection. */
  totalCount: Scalars['Int']['output'];
};

/** An edge in a connection. */
export type UserDeviceEdge = {
  __typename?: 'UserDeviceEdge';
  /** A cursor for use in pagination. */
  cursor: Scalars['Cursor']['output'];
  /** The item at the end of the edge. */
  node?: Maybe<UserDevice>;
};

/** Ordering options for UserDevice connections */
export type UserDeviceOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order UserDevices. */
  field: UserDeviceOrderField;
};

/** Properties by which UserDevice connections can be ordered. */
export enum UserDeviceOrderField {
  CreatedAt = 'createdAt'
}

/** UserDeviceSimpleStatus is enum for the field status */
export enum UserDeviceSimpleStatus {
  Active = 'active',
  Disabled = 'disabled',
  Inactive = 'inactive',
  Processing = 'processing'
}

/**
 * UserDeviceWhereInput is used for filtering UserDevice objects.
 * Input was generated by ent.
 */
export type UserDeviceWhereInput = {
  and?: InputMaybe<Array<UserDeviceWhereInput>>;
  /** app_version field predicates */
  appVersion?: InputMaybe<Scalars['String']['input']>;
  appVersionContains?: InputMaybe<Scalars['String']['input']>;
  appVersionContainsFold?: InputMaybe<Scalars['String']['input']>;
  appVersionEqualFold?: InputMaybe<Scalars['String']['input']>;
  appVersionGT?: InputMaybe<Scalars['String']['input']>;
  appVersionGTE?: InputMaybe<Scalars['String']['input']>;
  appVersionHasPrefix?: InputMaybe<Scalars['String']['input']>;
  appVersionHasSuffix?: InputMaybe<Scalars['String']['input']>;
  appVersionIn?: InputMaybe<Array<Scalars['String']['input']>>;
  appVersionIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  appVersionLT?: InputMaybe<Scalars['String']['input']>;
  appVersionLTE?: InputMaybe<Scalars['String']['input']>;
  appVersionNEQ?: InputMaybe<Scalars['String']['input']>;
  appVersionNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  appVersionNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** created_at field predicates */
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  createdAtGT?: InputMaybe<Scalars['Time']['input']>;
  createdAtGTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  createdAtLT?: InputMaybe<Scalars['Time']['input']>;
  createdAtLTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  createdAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  /** created_by field predicates */
  createdBy?: InputMaybe<Scalars['Int']['input']>;
  createdByGT?: InputMaybe<Scalars['Int']['input']>;
  createdByGTE?: InputMaybe<Scalars['Int']['input']>;
  createdByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  createdByLT?: InputMaybe<Scalars['Int']['input']>;
  createdByLTE?: InputMaybe<Scalars['Int']['input']>;
  createdByNEQ?: InputMaybe<Scalars['Int']['input']>;
  createdByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  /** device_model field predicates */
  deviceModel?: InputMaybe<Scalars['String']['input']>;
  deviceModelContains?: InputMaybe<Scalars['String']['input']>;
  deviceModelContainsFold?: InputMaybe<Scalars['String']['input']>;
  deviceModelEqualFold?: InputMaybe<Scalars['String']['input']>;
  deviceModelGT?: InputMaybe<Scalars['String']['input']>;
  deviceModelGTE?: InputMaybe<Scalars['String']['input']>;
  deviceModelHasPrefix?: InputMaybe<Scalars['String']['input']>;
  deviceModelHasSuffix?: InputMaybe<Scalars['String']['input']>;
  deviceModelIn?: InputMaybe<Array<Scalars['String']['input']>>;
  deviceModelIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  deviceModelLT?: InputMaybe<Scalars['String']['input']>;
  deviceModelLTE?: InputMaybe<Scalars['String']['input']>;
  deviceModelNEQ?: InputMaybe<Scalars['String']['input']>;
  deviceModelNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  deviceModelNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** device_name field predicates */
  deviceName?: InputMaybe<Scalars['String']['input']>;
  deviceNameContains?: InputMaybe<Scalars['String']['input']>;
  deviceNameContainsFold?: InputMaybe<Scalars['String']['input']>;
  deviceNameEqualFold?: InputMaybe<Scalars['String']['input']>;
  deviceNameGT?: InputMaybe<Scalars['String']['input']>;
  deviceNameGTE?: InputMaybe<Scalars['String']['input']>;
  deviceNameHasPrefix?: InputMaybe<Scalars['String']['input']>;
  deviceNameHasSuffix?: InputMaybe<Scalars['String']['input']>;
  deviceNameIn?: InputMaybe<Array<Scalars['String']['input']>>;
  deviceNameIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  deviceNameLT?: InputMaybe<Scalars['String']['input']>;
  deviceNameLTE?: InputMaybe<Scalars['String']['input']>;
  deviceNameNEQ?: InputMaybe<Scalars['String']['input']>;
  deviceNameNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  deviceNameNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** device_uid field predicates */
  deviceUID?: InputMaybe<Scalars['String']['input']>;
  deviceUIDContains?: InputMaybe<Scalars['String']['input']>;
  deviceUIDContainsFold?: InputMaybe<Scalars['String']['input']>;
  deviceUIDEqualFold?: InputMaybe<Scalars['String']['input']>;
  deviceUIDGT?: InputMaybe<Scalars['String']['input']>;
  deviceUIDGTE?: InputMaybe<Scalars['String']['input']>;
  deviceUIDHasPrefix?: InputMaybe<Scalars['String']['input']>;
  deviceUIDHasSuffix?: InputMaybe<Scalars['String']['input']>;
  deviceUIDIn?: InputMaybe<Array<Scalars['String']['input']>>;
  deviceUIDLT?: InputMaybe<Scalars['String']['input']>;
  deviceUIDLTE?: InputMaybe<Scalars['String']['input']>;
  deviceUIDNEQ?: InputMaybe<Scalars['String']['input']>;
  deviceUIDNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  /** user edge predicates */
  hasUser?: InputMaybe<Scalars['Boolean']['input']>;
  hasUserWith?: InputMaybe<Array<UserWhereInput>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  not?: InputMaybe<UserDeviceWhereInput>;
  or?: InputMaybe<Array<UserDeviceWhereInput>>;
  /** status field predicates */
  status?: InputMaybe<UserDeviceSimpleStatus>;
  statusIn?: InputMaybe<Array<UserDeviceSimpleStatus>>;
  statusIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  statusNEQ?: InputMaybe<UserDeviceSimpleStatus>;
  statusNotIn?: InputMaybe<Array<UserDeviceSimpleStatus>>;
  statusNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** system_name field predicates */
  systemName?: InputMaybe<Scalars['String']['input']>;
  systemNameContains?: InputMaybe<Scalars['String']['input']>;
  systemNameContainsFold?: InputMaybe<Scalars['String']['input']>;
  systemNameEqualFold?: InputMaybe<Scalars['String']['input']>;
  systemNameGT?: InputMaybe<Scalars['String']['input']>;
  systemNameGTE?: InputMaybe<Scalars['String']['input']>;
  systemNameHasPrefix?: InputMaybe<Scalars['String']['input']>;
  systemNameHasSuffix?: InputMaybe<Scalars['String']['input']>;
  systemNameIn?: InputMaybe<Array<Scalars['String']['input']>>;
  systemNameIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  systemNameLT?: InputMaybe<Scalars['String']['input']>;
  systemNameLTE?: InputMaybe<Scalars['String']['input']>;
  systemNameNEQ?: InputMaybe<Scalars['String']['input']>;
  systemNameNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  systemNameNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** system_version field predicates */
  systemVersion?: InputMaybe<Scalars['String']['input']>;
  systemVersionContains?: InputMaybe<Scalars['String']['input']>;
  systemVersionContainsFold?: InputMaybe<Scalars['String']['input']>;
  systemVersionEqualFold?: InputMaybe<Scalars['String']['input']>;
  systemVersionGT?: InputMaybe<Scalars['String']['input']>;
  systemVersionGTE?: InputMaybe<Scalars['String']['input']>;
  systemVersionHasPrefix?: InputMaybe<Scalars['String']['input']>;
  systemVersionHasSuffix?: InputMaybe<Scalars['String']['input']>;
  systemVersionIn?: InputMaybe<Array<Scalars['String']['input']>>;
  systemVersionIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  systemVersionLT?: InputMaybe<Scalars['String']['input']>;
  systemVersionLTE?: InputMaybe<Scalars['String']['input']>;
  systemVersionNEQ?: InputMaybe<Scalars['String']['input']>;
  systemVersionNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  systemVersionNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_at field predicates */
  updatedAt?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedAtLT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtLTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_by field predicates */
  updatedBy?: InputMaybe<Scalars['Int']['input']>;
  updatedByGT?: InputMaybe<Scalars['Int']['input']>;
  updatedByGTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedByLT?: InputMaybe<Scalars['Int']['input']>;
  updatedByLTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByNEQ?: InputMaybe<Scalars['Int']['input']>;
  updatedByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** user_id field predicates */
  userID?: InputMaybe<Scalars['ID']['input']>;
  userIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  userIDIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  userIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  userIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  userIDNotNil?: InputMaybe<Scalars['Boolean']['input']>;
};

/** An edge in a connection. */
export type UserEdge = {
  __typename?: 'UserEdge';
  /** A cursor for use in pagination. */
  cursor: Scalars['Cursor']['output'];
  /** The item at the end of the edge. */
  node?: Maybe<User>;
};

/** UserGender is enum for the field gender */
export enum UserGender {
  Female = 'female',
  Male = 'male',
  Privacy = 'privacy'
}

export type UserIdentity = Node & {
  __typename?: 'UserIdentity';
  /** 用户名、邮箱、手机、unionid、qq */
  code?: Maybe<Scalars['String']['output']>;
  /** 扩展标识码,比如微信的openID */
  codeExtend?: Maybe<Scalars['String']['output']>;
  createdAt: Scalars['Time']['output'];
  createdBy: Scalars['Int']['output'];
  id: Scalars['ID']['output'];
  /** 身份标识类型 手机、邮箱、用户名、微信、qq */
  kind: UserIdentityKind;
  /** 状态,部分登陆方式需要验证通过才可启用 */
  status?: Maybe<UserIdentitySimpleStatus>;
  updatedAt?: Maybe<Scalars['Time']['output']>;
  updatedBy?: Maybe<Scalars['Int']['output']>;
  user?: Maybe<User>;
  userID?: Maybe<Scalars['ID']['output']>;
};

/** UserIdentityKind is enum for the field kind */
export enum UserIdentityKind {
  Email = 'email',
  Name = 'name',
  Phone = 'phone',
  Qq = 'qq',
  Wechat = 'wechat'
}

/** Ordering options for UserIdentity connections */
export type UserIdentityOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order UserIdentities. */
  field: UserIdentityOrderField;
};

/** Properties by which UserIdentity connections can be ordered. */
export enum UserIdentityOrderField {
  CreatedAt = 'createdAt'
}

/** UserIdentitySimpleStatus is enum for the field status */
export enum UserIdentitySimpleStatus {
  Active = 'active',
  Disabled = 'disabled',
  Inactive = 'inactive',
  Processing = 'processing'
}

/**
 * UserIdentityWhereInput is used for filtering UserIdentity objects.
 * Input was generated by ent.
 */
export type UserIdentityWhereInput = {
  and?: InputMaybe<Array<UserIdentityWhereInput>>;
  /** code field predicates */
  code?: InputMaybe<Scalars['String']['input']>;
  codeContains?: InputMaybe<Scalars['String']['input']>;
  codeContainsFold?: InputMaybe<Scalars['String']['input']>;
  codeEqualFold?: InputMaybe<Scalars['String']['input']>;
  /** code_extend field predicates */
  codeExtend?: InputMaybe<Scalars['String']['input']>;
  codeExtendContains?: InputMaybe<Scalars['String']['input']>;
  codeExtendContainsFold?: InputMaybe<Scalars['String']['input']>;
  codeExtendEqualFold?: InputMaybe<Scalars['String']['input']>;
  codeExtendGT?: InputMaybe<Scalars['String']['input']>;
  codeExtendGTE?: InputMaybe<Scalars['String']['input']>;
  codeExtendHasPrefix?: InputMaybe<Scalars['String']['input']>;
  codeExtendHasSuffix?: InputMaybe<Scalars['String']['input']>;
  codeExtendIn?: InputMaybe<Array<Scalars['String']['input']>>;
  codeExtendIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  codeExtendLT?: InputMaybe<Scalars['String']['input']>;
  codeExtendLTE?: InputMaybe<Scalars['String']['input']>;
  codeExtendNEQ?: InputMaybe<Scalars['String']['input']>;
  codeExtendNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  codeExtendNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  codeGT?: InputMaybe<Scalars['String']['input']>;
  codeGTE?: InputMaybe<Scalars['String']['input']>;
  codeHasPrefix?: InputMaybe<Scalars['String']['input']>;
  codeHasSuffix?: InputMaybe<Scalars['String']['input']>;
  codeIn?: InputMaybe<Array<Scalars['String']['input']>>;
  codeIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  codeLT?: InputMaybe<Scalars['String']['input']>;
  codeLTE?: InputMaybe<Scalars['String']['input']>;
  codeNEQ?: InputMaybe<Scalars['String']['input']>;
  codeNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  codeNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** created_at field predicates */
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  createdAtGT?: InputMaybe<Scalars['Time']['input']>;
  createdAtGTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  createdAtLT?: InputMaybe<Scalars['Time']['input']>;
  createdAtLTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  createdAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  /** created_by field predicates */
  createdBy?: InputMaybe<Scalars['Int']['input']>;
  createdByGT?: InputMaybe<Scalars['Int']['input']>;
  createdByGTE?: InputMaybe<Scalars['Int']['input']>;
  createdByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  createdByLT?: InputMaybe<Scalars['Int']['input']>;
  createdByLTE?: InputMaybe<Scalars['Int']['input']>;
  createdByNEQ?: InputMaybe<Scalars['Int']['input']>;
  createdByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  /** user edge predicates */
  hasUser?: InputMaybe<Scalars['Boolean']['input']>;
  hasUserWith?: InputMaybe<Array<UserWhereInput>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** kind field predicates */
  kind?: InputMaybe<UserIdentityKind>;
  kindIn?: InputMaybe<Array<UserIdentityKind>>;
  kindNEQ?: InputMaybe<UserIdentityKind>;
  kindNotIn?: InputMaybe<Array<UserIdentityKind>>;
  not?: InputMaybe<UserIdentityWhereInput>;
  or?: InputMaybe<Array<UserIdentityWhereInput>>;
  /** status field predicates */
  status?: InputMaybe<UserIdentitySimpleStatus>;
  statusIn?: InputMaybe<Array<UserIdentitySimpleStatus>>;
  statusIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  statusNEQ?: InputMaybe<UserIdentitySimpleStatus>;
  statusNotIn?: InputMaybe<Array<UserIdentitySimpleStatus>>;
  statusNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_at field predicates */
  updatedAt?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedAtLT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtLTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_by field predicates */
  updatedBy?: InputMaybe<Scalars['Int']['input']>;
  updatedByGT?: InputMaybe<Scalars['Int']['input']>;
  updatedByGTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedByLT?: InputMaybe<Scalars['Int']['input']>;
  updatedByLTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByNEQ?: InputMaybe<Scalars['Int']['input']>;
  updatedByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** user_id field predicates */
  userID?: InputMaybe<Scalars['ID']['input']>;
  userIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  userIDIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  userIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  userIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  userIDNotNil?: InputMaybe<Scalars['Boolean']['input']>;
};

export type UserLoginProfile = Node & {
  __typename?: 'UserLoginProfile';
  /** 是否允许使用密码登陆控制台 */
  canLogin?: Maybe<Scalars['Boolean']['output']>;
  createdAt: Scalars['Time']['output'];
  createdBy: Scalars['Int']['output'];
  id: Scalars['ID']['output'];
  /** 最后登陆时间 */
  lastLoginAt?: Maybe<Scalars['Time']['output']>;
  lastLoginIP?: Maybe<Scalars['String']['output']>;
  /** 是否开启多因素验证 */
  mfaEnabled?: Maybe<Scalars['Boolean']['output']>;
  /** 多因素验证状态 */
  mfaStatus?: Maybe<UserLoginProfileSimpleStatus>;
  /** 下次登陆时需要重置密码 */
  passwordReset?: Maybe<Scalars['Boolean']['output']>;
  /** 设置密码:keep-保持不变,customer-客户自行设置,auto-自动生成 */
  setKind: UserLoginProfileSetKind;
  updatedAt?: Maybe<Scalars['Time']['output']>;
  updatedBy?: Maybe<Scalars['Int']['output']>;
  user?: Maybe<User>;
  userID?: Maybe<Scalars['ID']['output']>;
  /** 是否开启设备认证 */
  verifyDevice: Scalars['Boolean']['output'];
};

/** Ordering options for UserLoginProfile connections */
export type UserLoginProfileOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order UserLoginProfiles. */
  field: UserLoginProfileOrderField;
};

/** Properties by which UserLoginProfile connections can be ordered. */
export enum UserLoginProfileOrderField {
  CreatedAt = 'createdAt'
}

/** UserLoginProfileSetKind is enum for the field set_kind */
export enum UserLoginProfileSetKind {
  Auto = 'auto',
  Customer = 'customer',
  Keep = 'keep'
}

/** UserLoginProfileSimpleStatus is enum for the field mfa_status */
export enum UserLoginProfileSimpleStatus {
  Active = 'active',
  Disabled = 'disabled',
  Inactive = 'inactive',
  Processing = 'processing'
}

/**
 * UserLoginProfileWhereInput is used for filtering UserLoginProfile objects.
 * Input was generated by ent.
 */
export type UserLoginProfileWhereInput = {
  and?: InputMaybe<Array<UserLoginProfileWhereInput>>;
  /** can_login field predicates */
  canLogin?: InputMaybe<Scalars['Boolean']['input']>;
  canLoginIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  canLoginNEQ?: InputMaybe<Scalars['Boolean']['input']>;
  canLoginNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** created_at field predicates */
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  createdAtGT?: InputMaybe<Scalars['Time']['input']>;
  createdAtGTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  createdAtLT?: InputMaybe<Scalars['Time']['input']>;
  createdAtLTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  createdAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  /** created_by field predicates */
  createdBy?: InputMaybe<Scalars['Int']['input']>;
  createdByGT?: InputMaybe<Scalars['Int']['input']>;
  createdByGTE?: InputMaybe<Scalars['Int']['input']>;
  createdByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  createdByLT?: InputMaybe<Scalars['Int']['input']>;
  createdByLTE?: InputMaybe<Scalars['Int']['input']>;
  createdByNEQ?: InputMaybe<Scalars['Int']['input']>;
  createdByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  /** user edge predicates */
  hasUser?: InputMaybe<Scalars['Boolean']['input']>;
  hasUserWith?: InputMaybe<Array<UserWhereInput>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** mfa_enabled field predicates */
  mfaEnabled?: InputMaybe<Scalars['Boolean']['input']>;
  mfaEnabledIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  mfaEnabledNEQ?: InputMaybe<Scalars['Boolean']['input']>;
  mfaEnabledNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** mfa_status field predicates */
  mfaStatus?: InputMaybe<UserLoginProfileSimpleStatus>;
  mfaStatusIn?: InputMaybe<Array<UserLoginProfileSimpleStatus>>;
  mfaStatusIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  mfaStatusNEQ?: InputMaybe<UserLoginProfileSimpleStatus>;
  mfaStatusNotIn?: InputMaybe<Array<UserLoginProfileSimpleStatus>>;
  mfaStatusNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  not?: InputMaybe<UserLoginProfileWhereInput>;
  or?: InputMaybe<Array<UserLoginProfileWhereInput>>;
  /** password_reset field predicates */
  passwordReset?: InputMaybe<Scalars['Boolean']['input']>;
  passwordResetIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  passwordResetNEQ?: InputMaybe<Scalars['Boolean']['input']>;
  passwordResetNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** set_kind field predicates */
  setKind?: InputMaybe<UserLoginProfileSetKind>;
  setKindIn?: InputMaybe<Array<UserLoginProfileSetKind>>;
  setKindNEQ?: InputMaybe<UserLoginProfileSetKind>;
  setKindNotIn?: InputMaybe<Array<UserLoginProfileSetKind>>;
  /** updated_at field predicates */
  updatedAt?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedAtLT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtLTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_by field predicates */
  updatedBy?: InputMaybe<Scalars['Int']['input']>;
  updatedByGT?: InputMaybe<Scalars['Int']['input']>;
  updatedByGTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedByLT?: InputMaybe<Scalars['Int']['input']>;
  updatedByLTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByNEQ?: InputMaybe<Scalars['Int']['input']>;
  updatedByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** user_id field predicates */
  userID?: InputMaybe<Scalars['ID']['input']>;
  userIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  userIDIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  userIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  userIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  userIDNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** verify_device field predicates */
  verifyDevice?: InputMaybe<Scalars['Boolean']['input']>;
  verifyDeviceNEQ?: InputMaybe<Scalars['Boolean']['input']>;
};

/** 查看用户mfa信息 */
export type UserMfaInfo = {
  __typename?: 'UserMfaInfo';
  /** 用户账号 */
  accountName: Scalars['String']['output'];
  /** 是否启用mfa */
  mfaEnabled: Scalars['Boolean']['output'];
  /** 二维码 */
  qrCodeUri: Scalars['String']['output'];
  /** 秘钥 */
  secret: Scalars['String']['output'];
};

/** Ordering options for User connections */
export type UserOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order Users. */
  field: UserOrderField;
};

/** Properties by which User connections can be ordered. */
export enum UserOrderField {
  CreatedAt = 'createdAt'
}

export type UserPassword = Node & {
  __typename?: 'UserPassword';
  createdAt: Scalars['Time']['output'];
  createdBy: Scalars['Int']['output'];
  id: Scalars['ID']['output'];
  /** 场景: login 普通登陆 */
  scene: UserPasswordScene;
  /** 生效状态,默认生效 */
  status?: Maybe<UserPasswordSimpleStatus>;
  updatedAt?: Maybe<Scalars['Time']['output']>;
  updatedBy?: Maybe<Scalars['Int']['output']>;
  user?: Maybe<User>;
  userID?: Maybe<Scalars['ID']['output']>;
};

/** Ordering options for UserPassword connections */
export type UserPasswordOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order UserPasswords. */
  field: UserPasswordOrderField;
};

/** Properties by which UserPassword connections can be ordered. */
export enum UserPasswordOrderField {
  CreatedAt = 'createdAt'
}

export type UserPasswordPolicy = Node & {
  __typename?: 'UserPasswordPolicy';
  /** 是否允许包含用户名 */
  allowIncludeUserName?: Maybe<Scalars['Boolean']['output']>;
  /** 密码错误多少次出现验证码，最大5次，0代表不出现验证码 */
  captchaTimes?: Maybe<Scalars['Int']['output']>;
  createdAt: Scalars['Time']['output'];
  createdBy: Scalars['Int']['output'];
  id: Scalars['ID']['output'];
  /** 最少包含的不同字符数，最多8个，0代表不限制 */
  includeChar?: Maybe<Scalars['Int']['output']>;
  /** 必须包含的元素，异或：1-小写字母，2-大写字母，4-数字，8-符号 */
  includeElement?: Maybe<Scalars['Int']['output']>;
  /** 有效天数，最大1095天，0代表不过期 */
  invalidDay?: Maybe<Scalars['Int']['output']>;
  /** 过期后是否限制登录 */
  invalidLoginLimit?: Maybe<Scalars['Boolean']['output']>;
  /** 密码最短长度，长度应在6-32位之间 */
  length?: Maybe<Scalars['Int']['output']>;
  org?: Maybe<Org>;
  /** 一小时内密码错误最多尝试次数，最大32次，0代表不限次数 */
  retry?: Maybe<Scalars['Int']['output']>;
  /** 租户id */
  tenantID?: Maybe<Scalars['ID']['output']>;
  updatedAt?: Maybe<Scalars['Time']['output']>;
  updatedBy?: Maybe<Scalars['Int']['output']>;
};

/** Ordering options for UserPasswordPolicy connections */
export type UserPasswordPolicyOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order UserPasswordPolicies. */
  field: UserPasswordPolicyOrderField;
};

/** Properties by which UserPasswordPolicy connections can be ordered. */
export enum UserPasswordPolicyOrderField {
  CreatedAt = 'createdAt'
}

/**
 * UserPasswordPolicyWhereInput is used for filtering UserPasswordPolicy objects.
 * Input was generated by ent.
 */
export type UserPasswordPolicyWhereInput = {
  /** allow_include_user_name field predicates */
  allowIncludeUserName?: InputMaybe<Scalars['Boolean']['input']>;
  allowIncludeUserNameIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  allowIncludeUserNameNEQ?: InputMaybe<Scalars['Boolean']['input']>;
  allowIncludeUserNameNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  and?: InputMaybe<Array<UserPasswordPolicyWhereInput>>;
  /** captcha_times field predicates */
  captchaTimes?: InputMaybe<Scalars['Int']['input']>;
  captchaTimesGT?: InputMaybe<Scalars['Int']['input']>;
  captchaTimesGTE?: InputMaybe<Scalars['Int']['input']>;
  captchaTimesIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  captchaTimesIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  captchaTimesLT?: InputMaybe<Scalars['Int']['input']>;
  captchaTimesLTE?: InputMaybe<Scalars['Int']['input']>;
  captchaTimesNEQ?: InputMaybe<Scalars['Int']['input']>;
  captchaTimesNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  captchaTimesNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** created_at field predicates */
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  createdAtGT?: InputMaybe<Scalars['Time']['input']>;
  createdAtGTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  createdAtLT?: InputMaybe<Scalars['Time']['input']>;
  createdAtLTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  createdAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  /** created_by field predicates */
  createdBy?: InputMaybe<Scalars['Int']['input']>;
  createdByGT?: InputMaybe<Scalars['Int']['input']>;
  createdByGTE?: InputMaybe<Scalars['Int']['input']>;
  createdByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  createdByLT?: InputMaybe<Scalars['Int']['input']>;
  createdByLTE?: InputMaybe<Scalars['Int']['input']>;
  createdByNEQ?: InputMaybe<Scalars['Int']['input']>;
  createdByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  /** org edge predicates */
  hasOrg?: InputMaybe<Scalars['Boolean']['input']>;
  hasOrgWith?: InputMaybe<Array<OrgWhereInput>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** include_char field predicates */
  includeChar?: InputMaybe<Scalars['Int']['input']>;
  includeCharGT?: InputMaybe<Scalars['Int']['input']>;
  includeCharGTE?: InputMaybe<Scalars['Int']['input']>;
  includeCharIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  includeCharIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  includeCharLT?: InputMaybe<Scalars['Int']['input']>;
  includeCharLTE?: InputMaybe<Scalars['Int']['input']>;
  includeCharNEQ?: InputMaybe<Scalars['Int']['input']>;
  includeCharNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  includeCharNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** include_element field predicates */
  includeElement?: InputMaybe<Scalars['Int']['input']>;
  includeElementGT?: InputMaybe<Scalars['Int']['input']>;
  includeElementGTE?: InputMaybe<Scalars['Int']['input']>;
  includeElementIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  includeElementIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  includeElementLT?: InputMaybe<Scalars['Int']['input']>;
  includeElementLTE?: InputMaybe<Scalars['Int']['input']>;
  includeElementNEQ?: InputMaybe<Scalars['Int']['input']>;
  includeElementNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  includeElementNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** invalid_day field predicates */
  invalidDay?: InputMaybe<Scalars['Int']['input']>;
  invalidDayGT?: InputMaybe<Scalars['Int']['input']>;
  invalidDayGTE?: InputMaybe<Scalars['Int']['input']>;
  invalidDayIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  invalidDayIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  invalidDayLT?: InputMaybe<Scalars['Int']['input']>;
  invalidDayLTE?: InputMaybe<Scalars['Int']['input']>;
  invalidDayNEQ?: InputMaybe<Scalars['Int']['input']>;
  invalidDayNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  invalidDayNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** invalid_login_limit field predicates */
  invalidLoginLimit?: InputMaybe<Scalars['Boolean']['input']>;
  invalidLoginLimitIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  invalidLoginLimitNEQ?: InputMaybe<Scalars['Boolean']['input']>;
  invalidLoginLimitNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** length field predicates */
  length?: InputMaybe<Scalars['Int']['input']>;
  lengthGT?: InputMaybe<Scalars['Int']['input']>;
  lengthGTE?: InputMaybe<Scalars['Int']['input']>;
  lengthIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  lengthIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  lengthLT?: InputMaybe<Scalars['Int']['input']>;
  lengthLTE?: InputMaybe<Scalars['Int']['input']>;
  lengthNEQ?: InputMaybe<Scalars['Int']['input']>;
  lengthNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  lengthNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  not?: InputMaybe<UserPasswordPolicyWhereInput>;
  or?: InputMaybe<Array<UserPasswordPolicyWhereInput>>;
  /** retry field predicates */
  retry?: InputMaybe<Scalars['Int']['input']>;
  retryGT?: InputMaybe<Scalars['Int']['input']>;
  retryGTE?: InputMaybe<Scalars['Int']['input']>;
  retryIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  retryIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  retryLT?: InputMaybe<Scalars['Int']['input']>;
  retryLTE?: InputMaybe<Scalars['Int']['input']>;
  retryNEQ?: InputMaybe<Scalars['Int']['input']>;
  retryNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  retryNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** tenant_id field predicates */
  tenantID?: InputMaybe<Scalars['ID']['input']>;
  tenantIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  tenantIDIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  tenantIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  tenantIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  tenantIDNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_at field predicates */
  updatedAt?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedAtLT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtLTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_by field predicates */
  updatedBy?: InputMaybe<Scalars['Int']['input']>;
  updatedByGT?: InputMaybe<Scalars['Int']['input']>;
  updatedByGTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedByLT?: InputMaybe<Scalars['Int']['input']>;
  updatedByLTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByNEQ?: InputMaybe<Scalars['Int']['input']>;
  updatedByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByNotNil?: InputMaybe<Scalars['Boolean']['input']>;
};

/** UserPasswordScene is enum for the field scene */
export enum UserPasswordScene {
  Login = 'login'
}

/** UserPasswordSimpleStatus is enum for the field status */
export enum UserPasswordSimpleStatus {
  Active = 'active',
  Disabled = 'disabled',
  Inactive = 'inactive',
  Processing = 'processing'
}

/**
 * UserPasswordWhereInput is used for filtering UserPassword objects.
 * Input was generated by ent.
 */
export type UserPasswordWhereInput = {
  and?: InputMaybe<Array<UserPasswordWhereInput>>;
  /** created_at field predicates */
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  createdAtGT?: InputMaybe<Scalars['Time']['input']>;
  createdAtGTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  createdAtLT?: InputMaybe<Scalars['Time']['input']>;
  createdAtLTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  createdAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  /** created_by field predicates */
  createdBy?: InputMaybe<Scalars['Int']['input']>;
  createdByGT?: InputMaybe<Scalars['Int']['input']>;
  createdByGTE?: InputMaybe<Scalars['Int']['input']>;
  createdByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  createdByLT?: InputMaybe<Scalars['Int']['input']>;
  createdByLTE?: InputMaybe<Scalars['Int']['input']>;
  createdByNEQ?: InputMaybe<Scalars['Int']['input']>;
  createdByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  /** user edge predicates */
  hasUser?: InputMaybe<Scalars['Boolean']['input']>;
  hasUserWith?: InputMaybe<Array<UserWhereInput>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  not?: InputMaybe<UserPasswordWhereInput>;
  or?: InputMaybe<Array<UserPasswordWhereInput>>;
  /** scene field predicates */
  scene?: InputMaybe<UserPasswordScene>;
  sceneIn?: InputMaybe<Array<UserPasswordScene>>;
  sceneNEQ?: InputMaybe<UserPasswordScene>;
  sceneNotIn?: InputMaybe<Array<UserPasswordScene>>;
  /** status field predicates */
  status?: InputMaybe<UserPasswordSimpleStatus>;
  statusIn?: InputMaybe<Array<UserPasswordSimpleStatus>>;
  statusIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  statusNEQ?: InputMaybe<UserPasswordSimpleStatus>;
  statusNotIn?: InputMaybe<Array<UserPasswordSimpleStatus>>;
  statusNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_at field predicates */
  updatedAt?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedAtLT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtLTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_by field predicates */
  updatedBy?: InputMaybe<Scalars['Int']['input']>;
  updatedByGT?: InputMaybe<Scalars['Int']['input']>;
  updatedByGTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedByLT?: InputMaybe<Scalars['Int']['input']>;
  updatedByLTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByNEQ?: InputMaybe<Scalars['Int']['input']>;
  updatedByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** user_id field predicates */
  userID?: InputMaybe<Scalars['ID']['input']>;
  userIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  userIDIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  userIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  userIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  userIDNotNil?: InputMaybe<Scalars['Boolean']['input']>;
};

/** UserSimpleStatus is enum for the field status */
export enum UserSimpleStatus {
  Active = 'active',
  Disabled = 'disabled',
  Inactive = 'inactive',
  Processing = 'processing'
}

/** UserUserType is enum for the field user_type */
export enum UserUserType {
  Account = 'account',
  Member = 'member'
}

/**
 * UserWhereInput is used for filtering User objects.
 * Input was generated by ent.
 */
export type UserWhereInput = {
  and?: InputMaybe<Array<UserWhereInput>>;
  /** citizenship_id field predicates */
  citizenshipID?: InputMaybe<Scalars['ID']['input']>;
  citizenshipIDIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  citizenshipIDIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  citizenshipIDNEQ?: InputMaybe<Scalars['ID']['input']>;
  citizenshipIDNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  citizenshipIDNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** created_at field predicates */
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  createdAtGT?: InputMaybe<Scalars['Time']['input']>;
  createdAtGTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  createdAtLT?: InputMaybe<Scalars['Time']['input']>;
  createdAtLTE?: InputMaybe<Scalars['Time']['input']>;
  createdAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  createdAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  /** created_by field predicates */
  createdBy?: InputMaybe<Scalars['Int']['input']>;
  createdByGT?: InputMaybe<Scalars['Int']['input']>;
  createdByGTE?: InputMaybe<Scalars['Int']['input']>;
  createdByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  createdByLT?: InputMaybe<Scalars['Int']['input']>;
  createdByLTE?: InputMaybe<Scalars['Int']['input']>;
  createdByNEQ?: InputMaybe<Scalars['Int']['input']>;
  createdByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  /** creation_type field predicates */
  creationType?: InputMaybe<UserCreationType>;
  creationTypeIn?: InputMaybe<Array<UserCreationType>>;
  creationTypeNEQ?: InputMaybe<UserCreationType>;
  creationTypeNotIn?: InputMaybe<Array<UserCreationType>>;
  /** deleted_at field predicates */
  deletedAt?: InputMaybe<Scalars['Time']['input']>;
  deletedAtGT?: InputMaybe<Scalars['Time']['input']>;
  deletedAtGTE?: InputMaybe<Scalars['Time']['input']>;
  deletedAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  deletedAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  deletedAtLT?: InputMaybe<Scalars['Time']['input']>;
  deletedAtLTE?: InputMaybe<Scalars['Time']['input']>;
  deletedAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  deletedAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  deletedAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** display_name field predicates */
  displayName?: InputMaybe<Scalars['String']['input']>;
  displayNameContains?: InputMaybe<Scalars['String']['input']>;
  displayNameContainsFold?: InputMaybe<Scalars['String']['input']>;
  displayNameEqualFold?: InputMaybe<Scalars['String']['input']>;
  displayNameGT?: InputMaybe<Scalars['String']['input']>;
  displayNameGTE?: InputMaybe<Scalars['String']['input']>;
  displayNameHasPrefix?: InputMaybe<Scalars['String']['input']>;
  displayNameHasSuffix?: InputMaybe<Scalars['String']['input']>;
  displayNameIn?: InputMaybe<Array<Scalars['String']['input']>>;
  displayNameLT?: InputMaybe<Scalars['String']['input']>;
  displayNameLTE?: InputMaybe<Scalars['String']['input']>;
  displayNameNEQ?: InputMaybe<Scalars['String']['input']>;
  displayNameNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  /** first_name field predicates */
  firstName?: InputMaybe<Scalars['String']['input']>;
  firstNameContains?: InputMaybe<Scalars['String']['input']>;
  firstNameContainsFold?: InputMaybe<Scalars['String']['input']>;
  firstNameEqualFold?: InputMaybe<Scalars['String']['input']>;
  firstNameGT?: InputMaybe<Scalars['String']['input']>;
  firstNameGTE?: InputMaybe<Scalars['String']['input']>;
  firstNameHasPrefix?: InputMaybe<Scalars['String']['input']>;
  firstNameHasSuffix?: InputMaybe<Scalars['String']['input']>;
  firstNameIn?: InputMaybe<Array<Scalars['String']['input']>>;
  firstNameIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  firstNameLT?: InputMaybe<Scalars['String']['input']>;
  firstNameLTE?: InputMaybe<Scalars['String']['input']>;
  firstNameNEQ?: InputMaybe<Scalars['String']['input']>;
  firstNameNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  firstNameNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** gender field predicates */
  gender?: InputMaybe<UserGender>;
  genderIn?: InputMaybe<Array<UserGender>>;
  genderNEQ?: InputMaybe<UserGender>;
  genderNotIn?: InputMaybe<Array<UserGender>>;
  /** addresses edge predicates */
  hasAddresses?: InputMaybe<Scalars['Boolean']['input']>;
  hasAddressesWith?: InputMaybe<Array<UserAddrWhereInput>>;
  /** citizenship edge predicates */
  hasCitizenship?: InputMaybe<Scalars['Boolean']['input']>;
  hasCitizenshipWith?: InputMaybe<Array<CountryWhereInput>>;
  /** devices edge predicates */
  hasDevices?: InputMaybe<Scalars['Boolean']['input']>;
  hasDevicesWith?: InputMaybe<Array<UserDeviceWhereInput>>;
  /** identities edge predicates */
  hasIdentities?: InputMaybe<Scalars['Boolean']['input']>;
  hasIdentitiesWith?: InputMaybe<Array<UserIdentityWhereInput>>;
  /** login_profile edge predicates */
  hasLoginProfile?: InputMaybe<Scalars['Boolean']['input']>;
  hasLoginProfileWith?: InputMaybe<Array<UserLoginProfileWhereInput>>;
  /** oauth_clients edge predicates */
  hasOauthClients?: InputMaybe<Scalars['Boolean']['input']>;
  hasOauthClientsWith?: InputMaybe<Array<OauthClientWhereInput>>;
  /** passwords edge predicates */
  hasPasswords?: InputMaybe<Scalars['Boolean']['input']>;
  hasPasswordsWith?: InputMaybe<Array<UserPasswordWhereInput>>;
  /** permissions edge predicates */
  hasPermissions?: InputMaybe<Scalars['Boolean']['input']>;
  hasPermissionsWith?: InputMaybe<Array<PermissionWhereInput>>;
  /** user_quota edge predicates */
  hasUserQuota?: InputMaybe<Scalars['Boolean']['input']>;
  hasUserQuotaWith?: InputMaybe<Array<QuotaWhereInput>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** lang field predicates */
  lang?: InputMaybe<Scalars['String']['input']>;
  langContains?: InputMaybe<Scalars['String']['input']>;
  langContainsFold?: InputMaybe<Scalars['String']['input']>;
  langEqualFold?: InputMaybe<Scalars['String']['input']>;
  langGT?: InputMaybe<Scalars['String']['input']>;
  langGTE?: InputMaybe<Scalars['String']['input']>;
  langHasPrefix?: InputMaybe<Scalars['String']['input']>;
  langHasSuffix?: InputMaybe<Scalars['String']['input']>;
  langIn?: InputMaybe<Array<Scalars['String']['input']>>;
  langIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  langLT?: InputMaybe<Scalars['String']['input']>;
  langLTE?: InputMaybe<Scalars['String']['input']>;
  langNEQ?: InputMaybe<Scalars['String']['input']>;
  langNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  langNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** last_name field predicates */
  lastName?: InputMaybe<Scalars['String']['input']>;
  lastNameContains?: InputMaybe<Scalars['String']['input']>;
  lastNameContainsFold?: InputMaybe<Scalars['String']['input']>;
  lastNameEqualFold?: InputMaybe<Scalars['String']['input']>;
  lastNameGT?: InputMaybe<Scalars['String']['input']>;
  lastNameGTE?: InputMaybe<Scalars['String']['input']>;
  lastNameHasPrefix?: InputMaybe<Scalars['String']['input']>;
  lastNameHasSuffix?: InputMaybe<Scalars['String']['input']>;
  lastNameIn?: InputMaybe<Array<Scalars['String']['input']>>;
  lastNameIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  lastNameLT?: InputMaybe<Scalars['String']['input']>;
  lastNameLTE?: InputMaybe<Scalars['String']['input']>;
  lastNameNEQ?: InputMaybe<Scalars['String']['input']>;
  lastNameNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  lastNameNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** middle_name field predicates */
  middleName?: InputMaybe<Scalars['String']['input']>;
  middleNameContains?: InputMaybe<Scalars['String']['input']>;
  middleNameContainsFold?: InputMaybe<Scalars['String']['input']>;
  middleNameEqualFold?: InputMaybe<Scalars['String']['input']>;
  middleNameGT?: InputMaybe<Scalars['String']['input']>;
  middleNameGTE?: InputMaybe<Scalars['String']['input']>;
  middleNameHasPrefix?: InputMaybe<Scalars['String']['input']>;
  middleNameHasSuffix?: InputMaybe<Scalars['String']['input']>;
  middleNameIn?: InputMaybe<Array<Scalars['String']['input']>>;
  middleNameIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  middleNameLT?: InputMaybe<Scalars['String']['input']>;
  middleNameLTE?: InputMaybe<Scalars['String']['input']>;
  middleNameNEQ?: InputMaybe<Scalars['String']['input']>;
  middleNameNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  middleNameNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  not?: InputMaybe<UserWhereInput>;
  or?: InputMaybe<Array<UserWhereInput>>;
  /** principal_name field predicates */
  principalName?: InputMaybe<Scalars['String']['input']>;
  principalNameContains?: InputMaybe<Scalars['String']['input']>;
  principalNameContainsFold?: InputMaybe<Scalars['String']['input']>;
  principalNameEqualFold?: InputMaybe<Scalars['String']['input']>;
  principalNameGT?: InputMaybe<Scalars['String']['input']>;
  principalNameGTE?: InputMaybe<Scalars['String']['input']>;
  principalNameHasPrefix?: InputMaybe<Scalars['String']['input']>;
  principalNameHasSuffix?: InputMaybe<Scalars['String']['input']>;
  principalNameIn?: InputMaybe<Array<Scalars['String']['input']>>;
  principalNameLT?: InputMaybe<Scalars['String']['input']>;
  principalNameLTE?: InputMaybe<Scalars['String']['input']>;
  principalNameNEQ?: InputMaybe<Scalars['String']['input']>;
  principalNameNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  /** register_ip field predicates */
  registerIP?: InputMaybe<Scalars['String']['input']>;
  registerIPContains?: InputMaybe<Scalars['String']['input']>;
  registerIPContainsFold?: InputMaybe<Scalars['String']['input']>;
  registerIPEqualFold?: InputMaybe<Scalars['String']['input']>;
  registerIPGT?: InputMaybe<Scalars['String']['input']>;
  registerIPGTE?: InputMaybe<Scalars['String']['input']>;
  registerIPHasPrefix?: InputMaybe<Scalars['String']['input']>;
  registerIPHasSuffix?: InputMaybe<Scalars['String']['input']>;
  registerIPIn?: InputMaybe<Array<Scalars['String']['input']>>;
  registerIPLT?: InputMaybe<Scalars['String']['input']>;
  registerIPLTE?: InputMaybe<Scalars['String']['input']>;
  registerIPNEQ?: InputMaybe<Scalars['String']['input']>;
  registerIPNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  /** status field predicates */
  status?: InputMaybe<UserSimpleStatus>;
  statusIn?: InputMaybe<Array<UserSimpleStatus>>;
  statusIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  statusNEQ?: InputMaybe<UserSimpleStatus>;
  statusNotIn?: InputMaybe<Array<UserSimpleStatus>>;
  statusNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_at field predicates */
  updatedAt?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtGTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedAtLT?: InputMaybe<Scalars['Time']['input']>;
  updatedAtLTE?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNEQ?: InputMaybe<Scalars['Time']['input']>;
  updatedAtNotIn?: InputMaybe<Array<Scalars['Time']['input']>>;
  updatedAtNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** updated_by field predicates */
  updatedBy?: InputMaybe<Scalars['Int']['input']>;
  updatedByGT?: InputMaybe<Scalars['Int']['input']>;
  updatedByGTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  updatedByLT?: InputMaybe<Scalars['Int']['input']>;
  updatedByLTE?: InputMaybe<Scalars['Int']['input']>;
  updatedByNEQ?: InputMaybe<Scalars['Int']['input']>;
  updatedByNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  updatedByNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  /** user_type field predicates */
  userType?: InputMaybe<UserUserType>;
  userTypeIn?: InputMaybe<Array<UserUserType>>;
  userTypeNEQ?: InputMaybe<UserUserType>;
  userTypeNotIn?: InputMaybe<Array<UserUserType>>;
};

export type OrgUserListQueryVariables = Exact<{
  gid: Scalars['GID']['input'];
  first?: InputMaybe<Scalars['Int']['input']>;
  orderBy?: InputMaybe<UserOrder>;
  where?: InputMaybe<UserWhereInput>;
}>;


export type OrgUserListQuery = { __typename?: 'Query', node?: { __typename?: 'App' } | { __typename?: 'AppAction' } | { __typename?: 'AppDict' } | { __typename?: 'AppDictItem' } | { __typename?: 'AppMenu' } | { __typename?: 'AppPolicy' } | { __typename?: 'AppPolicyView' } | { __typename?: 'AppRes' } | { __typename?: 'AppRole' } | { __typename?: 'Country' } | { __typename?: 'Currency' } | { __typename?: 'FileIdentity' } | { __typename?: 'FileIdentityForApp' } | { __typename?: 'FileSource' } | { __typename?: 'OauthClient' } | { __typename?: 'Org', id: string, users: { __typename?: 'UserConnection', totalCount: number, pageInfo: { __typename?: 'PageInfo', hasNextPage: boolean, hasPreviousPage: boolean, startCursor?: any | null, endCursor?: any | null }, edges?: Array<{ __typename?: 'UserEdge', cursor: any, node?: { __typename?: 'User', id: string, displayName: string, addresses?: Array<{ __typename?: 'UserAddr', email?: string | null }> | null } | null } | null> | null } } | { __typename?: 'OrgPolicy' } | { __typename?: 'OrgRole' } | { __typename?: 'OrgUserPreference' } | { __typename?: 'Permission' } | { __typename?: 'Quota' } | { __typename?: 'QuotaItem' } | { __typename?: 'Region' } | { __typename?: 'User' } | { __typename?: 'UserAddr' } | { __typename?: 'UserDevice' } | { __typename?: 'UserIdentity' } | { __typename?: 'UserLoginProfile' } | { __typename?: 'UserPassword' } | { __typename?: 'UserPasswordPolicy' } | null };


export const OrgUserListDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"orgUserList"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"gid"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"GID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"first"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"Int"}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"orderBy"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"UserOrder"}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"where"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"UserWhereInput"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"node"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"gid"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"InlineFragment","typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"Org"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"users"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"first"},"value":{"kind":"Variable","name":{"kind":"Name","value":"first"}}},{"kind":"Argument","name":{"kind":"Name","value":"orderBy"},"value":{"kind":"Variable","name":{"kind":"Name","value":"orderBy"}}},{"kind":"Argument","name":{"kind":"Name","value":"where"},"value":{"kind":"Variable","name":{"kind":"Name","value":"where"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"totalCount"}},{"kind":"Field","name":{"kind":"Name","value":"pageInfo"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"hasNextPage"}},{"kind":"Field","name":{"kind":"Name","value":"hasPreviousPage"}},{"kind":"Field","name":{"kind":"Name","value":"startCursor"}},{"kind":"Field","name":{"kind":"Name","value":"endCursor"}}]}},{"kind":"Field","name":{"kind":"Name","value":"edges"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"cursor"}},{"kind":"Field","name":{"kind":"Name","value":"node"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"displayName"}},{"kind":"Field","name":{"kind":"Name","value":"addresses"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"email"}}]}}]}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<OrgUserListQuery, OrgUserListQueryVariables>;