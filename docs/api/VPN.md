# VPN

## 创建网关 - CreateVPNGW

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ProjectID** | string | 项目组ID | [] | [] | No |
| **TagKeyValuePairs** | array<string> | 标签键值对 | [] | [] | **Yes** |
| **Name** | string | VPN网关名称 | [] | [] | **Yes** |
| **Remark** | string | VPN网关备注 | [] | [] | No |
| **ChargeType** | string | 计费类型 | [] | [] | **Yes** |
| **Quantity** | integer | 购买数量 | [] | [] | **Yes** |
| **EIPID** | string | 外网IP | [] | [] | **Yes** |
| **VMType** | string | 计算集群类型 | [] | [] | **Yes** |
| **DiskSetType** | string | 存储集群 | [] | [] | No |
| **VPCID** | string | vpcID | [] | [] | **Yes** |
| **SubnetID** | string | 子网ID | [] | [] | **Yes** |
| **HighAvailability** | string | 高可用 | [] | [] | No |
| **TerminationPolicy** | integer | 是否开启删除保护, 0:开启; 1:关闭 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **VPNGWID** | string | 网关ID | [] |


## 创建对端网关 - CreateRemoteVPNGW

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ProjectID** | string | 项目组ID | [] | [] | No |
| **TagKeyValuePairs** | array<string> | 标签键值对 | [] | [] | **Yes** |
| **Name** | string | 对端网关名称 | [] | [] | **Yes** |
| **Remark** | string | 对端网关备注 | [] | [] | No |
| **IPAddress** | string | 对端网关IP地址 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **RemoteVPNGWID** | string | 对端网关ID | [] |


## 创建隧道 - CreateVPNTunnel

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ProjectID** | string | 项目组ID | [] | [] | No |
| **TagKeyValuePairs** | array<string> | 标签键值对 | [] | [] | **Yes** |
| **Name** | string | 隧道名称 | [] | [] | **Yes** |
| **Remark** | string | 隧道备注 | [] | [] | No |
| **VPNGWID** | string | 网关ID | [] | [] | **Yes** |
| **RemoteVPNGWID** | string | 对端网关ID | [] | [] | **Yes** |
| **LocalSubnetIDs** | array<string> | 本端子网ID | [] | [] | **Yes** |
| **RemoteSubnets** | array<string> | 对端子网 | [] | [] | **Yes** |
| **PreSharedKey** | string | 共享秘钥 | [] | [] | **Yes** |
| **IKEVersion** | string | IKE版本 | [] | [] | **Yes** |
| **IKEAuthenticationAlgorithm** | string | IKE认证算法 | [] | [] | **Yes** |
| **IKEEncryptionAlgorithm** | string | IKE加密算法 | [] | [] | **Yes** |
| **IKEExchangeMode** | string | IKE交换模式 | [] | [] | No |
| **IKEDhGroup** | integer | IKE DH组 | [] | [] | No |
| **IKELocalLabel** | string | IKE本地标识 | [] | [] | No |
| **IKERemoteLabel** | string | IKE对端标识 | [] | [] | No |
| **IKESALifetime** | integer | IKE SA生存周期 | [] | [] | **Yes** |
| **IPSecAuthenticationAlgorithm** | string | IPSec认证算法 | [] | [] | **Yes** |
| **IPSecEncryptionAlgorithm** | string | IPSec加密算法 | [] | [] | No |
| **IPSecProtocol** | string | IPSec安全传输协议 | [] | [] | **Yes** |
| **IPSecPFSDhGroup** | integer | PFS DH组 | [] | [] | No |
| **IPSecSALifetime** | integer | 生存周期 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **VPNTunnelID** | string | 隧道ID | [] |


## 更新隧道信息 - UpdateVPNTunnel

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **VPNTunnelID** | string | 隧道ID | [] | [] | **Yes** |
| **LocalSubnetIDs** | array<string> | 本端标识 | [] | [] | **Yes** |
| **RemoteSubnets** | array<string> | 对端标识 | [] | [] | **Yes** |
| **PreSharedKey** | string | 共享秘钥 | [] | [] | **Yes** |
| **IKEVersion** | string | IKE版本 | [] | [] | **Yes** |
| **IKEAuthenticationAlgorithm** | string | IKE认证算法 | [] | [] | **Yes** |
| **IKEEncryptionAlgorithm** | string | IKE加密算法 | [] | [] | **Yes** |
| **IKEExchangeMode** | string | IKE交换模式 | [] | [] | No |
| **IKEDhGroup** | integer | IKE DH组 | [] | [] | No |
| **IKELocalLabel** | string | 本端标识 | [] | [] | No |
| **IKERemoteLabel** | string | 对端标识 | [] | [] | No |
| **IKESALifetime** | integer | 生存周期 | [] | [] | **Yes** |
| **IPSecAuthenticationAlgorithm** | string | IPSec认证算法 | [] | [] | **Yes** |
| **IPSecEncryptionAlgorithm** | string | IPSec加密算法 | [] | [] | No |
| **IPSecProtocol** | string | IPSec安全传输协议 | [] | [] | **Yes** |
| **IPSecPFSDhGroup** | integer | DH组 | [] | [] | No |
| **IPSecSALifetime** | integer | 生存周期 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 删除隧道 - DeleteVPNTunnel

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **VPNTunnelID** | string | 隧道ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 删除对端网关 - DeleteRemoteVPNGW

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **RemoteVPNGWID** | string | 对端网关ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 从VPN解绑EIP - UnbindEIPFromVPN

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户 | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **VPNGWID** | string | 网关ID | [] | [] | **Yes** |
| **EIPID** | string | EIPID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 绑定EIP到VPN - BindEIPToVPN

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **VPNGWID** | string | 网关ID | [] | [] | **Yes** |
| **EIPID** | string | 外网IP | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 删除网关 - DeleteVPNGW

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **VPNGWID** | string | 网关ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取价格 - GetPrice

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ChargeType** | string | 计费类型 | [] | [] | **Yes** |
| **Quantity** | integer | 购买数量 | [] | [] | **Yes** |
| **Count** | integer | 购买数量 | [] | [] | **Yes** |
| **VMType** | string | 计算集群类型 | [] | [] | **Yes** |
| **HighAvailability** | string | 高可用 | [] | [] | No |
| **VPNGWID** | string | VPNGWID | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Info** | object | 订单详情 | [] |

### 数据模型

#### VpnPriceInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Price** | double | VPN价格 | [] |
| **ChargeType** | string | 支付类型 | [] |
| **PurchaseValue** | string | 价格 | [] |

## 获取网关信息 - DescribeVPNGW

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **Offset** | integer | 分页页码 | [] | [] | No |
| **Limit** | integer | 分页数量 | [] | [] | No |
| **ProjectIDs** | array<string> | 项目组ID | [] | [] | No |
| **VPNGWIDs** | array<string> | 网关ID | [] | [] | No |
| **Keyword** | string | 搜索关键字 | [] | [] | No |
| **Status** | array<string> | VPNGW状态 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数 | [] |
| **Infos** | array<[*VPNGWInfo*](#VPNGWInfo)> | vpn详情 | [] |

### 数据模型

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

#### VPNGWInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string | 名称 | [] |
| **Remark** | string | 备注 | [] |
| **Region** | string | 地域 | [] |
| **RegionAlias** | string | 地域别名 | [] |
| **CompanyID** | integer | 租户ID | [] |
| **CompanyName** | string | 租户名 | [] |
| **Email** | string | 租户邮箱 | [] |
| **Reason** | string | 错误原因 | [] |
| **ProjectID** | string | 项目组ID | [] |
| **ProjectName** | string | 项目组名称 | [] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> | 标签 | [] |
| **ChargeType** | string | 支付类型 | [] |
| **VPNGWID** | string | 网关ID | [] |
| **VPCID** | string | vpcID | [] |
| **SubnetID** | string | 子网ID | [] |
| **EIPID** | string | eipID | [] |
| **State** | string | 状态 | [] |
| **TunnelCount** | integer | 隧道数量 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **ExpireTime** | integer | 过期时间 | [] |
| **EIPIP** | string | 外网IP | [] |
| **Arch** | string | 架构 | [] |
| **VMType** | string | 计算集群类型 | [] |
| **HighAvailability** | string | 高可用模式 | [] |
| **StorageSetType** | string | 存储集群 | [] |
| **TerminationPolicy** | integer | 是否开启删除保护 | [] |

## 获取对端网关信息 - DescribeRemoteVPNGW

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **Offset** | integer | 分页页码 | [] | [] | No |
| **Limit** | integer | 分页数量 | [] | [] | No |
| **ProjectIDs** | array<string> | 项目组ID | [] | [] | No |
| **RemoteVPNGWIDs** | array<string> | 对端网关ID | [] | [] | No |
| **Keyword** | string | 搜索关键字 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 对端网关总数 | [] |
| **Infos** | array<[*RemoteVPNGWInfo*](#RemoteVPNGWInfo)> | 对端网关详情 | [] |

### 数据模型

#### RemoteVPNGWInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Region** | string | 地域 | [] |
| **CompanyID** | integer | 租户ID | [] |
| **Email** | string | 邮箱 | [] |
| **Name** | string | 名称 | [] |
| **Remark** | string | 备注 | [] |
| **ProjectID** | string | 项目组ID | [] |
| **ProjectName** | string | 项目组名称 | [] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> | 标签 | [] |
| **State** | string | 状态 | [] |
| **IPAddress** | string | IP地址 | [] |
| **RemoteVPNGWID** | string | 对端网关ID | [] |
| **TunnelCount** | integer | 隧道数量 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

## 获取隧道信息 - DescribeVPNTunnel

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **Offset** | integer | 分页页码 | [] | [] | No |
| **Limit** | integer | 分页数量 | [] | [] | No |
| **ProjectIDs** | array<string> | 项目组ID | [] | [] | No |
| **VPNTunnelIDs** | array<string> | 隧道ID | [] | [] | No |
| **Keyword** | string | 搜索关键字 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 隧道数 | [] |
| **Infos** | array<[*VPNTunnelInfo*](#VPNTunnelInfo)> | 隧道详情 | [] |

### 数据模型

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

#### VPNTunnelInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Region** | string | 地域 | [] |
| **CompanyID** | integer | 租户ID | [] |
| **Email** | string | 邮箱 | [] |
| **Name** | string | 名称 | [] |
| **Remark** | string | 备注 | [] |
| **ProjectID** | string | 项目组ID | [] |
| **ProjectName** | string | 项目组名称 | [] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> | 标签 | [] |
| **VPNTunnelID** | string | 隧道ID | [] |
| **State** | string | 状态 | [] |
| **ConnectState** | string | 连接状态 | [] |
| **VPNGWID** | string | 网关ID | [] |
| **RemoteVPNGWID** | string | 对端网关ID | [] |
| **LocalSubnetIDs** | array<string> | 本端标识 | [] |
| **RemoteSubnets** | array<string> | 对端标识 | [] |
| **VPCID** | string | vpcID | [] |
| **PreSharedKey** | string | 共享秘钥 | [] |
| **IKEVersion** | string | IKE版本 | [] |
| **IKEAuthenticationAlgorithm** | string | IKE认证算法 | [] |
| **IKEEncryptionAlgorithm** | string | IKE加密算法 | [] |
| **IKEExchangeMode** | string | IKE交换模式 | [] |
| **IKEDhGroup** | integer | DH组 | [] |
| **IKELocalLabel** | string | 本端标识 | [] |
| **IKERemoteLabel** | string | 对端标识 | [] |
| **IKESALifetime** | integer | 生命周期 | [] |
| **IPSecAuthenticationAlgorithm** | string | IPSec认证算法 | [] |
| **IPSecEncryptionAlgorithm** | string | IPSec加密算法 | [] |
| **IPSecProtocol** | string | 安全协议 | [] |
| **IPSecPFSDhGroup** | integer | DH组 | [] |
| **IPSecSALifetime** | integer | IPSec生命周期 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |
| **DeleteTime** | integer | 删除时间 | [] |

## 获取隧道配置 - GetVPNTunnelConfig

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **VPNGWID** | string | 网关ID | [] | [] | **Yes** |
| **TunnelID** | string | 隧道ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Info** | string | 配置内容 | [] |


## 升级为高可用版本 - UpgradeVPNGWToHA

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **VPNGWID** | string | VPN网关ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|

