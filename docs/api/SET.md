# SET

## 设置计算集群别名 - AliasSet

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **SetType** | string | 计算集群类型 | [] | [] | **Yes** |
| **Alias** | string | 集群类型别名 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 设置存储集群别名 - AliasStorageSet

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **StorageSetType** | string | 存储集群类型 | [] | [] | **Yes** |
| **Alias** | string | 存储集群类型别名 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 查询存储集群信息 - DescribeStorageSet

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **SetArch** | string | 集群架构 | [] | [] | No |
| **SetType** | string | 集群类型 | [] | [] | No |
| **SetIDs** | array<string> | 集群ID列表 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer |  | [] |
| **Infos** | array<[*StorageSetInfo*](#StorageSetInfo)> | 存储集群信息 | [] |

### 数据模型

#### StorageSetBlockBackupStatus

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Health** | string | 健康 | [] |
| **BackupVolumeTotalCount** | integer | 备份卷总数 | [] |
| **AbnormalVolumeIDs** | array<string> | 异常备份卷ID列表 | [] |

#### StorageSetInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Region** | string | 地域 | [] |
| **SetAlias** | string | 集群别名 | [] |
| **CreateTime** | integer |  | [] |
| **UpdateTime** | integer |  | [] |
| **SetID** | string | 集群ID | [] |
| **SetArch** | string | 集群架构 | [] |
| **SetType** | string | 集群类型 | [] |
| **SetProvider** | string | 集群制备器 | [] |
| **StorageCap** | integer | 存储集群容量 | [] |
| **StorageUsed** | integer | 存储集群使用量 | [] |
| **StoragePhysicUsed** | integer | 物理使用量 | [] |
| **Redundancy** | string | 存储集群冗余策略 | [] |
| **SetBlockBackupRole** | string |  | [] |
| **BlockBackupPeerSetID** | string | 卷备份对端集群 | [] |
| **BlockBackupStatus** | object | 集群卷备份状态 | [] |
| **DOSMultisiteRole** | string | dos 多站点角色，值为：'master'、'slave'、'' | [] |
| **Authorized** | string |  | [] |
| **VmSetList** | array<[*VmSetItem*](#VmSetItem)> | 可以使用该存储集群的计算集群列表，物理绑定，硬限制 | [] |
| **BoundVMSetList** | array<[*VmSetItem*](#VmSetItem)> | 当前存储集群绑定的计算集群列表，逻辑绑定，软限制 | [] |
| **Cache** | string | 是否开启了缓存 | [] |

#### VmSetItem

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **VmSetID** | string |  | [] |
| **VmSetName** | string |  | [] |
| **VmSetType** | string |  | [] |

## 查询存储类型 - DescribeStorageType

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer |  | [] |
| **Infos** | array<[*StorageSetTypeInfo*](#StorageSetTypeInfo)> | 存储集群信息 | [] |

### 数据模型

#### StorageSetTypeInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Region** | string | 地域 | [] |
| **StorageTypeAlias** | string | 存储集群别名 | [] |
| **Permission** | string | 权限 | [] |
| **StorageType** | string | 存储集群类型 | [] |
| **SetArch** | string | 存储集群架构 | [] |

## 获取虚拟机Set信息 - DescribeVMSet

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **SetIDs** | array<string> | 集群ID列表 | [] | [] | No |
| **FilterRunnableSetsByVMID** | string |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer |  | [] |
| **Infos** | array<[*SetInfo*](#SetInfo)> | 集群信息 | [] |

### 数据模型

#### CPUModelInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **CPUModel** | string | CPUModel | [] |
| **SupportedHosts** | array<string> | 支持此CPUModel的物理机 | [] |
| **Disabled** | bool | 是否被禁用 | [] |

#### PFNodeInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **NodeID** | string |  | [] |
| **PFSummaries** | array<[*PFSummary*](#PFSummary)> |  | [] |

#### PFSummary

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Vendor** | string | 厂商，例如：Mellanox Technologies | [] |
| **Product** | string | 型号，例如：MT27800 Family [ConnectX-5] | [] |
| **VFCount** | integer | vf总量 | [] |
| **VFUsed** | integer | vf已使用数量 | [] |
| **PFCode** | string | 物理网卡型号的标准编号，例如：15b3-1018 | [] |

#### SetGPUUsedInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **HDName** | string | HD名称 | [] |
| **MdevName** | string | Mdev名称 | [] |
| **Count** | integer | 数量 | [] |
| **Used** | integer | 使用量 | [] |
| **Allocable** | integer | 可分配量 | [] |

#### SetInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Region** | string | 地域 | [] |
| **SetAlias** | string | 集群别名 | [] |
| **CreateTime** | integer |  | [] |
| **UpdateTime** | integer |  | [] |
| **SetID** | string | 集群ID | [] |
| **SetArch** | string | 集群架构 | [] |
| **SetType** | string | 集群类型 | [] |
| **CPUCount** | integer | CPU数量 | [] |
| **CPUPhysicCount** | integer | 物理CPU数量 | [] |
| **CPUUsed** | integer | CPU使用量 | [] |
| **CPUAllocable** | integer | CPU可分配量 | [] |
| **CPUAllocationRatio** | float | CPU分配比例 | [] |
| **MemoryCap** | integer | 内存容量 | [] |
| **MemoryUsed** | integer | 内存使用量 | [] |
| **MemoryAllocable** | integer | 内存可分配量 | [] |
| **MemoryPhysicUsed** | integer | 内存物理使用量 | [] |
| **MemoryReserved** | integer | 内存预留量 | [] |
| **GPUCount** | integer | GPU数量 | [] |
| **GPUUsed** | integer | GPU使用量 | [] |
| **GPUAllocable** | integer | GPU可分配量 | [] |
| **VGPUInfos** | array<[*SetVGPUUsedInfo*](#SetVGPUUsedInfo)> | 集群VGPU使用信息 | [] |
| **GPUInfos** | array<[*SetGPUUsedInfo*](#SetGPUUsedInfo)> | 集群GPU使用信息 | [] |
| **CPUManufacturer** | array<string> | cpu型号 | [] |
| **StorageClassList** | array<[*StorageClassItem*](#StorageClassItem)> | 当前计算集群可用的存储集群列表，物理绑定，硬限制 | [] |
| **CPUModelIntersection** | array<[*CPUModelInfo*](#CPUModelInfo)> | 集群cpu model交集 | [] |
| **CPUModelDiffsection** | array<[*CPUModelInfo*](#CPUModelInfo)> | 集群cpu model差集 | [] |
| **BoundStorageClassList** | array<[*StorageClassItem*](#StorageClassItem)> | 当前计算集群绑定的存储集群列表，逻辑绑定，软限制 | [] |
| **BoundImageIDs** | string | 当前计算集群绑定的镜像ID列表 | [] |
| **PFNodeInfos** | array<[*PFNodeInfo*](#PFNodeInfo)> | 物理网卡信息 | [] |

#### SetVGPUUsedInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **MdevName** | string | Mdev名称 | [] |
| **Count** | integer | 数量 | [] |
| **Used** | integer | 使用量 | [] |
| **Allocable** | integer | 可分配量 | [] |

#### StorageClassItem

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **StorageSetID** | string |  | [] |
| **StorageSetName** | string |  | [] |
| **StorageSetType** | string |  | [] |

## 查询主机机型 - DescribeVMType

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer |  | [] |
| **Infos** | array<[*SetTypeInfo*](#SetTypeInfo)> | 集群信息 | [] |

### 数据模型

#### SetTypeInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Region** | string | 地域 | [] |
| **VMType** | string | 集群名称 | [] |
| **VMTypeAlias** | string | 集群别名 | [] |
| **Permission** | string | 权限 | [] |
| **SetArch** | string | 集群架构 | [] |

## 设置计算集群的超分比例 - UpdateComputeSetCPUAllocationRatio

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **SetID** | string | 集群ID | [] | [] | **Yes** |
| **CPUAllocationRatio** | float | CPU分配比例 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 设置计算CPU模型 - UpdateComputeSetCPUModels

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **SetID** | string | 集群ID | [] | [] | **Yes** |
| **DisabledCPUModels** | array<string> | 禁用的CPU模型 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 更新计算集群绑定的存储集群 - UpdateVMSetBoundStorageSet

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **SetID** | string | 计算集群ID | [] | [] | **Yes** |
| **StorageSetIDs** | array<string> | 存储集群ID列表 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 更新计算集群绑定的镜像 - UpdateVMSetBoundImage

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **SetID** | string | 计算集群ID | [] | [] | **Yes** |
| **ImageIDs** | string | 镜像ID列表 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 修改资源权限 - UpdateResourcePermission

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ResourceID** | string | 资源ID | [] | [] | **Yes** |
| **ResourceType** | string | 资源类型 | [] | [] | **Yes** |
| **Permission** | string | 权限 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 查询正在使用资源的用户 - DescribeResourceUsers

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ResourceID** | string | 资源 ID | [] | [] | **Yes** |
| **ResourceType** | string | 资源 类型 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer |  | [] |
| **Infos** | array<int32> | 租户ID | [] |

