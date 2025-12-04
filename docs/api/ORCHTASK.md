# ORCHTASK

## 查询支持的编排任务 - DescribeOrchTaskType

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer |  | [] |
| **Infos** | array<[*TaskTypeInfo*](#TaskTypeInfo)> |  | [] |

### 数据模型

#### TaskTypeInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TaskName** | string | 任务名称 | [] |
| **TaskType** | string | 任务类型 | [] |
| **ResourceTypes** | array<string> | 支持的资源类型 | [] |

## 创建编排任务 - CreateOrchTask

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **Name** | string | 任务名称 | [] | [] | **Yes** |
| **Remark** | string | 任务备注 | [] | [] | No |
| **ProjectID** | string | 项目组ID | [] | [] | No |
| **TaskType** | string | 任务类型 | [] | [] | **Yes** |
| **ResourceType** | string | 资源类型 | [] | [] | **Yes** |
| **Condition** | string | 资源筛选条件 | [] | [] | No |
| **Steps** | array<string> | 步骤 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TaskID** | string | 任务ID | [] |


## 删除编排任务 - DeleteOrchTask

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string |  | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **TaskID** | string | 任务ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 查询编排任务 - DescribeOrchTask

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | No |
| **Keyword** | string | 搜索关键词 | [] | [] | No |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **TaskIDs** | array<string> | 任务IDs | [] | [] | No |
| **Status** | string | 任务状态 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer |  | [] |
| **Infos** | array<[*OrchTaskInfo*](#OrchTaskInfo)> |  | [] |

### 数据模型

#### OrchTaskInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TaskID** | string | 任务ID | [] |
| **Name** | string | 任务名称 | [] |
| **Remark** | string | 任务备注 | [] |
| **Region** | string | 地域 | [] |
| **RegionAlias** | string | 地域别名 | [] |
| **CompanyID** | integer | 租户ID | [] |
| **CompanyName** | string | 租户名称 | [] |
| **Email** | string | 租户邮箱 | [] |
| **ProjectID** | string | 项目组ID | [] |
| **ProjectName** | string | 项目组名称 | [] |
| **Reason** | string | 失败原因 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |
| **Status** | string | 任务状态 | [] |
| **TaskType** | string | 任务类型 | [] |
| **ResourceType** | string | 资源类型 | [] |
| **Condition** | string | 资源筛选条件 | [] |
| **Execution** | object | 最新执行 | [] |
| **Steps** | array<[*TaskStep*](#TaskStep)> | 步骤信息 | [] |

#### TaskExecution

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Result** | string | 最新执行结果 | [] |
| **Message** | string | 最新执行信息 | [] |
| **StartsAt** | integer | 最新执行开始时间 | [] |
| **EndsAt** | integer | 最新执行结束时间 | [] |

#### TaskStep

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Order** | integer | 步骤次序 | [] |
| **Resources** | array<[*TaskStepResource*](#TaskStepResource)> | 步骤包含的资源信息 | [] |
| **Delay** | integer | 步骤延迟 | [] |

#### TaskStepResource

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ResourceID** | string | 资源ID | [] |
| **Name** | string | 资源名称 | [] |
| **Remark** | string | 资源备注 | [] |
| **Status** | string | 资源状态 | [] |
| **Execution** | object | 最新执行 | [] |

## 更新编排任务 - UpdateOrchTask

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **TaskID** | string | 任务ID | [] | [] | **Yes** |
| **Condition** | string | 资源筛选条件 | [] | [] | No |
| **Steps** | array<string> | 步骤 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 操作编排任务 - OperateOrchTask

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string |  | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **TaskID** | string | 任务ID | [] | [] | **Yes** |
| **Operation** | string | 任务操作 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|

