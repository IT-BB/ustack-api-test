# K8S

## 代理k8s集群请求 - ForwardCluster

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | Region | [] | [] | **Yes** |
| **ClusterID** | string | 集群名称 | [] | [] | **Yes** |
| **Language** | string | 语言 | [] | [] | No |
| **Method** | string | 请求类型 | [] | [] | No |
| **Path** | string | 请求路径 | [] | [] | No |
| **Version** | string | 请求版本 | [] | [] | No |
| **RequestBody** | string | 转发Body数据 | [] | [] | No |
| **EncodedBody** | bool | 转发Body数据是否被编码 | [] | [] | No |
| **Accept** | string | 响应体格式 | [] | [] | No |
| **ContentType** | string | 请求体格式 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Info** | object | 响应信息 | [] |

### 数据模型

#### ResponseInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ResponseBody** | string | 响应体 | [] |
| **ResponseCode** | integer | 响应码 | [] |

## 申请console会话 - AllocateK8SSession

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ClusterID** | string | 集群Id | [] | [] | **Yes** |
| **NameSpace** | string | 名字空间 | [] | [] | No |
| **PodName** | string | Pod名称 | [] | [] | **Yes** |
| **ContainerName** | string | 容器名称 | [] | [] | No |
| **CmdName** | string | 执行命令 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **SessionID** | string |  | [] |


## 申请console会话 - AllocateK8STerminal

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ClusterID** | string | 集群Id | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **SessionID** | string |  | [] |


## 获取K8S价格 - GetClusterPrice

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **ChargeType** | string | 计费类型 | [] | [] | **Yes** |
| **Quantity** | integer | 计费周期 | [] | [] | No |
| **ComputeclassType** | string | 计算集群类型 | [] | [] | **Yes** |
| **StorageclassType** | string | 存储集群类型 | [] | [] | **Yes** |
| **CPU** | integer | CPU核数 | [] | [] | No |
| **Memory** | integer | 内存大小 | [] | [] | No |
| **HighAvailability** | string | 高可用 | [] | [] | No |
| **ClusterID** | string | ClusterID | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*K8SPriceInfo*](#K8SPriceInfo)> |  | [] |

### 数据模型

#### K8SPriceInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Price** | double | 价格 | [] |
| **ChargeType** | string | 计费类型 | [] |

## 创建k8s集群 - CreateCluster

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Name** | string | 名称 | [] | [] | **Yes** |
| **Remark** | string | 备注 | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ProjectID** | string | 项目组ID | [] | [] | No |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **TagKeyValuePairs** | array<string> | 标签键值对 | [] | [] | **Yes** |
| **ChargeType** | string | 计费类型 | [] | [] | **Yes** |
| **Quantity** | integer | 计费数量 | [] | [] | **Yes** |
| **ApplicationName** | string | 审批名称 | [] | [] | No |
| **ApplicationReason** | string | 审批理由 | [] | [] | No |
| **VpcID** | string | VpcID | [] | [] | **Yes** |
| **SubnetID** | string | 子网Id | [] | [] | **Yes** |
| **ServiceCIDR** | string | Service CIDR 或 Serivce 所在子网Id | [] | [] | No |
| **EIPID** | string | APIServer EIPID | [] | [] | No |
| **K8SVersion** | string | k8s版本号。可为1.25.0 | [] | [] | **Yes** |
| **StorageclassType** | string | 存储集群类型 | [] | [] | **Yes** |
| **ComputeclassType** | string | 计算集群类型 | [] | [] | **Yes** |
| **HighAvailability** | string | 高可用类型 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ClusterID** | string | 集群Id | [] |


## 查询k8s集群 - DescribeCluster

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ClusterIDs** | array<string> | 集群Id | [] | [] | No |
| **ProjectIDs** | array<string> | 项目组ID | [] | [] | No |
| **Keyword** | string | 关键词 | [] | [] | No |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数 | [] |
| **Infos** | array<[*ClusterInfo*](#ClusterInfo)> | 详情 | [] |

### 数据模型

#### ClusterInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string | 名称 | [] |
| **Remark** | string | 备注 | [] |
| **Region** | string | 地域 | [] |
| **RegionAlias** | string | 地域别名 | [] |
| **CompanyID** | integer | 租户ID | [] |
| **CompanyName** | string | 租户名称 | [] |
| **Email** | string | 租户邮箱 | [] |
| **ProjectID** | string | 项目组ID | [] |
| **ProjectName** | string | 项目组名称 | [] |
| **Reason** | string | 失败原因 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |
| **ExpireTime** | integer | 过期时间 | [] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> | 标签 | [] |
| **ChargeType** | string | 计费周期 | [] |
| **ClusterID** | string | 集群Id | [] |
| **VpcID** | string | VpcID | [] |
| **VpcName** | string | VPC名称 | [] |
| **SubnetIDs** | array<string> | 子网Id 列表 | [] |
| **SubnetNames** | array<string> | 子网名称列表 | [] |
| **ServiceCIDR** | string | Service CIDR 或 Serivce 所在子网Id | [] |
| **K8SVersion** | string | k8s版本号。可为1.20.6。 | [] |
| **StorageclassType** | string | 存储集群类型 | [] |
| **ComputeclassType** | string | 计算集群类型 | [] |
| **Status** | string | 状态 | [] |
| **State** | string | 阶段 | [] |
| **Addresses** | array<[*IPAddress*](#IPAddress)> | apiserver接入地址列表 | [] |
| **Kubeconfig** | string | 连接用户集群k8s 的Config | [] |
| **KubeconfigWan** | string | 连接用户集群k8s 的外网Config | [] |
| **Registry** | string | Registry地址 | [] |
| **APIServerEIPID** | string | 集群EIP ID | [] |
| **MaxCPU** | integer | 集群最大CPU数量 | [] |
| **MaxMemory** | integer | 集群最大内存容量 | [] |
| **MaxPods** | integer | 集群最大pod数量 | [] |
| **MaxStorage** | integer | 集群最大存储容量 | [] |
| **MaxStoragAllocation** | string | 集群最大存储容量分配策略 | [] |
| **AllocatedCPU** | integer | 集群已分配CPU数量 | [] |
| **AllocatedMemory** | integer | 集群已分配内存容量 | [] |
| **AllocatedPods** | integer | 集群已分配pod数量 | [] |
| **ControllerPlaneStatus** | object | 集群控制面状态 | [] |

#### ComponentStatus

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Status** | string | 组件状态 | [] |
| **Details** | array<[*ContainerStatus*](#ContainerStatus)> | 组件列表 | [] |

#### ContainerStatus

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **PodName** | string | pod名 | [] |
| **Container** | string | 容器名 | [] |
| **Reason** | string | 状态或原因 | [] |
| **Ready** | bool | 是否就绪 | [] |

#### ControllerPlaneStatus

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ETCD** | object | 集群ETCD状态 | [] |
| **KubeApiserver** | object | 集群KubeApiserver状态 | [] |
| **KubeTerminal** | object | 集群Kubeutil状态 | [] |
| **KubeControllerManager** | object | 集群KubeControllerManager状态 | [] |
| **KubeScheduler** | object | 集群KubeScheduler状态 | [] |
| **CSI** | object | 集群存储插件状态 | [] |
| **CCM** | object | 集群svc插件状态 | [] |
| **VK** | object | 集群默认超级节点状态 | [] |

#### IPAddress

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Type** | string | 地址类型，可为 advertise, public 等 | [] |
| **Ip** | string | Ip地址 | [] |
| **Port** | integer | 端口 | [] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

## 更新k8s集群 - UpdateCluster

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | Region | [] | [] | **Yes** |
| **ClusterID** | string | 集群Id | [] | [] | **Yes** |
| **SubnetIDs** | array<string> | 子网Id 列表 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 绑定k8s集群外网IP - AttachClusterEIP

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | Region | [] | [] | **Yes** |
| **ClusterID** | string | 集群Id | [] | [] | **Yes** |
| **EIPID** | string | APIServer EIPID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 解绑k8s集群外网IP - DetachClusterEIP

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | Region | [] | [] | **Yes** |
| **ClusterID** | string | 集群Id | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 更新k8s集群容量配额 - UpdateClusterCapacity

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | Region | [] | [] | **Yes** |
| **ClusterID** | string | 集群Id | [] | [] | **Yes** |
| **MaxCPU** | integer | 集群最大CPU数量 | [] | [] | No |
| **MaxMemory** | integer | 集群最大内存容量 | [] | [] | No |
| **MaxPods** | integer | 集群最大内存容量 | [] | [] | No |
| **MaxStorage** | integer | 集群最大存储容量 | [] | [] | No |
| **MaxStorageAllocation** | string | 集群最大存储容量分配策略 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 删除k8s集群 - DeleteCluster

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | Region | [] | [] | **Yes** |
| **ClusterID** | string | 集群Id | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 创建k8s集群SuperNode - CreateSuperNode

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ClusterID** | string | 集群Id | [] | [] | **Yes** |
| **NodeName** | string | 节点名称 | [] | [] | **Yes** |
| **StorageclassType** | string | 存储集群类型 | [] | [] | No |
| **ComputeclassType** | string | 计算集群类型 | [] | [] | **Yes** |
| **MaxCPU** | integer | 节点最大CPU数量 | [] | [] | No |
| **MaxMemory** | integer | 节点最大内存容量 | [] | [] | No |
| **MaxPods** | integer | 节点最大pod数量 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **NodeName** | string | 节点名称 | [] |


## 查询k8s集群Node - DescribeSuperNode

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ClusterID** | string | 集群Id | [] | [] | **Yes** |
| **NodeNames** | array<string> | 节点名称列表 | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数 | [] |
| **Infos** | array<[*SuperNodeInfo*](#SuperNodeInfo)> | 详情 | [] |

### 数据模型

#### SuperNodeGPUInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **GPUMdevName** | string | GPU的规格 | [] |
| **GPU** | integer | GPU的颗数 | [] |

#### SuperNodeInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ClusterID** | string | 集群Id | [] |
| **NodeName** | string | 节点名称 | [] |
| **VpcID** | string | VpcID | [] |
| **VpcName** | string | VpcName | [] |
| **SubnetIDs** | array<string> | 子网Id 列表 | [] |
| **SubnetNames** | array<string> | 子网名称列表 | [] |
| **StorageclassType** | string | 存储集群类型 | [] |
| **ComputeclassType** | string | 计算集群类型 | [] |
| **OSType** | string | 系统类型 | [] |
| **ArchType** | string | 架构类型 | [] |
| **Status** | string | 状态 | [] |
| **Unschedulable** | bool | 不可调度 | [] |
| **IsDefault** | bool | 默认节点 | [] |
| **YamlData** | string | yaml数据 | [] |
| **MaxCPU** | integer | 节点最大CPU数量 | [] |
| **MaxMemory** | integer | 节点最大内存容量 | [] |
| **MaxPods** | integer | 节点最大pod数量 | [] |
| **UsedCPU** | float | 节点已使用CPU数量 | [] |
| **UsedMemory** | float | 节点已使用内存容量 | [] |
| **UsedPods** | float | 节点已使用pod数量 | [] |
| **AdjustableMaxCPU** | integer | 节点可调最大CPU数量 | [] |
| **AdjustableMaxMemory** | integer | 节点可调最大内存容量 | [] |
| **AdjustableMaxPods** | integer | 节点可调最大pod数量 | [] |
| **GPUInfos** | array<[*SuperNodeGPUInfo*](#SuperNodeGPUInfo)> | 节点GPU信息 | [] |

## 更新k8s集群SuperNode - UpdateSuperNode

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ClusterID** | string | 集群Id | [] | [] | **Yes** |
| **NodeName** | string | 节点名称 | [] | [] | **Yes** |
| **MaxCPU** | integer | 节点最大CPU数量 | [] | [] | No |
| **MaxMemory** | integer | 节点最大内存容量 | [] | [] | No |
| **MaxPods** | integer | 节点最大pod数量 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 删除k8s集群SuperNode - DeleteSuperNode

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ClusterID** | string | 集群Id | [] | [] | **Yes** |
| **NodeName** | string | 节点名称 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 查询容器日志 - GetContainerLogs

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | Region | [] | [] | **Yes** |
| **ClusterID** | string | 集群id | [] | [] | **Yes** |
| **Namespace** | string | 命名空间 | [] | [] | **Yes** |
| **Pod** | string | pod名字 | [] | [] | **Yes** |
| **Container** | string | 容器名 | [] | [] | **Yes** |
| **Keywords** | array<string> | 关键词 | [] | [] | No |
| **KeywordsIgnore** | array<string> | 关键词排除 | [] | [] | No |
| **LineLimit** | integer | 数量限制 | [] | [] | No |
| **Start** | integer | 起始时间,unixnano | [] | [] | No |
| **End** | integer | 结束时间,unixnano | [] | [] | No |
| **NotInCluster** | bool | 是否在集群内 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数 | [] |
| **Infos** | array<[*ContainerLogInfo*](#ContainerLogInfo)> | 详情 | [] |

### 数据模型

#### ContainerLogInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Time** | integer | 时间 | [] |
| **Body** | string | 内容 | [] |

## 更新k8s集群SuperNodeGPU - UpdateSuperNodeGPU

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ClusterID** | string | 集群Id | [] | [] | **Yes** |
| **NodeName** | string | 节点名称 | [] | [] | **Yes** |
| **GPUMdevName** | string | GPU的规格 | [] | [] | **Yes** |
| **GPU** | integer | GPU的颗数 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取K8S修改配置后的差价 - GetClusterPaymentOfPremium

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **ClusterID** | string | ClusterID | [] | [] | No |
| **SuperNode** | string | 超级节点名称 | [] | [] | No |
| **GPUMdevName** | string | GPU的规格 | [] | [] | No |
| **GPU** | integer | GPU的颗数 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Price** | double | 价格 | [] |
| **OrderType** | string | 计费类型 | [] |

