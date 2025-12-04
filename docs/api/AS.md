# AS

## 创建伸缩组 - CreateASGroup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ProjectID** | string | 项目组ID | [] | [] | No |
| **TagKeyValuePairs** | array<string> | 标签键值对 | [] | [] | **Yes** |
| **Name** | string | 名称 | [] | [] | **Yes** |
| **Remark** | string | 备注 | [] | [] | No |
| **AsType** | string | 伸缩类型, 枚举值：[虚拟机类型VM, 监听器类型VS, EIP类型EIP] | [] | [] | **Yes** |
| **LBID** | string | 负载均衡ID, 只有当伸缩类型是“VS”时才必填 | [] | [] | No |
| **VSID** | string | 监听器ID, 只有当伸缩类型是”VS“时才必填 | [] | [] | No |
| **Port** | integer | 监听器端口, 只有当伸缩类型是”VS“时才必填 | [] | [] | No |
| **Weight** | integer | 负载均衡权重, 只有当伸缩类型是”VS“时才必填 | [] | [] | No |
| **MinInstance** | integer | 最小成员数量, 只有当伸缩模式是”水平伸缩“时才必填 | [] | [] | No |
| **MaxInstance** | integer | 最大成员数量, 只有当伸缩模式是”水平伸缩“时才必填 | [] | [] | No |
| **AllowDecrease** | bool | 是否允许缩容, 只有当伸缩模式是”水平伸缩“时才必填 | [] | [] | No |
| **TemplateID** | string | 虚拟机模板ID, 只有当伸缩模式是”水平伸缩“时才必填 | [] | [] | No |
| **WarmTime** | integer | 预热时间, 只有当伸缩模式是”水平伸缩“时才必填 | [] | [] | No |
| **AsMetric** | string | 伸缩指标 | [] | [] | **Yes** |
| **Threshold** | string | 伸缩阈值 | [] | [] | **Yes** |
| **AsMode** | string | 伸缩模式, 枚举值：[水平伸缩Horizontal, 垂直伸缩Vertical] | [] | [] | **Yes** |
| **MaxSpecification** | string | 最大规格, 只有当伸缩模式是”垂直伸缩“时才必填 | [] | [] | No |
| **VerticalResourceID** | string | 垂直伸缩资源ID, 只有当伸缩模式是”垂直伸缩“时才必填 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **GroupID** | string | 伸缩组ID | [] |


## 查询伸缩组 - DescribeASGroup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ProjectIDs** | array<string> | 项目组ID | [] | [] | No |
| **GroupIDs** | array<string> | 伸缩组ID | [] | [] | No |
| **Limit** | integer | 分页参数 | [] | [] | No |
| **Offset** | integer | 分页参数 | [] | [] | No |
| **Keyword** | string | 查询关键字 | [] | [] | No |
| **AsMode** | string | 伸缩模式, 枚举值：[水平伸缩Horizontal, 垂直伸缩Vertical] | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数 | [] |
| **Infos** | array<[*ASGroupInfo*](#ASGroupInfo)> | 伸缩组详情 | [] |

### 数据模型

#### ASGroupInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string | 名称 | [] |
| **Remark** | string | 备注 | [] |
| **Region** | string | 地域 | [] |
| **RegionAlias** | string | 地域别名 | [] |
| **CompanyID** | integer | 租户ID | [] |
| **CompanyName** | string | 租户名称 | [] |
| **Email** | string | 邮箱 | [] |
| **ProjectID** | string | 项目组ID | [] |
| **ProjectName** | string | 项目组名称 | [] |
| **Reason** | string | 原因 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> | 标签 | [] |
| **GroupID** | string | 伸缩组ID | [] |
| **AsMode** | string | 伸缩模式 | [] |
| **AsType** | string | 伸缩类型 | [] |
| **LBID** | string | 负载均衡ID | [] |
| **VSID** | string | 监听器ID | [] |
| **Port** | integer | 监听器端口, 只有当伸缩类型是”VS“时才必填 | [] |
| **Weight** | integer | 负载均衡权重, 只有当伸缩类型是”VS“时才必填 | [] |
| **MinInstance** | integer | 最小实例数 | [] |
| **MaxInstance** | integer | 最大实例数 | [] |
| **AllowDecrease** | bool | 是否允许缩容 | [] |
| **TemplateID** | string | 虚拟机模板ID | [] |
| **WarmTime** | integer | 预热时间 | [] |
| **MaxSpecification** | string | 最大规格 | [] |
| **VerticalResourceID** | string | 垂直伸缩对应资源ID | [] |
| **AsMetric** | string | 伸缩指标 | [] |
| **Threshold** | string | 指标阈值 | [] |
| **Mode** | string | 禁用启用模式 | [] |
| **InstanceInfos** | array<[*ASInstanceInfo*](#ASInstanceInfo)> | 伸缩成员 | [] |
| **LBInfos** | array<[*LoadBalancerInfo*](#LoadBalancerInfo)> | 关联LB信息 | [] |
| **Status** | string | 状态 | [] |

#### ASInstanceIPInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **IP** | string | 成员IP地址 | [] |
| **Type** | string | 成员类型 | [] |

#### ASInstanceInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ID** | string | 伸缩成员ID | [] |
| **Name** | string | 名称 | [] |
| **TemplateID** | string | 虚拟机模板ID | [] |
| **CreateTime** | integer |  | [] |
| **VMState** | string | 虚拟机状态 当伸缩组类型是VM类型时，VMState为虚拟机的状态，RSState为空 | [] |
| **RSState** | string | RS状态 当伸缩组类型是VS类型时，RSState为RS的状态，VMState为空 | [] |
| **IPInfos** | array<[*ASInstanceIPInfo*](#ASInstanceIPInfo)> | 伸缩成员IP信息 | [] |
| **Origin** | string | 来源 | [] |
| **TemplateName** | string | 虚拟机模板名称 | [] |

#### LoadBalancerInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string | 名称 | [] |
| **LBID** | string | 负载均衡ID | [] |
| **VServerID** | string | 监听器ID | [] |
| **Port** | integer | 监听端口 | [] |
| **Weight** | integer | 负载均衡权重 | [] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

## 删除伸缩组 - DeleteASGroup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | No |
| **GroupID** | string | 伸缩组ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 禁用伸缩组 - DisableASGroup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | No |
| **GroupID** | string | 伸缩组ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 启用伸缩组 - EnableASGroup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | No |
| **GroupID** | string | 伸缩组ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 更新伸缩组 - UpdateASGroup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **GroupID** | string | 伸缩组ID | [] | [] | **Yes** |
| **Port** | integer | 监听器端口, 只有当伸缩类型是”VS“时才必填 | [] | [] | No |
| **Weight** | integer | 负载均衡权重, 只有当伸缩类型是”VS“时才必填 | [] | [] | No |
| **MinInstance** | integer | 最小成员数量, 只有当伸缩模式是”水平伸缩“时才必填 | [] | [] | No |
| **MaxInstance** | integer | 最大成员数量, 只有当伸缩模式是”水平伸缩“时才必填 | [] | [] | No |
| **AllowDecrease** | bool | 是否允许缩容, 只有当伸缩模式是”水平伸缩“时才必填 | [] | [] | No |
| **TemplateID** | string | 虚拟机模板ID, 只有当伸缩模式是”水平伸缩“时才必填 | [] | [] | No |
| **WarmTime** | integer | 预热时间, 只有当伸缩模式是”水平伸缩“时才必填 | [] | [] | No |
| **AsMetric** | string | 伸缩指标 | [] | [] | **Yes** |
| **Threshold** | string | 伸缩阈值 | [] | [] | **Yes** |
| **MaxSpecification** | string | 最大规格, 只有当伸缩模式是”垂直伸缩“时才必填 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 伸缩组关联lb - AttachLoadBalancer

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **Region** | string | 地域 | [] | [] | No |
| **ProjectID** | string | 项目组ID | [] | [] | No |
| **ASGroupID** | string | 伸缩组ID | [] | [] | **Yes** |
| **LBID** | string | 负载均衡ID | [] | [] | **Yes** |
| **VServerID** | string | 监听器ID | [] | [] | **Yes** |
| **Weight** | integer | 负载均衡权重 | [] | [] | No |
| **Port** | integer | 负载均衡端口 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 伸缩组解关联lb - DetachLoadBalancer

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **ASGroupID** | string | 伸缩组ID | [] | [] | **Yes** |
| **LBID** | string | 负载均衡ID | [] | [] | **Yes** |
| **VServerID** | string | 监听器ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 添加伸缩成员 - AddASMember

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **GroupID** | string | 伸缩组ID | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **ResourceID** | string | 资源ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 移除伸缩成员 - RemoveASMember

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **GroupID** | string | 伸缩组ID | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **ResourceID** | string | 资源ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|

