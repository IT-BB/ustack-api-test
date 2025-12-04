# FS

## 创建文件存储服务 - CreateFS

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
| **VMType** | string |  | [] | [] | **Yes** |
| **DiskSetType** | string | 存储集群 | [] | [] | **Yes** |
| **DiskSpace** | integer | 存储容量, 最小值100G | [] | [] | **Yes** |
| **VPCID** | string | vpcID | [] | [] | **Yes** |
| **SubnetID** | string | 子网ID | [] | [] | **Yes** |
| **EIPID** | string | EIPID | [] | [] | No |
| **WANSGID** | string | 外网安全组 | [] | [] | No |
| **BackupID** | string | 备份ID | [] | [] | No |
| **TerminationPolicy** | integer | 是否开启删除保护, 0:开启; 1:关闭 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **FSID** | string | 文件存储实例ID | [] |


## 删除文件存储服务 - DeleteFS

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **FSID** | string | 文件存储实例ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 文件存储扩容 - UpgradeFS

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **FSID** | string | FSID | [] | [] | **Yes** |
| **DiskSpace** | integer | 磁盘大小 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取文件存储列表 - DescribeFS

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ProjectIDs** | array<string> | 项目组ID | [] | [] | No |
| **Keyword** | string | 搜索关键词 | [] | [] | No |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **FSIDs** | array<string> | 文件存储实例ID | [] | [] | No |
| **Status** | array<string> | 文件存储状态 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer |  | [] |
| **Infos** | array<[*FSInfo*](#FSInfo)> |  | [] |

### 数据模型

#### FSInfo

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
| **FSID** | string | 文件存储实例ID | [] |
| **VMTypeAlias** | string | 计算集群名称 | [] |
| **VPCName** | string | VPC名称 | [] |
| **SubnetName** | string | 子网名称 | [] |
| **VPCID** | string | vpcID | [] |
| **SubnetID** | string | 子网ID | [] |
| **Status** | string | 状态 | [] |
| **DiskID** | string | 磁盘ID | [] |
| **StorageSetType** | string | 存储集群类型 | [] |
| **DiskSpace** | integer | 存储容量 | [] |
| **Endpoints** | array<[*FsEndpoint*](#FsEndpoint)> | 网络信息 | [] |
| **IOPS** | integer | IOPS | [] |
| **Bandwidth** | integer | 带宽 | [] |
| **StorageSetArch** | string | 存储集群架构 | [] |
| **TerminationPolicy** | integer | 是否开启删除保护 | [] |
| **VMType** | string | 计算集群类型 | [] |
| **StorageSetTypeAlias** | string | 存储集群名称 | [] |
| **SetArch** | string | 计算集群架构 | [] |

#### FsEndpoint

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **IP** | string | IP | [] |
| **IPID** | string | IPID | [] |
| **IPVersion** | string | IP版本 | [] |
| **NicType** | string | 网卡类型 | [] |
| **Endpoint** | string | 连接信息 | [] |
| **SGID** | string | 安全组ID | [] |
| **SGName** | string | 安全组名称 | [] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

## 获取文件存储价格 - GetFSPrice

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ChargeType** | string | 计费类型 | [] | [] | **Yes** |
| **VMType** | string | 计算集群 | [] | [] | **Yes** |
| **DiskSetType** | string | 存储集群 | [] | [] | **Yes** |
| **DiskSpace** | integer | 磁盘大小 | [] | [] | **Yes** |
| **Quantity** | integer | 数量，获取扩容价格时非必填 | [] | [] | No |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **FSID** | string | 文件存储实例ID | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*FsPriceInfo*](#FsPriceInfo)> | 价格信息 | [] |

### 数据模型

#### FsPriceInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ChargeType** | string | 计费类型 | [] |
| **OrderType** | string | 账单类型 | [] |
| **Price** | double | 价格 | [] |

## 获取文件存储目录文件 - DescribeFSFile

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **FSID** | string | 文件存储实例ID | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **FilePath** | string | 文件路径 | [] | [] | **Yes** |
| **ShowHidden** | bool | 是否显示隐藏文件 | [] | [] | No |
| **Limit** | integer | 返回数量 | [] | [] | No |
| **Offset** | integer | 偏移量 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 文件数量 | [] |
| **Infos** | array<[*FileInfo*](#FileInfo)> | 文件信息 | [] |

### 数据模型

#### FileInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string | 文件名 | [] |
| **Size** | integer | 文件大小 | [] |
| **IsDir** | bool | 是否为目录 | [] |
| **ModTime** | integer | 文件修改时间 | [] |
| **FilePath** | string | 文件路径 | [] |

## 删除文件存储目录文件 - DeleteFSFile

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **FSID** | string | 文件存储实例ID | [] | [] | **Yes** |
| **CompanyID** | integer |  | [] | [] | No |
| **FilePath** | string | 文件路径 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 创建文件存储会话 - FSLogin

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **FSID** | string | 文件存储实例ID | [] | [] | **Yes** |
| **CompanyID** | integer |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **SessionID** | string | 会话ID | [] |


## 创建目录 - CreateFSDir

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **FSID** | string | 文件存储实例ID | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **DirPath** | string | 文件路径 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|

