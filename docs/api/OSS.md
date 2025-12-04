# OSS

## 创建对象存储服务 - CreateOSS

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer |  | [] | [] | **Yes** |
| **ProjectID** | string |  | [] | [] | No |
| **TagKeyValuePairs** | array<string> | 标签键值对 | [] | [] | **Yes** |
| **Name** | string |  | [] | [] | **Yes** |
| **Remark** | string |  | [] | [] | No |
| **ChargeType** | string | 计费类型 | [] | [] | **Yes** |
| **Quantity** | integer | 计费时长 | [] | [] | **Yes** |
| **VMType** | string |  | [] | [] | **Yes** |
| **DiskSetType** | string | 存储集群 | [] | [] | **Yes** |
| **DiskSpace** | integer | 存储容量, 最小值100G | [] | [] | **Yes** |
| **VPCID** | string | vpcID | [] | [] | **Yes** |
| **SubnetID** | string | 子网ID | [] | [] | **Yes** |
| **EIPID** | string |  | [] | [] | No |
| **WANSGID** | string | 外网安全组 | [] | [] | No |
| **Password** | string |  | [] | [] | **Yes** |
| **CPU** | integer |  | [] | [] | No |
| **BackupID** | string | 备份ID | [] | [] | No |
| **TerminationPolicy** | integer | 是否开启删除保护, 0:开启; 1:关闭 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **OSSID** | string | 对象存储ID | [] |


## 删除对象存储服务 - DeleteOSS

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string |  | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **OSSID** | string | 对象存储ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 对象存储扩容 - UpgradeOSS

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **OSSID** | string | 对象存储ID | [] | [] | **Yes** |
| **DiskSpace** | integer | 磁盘大小 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **CPU** | integer |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取对象存储列表 - DescribeOSS

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | No |
| **ProjectIDs** | array<string> | 项目组IDs | [] | [] | No |
| **Keyword** | string | 搜索关键词 | [] | [] | No |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **OSSIDs** | array<string> | 对象存储IDs | [] | [] | No |
| **VpcIDs** | array<string> | vpcIDs | [] | [] | No |
| **Filter** | string | 过滤标志 | [] | [] | No |
| **Status** | array<string> | OSS状态 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer |  | [] |
| **Infos** | array<[*OSSInfo*](#OSSInfo)> |  | [] |

### 数据模型

#### OSSInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **OSSID** | string | 对象存储ID | [] |
| **Name** | string | 名称 | [] |
| **Remark** | string | 描述 | [] |
| **Region** | string | 地域 | [] |
| **RegionAlias** | string | 地域别名 | [] |
| **CompanyID** | integer | 租户ID | [] |
| **CompanyName** | string | 租户名称 | [] |
| **Email** | string | 邮箱 | [] |
| **ProjectID** | string | 项目组ID | [] |
| **ProjectName** | string | 项目组名称 | [] |
| **Reason** | string | 失败原因 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> | 标签 | [] |
| **ChargeType** | string | 计费类型 | [] |
| **ExpireTime** | integer | 过期时间 | [] |
| **VMTypeAlias** | string | 计算集群名称 | [] |
| **VPCName** | string | VPC名称 | [] |
| **SubnetName** | string | 子网名称 | [] |
| **VPCID** | string | vpcID | [] |
| **SubnetID** | string | 子网ID | [] |
| **Endpoints** | array<[*OssEndpoint*](#OssEndpoint)> |  | [] |
| **Status** | string | 状态 | [] |
| **DiskSpace** | integer | 磁盘大小 | [] |
| **StorageSetType** | string | 存储集群类型 | [] |
| **IOPS** | integer | IOPS | [] |
| **Bandwidth** | integer | 带宽 | [] |
| **CPU** | integer | CPU数量 | [] |
| **SetArch** | string | 计算集群架构 | [] |
| **StorageSetArch** | string | 存储集群架构 | [] |
| **TerminationPolicy** | integer | 是否开启删除保护 | [] |
| **VMType** | string | 计算集群类型 | [] |
| **StorageSetTypeAlias** | string | 存储集群名称 | [] |

#### OssEndpoint

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **IP** | string | IP地址 | [] |
| **IPID** | string | IPID | [] |
| **IPVersion** | string | IP协议版本 | [] |
| **NicType** | string | 网卡类型 | [] |
| **Endpoint** | string | 入口地址 | [] |
| **SGID** | string | 安全组ID | [] |
| **SGName** | string | 安全组名称 | [] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

## 重置对象存储密码 - ResetOSSPassword

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **OSSID** | string | 对象存储ID | [] | [] | **Yes** |
| **Password** | string | 密码 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取对象存储价格 - GetOSSPrice

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ChargeType** | string | 计费类型 | [] | [] | No |
| **VMType** | string | 计算集群 | [] | [] | **Yes** |
| **DiskSetType** | string | 存储集群 | [] | [] | **Yes** |
| **DiskSpace** | integer | 磁盘大小 | [] | [] | **Yes** |
| **Quantity** | integer | 计费周期数量 | [] | [] | No |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **OSSID** | string | 对象存储ID | [] | [] | No |
| **CPU** | integer |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*OssPriceInfo*](#OssPriceInfo)> |  | [] |

### 数据模型

#### OssPriceInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ChargeType** | string | 计费类型 | [] |
| **OrderType** | string | 订单类型 | [] |
| **Price** | double | 价格 | [] |

## 对象存储降配 - DowngradeOSS

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **OSSID** | string | 对象存储ID | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **CPU** | integer |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|

