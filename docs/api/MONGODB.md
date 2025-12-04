# MONGODB

## 创建 MongoDB - CreateMongoDB

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **ProjectID** | string | 项目组ID | [] | [] | No |
| **TagKeyValuePairs** | array<string> | 标签键值对 | [] | [] | **Yes** |
| **Name** | string | 名字 | [] | [] | **Yes** |
| **Remark** | string | 描述 | [] | [] | No |
| **ChargeType** | string | 计费类型 | [] | [] | **Yes** |
| **Quantity** | integer | 计费时长 | [] | [] | **Yes** |
| **ComputeSetClass** | string | 虚拟机所在宿主机的集群类型 | [] | [] | **Yes** |
| **CPU** | integer | CPU | [] | [] | **Yes** |
| **Memory** | integer | 内存 | [] | [] | **Yes** |
| **StorageSetClass** | string | 存储集群 | [] | [] | **Yes** |
| **DiskSpace** | integer | 存储容量 | [] | [] | **Yes** |
| **VPCID** | string | vpcID | [] | [] | **Yes** |
| **SubnetID** | string | 子网ID | [] | [] | **Yes** |
| **Password** | string | root 密码 | [] | [] | **Yes** |
| **ConfigFileID** | string | 配置文件ID | [] | [] | **Yes** |
| **BackupID** | string | 备份文件ID | [] | [] | No |
| **Version** | string | MongoDB 版本，必须为 MongoDB 5.0/6.0/7.0 | [] | [] | **Yes** |
| **ClusterType** | string | 实例架构类型 | [] | [] | **Yes** |
| **NodeCount** | integer | 副本集节点数量或者分片集群中每个分片内节点数量 | [] | [] | **Yes** |
| **RestoreTime** | integer | 备份恢复的时间点 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **MongoDBID** | string | MongoDBID | [] |


## 查询 MongoDB 信息 - DescribeMongoDB

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **ProjectIDs** | array<string> | 项目组ID | [] | [] | No |
| **Keyword** | string | 搜索关键词 | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **MongoDBIDs** | array<string> | MongoDB IDs | [] | [] | No |
| **ClusterType** | string | 实例架构类型 | [] | [] | No |
| **VPCID** | string | VPC ID | [] | [] | No |
| **SubnetID** | string | Subnet ID | [] | [] | No |
| **Status** | array<string> | MongoDB 状态 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数 | [] |
| **Infos** | array<[*MongoDBInfo*](#MongoDBInfo)> | MongoDB 信息 | [] |

### 数据模型

#### ConnectInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ConnType** | string | 链接类型 Internal Internet | [] |
| **ConnectURL** | string |  | [] |

#### MongoDBInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **MongoDBID** | string | MongoDBID | [] |
| **Name** | string | 名字 | [] |
| **Remark** | string | 描述 | [] |
| **Region** | string | 地域 | [] |
| **RegionAlias** | string | 地域别名 | [] |
| **CompanyID** | integer | 租户ID | [] |
| **CompanyName** | string | 租户名 | [] |
| **Email** | string | 租户邮箱 | [] |
| **ProjectID** | string | 项目组ID | [] |
| **ProjectName** | string | 项目组名 | [] |
| **Reason** | string | 失败原因 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> | 标签 | [] |
| **ExpireTime** | integer | 过期时间 | [] |
| **ChargeType** | string | 计费类型 | [] |
| **ComputeSetClass** | string | 计算集群名称 | [] |
| **StorageSetClass** | string | 存储集群 | [] |
| **CPU** | integer | CPU | [] |
| **Memory** | integer | 内存 | [] |
| **DiskID** | string | 磁盘ID | [] |
| **DiskSpace** | integer | 磁盘容量 | [] |
| **VPCName** | string | VPC 名称 | [] |
| **SubnetName** | string | Subnet 名称 | [] |
| **VPCID** | string | VPCID | [] |
| **SubnetID** | string | SubnetID | [] |
| **Status** | string | 状态 | [] |
| **Version** | string | MongoDB 版本 | [] |
| **ClusterType** | string | 实例架构类型 | [] |
| **ConnectURLs** | array<[*ConnectInfo*](#ConnectInfo)> | 集群访问地址 | [] |
| **NodeCount** | integer |  | [] |
| **Statistics** | object | MongoDB 节点信息统计 | [] |

#### MongoDBStatistics

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Total** | integer |  | [] |
| **Ready** | integer |  | [] |
| **CurrentOperatorNode** | string | 当前操作的节点,如果是针对集群操作,不会记录 | [] |
| **NodeTypeStatistics** | array<[*NodeTypeStatistics*](#NodeTypeStatistics)> |  | [] |

#### NodeTypeStatistics

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **NodeType** | string |  | [] |
| **NodeTypeStatusStatistics** | array<[*NodeTypeStatusStatistics*](#NodeTypeStatusStatistics)> |  | [] |

#### NodeTypeStatusStatistics

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Status** | string |  | [] |
| **Count** | integer |  | [] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

## 查询 MongoDB 信息 - DescribeMongoDBDetail

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **MongoDBID** | string | MongoDB ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **DataReplicaInfos** | array<[*ReplicaInfo*](#ReplicaInfo)> |  | [] |

### 数据模型

#### MongoDBNodeInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **NodeID** | string |  | [] |
| **NodeRole** | string |  | [] |
| **CPUUtilization** | float | CPU 利用率 | [] |
| **MemUsage** | float | 内存利用率 | [] |
| **SpaceUsage** | float | 空间利用率 | [] |
| **Status** | string |  | [] |

#### ReplicaInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ReplicaID** | string |  | [] |
| **ReplicaType** | string |  | [] |
| **Status** | string |  | [] |
| **CPU** | integer |  | [] |
| **Memory** | integer |  | [] |
| **NodeInfos** | array<[*MongoDBNodeInfo*](#MongoDBNodeInfo)> |  | [] |

## 副本集群设置 Primary 节点 - SetMongoDBPrimary

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **MongoDBID** | string | MongoDBID | [] | [] | **Yes** |
| **NodeID** | string | NodeID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 水平升/降 MongoDB - ScaleMongoDB

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **MongoDBID** | string | MongoDBID | [] | [] | **Yes** |
| **ScaleType** | string |  | [] | [] | **Yes** |
| **NodeCount** | integer | NodeCount | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 升/降级 MongoDB - ResizeMongoDB

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **MongoDBID** | string | MongoDBID | [] | [] | **Yes** |
| **ResizeType** | string |  | [] | [] | **Yes** |
| **CPU** | integer | CPU | [] | [] | **Yes** |
| **Memory** | integer | 内存 | [] | [] | **Yes** |
| **DiskSpace** | integer | 存储容量 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 查询 MongoDB 价格 - GetMongoDBPrice

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **ProjectID** | string | 项目组ID | [] | [] | No |
| **ChargeType** | string | 计费类型 | [] | [] | **Yes** |
| **Quantity** | integer | 计费时长 | [] | [] | **Yes** |
| **ComputeSetClass** | string | 虚拟机所在宿主机的集群类型 | [] | [] | **Yes** |
| **CPU** | integer | CPU | [] | [] | **Yes** |
| **Memory** | integer | 内存 | [] | [] | **Yes** |
| **StorageSetClass** | string | 存储集群 | [] | [] | **Yes** |
| **DiskSpace** | integer | 存储容量 | [] | [] | **Yes** |
| **Version** | string | MongoDB 版本，必须为 MongoDB 5.0/6.0/7.0 | [] | [] | **Yes** |
| **ClusterType** | string | 实例架构类型 | [] | [] | **Yes** |
| **NodeCount** | integer | 副本集节点数量或者分片集群中每个分片内节点数量 | [] | [] | **Yes** |
| **MongoDBID** | string |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*MongoDBPriceInfo*](#MongoDBPriceInfo)> | 价格信息 | [] |

### 数据模型

#### MongoDBPriceInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ChargeType** | string | 计费类型 | [] |
| **OrderType** | string | 订单类型 | [] |
| **Price** | double | 价格 | [] |

## 删除 MongoDB - DeleteMongoDB

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **MongoDBID** | string | MongoDBID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 修改 MongoDB 密码 - ResetMongoDBPassword

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **MongoDBID** | string | MongoDBID | [] | [] | **Yes** |
| **Password** | string | 密码 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 重启 MongoDB 实例 - RestartMongoDB

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **MongoDBID** | string | MongoDBID | [] | [] | **Yes** |
| **NodeID** | string | NodeID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 创建 MongoDB 参数模板 - CreateMongoDBConfigTpl

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **Name** | string | 名称 | [] | [] | **Yes** |
| **Remark** | string | 描述 | [] | [] | No |
| **Version** | string | 版本 | [] | [] | **Yes** |
| **ClusterType** | string | 实例架构类型 | [] | [] | **Yes** |
| **SrcTplID** | string | 源模版ID | [] | [] | No |
| **Params** | string | 自定义配置项 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TplID** | string | 模版ID | [] |


## 应用 MongoDB 参数模板 - ApplyMongoDBConfigTpl

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **TplID** | string | 模版ID | [] | [] | **Yes** |
| **MongoDBID** | string | MongoDBID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 删除 MongoDB 参数模板 - DeleteMongoDBConfigTpl

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **TplID** | string | 模版ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 更新 MongoDB 参数模板 - UpdateMongoDBConfigTpl

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **TplID** | string | 模版ID | [] | [] | **Yes** |
| **Params** | string | 配置参数 | [] | [] | No |
| **Name** | string | 名称 | [] | [] | No |
| **Remark** | string | 描述 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 查询 MongoDB 参数模板列表 - DescribeMongoDBConfigTpl

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **TplIDs** | array<string> | 模版ID | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数 | [] |
| **Infos** | array<[*MongoDBParamTplInfo*](#MongoDBParamTplInfo)> | 模版信息 | [] |

### 数据模型

#### MongoDBParamTplInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **Email** | string | Email | [] |
| **ID** | string | 模版ID | [] |
| **Name** | string | 名称 | [] |
| **Region** | string | 地域 | [] |
| **Type** | string | 类型 | [] |
| **UpdateTime** | integer | 更新时间 | [] |
| **Version** | string | 版本 | [] |
| **ClusterType** | string | 实例架构类型 | [] |
| **Remark** | string |  | [] |

## 获取 MongoDB 配置参数 - DescribeMongoDBConfigParam

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户 ID | [] | [] | No |
| **MongoDBID** | string | MongoDBID | [] | [] | No |
| **TplID** | string | 模版 ID | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数 | [] |
| **Infos** | array<[*MongoDBConfigParamInfo*](#MongoDBConfigParamInfo)> |  | [] |

### 数据模型

#### MongoDBConfigParamInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **NodeType** | string | 节点类型 | [] |
| **Description** | string |  参数描述 | [] |
| **Key** | string | 参数名称 | [] |
| **Value** | string | 参数值 | [] |
| **ValueRange** | string | 参数值范围 | [] |
| **ValueType** | string | 参数类型 | [] |
| **RequireRestart** | bool | 是否需要重启 | [] |

## 更新 MongoDB 配置参数 - UpdateMongoDBConfigParam

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **MongoDBID** | string | MongoDBID 会立即应用配置 | [] | [] | No |
| **TplID** | string | TplID | [] | [] | No |
| **Params** | string | 配置参数 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|

