# OPERATIONS

## 设置PaaS产品硬盘QoS - UpdatePaaSDiskQoS

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string |  | [] | [] | **Yes** |
| **CompanyID** | integer |  | [] | [] | No |
| **ResourceID** | string |  | [] | [] | **Yes** |
| **DiskIOPS** | integer |  | [] | [] | No |
| **DiskBandwidth** | integer |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 开关数据库审计 - UpdateAuditLog

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | No |
| **ResourceID** | string | 资源ID | [] | [] | No |
| **ResourceType** | string | 资源类型 | [] | [] | No |
| **AuditStatus** | integer | 开关, 0:关闭  1:开启 | [] | [] | No |
| **AuditEvents** | string | 审计日志事件, QUERY, QUERY_DDL, QUERY_DML, QUERY_DCL, 多个事件用逗号分隔 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取审计日志 - DescribeAuditLog

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | No |
| **ResourceID** | string | 资源ID | [] | [] | No |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **BeginTime** | integer | 开始时间 | [] | [] | **Yes** |
| **EndTime** | integer | 结束时间 | [] | [] | **Yes** |
| **IgnoreQuery** | string | 忽略审计语句关键字, select,show, 多个关键字用逗号分隔 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数 | [] |
| **Infos** | array<[*AuditLogInfos*](#AuditLogInfos)> |  | [] |
| **AuditEvents** | string | 审计日志事件, QUERY, QUERY_DDL, QUERY_DML, QUERY_DCL, 多个事件用逗号分隔 | [] |

### 数据模型

#### AuditLogInfos

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **LogType** | string | 日志类型 | [] |
| **ClientIP** | string | 客户端IP | [] |
| **AccountName** | string | 账户名 | [] |
| **SQLCommand** | string | 执行命令 | [] |
| **ExecutionDB** | string | 操作数据库 | [] |
| **ExecutionTime** | integer | 执行时间 | [] |
| **ExecutionResult** | integer | 执行结果 | [] |

## 修改 PaaS产品 删除保护 - UpdateTerminationPolicy

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ResourceID** | string | 资源ID | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **TerminationPolicy** | integer | 是否开启删除保护 | [] | [] | No |
| **SlaveIDs** | array<string> | 从库ID | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取连接信息 - GetConnectionInfo

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | No |
| **ResourceID** | string | 资源ID | [] | [] | No |
| **ProductType** | string | 资源类型 | [] | [] | No |
| **BeginTime** | integer | 开始时间 | [] | [] | No |
| **EndTime** | integer | 结束时间 | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **CompanyID** | integer | 租户ID | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数 | [] |
| **Infos** | array<[*ConnectionInfos*](#ConnectionInfos)> |  | [] |

### 数据模型

#### ConnectionInfos

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ResponseAdd** | string | 响应地址 | [] |
| **RequestAdd** | string | 请求地址 | [] |
| **RequestResourceID** | string | 请求地址机器ID | [] |
| **RequestResourceName** | string | 请求地址机器名称 | [] |
| **ConnectionStatus** | string | 连接状态 | [] |
| **SearchTime** | integer | 查询时间 | [] |
| **RequestResourceCompanyID** | integer | 请求地址机器租户ID | [] |

## 查询参数修改记录 - DescribeParametersHistories

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **BeginTime** | integer | 开始时间 | [] | [] | **Yes** |
| **EndTime** | integer | 结束时间 | [] | [] | **Yes** |
| **DatabaseID** | string | 数据库 实例 ID | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **ProductType** | string | 产品类型 | [] | [] | **Yes** |
| **Keyword** | string | 关键词 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数 | [] |
| **Infos** | array<[*ParametersHistoriesInfos*](#ParametersHistoriesInfos)> |  | [] |

### 数据模型

#### ParametersHistoriesInfos

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **NewValue** | string | 新值 | [] |
| **OldValue** | string | 旧值 | [] |
| **ParameterName** | string | 参数名称 | [] |
| **Updatetime** | integer | 更新时间 | [] |

## Paas 实例关机 - StopPaaSInstance

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | No |
| **ResourceID** | string | 资源ID | [] | [] | **Yes** |
| **ProductType** | string | 资源类型 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## Paas 实例开机 - StartPaaSInstance

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | No |
| **ResourceID** | string | 资源ID | [] | [] | **Yes** |
| **ProductType** | string | 资源类型 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|

