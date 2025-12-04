# VIP

## 申请VIP - AllocateVIP

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ProjectID** | string | 项目ID | [] | [] | No |
| **TagKeyValuePairs** | array<string> | 标签键值对 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **Name** | string | 名称 | [] | [] | **Yes** |
| **Remark** | string | 备注 | [] | [] | No |
| **ChargeType** | string | 计费类型，带宽不为空时必填 | [] | [] | No |
| **Quantity** | integer | 计费数量 | [] | [] | **Yes** |
| **VPCID** | string | vpcID | [] | [] | **Yes** |
| **SubnetID** | string | 子网ID | [] | [] | **Yes** |
| **IP** | string | IP地址 | [] | [] | No |
| **AssociatedResourceIDs** | array<string> | 关联资源ID | [] | [] | **Yes** |
| **AssociatedResourceType** | string | 关联资源类型 | [] | [] | No |
| **VIPType** | string | VIP类型 LAN/WAN | [] | [] | **Yes** |
| **Bandwidth** | integer | 带宽 | [] | [] | No |
| **SegmentID** | string | 网段，带宽不为空时必填 | [] | [] | No |
| **IPVersion** | string | IP协议，空值为所选VPC/外网默认网络的网络协议, 枚举支持 IPv4 IPv6 ALL和空值 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **VIPID** | string | vipID | [] |


## 释放VIP - ReleaseVIP

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **VIPID** | string | vipID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 更新VIP绑定资源 - UpdateVIPBindResource

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **VIPID** | string | vipID | [] | [] | **Yes** |
| **AssociatedResourceIDs** | array<string> | 关联资源ID | [] | [] | **Yes** |
| **AssociatedResourceType** | string | 关联资源类型 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取VIP列表 - DescribeVIP

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **ProjectIDs** | array<string> | 项目组ID | [] | [] | No |
| **VIPIDs** | array<string> | vipID | [] | [] | No |
| **VIPType** | string | VIP类型 LAN/WAN | [] | [] | No |
| **Keyword** | string | 关键词 | [] | [] | No |
| **Limit** | integer | 分页大小 | [] | [] | No |
| **Offset** | integer | 分页偏移量 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*VIPInfo*](#VIPInfo)> | vip信息 | [] |
| **TotalCount** | integer | 总数 | [] |

### 数据模型

#### AssociatedResources

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ResourceID** | string |  | [] |
| **ResourceName** | string |  | [] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

#### VIPInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **VIPID** | string | vipID | [] |
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
| **Status** | string | 状态 | [] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> | 标签 | [] |
| **ChargeType** | string | 计费类型 | [] |
| **ExpireTime** | integer | 过期时间 | [] |
| **VPCID** | string | vpcID | [] |
| **VPCName** | string | vpc名称 | [] |
| **SubnetID** | string | 子网ID | [] |
| **SubnetName** | string | 子网名称 | [] |
| **IP** | string | IP地址 | [] |
| **AssociatedResources** | array<[*AssociatedResources*](#AssociatedResources)> | 关联资源ID | [] |
| **AssociatedResourceType** | string | 关联资源类型 | [] |
| **VIPType** | string | vip类型 | [] |
| **SegmentID** | string | 外网线路ID | [] |
| **Bandwidth** | integer | 带宽 | [] |
| **SegmentName** | string | 外网线路名称 | [] |
| **SegmentNetwork** | string | 外网网段 | [] |
| **VPCNetwork** | string | VPC网段 | [] |
| **SubnetNetwork** | string | 子网网段 | [] |

## 修改外网VIP的带宽 - UpdateVIPBandwidth

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **VIPID** | string | 外网vipID | [] | [] | **Yes** |
| **Bandwidth** | integer | 带宽 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取外网VIP差价 - GetVIPDiffPrice

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **VIPID** | string | vipID | [] | [] | **Yes** |
| **Bandwidth** | integer | 带宽 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*VipPriceInfo*](#VipPriceInfo)> | vip价格信息 | [] |

### 数据模型

#### VipPriceInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Price** | double | 价格 | [] |
| **ChargeType** | string | 计费类型 | [] |
| **OrderType** | string | 订单类型 | [] |

## 获取外网VIP价格 - GetVIPPrice

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **ChargeType** | string | 计费类型 | [] | [] | **Yes** |
| **Quantity** | integer | 计费数量 | [] | [] | **Yes** |
| **SegmentID** | string | 网段 | [] | [] | **Yes** |
| **Bandwidth** | integer | 带宽 | [] | [] | **Yes** |
| **Count** | integer | 数量 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*VipPriceInfo*](#VipPriceInfo)> | vip价格信息 | [] |

### 数据模型

#### VipPriceInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Price** | double | 价格 | [] |
| **ChargeType** | string | 计费类型 | [] |
| **OrderType** | string | 订单类型 | [] |
