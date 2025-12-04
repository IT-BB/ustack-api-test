# FLATNETWORK

## 创建扁平网络 - CreateFlatNetwork

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **Name** | string | 名称 | [] | [] | **Yes** |
| **Remark** | string | 备注 | [] | [] | No |
| **TagKeyValuePairs** | array<string> | 标签键值对 | [] | [] | **Yes** |
| **Permission** | string | 权限 | [] | [] | **Yes** |
| **Device** | string | 网卡名称 | [] | [] | **Yes** |
| **Vlan** | string | Vlan | [] | [] | No |
| **CIDR** | string | 网段 | [] | [] | No |
| **IPRange** | string | 可用IP范围 | [] | [] | No |
| **EnableDHCP** | bool | 是否开启DHCP | [] | [] | No |
| **DHCPServerIP** | string | dhcp服务IP | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **FlatNetworkID** | string | 扁平网络ID | [] |


## 删除扁平网络 - DeleteFlatNetwork

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **FlatNetworkID** | string | 扁平网络ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 查询扁平网络 - DescribeFlatNetwork

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **Keyword** | string | 关键词 | [] | [] | No |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **FlatNetworkIDs** | array<string> | 扁平网络ID列表 | [] | [] | No |
| **IPVersion** | string | 网络协议 | [] | [] | No |
| **Offset** | integer | 分页偏移量 | [] | [] | No |
| **Limit** | integer | 分页大小 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*FlatNetwork*](#FlatNetwork)> | 详情 | [] |
| **TotalCount** | integer | 总数 | [] |

### 数据模型

#### FlatNetwork

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Region** | string | 地域 | [] |
| **Name** | string | 名称 | [] |
| **Description** | string | 描述 | [] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> | 标签 | [] |
| **Permission** | string | 权限 | [] |
| **FlatNetworkID** | string | 扁平网络ID | [] |
| **IPVersion** | string | IP版本 | [] |
| **CIDR** | string | 网段 | [] |
| **IPRanges** | string | IP范围 | [] |
| **UsedCount** | integer | 已用数量 | [] |
| **AvailableCount** | integer | 可用数量 | [] |
| **Device** | string | 网卡 | [] |
| **Vlan** | string | Vlan | [] |
| **Platform** | bool | 平台使用 | [] |
| **Status** | string | 状态 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |
| **EnableDHCP** | bool | 是否开启DHCP | [] |
| **DHCPServerIP** | string | dhcp服务IP | [] |
| **Routing** | string | 路由 | [] |
| **RegionAlias** | string | 地域别名 | [] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

## 更新扁平网络 - UpdateFlatNetwork

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **FlatNetworkID** | string | 线路ID | [] | [] | **Yes** |
| **IPRange** | string | 可用IP范围 | [] | [] | No |
| **Device** | string | 网卡名称 | [] | [] | No |
| **Vlan** | string | Vlan | [] | [] | No |
| **DHCPServerIP** | string | DHCP服务IP，传值表示启用DHCP，传空表示禁用DHCP | [] | [] | No |
| **UpdateMode** | string | 更新模式，传IPRange只会更新IPRange字段；传Device只会更新Device、Vlan字段；传DHCPServerIP只会更新DHCPServerIP字段；传All全部更新；传空只会更新IPRange，这是为了向前兼容 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 创建扁平网络路由 - CreateFlatNetworkRoute

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **Remark** | string | 备注 | [] | [] | No |
| **FlatNetworkID** | string | 扁平网络ID | [] | [] | **Yes** |
| **Destination** | string | 目的地址 | [] | [] | **Yes** |
| **NextHop** | string | 下一跳 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 更新扁平网络路由 - UpdateFlatNetworkRoute

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **Remark** | string | 备注 | [] | [] | No |
| **FlatNetworkID** | string | 扁平网络ID | [] | [] | **Yes** |
| **Destination** | string | 目的地址 | [] | [] | **Yes** |
| **NextHop** | string | 下一跳 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 删除扁平网络路由 - DeleteFlatNetworkRoute

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **FlatNetworkID** | string | 扁平网络ID | [] | [] | **Yes** |
| **Destination** | string | 目的地址 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 查询扁平网络路由 - DescribeFlatNetworkRoute

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **FlatNetworkID** | string | 扁平网络ID | [] | [] | **Yes** |
| **Offset** | integer | 分页偏移量 | [] | [] | No |
| **Limit** | integer | 分页大小 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数 | [] |
| **Infos** | array<[*FlatNetworkRoute*](#FlatNetworkRoute)> | 详情 | [] |

### 数据模型

#### FlatNetworkRoute

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **FlatNetworkID** | string | 扁平网络ID | [] |
| **FlatNetworkName** | string | 扁平网络名称 | [] |
| **Destination** | string | 目的地址 | [] |
| **NextHop** | string | 下一跳 | [] |
| **NextHopName** | string | 下一跳名称 | [] |
| **NextHopAddr** | string | 下一跳地址 | [] |
| **NextHopType** | string | 下一跳类型 | [] |
| **Remark** | string | 备注 | [] |
