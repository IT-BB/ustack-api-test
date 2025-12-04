# VPC

## 创建VPC - CreateVPC

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | **Yes** |
| **ProjectID** | string | 项目组ID，用于资源分组管理 | [project-5kyt09jzwcqh6j] | [ListProjects] | No |
| **TagKeyValuePairs** | array<string> | 标签键值对，用于资源标记和分类管理。格式为 key:value 的字符串，传入​​ Base64 编码的字符串。 | [dGVzdA==:dGVzdC0x] | [] | **No** |
| **Name** | string | VPC名称，长度为1-50个字符，名称只能包含中英文、数字、点（.），下划线（_）和中划线（-） | [vpc-test] | [] | **Yes** |
| **Remark** | string | 备注，用于进行说明和注释，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [这是一个备注] | [] | No |
| **Network** | string | VPC默认网段, 目前仅支持 IPv4 cidr | [] | [] | **Yes** |
| **ExpandNetwork** | string | VPC拓展网段, 目前仅支持 IPv6 cidr | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **VPCID** | string | 对应创建的vpcID | [vpc-awh9qy2g5f04gg] |


## 添加 VPC 网络 - AddVPCNetwork

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **VPCID** | string | vpcID | [] | [] | **Yes** |
| **Network** | string | 需要加载的 vpc cidr | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 移除 VPC 网络 - RemoveVPCNetwork

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **VPCID** | string | vpcID | [] | [] | **Yes** |
| **Network** | string | 需要移除的 vpc cidr | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 删除VPC - DeleteVPC

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | **Yes** |
| **VPCID** | string | VPC ID | [vpc-awh9qy2g5f04gg] | [DescribeVPC] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取VPC信息 - DescribeVPC

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **VPCIDs** | array<string> | VPC组ID列表，查询指定 ID 的VPC信息，多个元素之间用,分隔。 | [vpc-lj1nck407g85o7] | [DescribeVPC] | No |
| **ProjectIDs** | array<string> | 项目组ID列表，用于资源分组管理，多个元素之间用,分隔。 | [project-5kyt09jzwcqh6j] | [ListProjects] | No |
| **Keyword** | string | 搜索关键词 | [example_key] | [] | No |
| **Offset** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 默认值：0 | [0] | [] | No |
| **Limit** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。默认值：10 | [20] | [] | No |
| **PeerID** | string | 打通的资源ID，网络互通功能用于实现同租户两个 VPC 之间的网络互通，租户可以通过网络互通功能在两个 VPC 之间建立连接，如此就可以使用私有 IP 地址在两个 VPC 之间进行通信。 | [vpc-l3m0g6u2re0g25] | [DescribeVPC] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*VPCDetailInfo*](#VPCDetailInfo)> | 查询到的详情信息 | [] |
| **TotalCount** | integer | 查询到的总记录数 | [] |

### 数据模型

#### PeeringInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **PeerID** | string | 联通的对端 VPC ID | [vpc-78lvjhgaggn0tf] |
| **PeerName** | string | 联通的对端VPC名称 | [test-vpc] |
| **PeerType** | string | 联通的对端类型 | [VPC] |
| **PeerCIDRs** | array<string> | 对端CIDR（网段） | [172.16.0.0/16] |
| **Status** | string | 对端联通的状态 | [Activate] |

#### SubnetInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **SubnetID** | string | 子网ID | [subnet-v0wzld8nnuobbr] |
| **Region** | string | 地域 | [pro] |
| **Name** | string | 子网名称 | [test] |
| **Remark** | string | 备注，用于进行说明和注释，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [这是一个备注] |
| **ProjectID** | string | 项目组ID，用于资源分组管理 | [project-5kyt09jzwcqh6j] |
| **ProjectName** | string | 项目组名称 | [default] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> | 标签 | [cluster-resouce.system:offline-node] |
| **CreateTime** | integer | 创建时间，Unix 时间戳（秒级） | [1754644992] |
| **UpdateTime** | integer | 更新时间，Unix 时间戳（秒级） | [1754644992] |
| **Reason** | string | 创建失败原因 | [子网已存在] |
| **Network** | string | 子网网段，包含在 指定的 VPC网络 内。示例：10.0.0.0/22 | [10.0.0.0/22] |
| **Subnet** | string | 子网网段，包含在 指定的 VPC网络 内。示例：10.0.0.0/22 | [10.0.0.0/22] |
| **VPCID** | string | 子网所属的vpcID | [vpc-l3m0g6u2re0g25] |
| **State** | string | 子网状态 | [Available] |
| **ExpandNetwork** | string | 拓展网段 | [] |
| **IPVersion** | string | IP版本 | [] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

#### VPCDetailInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **VPCID** | string | vpcID | [vpc-lj1nck407g85o7] |
| **Name** | string | VPC名称 | [test-vpc] |
| **Remark** | string | 备注，用于进行说明和注释，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [这是一个备注] |
| **Region** | string | 地域 | [Region-A] |
| **RegionAlias** | string | 地域别名 | [地域A] |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] |
| **CompanyName** | string | VPC 所属的租户名称 | [ucloud] |
| **Email** | string | 租户邮箱，租户注册时提供的用于接收通知和联系的邮箱 | [example@ucloud.cn] |
| **ProjectID** | string | 项目组ID，用于资源分组管理 | [project-5kyt09jzwcqh6j] |
| **ProjectName** | string | 项目组名称 | [project01] |
| **Reason** | string | 创建失败原因 | [VPC已存在] |
| **CreateTime** | integer | 创建时间，Unix 时间戳（秒级） | [1754644992] |
| **UpdateTime** | integer | 更新时间，Unix 时间戳（秒级） | [1754644992] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> | 标签 | [cluster-resouce.system:offline-node] |
| **SubnetCount** | integer | 该 VPC 下子网数量 | [3] |
| **Network** | string | VPC默认网段 | [] |
| **SubnetInfos** | array<[*SubnetInfo*](#SubnetInfo)> | 子网信息详情 | [] |
| **PeeringInfos** | array<[*PeeringInfo*](#PeeringInfo)> | 对端的 VPC 信息 | [] |
| **State** | string | VPC 状态 | [Available] |
| **ExpandSubnetCount** | integer | 拓展VPC网络相关的子网数量 | [] |
| **ExpandNetwork** | string | VPC拓展网段 | [] |

## 创建子网 - CreateSubnet

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | **Yes** |
| **ProjectID** | string | 项目组ID，用于资源分组管理 | [project-5kyt09jzwcqh6j] | [ListProjects] | No |
| **TagKeyValuePairs** | array<string> | 标签键值对，用于资源标记和分类管理。格式为 key:value 的字符串，传入​​ Base64 编码的字符串。 | [dGVzdA==:dGVzdC0x] | [] | **No** |
| **Name** | string | 子网名称，名称只能包含中英文、数字及-_. | [subset-test] | [] | **Yes** |
| **Remark** | string | 备注，用于进行说明和注释，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [这是一个备注] | [] | No |
| **Network** | string | 子网网段，包含在 指定的 VPC网络 内。示例：10.0.0.0/22 | [10.0.0.0/22] | [] | **Yes** |
| **VPCID** | string | VPC ID | [vpc-lj1nck407g85o7] | [DescribeVPC] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **SubnetID** | string | 子网ID | [subnet-7aedosr3p0bfu8] |


## 删除子网 - DeleteSubnet

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **SubnetID** | string | 子网ID | [subnet-7aedosr3p0bfu8] | [DescribeSubnet] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 添加子网网络 - AddSubnetNetwork

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **SubnetID** | string | SubnetID | [] | [] | **Yes** |
| **Network** | string | 需要加载的子网 cidr | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 移除子网网络 - RemoveSubnetNetwork

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **SubnetID** | string | SubnetID | [] | [] | **Yes** |
| **Network** | string | 需要移除的子网 cidr | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取子网 - DescribeSubnet

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **SubnetIDs** | array<string> | 子网ID列表，用于查询指定的子网，多个元素之间用,分隔。 | [subnet-v0wzld8nnuobbr] | [DescribeSubnet] | No |
| **Offset** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 默认值：0 | [0] | [] | No |
| **Limit** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。默认值：10 | [20] | [] | No |
| **VPCID** | string | VPC ID，用于查询指定 VPC 下的子网 | [vpc-l3m0g6u2re0g25] | [DescribeVPC] | No |
| **Keyword** | string | 搜索关键词 | [example_key] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*SubnetInfo*](#SubnetInfo)> | 子网信息 | [] |
| **TotalCount** | integer | 查询到的总记录数 | [] |

### 数据模型

#### SubnetInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **SubnetID** | string | 子网ID | [subnet-v0wzld8nnuobbr] |
| **Region** | string | 地域 | [pro] |
| **Name** | string | 子网名称 | [test] |
| **Remark** | string | 备注，用于进行说明和注释，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [这是一个备注] |
| **ProjectID** | string | 项目组ID，用于资源分组管理 | [project-5kyt09jzwcqh6j] |
| **ProjectName** | string | 项目组名称 | [default] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> | 标签 | [cluster-resouce.system:offline-node] |
| **CreateTime** | integer | 创建时间，Unix 时间戳（秒级） | [1754644992] |
| **UpdateTime** | integer | 更新时间，Unix 时间戳（秒级） | [1754644992] |
| **Reason** | string | 创建失败原因 | [子网已存在] |
| **Network** | string | 子网网段，包含在 指定的 VPC网络 内。示例：10.0.0.0/22 | [10.0.0.0/22] |
| **Subnet** | string | 子网网段，包含在 指定的 VPC网络 内。示例：10.0.0.0/22 | [10.0.0.0/22] |
| **VPCID** | string | 子网所属的vpcID | [vpc-l3m0g6u2re0g25] |
| **State** | string | 子网状态 | [Available] |
| **ExpandNetwork** | string | 拓展网段 | [] |
| **IPVersion** | string | IP版本 | [] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

## 获取子网可用IP数量 - GetSubnetAvailableIPQuota

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | **Yes** |
| **SubnetID** | string | 子网ID | [subnet-7aedosr3p0bfu8] | [DescribeSubnet] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **UsedCount** | integer | 已用子网数量 | [1] |
| **AvailableCount** | integer | 可用子网数量 | [2] |


## 获取子网中申请出来的 IP 列表 - ListAllocatedIPsInSubnet

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **SubnetID** | string | 子网ID | [subnet-7aedosr3p0bfu8] | [DescribeSubnet] | **Yes** |
| **OnlyCount** | bool | 是否仅返回 IP 数量 | [true] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **AllocatedIPCount** | integer | 已分配 IP 数量 | [8] |
| **AllocatedIPNames** | array<string> | 已分配 IP 名称列表 | [] |


## 删除子网路由 - DeleteSubnetRoute

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | **Yes** |
| **SubnetID** | string | 子网ID | [subnet-v0wzld8nnuobbr] | [DescribeSubnet] | **Yes** |
| **Destination** | string | 目的地址，路由策略的目的地址，格式：<网络前缀>/<前缀长度> | [0.0.0.0/16] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 更新子网路由 - UpdateSubnetRoute

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | **Yes** |
| **Remark** | string | 备注，用于进行说明和注释，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [这是一个备注] | [] | No |
| **SubnetID** | string | 子网ID，当前所属的子网ID | [subnet-v0wzld8nnuobbr] | [DescribeSubnet] | **Yes** |
| **Destination** | string | 目的地址，格式：<网络前缀>/<前缀长度> | [0.0.0.0/16] | [] | **Yes** |
| **NextHop** | string | 下一跳的 IP 地址 | [10.0.128.1] | [] | **Yes** |
| **NextHopType** | string | 下一跳的类型，取值范围：1、VIP（虚拟 IP 地址）2、VM（虚拟机）3、Custom（自定义） | [VIP VM] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 查询子网路由 - DescribeSubnetRoute

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **SubnetID** | string | 子网ID，当前所属的子网ID | [subnet-v0wzld8nnuobbr] | [DescribeSubnet] | **Yes** |
| **Offset** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 默认值：0 | [0] | [] | No |
| **Limit** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。默认值：10 | [20] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 查询到的总记录数 | [] |
| **Infos** | array<[*SubnetRouteInfo*](#SubnetRouteInfo)> | 子网路由信息 | [] |

### 数据模型

#### SubnetRouteInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **SubnetID** | string | 子网ID | [subnet-v0wzld8nnuobbr] |
| **SubnetName** | string | 子网名称 | [test] |
| **Destination** | string | 目的地址 | [172.16.0.0/16] |
| **NextHop** | string | 下一跳地址 | [10.0.128.1] |
| **NextHopName** | string | 下一跳名称，仅下一跳为 VM 和 VIP 时显示 | [vm-k24t7bkn7fhook] |
| **NextHopAddr** | string | 下一跳的 IP 地址 | [10.0.128.1] |
| **NextHopType** | string | 下一跳的类型，取值范围：1、VIP（虚拟 IP 地址）2、VM（虚拟机）3、Custom（自定义） | [VIP VM] |
| **Remark** | string | 备注，用于进行说明和注释，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [这是一个备注] |

## 创建子网路由 - CreateSubnetRoute

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | **Yes** |
| **Remark** | string | 备注，用于进行说明和注释，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [这是一个备注] | [] | No |
| **SubnetID** | string | 子网ID | [subnet-v0wzld8nnuobbr] | [DescribeSubnet] | **Yes** |
| **Destination** | string | 目的地址，格式：<网络前缀>/<前缀长度> | [0.0.0.0/16] | [] | **Yes** |
| **NextHop** | string | 下一跳的 IP 地址 | [10.0.128.1] | [] | **Yes** |
| **NextHopType** | string | 下一跳的类型，取值范围：1、VIP（虚拟 IP 地址）2、VM（虚拟机）3、Custom（自定义） | [VIP VM] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 创建VPC对等连接 - AssociateVPCPeering

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | **Yes** |
| **VPCID** | string | 源vpcID | [vpc-l3m0g6u2re0g25] | [DescribeVPC] | **Yes** |
| **PeerType** | string | 对端类型，取值范围：1、VPC（虚拟私有云）2、DC（专线接入）  | [VPC] | [] | **Yes** |
| **PeerID** | string | 对端的资源 ID | [vpc-78lvjhgaggn0tf] | [DescribeVPC] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 删除VPC对等连接 - DissociateVPCPeering

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | **Yes** |
| **VPCID** | string | 源vpcID | [vpc-l3m0g6u2re0g25] | [DescribeVPC] | **Yes** |
| **PeerID** | string | 对端的资源 ID | [vpc-78lvjhgaggn0tf] | [DescribeVPC] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 更新产品内网IP - ReplaceIP

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **ResourceType** | string | 资源类型，取值范围：1、OSS（对象存储）2、FS（文件存储）3、MYSQL（MySQL）4、REDIS（Redis） | [OSS] | [] | **Yes** |
| **ResourceID** | string | 目的资源ID | [mysql-duc5j1x0mwfeqk] | [DescribeOSS DescribeFS DescribeMySQL DescribeRedis] | **Yes** |
| **IP** | string |  IP地址，新的内网 IP 地址 | [10.0.40.13] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|

