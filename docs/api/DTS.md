# DTS

## 创建数据传输任务 - CreateDTSTask

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **ProjectID** | string | 项目ID | [] | [] | No |
| **TagKeyValuePairs** | array<string> | 标签键值对 | [] | [] | **Yes** |
| **Name** | string | 名称 | [] | [] | **Yes** |
| **Remark** | string | 备注 | [] | [] | No |
| **ChargeType** | string | 计费类型 | [] | [] | **Yes** |
| **Quantity** | integer | 计费数量 | [] | [] | **Yes** |
| **MaxRPS** | integer | 最大每秒同步记录数 | [] | [] | **Yes** |
| **ComputeClass** | string | 计算集群 | [] | [] | **Yes** |
| **StorageClass** | string | 存储集群 | [] | [] | **Yes** |
| **EIPID** | string | 外网IP, 源端或目的端存在外部数据源时必填, 可以与源或目的端网络通信 | [] | [] | No |
| **CPU** | integer | CPU | [] | [] | **Yes** |
| **Memory** | integer | 内存 | [] | [] | **Yes** |
| **TaskMode** | string | 任务类型, 全量,增量,全量加增量 | [] | [] | **Yes** |
| **SourceEndpointInstanceType** | string | 源实例类型, 平台内部或外部数据库 | [] | [] | **Yes** |
| **SourceEndpointInstanceID** | string | 源实例 ID, 当 SourceEndpointInstanceType 为 Internal 时需要传入实例 ID | [] | [] | No |
| **SourceEngine** | string | 源实例的数据库类型 | [] | [] | **Yes** |
| **SourceEndpointIP** | string | 源实例地址, 当 SourceEndpointInstanceType 为 External 时需要填入 | [] | [] | No |
| **SourceEndpointPort** | integer | 源实例端口, 当 SourceEndpointInstanceType 为 External 时需要填入 | [] | [] | No |
| **SourceEndpointUserName** | string | 源实例用户名 | [] | [] | No |
| **SourceEndpointPassword** | string | 源实例密码 | [] | [] | No |
| **SourceEndpointBinlogName** | string | 增量时需要指定的 binlog name，可以通过 show master status 获取，或者全量+增量任务会自动设置 | [] | [] | No |
| **SourceEndpointBinlogPos** | integer | 增量时需要指定的 binlog pos，可以通过 show master status 获取，或者全量+增量任务会自动设置 | [] | [] | No |
| **SourceEndpointBinlogGTID** | string | 增量时需要指定的 binlog gtid，可以通过 show master status 获取，或者全量+增量任务会自动设置 | [] | [] | No |
| **SourceEndpointServerID** | string | 增量时需要指定的 serverID，不能和现有的 slave 重复 | [] | [] | No |
| **DestinationEngine** | string | 目标实例的数据库类型,需要与源端一致 | [] | [] | **Yes** |
| **DestinationEndpointInstanceType** | string | 目标实例类型, 平台内部或外部数据库 | [] | [] | **Yes** |
| **DestinationEndpointInstanceID** | string | 目标实例 ID, 当 DestinationEndpointInstanceType 为 Internal 时需要传入实例 ID | [] | [] | No |
| **DestinationEndpointIP** | string | 目标实例地址, 当 DestinationEndpointInstanceType 为 External 时需要填入 | [] | [] | No |
| **DestinationEndpointPort** | integer | 目标实例端口, 当 DestinationEndpointInstanceType 为 External 时需要填入 | [] | [] | No |
| **DestinationEndpointUserName** | string | 目标实例用户名 | [] | [] | No |
| **DestinationEndpointPassword** | string | 目标实例密码 | [] | [] | No |
| **DestinationEndpointBinlogGTID** | string | 双向同步时, 需要配置目标库的 GTID, 可以通过 show master status 获取, 用于反向同步 | [] | [] | No |
| **DataMarkTable** | string | 双向同步时, 同步服务需要记录同步数据, 如 database.mark_table, 如果没有指定, 同步服务会自己创建, 请确保账户有相应权限 | [] | [] | No |
| **Databases** | string | 需要同步的数据库, 未指定同步所有数据库, db1,db2 | [] | [] | No |
| **Tables** | string | 需要同步的数据表, 未指定则同步所有数据表, db1.*, db2.table_1 | [] | [] | No |
| **IgnoreDatabases** | string | 需要忽略的数据库 | [] | [] | No |
| **IgnoreTables** | string | 需要忽略的数据表 | [] | [] | No |
| **HeartbeatTable** | string | 增量同步或双向同步时, 同步服务需要定时推进 binlog 位点, 如 database.heartbeat_table, 如果没有指定, 同步服务会自己创建, 请确保账户有相应权限 | [] | [] | No |
| **HeartbeatInterval** | integer | 多少秒推进一次位点, 默认 30s | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **DTSID** | string |  | [] |


## 更新数据传输任务配置 - UpdateDTSTaskConfigure

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **DTSID** | string | 任务 ID | [] | [] | **Yes** |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **MaxRPS** | integer | MaxRPS | [] | [] | **Yes** |
| **SourceEndpointInstanceType** | string | 源实例类型, 平台内部或外部数据库 | [] | [] | **Yes** |
| **SourceEndpointInstanceID** | string | 源实例 ID, 当 SourceEndpointInstanceType 为 Internal 时需要传入实例 ID | [] | [] | No |
| **SourceEngine** | string | 源实例的数据库类型 | [] | [] | **Yes** |
| **SourceEndpointIP** | string | 源实例地址, 当 SourceEndpointInstanceType 为 External 时需要填入 | [] | [] | No |
| **SourceEndpointPort** | integer | 源实例端口, 当 SourceEndpointInstanceType 为 External 时需要填入 | [] | [] | No |
| **SourceEndpointUserName** | string | 源实例用户名 | [] | [] | No |
| **SourceEndpointPassword** | string | 源实例密码 | [] | [] | No |
| **SourceEndpointBinlogName** | string | 增量时需要指定的 binlog name，可以通过 show master status 获取，或者全量+增量任务会自动设置 | [] | [] | No |
| **SourceEndpointBinlogPos** | integer | 增量时需要指定的 binlog pos，可以通过 show master status 获取，或者全量+增量任务会自动设置 | [] | [] | No |
| **SourceEndpointBinlogGTID** | string | 增量时需要指定的 binlog gtid，可以通过 show master status 获取，或者全量+增量任务会自动设置 | [] | [] | No |
| **SourceEndpointServerID** | string | 增量时需要指定的 serverID，不能和现有的 slave 重复 | [] | [] | No |
| **DestinationEngine** | string | 目标实例的数据库类型,需要与源端一致 | [] | [] | **Yes** |
| **DestinationEndpointInstanceType** | string | 目标实例类型, 平台内部或外部数据库 | [] | [] | **Yes** |
| **DestinationEndpointInstanceID** | string | 目标实例 ID, 当 DestinationEndpointInstanceType 为 Internal 时需要传入实例 ID | [] | [] | No |
| **DestinationEndpointIP** | string | 目标实例地址, 当 DestinationEndpointInstanceType 为 External 时需要填入 | [] | [] | No |
| **DestinationEndpointPort** | integer | 目标实例端口, 当 DestinationEndpointInstanceType 为 External 时需要填入 | [] | [] | No |
| **DestinationEndpointUserName** | string | 目标实例用户名 | [] | [] | No |
| **DestinationEndpointPassword** | string | 目标实例密码 | [] | [] | No |
| **DestinationEndpointBinlogGTID** | string | 双向同步时, 需要配置目标库的 GTID, 可以通过 show master status 获取, 用于反向同步 | [] | [] | No |
| **DataMarkTable** | string | 双向同步时, 同步服务需要记录同步数据, 如 database.mark_table, 如果没有指定, 同步服务会自己创建, 请确保账户有相应权限 | [] | [] | No |
| **Databases** | string | 需要同步的数据库, 未指定同步所有数据库, db1,db2 | [] | [] | No |
| **Tables** | string | 需要同步的数据表, 未指定则同步所有数据表, db1.*, db2.table_1 | [] | [] | No |
| **IgnoreDatabases** | string | 需要忽略的数据库 | [] | [] | No |
| **IgnoreTables** | string | 需要忽略的数据表 | [] | [] | No |
| **HeartbeatTable** | string | 增量同步或双向同步时, 同步服务需要定时推进 binlog 位点, 如 database.heartbeat_table, 如果没有指定, 同步服务会自己创建, 请确保账户有相应权限 | [] | [] | No |
| **HeartbeatInterval** | integer | 多少秒推进一次位点, 默认 30s | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取传输任务列表 - DescribeDTSTask

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer |  | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **DTSIDs** | array<string> |  | [] | [] | No |
| **ProjectIDs** | array<string> |  | [] | [] | No |
| **Keyword** | string |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer |  | [] |
| **Infos** | array<[*DTSTaskInfo*](#DTSTaskInfo)> |  | [] |

### 数据模型

#### DTSTaskInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string | 名称 | [] |
| **Remark** | string | 备注 | [] |
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
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> | 标签 | [] |
| **ChargeType** | string | 计费类型 | [] |
| **ExpireTime** | integer | 过期时间 | [] |
| **DTSID** | string | 任务ID | [] |
| **Status** | string | 状态 | [] |
| **TaskMode** | string | 任务类型 | [] |
| **TaskStages** | array<[*TaskStage*](#TaskStage)> | 任务阶段以及对应的状态 | [] |
| **MaxRPS** | integer | 每秒最大同步数据 | [] |
| **EIPID** | string | EIPID | [] |
| **EIPName** | string | EIP name | [] |
| **ComputeClass** | string | 计算集群 | [] |
| **StorageClass** | string | 存储集群 | [] |
| **EIP** | string | EIP 地址 | [] |
| **CPU** | integer | CPU | [] |
| **Memory** | integer | 内存,单位 GB | [] |
| **Arch** | string | 架构 | [] |
| **AllowStart** | integer | 是否允许启动, 0: 不允许, 1: 允许 | [] |

#### TaskStage

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Stage** | string |  | [] |
| **Status** | string |  | [] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

## 获取传输任务配置 - GetDTSTaskConfigure

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **DTSID** | string | 任务 ID | [] | [] | **Yes** |
| **Region** | string | 地域 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **DTSID** | string | 任务 ID | [] |
| **SourceEndpointInstanceType** | string | 源实例类型, 平台内部或外部数据库 | [] |
| **SourceEndpointInstanceID** | string | 源实例 ID | [] |
| **SourceEngine** | string | 源实例的数据库类型 | [] |
| **SourceEndpointIP** | string | 源实例地址 | [] |
| **SourceEndpointPort** | integer | 源实例端口 | [] |
| **SourceEndpointUserName** | string | 源实例用户名 | [] |
| **SourceEndpointPassword** | string | 源实例密码 | [] |
| **SourceEndpointBinlogName** | string | 增量时需要指定的 binlog name | [] |
| **SourceEndpointBinlogPos** | integer | 增量时需要指定的 binlog pos | [] |
| **SourceEndpointBinlogGTID** | string | 增量时需要指定的 binlog gtid | [] |
| **SourceEndpointServerID** | string | 增量时需要指定的 serverID | [] |
| **SourceEndpointInstanceName** | string | 源实例名称 | [] |
| **DestinationEndpointInstanceType** | string | 目标实例类型 | [] |
| **DestinationEndpointInstanceID** | string | 目标实例 ID | [] |
| **DestinationEndpointIP** | string | 目标实例地址 | [] |
| **DestinationEndpointPort** | integer | 目标实例端口 | [] |
| **DestinationEndpointUserName** | string | 目标实例用户名 | [] |
| **DestinationEndpointPassword** | string | 目标实例密码 | [] |
| **DestinationEndpointBinlogGTID** | string | 双向同步时, 需要配置目标库的 GTID | [] |
| **DestinationEngine** | string | 目标实例的数据库类型 | [] |
| **DestinationEndpointInstanceName** | string | 目标实例名称 | [] |
| **DataMarkTable** | string | 双向同步时, 同步服务需要记录同步数据, 如 database.mark_table, 如果没有指定, 同步服务会自己创建, 请确保账户有相应权限 | [] |
| **Databases** | string | 需要同步的数据库, 未指定同步所有数据库, db1,db2 | [] |
| **Tables** | string | 需要同步的数据表, 未指定则同步所有数据表, db1.*, db2.table_1 | [] |
| **IgnoreDatabases** | string | 需要忽略的数据库 | [] |
| **IgnoreTables** | string | 需要忽略的数据表 | [] |
| **HeartbeatTable** | string | 增量同步或双向同步时, 同步服务需要定时推进 binlog 位点, 如 database.heartbeat_table, 如果没有指定, 同步服务会自己创建, 请确保账户有相应权限 | [] |
| **HeartbeatInterval** | integer | 多少秒推进一次位点, 默认 30s | [] |


## 启动 DTS 任务 - StartDTSTask

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **DTSID** | string | 任务 ID | [] | [] | **Yes** |
| **Region** | string | 地域 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 暂停 DTS 任务 - SuspendDTSTask

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **DTSID** | string | 任务 ID | [] | [] | **Yes** |
| **Region** | string | 地域 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 删除 DTS 任务 - DeleteDTSTask

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **DTSID** | string | 任务 ID | [] | [] | **Yes** |
| **Region** | string | 地域 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 创建数据校验任务 - CreateDataCheckTask

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **DTSID** | string | 需要进行校验的任务 ID | [] | [] | **Yes** |
| **SampleInterval** | integer | 抽样频率, 每多少条记录采样一次 | [] | [] | **Yes** |
| **Region** | string | 地域 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **DTSID** | string | 校验任务 ID | [] |


## 获取数据校验任务列表 - DescribeDataCheckTask

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **DTSID** | string | DTS 任务ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数量 | [] |
| **Infos** | array<[*DataCheckTaskOverview*](#DataCheckTaskOverview)> |  | [] |

### 数据模型

#### DataCheckTaskOverview

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TaskID** | string | 校验任务 ID | [] |
| **SampleInterval** | integer | 采样频率 | [] |
| **Status** | string | 任务状态 | [] |
| **Result** | string | 校验结果, Consistent, Inconsistent | [] |
| **StartTime** | integer | 开始时间 | [] |
| **EndTime** | integer | 结束时间 | [] |

## 获取数据校验任务详细结果 - GetDataCheckTaskResult

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **TaskID** | string | 校验任务ID | [] | [] | **Yes** |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **DTSID** | string | 数据传输任务ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Result** | string | 校验结果 | [] |
| **Logs** | array<string> | 数据表详细差异信息 | [] |


## 获取数据传输任务价格 - GetDTSPrice

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **ChargeType** | string | 计费类型 | [] | [] | **Yes** |
| **Quantity** | integer | 计费数量 | [] | [] | **Yes** |
| **Count** | integer | 数量 | [] | [] | **Yes** |
| **StorageClass** | string | 存储集群 | [] | [] | **Yes** |
| **ComputeClass** | string | 计算集群 | [] | [] | **Yes** |
| **CPU** | integer | CPU | [] | [] | **Yes** |
| **Memory** | integer | 内存 | [] | [] | **Yes** |
| **DTSID** | string | 实例 ID, 升级配置时计算费用 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*DTSPriceInfo*](#DTSPriceInfo)> |  | [] |

### 数据模型

#### DTSPriceInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ChargeType** | string | 计费类型 | [] |
| **OrderType** | string | 账单类型 | [] |
| **Price** | double | 价格 | [] |

## 获取任务日志 - DescribeDTSLog

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **DTSID** | string | DTS 任务ID | [] | [] | **Yes** |
| **DTSStage** | string | 任务阶段 | [] | [] | **Yes** |
| **Type** | string | 日志类型 | [] | [] | **Yes** |
| **StartTime** | integer | 开始时间, 默认为当前时间前一小时 | [] | [] | No |
| **EndTime** | integer | 结束时间, 默认为当前时间 | [] | [] | No |
| **LogLimit** | integer | 返回数量, 默认为 1000 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer |  | [] |
| **Logs** | array<string> |  | [] |


## 更新 DTS 实例规格 - UpdateDTSInstanceSpec

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **DTSID** | string | DTS 任务ID | [] | [] | **Yes** |
| **CPU** | integer | CPU | [] | [] | **Yes** |
| **Memory** | integer | 内存 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 执行数据传输预检查 - RunDTSPrecheck

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **TaskMode** | string | 任务类型, 全量,增量,全量加增量 | [] | [] | **Yes** |
| **SourceEndpointInstanceType** | string | 源实例类型, 平台内部或外部数据库 | [] | [] | **Yes** |
| **SourceEndpointInstanceID** | string | 源实例 ID, 当 SourceEndpointInstanceType 为 Internal 时需要传入实例 ID | [] | [] | No |
| **SourceEngine** | string | 源实例的数据库类型 | [] | [] | **Yes** |
| **SourceEndpointIP** | string | 源实例地址, 当 SourceEndpointInstanceType 为 External 时需要填入 | [] | [] | No |
| **SourceEndpointPort** | integer | 源实例端口, 当 SourceEndpointInstanceType 为 External 时需要填入 | [] | [] | No |
| **SourceEndpointUserName** | string | 源实例用户名 | [] | [] | No |
| **SourceEndpointPassword** | string | 源实例密码 | [] | [] | No |
| **SourceEndpointBinlogName** | string | 增量时需要指定的 binlog name，可以通过 show master status 获取，或者全量+增量任务会自动设置 | [] | [] | No |
| **SourceEndpointBinlogPos** | integer | 增量时需要指定的 binlog pos，可以通过 show master status 获取，或者全量+增量任务会自动设置 | [] | [] | No |
| **SourceEndpointBinlogGTID** | string | 增量时需要指定的 binlog gtid，可以通过 show master status 获取，或者全量+增量任务会自动设置 | [] | [] | No |
| **SourceEndpointServerID** | string | 增量时需要指定的 serverID，不能和现有的 slave 重复 | [] | [] | No |
| **DestinationEngine** | string | 目标实例的数据库类型,需要与源端一致 | [] | [] | **Yes** |
| **DestinationEndpointInstanceType** | string | 目标实例类型, 平台内部或外部数据库 | [] | [] | **Yes** |
| **DestinationEndpointInstanceID** | string | 目标实例 ID, 当 DestinationEndpointInstanceType 为 Internal 时需要传入实例 ID | [] | [] | No |
| **DestinationEndpointIP** | string | 目标实例地址, 当 DestinationEndpointInstanceType 为 External 时需要填入 | [] | [] | No |
| **DestinationEndpointPort** | integer | 目标实例端口, 当 DestinationEndpointInstanceType 为 External 时需要填入 | [] | [] | No |
| **DestinationEndpointUserName** | string | 目标实例用户名 | [] | [] | No |
| **DestinationEndpointPassword** | string | 目标实例密码 | [] | [] | No |
| **DestinationEndpointBinlogGTID** | string | 双向同步时, 需要配置目标库的 GTID, 可以通过 show master status 获取, 用于反向同步 | [] | [] | No |
| **Databases** | string | 需要同步的数据库, 未指定同步所有数据库, db1,db2 | [] | [] | No |
| **Tables** | string | 需要同步的数据表, 未指定则同步所有数据表, db1.*, db2.table_1 | [] | [] | No |
| **IgnoreDatabases** | string | 需要忽略的数据库 | [] | [] | No |
| **IgnoreTables** | string | 需要忽略的数据表 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **PrecheckItems** | array<[*DTSPrecheckItem*](#DTSPrecheckItem)> |  | [] |

### 数据模型

#### DTSPrecheckItem

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **PreCheckName** | string | 检查项 | [] |
| **Passed** | bool | 检查是否通过 | [] |
| **Reason** | string | 未通过原因 | [] |
