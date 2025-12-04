# MULTICAST

## 创建组播组 - CreateMulticastGroup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ProjectID** | string | 项目ID | [] | [] | No |
| **TagKeyValuePairs** | array<string> | 标签键值对 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **Name** | string | 名称 | [] | [] | **Yes** |
| **Remark** | string | 备注 | [] | [] | No |
| **VPCID** | string | vpcID | [] | [] | **Yes** |
| **MulticastIP** | string | 组播组IP | [] | [] | **Yes** |
| **MulticastPort** | integer | 组播组端口 | [] | [] | **Yes** |
| **MulticastSource** | string | 发送方 虚拟机ID | [] | [] | **Yes** |
| **MulticastDest** | array<string> | 接收方 虚拟机ID列表 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **MulticastGroupID** | string | 组播组ID | [] |


## 删除组播组 - DeleteMulticastGroup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **MulticastGroupID** | string | 组播组ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 更新组播组 - UpdateMulticastGroup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **MulticastGroupID** | string | 组播组ID | [] | [] | **Yes** |
| **MulticastDest** | array<string> | 接收方 虚拟机ID列表 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取组播组列表 - DescribeMulticastGroup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **MulticastGroupIDs** | array<string> | 组播组ID列表 | [] | [] | No |
| **ProjectIDs** | array<string> | 项目组ID列表 | [] | [] | No |
| **VPCID** | string | vpcID | [] | [] | No |
| **Offset** | integer | 分页偏移量 | [] | [] | No |
| **Limit** | integer | 分页大小 | [] | [] | No |
| **Keyword** | string | 关键词 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*MulticastGroupInfo*](#MulticastGroupInfo)> | 详情 | [] |
| **TotalCount** | integer | 总数 | [] |

### 数据模型

#### MulticastGroupInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **MulticastGroupID** | string | 组播组ID | [] |
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
| **Status** | string | 状态 | [] |
| **VPCID** | string | vpcID | [] |
| **VPCName** | string | vpc名称 | [] |
| **MulticastIP** | string | 组播组IP | [] |
| **MulticastPort** | integer | 组播组端口 | [] |
| **MulticastSource** | string | 发送方 虚拟机ID | [] |
| **MulticasrDest** | array<string> | 接收方 虚拟机ID列表 | [] |
| **VPCNetwork** | string | VPC网段 | [] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |
