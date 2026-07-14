# 动态附件

邮件通知支持在发送时携带动态附件。附件来源可以是本地文件路径，也可以是 HTTP(S) URL。

## 使用方式

通过 Alert 的 `annotations` 传递附件路径，使用保留键 `__attachments__`，多个附件以分号 `;` 分隔。

```json
POST /api/v2/alerts
[
  {
    "labels": {
      "tenant": "1",
      "alertname": "MonthlyReport",
      "user": "100"
    },
    "annotations": {
      "summary": "月度报表",
      "__attachments__": "https://oss.example.com/reports/2026-07.pdf;/data/exports/summary.csv"
    }
  }
]
```

使用分号而非逗号作为分隔符，是为了避免与 URL 查询参数或文件路径中的逗号产生冲突。

## 附件路径类型

| 类型 | 示例 | 说明 |
|------|------|------|
| HTTP(S) URL | `https://oss.example.com/report.pdf` | 发送时实时下载，附加到邮件 |
| 本地文件路径 | `/data/attachments/report.pdf` | 直接读取本地文件附加 |

两种类型可在同一次请求中混合使用，以分号分隔即可。

## 与静态附件的关系

模板中可配置静态附件（在模板管理界面上传），静态附件在模板启用时下载到本地缓存。

发送时的附件合并顺序：

1. **静态附件** — 来自模板配置，通过 `Headers["Attachments"]` 传递
2. **动态附件** — 来自 Alert annotations 中的 `__attachments__`

两者互不覆盖，按顺序依次附加到邮件。

## HTTP 附件处理机制

对于 HTTP(S) URL 类型的附件：

- **超时控制**：单个附件下载超时 30 秒
- **大小限制**：单个附件最大 50MB，超出部分截断
- **文件名提取**：优先从响应头 `Content-Disposition` 的 `filename` 参数获取；若无，则取 URL 路径的最后一段
- **Content-Type**：从响应头 `Content-Type` 获取；若未提供，默认 `application/octet-stream`
- **错误处理**：下载失败（网络错误、非 200 状态码等）会导致本次通知发送失败，触发重试

## 多 Alert 合并

当多条 Alert 被分组到同一通知时，各 Alert 的 `__attachments__` 会被合并并去重。

```json
[
  {
    "labels": {"tenant": "1", "alertname": "Report"},
    "annotations": {"__attachments__": "https://example.com/a.pdf"}
  },
  {
    "labels": {"tenant": "1", "alertname": "Report"},
    "annotations": {"__attachments__": "https://example.com/a.pdf;https://example.com/b.csv"}
  }
]
```

最终邮件携带 `a.pdf` 和 `b.csv` 两个附件（`a.pdf` 去重后只附加一次）。

## 注意事项

- `__attachments__` 的值以分号 `;` 分隔，分号前后的空格会自动 trim
- HTTP 附件在每次通知发送时都会重新下载，不会缓存
- 若 URL 需要认证，需确保 URL 本身包含必要的认证参数（如 token），或该地址为内网可直连地址
- 空值或纯空白的路径条目会被自动跳过，不会导致错误
