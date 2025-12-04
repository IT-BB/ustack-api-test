# NATGW

## 创建NAT网关 - CreateNATGW

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **ProjectID** | string | 项目组ID | [] | [] | No |
| **TagKeyValuePairs** | array<string> | 标签键值对 | [] | [] | **Yes** |
| **Remark** | string | 备注 | [] | [] | No |
| **Name** | string | 名称 | [] | [] | **Yes** |
| **ChargeType** | string | 计费类型 | [] | [] | **Yes** |
| **Quantity** | integer | 购买数量 | [] | [] | **Yes** |
| **VPCID** | string | vpcID | [] | [] | **Yes** |
| **SubnetID** | string | 子网ID | [] | [] | **Yes** |
| **SGID** | string | 安全组ID | [] | [] | **Yes** |
| **EIPID** | string | eipID | [] | [] | **Yes** |
| **VMType** | string | 计算集群类型 | [] | [] | **Yes** |
| **DiskSetType** | string | 存储集群类型 | [] | [] | **Yes** |
| **HighAvailability** | string | 高可用 | [] | [] | No |
| **TerminationPolicy** | integer | 是否开启删除保护, 0:开启; 1:关闭 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **NATGWID** | string | 网关ID | [] |


## 删除NAT网关 - DeleteNATGW

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | string | 租户ID | [] | [] | No |
| **NATGWID** | string | 网关ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 添加NAT网关规则 - CreateNATGWRule

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **NATGWID** | string | 网关ID | [] | [] | **Yes** |
| **BindResourceType** | string | SNAT类型 | [] | [] | **Yes** |
| **BindResourceID** | string | 绑定资源ID | [] | [] | **Yes** |
| **EIPID** | string | EIPID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **RuleID** | string | SNATID | [] |


## 修改NAT网关规则 - UpdateNATGWRule

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | string | 租户ID | [] | [] | No |
| **NATGWID** | string | 网关ID | [] | [] | **Yes** |
| **RuleID** | string | SNAT规则ID | [] | [] | **Yes** |
| **EIPID** | string | 外网IP | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 删除NAT网关规则 - DeleteNATGWRule

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | string | 租户ID | [] | [] | No |
| **NATGWID** | string | 网关ID | [] | [] | **Yes** |
| **RuleID** | string | SNATID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 创建端口转发 - CreateNATGWPolicy

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Protocol** | string | DNAT协议 | [] | [] | **Yes** |
| **NATGWID** | string | 网关ID | [] | [] | **Yes** |
| **SrcIP** | string | 源IP | [] | [] | **Yes** |
| **SrcPort** | string | 源端口 | [] | [] | **Yes** |
| **DstResourceID** | string | 目的虚拟机ID | [] | [] | **Yes** |
| **DstPort** | string | 目的虚拟机端口 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **PolicyID** | string | DNAT规则ID | [] |


## 更新端口转发 - UpdateNATGWPolicy

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **PolicyID** | string | DNAT规则ID | [] | [] | **Yes** |
| **Protocol** | string | DNAT协议 | [] | [] | **Yes** |
| **SrcIP** | string | 源IP | [] | [] | **Yes** |
| **SrcPort** | string | 源端口 | [] | [] | **Yes** |
| **DstResourceID** | string | 目的虚拟机ID | [] | [] | **Yes** |
| **DstPort** | string | 目的虚拟机端口 | [] | [] | **Yes** |
| **NATGWID** | string | 网关ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 删除端口转发 - DeleteNATGWPolicy

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **PolicyID** | string | DNAT规则ID | [] | [] | **Yes** |
| **NATGWID** | string | 网关ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 绑定EIP到NAT网关 - BindEIPToNATGW

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ResourceID** | string | 网关ID | [] | [] | **Yes** |
| **EIPID** | string | EIPID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 从NAT网关上解绑EIP - UnbindEIPFromNATGW

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ResourceID** | string | 网关ID | [] | [] | **Yes** |
| **EIPID** | string | EIPID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 修改NAT网关的安全组 - UpdateSGFromNATGW

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ResourceID** | string | 网关ID | [] | [] | **Yes** |
| **SGID** | string | 安全组ID | [] | [] | **Yes** |
| **NICType** | string | 网卡类型 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取NAT网关 - DescribeNATGW

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | No |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Keyword** | string | 搜索关键字 | [] | [] | No |
| **ProjectIDs** | array<string> | 项目组ID | [] | [] | No |
| **NATGWIDs** | array<string> | 网关ID | [] | [] | No |
| **Offset** | integer | 分页页码 | [] | [] | No |
| **Limit** | integer | 分页数量 | [] | [] | No |
| **Status** | array<string> | NATGW状态 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*NATGWInfo*](#NATGWInfo)> | 网关详情 | [] |
| **TotalCount** | integer | 资源总数 | [] |

### 数据模型

#### NATGWEIPInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **IPID** | string | EIPID | [] |
| **IP** | string | IP地址 | [] |
| **IPStatus** | string | IP状态 | [] |

#### NATGWInfo

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
| **DeleteTime** | integer | 删除时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> | 标签 | [] |
| **ChargeType** | string | 支付类型 | [] |
| **ExpireTime** | integer | 过期时间 | [] |
| **EIPInfos** | array<[*NATGWEIPInfo*](#NATGWEIPInfo)> | EIP信息 | [] |
| **VPCID** | string | vpcID | [] |
| **VPCName** | string | VPC名称 | [] |
| **SubnetID** | string | 子网ID | [] |
| **SubnetName** | string | 子网名称 | [] |
| **SGID** | string | 安全组ID | [] |
| **SGName** | string | 安全组名称 | [] |
| **NATGWID** | string | 网关ID | [] |
| **NATGWStatus** | string | 网关状态 | [] |
| **VMType** | string | 计算集群类型 | [] |
| **HighAvailability** | string | 高可用模式 | [] |
| **StorageSetType** | string | 存储集群 | [] |
| **TerminationPolicy** | integer | 是否开启删除保护 | [] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

## 获取NAT网关价格 - GetNATGWPrice

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **VMType** | string | 计算集群类型 | [] | [] | **Yes** |
| **ChargeType** | string | 计费类型 | [] | [] | **Yes** |
| **Quantity** | integer | 购买数量 | [] | [] | **Yes** |
| **HighAvailability** | string | 高可用 | [] | [] | No |
| **NATGWID** | string | NATGWID | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*NatGwPriceInfo*](#NatGwPriceInfo)> | 订单详情 | [] |

### 数据模型

#### NatGwPriceInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Price** | double | 价格 | [] |
| **ChargeType** | string | 支付类型 | [] |

## 查询端口转发 - DescribeNATGWPolicy

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | No |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **NATGWID** | string | 网关ID | [] | [] | No |
| **PolicyIDs** | array<string> | DNAT规则ID | [] | [] | No |
| **Limit** | integer | 分页数量 | [] | [] | No |
| **Offset** | integer | 分页页码 | [] | [] | No |
| **Keyword** | string | 搜索关键字 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*NATGWPolicyInfo*](#NATGWPolicyInfo)> | DNAT详情 | [] |
| **TotalCount** | integer | DNAT总数 | [] |

### 数据模型

#### NATGWPolicyInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **NATGWID** | string | 网关ID | [] |
| **PolicyID** | string | DNAT规则ID | [] |
| **Protocol** | string | DNAT协议 | [] |
| **SrcIP** | string | 源IP | [] |
| **SrcPort** | string | 源端口 | [] |
| **DstIP** | string | 目的IP | [] |
| **DstPort** | string | 目的端口 | [] |
| **DstResourceID** | string | 目的虚拟机ID | [] |
| **PolicyStatus** | string | 规则状态 | [] |
| **State** | string | 规则状态 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |

## 获取NAT网关规则 - DescribeNATGWRule

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | No |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **NATGWID** | string | 网关ID | [] | [] | No |
| **Keyword** | string | 搜索关键字 | [] | [] | No |
| **RuleIDs** | array<string> | SNAT规则ID | [] | [] | No |
| **Limit** | integer | 分页参数 | [] | [] | No |
| **Offset** | integer | 分页参数 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*NATGWRuleInfo*](#NATGWRuleInfo)> | SNAT详情 | [] |
| **TotalCount** | integer | SNAT总数 | [] |

### 数据模型

#### NATGWEIPInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **IPID** | string | EIPID | [] |
| **IP** | string | IP地址 | [] |
| **IPStatus** | string | IP状态 | [] |

#### NATGWRuleInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **RuleID** | string | SNAT规则ID | [] |
| **NATGWID** | string | 网关ID | [] |
| **BindResourceID** | string | 绑定资源ID | [] |
| **BindResourceType** | string | 绑定资源类型 | [] |
| **Address** | string | SNAT代理IP地址 | [] |
| **BindResourceName** | string | 绑定资源名称 | [] |
| **EIP** | string | 外网IP | [] |
| **EIPInfos** | array<[*NATGWEIPInfo*](#NATGWEIPInfo)> | EIP详情 | [] |
| **RuleStatus** | string | SNAT状态 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |

## 升级为高可用版本 - UpgradeNATGWToHA

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **NATGWID** | string | NATGW网关ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|

