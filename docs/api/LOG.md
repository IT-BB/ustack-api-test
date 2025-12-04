# LOG

## 获取操作日志 - DescribeOPLogs

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **ResourceID** | string | 资源ID | [] | [] | No |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **ProductType** | string | 产品类型 | [] | [] | No |
| **IsSuccess** | string | 是否成功 | [] | [] | No |
| **BeginTime** | integer | 开始时间 | [] | [] | **Yes** |
| **EndTime** | integer | 结束时间 | [] | [] | **Yes** |
| **Keyword** | string | 关键词 | [] | [] | No |
| **Limit** | integer | 分页大小 | [] | [] | **Yes** |
| **Offset** | integer | 分页偏移量 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*OPLogInfo*](#OPLogInfo)> | 详情 | [] |
| **TotalCount** | integer | 总数 | [] |

### 数据模型

#### HttpInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **RequestMethods** | string | 请求方法 | [] |
| **RequestProtocolType** | string | 请求协议 | [] |
| **RequestHeader** | string | 请求头 | [] |
| **RequestBody** | string | 请求体 | [] |
| **RequestURI** | string | 请求URI | [] |
| **RequestParams** | string | 请求参数 | [] |
| **ResponseHeader** | string | 响应头 | [] |
| **ResponseBody** | string | 响应体 | [] |
| **ResponseStatusCode** | integer | 状态码 | [] |

#### OPLogInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **OPLogID** | integer | 操作日志ID | [] |
| **CompanyID** | integer | 租户ID | [] |
| **Region** | string | 地域 | [] |
| **APIName** | string | API名称 | [] |
| **UserEmail** | string | 账号邮箱 | [] |
| **ProductType** | string | 产品类型 | [] |
| **ResourceID** | string | 资源ID | [] |
| **RetCode** | integer | 返回码 | [] |
| **Message** | string | 返回信息 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |
| **ClientIp** | string | 客户端IP | [] |
| **RegionAlias** | string | 地域别称 | [] |
| **HttpInfo** | object | 请求响应信息 | [] |
