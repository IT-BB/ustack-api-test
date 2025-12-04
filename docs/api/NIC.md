# NIC

## 创建网卡 - CreateNIC

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | **Yes** |
| **ProjectID** | string | 项目组ID，用于资源分组管理 | [project-5kyt09jzwcqh6j] | [ListProjects] | No |
| **TagKeyValuePairs** | array<string> | 标签键值对，用于资源标记和分类管理。格式为 key:value 的字符串，传入​​ Base64 编码的字符串。 | [dGVzdA==:dGVzdC0x] | [] | **Yes** |
| **Name** | string | 网卡名称，名称只能包含中英文、数字及-_. | [NIC-test] | [] | **Yes** |
| **Remark** | string | 备注，用于进行说明和注释，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [这是一个备注] | [] | No |
| **ChargeType** | string | 计费类型，用于指定计费模式。取值范围：1、Dynamic（按小时计费）2、Month（按月计费）3、Year（按年计费） | [] | [] | **Yes** |
| **Quantity** | integer | 计费数量，指定计费周期的数量。按时/月/年计费时表示创建 Quantity 个时/月/年的磁盘，按小时计费时强制为1。取值范围：正整数 | [] | [] | **Yes** |
| **VPCID** | string | VPC ID，仅网卡类型为内网网卡时必须指定 | [vpc-l3m0g6u2re0g25] | [DescribeVPC] | No |
| **SubnetID** | string | 子网ID，仅网卡类型为内网网卡时必须指定 | [subnet-fe435gtgb434] | [DescribeSubnet] | No |
| **SGID** | string | 安全组ID，当前网卡需要绑定的安全组，用于管控进出网卡的网络流量；支持暂不绑定操作，即当前网卡暂不绑定安全组。 | [sg-v55mgokwcc11c5] | [DescribeSecurityGroup] | No |
| **IP** | string | IP 地址，当前网卡的 IP 地址，默认会从所属网络的 IP 地址段中自动分配 IP 地址，如需自定义 IP 地址，可在 IP 地址栏中输入指定的 IP 地址。 | [10.0.0.1] | [] | No |
| **NICType** | string | 网卡类型，包括内网网卡和外网网卡，分别从 VPC 和外网网段中分配 IP 地址。取值范围：1、LAN（内网）2、WAN（外网）3、Flat（扁平网络）. | [LAN] | [] | **Yes** |
| **SegmentID** | string | 外网线路ID，仅网卡类型为外网网卡时必须指定 | [segment-rwcyd8niajdmp2] | [DescribeSegment] | No |
| **Bandwidth** | integer | 带宽，单位：Mb，仅网卡类型为外网网卡时必须指定，有效范围1-20000 | [1000] | [] | No |
| **FlatNetworkID** | string | 扁平网络ID | [faltip-dqwfrgtb23rttyr] | [] | No |
| **MAC** | string | MAC地址，格式为6组2位十六进制数，以冒号或连字符分隔，不指定时自动分配 | [00:1A:2B:3C:4D:5E] | [] | No |
| **InAverageBandwidth** | integer | 入向平均带宽，仅网卡类型为扁平网络时指定，0 表示不限制，范围：0-20000 | [] | [] | No |
| **OutAverageBandwidth** | integer | 出向平均带宽，仅网卡类型为扁平网络时指定，0 表示不限制，范围：0-20000 | [] | [] | No |
| **IPVersion** | string | 协议，空值为所选VPC/外网默认网络的网络协议, 枚举支持 IPv4 IPv6 ALL 和空值 | [] | [] | No |
| **ExpandIP** | string | 拓展IP | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **NICID** | string | 弹性网卡ID | [elasticnic-14is5z8tu96h5a] |


## 删除网卡 - DeleteNIC

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | **Yes** |
| **NICID** | string | 弹性网卡ID | [elasticnic-14is5z8tu96h5a] | [DescribeNIC] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 修改弹性外网网卡的IP带宽 - UpdateNICIPBandwidth

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | **Yes** |
| **NICID** | string | 弹性网卡ID | [elasticnic-14is5z8tu96h5a] | [DescribeNIC] | No |
| **Bandwidth** | integer | 带宽，单位：Mb，仅网卡类型为外网网卡时必须指定，有效范围1-20000 | [1000] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 查询网卡信息 - DescribeNIC

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **Offset** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。默认值：10 | [20] | [] | No |
| **Limit** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 默认值：0 | [0] | [] | No |
| **Keyword** | string | 搜索关键词 | [example_key] | [] | No |
| **NICIDs** | array<string> | 网卡ID列表，用于查询指定的网卡，多个元素之间用,分隔。 | [elasticnic-12mxwtmk38niid] | [CreateNIC DescribeNIC] | No |
| **Status** | string | 状态，查询指定状态的网卡。取值范围：1、Bounding（绑定中）2、Bound（绑定）3、Unbounding（解绑中）4、Free（未绑定）5、Failed（创建失败） | [Free Failed Bound] | [] | No |
| **BindResourceID** | string | 绑定的资源ID，如果指定了此参数,查询此资源绑定的网卡 | [vm-t5g2g75hewr324] | [DescribeVMInstance] | No |
| **ProjectIDs** | array<string> | 项目组ID列表，用于资源分组管理，多个元素之间用,分隔。 | [project-5kyt09jzwcqh6j] | [ListProjects] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*NICInfo*](#NICInfo)> | 网卡信息详情 | [] |
| **TotalCount** | integer | 查询到的总记录数 | [] |

### 数据模型

#### NICIPInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **IPID** | string | ip ID | [ip-rpyeap76usq3lg] |
| **IP** | string | ip 地址 | [192.168.179.21] |
| **Bandwidth** | integer | ip 带宽，单位：Mb | [8192] |
| **IPVersion** | string | 协议 | [] |

#### NICInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **NICID** | string | 网卡ID | [elasticnic-b091lh96gugzmb] |
| **Name** | string | 网卡名称 | [test] |
| **Remark** | string | 备注，用于进行说明和注释，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [这是一个备注] |
| **Region** | string | 地域ID | [cn] |
| **RegionAlias** | string | 地域别名 | [中国] |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] |
| **CompanyName** | string | 租户名称 | [pony.ma] |
| **Email** | string | 租户邮箱 | [pony.ma@ucloud.cn] |
| **ProjectID** | string | 项目组ID，用于资源分组管理 | [project-5kyt09jzwcqh6j] |
| **ProjectName** | string | 项目组名称 | [project01] |
| **Reason** | string | 创建失败原因 | [没有可用地址] |
| **CreateTime** | integer | 创建时间，Unix 时间戳（秒级） | [1754644992] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> | 标签 | [cluster-resouce.system:offline-node] |
| **ChargeType** | string | 计费类型 | [Month] |
| **ExpireTime** | integer | 过期时间，Unix 时间戳（秒级） | [1754644992] |
| **NICType** | string | 网卡类型，包括内网网卡和外网网卡，分别从 VPC 和外网网段中分配 IP 地址。取值范围：LAN（内网），WAN（外网），Flat（扁平网络）. | [LAN] |
| **NICStatus** | string | 网卡状态，查询指定状态的网卡。取值范围：Bounding（绑定中）、Bound（绑定）、Unbounding（解绑中）、Free（未绑定）、Failed（创建失败） | [Free Failed Bound] |
| **MAC** | string | MAC地址，格式为6组2位十六进制数，以冒号或连字符分隔，不指定时自动分配 | [00:1A:2B:3C:4D:5E] |
| **SGID** | string | 绑定的安全组ID | [sg-z0d7ecehy5o8312] |
| **SGName** | string | 绑定的安全组名称 | [test] |
| **IPID** | string | 锁绑定的IP资源的 id | [ip-rpyeap76usq3lg] |
| **IP** | string | IP地址 | [192.168.179.21] |
| **VPCID** | string | 内网型网卡所属的VPC ID | [vpc-el7miluuojdjyu] |
| **VPCName** | string | 内网型网卡所属的VPC名称 | [test] |
| **SubnetID** | string | 内网型网卡的子网ID | [subnet-d7nod53eb0rwq2] |
| **SubnetName** | string | 内网型网卡的子网名称 | [test-subnet] |
| **SegmentID** | string | 外网型网卡的外网线路ID | [segment-rwcyd8niajdmp2] |
| **SegmentName** | string | 外网型网卡的外网线路名称 | [test-waiwang] |
| **SegmentNetwork** | string | 外网型网卡的外网线路网段 | [192.168.179.0/24] |
| **Bandwidth** | integer | 外网网卡带宽，单位：Mb | [1000] |
| **FlatNetworkID** | string | 扁平网络ID | [faltip-dqwfrgtb23rttyr] |
| **FlatNetworkName** | string | 扁平网络名称 | [test-faltip] |
| **FlatNetwork** | string | 扁平网络网段 | [192.168.179.0/24] |
| **BindResourceID** | string | 网卡所绑定的资源ID | [vm-13fggbr392gfg92] |
| **BindResourceType** | string | 绑定的资源类型 | [VM] |
| **BindResourceName** | string | 绑定的资源名称 | [test-vm] |
| **PFCode** | string | 物理网卡型号的标准编号 | [intel-x710] |
| **PFVendor** | string | 物理网卡厂商 | [Intel] |
| **PFProduct** | string | 物理网卡型号 | [X710-DA4] |
| **IsElastic** | bool | 是否弹性网卡 | [true] |
| **InTrafficShaping** | object | 入方向策略，格式：Enable:true(是否启用，入向平均带宽) | [Enable:false,AverageBandwidth:0] |
| **OutTrafficShaping** | object | 出方向策略，格式：Enable:true(是否启用，出向平均带宽) | [Enable:false,AverageBandwidth:0] |
| **IPInfos** | array<[*NICIPInfo*](#NICIPInfo)> | IP 信息 | [] |

#### TrafficShaping

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Enable** | bool | 是否开启 | [] |
| **AverageBandwidth** | integer | 平均带宽 | [] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

## 绑定网卡 - AttachNIC

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **NICID** | string | 网卡ID | [elasticnic-12mxwtmk38niid] | [DescribeNIC] | **Yes** |
| **ResourceID** | string | 资源ID，网卡所绑定的资源 ID | [vm-7rli7wr5jpk8h5] | [DescribeVMInstance] | **Yes** |
| **ResourceType** | string | 资源类型，网卡所绑定的资源类型，目前仅支持虚拟机（VM） | [VM] | [] | **Yes** |
| **PFCode** | string | 物理网卡型号的标准编号 | [intel-x710] | [DescribeNIC] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 解绑网卡 - DetachNIC

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | **Yes** |
| **NICID** | string | 网卡ID | [elasticnic-12mxwtmk38niid] | [DescribeNIC] | **Yes** |
| **ResourceID** | string | 资源ID | [vm-7rli7wr5jpk8h5] | [DescribeVMInstance] | **Yes** |
| **ResourceType** | string | 资源类型，可以不指定。（当前仅支持操作虚拟机和网卡的绑定与解绑） | [VM] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取弹性IP价格 - GetCreateNICPrice

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | **Yes** |
| **ChargeType** | string | 计费类型，用于指定计费模式。取值范围：1、Dynamic（按小时计费）2、Month（按月计费）3、Year（按年计费） | [Month] | [] | **Yes** |
| **Quantity** | integer | 计费数量，指定计费周期的数量。按月/年计费时表示创建 Quantity 个月/年的磁盘，按小时计费时强制为1。取值范围：正整数 | [1] | [] | **Yes** |
| **OpertatorName** | string | 外网网段ID | [segment-rwcyd8niajdmp2] | [DescribeSegment] | **Yes** |
| **Bandwidth** | integer | 带宽，单位:Mb，有效范围1-20000 | [1000] | [] | **Yes** |
| **Count** | integer | 购买数量 | [1] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*NicPriceInfo*](#NicPriceInfo)> | 网卡价格信息详情 | [] |

### 数据模型

#### NicPriceInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Price** | double | 总价格，单位（人民币，¥） | [1200] |
| **ChargeType** | string | 计费类型，取值范围：1、Dynamic（按小时计费）2、Month（按月计费）3、Year（按年计费） | [Month] |
| **OrderType** | string | 订单类型，改配时OrderType会返回改配类型，创建网卡时为空。取值范围：1、UPGRADE（升级）2、DOWNGRADE（降级）3、ORDER_TYPE_NONE（平级） | [UPGRADE] |

## 获取更新弹性网卡价格 - GetUpdateNICPrice

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **NICID** | string | 网卡ID | [elasticnic-12mxwtmk38niid] | [DescribeNIC] | **Yes** |
| **Bandwidth** | integer | 带宽，单位:Mb。有效范围1-20000 | [1000] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*NicPriceInfo*](#NicPriceInfo)> | 修改带宽后网卡价格详情 | [] |

### 数据模型

#### NicPriceInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Price** | double | 总价格，单位（人民币，¥） | [1200] |
| **ChargeType** | string | 计费类型，取值范围：1、Dynamic（按小时计费）2、Month（按月计费）3、Year（按年计费） | [Month] |
| **OrderType** | string | 订单类型，改配时OrderType会返回改配类型，创建网卡时为空。取值范围：1、UPGRADE（升级）2、DOWNGRADE（降级）3、ORDER_TYPE_NONE（平级） | [UPGRADE] |

## 修改网卡的MAC - UpdateNICMAC

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **NICID** | string | 网卡ID | [elasticnic-12mxwtmk38niid] | [DescribeNIC] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | **Yes** |
| **MAC** | string | 自定义MAC，新的 MAC 地址，格式为6组2位十六进制数，以冒号或连字符分隔 | [00:1A:2B:3C:4D:5E] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 查询MAC是否使用中 - CheckMACInUse

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **MAC** | string | MAC地址，格式为6组2位十六进制数，以冒号或连字符分隔 | [00:1A:2B:3C:4D:5E] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 修改网卡的物理型号 - UpdateNICPF

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **NICID** | string | 网卡ID | [elasticnic-12mxwtmk38niid] | [] | **Yes** |
| **PFCode** | string | 物理网卡型号，PFCode传空字符串表示取消网卡的sriov属性。操作前提：1. 网卡必须绑定给虚拟机。2. 虚拟机必须关机 | [SR-IOV] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 更新网卡流量整形信息 - UpdateNICTrafficShaping

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **NICID** | string | 网卡ID | [elasticnic-12mxwtmk38niid] | [DescribeNIC] | **Yes** |
| **InAverageBandwidth** | integer | 入向平均带宽，-1 表示不限速, 0 表示不变 | [] | [] | No |
| **OutAverageBandwidth** | integer | 出向平均带宽，-1 表示不限速, 0 表示不变 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|

