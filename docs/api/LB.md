# LB

## 获取负载均衡价格 - GetLBPrice

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Specification** | string |  | [] | [] | No |
| **ChargeType** | string | 计费类型 | [] | [] | **Yes** |
| **Quantity** | integer | 计费周期 | [] | [] | No |
| **VMType** | string | 集群类型 | [] | [] | **Yes** |
| **CPU** | integer | CPU核数 | [] | [] | No |
| **HighAvailability** | string | 高可用 | [] | [] | No |
| **LBID** | string | LBID | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*LbPriceInfo*](#LbPriceInfo)> |  | [] |

### 数据模型

#### LbPriceInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Price** | double | 价格 | [] |
| **ChargeType** | string | 计费类型 | [] |
| **OrderType** | string | 订单类型 | [] |

## 创建负载均衡 - CreateLB

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **ProjectID** | string | 项目组ID | [] | [] | No |
| **TagKeyValuePairs** | array<string> | 标签键值对 | [] | [] | **Yes** |
| **Name** | string | 名称 | [] | [] | **Yes** |
| **Remark** | string | 备注 | [] | [] | No |
| **ApplicationName** | string | 审批名称 | [] | [] | No |
| **ApplicationReason** | string | 审批理由 | [] | [] | No |
| **ChargeType** | string | 计费类型 | [] | [] | **Yes** |
| **Quantity** | integer | 计费周期 | [] | [] | **Yes** |
| **VPCID** | string | VPCID | [] | [] | **Yes** |
| **SubnetID** | string | 子网ID | [] | [] | **Yes** |
| **SGID** | string | 安全组ID | [] | [] | No |
| **VMType** | string | 集群类型 | [] | [] | **Yes** |
| **DiskSetType** | string | 存储集群 | [] | [] | No |
| **LBType** | string | LB类型 | [] | [] | **Yes** |
| **EIPID** | string | EIPID, LB类型为外网时必填 | [] | [] | No |
| **CPU** | integer | CPU核数 | [] | [] | No |
| **HighAvailability** | string | 高可用 | [] | [] | No |
| **TerminationPolicy** | integer | 是否开启删除保护, 0:开启; 1:关闭 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **LBID** | string |  | [] |


## 更新负载均衡日志是否上传至S3 - UpdateLBLog

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **LBID** | string | LBID | [] | [] | **Yes** |
| **LogAccessEnable** | string | 访问日志开关 | [] | [] | **Yes** |
| **LogOssID** | string | 日志转储后端OSS，访问日志开关为On时必填 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 负载均衡日志实时查看开关 - UpdateLBAccessLogForLive

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **LBID** | string | LBID | [] | [] | **Yes** |
| **Switch** | string | 是否开启实时查看访问日志 | [] | [] | No |
| **LogMachineIP** | string | 访问日志机器IP | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **WebSocketURL** | string | WebSocket连接URL，只有当Switch=On时才会返回 | [] |


## 删除负载均衡 - DeleteLB

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **LBID** | string | LBID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取负载均衡信息 - DescribeLB

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **ProjectIDs** | array<string> | 项目组ID | [] | [] | No |
| **Keyword** | string | 关键字 | [] | [] | No |
| **VPCID** | string | vpcID | [] | [] | No |
| **SubnetID** | string | 子网ID | [] | [] | No |
| **LBIDs** | array<string> | lbID | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **Status** | array<string> | LB状态 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer |  | [] |
| **Infos** | array<[*LBInfo*](#LBInfo)> |  | [] |

### 数据模型

#### LBInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string | 名称 | [] |
| **Remark** | string | 备注 | [] |
| **Region** | string | 地域ID | [] |
| **RegionAlias** | string | 地域名称 | [] |
| **CompanyID** | integer | 租户ID | [] |
| **CompanyName** | string | 租户名称 | [] |
| **Email** | string | 邮箱 | [] |
| **ProjectID** | string | 项目组ID | [] |
| **ProjectName** | string | 项目组名称 | [] |
| **Reason** | string | 失败原因 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |
| **DeleteTime** | integer | 删除 | [] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> |  | [] |
| **ChargeType** | string | 计费类型 | [] |
| **ExpireTime** | integer | 过期时间 | [] |
| **VMType** | string | 集群类型 | [] |
| **VPCID** | string | vpcID | [] |
| **VPCName** | string | VPC名称 | [] |
| **SubnetID** | string | 子网ID | [] |
| **SubnetName** | string | 子网名称 | [] |
| **SGID** | string | 安全组ID | [] |
| **SGName** | string | 安全组名称 | [] |
| **LBID** | string | lbID | [] |
| **LBType** | string | lb类型 | [] |
| **LBStatus** | string | lb状态 | [] |
| **VSCount** | integer | vs数量 | [] |
| **PrivateIP** | string | 内网IP | [] |
| **PublicIP** | string | 外网IP | [] |
| **LogAccessEnable** | string | 访问日志 | [] |
| **LogOssID** | string | 日志存储OSSID | [] |
| **CPU** | integer |  | [] |
| **HighAvailability** | string |  | [] |
| **SetArch** | string | 计算集群架构 | [] |
| **StorageSetType** | string | 存储集群 | [] |
| **TerminationPolicy** | integer | 是否开启删除保护 | [] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

## 绑定 eip 到 LB - BindEIPToLB

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ResourceID** | string | 资源ID | [] | [] | **Yes** |
| **EIPID** | string | EIPID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 从 LB 解绑 eip - UnbindEIPFromLB

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ResourceID** | string | 资源ID | [] | [] | **Yes** |
| **EIPID** | string | EIPID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 新增/更新 SG 到 LB - UpdateSGFromLB

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ResourceID** | string | 资源ID | [] | [] | **Yes** |
| **SGID** | string | 安全组ID | [] | [] | **Yes** |
| **NICType** | string | 网络类型 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 创建VS - CreateVS

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **LBID** | string | LBID | [] | [] | **Yes** |
| **Protocol** | string | VS协议 | [] | [] | **Yes** |
| **Port** | integer | VS端口 | [] | [] | **Yes** |
| **Scheduler** | string | 调度算法 | [] | [] | **Yes** |
| **HealthcheckType** | string | 健康检查类型 | [] | [] | **Yes** |
| **SSLMode** | string | SSL模式 | [] | [] | **Yes** |
| **KeepaliveTimeout** | integer | 连接空闲超时时间 | [] | [] | No |
| **PersistenceType** | string | 会话保持类型 | [] | [] | No |
| **PersistenceKey** | string | 会话保持KEY | [] | [] | No |
| **Path** | string | HTTP 健康检查的路径 | [] | [] | No |
| **Domain** | string | HTTP 健康检查时校验请求的 HOST 字段中的域名 | [] | [] | No |
| **ServerCertificateID** | string | 服务器证书ID | [] | [] | No |
| **CACertificateID** | string | CA证书ID，用于验证客户端证书的签名 | [] | [] | No |
| **RedirectEnable** | string | http重定向开关 | [] | [] | No |
| **RedirectVsID** | string | http重定向vsid | [] | [] | No |
| **ProxyProtocolEnable** | string | tcp proxy protocol开关，协议为 TCP 时必填 | [] | [] | No |
| **BackendProtocol** | string | 后端协议，枚举值：空、HTTP、HTTPS。默认为空，表示 http, 仅在 Protocol 为 HTTPS 时有效。 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **VSID** | string | VSID | [] |


## 删除VS - DeleteVS

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | string | 租户ID | [] | [] | No |
| **LBID** | string | LBID | [] | [] | **Yes** |
| **VSID** | string | VSID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取VS信息 - DescribeVS

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **LBID** | string | LBID | [] | [] | No |
| **VSIDs** | array<string> | VSIDs | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **Origin** | string | vs来源，枚举值有Default/Service/Ingress，传空表示不筛选来源 | [] | [] | No |
| **Keyword** | string | 搜索关键字 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer |  | [] |
| **Infos** | array<[*LbVSInfo*](#LbVSInfo)> |  | [] |

### 数据模型

#### LbRSInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string | 名称 | [] |
| **LBID** | string | LBID | [] |
| **VSID** | string | VSID | [] |
| **RSID** | string | RSID | [] |
| **BindResourceID** | string | 绑定的资源ID | [] |
| **HealthCheckAddress** | string | 健康检查地址 | [] |
| **IP** | string | 服务节点的内网 IP 地址 | [] |
| **Port** | integer | 服务节点暴露的服务端口号 | [] |
| **Weight** | integer | 服务节点的权重 | [] |
| **RSMode** | string | 节点模式 | [] |
| **RSStatus** | string | 节点状态 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |
| **BindResourceName** | string | 绑定的资源名称 | [] |

#### LbVSInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **VSID** | string | VSID | [] |
| **LBID** | string | LBID | [] |
| **Protocol** | string | 协议 | [] |
| **Port** | integer | 端口 | [] |
| **RedirectVsID** | string | http重定向vsid | [] |
| **Scheduler** | string | 负载均衡的调度算法 | [] |
| **PersistenceType** | string | 会话保持类型 | [] |
| **HealthcheckType** | string | 负载均衡的健康检查类型 | [] |
| **KeepaliveTimeout** | integer | 负载均衡的连接空闲超时时间 | [] |
| **VSStatus** | string | VServer 的资源状态 | [] |
| **RSHealthStatus** | string | 健康检查状态 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |
| **SSLMode** | string | SSL认证模式 | [] |
| **PersistenceKey** | string | 会话保持KEY | [] |
| **Path** | string | HTTP 健康检查的路径 | [] |
| **Domain** | string | HTTP 健康检查时校验请求的 HOST 字段中的域名 | [] |
| **RSInfos** | array<[*LbRSInfo*](#LbRSInfo)> | RSInfos | [] |
| **VSPolicyInfos** | array<[*LbVSPolicyInfo*](#LbVSPolicyInfo)> | VSPolicyInfos | [] |
| **ServerCertificateID** | string | 服务器证书ID | [] |
| **CACertificateID** | string | CA证书ID | [] |
| **ProxyProtocolEnable** | string | tcp proxy protocol开关 | [] |
| **Origin** | string | vs来源，枚举值：''、'Service'、'Ingress' | [] |
| **BackendProtocol** | string | 后端协议，枚举值：空、HTTP、HTTPS。默认为空，表示 http, 仅在 Protocol 为 HTTPS 时有效。 | [] |

#### LbVSPolicyInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **LBID** | string | LBID | [] |
| **VSID** | string | VSID | [] |
| **PolicyID** | string | PolicyID | [] |
| **Domain** | string | 内容转发规则关联的请求域名 | [] |
| **Path** | string | 内容转发规则关联的请求访问路径 | [] |
| **PolicyStatus** | string | 状态 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |
| **RSInfos** | array<[*LbRSInfo*](#LbRSInfo)> | RSInfos | [] |
| **RedirectPolicyID** | string | http重定向VsPolicyId | [] |
| **SSLMode** | string | sni协议开启时，ssl模式配置 | [] |
| **ServerCertificateID** | string | sni协议开启时，服务器证书ID | [] |
| **CACertificateID** | string | sni协议开启时，CA证书ID | [] |

## 更新VS - UpdateVS

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **LBID** | string | LBID | [] | [] | **Yes** |
| **VSID** | string | VSID | [] | [] | **Yes** |
| **Port** | integer | VServer 的监听端口 | [] | [] | No |
| **Scheduler** | string | 负载均衡的调度算法 | [] | [] | No |
| **HealthcheckType** | string | 健康检查类型 | [] | [] | No |
| **KeepaliveTimeout** | integer | 连接空闲超时时间 | [] | [] | No |
| **SSLMode** | string | SSL认证模式 | [] | [] | **Yes** |
| **PersistenceType** | string | 会话保持类型 | [] | [] | No |
| **PersistenceKey** | string | 会话保持KEY | [] | [] | No |
| **Path** | string | HTTP 健康检查的路径 | [] | [] | No |
| **Domain** | string | HTTP 健康检查时校验请求的 HOST 字段中的域名 | [] | [] | No |
| **ServerCertificateID** | string | 服务器证书ID | [] | [] | No |
| **CACertificateID** | string | CA证书ID | [] | [] | No |
| **RedirectEnable** | string | http重定向开关，枚举值：On、Off | [] | [] | No |
| **RedirectVsID** | string | http重定向vsid | [] | [] | No |
| **ProxyProtocolEnable** | string | tcp proxy protocol开关，协议为 TCP 时必填，枚举值：On、Off | [] | [] | No |
| **BackendProtocol** | string | 后端协议，枚举值：空、HTTP、HTTPS。默认为空，表示 http, 仅在 Protocol 为 HTTPS 时有效。 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 添加服务节点 - CreateRS

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **LBID** | string | LBID | [] | [] | **Yes** |
| **VSID** | string | VSID | [] | [] | **Yes** |
| **BindResourceID** | string | 服务节点的资源 ID ，仅支持添加与 LB 相同 VPC 的虚拟机资源 | [] | [] | **Yes** |
| **Port** | integer | 服务节点暴露的服务端口号 | [] | [] | **Yes** |
| **Weight** | integer | 服务节点的权重 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **RSID** | string | RSID | [] |


## 删除服务节点 - DeleteRS

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **LBID** | string | LBID | [] | [] | **Yes** |
| **VSID** | string | VSID | [] | [] | **Yes** |
| **RSID** | string | RSID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取RS信息 - DescribeRS

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **LBID** | string | LBID | [] | [] | **Yes** |
| **VSID** | string | VSID | [] | [] | **Yes** |
| **RSIDs** | array<string> | RSIDs | [] | [] | No |
| **BindResourceID** | string | RS绑定的资源ID | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer |  | [] |
| **Infos** | array<[*LbRSInfo*](#LbRSInfo)> |  | [] |

### 数据模型

#### LbRSInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string | 名称 | [] |
| **LBID** | string | LBID | [] |
| **VSID** | string | VSID | [] |
| **RSID** | string | RSID | [] |
| **BindResourceID** | string | 绑定的资源ID | [] |
| **HealthCheckAddress** | string | 健康检查地址 | [] |
| **IP** | string | 服务节点的内网 IP 地址 | [] |
| **Port** | integer | 服务节点暴露的服务端口号 | [] |
| **Weight** | integer | 服务节点的权重 | [] |
| **RSMode** | string | 节点模式 | [] |
| **RSStatus** | string | 节点状态 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |
| **BindResourceName** | string | 绑定的资源名称 | [] |

## 更改RS - UpdateRS

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **LBID** | string | LBID | [] | [] | **Yes** |
| **VSID** | string | VSID | [] | [] | **Yes** |
| **RSID** | string | RSID | [] | [] | **Yes** |
| **Port** | integer | 服务节点暴露的服务端口号 | [] | [] | No |
| **Weight** | integer | 服务节点的权重 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 禁用节点 - DisableRS

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **LBID** | string | LBID | [] | [] | **Yes** |
| **VSID** | string | VSID | [] | [] | **Yes** |
| **RSID** | string | RSID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 启用节点 - EnableRS

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **LBID** | string | LBID | [] | [] | **Yes** |
| **VSID** | string | VSID | [] | [] | **Yes** |
| **RSID** | string | RSID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 创建转发规则 - CreateVSPolicy

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **LBID** | string | LBID | [] | [] | **Yes** |
| **VSID** | string | VSID | [] | [] | **Yes** |
| **RSIDs** | array<string> | RSIDs | [] | [] | No |
| **Path** | string | 内容转发规则关联的请求访问路径，如 '/' 。域名和路径至少需要指定一项，且域名和路径的组合在一个 VServer 中必须唯一。 | [] | [] | No |
| **Domain** | string | 内容转发规则关联的请求域名，值可为空，即代表仅匹配路径。域名和路径至少需要指定一项，且域名和路径的组合在一个 VServer 中必须唯一。 | [] | [] | No |
| **SSLMode** | string | sni协议开启时，ssl模式配置 | [] | [] | No |
| **ServerCertificateID** | string | sni协议开启时，服务器证书ID | [] | [] | No |
| **CACertificateID** | string | sni协议开启时，CA证书ID | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **PolicyID** | string | PolicyID | [] |


## 删除转发规则 - DeleteVSPolicy

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **LBID** | string | LBID | [] | [] | **Yes** |
| **VSID** | string | VSID | [] | [] | **Yes** |
| **PolicyID** | string | PolicyID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 查询转发规则 - DescribeVSPolicy

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **LBID** | string | LBID | [] | [] | **Yes** |
| **VSID** | string | VSID | [] | [] | **Yes** |
| **PolicyIDs** | array<string> | PolicyIDs | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer |  | [] |
| **Infos** | array<[*LbVSPolicyInfo*](#LbVSPolicyInfo)> |  | [] |

### 数据模型

#### LbRSInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string | 名称 | [] |
| **LBID** | string | LBID | [] |
| **VSID** | string | VSID | [] |
| **RSID** | string | RSID | [] |
| **BindResourceID** | string | 绑定的资源ID | [] |
| **HealthCheckAddress** | string | 健康检查地址 | [] |
| **IP** | string | 服务节点的内网 IP 地址 | [] |
| **Port** | integer | 服务节点暴露的服务端口号 | [] |
| **Weight** | integer | 服务节点的权重 | [] |
| **RSMode** | string | 节点模式 | [] |
| **RSStatus** | string | 节点状态 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |
| **BindResourceName** | string | 绑定的资源名称 | [] |

#### LbVSPolicyInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **LBID** | string | LBID | [] |
| **VSID** | string | VSID | [] |
| **PolicyID** | string | PolicyID | [] |
| **Domain** | string | 内容转发规则关联的请求域名 | [] |
| **Path** | string | 内容转发规则关联的请求访问路径 | [] |
| **PolicyStatus** | string | 状态 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |
| **RSInfos** | array<[*LbRSInfo*](#LbRSInfo)> | RSInfos | [] |
| **RedirectPolicyID** | string | http重定向VsPolicyId | [] |
| **SSLMode** | string | sni协议开启时，ssl模式配置 | [] |
| **ServerCertificateID** | string | sni协议开启时，服务器证书ID | [] |
| **CACertificateID** | string | sni协议开启时，CA证书ID | [] |

## 更新VS转发规则 - UpdateVSPolicy

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **LBID** | string | LBID | [] | [] | **Yes** |
| **VSID** | string | VSID | [] | [] | **Yes** |
| **PolicyID** | string | PolicyID | [] | [] | **Yes** |
| **Domain** | string | 内容转发规则关联的请求域名 | [] | [] | No |
| **Path** | string | 内容转发规则关联的请求访问路径 | [] | [] | No |
| **RSIDs** | array<string> | RSIDs | [] | [] | No |
| **SSLMode** | string | sni协议开启时，ssl模式配置 | [] | [] | No |
| **ServerCertificateID** | string | sni协议开启时，服务器证书ID | [] | [] | No |
| **CACertificateID** | string | sni协议开启时，CA证书ID | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 创建证书 - CreateCertificate

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **ProjectID** | string | ProjectID | [] | [] | No |
| **TagKeyValuePairs** | array<string> | 标签键值对 | [] | [] | **Yes** |
| **Name** | string | 名称 | [] | [] | **Yes** |
| **Remark** | string | 备注 | [] | [] | No |
| **CertificateType** | string | 证书类型 | [] | [] | **Yes** |
| **Certificate** | string | 证书内容 | [] | [] | **Yes** |
| **PrivateKey** | string | 私钥内容，证书类型为ServerCrt时必填 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **CertificateID** | string | 证书ID | [] |


## 删除证书 - DeleteCertificate

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **CertificateID** | string | 证书ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 查询证书 - DescribeCertificate

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Keyword** | string | 搜索关键字 | [] | [] | No |
| **ProjectIDs** | array<string> | 项目组ID | [] | [] | No |
| **CertificateType** | string | 证书类型 | [] | [] | No |
| **CertificateIDs** | array<string> | 证书ID | [] | [] | No |
| **LBID** | string | LBID | [] | [] | No |
| **VSID** | string | VSID | [] | [] | No |
| **Origin** | string | 证书来源，枚举值有Default/Ingress，传空表示不筛选来源 | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer |  | [] |
| **Infos** | array<[*LbCertificateInfo*](#LbCertificateInfo)> |  | [] |

### 数据模型

#### LbBindVSInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **LBID** | string | LBID | [] |
| **LBName** | string | LB名称 | [] |
| **VSID** | string | VSID | [] |
| **Protocol** | string | VS的协议 | [] |
| **Port** | integer | VS的端口 | [] |

#### LbCertificateInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Region** | string | 地域 | [] |
| **CompanyID** | integer | 租户ID | [] |
| **Email** | string | 邮箱 | [] |
| **ProjectID** | string | 项目组 | [] |
| **ProjectName** | string | 项目组名称 | [] |
| **Name** | string | 名称 | [] |
| **Remark** | string | 备注 | [] |
| **State** | string | 状态 | [] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> |  | [] |
| **VSInfos** | array<[*LbBindVSInfo*](#LbBindVSInfo)> | VSInfos | [] |
| **CertificateID** | string | 证书ID | [] |
| **CertificateType** | string | 证书类型 | [] |
| **CertificateContent** | string | 证书内容 | [] |
| **Privatekey** | string | 私钥内容 | [] |
| **CommonName** | string | 主域名 | [] |
| **SubjectAlternativeNames** | array<string> | 备域名 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **ExpireTime** | integer | 证书内容的过期时间 | [] |
| **Fingerprint** | string | 证书指纹 | [] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

## 升级LB - UpgradeLB

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **CPU** | integer | CPU核数 | [] | [] | **Yes** |
| **LBID** | string | LBID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 升级为高可用版本 - UpgradeLBToHA

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **LBID** | string | LBID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 降级LB - DowngradeLB

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **CPU** | integer | CPU核数 | [] | [] | **Yes** |
| **LBID** | string | LBID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|

