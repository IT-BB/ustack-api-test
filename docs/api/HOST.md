# HOST

## 获取物理机节点信息 - DescribeNode

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | No |
| **Keyword** | string | 搜索关键词 | [example_key] | [] | No |
| **Limit** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 默认值：0 | [0] | [] | No |
| **Offset** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。默认值：10 | [20] | [] | No |
| **HostIDs** | array<string> | 宿主机ID列表，用于查询指定的节点信息，多个元素之间用,分隔。 | [10.10.1.135] | [DescribeVMHost] | No |
| **SetID** | string | 集群ID，用于查询指定计算集群下的节点信息 | [compute-set-j1jg] | [DescribeVMSet] | No |
| **NodeType** | string | 节点类型，用于查询指定类型的节点信息。取值范围：1、manager 2、compute 3、storage | [manager] | [] | No |
| **SortBy** | string | 排序字段，目前仅支持通过NodeIP进行排序 | [NodeIP] | [] | No |
| **Sort** | string | 排序，取值范围：1、Descending（降序） 2、Ascending（升序） | [Descending] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 查询到的总记录数 | [] |
| **Infos** | array<[*NodeInfo*](#NodeInfo)> | 查询到的节点信息详情 | [] |

### 数据模型

#### CPU

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Manufacturer** | string | 制造商 | [] |
| **Name** | string | 名称 | [] |
| **Sockets** | integer | 插槽数 | [] |
| **Cores** | integer | 核数 | [] |
| **Threads** | integer | 线程数 | [] |

#### Compute

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **SetID** | string | 集群ID | [] |
| **SetType** | string | 集群类型 | [] |
| **Weight** | integer | 权重 | [] |
| **CPUType** | string | CPU类型 | [] |
| **CPUCount** | integer | CPU数量 | [] |
| **CPUUsed** | integer | CPU已用 | [] |
| **CPUAllocable** | integer | CPU可分配 | [] |
| **MemoryCap** | integer | 内存总量 | [] |
| **MemoryUsed** | integer | 内存已用 | [] |
| **MemoryAllocable** | integer | 内存可分配 | [] |
| **MemoryPhysicUsed** | integer | 物理内存已用 | [] |
| **MemoryReserved** | integer | 内存预留量 | [] |
| **GPUType** | string | GPU类型 | [] |
| **GPUCount** | integer | GPU数量 | [] |
| **GPUUsed** | integer | GPU已用 | [] |
| **GPUAllocable** | integer | GPU可分配 | [] |
| **VGPUInfos** | array<[*HostVGPUUsedInfo*](#HostVGPUUsedInfo)> | vGPU详情 | [] |
| **GPUInfos** | array<[*HostGPUUsedInfo*](#HostGPUUsedInfo)> | GPU详情 | [] |
| **VGPUDriverNotReady** | bool | vGPU驱动是否就绪 | [] |
| **ResourceCount** | integer | 资源总数 | [] |

#### HostGPUUsedInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **HDName** | string | GPU硬件设备名称 | [] |
| **MdevName** | string | GPU虚拟设备名称 | [] |
| **Count** | integer | GPU总量 | [] |
| **Used** | integer | 已用GPU数量 | [] |
| **Allocable** | integer | 可用GPU数量 | [] |

#### HostVGPUUsedInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **MdevName** | string | VGPU虚拟设备名称 | [] |
| **Count** | integer | VGPU总量 | [] |
| **Used** | integer | 已用VGPU数量 | [] |
| **Allocable** | integer | 可用VGPU数量 | [] |

#### IP

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Address** | string | IP地址 | [] |
| **Mask** | string | 掩码 | [] |

#### Location

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Drive** | string | 物理设备名 | [] |
| **RAIDLevel** | string | raid级别 | [] |
| **ControllerID** | string | raid卡的controller ID | [] |
| **Enclosure** | string | 磁盘柜编号 | [] |
| **Slot** | string | 磁盘插槽号 | [] |
| **Status** | string | 磁盘状态 | [] |
| **Type** | string | 磁盘类型，取值范围：1、HDD 2、SSD | [] |
| **Model** | string | 磁盘型号 | [] |
| **SerialNumber** | string | 磁盘序列号 | [] |
| **Size** | integer | 磁盘大小，单位（GiB） | [] |

#### LogicalVolume

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **LVName** | string | 逻辑卷名 | [] |
| **LVDMPath** | string | 逻辑卷路径 | [] |
| **LVTAGs** | string | 逻辑卷标签 | [] |
| **VGName** | string | 卷组名 | [] |
| **LVMSize** | integer | 逻辑卷大小 | [] |

#### Memory

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalSize** | integer | 总内存大小，单位（GiB） | [] |

#### NIC

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string | 网卡名称 | [] |
| **MAC** | string | MAC地址 | [] |
| **IPs** | array<[*IP*](#IP)> | IP详情 | [] |

#### NICStatus

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Linked** | bool | 连接状态 | [] |
| **Speed** | string | 实际速率 | [] |
| **Duplex** | string | 双工模式 | [] |
| **AutoNegotiation** | bool | 是否开启自动协商 | [] |
| **Mac** | string | 网卡mac | [] |
| **EnableBonding** | bool | 是否开启网卡绑定 | [] |

#### NTPInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Status** | string | NTP时间是否已同步 | [] |
| **Servers** | array<string> | NTP上游服务器列表 | [] |

#### Network

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **NICs** | array<[*NIC*](#NIC)> | 虚拟网卡详情 | [] |
| **PhyProductInfos** | array<[*PhyProductInfo*](#PhyProductInfo)> | 物理网卡详情 | [] |

#### NodeInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [] |
| **NodeID** | string | 节点ID | [] |
| **NodeIP** | string | 节点IP地址 | [] |
| **NodeIPv6** | string | 节点IPv6地址 | [] |
| **NodeStatus** | string | 节点状态 | [] |
| **SerialNumber** | string | 序列号 | [] |
| **Arch** | string | 架构 | [] |
| **CPU** | object | CPU详情 | [] |
| **Memory** | object | 内存详情 | [] |
| **PhysicalDisks** | array<[*PhysicalDisk*](#PhysicalDisk)> | 物理磁盘详情 | [] |
| **Network** | object | 主机网络详情 | [] |
| **NTP** | object | NTP详情 | [] |
| **Compute** | object | 资源用量情况（计算节点独有） | [] |
| **Types** | array<string> | 节点类型 | [] |
| **OSDInfos** | array<[*OSDStat*](#OSDStat)> | 存储节点独有 OSD 详情 | [] |
| **UUID** | string | UUID | [] |
| **VCPUBindingEnabled** | bool | 是否支持vcpu绑定 | [] |
| **NumaNodes** | integer | 节点numa数量 | [] |

#### OSDStat

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **OSDID** | string | OSD ID | [] |
| **OSDName** | string | OSD名称 | [] |
| **HostIP** | string | 宿主机IP | [] |
| **DiskType** | string | 磁盘类型 | [] |
| **Status** | string | 状态 | [] |
| **Size** | integer | 大小 | [] |
| **Used** | integer | 使用量 | [] |
| **HostName** | string | 主机名 | [] |
| **StorageType** | string | 存储类型 | [] |

#### PFInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **PCI** | string | PCI位置 | [] |
| **Vendor** | string | 网卡厂商 | [] |
| **Product** | string | 网卡型号 | [] |
| **Code** | string | 网卡型号的标准编号 | [] |
| **DeviceName** | string | 网卡名 | [] |
| **Driver** | string | 驱动 | [] |
| **VFPhyCount** | integer | 物理限制的VF数量 | [] |
| **VFLogicCount** | integer | 逻辑限制的VF数量，VFLogicCount <= VFPhyCount，如果VFLogicCount不为0，则以VFLogicCount为准，否则VFPhyCount | [] |
| **NICStatus** | object | 网卡状态 | [] |
| **UsedVFs** | array<[*VFInfo*](#VFInfo)> | 已使用的VF网卡 | [] |

#### PhyProductInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Vendor** | string | 网卡厂商 | [] |
| **Product** | string | 网卡型号 | [] |
| **Code** | string | 网卡型号的标准编号 | [] |
| **PFInfos** | array<[*PFInfo*](#PFInfo)> | 节点下的所有同型号的网卡 | [] |
| **ReservedVFs** | array<[*VFInfo*](#VFInfo)> | 节点在该型号下的预占VF，表示一些已经占用了VF名额的网卡，但不确定属于哪一个PF。例如绑定中的VF卡、虚拟机Eth0没有绑定IP的VF卡 | [] |

#### PhysicalDisk

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **SerialNumber** | string | 设备序列号 | [] |
| **WWID** | string | WWID（全球唯一磁盘标识符） | [] |
| **Arch** | string | 存储架构，取值范围：1、HDD 2、SSD | [] |
| **Drive** | string | 设备名 | [] |
| **Partitions** | array<[*PhysicalDiskPartition*](#PhysicalDiskPartition)> | 磁盘分区 | [] |
| **Size** | integer | 磁盘大小，单位（GiB）  | [] |
| **LVs** | array<[*LogicalVolume*](#LogicalVolume)> | 逻辑卷 | [] |
| **Location** | array<[*Location*](#Location)> | 真实物理磁盘厂商位置等信息 | [] |
| **HealthStatus** | string | 磁盘健康状态，取值范围：1、healthy 2、unhealthy | [] |
| **MountStatus** | string | 磁盘挂载状态，取值范围：1、mounted 2、mounting 3、mount_failed 4、unmounted 5、unmounting 6、unmount_failed | [] |
| **DiskUsage** | string | 磁盘用途，取值范围：1、boot 2、cache 3、data | [] |

#### PhysicalDiskPartition

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Partition** | string | 分区名 | [] |
| **Mountpoint** | string | 挂载点 | [] |
| **Size** | integer | 大小，单位（GiB） | [] |
| **Used** | integer | 已用大小，单位（GiB） | [] |

#### VFInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **PCI** | string | PCI位置 | [] |
| **RelatedVM** | string | 关联虚拟机 | [] |
| **RelatedNIC** | string | 关联网卡 | [] |

## 获取虚拟机物理机信息 - DescribeVMHost

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **Keyword** | string | 搜索关键词 | [example_key] | [] | No |
| **Limit** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 默认值：0 | [0] | [] | No |
| **Offset** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。默认值：10 | [20] | [] | No |
| **HostIDs** | array<string> | 物理机ID列表，用于查询指定的节点信息，多个元素之间用,分隔。 | [10.10.1.135] | [DescribeVMHost] | No |
| **SetID** | string | 计算集群ID，用于查询指定集群的节点 | [compute-set-j1jg] | [DescribeVMSet] | No |
| **FilterRunnableHostsByVMID** | string | 虚拟机ID，用于筛选支持某个虚拟机的宿主机 | [vm-lac3eb467pqywq] | [DescribeVMInstance] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 查询到的总记录数 | [] |
| **Infos** | array<[*HostInfo*](#HostInfo)> | 物理宿主机信息详情 | [] |

### 数据模型

#### HostGPUUsedInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **HDName** | string | GPU硬件设备名称 | [] |
| **MdevName** | string | GPU虚拟设备名称 | [] |
| **Count** | integer | GPU总量 | [] |
| **Used** | integer | 已用GPU数量 | [] |
| **Allocable** | integer | 可用GPU数量 | [] |

#### HostInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [] |
| **HostID** | string | 物理机ID | [] |
| **SetID** | string | 计算集群ID | [] |
| **SetType** | string | 集群类型 | [] |
| **HostIP** | string | 物理机IP地址 | [] |
| **HostIPv6** | string | 物理机IPv6地址 | [] |
| **HostStatus** | string | 物理机状态 | [] |
| **Weight** | integer | 负载均衡权重 | [] |
| **CPUType** | string | CPU类型 | [] |
| **CPUCount** | integer | CPU总数 | [] |
| **CPUUsed** | integer | 已用CPU数量 | [] |
| **CPUAllocable** | integer | 可用CPU数量 | [] |
| **MemoryCap** | integer | 总内存容量（GiB） | [] |
| **MemoryUsed** | integer | 已用内存（GiB） | [] |
| **MemoryAllocable** | integer | 可用内存（GiB） | [] |
| **MemoryPhysicUsed** | integer | 物理机使用的内存（GiB） | [] |
| **MemoryReserved** | integer | 物理机预留的内存（GiB） | [] |
| **GPUType** | string | GPU类型 | [] |
| **GPUCount** | integer | GPU总量 | [] |
| **GPUUsed** | integer | 已用GPU数量 | [] |
| **GPUAllocable** | integer | 可用GPU数量 | [] |
| **VGPUInfos** | array<[*HostVGPUUsedInfo*](#HostVGPUUsedInfo)> | 虚拟GPU信息 | [] |
| **GPUInfos** | array<[*HostGPUUsedInfo*](#HostGPUUsedInfo)> | 物理机已用GPU信息 | [] |
| **CreateTime** | integer | 创建时间，Unix 时间戳（秒级） | [] |
| **UpdateTime** | integer | 更新时间，Unix 时间戳（秒级） | [] |
| **VGPUDriverNotReady** | bool | 虚拟GPU驱动未就绪 | [] |
| **ResourceCount** | integer | 资源数量，运行在该物理机上的资源数量 | [] |
| **VCPUBindingEnabled** | bool | 是否启用VCPU绑定 | [] |
| **NumaNodes** | integer | NUMA节点数量 | [] |

#### HostVGPUUsedInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **MdevName** | string | VGPU虚拟设备名称 | [] |
| **Count** | integer | VGPU总量 | [] |
| **Used** | integer | 已用VGPU数量 | [] |
| **Allocable** | integer | 可用VGPU数量 | [] |

## 锁定物理机 - LockHost

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **MemberID** | integer | 用户账号ID，标识登录用户的唯一身份标识符 | [200000234] | [DescribeMember] | No |
| **HostID** | string | 物理机ID | [10.10.1.135] | [DescribeVMHost] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 解锁物理机 - UnlockHost

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **MemberID** | integer | 用户账号ID，标识登录用户的唯一身份标识符 | [200000234] | [DescribeMember] | No |
| **HostID** | string | 物理机ID | [10.10.1.135] | [DescribeVMHost] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取物理机上虚拟机信息 - DescribeHostVMInstance

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **Keyword** | string | 搜索关键词 | [example_key] | [] | No |
| **Limit** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。默认值：10 | [20] | [] | No |
| **Offset** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 默认值：0 | [0] | [] | No |
| **HostID** | string | 物理机ID，用于查询指定物理机上的虚拟机信息 | [10.10.1.135] | [DescribeVMHost] | No |
| **SetID** | string | 计算集群ID，用于查询指定集群上的虚拟机信息 | [compute-set-j1jg] | [DescribeVMSet] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 查询到的总记录数 | [] |
| **Infos** | array<[*CIInfo*](#CIInfo)> | 计算实例的信息详情 | [] |

### 数据模型

#### CIInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [] |
| **RegionAlias** | string | 地域别名 | [] |
| **Name** | string | 虚拟机名称 | [] |
| **Remark** | string | 虚拟机描述 | [] |
| **Email** | string | 所属的租户邮箱 | [] |
| **BindResourceType** | string | 绑定的资源类型 | [] |
| **CompanyID** | integer | 所属的租户ID | [] |
| **VMID** | string | 虚拟机ID | [] |
| **SetArch** | string | 计算集群架构 | [] |
| **SetID** | string | 计算集群ID | [] |
| **HostID** | string | 物理机ID | [] |
| **HostIP** | string | 物理机IP地址 | [] |
| **HostIPv6** | string | 物理机IPv6地址 | [] |
| **Usage** | string | 虚拟机用途 | [] |
| **ImageID** | string | 虚拟机使用的镜像ID | [] |
| **CPU** | integer | CPU 核数 | [] |
| **Memory** | integer | 内存大小，单位（MiB） | [] |
| **GPU** | integer | GPU 数量 | [] |
| **Status** | string | 虚拟机运行状态 | [] |
| **CreateTime** | integer | 创建时间，Unix 时间戳（秒级） | [] |
| **BindResourceID** | string | 绑定资源的 ID | [] |
| **InternalIP** | string | 内网IP地址 | [] |
| **VGPUID** | string | 虚拟GPU的ID | [] |
| **MdevName** | string | VGPU虚拟设备名称 | [] |
| **GPUMdevName** | string | GPU虚拟设备名称 | [] |
| **ImageName** | string | 镜像的名称 | [] |
| **BindResourceName** | string | 绑定资源的名称 | [] |
| **CanMigrateAbort** | bool | 是否可以取消迁移 | [] |
| **VCPUBindingNode** | string | VCPU绑定的物理节点 | [] |

## 获取物理机上Pod信息 - DescribeHostPods

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **Keyword** | string | 搜索关键词 | [example_key] | [] | No |
| **Limit** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。默认值：10 | [20] | [] | No |
| **Offset** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 默认值：0 | [0] | [] | No |
| **HostID** | string | 物理机ID，用于查询指定物理机上的虚拟机信息 | [10.10.1.135] | [DescribeVMHost] | No |
| **SetID** | string | 计算集群ID，用于查询指定集群上的虚拟机信息 | [compute-set-j1jg] | [DescribeVMSet] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 查询到的总记录数 | [] |
| **Infos** | array<[*PodInfo*](#PodInfo)> | Pod的信息详情 | [] |

### 数据模型

#### ContainerInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string | container的名称 | [] |
| **CPURequest** | integer | CPU资源的下限，单位：核数 | [] |
| **MemoryRequest** | integer | 存储资源的下限，单位：GiB | [] |
| **CPULimit** | integer | CPU资源的上限，单位：核数 | [] |
| **MemoryLimit** | integer | 存储资源的上限，单位：GiB | [] |

#### PodInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string | Pod 的名称 | [] |
| **NameSpace** | string | 所属的命名空间 | [] |
| **Phase** | string | Pod 的状态 | [] |
| **PodIP** | string | Pod 的 IP 地址 | [] |
| **HostIP** | string | 物理机IP地址 | [] |
| **HostID** | string | 物理机ID | [] |
| **CPURequest** | integer | CPU资源的下限，单位：核数 | [] |
| **MemoryRequest** | integer | 存储资源的下限，单位：GiB | [] |
| **CPULimit** | integer | CPU资源的上限，单位：核数 | [] |
| **MemoryLimit** | integer | 存储资源的上限，单位：GiB | [] |
| **Containers** | array<[*ContainerInfo*](#ContainerInfo)> | container 详情信息 | [] |
| **CreateTime** | integer | 创建时间，Unix 时间戳（秒级） | [] |

## 虚机迁移 - MigrateVMInstance

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **CIID** | string | 计算实例ID | [vm-6t5c7jlz0zdq3r] | [DescribeHostVMInstance] | **Yes** |
| **SetID** | string | 计算集群ID，源计算实例所属的计算集群ID | [compute-set-j1jg] | [DescribeVMSet] | No |
| **HostIP** | string | 物理机IP地址，仅在线迁移时需要指定 | [10.10.1.137] | [DescribeVMHostRequest] | No |
| **MigrateHostIP** | string | 目标宿主IP地址 | [10.10.1.135] | [DescribeVMHostRequest] | No |
| **MigrateSetID** | string | 目标集群ID | [compute-set-j1jg] | [DescribeVMSet] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 取消虚机迁移 - AbortMigrateVMInstance

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **VMID** | string | 虚拟机ID | [vm-6t5c7jlz0zdq3r] | [DescribeHostVMInstance] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 磁盘点灯 - DiskLightOn

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **HostIP** | string | 物理机IP地址 | [10.10.1.137] | [DescribeVMHostRequest] | No |
| **ControllerID** | string | raid卡的controllerID | [1] | [] | No |
| **Enclosure** | string | 磁盘柜编号 | [1] | [] | No |
| **Slot** | string | 磁盘插槽号 | [1] | [] | No |
| **Duration** | integer | 持续时长，单位：秒 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 磁盘关灯 - DiskLightOff

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **HostIP** | string | 物理机IP地址 | [10.10.1.137] | [DescribeVMHostRequest] | No |
| **ControllerID** | string | raid卡的controllerID | [1] | [] | No |
| **Enclosure** | string | 磁盘柜编号 | [1] | [] | No |
| **Slot** | string | 磁盘插槽号 | [1] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取节点 CPU 电源模式 - GetNodeCPUGovernor

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **HostID** | string | 节点ID | [10.10.1.135] | [DescribeVMHost] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Governor** | string | ​​当前生效的电源模式 | [] |
| **AvailableGovernors** | array<string> | 支持的模式列表​ | [] |
| **Infos** | array<[*CPUPowerInfo*](#CPUPowerInfo)> | 每个CPU核心的频率信息​ | [] |

### 数据模型

#### CPUPowerInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Driver** | string | CPU频率驱动类型​​ | [] |
| **CurrentFrequency** | string | CPU核心的当前运行频率​ | [] |

## 更新节点 CPU 电源模式 - UpdateNodeCPUGovernor

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **HostID** | string | 节点ID | [10.10.1.135] | [DescribeVMHost] | **Yes** |
| **Governor** | string | 电源模式，取值范围：1、performance（性能，最高频率） 2、powersave（节能，最低频率） 3、ondemand（自动，部分机器不支持） | [performance] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 调整逻辑VF数量限制 - UpdateVFLogicCount

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **HostIP** | string | 节点IP地址 | [DescribeVMHost] | [] | No |
| **PCI** | string | 网卡PCI地址 | [0000:03:00.0] | [DescribeNode] | No |
| **VFLogicCount** | integer | 逻辑VF数量限制，注意VFLogicCount <= 物理VF数量限制，VFLogicCount >= 已使用VF数量 | [2] | [DescribeNode] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 开启节点NUMA调度 - OpenHostNUMASchedule

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **HostID** | string | 物理机ID | [10.10.1.135] | [DescribeVMHost] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 关闭节点NUMA调度 - CloseHostNUMASchedule

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **HostID** | string | 物理机ID | [10.10.1.135] | [DescribeVMHost] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取节点NUMANode信息 - DescribeNodeNUMAInfo

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | No |
| **SetID** | string | 计算集群ID | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 查询到的总记录数 | [] |
| **Infos** | array<[*PCPUBindingInfo*](#PCPUBindingInfo)> | 详情 | [] |

### 数据模型

#### NUMANode

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **NUMAID** | integer |  | [] |
| **CPUs** | array<[*NUMANodeCPU*](#NUMANodeCPU)> | NUMANode cpu 列表 | [] |

#### NUMANodeCPU

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **CPUID** | integer |  | [] |
| **RunningVMs** | array<[*NUMANodeVMInfo*](#NUMANodeVMInfo)> |  | [] |
| **NotRunningVMs** | array<[*NUMANodeVMInfo*](#NUMANodeVMInfo)> |  | [] |
| **Reserved** | bool |  | [] |

#### NUMANodeVMInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **VMID** | string |  | [] |
| **Name** | string |  | [] |

#### PCPUBindingInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **NodeID** | string | 节点ID | [] |
| **NUMANodes** | array<[*NUMANode*](#NUMANode)> | CPU详情 | [] |
