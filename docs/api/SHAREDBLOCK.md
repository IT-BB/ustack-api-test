# SHAREDBLOCK

## 分配外置存储集群硬盘 - AllocateExternalStorageSetDisk

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **TargetCompanyID** | integer | 接受分配的租户ID | [] | [] | No |
| **DiskID** | string | 硬盘ID | [] | [] | **Yes** |
| **ProjectID** | string | 项目组ID | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 绑定外置存储 - AttachExternalDisk

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户 | [] | [] | No |
| **DiskID** | string | 磁盘ID | [] | [] | **Yes** |
| **ResourceID** | string | 资源ID | [] | [] | **Yes** |
| **ResourceType** | string | 资源类型 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 创建外置存储集群 - CreateExternalStorageSet

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Name** | string | 名称 | [] | [] | **Yes** |
| **CSIPluginAddress** | string |  | [] | [] | **Yes** |
| **UserName** | string | 用户名 | [] | [] | No |
| **Password** | string | 密码 | [] | [] | No |
| **EnableTargetIQNs** | array<string> | IQN启用列表 | [] | [] | No |
| **DisableTargetIQNs** | array<string> | IQN禁用列表 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **SetID** | string |  | [] |


## 更新外置存储集群 - UpdateExternalStorageSet

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **SetID** | string | 存储集群ID | [] | [] | **Yes** |
| **EnableTargetIQNs** | array<string> | IQN启用列表 | [] | [] | No |
| **DisableTargetIQNs** | array<string> | IQN禁用列表 | [] | [] | No |
| **AddTargetIQNs** | array<string> | IQN新增列表 | [] | [] | No |
| **RemoveTargetIQNs** | array<string> | IQN移除列表 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 删除外置存储集群 - DeleteExternalStorageSet

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **SetID** | string | 存储集群ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 查询外置存储集群硬盘 - DescribeExternalDisk

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户 | [] | [] | No |
| **Keyword** | string | 搜索关键词 | [] | [] | No |
| **ProjectIDs** | array<string> | 项目组IDs | [] | [] | No |
| **CompanyOnly** | integer | 过滤未分配给租户的记录 | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **SetID** | string | 存储集群ID | [] | [] | No |
| **SetType** | string | 存储集群 | [] | [] | No |
| **AttachResourceID** | string | 挂载的资源ID | [] | [] | No |
| **ComputeSetID** | string | 计算集群ID | [] | [] | No |
| **ShareAbleFilter** | string | 筛选共享盘，true 仅返回共享盘，false 仅返回普通盘，空 都返回 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer |  | [] |
| **Infos** | array<[*ExternalDiskInfo*](#ExternalDiskInfo)> |  | [] |

### 数据模型

#### AttachExternalShareBlock

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **AttachResourceID** | string |  | [] |
| **AttachResourceName** | string |  | [] |
| **AttachResourceType** | string |  | [] |
| **AttachStatus** | string |  | [] |

#### ExternalDiskInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string |  | [] |
| **Remark** | string |  | [] |
| **Region** | string |  | [] |
| **RegionAlias** | string |  | [] |
| **CompanyID** | integer |  | [] |
| **CompnayName** | string |  | [] |
| **Email** | string |  | [] |
| **ProjectID** | string |  | [] |
| **ProjectName** | string |  | [] |
| **Reason** | string |  | [] |
| **CreateTime** | integer |  | [] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> |  | [] |
| **SetAlias** | string |  | [] |
| **AttachResourceName** | string |  | [] |
| **AttachResourceType** | string |  | [] |
| **DiskID** | string |  | [] |
| **DiskType** | string |  | [] |
| **Size** | integer |  | [] |
| **DiskStatus** | string |  | [] |
| **LUNID** | string |  | [] |
| **SetID** | string |  | [] |
| **SetProvider** | string |  | [] |
| **SetType** | string |  | [] |
| **SetArch** | string |  | [] |
| **AttachResourceID** | string |  | [] |
| **ComputeSetIDs** | string |  | [] |
| **ShareAble** | bool |  | [] |
| **AttachShareBlockInfos** | array<[*AttachExternalShareBlock*](#AttachExternalShareBlock)> |  | [] |
| **Percent** | float | 迁移进度 | [] |
| **UsedForVirtSC** | bool |  | [] |
| **IQN** | string |  | [] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

## 查询外置存储集群 - DescribeExternalStorageSet

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Offset** | integer | 分页偏移量 | [] | [] | No |
| **Limit** | integer | 分页大小 | [] | [] | No |
| **Provisioner** | string | 存储集群提供者 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer |  | [] |
| **Infos** | array<[*ExternalStorageSetInfo*](#ExternalStorageSetInfo)> |  | [] |

### 数据模型

#### ExternalStorageSetInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Region** | string |  | [] |
| **SetAlias** | string |  | [] |
| **Status** | string |  | [] |
| **Reason** | string |  | [] |
| **SetID** | string |  | [] |
| **SetType** | string |  | [] |
| **SetArch** | string |  | [] |
| **CSIPluginAddress** | string |  | [] |
| **StorageCap** | integer |  | [] |
| **UserName** | string |  | [] |
| **LUNCount** | integer |  | [] |
| **CreateTime** | integer |  | [] |
| **EnableTargetIQNs** | array<string> |  | [] |
| **DisableTargetIQNs** | array<string> |  | [] |

## 查询外置存储集群类型 - DescribeExternalStorageType

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer |  | [] |
| **Infos** | array<[*ExternalStorageTypeInfo*](#ExternalStorageTypeInfo)> |  | [] |

### 数据模型

#### ExternalStorageTypeInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Region** | string |  | [] |
| **SetAlias** | string |  | [] |
| **SetID** | string |  | [] |
| **SetType** | string |  | [] |
| **SetArch** | string |  | [] |
| **SetUsage** | string |  | [] |
| **StorageCap** | integer |  | [] |
| **StorageUsed** | integer |  | [] |
| **StoragePhysicUsed** | integer |  | [] |

## 解绑外置存储 - DetachExternalDisk

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **DiskID** | string | 磁盘ID | [] | [] | **Yes** |
| **ResourceID** | string | 资源ID | [] | [] | **Yes** |
| **ResourceType** | string | 资源类型 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 扫描外置存储集群硬盘 - ScanISCSIDisk

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **SetID** | string | 集群ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer |  | [] |
| **Infos** | array<[*ExternalDiskInfo*](#ExternalDiskInfo)> |  | [] |

### 数据模型

#### AttachExternalShareBlock

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **AttachResourceID** | string |  | [] |
| **AttachResourceName** | string |  | [] |
| **AttachResourceType** | string |  | [] |
| **AttachStatus** | string |  | [] |

#### ExternalDiskInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string |  | [] |
| **Remark** | string |  | [] |
| **Region** | string |  | [] |
| **RegionAlias** | string |  | [] |
| **CompanyID** | integer |  | [] |
| **CompnayName** | string |  | [] |
| **Email** | string |  | [] |
| **ProjectID** | string |  | [] |
| **ProjectName** | string |  | [] |
| **Reason** | string |  | [] |
| **CreateTime** | integer |  | [] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> |  | [] |
| **SetAlias** | string |  | [] |
| **AttachResourceName** | string |  | [] |
| **AttachResourceType** | string |  | [] |
| **DiskID** | string |  | [] |
| **DiskType** | string |  | [] |
| **Size** | integer |  | [] |
| **DiskStatus** | string |  | [] |
| **LUNID** | string |  | [] |
| **SetID** | string |  | [] |
| **SetProvider** | string |  | [] |
| **SetType** | string |  | [] |
| **SetArch** | string |  | [] |
| **AttachResourceID** | string |  | [] |
| **ComputeSetIDs** | string |  | [] |
| **ShareAble** | bool |  | [] |
| **AttachShareBlockInfos** | array<[*AttachExternalShareBlock*](#AttachExternalShareBlock)> |  | [] |
| **Percent** | float | 迁移进度 | [] |
| **UsedForVirtSC** | bool |  | [] |
| **IQN** | string |  | [] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

## 扫描 FCSAN 外置存储集群硬盘 - ScanFCSAN

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*FCSANSetInfo*](#FCSANSetInfo)> |  | [] |

### 数据模型

#### AttachExternalShareBlock

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **AttachResourceID** | string |  | [] |
| **AttachResourceName** | string |  | [] |
| **AttachResourceType** | string |  | [] |
| **AttachStatus** | string |  | [] |

#### ExternalDiskInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string |  | [] |
| **Remark** | string |  | [] |
| **Region** | string |  | [] |
| **RegionAlias** | string |  | [] |
| **CompanyID** | integer |  | [] |
| **CompnayName** | string |  | [] |
| **Email** | string |  | [] |
| **ProjectID** | string |  | [] |
| **ProjectName** | string |  | [] |
| **Reason** | string |  | [] |
| **CreateTime** | integer |  | [] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> |  | [] |
| **SetAlias** | string |  | [] |
| **AttachResourceName** | string |  | [] |
| **AttachResourceType** | string |  | [] |
| **DiskID** | string |  | [] |
| **DiskType** | string |  | [] |
| **Size** | integer |  | [] |
| **DiskStatus** | string |  | [] |
| **LUNID** | string |  | [] |
| **SetID** | string |  | [] |
| **SetProvider** | string |  | [] |
| **SetType** | string |  | [] |
| **SetArch** | string |  | [] |
| **AttachResourceID** | string |  | [] |
| **ComputeSetIDs** | string |  | [] |
| **ShareAble** | bool |  | [] |
| **AttachShareBlockInfos** | array<[*AttachExternalShareBlock*](#AttachExternalShareBlock)> |  | [] |
| **Percent** | float | 迁移进度 | [] |
| **UsedForVirtSC** | bool |  | [] |
| **IQN** | string |  | [] |

#### FCSANSetInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **SetID** | string |  | [] |
| **SetArch** | string |  | [] |
| **SetType** | string |  | [] |
| **SetUsage** | string |  | [] |
| **Disks** | array<[*ExternalDiskInfo*](#ExternalDiskInfo)> |  | [] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

## 外置存储盘设置为可共享的磁盘 - SetShareAbleExternalStorage

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **DiskID** | string | 磁盘ID | [] | [] | **Yes** |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|

