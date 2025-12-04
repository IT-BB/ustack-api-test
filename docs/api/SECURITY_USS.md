# SECURITY_USS

## 获取主机安全列表 - DescribeUss

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **ResourceIDs** | array<string> | 资源ID | [] | [] | No |
| **ProjectIDs** | array<string> | 项目组ID | [] | [] | No |
| **Limit** | integer | 分页大小 | [] | [] | No |
| **Offset** | integer | 分页偏移量 | [] | [] | No |
| **Keyword** | string | 关键词 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数 | [] |
| **Infos** | array<[*UssInfo*](#UssInfo)> | 详情 | [] |

### 数据模型

#### SecurityProductIPInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **IPID** | string | ipID | [] |
| **IP** | string | IP地址 | [] |
| **InterfaceID** | string | 网卡ID | [] |
| **MAC** | string | MAC地址 | [] |
| **NicType** | string | 网卡类型 内网/外网 | [] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

#### UssInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ResourceID** | string | 资源ID | [] |
| **Name** | string | 资源名称 | [] |
| **Remark** | string | 备注 | [] |
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
| **VersionName** | string | 规格 | [] |
| **Key** | string | 规格的key 用于API传值 | [] |
| **Assets** | integer | 虚拟机终端数 | [] |
| **VPCID** | string | vpcID | [] |
| **VPCName** | string | vpc名称 | [] |
| **SubnetID** | string | 子网ID | [] |
| **SubnetName** | string | 子网名称 | [] |
| **VMType** | string | 计算集群 | [] |
| **DiskSetType** | string | 存储集群 | [] |
| **CPU** | integer | CPU核数 | [] |
| **Memory** | integer | 内存大小 | [] |
| **DiskSize** | integer | 系统盘大小 | [] |
| **Status** | string | 状态 | [] |
| **IPInfos** | array<[*SecurityProductIPInfo*](#SecurityProductIPInfo)> | IP信息 | [] |

## 创建主机安全 - CreateUss

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **ProjectID** | string | 项目组ID | [] | [] | No |
| **TagKeyValuePairs** | array<string> | 标签键值对 | [] | [] | **Yes** |
| **Name** | string | 名称 | [] | [] | **Yes** |
| **Remark** | string | 备注 | [] | [] | No |
| **Quantity** | integer | 计费时长 | [] | [] | **Yes** |
| **VersionNameEN** | string | 规格英文 | [] | [] | **Yes** |
| **VMType** | string | 计算集群 | [] | [] | **Yes** |
| **DiskSetType** | string | 存储集群 | [] | [] | **Yes** |
| **VPCID** | string | vpcID | [] | [] | **Yes** |
| **SubnetID** | string | 子网ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ResourceID** | string | 资源ID | [] |


## 删除主机安全 - DeleteUss

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **ResourceID** | string | 资源ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|

