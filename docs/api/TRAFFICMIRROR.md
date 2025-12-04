# TRAFFICMIRROR

## 创建流量镜像 - CreateTrafficMirror

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **Name** | string | 名称 | [] | [] | **Yes** |
| **Remark** | string | 描述 | [] | [] | No |
| **MaxLength** | integer | 报文截取最大长度 | [] | [] | No |
| **Sources** | array<string> | 源设备信息 | [] | [] | **Yes** |
| **Destination** | string | 目的设备信息 | [] | [] | **Yes** |
| **Enable** | bool | 是否启用 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TrafficMirrorID** | string | 流量镜像ID | [] |


## 删除流量镜像 - DeleteTrafficMirror

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **TrafficMirrorID** | string | 流量镜像ID | [] | [] | **Yes** |
| **Region** | string | 地域 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 更新流量镜像 - UpdateTrafficMirror

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **TrafficMirrorID** | string | 流量镜像ID | [] | [] | **Yes** |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **MaxLength** | integer | 报文截取最大长度 | [] | [] | No |
| **Destination** | string | 目的设备信息 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 更新流量镜像源设备信息 - UpdateTrafficMirrorSources

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **TrafficMirrorID** | string | 流量镜像ID | [] | [] | **Yes** |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **Sources** | array<string> | 源设备信息 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 是否启用流量镜像 - UpdateTrafficMirrorEnable

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **TrafficMirrorID** | string | 流量镜像ID | [] | [] | **Yes** |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **Enable** | bool | 是否启用 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 更新流量镜像规则 - UpdateTrafficMirrorRule

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **TrafficMirrorID** | string | 流量镜像ID | [] | [] | **Yes** |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **Rules** | array<string> | 规则信息 | [] | [] | No |
| **Direction** | string | 方向 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 查询流量镜像 - DescribeTrafficMirror

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **TrafficMirrorIDs** | array<string> | 流量镜像ID | [] | [] | No |
| **Keyword** | string |  | [] | [] | No |
| **Status** | string | 状态 | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数 | [] |
| **Infos** | array<[*TrafficMirrorInfo*](#TrafficMirrorInfo)> |  | [] |

### 数据模型

#### TrafficMirrorDestination

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **DstType** | string | 目标类型 | [] |
| **DstDevice** | string | 目标设备 | [] |
| **DstDeviceName** | string | 目标设备名称 | [] |
| **DstLinkName** | string | 目标链路名称 | [] |
| **MAC** | string | mac 地址 | [] |

#### TrafficMirrorInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TrafficMirrorID** | string | 流量镜像ID | [] |
| **Name** | string | 流量镜像名称 | [] |
| **Remark** | string | 流量镜像描述 | [] |
| **Region** | string | 地域 | [] |
| **RegionAlias** | string | 地域别名 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **Enable** | bool | 是否启用 | [] |
| **MaxLength** | integer | 报文截取最大长度 | [] |
| **Destination** | object | 目的设备信息 | [] |
| **SourceCount** | integer | 源设备数量 | [] |
| **IngressRuleCount** | integer | 入方向规则组数量 | [] |
| **EgressRuleCount** | integer | 出方向规则组数量 | [] |
| **Status** | string | 状态 | [] |
| **Reason** | string | 原因 | [] |
| **Sources** | array<[*TrafficMirrorSource*](#TrafficMirrorSource)> | 源设备信息 | [] |
| **IngressRules** | array<[*TrafficMirrorRule*](#TrafficMirrorRule)> | 入方向规则组 | [] |
| **EgressRules** | array<[*TrafficMirrorRule*](#TrafficMirrorRule)> | 出方向规则组 | [] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> | 标签 | [] |

#### TrafficMirrorRule

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Action** | string | 规则动作 | [] |
| **ProtocolType** | string | 协议类型 | [] |
| **SrcNetwork** | string | 网络 | [] |
| **SrcPortRange** | string | 端口 | [] |
| **DstNetwork** | string | 网络 | [] |
| **DstPortRange** | string | 端口 | [] |
| **Priority** | integer | 优先级 | [] |

#### TrafficMirrorSource

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **SrcDevice** | string | 源设备 | [] |
| **SrcDeviceName** | string | 源设备名称 | [] |
| **SrcDeviceStatus** | string | 源设备状态 | [] |
| **Direction** | string | 方向 | [] |
| **AssociatedResourceID** | string | 关联的资源ID | [] |
| **AssociatedResourceName** | string | 关联的资源名称 | [] |
| **AssociatedResourceStatus** | string | 关联的资源状态 | [] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

## 查询流量镜像源设备信息 - DescribeTrafficMirrorSources

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **Keyword** | string |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数 | [] |
| **Infos** | array<[*TrafficMirrorSourceInfo*](#TrafficMirrorSourceInfo)> |  | [] |

### 数据模型

#### TrafficMirrorSimpleSource

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **SrcDevice** | string | 源设备 | [] |
| **SrcDeviceName** | string | 源设备名称 | [] |
| **Direction** | string | 方向 | [] |
| **Available** | bool | 是否可用 | [] |

#### TrafficMirrorSourceInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **AssociatedResourceID** | string | 关联的资源ID | [] |
| **AssociatedResourceName** | string | 关联的资源名称 | [] |
| **Sources** | array<[*TrafficMirrorSimpleSource*](#TrafficMirrorSimpleSource)> | 源设备信息 | [] |
