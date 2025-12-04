# IP

## 申请弹性IP - AllocateEIP

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 所属的地域 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [DescribeUser] | **Yes** |
| **ProjectID** | string | 项目组ID，用于资源分组管理。格式为 project- 开头的字符串 | [project-5acijnlza4lp57] | [ListProjects] | **Yes** |
| **TagKeyValuePairs** | array<string> | 标签键值对，用于资源标记和分类管理。格式为 key:value 的字符串数组 | [] | [DescribeTag] | No |
| **Name** | string | 外网IP名称，长度为1-50个字符，支持 Unicode 中 letter 分类下的字符（其中包括英文、中文等），ASCII 数字（0-9）。支持下划线（_）、半角句号（.）或者短划线（-）。必须以 Unicode 中 letter 分类下的字符开头。默认值：空。 | [testDiskName] | [] | **Yes** |
| **Remark** | string | 外网IP描述，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [用于测试的外网IP] | [] | No |
| **ChargeType** | string | 计费类型，用于指定外网IP的计费模式。取值范围：hour（按小时计费）、month（按月计费）、year（按年计费）。 | [hour] | [] | **Yes** |
| **Quantity** | integer | 计费数量，指定计费周期的数量。按月/年计费时表示创建 Quantity 个月/年的外网IP，按小时计费时强制为1。取值范围：正整数，最小值为1。 | [1] | [] | **Yes** |
| **OperatorName** | string | 网段 | [segment-clbkkliyi7yr4] | [DescribeSegment] | **Yes** |
| **Bandwidth** | integer | 带宽，单位：Mbps | [1] | [DescribeProductSpecification] | **Yes** |
| **IP** | string | IP地址，指定要绑定的IP地址。如果不指定，系统会自动分配一个可用的IP地址 | [192.168.171.180] | [DescribeSegment] | No |
| **IPVersion** | string | IP版本，指定IP地址的版本。取值范围：IPv4、IPv6。 | [IPv4] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **EIPID** | string | 外网IP ID，弹性IP的唯一标识符。 | [] |


## 释放弹性IP - ReleaseEIP

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 所属的地域，您可以调用 DescribeRegion 方法 查看最新的地域列表。 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [DescribeUser] | **Yes** |
| **EIPID** | string | 外网IP ID，弹性IP的唯一标识符 | [eip-9yxs1wet7eoqs4] | [DescribeEIP] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 调整带宽 - ModifyEIPBandwidth

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 所属的地域，您可以调用 DescribeRegion 方法 查看最新的地域列表。 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [DescribeUser] | **Yes** |
| **EIPID** | string | 外网IP ID，弹性IP的唯一标识符 | [eip-9yxs1wet7eoqs4] | [DescribeEIP] | **Yes** |
| **Bandwidth** | integer | 带宽，单位：Mbps | [100] | [DescribeProductSpecification] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取弹性IP - DescribeEIP

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 所属的地域，您可以调用 DescribeRegion 方法 查看最新的地域列表。 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [DescribeUser] | No |
| **EIPIDs** | array<string> | 外网IP ID，弹性IP的唯一标识符 | [eip-9yxs1wet7eoqs4,eip-sa2a1wet7fasf2] | [DescribeEIP] | No |
| **ProjectIDs** | array<string> | 项目组ID，用于资源分组管理。格式为 project- 开头的字符串 | [project-5acijnlza4lp57,project-5acijnlza4lp57] | [ListProjects] | No |
| **Keyword** | string | 关键词，用于模糊搜索账号名、邮箱等字段，支持中英文搜索。长度限制：1-50字符。 | [eip] | [] | No |
| **IPVersion** | string | IP版本，指定IP地址的版本。取值范围：IPv4、IPv6。 | [IPv4] | [] | No |
| **BindResourceID** | string | 绑定的资源ID，指定要绑定的资源的唯一标识符。 | [vm-5acijnlza4lp57] | [] | No |
| **IPStatus** | string | EIP绑定状态，状态为已绑定（Bound），绑定中（Bounding）和未绑定（Free）。 | [Bound] | [] | No |
| **OperatorName** | string | 网段 | [segment-clbkkliyi7yr4] | [DescribeSegmennt] | No |
| **Offset** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 | [0] | [] | No |
| **Limit** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。 | [20] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*EIPInfo*](#EIPInfo)> | EIP信息列表 | [] |
| **TotalCount** | integer | 总记录数 | [] |

### 数据模型

#### EIPInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string | EIP名称 | [] |
| **Remark** | string | EIP备注 | [] |
| **Region** | string | EIP所属的地域 | [] |
| **RegionAlias** | string | EIP所属的地域别名 | [] |
| **CompanyID** | integer | 租户ID | [] |
| **CompanyName** | string | 租户名称 | [] |
| **Email** | string | 租户邮箱 | [] |
| **ProjectID** | string | 项目ID | [] |
| **ProjectName** | string | 项目名称 | [] |
| **Reason** | string | 失败原因 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> | 标签 | [] |
| **ChargeType** | string | EIP计费类型 | [] |
| **ExpireTime** | integer | 过期时间 | [] |
| **EIPID** | string | 外网IP ID，弹性IP的唯一标识符 | [] |
| **IP** | string | IP地址 | [] |
| **Mode** | string | EIP模式 | [] |
| **Bandwidth** | integer | EIP带宽 | [] |
| **IPVersion** | string | IP版本，IPv4和IPv6 | [] |
| **ISDefaultGW** | integer | 是否是默认网关 | [0] |
| **CanDefaultGW** | integer | 是否可以作为默认网关 | [1] |
| **IsElastic** | string | 是否弹性 | [N] |
| **NICID** | string | 网卡ID | [nic-5acijnlza4lp57] |
| **OperatorName** | string | 网段 | [segment-clbkkliyi7yr4] |
| **OperatorAlias** | string | 网段别名 | [segment-clbkkliyi7yr4] |
| **BindResourceID** | string | 绑定的资源ID | [vm-5acijnlza4lp57] |
| **BindResourceName** | string | 绑定的资源名称 | [1] |
| **BindResourceType** | string | 绑定的资源类型 | [VM] |
| **BindResourceProjectID** | string | 绑定的项目组ID | [project-5acijnlza4lp57] |
| **BindTime** | integer | 绑定时间 | [] |
| **Status** | string | EIP绑定状态，EIP绑定状态，状态为已绑定（Bound），绑定中（Bounding）和未绑定（Free）。 | [Bound] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

## 绑定弹性IP - BindEIP

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 所属的地域，您可以调用 DescribeRegion 方法 查看最新的地域列表。 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [DescribeUser] | **Yes** |
| **EIPID** | string | 外网IP ID，弹性IP的唯一标识符 | [eip-sa2a1wet7fasf2] | [DescribeEIP] | **Yes** |
| **ResourceType** | string | 资源类型，EIP绑定的资源类型 | [VM] | [ListProductResources] | **Yes** |
| **Mode** | string | EIP模式，取值范围：NAT、直通（Direct） | [NAT Direct] | [] | No |
| **ResourceID** | string | 资源ID，例如，如果ResourceType为VM，ResourceID就是VM的唯一标识符 | [vm-5acijnlza4lp57] | [DescribeVMInstance] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 解绑弹性IP - UnBindEIP

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 所属的地域，您可以调用 DescribeRegion 方法 查看最新的地域列表。 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [DescribeUser] | **Yes** |
| **EIPID** | string | 外网IP ID，弹性IP的唯一标识符 | [eip-sa2a1wet7fasf2] | [DescribeEIP] | **Yes** |
| **ResourceType** | string | 资源类型，EIP绑定的资源类型 | [VM] | [ListProductResources] | **Yes** |
| **ResourceID** | string | 资源ID，例如，如果ResourceType为VM，ResourceID就是VM的唯一标识符 | [vm-5acijnlza4lp57] | [DescribeVMInstance] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取EIP差价 - GetEIPDiffPrice

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 所属的地域，您可以调用 DescribeRegion 方法 查看最新的地域列表。 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [DescribeUser] | **Yes** |
| **EIPID** | string | 外网IP ID，弹性IP的唯一标识符 | [eip-sa2a1wet7fasf2] | [DescribeEIP] | **Yes** |
| **Bandwidth** | integer | 带宽，单位：Mbps | [100] | [DescribeProductSpecification] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*IpPriceInfo*](#IpPriceInfo)> | IP价格信息 | [] |

### 数据模型

#### IpPriceInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Price** | double | 价格 | [] |
| **ChargeType** | string | 计费类型 | [] |
| **OrderType** | string | 订单类型 | [UPGRADE] |

## 获取弹性IP价格 - GetEIPPrice

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 所属的地域，您可以调用 DescribeRegion 方法 查看最新的地域列表。 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [DescribeUser] | **Yes** |
| **ChargeType** | string | 计费类型，用于指定外网IP的计费模式。取值范围：hour（按小时计费）、month（按月计费）、year（按年计费）。 | [hour] | [] | **Yes** |
| **Quantity** | integer | 计费数量，指定计费周期的数量。按月/年计费时表示创建 Quantity 个月/年的外网IP，按小时计费时强制为1。取值范围：正整数，最小值为1。 | [1] | [] | **Yes** |
| **OperatorName** | string | 网段 | [segment-clbkkliyi7yr4] | [DescribeSegmennt] | **Yes** |
| **Bandwidth** | integer | 带宽，单位：Mbps | [1] | [DescribeProductSpecification] | **Yes** |
| **Count** | integer | EIP数量 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*IpPriceInfo*](#IpPriceInfo)> | IP价格信息 | [] |

### 数据模型

#### IpPriceInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Price** | double | 价格 | [] |
| **ChargeType** | string | 计费类型 | [] |
| **OrderType** | string | 订单类型 | [UPGRADE] |

## 查询IP是否使用中 - CheckIPInuse

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 所属的地域，您可以调用 DescribeRegion 方法 查看最新的地域列表。 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [DescribeUser] | No |
| **SubnetID** | string | 子网ID | [subnet-lzh7o9rpoq1r3t] | [DescribeSubnet] | No |
| **SegmentID** | string | 网段ID | [segment-clbkkliyi7yr4] | [DescribeSegmennt] | No |
| **FlatNetworkID** | string | 扁平网络ID | [] | [] | No |
| **IP** | string | IP地址 | [192.168.170.18] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|

