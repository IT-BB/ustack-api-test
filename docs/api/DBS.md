# DBS

## 绑定存储系统到DBS - BindStorageToDBS

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Name** | string | 名称 | [] | [] | **Yes** |
| **Remark** | string | 备注 | [] | [] | No |
| **StorageType** | string | 存储集群 | [] | [] | **Yes** |
| **OSSID** | string | 对象存储ID | [] | [] | No |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **Endpoint** | string | 外部服务地址 | [] | [] | No |
| **AccessKey** | string | 外部服务密钥 | [] | [] | No |
| **SecretKey** | string | 外部服务密钥 | [] | [] | No |
| **BucketName** | string | 外部服务存储桶 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **StorageID** | string | 存储池ID | [] |


## 从DBS解绑存储系统 - UnbindStorageFromDBS

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **StorageID** | string | 存储池ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取DBS存储系统 - DescribeDBSStorage

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数 | [] |
| **Infos** | array<[*DBSStorageInfo*](#DBSStorageInfo)> | 存储池列表 | [] |

### 数据模型

#### DBSStorageInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **StorageID** | string | 存储池ID | [] |
| **OSSID** | string | 对象存储ID | [] |
| **OSSName** | string | 对象存储名称 | [] |
| **Name** | string | 名称 | [] |
| **Remark** | string | 备注 | [] |
| **CompanyID** | integer | 租户ID | [] |
| **Email** | string | Email | [] |
| **Status** | string | 状态 | [] |
| **StorageType** | string | 存储集群 | [] |
| **Space** | double | 总空间 | [] |
| **SpaceUsed** | double | 已用空间 | [] |
| **RelatedPlanIDs** | array<string> | 备份计划ID | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |
| **Region** | string | 地域 | [] |
| **Endpoint** | string | 外部服务地址 | [] |
| **AccessKey** | string | 外部服务密钥 | [] |
| **SecretKey** | string | 外部服务密钥 | [] |
| **BucketName** | string | 外部服务存储桶 | [] |

## 更新DBS存储系统 - UpdateDBSStorage

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **StorageID** | string | 存储池ID | [] | [] | **Yes** |
| **Name** | string | 名称 | [] | [] | **Yes** |
| **Remark** | string | 备注 | [] | [] | No |
| **CompanyID** | string | 租户ID | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取备份计划 - DescribeDBSBackupPlan

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **PlanID** | string | 备份计划ID | [] | [] | No |
| **SrcResourceID** | string | 备份源ID | [] | [] | No |
| **BackupID** | string | 备份ID | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数 | [] |
| **Infos** | array<[*DBSPlanInfo*](#DBSPlanInfo)> | 备份计划信息 | [] |

### 数据模型

#### DBSPlanInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **PlanID** | string | 备份计划ID | [] |
| **Name** | string | 名称 | [] |
| **Remark** | string | 备注 | [] |
| **CompanyID** | integer | 租户ID | [] |
| **CompanyName** | string | 租户名称 | [] |
| **CompanyEmail** | string | 租户Email | [] |
| **Email** | string | Email | [] |
| **Status** | string | 状态 | [] |
| **SrcName** | string | 备份源名称 | [] |
| **SrcType** | string | 备份源类型 | [] |
| **SrcRegion** | string | 备份源地域 | [] |
| **SrcID** | string | 备份源ID | [] |
| **ScheduleType** | string | 调度类型 | [] |
| **TimerType** | string | 定时器类型 | [] |
| **TimerID** | string | 定时器ID | [] |
| **TimerDays** | array<int32> | 执行日期 | [] |
| **TimerWeeks** | array<int32> | 执行星期 | [] |
| **TimerHours** | array<int32> | 执行小时 | [] |
| **BackupType** | string | 备份类型 | [] |
| **StorageName** | string | 存储池名称 | [] |
| **StorageID** | string | 存储池ID | [] |
| **RetentionTime** | integer | 保留时间 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |
| **BackupTables** | array<string> | 备份库表 | [] |
| **BackupCount** | integer | 备份数量 | [] |
| **Reason** | string | 失败原因 | [] |
| **IncrementalInterval** | integer | 增量备份周期 | [] |
| **IncrementalSrcResourceID** | string | 增量备份源ID | [] |
| **RelatedPlanID** | string | 关联的计划 | [] |
| **StorageSetArch** | string | 备份资源存储集群架构 | [] |
| **PauseReason** | string | 暂停原因 | [] |

## 创建备份计划 - CreateDBSBackupPlan

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **Name** | string | 名称 | [] | [] | **Yes** |
| **Remark** | string | 备注 | [] | [] | No |
| **SrcResourceType** | string | 备份源头类型 | [] | [] | **Yes** |
| **SrcResourceRegion** | string | 备份源地域 | [] | [] | **Yes** |
| **SrcResourceID** | string | 备份源ID | [] | [] | **Yes** |
| **ScheduleType** | string | 调度类型 | [] | [] | **Yes** |
| **TimerType** | string | 定时器类型，备份策略为计时器时必填 | [] | [] | No |
| **TimerDays** | array<int32> | 执行日期，与定时器类型对应的必填 | [] | [] | No |
| **TimerWeeks** | array<int32> | 执行星期，与定时器类型对应的必填 | [] | [] | No |
| **TimerHours** | array<int32> | 执行小时，与定时器类型对应的必填 | [] | [] | No |
| **BackupType** | string | 备份类型 | [] | [] | **Yes** |
| **StorageID** | string | 存储池ID，非快照备份时必填 | [] | [] | No |
| **RetentionTime** | integer | 保留时间 | [] | [] | **Yes** |
| **BackupTables** | string | 备份库表 | [] | [] | No |
| **IncrementalSrcResourceID** | string | 增量备份源ID | [] | [] | No |
| **IncrementalInterval** | integer | 增量备份间隔 单位分钟, 最低5分钟 | [] | [] | No |
| **RelatedPlanID** | string | 关联的计划 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **PlanID** | string | 备份计划ID | [] |


## 删除备份计划 - DeleteDBSBackupPlan

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **PlanID** | string | 计划ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 更新备份计划 - UpdateDBSBackupPlan

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **PlanID** | string | 计划ID | [] | [] | **Yes** |
| **Name** | string | 名称 | [] | [] | **Yes** |
| **Remark** | string | 备注 | [] | [] | No |
| **SrcResourceType** | string | 源数据类型 | [] | [] | **Yes** |
| **SrcResourceRegion** | string | 源数据地域 | [] | [] | **Yes** |
| **SrcResourceID** | string | 源资源ID | [] | [] | **Yes** |
| **ScheduleType** | string | 调度类型 | [] | [] | **Yes** |
| **TimerType** | string | 定时器类型，备份策略为计时器时必填 | [] | [] | No |
| **TimerDays** | array<int32> | 执行日期，与定时器类型对应的必填 | [] | [] | No |
| **TimerWeeks** | array<int32> | 执行星期，与定时器类型对应的必填 | [] | [] | No |
| **TimerHours** | array<int32> | 执行小时，与定时器类型对应的必填 | [] | [] | No |
| **BackupType** | string | 备份类型 | [] | [] | **Yes** |
| **StorageID** | string | 存储系统 | [] | [] | No |
| **RetentionTime** | integer | 保留时间 | [] | [] | **Yes** |
| **BackupTables** | string | 备份库表 | [] | [] | No |
| **IncrementalInterval** | integer | 增量备份间隔 单位分钟 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 更新备份计划的名称和remark - UpdateDBSBackupPlanSimple

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **PlanID** | string | 计划ID | [] | [] | **Yes** |
| **Name** | string | 名称 | [] | [] | **Yes** |
| **Remark** | string | 备注 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 手动执行备份计划 - ExecDBSBackupPlan

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **PlanID** | string | 计划ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **BackupID** | string | 备份ID | [] |


## 获取备份 - DescribeDBSBackup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **SrcResourceRegion** | string | 源数据地域 | [] | [] | No |
| **SrcResourceID** | string | 源资源ID | [] | [] | No |
| **PlanID** | string | 计划ID | [] | [] | No |
| **StorageID** | string | 存储池ID | [] | [] | No |
| **BackupIDs** | array<string> | 备份IDs | [] | [] | No |
| **BackupType** | string | 备份类型 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数 | [] |
| **Infos** | array<[*DBSBackupInfo*](#DBSBackupInfo)> | 备份列表 | [] |

### 数据模型

#### DBSBackupInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **BackupID** | string | 备份ID | [] |
| **PlanName** | string | 计划名称 | [] |
| **PlanID** | string | 计划ID | [] |
| **StorageName** | string | 存储池名称 | [] |
| **StorageID** | string | 存储池ID | [] |
| **CompanyName** | string | 租户名称 | [] |
| **CompanyEmail** | string | 租户Email | [] |
| **CompanyID** | integer | 租户ID | [] |
| **SrcName** | string | 备份源名称 | [] |
| **SrcType** | string | 备份源类型 | [] |
| **SrcRegion** | string | 备份源地域 | [] |
| **SrcID** | string | 备份源ID | [] |
| **Status** | string | 备份状态 | [] |
| **Reason** | string | 失败原因 | [] |
| **Size** | integer | 备份大小 | [] |
| **Path** | string | 备份存放路径 | [] |
| **RetentionTime** | integer | 保留时间 | [] |
| **ExpireTime** | integer | 过期时间 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |
| **SrcDiskSpace** | integer | 源磁盘空间 | [] |
| **SrcDiskSetType** | string | 源存储集群 | [] |
| **InstanceVersion** | string | 备份实例版本 | [] |
| **ClusterType** | string | 集群类型 | [] |
| **ShardCount** | integer | 分片数 | [] |

## 删除备份 - DeleteDBSBackup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **BackupID** | string | 备份ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 查看可恢复时间段详情 - DescribeDBSRestoreRangeInfo

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **BackupSourceID** | string | 备份源ID | [] | [] | **Yes** |
| **BackupType** | string | 备份类型, Physical,Logical,Snapshot 返回的是时间点，Incremental 返回的是时间段 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **BeginTimestamp** | integer | 最早可恢复时间 | [] |
| **EndTimestamp** | integer | 最晚可恢复时间 | [] |
| **SourceInstanceID** | string | 备份源ID | [] |
| **RestoreRangeInfos** | array<[*RestoreRangeInfo*](#RestoreRangeInfo)> | 恢复时间信息 | [] |

### 数据模型

#### RestoreRangeInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **BeginTimestamp** | integer | 开始时间 | [] |
| **EndTimestamp** | integer | 结束时间 | [] |
| **BackupID** | string | 备份ID | [] |

## 创建DBS网关 - CreateDBSGateway

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Name** | string | 名称 | [] | [] | **Yes** |
| **Remark** | string | 备注 | [] | [] | No |
| **EIPID** | string | 外网IPID | [] | [] | No |
| **ProjectID** | string | 项目组ID | [] | [] | No |
| **TagKeyValuePairs** | array<string> | 标签键值对 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **DBSGatewayID** | string | DBS网关ID | [] |


## 获取DBS网关 - DescribeDBSGateway

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Keyword** | string | 搜索关键词 | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **ProjectIDs** | array<string> | 项目组 ID | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数 | [] |
| **Infos** | array<[*DBSGateway*](#DBSGateway)> | DBS网关信息 | [] |

### 数据模型

#### DBSGateway

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **DBSGatewayID** | string | DBS网关ID | [] |
| **Status** | string | DBS网关状态 | [] |
| **CompanyID** | integer | 租户ID | [] |
| **CompanyName** | string | 租户名称 | [] |
| **Name** | string | 名称 | [] |
| **Remark** | string | 备注 | [] |
| **IP** | string | IP | [] |
| **Bandwidth** | integer | 带宽 | [] |
| **OperatorAlias** | string | 外网线路别名 | [] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> | 标签 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |
| **Region** | string | 地域 | [] |
| **ProjectID** | string | 项目组ID | [] |
| **ProjectName** | string | 项目名称 | [] |
| **Email** | string | 租户邮箱 | [] |
| **EIPID** | string | EIPID | [] |
| **EIPName** | string | EIPName | [] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

## 换绑备份网关 - ChangeDBSGatewayEIP

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **DBSGatewayID** | string | DBS网关ID | [] | [] | **Yes** |
| **EIPID** | string | 外网IPID | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 解绑备份网关 - DeleteDBSGateway

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **DBSGatewayID** | string | DBS网关ID | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 暂停备份 - PauseDBSBackup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | No |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **PlanID** | string | 计划ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 恢复定时备份 - ResumeDBSBackup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | No |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **PlanID** | string | 计划ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|

