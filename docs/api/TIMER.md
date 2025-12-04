# TIMER

## 创建定时器 - CreateTimer

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ProjectID** | string | 项目组 | [] | [] | No |
| **TagKeyValuePairs** | array<string> | 标签键值对 | [] | [] | **Yes** |
| **Name** | string | 名称 | [] | [] | **Yes** |
| **Remark** | string | 描述 | [] | [] | No |
| **Task** | string |  | [] | [] | **Yes** |
| **Type** | string |  | [] | [] | **Yes** |
| **Days** | array<int32> | 日期 | [] | [] | No |
| **Weeks** | array<int32> | 星期 | [] | [] | No |
| **Hours** | array<int32> | 小时 | [] | [] | No |
| **FixedInterval** | integer | 间隔 | [] | [] | No |
| **TriggerTime** | integer | 单次触发时间戳 | [] | [] | No |
| **ResourceIDs** | array<string> | 资源IDs | [] | [] | No |
| **ResourceType** | string | 资源类型 | [] | [] | No |
| **KeepNum** | integer | 保留数量 | [] | [] | No |
| **Retention** | integer | 保留时长 | [] | [] | No |
| **Annotations** | array<string> | 特定参数 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TimerID** | string | 定时器ID | [] |


## 查询定时器 - DescribeTimer

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **Keyword** | string | 关键词 | [] | [] | No |
| **ProjectIDs** | array<string> | 项目组IDs | [] | [] | No |
| **TimerIDs** | array<string> | 定时器IDs | [] | [] | No |
| **Types** | array<string> | 类型 | [] | [] | No |
| **TimerStatus** | array<string> | 状态 | [] | [] | No |
| **Task** | string | 任务类型 | [] | [] | No |
| **SetID** | string | 集群ID | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数 | [] |
| **Infos** | array<[*OldTimerInfo*](#OldTimerInfo)> |  | [] |

### 数据模型

#### OldTimerInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Region** | string | 地域 | [] |
| **CompanyID** | integer | 租户ID | [] |
| **Email** | string | 邮箱 | [] |
| **Name** | string | 名称 | [] |
| **Remark** | string | 描述 | [] |
| **ProjectID** | string | 项目组ID | [] |
| **ProjectName** | string | 项目组名称 | [] |
| **TimerID** | string | 定时器ID | [] |
| **Task** | string | 任务类型 | [] |
| **Type** | string | 定时器类型 | [] |
| **TimerStatus** | string | 状态 | [] |
| **Days** | array<int32> | 日期 | [] |
| **Weeks** | array<int32> | 星期 | [] |
| **Hours** | array<int32> | 小时 | [] |
| **FixedInterval** | integer | 间隔 | [] |
| **TriggerTime** | integer | 触发时间 | [] |
| **KeepNum** | integer | 保留数量 | [] |
| **Retention** | integer | 保留时长 | [] |
| **ResourceType** | string | 资源类型 | [] |
| **TimerResources** | array<[*TimerResources*](#TimerResources)> | 资源ID | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |
| **Annotations** | array<string> | 特定参数 | [] |
| **RegionAlias** | string | 地域别称 | [] |

#### TimerResources

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ResourceID** | string |  | [] |
| **ResourceName** | string |  | [] |

## 查询定时器执行记录 - DescribeTimerTask

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **TimerID** | string |  | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数 | [] |
| **Infos** | array<[*TaskInfo*](#TaskInfo)> |  | [] |

### 数据模型

#### TaskInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TaskID** | string | 任务执行记录ID | [] |
| **TimerID** | string | 定时器ID | [] |
| **ResourceID** | string | 资源ID | [] |
| **Task** | string | 任务类型 | [] |
| **TaskStatus** | string | 任务执行状态 | [] |
| **TargetID** | string | 任务执行结果ID | [] |
| **Remark** | string | 任务执行备注 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **BeginTime** | integer | 开始时间 | [] |
| **EndTime** | integer | 结束时间 | [] |
| **ErrCode** | integer | 错误码 | [] |

## 更新定时器 - UpdateTimer

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ProjectID** | string | 项目组ID | [] | [] | No |
| **TimerID** | string | 定时器ID | [] | [] | **Yes** |
| **Name** | string | 名称 | [] | [] | **Yes** |
| **Remark** | string | 备注 | [] | [] | No |
| **Task** | string | 任务 | [] | [] | **Yes** |
| **Type** | string | 类型 | [] | [] | **Yes** |
| **Days** | array<int32> | 日期 | [] | [] | No |
| **Weeks** | array<int32> | 星期 | [] | [] | No |
| **Hours** | array<int32> | 小时 | [] | [] | No |
| **FixedInterval** | integer | 间隔 | [] | [] | No |
| **TriggerTime** | integer | 触发时间 | [] | [] | No |
| **ResourceIDs** | array<string> | 资源IDs | [] | [] | No |
| **ResourceType** | string | 资源类型 | [] | [] | No |
| **KeepNum** | integer | 保留数量 | [] | [] | No |
| **Retention** | integer | 保留时长 | [] | [] | No |
| **Annotations** | array<string> | 特定参数 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 删除定时器 - DeleteTimer

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **TimerID** | string | 定时器ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|

