# REDIS

## 创建redis实例 - CreateRedis

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **ProjectID** | string |  | [] | [] | No |
| **TagKeyValuePairs** | array<string> | 标签键值对 | [] | [] | **Yes** |
| **Name** | string | 名称 | [] | [] | **Yes** |
| **Remark** | string | 描述 | [] | [] | No |
| **VPCID** | string | vpcID | [] | [] | **Yes** |
| **SubnetID** | string | 子网ID | [] | [] | **Yes** |
| **EIPID** | string |  | [] | [] | No |
| **WANSGID** | string |  | [] | [] | No |
| **ChargeType** | string | 计费类型 | [] | [] | **Yes** |
| **Quantity** | integer | 计费时长 | [] | [] | No |
| **Count** | string | 购买数量 | [] | [] | **Yes** |
| **Memory** | integer | 内存容量 | [] | [] | **Yes** |
| **Password** | string | 密码 | [] | [] | No |
| **VMType** | string | 集群类型 | [] | [] | **Yes** |
| **DiskSetType** | string | 存储集群 | [] | [] | No |
| **ConfigID** | string | 配置文件ID | [] | [] | **Yes** |
| **HighAvailability** | string | 高可用 | [] | [] | **Yes** |
| **BackupID** | string | 备份ID, 用于从备份创建 | [] | [] | No |
| **Version** | string | Redis 版本，默认为'4.0'，也可以为'7.0' | [] | [] | No |
| **Threads** | integer | 数量线程，默认为2 | [] | [] | No |
| **TerminationPolicy** | integer | 是否开启删除保护, 0:开启; 1:关闭 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **RedisID** | string |  | [] |


## 查询redis实例 - DescribeRedis

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Offset** | integer | 分页偏移量 | [] | [] | No |
| **Limit** | integer | 分页数量 | [] | [] | No |
| **Keyword** | string | 搜索关键字 | [] | [] | No |
| **RedisIDs** | array<string> | Redis ID | [] | [] | No |
| **ProjectIDs** | array<string> | 项目组ID | [] | [] | No |
| **Status** | array<string> | Redis状态 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer |  | [] |
| **Infos** | array<[*RedisInfo*](#RedisInfo)> |  | [] |

### 数据模型

#### MemoryUsage

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Size** | integer | 内存大小 mb | [] |
| **Used** | integer | 已使用的内存大小 mb | [] |

#### RedisEndpoint

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **IP** | string | IP地址 | [] |
| **IPVersion** | string | IP版本 | [] |
| **NicType** | string | 网卡类型 | [] |
| **Endpoint** | string | 入口 | [] |
| **SGID** | string | 安全组ID | [] |
| **SGName** | string | 安全组名称 | [] |

#### RedisInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **RedisID** | string | RedisID | [] |
| **Name** | string | 名称 | [] |
| **Remark** | string | 描述 | [] |
| **Region** | string | 地域 | [] |
| **RegionAlias** | string | 地域别名 | [] |
| **CompanyID** | integer | 租户ID | [] |
| **CompanyName** | string | 租户名称 | [] |
| **Email** | string |  | [] |
| **ProjectID** | string | 项目组ID | [] |
| **ProjectName** | string | 项目组名称 | [] |
| **Reason** | string | 失败原因 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> | 标签 | [] |
| **ChargeType** | string | 计费类型 | [] |
| **ExpireTime** | integer | 过期时间 | [] |
| **VPCID** | string | vpcID | [] |
| **VPCName** | string | VPC名称 | [] |
| **SubnetID** | string | 子网ID | [] |
| **SubnetName** | string | 子网名称 | [] |
| **Memory** | integer | 内存 | [] |
| **SetType** | string | 计算集群 | [] |
| **Arch** | string | 计算集群架构 | [] |
| **Endpoints** | array<[*RedisEndpoint*](#RedisEndpoint)> |  | [] |
| **Status** | string | 状态 | [] |
| **ConfigID** | string | 配置ID | [] |
| **StorageSetType** | string | 存储集群 | [] |
| **HighAvailability** | string | 高可用类型 | [] |
| **Role** | string | 角色 | [] |
| **MemoryUsage** | object | 内存使用情况 | [] |
| **SlaveInfos** | array<[*SlaveRedisInfo*](#SlaveRedisInfo)> |  | [] |
| **SlaveCount** | integer |  | [] |
| **Version** | string | redis版本 | [] |
| **Threads** | integer | 数量线程 | [] |
| **TerminationPolicy** | integer | 是否开启删除保护 | [] |

#### SlaveRedisInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **RedisID** | string | RedisID | [] |
| **Name** | string | 名称 | [] |
| **Remark** | string | 描述 | [] |
| **VPCID** | string |  | [] |
| **VPCName** | string | VPC名称 | [] |
| **SubnetID** | string | 子网ID | [] |
| **SubnetName** | string | 子网名称 | [] |
| **ChargeType** | string | 计费类型 | [] |
| **CreateTime** | integer |  | [] |
| **UpdateTime** | integer |  | [] |
| **ExpireTime** | integer | 过期时间 | [] |
| **Memory** | integer | 内存大小 | [] |
| **SetType** | string | 计算集群 | [] |
| **Arch** | string | 计算集群架构 | [] |
| **Endpoints** | array<[*RedisEndpoint*](#RedisEndpoint)> |  | [] |
| **Region** | string | 地域 | [] |
| **Status** | string | 状态 | [] |
| **ConfigID** | string | 配置ID | [] |
| **CompanyID** | integer | 租户ID | [] |
| **Email** | string | 邮箱 | [] |
| **ProjectID** | string | 项目组 | [] |
| **ProjectName** | string | 项目组名称 | [] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> | 标签 | [] |
| **StorageSetType** | string | 存储集群类型 | [] |
| **HighAvailability** | string | 高可用类型 | [] |
| **Role** | string | 角色 | [] |
| **MasterID** | string | 主RedisID | [] |
| **Reason** | string | 失败原因 | [] |
| **MemoryUsage** | object | 内存使用情况 | [] |
| **Version** | string | redis版本 | [] |
| **Threads** | integer | 数量线程 | [] |
| **TerminationPolicy** | integer | 是否开启删除保护 0:开启; 1:关闭 | [] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

## 获取redis创建升级价格 - GetRedisPrice

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ChargeType** | string | 计费类型 | [] | [] | No |
| **VMType** | string | 计算集群 | [] | [] | **Yes** |
| **Memory** | integer | 内存大小 | [] | [] | **Yes** |
| **DiskSetType** | string | 存储集群 | [] | [] | No |
| **Quantity** | integer | 计费周期数量 | [] | [] | No |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **RedisID** | string | Redis ID | [] | [] | No |
| **HighAvailability** | string | 高可用类型 | [] | [] | **Yes** |
| **Threads** | integer | 数量线程 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Info** | object |  | [] |

### 数据模型

#### RedisPriceInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ChargeType** | string | 计费类型 | [] |
| **OrderType** | string | 订单类型 | [] |
| **Price** | double | 价格 | [] |

## 升级redis内存 - UpgradeRedis

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Memory** | integer | 内存大小 | [] | [] | **Yes** |
| **RedisID** | string | Redis ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 更新redis密码 - UpdateRedisPassword

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **RedisID** | string | Redis ID | [] | [] | **Yes** |
| **Password** | string | 新密码 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 更新配置项 - UpdateRedisConfigParams

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **ConfigID** | string |  | [] | [] | **Yes** |
| **Params** | array<string> | Redis配置项 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*RedisConfigOpLogInfo*](#RedisConfigOpLogInfo)> |  | [] |

### 数据模型

#### RedisConfigOpLogInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ParameterName** | string | 参数名称 | [] |
| **OldValue** | string | 修改前的值 | [] |
| **NewValue** | string | 修改后的值 | [] |
| **UpdateTime** | integer | 更新时间 | [] |

## 创建配置文件 - CreateRedisConfigFile

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **Name** | string | 名称 | [] | [] | **Yes** |
| **Description** | string | 描述 | [] | [] | No |
| **Version** | string | 版本 | [] | [] | **Yes** |
| **ConfigID** | string | 原配置ID | [] | [] | No |
| **Params** | string | 配置项 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ConfigID** | string |  | [] |


## 查询配置文件列表 - DescribeRedisConfigFile

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ConfigID** | string | 配置ID | [] | [] | No |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **Version** | string | redis版本，默认为4.0 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer |  | [] |
| **Infos** | array<[*RedisConfigFile*](#RedisConfigFile)> |  | [] |

### 数据模型

#### RedisConfigFile

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ConfigID** | string | 配置ID | [] |
| **CompanyID** | integer | 租户 | [] |
| **CompanyName** | string | 租户名称 | [] |
| **CompanyEmail** | string | 租户邮箱 | [] |
| **Name** | string | 配置名称 | [] |
| **Description** | string | 配置描述 | [] |
| **Version** | string | 版本 | [] |
| **FileType** | string | 类型 | [] |
| **CreateTime** | integer |  | [] |
| **UpdateTime** | integer |  | [] |

## 查询配置文件详情 - DescribeRedisConfigParams

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **ConfigID** | string | 配置ID | [] | [] | **Yes** |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ConfigID** | string |  | [] |
| **TotalCount** | integer |  | [] |
| **Infos** | array<[*RedisRedisConfigParam*](#RedisRedisConfigParam)> |  | [] |

### 数据模型

#### RedisRedisConfigParam

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ConfigID** | string | 配置ID | [] |
| **Key** | string | 配置项 | [] |
| **Value** | string | 配置值 | [] |
| **ValueType** | string | 值类型 | [] |
| **Description** | string | 配置项描述 | [] |
| **CreateTime** | integer |  | [] |
| **UpdateTime** | integer |  | [] |
| **NeedRestart** | integer | 是否需要重启，0:不需要；1:需要 | [] |

## 删除配置文件 - DeleteRedisConfigFile

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **ConfigID** | string | 配置ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 应用Redis参数模板 - ApplyRedisConfigFile

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **ConfigID** | string | 配置ID | [] | [] | **Yes** |
| **RedisID** | string | RedisID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 清空数据 - FlushRedis

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **RedisID** | string | Redis ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 删除redis实例 - DeleteRedis

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **RedisID** | string | Redis ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 升级至主备版本 - UpgradeRedisToHA

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **RedisID** | string | Redis ID | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 创建Redis从库 - CreateSlaveRedis

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **Name** | string | 名称 | [] | [] | **Yes** |
| **Remark** | string | 描述 | [] | [] | No |
| **Memory** | integer | 内存容量 | [] | [] | **Yes** |
| **RedisID** | string | 主库ID | [] | [] | **Yes** |
| **EIPID** | string | EIP ID | [] | [] | No |
| **WANSGID** | string | 安全组ID | [] | [] | No |
| **ChargeType** | string | 计费类型 | [] | [] | **Yes** |
| **Quantity** | integer | 计费时长 | [] | [] | No |
| **VMType** | string | 集群类型 | [] | [] | **Yes** |
| **TagKeyValuePairs** | array<string> | 标签键值对 | [] | [] | **Yes** |
| **Password** | string | 密码 | [] | [] | No |
| **Version** | string | redis版本，默认为4.0 | [] | [] | No |
| **Threads** | integer | 数量线程，Redis7.0默认为2, Redis4.0默认为1 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **RedisID** | string |  | [] |


## 慢日志查询 - DescribeRedisSlowlog

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **RedisID** | string | Redis ID | [] | [] | **Yes** |
| **StartTime** | integer | 开始时间戳 | [] | [] | **Yes** |
| **EndTime** | integer | 结束时间戳 | [] | [] | **Yes** |
| **Limit** | integer | 分页 | [] | [] | No |
| **Offset** | integer | 偏移量 | [] | [] | No |
| **Keyword** | string | 关键词 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数 | [] |
| **Infos** | array<[*RedisSlowlog*](#RedisSlowlog)> | 慢日志信息 | [] |

### 数据模型

#### RedisSlowlog

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **StartTime** | integer | 执行时间 | [] |
| **ExecSpend** | integer | 耗时（微秒） | [] |
| **Command** | string | 执行命令 | [] |
| **ClientAddr** | string | 客户端地址 | [] |
| **ClientName** | string | 客户端名称 | [] |

## 降级redis内存 - DowngradeRedis

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Memory** | integer | 内存大小 | [] | [] | **Yes** |
| **RedisID** | string | Redis ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|

