# ISOLATIONGROUP

## 创建隔离组 - CreateIsolationGroup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **ProjectID** | string | 项目组ID | [] | [] | No |
| **TagKeyValuePairs** | array<string> | 标签键值对 | [] | [] | **Yes** |
| **Name** | string | 隔离组名字 | [] | [] | **Yes** |
| **Remark** | string | 备注 | [] | [] | No |
| **SetID** | string | 计算集群 | [] | [] | **Yes** |
| **PolicyType** | string | 隔离组策略类型 可以是隔离组可以是单个节点 | [] | [] | **Yes** |
| **PolicyObjType** | string | 隔离组策略对象类型 具体的隔离组或者单个节点的名字 | [] | [] | **Yes** |
| **PolicyObj** | string | 隔离组策略对象 | [] | [] | No |
| **IsForce** | bool | 是否强制执行 | [] | [] | No |
| **IsEnable** | bool | 是否启用 | [] | [] | No |
| **PolicyObjIsSelf** | bool | 创建的时候PolicyObj是否为自身 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **IGID** | string | 隔离组ID | [] |


## 获取隔离组信息 - DescribeIsolationGroups

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **ProjectIDs** | array<string> | 项目组ID | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **SetID** | string | 通过计算集群ID筛选 | [] | [] | No |
| **SetType** | string | 通过计算集群类型筛选 | [] | [] | No |
| **PolicyType** | string | 通过策略筛选 | [] | [] | No |
| **PolicyObjType** | string | 通过策略对象类型筛选 | [] | [] | No |
| **PolicyObj** | string | 通过策略对象筛选 | [] | [] | No |
| **IGIDs** | array<string> | 通过隔离组ID筛选 | [] | [] | No |
| **Keyword** | string | 关键词 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer |  | [] |
| **Infos** | array<[*IG*](#IG)> | 详情 | [] |

### 数据模型

#### IG

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
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> | 标签 | [] |
| **IGID** | string | 隔离组ID | [] |
| **SetID** | string | 计算集群ID | [] |
| **SetType** | string | 计算集群类型 | [] |
| **PolicyType** | string | 策略类型 | [] |
| **PolicyObjType** | string | 策略对象类型 | [] |
| **PolicyObj** | array<string> | 策略对象 | [] |
| **VMIDs** | array<string> | 虚拟机ID | [] |
| **IsEnable** | bool | 是否启用 | [] |
| **IsForce** | bool | 是否强制执行 | [] |
| **Status** | string | 状态 | [] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

## 删除隔离组 - DeleteIsolationGroup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **IGID** | string | 隔离组ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 更新隔离组 - UpdateIsolationGroup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **IGID** | string | 隔离组ID | [] | [] | **Yes** |
| **PolicyType** | string | 策略类型 | [] | [] | No |
| **PolicyObj** | array<string> | 策略对象 | [] | [] | No |
| **IsForce** | bool | 是否强制执行 | [] | [] | No |
| **IsEnable** | bool | 是否启用 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 添加VM到隔离组 - AddVMToIsolationGroup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **IGID** | string | 隔离组ID | [] | [] | **Yes** |
| **VMID** | string | 虚拟机ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 从隔离组移除VM - RemoveVMFromIsolationGroup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **IGID** | string | 隔离组ID | [] | [] | **Yes** |
| **VMID** | string | 虚拟机ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 添加节点到隔离组 - AddNodesToIsolationGroup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **IGID** | string | 隔离组ID | [] | [] | **Yes** |
| **NodeIDs** | array<string> | 节点ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 从隔离组移除节点 - RemoveNodesFromIsolationGroup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **IGID** | string | 隔离组ID | [] | [] | **Yes** |
| **NodeIDs** | array<string> | 节点ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|

