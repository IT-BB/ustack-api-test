# FLATIP

## 申请扁平网络IP - AllocateFlatIP

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **ProjectID** | string | 项目ID | [] | [] | No |
| **TagKeyValuePairs** | array<string> | 标签键值对 | [] | [] | **Yes** |
| **Name** | string | 名称 | [] | [] | **Yes** |
| **Remark** | string | 备注 | [] | [] | No |
| **FlatNetworkID** | string | 扁平网络 | [] | [] | **Yes** |
| **IP** | string | IP地址 | [] | [] | No |
| **IPVersion** | string | IP版本 | [] | [] | **Yes** |
| **Bandwidth** | integer | 带宽, 0 表示不限速 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **FlatIPID** | string | FlatIPID | [] |


## 释放扁平网络IP - ReleaseFlatIP

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **FlatIPID** | string | FlatIPID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取扁平网络IP - DescribeFlatIP

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **FlatIPIDs** | array<string> | FlatIPID | [] | [] | No |
| **ProjectIDs** | array<string> | 项目组ID | [] | [] | No |
| **Keyword** | string |  | [] | [] | No |
| **BindResourceID** | string | 绑定的资源ID | [] | [] | No |
| **IPStatus** | string | IP状态 | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*FlatIPInfo*](#FlatIPInfo)> |  | [] |
| **TotalCount** | integer |  | [] |

### 数据模型

#### FlatIPInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string | 名称 | [] |
| **Remark** | string | 备注 | [] |
| **Region** | string | 地域 | [] |
| **RegionAlias** | string | 地域名称 | [] |
| **CompanyID** | integer | 租户ID | [] |
| **CompanyName** | string | 租户名称 | [] |
| **Email** | string | 租户邮箱 | [] |
| **ProjectID** | string | 项目ID | [] |
| **ProjectName** | string | 项目名称 | [] |
| **Reason** | string | 失败原因 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> |  | [] |
| **FlatIPID** | string | FlatIPID | [] |
| **IP** | string | IP | [] |
| **Mode** | string | 模式 | [] |
| **IPVersion** | string | IP版本 | [] |
| **FlatNetworkID** | string | 扁平网络ID | [] |
| **BindResourceID** | string | 绑定的资源ID | [] |
| **BindResourceName** | string | 绑定的资源名称 | [] |
| **BindResourceType** | string | 绑定的资源类型 | [] |
| **BindResourceProjectID** | string | 绑定的项目组ID | [] |
| **BindTime** | integer | 绑定时间 | [] |
| **Status** | string | 状态 | [] |
| **ISDefaultGW** | integer | 是否是默认网关 | [] |
| **CanDefaultGW** | integer | 是否可以作为默认网关 | [] |
| **IsElastic** | string | 是否弹性 | [] |
| **Bandwidth** | integer | 带宽 | [] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

## 绑定扁平网络IP - BindFlatIP

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **FlatIPID** | string | FlatIPID | [] | [] | **Yes** |
| **ResourceType** | string | 资源类型 | [] | [] | **Yes** |
| **ResourceID** | string | 资源ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 解绑扁平网络IP - UnBindFlatIP

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **FlatIPID** | string | FlatIPID | [] | [] | **Yes** |
| **ResourceType** | string | 资源类型 | [] | [] | **Yes** |
| **ResourceID** | string | 资源ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 更换扁平网络 - ReplaceFlatNetwork

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **FlatNetworkID** | string | 扁平网络ID | [] | [] | **Yes** |
| **IP** | string | 指定IP地址 | [] | [] | No |
| **ResourceType** | string | 资源类型 | [] | [] | **Yes** |
| **ResourceID** | string | 资源ID | [] | [] | **Yes** |
| **Bandwidth** | integer | 带宽 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **FlatIPID** | string | flatIPID | [] |


## 更新扁平网络IP带宽 - UpdateFlatIPBandwidth

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **FlatIPID** | string | FlatIPID | [] | [] | **Yes** |
| **Bandwidth** | integer | 带宽, 0 表示不限速 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|

