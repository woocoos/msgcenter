# 友盟推送通知

本文档描述 msgcenter 中友盟（Umeng）推送通知器的实现和使用。

## 概述

友盟推送通知器支持向 Android、iOS 和 HarmonyOS 设备发送推送通知。它基于友盟 Push API 实现，支持多种推送类型和平台特定的 payload 结构。

## 配置

### 基本配置

在 `alertmanager.yaml` 中配置友盟推送：

```yaml
route:
  receiver: 'umeng-push'

receivers:
  - name: 'umeng-push'
    umengConfigs:
      - apps:
          myapp-android:
            appKey: "your_android_app_key"
            appMasterSecret: "your_android_master_secret"
            platform: "android"
            appSet: "myapp"
            aliasType: "uid"
            afterOpen: "go_activity"
            activity: "com.example.MainActivity"
          myapp-ios:
            appKey: "your_ios_app_key"
            appMasterSecret: "your_ios_master_secret"
            platform: "ios"
            appSet: "myapp"
            aliasType: "uid"
          myapp-harmonyos:
            appKey: "your_harmonyos_app_key"
            appMasterSecret: "your_harmonyos_master_secret"
            platform: "harmonyos"
            appSet: "myapp"
            aliasType: "uid"
```

### 配置字段说明

#### UmengConfig

| 字段 | 类型 | 说明 |
|------|------|------|
| `apps` | `map[string]*UmengAppConfig` | 应用配置映射，key 为应用名称 |
| `apiURL` | `string` | API 端点，默认 `https://msgapi.umeng.com/api/send` |
| `productionMode` | `*bool` | 是否使用生产模式 |

#### UmengAppConfig

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `appKey` | `string` | 是 | 友盟应用 Key |
| `appMasterSecret` | `string` | 是 | 友盟应用 Master Secret |
| `platform` | `string` | 是 | 平台标识：`android`、`ios`、`harmonyos` |
| `appSet` | `string` | 否 | 应用集名称，用于按业务应用分组 |
| `aliasType` | `string` | 否 | 别名类型，默认 `uid` |
| `afterOpen` | `string` | 否 | 点击通知后的行为（仅 Android） |
| `activity` | `string` | 否 | 要打开的 Activity（仅 Android） |

## 多平台支持

### 用户设备与推送的关系

推送通知的目标是**用户**，但实际发送时需要知道用户的**设备平台**。系统通过以下流程建立用户与设备的关联：

```
用户 (user_id)
  └─> 设备 (user_device 表)
        ├─> device_uid: 设备唯一标识
        ├─> system_name: 操作系统名称（如 "Android 13"、"iOS 16"）
        ├─> device_model: 设备型号
        └─> status: 设备状态（active/inactive）
```

**关键概念**：
- 一个用户可以有多个设备（如同时拥有 Android 手机和 iPad）
- 每个设备属于一个平台（Android/iOS/HarmonyOS）
- 推送时，系统查询用户的所有活跃设备，按平台分组后分别发送

### 推送路由流程

```
1. Alert 推送 → 提取 user 标签（如 user: "100,200"）
                    ↓
2. 查询 user_device 表 → 获取用户 100 和 200 的设备信息
                    ↓
3. 按平台分组：
   - 用户 100: Android 设备 1 台
   - 用户 200: iOS 设备 1 台，HarmonyOS 设备 1 台
                    ↓
4. 匹配配置中的应用：
   - Android 设备 → myapp-android 应用
   - iOS 设备 → myapp-ios 应用
   - HarmonyOS 设备 → myapp-harmonyos 应用
                    ↓
5. 分别向三个应用发送推送
```

### 设备状态要求

只有 `status = active` 的设备才会被纳入推送目标。如果用户没有活跃设备，系统会记录警告日志并跳过该用户。

### 平台识别

系统通过查询 `user_device` 表获取用户的设备信息，根据 `system_name` 字段识别设备平台：

- 包含 "android" → Android
- 包含 "ios"、"iphone"、"ipad" → iOS
- 包含 "harmony"、"鸿蒙" → HarmonyOS

### 平台特定 Payload

#### Android

```json
{
  "payload": {
    "display_type": "notification",
    "body": {
      "title": "消息通知",
      "ticker": "消息通知",
      "text": "您有一条新消息",
      "after_open": "go_activity",
      "activity": "com.example.MainActivity"
    },
    "extra": {
      "type": "msg_detail",
      "id": "123"
    }
  },
  "category": 1
}
```

#### iOS

```json
{
  "payload": {
    "display_type": "notification",
    "aps": {
      "alert": {
        "title": "消息通知",
        "body": "您有一条新消息"
      },
      "sound": "default"
    },
    "type": "msg_detail",
    "id": "123"
  }
}
```

**注意**：iOS 没有 `extra` 字段，自定义内容直接放在 payload 根级别。

#### HarmonyOS

```json
{
  "payload": {
    "display_type": "notification",
    "body": {
      "title": "消息通知",
      "body": "您有一条新消息"
    },
    "extra": {
      "type": "msg_detail",
      "id": "123"
    }
  },
  "channel_properties": {
    "harmony_channel_category": "MARKETING"
  },
  "policy": {
    "channel_strategy": {
      "default": 2
    }
  }
}
```

## 推送类型

系统根据目标用户数量自动选择推送类型：

| 用户数量 | 推送类型 | 说明 |
|----------|----------|------|
| 0 | `broadcast` | 广播，发送给所有设备 |
| 1 | `customizedcast` | 单播，通过 alias 发送给单个用户 |
| >1 | `customizedcast` | 列播，通过 alias 发送给多个用户 |

## 应用集路由

通过 alert 的 `appSet` 标签可以指定目标应用集：

```json
{
  "labels": {
    "alertname": "test",
    "user": "100",
    "appSet": "myapp"
  }
}
```

系统会根据 `appSet` 过滤配置中的应用，只向匹配的应用发送推送。

## 签名计算

友盟 API 要求通过 URL 参数传递签名：

```
sign = MD5(method + url + post-body + app_master_secret)
```

其中：
- `method`：HTTP 方法（POST，全大写）
- `url`：完整 URL（包括协议、主机、路径，不包括 query string）
- `post-body`：请求体的 JSON 字符串
- `app_master_secret`：应用的 Master Secret

示例：
```
POST https://msgapi.umeng.com/api/send?sign=abc123...
```

## Alert ID 传递

系统在保存 alert 到数据库后，会将数据库 ID 写入 alert 的 annotation：

```
__alert_id__: 123
```

在构建 payload 时，这个 ID 会被提取并添加到 `extra.id` 字段中，供客户端使用。

## Extra 字段

Payload 中的 `extra` 字段用于传递自定义信息：

| 字段 | 来源 | 说明 |
|------|------|------|
| `type` | annotation `msg_detail` | 消息类型，用于 App 端识别业务实例 |
| `id` | annotation `__alert_id__` | Alert 数据库 ID |

**注意**：iOS 平台没有 `extra` 字段，这些内容直接放在 payload 根级别。

## 错误处理

### HTTP 状态码

- `200 OK`：请求成功
- `4xx`：客户端错误，不重试
- `5xx`：服务器错误，重试

### API 响应

```json
{
  "ret": "SUCCESS" | "FAIL",
  "error_code": "2007",
  "error_msg": "签名不正确"
}
```

常见错误码：
- `2007`：签名不正确
- `1001`：无效的 app_key

## 日志

系统使用 `zap` 记录日志，组件名为 `umeng`：

- `Warn`：alias_type 为空、用户无设备、应用未配置、推送失败
- `Error`：数据库查询失败

## 集成测试

运行集成测试需要配置 `notify/umeng/testdata/.env.local`：

```bash
UMENG_ANDROID_APP_KEY=your_key
UMENG_ANDROID_APP_MASTER_SECRET=your_secret
UMENG_ANDROID_TEST_USER=101164

UMENG_IOS_APP_KEY=your_key
UMENG_IOS_APP_MASTER_SECRET=your_secret
UMENG_IOS_TEST_USER=101164
```

运行测试：

```bash
go test -v -run TestIntegration_RealAPI_Android ./notify/umeng/...
go test -v -run TestIntegration_RealAPI_IOS ./notify/umeng/...
```
