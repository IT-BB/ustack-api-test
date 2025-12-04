# SEGMENT

## 查询线路 - DescribeSegment

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **Keyword** | string | 搜索关键词 | [example_key] | [] | No |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **SegmentIDs** | array<string> | 线路ID列表，用于查询指定ID的外网线路，多个元素之间用,分隔。 | [segment-test] | [DescribeSegment] | No |
| **IPVersion** | string | 网络协议，用于查询指定 IP 版本的外网线路。取值范围：1、IPv4 2、IPv6 | [IPv4 IPv6] | [] | No |
| **Offset** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。默认值：10 | [20] | [] | No |
| **Limit** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 默认值：0 | [0] | [] | No |
| **ShowIPResourceCount** | bool | 是否展示 IP 资源数量 | [true] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*SegmentInfo*](#SegmentInfo)> | 外网线路详情 | [] |
| **TotalCount** | integer | 查询得到的总数 | [] |

### 数据模型

#### SegmentInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Region** | string | 地域，一批可共享的物理资源使用集合 | [] |
| **Name** | string | 外网线路名称 | [] |
| **Description** | string | 外网线路描述 | [] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> | 标签，用于资源标记和分类管理。 | [] |
| **Permission** | string | 访问权限，该外网线路可以被哪些租户访问 | [] |
| **SegmentID** | string | 外网线路ID | [] |
| **IPVersion** | string | IP版本 | [] |
| **Segment** | string | 外网网段 | [] |
| **IPRanges** | string | 可用IP范围 | [] |
| **UsedCount** | integer | 已用IP数量 | [] |
| **AvailableCount** | integer | 可用IP数量 | [] |
| **Gateway** | string | 网关，网关地址 | [] |
| **Device** | string | 网卡ID | [] |
| **Vlan** | string | Vlan，Vlan ID | [] |
| **Routing** | string | 路由 | [] |
| **Platform** | bool | 是否为平台保留使用的线路 | [] |
| **Status** | string | 当前外网线路状态 | [] |
| **CreateTime** | integer | 创建时间，Unix 时间戳（秒级） | [] |
| **UpdateTime** | integer | 更新时间，Unix 时间戳（秒级） | [] |
| **RegionAlias** | string | 地域别称 | [] |
| **FailedEIPCount** | integer | 失败 EIP 数量 | [] |
| **BoundedEIPCount** | integer | 已绑定的 EIP 数量 | [] |
| **RecycledEIPCount** | integer | 处于回收站的 EIP 数量 | [] |
| **DHCPServerIP** | string | dhcp服务IP | [] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

## 创建外网线路 - CreateSegment

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **Name** | string | 外网网络名称，长度为1-50个字符，名称只能包含中英文、数字、点（.），下划线（_）和中划线（-） | [segment-test] | [] | **Yes** |
| **Remark** | string | 备注，用于进行说明和注释，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [这是一个备注] | [] | No |
| **TagKeyValuePairs** | array<string> | 标签键值对，用于资源标记和分类管理。格式为 key:value 的字符串，传入​​ Base64 编码的字符串。 | [dGVzdA==:dGVzdC0x] | [] | **Yes** |
| **Permission** | string | 权限范围，指定该外网网络租户权限。如果设置为all则表示所有租户可用，如果设置租户ID的列表则指定用户可用 | [all 200000230,200000232] | [] | **Yes** |
| **Device** | string | 网卡名称，网卡格式为 bond0，bond1 等，若不确定请联系部署人员或网络管理员。 | [bond0] | [] | **Yes** |
| **Vlan** | string | Vlan，VLAN 范围为 0-4094 | [100] | [] | No |
| **Segment** | string | 网段，指定网段，支持 IPv4 和 IPv6 | [192.168.178.0/24] | [] | **Yes** |
| **IPRange** | string | IP范围，可用的 IP 地址范围 | [192.168.178.221-192.168.178.250] | [] | No |
| **EnableDHCP** | bool | 是否开启DHCP | [true] | [] | No |
| **DHCPServerIP** | string | DHCP服务IP，传值表示启用DHCP，传空表示禁用DHCP | [192.168.178.2] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **SegmentID** | string | 线路ID，创建成功返回的外网线路ID | [] |


## 删除外网线路 - DeleteSegment

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **SegmentID** | string | 线路ID，目的外网线路ID | [segment-35dehi0g8zemz] | [DescribeSegment] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 创建外网线路路由 - CreateSegmentRoute

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **Remark** | string | 备注，用于进行说明和注释，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [这是一个备注] | [] | No |
| **SegmentID** | string | 线路ID，目的外网线路ID | [segment-35dehi0g8zemz] | [DescribeSegment] | **Yes** |
| **Destination** | string | 目的地址，仅支持 CIDR 网段格式 | [10.0.0.0/16] | [] | **Yes** |
| **NextHop** | string | 下一跳的 IP 地址 | [10.0.128.1] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 更新外网线路路由 - UpdateSegmentRoute

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **Remark** | string | 备注，用于进行说明和注释，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [这是一个备注] | [] | No |
| **SegmentID** | string | 线路ID，目的外网线路ID | [segment-35dehi0g8zemz] | [DescribeSegment] | **Yes** |
| **Destination** | string | 目的地址，仅支持 CIDR 网段格式 | [10.0.0.0/16] | [DescribeSegmentRoute] | **Yes** |
| **NextHop** | string | 下一跳的 IP 地址 | [10.0.128.1] | [DescribeSegmentRoute] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 删除外网线路路由 - DeleteSegmentRoute

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **SegmentID** | string | 线路ID，目的外网线路ID | [segment-35dehi0g8zemz] | [DescribeSegmentRoute] | **Yes** |
| **Destination** | string | 目的地址，仅支持 CIDR 网段格式 | [10.0.0.0/16] | [DescribeSegmentRoute] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 查询外网线路路由 - DescribeSegmentRoute

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **SegmentID** | string | 线路ID，目的外网线路ID | [segment-35dehi0g8zemz] | [DescribeSegment] | **Yes** |
| **Offset** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。默认值：10 | [20] | [] | No |
| **Limit** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 默认值：0 | [0] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 查询得到的总数 | [] |
| **Infos** | array<[*SegmentRouteInfo*](#SegmentRouteInfo)> | 外网线路路由详情信息 | [] |

### 数据模型

#### SegmentRouteInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **SegmentID** | string | 外网线路ID | [] |
| **SegmentName** | string | 外网线路名称 | [] |
| **Destination** | string | 目的地址，外网线路路由的目的地址 | [] |
| **NextHop** | string | 下一跳，外网线路路由的下一跳地址 | [] |
| **NextHopName** | string | 下一跳名称，外网线路路由的下一条名称 | [] |
| **NextHopAddr** | string | 下一跳地址，外网线路路由的下一跳地址 | [] |
| **NextHopType** | string | 下一跳类型 | [] |
| **Remark** | string | 备注，用于进行说明和注释，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [这是一个备注] |

## 更新外网线路 - UpdateSegment

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **SegmentID** | string | 线路ID，目的外网线路ID | [segment-35dehi0g8zemz] | [DescribeSegment] | **Yes** |
| **IPRange** | string | IP范围，可用的 IP 地址范围 | [192.168.178.221-192.168.178.250] | [] | No |
| **Device** | string | 网卡名称，网卡格式为 bond0，bond1 等，若不确定请联系部署人员或网络管理员。 | [bond0] | [] | No |
| **Vlan** | string | Vlan，VLAN 范围为 0-4094 | [100] | [] | No |
| **DHCPServerIP** | string | DHCP服务IP，传值表示启用DHCP，传空表示禁用DHCP | [192.168.178.2] | [] | No |
| **UpdateMode** | string | 更新模式，传IPRange只会更新IPRange字段；传Device只会更新Device、Vlan字段；传DHCPServerIP只会更新DHCPServerIP字段；传All全部更新；传空只会更新IPRange，这是为了向前兼容 | [all] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 创建DirectConnect专线接入 - CreateDirectConnect

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **TagKeyValuePairs** | array<string> | 标签键值对，用于资源标记和分类管理。格式为 key:value 的字符串，传入​​ Base64 编码的字符串。 | [dGVzdA==:dGVzdC0x] | [] | **Yes** |
| **Name** | string | 专线接入名称，长度为1-50个字符，名称只能包含中英文、数字、点（.），下划线（_）和中划线（-） | [direct-fewfwf] | [] | **Yes** |
| **Remark** | string | 备注，用于进行说明和注释，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [这是一个备注] | [] | No |
| **Permission** | string | 权限范围，指定该外网网络租户权限。如果设置为all则表示所有租户可用，如果设置租户ID的列表则指定用户可用 | [all 200000230,200000232] | [] | **Yes** |
| **Device** | string | 网卡名称，网卡格式为 bond0，bond1 等，若不确定请联系部署人员或网络管理员。 | [bond0] | [] | **Yes** |
| **VLAN** | string | Vlan，VLAN 范围为 0-4094 | [100] | [] | No |
| **Bandwidth** | integer | 带宽，范围：1-20000 Mpbs，实际不超过物理带宽上限 | [20] | [] | No |
| **LocalGatewayIP** | string | 本端网关，UCloudStack侧网络接口互联的带掩码IP地址 | [192.168.1.1/24] | [] | **Yes** |
| **RemoteGatewayIP** | string | 远端网关，用户本地数据中心侧网络的互联的带掩码IP地址。远端网关与本地网关需要设置为同一网段的IP地址 | [192.168.1.2/24] | [] | **Yes** |
| **RemoteSubnetCIDRs** | array<string> | 远端网段，用户数据中心需要互连的网段 | [10.1.0.0/16] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **DirectConnectID** | string | 专线接入ID | [] |


## 删除DirectConnect专线接入 - DeleteDirectConnect

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **DirectConnectID** | string | 专线接入ID，目的专线接入的ID | [direct-ewg34g] | [DescribeDirectConnect] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 查询DirectConnect专线接入 - DescribeDirectConnect

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **DirectConnectIDs** | array<string> | 专线接入ID列表，用于查询指定的专线接入信息，多个元素之间用,分隔。 | [direct-wfebe3] | [DescribeDirectConnect] | No |
| **Keyword** | string | 搜索关键词 | [example_key] | [] | No |
| **Offset** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。默认值：10 | [20] | [] | No |
| **Limit** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 默认值：0 | [0] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*DirectConnectDetailInfo*](#DirectConnectDetailInfo)> | 专线接入的信息详情 | [] |
| **TotalCount** | integer | 查询得到的总数 | [] |

### 数据模型

#### DirectConnectDetailInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string | 专线接入的名称 | [] |
| **Remark** | string | 备注，用于进行说明和注释，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [] |
| **CreateTime** | integer | 创建时间，Unix 时间戳（秒级） | [] |
| **UpdateTime** | integer | 更新时间，Unix 时间戳（秒级） | [] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> | 标签，用于资源标记和分类管理。 | [] |
| **Region** | string | 地域，一批可共享的物理资源使用集合 | [] |
| **Permission** | string | 访问权限，该专线可以被哪些租户访问 | [] |
| **DirectConnectID** | string | 专线接入ID | [] |
| **Bandwidth** | integer | 带宽，单位:Mb | [] |
| **Device** | string | 网卡ID | [] |
| **VLAN** | string | VLAN，Vlan ID | [] |
| **LocalGatewayIP** | string | 本端网关IP | [] |
| **RemoteGatewayIP** | string | 远端网关IP | [] |
| **RemoteSubnetCIDRs** | array<string> | 远端网段 | [] |
| **Status** | string | 状态 | [] |
| **VPCInfos** | array<[*VPCInfo*](#VPCInfo)> | 关联vpc的信息 | [] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

#### VPCInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **VPCID** | string | vpcID | [] |
| **Name** | string | 名称，VPC 的名称 | [] |
| **Remark** | string | 备注，用于进行说明和注释，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空。 | [] |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [] |
| **Email** | string | 租户邮箱，租户注册时提供的用于接收通知和联系的邮箱。 | [] |

## 修改DirectConnect专线接入限速 - UpdateDirectConnectBandwidth

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **DirectConnectID** | string | 专线接入ID，目的专线接入的ID | [direct-ewg34g] | [DescribeDirectConnect] | **Yes** |
| **Bandwidth** | integer | 带宽，范围：1-20000 Mpbs，实际不超过物理带宽上限 | [20] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 修改DirectConnect专线接入远端子网网段 - UpdateDirectConnectRemoteSubnetCIDRs

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **DirectConnectID** | string | 专线接入ID，目的专线接入的ID | [direct-ewg34g] | [DescribeDirectConnect] | **Yes** |
| **CIDRs** | array<string> | 远端网段，用户数据中心需要互连的网段 | [10.1.0.0/16] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|

