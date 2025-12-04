# SG

## 创建安全组 - CreateSecurityGroup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **ProjectID** | string | 项目组ID，用于资源分组管理 | [project-5kyt09jzwcqh6j] | [ListProjects] | No |
| **TagKeyValuePairs** | array<string> | 标签键值对，用于资源标记和分类管理。格式为 key:value 的字符串，传入​​ Base64 编码的字符串。 | [dGVzdA==:dGVzdC0x] | [] | **Yes** |
| **Name** | string | 安全组名称，长度为1-50个字符，名称只能包含中英文、数字、点（.），下划线（_）和中划线（-） | [test-sggroup] | [] | **Yes** |
| **Remark** | string | 备注，用于进行说明和注释，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [这是一个备注] | [] | No |
| **Rule** | array<string> | 规则，标准格式：“协议｜端口号｜地址｜动作(ACCEPT,DROP)｜优先级(HIGH,MEDIUM,LOW)｜方向(出/入：0/1)｜备注” 示例：“TCP|12345|0.0.0.0/0|ACCEPT|HIGH|1|备注二” 可以通过端口组来统一管理协议和端口号，通过 ip 组统一管理 ip。格式：“端口组 id｜-｜ip 组 id｜动作｜优先级｜方向｜备注” 示例：“group:sgportgroup-11wzcaa6b6e3of|-|group:sgipgroup-yj1ujstljnml4o|DROP|MEDIUM|0|test3”  | [group:sgportgroup-11wzcaa6b6e3of|-|group:sgipgroup-yj1ujstljnml4o|DROP|MEDIUM|0|test3 group:sgportgroup-11wzcaa6b6e3of|-|0.0.0.0/0|ACCEPT|LOW|1|test4] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **SGID** | string | 创建的安全组ID | [sg-5w3v5d8ucmoga0] |


## 删除安全组 - DeleteSecurityGroup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **SGID** | string | 安全组ID | [sg-5w3v5d8ucmoga0] | [DescribeSecurityGroup] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取安全组 - DescribeSecurityGroup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **SGIDs** | array<string> | 安全组ID列表，用于查询指定的安全组，多个元素之间用,分隔。 | [sg-z0d7ecehy5o8312] | [DescribeSecurityGroup] | No |
| **ProjectIDs** | array<string> | 项目组ID列表，用于资源分组管理，多个元素之间用,分隔 | [project-5kyt09jzwcqh6j] | [ListProjects] | No |
| **Keyword** | string | 搜索关键词 | [example_key] | [] | No |
| **Status** | string | 状态，查询指定状态的安全组 | [Available] | [] | No |
| **Offset** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 默认值：0 | [0] | [] | No |
| **Limit** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。默认值：10 | [20] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*SGDetailInfo*](#SGDetailInfo)> | 安全组的详情记录 | [] |
| **TotalCount** | integer | 查询到的总记录数 | [] |

### 数据模型

#### SGDetailInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **SGID** | string | 目标安全组ID | [sg-9mdqfo7pvin9wz] |
| **Name** | string | 安全组的名称 | [test] |
| **Remark** | string | 备注，用于进行说明和注释，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [这是一个备注] |
| **Region** | string | 安全组所属的地域 | [region-1] |
| **RegionAlias** | string | 地域别名 | [test] |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] |
| **CompanyName** | string | 安全组所属的租户名称 | [ucloud] |
| **Email** | string | 租户邮箱，租户注册时提供的用于接收通知和联系的邮箱 | [example@ucloud.cn] |
| **ProjectID** | string | 项目组ID，用于资源分组管理 | [project-5kyt09jzwcqh6j] |
| **ProjectName** | string | 项目组名称 | [project01] |
| **Reason** | string | 创建失败原因 | [安全组已存在] |
| **CreateTime** | integer | 创建时间，Unix 时间戳（秒级） | [1754644992] |
| **UpdateTime** | integer | 更新时间，Unix 时间戳（秒级） | [1754644992] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> | 标签 | [cluster-resouce.system:offline-node] |
| **ResourceCount** | integer | 安全组已经绑定的资源数量 | [3] |
| **Rule** | array<[*SGRuleInfo*](#SGRuleInfo)> | 安全组规则详情 | [] |
| **RuleCount** | integer | 安全组规则数量 | [3] |
| **Status** | string | 安全组当前的状态 | [Available] |

#### SGRuleInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **RuleID** | string | 规则ID | [sgrule-4uft68ix19xjpd] |
| **SGID** | string | 目标安全组ID | [sg-9mdqfo7pvin9wz] |
| **ProjectID** | string | 项目组ID，用于资源分组管理 | [project-5kyt09jzwcqh6j] |
| **ProjectName** | string | 项目组名称 | [project01] |
| **Remark** | string | 备注，用于进行说明和注释，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [这是一个备注] |
| **ProtocolType** | string | 安全组协议的类型 | [TCP UDP] |
| **DstPort** | string | 目标端口 | [12345] |
| **SrcIP** | string | 源IP，安全组定义的 IP 地址 | [127.0.0.1] |
| **RuleAction** | string | 规则动作，取值范围：1、ACCEPT 2、DROP | [ACCEPT] |
| **Priority** | string | 优先级，取值范围：1、HIGH 2、MEDIUM 3、LOW | [HIGH] |
| **IsIn** | string | 方向，取值范围：1、0（出） 2、1（入） | [1] |
| **State** | string | 安全规则状态 | [Available] |
| **IPGroupID** | string | IP组ID | [sgipgroup-1egyy1321yy] |
| **IPGroupName** | string | IP组名称 | [test] |
| **IPGroupRules** | string | IP组规则 | [127.0.0.1] |
| **PortGroupID** | string | 端口组ID | [sgportgroup-11wzcaa6b6e3of] |
| **PortGroupName** | string | 端口组名称 | [test] |
| **PortGroupRules** | string | 端口组规则 | [tcp:12346] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

## 绑定安全组 - BindSecurityGroup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **SGID** | string | 安全组ID，指定所要使用的安全组 | [sg-9mdqfo7pvin9wz] | [DescribeSecurityGroup] | **Yes** |
| **ResourceID** | string | 资源ID，指定安全组绑定的资源的 id，例如：虚拟机的 id，网卡 id | [vm-ds7z7z6qz6at9g] | [DescribeVMInstance DescribeLB DescribeNATGW DescribeNIC DescribeOSS DescribeFS DescribeMySQL DescribeRedis] | **Yes** |
| **NICType** | string | 网卡类型，绑定网络的类型：1、内网（LAN） 2、外网（WAN） 3、扁平网络（Flat） | [LAN] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 解绑安全组 - UnBindSecurityGroup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **ResourceID** | string | 资源ID，指定安全组绑定的资源的 id，例如：虚拟机的 id，网卡 id | [vm-ds7z7z6qz6at9g] | [DescribeVMInstance DescribeLB DescribeNATGW DescribeNIC DescribeOSS DescribeFS DescribeMySQL DescribeRedis] | **Yes** |
| **NICType** | string | 网卡类型，绑定网络的类型：1、内网（LAN） 2、外网（WAN） 3、扁平网络（Flat） | [LAN] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 新建安全组规则 - CreateSecurityGroupRule

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **SGID** | string | 目标安全组ID | [sg-9mdqfo7pvin9wz] | [DescribeSecurityGroup] | **Yes** |
| **ProjectID** | string | 项目组ID，用于资源分组管理 | [project-5kyt09jzwcqh6j] | [ListProjects] | No |
| **Rules** | array<string> | 规则，标准格式：“协议｜端口号｜地址｜动作(ACCEPT,DROP)｜优先级(HIGH,MEDIUM,LOW)｜方向(出/入：0/1)｜备注” 示例：“TCP|12345|0.0.0.0/0|ACCEPT|HIGH|1|备注二” 可以通过端口组来统一管理协议和端口号，通过 ip 组统一管理 ip。格式：“端口组 id｜-｜ip 组 id｜动作｜优先级｜方向｜备注” 示例：“group:sgportgroup-11wzcaa6b6e3of|-|group:sgipgroup-yj1ujstljnml4o|DROP|MEDIUM|0|test3”  | [group:sgportgroup-11wzcaa6b6e3of|-|group:sgipgroup-yj1ujstljnml4o|DROP|MEDIUM|0|test3 group:sgportgroup-11wzcaa6b6e3of|-|0.0.0.0/0|ACCEPT|LOW|1|test4] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 删除安全组规则 - DeleteSecurityGroupRule

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **SGID** | string | 目标安全组ID | [sg-9mdqfo7pvin9wz] | [DescribeSecurityGroup] | **Yes** |
| **SGRuleID** | string | 安全组规则ID | [sgrule-bcbx5lheizypgn] | [DescribeSecurityGroupRule] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 更新安全组规则 - UpdateSecurityGroupRule

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **SGID** | string | 目标安全组ID | [sg-9mdqfo7pvin9wz] | [DescribeSecurityGroup] | **Yes** |
| **Remark** | string | 备注，用于进行说明和注释，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [这是一个备注] | [] | No |
| **Rules** | array<string> | 规则，标准格式：“协议｜端口号｜地址｜动作(ACCEPT,DROP)｜优先级(HIGH,MEDIUM,LOW)｜方向(出/入：0/1)｜备注” 示例：“TCP|12345|0.0.0.0/0|ACCEPT|HIGH|1|备注二” 可以通过端口组来统一管理协议和端口号，通过 ip 组统一管理 ip。格式：“端口组 id｜-｜ip 组 id｜动作｜优先级｜方向｜备注” 示例：“group:sgportgroup-11wzcaa6b6e3of|-|group:sgipgroup-yj1ujstljnml4o|DROP|MEDIUM|0|test3”  | [group:sgportgroup-11wzcaa6b6e3of|-|group:sgipgroup-yj1ujstljnml4o|DROP|MEDIUM|0|test3 group:sgportgroup-11wzcaa6b6e3of|-|0.0.0.0/0|ACCEPT|LOW|1|test4] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*SGRuleInfo*](#SGRuleInfo)> | 查询到的详情信息 | [] |
| **TotalCount** | integer | 查询到的总记录数 | [] |

### 数据模型

#### SGRuleInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **RuleID** | string | 规则ID | [sgrule-4uft68ix19xjpd] |
| **SGID** | string | 目标安全组ID | [sg-9mdqfo7pvin9wz] |
| **ProjectID** | string | 项目组ID，用于资源分组管理 | [project-5kyt09jzwcqh6j] |
| **ProjectName** | string | 项目组名称 | [project01] |
| **Remark** | string | 备注，用于进行说明和注释，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [这是一个备注] |
| **ProtocolType** | string | 安全组协议的类型 | [TCP UDP] |
| **DstPort** | string | 目标端口 | [12345] |
| **SrcIP** | string | 源IP，安全组定义的 IP 地址 | [127.0.0.1] |
| **RuleAction** | string | 规则动作，取值范围：1、ACCEPT 2、DROP | [ACCEPT] |
| **Priority** | string | 优先级，取值范围：1、HIGH 2、MEDIUM 3、LOW | [HIGH] |
| **IsIn** | string | 方向，取值范围：1、0（出） 2、1（入） | [1] |
| **State** | string | 安全规则状态 | [Available] |
| **IPGroupID** | string | IP组ID | [sgipgroup-1egyy1321yy] |
| **IPGroupName** | string | IP组名称 | [test] |
| **IPGroupRules** | string | IP组规则 | [127.0.0.1] |
| **PortGroupID** | string | 端口组ID | [sgportgroup-11wzcaa6b6e3of] |
| **PortGroupName** | string | 端口组名称 | [test] |
| **PortGroupRules** | string | 端口组规则 | [tcp:12346] |

## 获取安全组规则 - DescribeSecurityGroupRule

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **SGID** | string | 目标安全组ID | [sg-9mdqfo7pvin9wz] | [DescribeSecurityGroup] | **Yes** |
| **SGRuleIDs** | array<string> | 安全组规则ID列表，用于查询指定的安全组，多个元素之间用,分隔 | [sgrule-4uft68ix19xjpd] | [] | No |
| **Offset** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 默认值：0 | [0] | [] | No |
| **Limit** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。默认值：10 | [20] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*SGRuleInfo*](#SGRuleInfo)> | 查询到的详情信息 | [] |
| **TotalCount** | integer | 查询到的总记录数 | [] |

### 数据模型

#### SGRuleInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **RuleID** | string | 规则ID | [sgrule-4uft68ix19xjpd] |
| **SGID** | string | 目标安全组ID | [sg-9mdqfo7pvin9wz] |
| **ProjectID** | string | 项目组ID，用于资源分组管理 | [project-5kyt09jzwcqh6j] |
| **ProjectName** | string | 项目组名称 | [project01] |
| **Remark** | string | 备注，用于进行说明和注释，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [这是一个备注] |
| **ProtocolType** | string | 安全组协议的类型 | [TCP UDP] |
| **DstPort** | string | 目标端口 | [12345] |
| **SrcIP** | string | 源IP，安全组定义的 IP 地址 | [127.0.0.1] |
| **RuleAction** | string | 规则动作，取值范围：1、ACCEPT 2、DROP | [ACCEPT] |
| **Priority** | string | 优先级，取值范围：1、HIGH 2、MEDIUM 3、LOW | [HIGH] |
| **IsIn** | string | 方向，取值范围：1、0（出） 2、1（入） | [1] |
| **State** | string | 安全规则状态 | [Available] |
| **IPGroupID** | string | IP组ID | [sgipgroup-1egyy1321yy] |
| **IPGroupName** | string | IP组名称 | [test] |
| **IPGroupRules** | string | IP组规则 | [127.0.0.1] |
| **PortGroupID** | string | 端口组ID | [sgportgroup-11wzcaa6b6e3of] |
| **PortGroupName** | string | 端口组名称 | [test] |
| **PortGroupRules** | string | 端口组规则 | [tcp:12346] |

## 获取安全组关联资源 - DescribeSecurityGroupResource

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **SGID** | string | 目标安全组ID | [sg-9mdqfo7pvin9wz] | [DescribeSecurityGroup] | **Yes** |
| **Offset** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 默认值：0 | [0] | [] | No |
| **Limit** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。默认值：10 | [20] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*SGResourceInfo*](#SGResourceInfo)> | 查询到的详情信息 | [] |
| **TotalCount** | integer | 查询到的总记录数 | [] |

### 数据模型

#### SGResourceInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ResourceID** | string | 所绑定的资源ID | [vm-ds7z7z6qz6at9g] |
| **ResourceType** | string | 资源类型，取值范围：1、VM（虚拟机）2、LB（负载均衡）3、NATGW（NAT 网关）4、NIC（网卡）5、STOR（对象存储）6、FILESTOR（文件存储）7、MYSQL（MySQL）8、REDIS（Redis） | [VM] |
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] |
| **Name** | string | 安全组名称 | [test] |
| **SGID** | string | 目标安全组ID | [sg-9mdqfo7pvin9wz] |
| **NICType** | string | 网卡类型，绑定网络的类型：1、内网（LAN） 2、外网（WAN） 3、扁平网络（Flat） | [LAN] |

## 创建 IP 组 - CreateIPGroup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **ProjectID** | string | 项目组ID，用于资源分组管理 | [project-5kyt09jzwcqh6j] | [ListProjects] | No |
| **TagKeyValuePairs** | array<string> | 标签键值对，用于资源标记和分类管理。格式为 key:value 的字符串，传入​​ Base64 编码的字符串。 | [dGVzdA==:dGVzdC0x] | [] | **Yes** |
| **Name** | string | 安全组名称，长度为1-50个字符，名称只能包含中英文、数字、点（.），下划线（_）和中划线（-） | [test-sggroup] | [] | **Yes** |
| **Remark** | string | 备注，用于进行说明和注释，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [这是一个备注] | [] | No |
| **Rules** | string | 规则，格式：ip 地址 1，ip 地址 2。 | [127.0.0.1,0.0.0.0] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **IPGroupID** | string | IP组ID | [sgipgroup-hyma9lr3yv0vcw] |


## 删除 IP 组 - DeleteIPGroup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **IPGroupID** | string | IP组ID，指定被删除的IP组ID | [sgipgroup-hyma9lr3yv0vcw] | [DescribeIPGroup] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 更新 IP 组 - UpdateIPGroup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **IPGroupID** | string | IP组ID | [sgipgroup-cx0oe3chxrspye] | [DescribeIPGroup] | **Yes** |
| **Rules** | string | 新的规则，格式：ip 地址 1，ip 地址 2。示例：127.0.0.1,0.0.0.0 | [127.0.0.1,0.0.0.0] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 查询 IP 组 - DescribeIPGroup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **Offset** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 默认值：0 | [0] | [] | No |
| **Limit** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。默认值：10 | [20] | [] | No |
| **ProjectIDs** | array<string> | 项目组ID列表，用于资源分组管理，多个元素之间用,分隔 | [project-5kyt09jzwcqh6j] | [ListProjects] | No |
| **IPGroupIDs** | array<string> | IP组ID列表，查询指定 ID 的 IP 组信息，多个元素之间用,分隔 | [sgipgroup-cx0oe3chxrspye] | [DescribeIPGroup] | No |
| **Keyword** | string | 搜索关键词 | [example_key] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 查询到的总记录数 | [] |
| **Infos** | array<[*IPGroupInfo*](#IPGroupInfo)> | IP组的信息 | [] |

### 数据模型

#### IPGroupInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] |
| **RegionAlias** | string | 地域别名 | [地域 A] |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] |
| **ProjectID** | string | 项目组ID，用于资源分组管理 | [project-5kyt09jzwcqh6j] |
| **ProjectName** | string | 项目组名称 | [project01] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> | 标签 | [cluster-resouce.system:offline-node] |
| **Email** | string | 租户邮箱，租户注册时提供的用于接收通知和联系的邮箱 | [example@ucloud.cn] |
| **Name** | string | IP 组名称 | [IPgroup-test] |
| **Remark** | string | 备注，用于进行说明和注释，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [这是一个备注] |
| **Reason** | string | 创建失败原因 | [IP组已存在] |
| **IPGroupID** | string | IP组ID | [sgipgroup-cx0oe3chxrspye] |
| **Rules** | string | 所查询到的 IP 组规则 | [127.0.0.1,0.0.0.1] |
| **UsedSGCount** | integer | 已用安全组数量 | [1] |
| **UsedSGIDInfos** | array<[*UsedSGIDInfo*](#UsedSGIDInfo)> | 已用安全组信息 | [] |
| **Status** | string | 当前 IP 组的状态 | [Available] |
| **CreateTime** | integer | 创建时间，Unix 时间戳（秒级） | [1754644992] |
| **UpdateTime** | integer | 更新时间，Unix 时间戳（秒级） | [1754644992] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

#### UsedSGIDInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **SGName** | string | 安全组名称 | [test] |
| **SGID** | string | 目标安全组ID | [sg-9mdqfo7pvin9wz] |

## 创建端口组 - CreatePortGroup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **ProjectID** | string | 项目组ID，用于资源分组管理 | [project-5kyt09jzwcqh6j] | [ListProjects] | No |
| **TagKeyValuePairs** | array<string> | 标签键值对，用于资源标记和分类管理。格式为 key:value 的字符串，传入​​ Base64 编码的字符串。 | [dGVzdA==:dGVzdC0x] | [] | **Yes** |
| **Name** | string | 安全组名称，长度为1-50个字符，名称只能包含中英文、数字、点（.），下划线（_）和中划线（-） | [test-sggroup] | [] | **Yes** |
| **Remark** | string | 备注，用于进行说明和注释，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [这是一个备注] | [] | No |
| **Rules** | string | 端口组规则，格式：（协议：端口）。示例：TCP:12345 | [TCP:12345 UDP:22334] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **PortGroupID** | string | 端口组ID | [sgportgroup-eupz2egtkm79yh] |


## 删除端口组 - DeletePortGroup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **PortGroupID** | string | 端口组ID | [sgportgroup-eupz2egtkm79yh] | [DescribePortGroup] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 更新端口组 - UpdatePortGroup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **PortGroupID** | string | 端口组ID | [sgportgroup-eupz2egtkm79yh] | [DescribePortGroup] | **Yes** |
| **Rules** | string | 端口组规则，格式：（协议：端口）。示例：TCP:12345 | [TCP:12345 UDP:22334] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 查询端口组 - DescribePortGroup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **Offset** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 默认值：0 | [0] | [] | No |
| **Limit** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。默认值：10 | [20] | [] | No |
| **ProjectIDs** | array<string> | 项目组ID列表，用于资源分组管理，多个元素之间用,分隔 | [project-5kyt09jzwcqh6j] | [ListProjects] | No |
| **PortGroupIDs** | array<string> | 端口组ID列表，用于查询指定端口组，多个元素之间用,分隔 | [sgportgroup-eupz2egtkm79yh] | [DescribePortGroup] | No |
| **Keyword** | string | 搜索关键词 | [example_key] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 查询到的总记录数 | [] |
| **Infos** | array<[*PortGroupInfo*](#PortGroupInfo)> | 端口组的详情信息 | [] |

### 数据模型

#### PortGroupInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Region** | string | 地域ID，您可以调用 DescribeRegion 方法，查看最新的地域列表。 | [develop] |
| **RegionAlias** | string | 地域别名 | [地域 A] |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] |
| **ProjectID** | string | 项目组ID，用于资源分组管理 | [project-5kyt09jzwcqh6j] |
| **ProjectName** | string | 项目组名称 | [project01] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> | 标签 | [cluster-resouce.system:offline-node] |
| **Email** | string | 租户邮箱，租户注册时提供的用于接收通知和联系的邮箱 | [example@ucloud.cn] |
| **Name** | string | 端口组名称 | [test] |
| **Remark** | string | 备注，用于进行说明和注释，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [这是一个备注] |
| **Reason** | string | 创建失败原因 | [端口组已存在] |
| **PortGroupID** | string | 端口组ID | [sgportgroup-eupz2egtkm79yh] |
| **Rules** | string | 端口组规则，格式：（协议：端口）。示例：TCP:12345 | [TCP:12345 UDP:22334] |
| **UsedSGCount** | integer | 已用安全组数量 | [1] |
| **UsedSGIDInfos** | array<[*UsedSGIDInfo*](#UsedSGIDInfo)> | 已用安全组ID信息 | [] |
| **Status** | string | 端口组状态 | [Available] |
| **CreateTime** | integer | 创建时间，Unix 时间戳（秒级） | [1754644992] |
| **UpdateTime** | integer | 更新时间，Unix 时间戳（秒级） | [1754644992] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

#### UsedSGIDInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **SGName** | string | 安全组名称 | [test] |
| **SGID** | string | 目标安全组ID | [sg-9mdqfo7pvin9wz] |
