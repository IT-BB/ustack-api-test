# REGION

## 纳管新地域 - AddRegion

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **RegionID** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [] | **Yes** |
| **RegionToken** | string | 地域令牌 | [] | [] | **Yes** |
| **RegionIP** | string | 地域IP | [] | [] | **Yes** |
| **CMPIP** | string | CMP IP | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 修改租户地域授权 - UpdateCompanyRegion

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [CreateUser DescribeUser DescribeMember] | **Yes** |
| **Regions** | array<string> | 租户地域列表，用于指定租户账号在哪些地域下有操作权限，实现对多个地域的授权管理。 | [cn pro2134] | [DescribeRegion] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 修改管理员地域授权 - UpdateAdminRegion

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **MemberID** | integer | 管理员ID，标识管理员账号的唯一标识符，用于指定要修改的管理员。 | [200000237] | [ListAdmin] | **Yes** |
| **Regions** | array<string> | 管理员地域列表，用于指定管理员账号在哪些地域下有操作权限。 | [cn develop] | [DescribeRegion] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取租户已授权地域 - DescribeRegion

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [CreateUser DescribeUser DescribeMember] | No |
| **MemberID** | integer | 账号ID，用于标识资源所属的账号，实现多账号环境下的资源隔离。 | [200000237] | [CreateSubMember DescribeMember GetMemberInfo] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*RegionInfo*](#RegionInfo)> | 地域的详细信息列表 | [] |
| **TotalCount** | integer | 当前账号地域的总数 | [] |

### 数据模型

#### RegionInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Region** | string | 地域标识，用于唯一标识一个地理区域，通常由字母和数字组成，用于API调用时指定操作的目标地域 | [cn] |
| **RegionAlias** | string | 地域别名，地域的友好显示名称，便于用户识别和管理，通常使用中文或英文描述地域的地理位置 | [北京地域] |
| **EndPoints** | string | 管理端点地址，用于连接和管理该地域的基础设施服务，包含协议、IP地址和端口信息 | [] |
| **City** | string | 地域所在城市名称，标识该数据中心的地理位置，用于地理位置相关的资源调度和合规要求 | [北京] |
| **Address** | string | 地域数据中心的详细物理地址，包含具体的街道、建筑物等信息，用于运维和合规管理 | [上海市杨浦区xxxx] |
| **Remark** | string | 地域的备注信息，用于记录该地域的特殊说明、用途描述或其他重要信息 | [主要生产环境数据中心] |
| **CPUCount** | integer | CPU总核心数，表示该地域所有物理服务器的CPU核心总数，用于资源容量规划和监控 | [190] |
| **CPUUsed** | integer | CPU已使用核心数，表示当前已分配给虚拟机和容器的CPU核心数量，用于资源使用率监控 | [90] |
| **CPUAllocable** | integer | CPU可分配核心数，表示当前可用于新建虚拟机的CPU核心数量，计算方式为总数减去已用数量 | [100] |
| **MemoryCap** | integer | 内存总容量，以MB为单位，表示该地域所有物理服务器的内存总容量，用于内存资源规划 | [564] |
| **MemoryUsed** | integer | 内存已使用量，以MB为单位，表示当前已分配给虚拟机和系统的内存总量，包含虚拟化开销 | [175] |
| **MemoryPhysicUsed** | integer | 物理内存已使用量，以MB为单位，表示物理服务器实际使用的内存量，不包含虚拟化预分配 | [196] |
| **MemoryReserved** | integer | 内存保留量，以MB为单位，表示为系统稳定性和突发需求预留的内存容量，不可分配给用户 | [12] |
| **MemoryAllocable** | integer | 内存可分配量，以MB为单位，表示当前可用于创建新虚拟机的内存容量，已扣除保留量 | [357] |
| **VGPUCount** | integer | 虚拟GPU总数量，表示通过GPU虚拟化技术创建的vGPU实例总数，用于GPU资源的细粒度分配 | [0] |
| **VGPUUsed** | integer | 虚拟GPU已使用数量，表示当前已分配给虚拟机的vGPU实例数量，用于GPU资源使用率监控 | [0] |
| **GPUCount** | integer | 物理GPU总数量，表示该地域所有物理服务器上安装的GPU卡总数，包含各种型号的GPU | [0] |
| **GPUAllocable** | integer | GPU可分配数量，表示当前可用于虚拟化或直通分配的GPU数量，用于GPU资源调度 | [0] |
| **GPUUsed** | integer | GPU已使用数量，表示当前已分配给虚拟机或容器的GPU数量，包含直通和虚拟化方式 | [0] |
| **StorageCount** | integer | 存储总容量，以GB为单位，表示该地域所有存储设备的总容量，包含SSD、HDD等各类存储 | [13933] |
| **StorageUsed** | integer | 存储已使用量，以GB为单位，表示当前已分配给虚拟机和系统的存储空间总量 | [6540] |
| **StoragePhysicUsed** | integer | 物理存储已使用量，以GB为单位，表示存储设备实际占用的物理空间，不包含预分配空间 | [719] |
| **CreateTime** | integer | 创建时间，Unix时间戳格式，表示该地域首次纳管到系统中的时间，用于审计和统计 | [1753789360] |
| **UpdateTime** | integer | 更新时间，Unix时间戳格式，表示该地域信息最后一次更新的时间，用于数据同步和缓存 | [1754546800] |
| **ExpireTime** | integer | 过期时间，Unix时间戳格式，表示该地域授权或许可证的过期时间，用于合规管理 | [0] |
| **State** | string | 地域状态，表示该地域当前的运行状态 | [authorized] |
| **Connectivity** | string | 连接性状态，表示该地域与中心管理系统的网络连接状态 | [] |

## 更新地域 - UpdateRegion

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，用于标识资源所属的地域，实现多地域环境下的资源隔离。 | [cn] | [DescribeRegion] | **Yes** |
| **Name** | string | 地域名称 | [北京] | [] | **Yes** |
| **City** | string | 城市，地域所在的城市信息 | [北京] | [] | **Yes** |
| **Address** | string | 地址，地域的地址信息 | [上海市杨浦区xxxx] | [] | No |
| **Remark** | string | 备注，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [This is region_remark] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 修改地域下资源名称和备注 - ModifyNameAndRemark

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，用于标识资源所属的地域，实现多地域环境下的资源隔离。 | [cn] | [DescribeRegion] | **Yes** |
| **ResourceID** | string | 资源ID，该地域下租户所拥有资源的资源ID | [disk-xwbm****s0hxy0] | [DescribeTenantResources] | **Yes** |
| **Name** | string | 资源名称，对应ResourceID的资源名称 | [disk_name] | [DescribeTenantResources] | **Yes** |
| **Remark** | string | 备注，对应ResourceID的资源备注，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [This is disk_remark] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|

