# PLATFORM_STORAGE_MANAGEMENT

## 查询平台通用存储状态 - DescribePlatformStorage

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Status** | string | 状态 | [] |
| **Total** | integer | 总量，单位Gi | [] |
| **Used** | integer | 已使用，单位Gi | [] |
| **DefaultSetType** | string | 默认存储集群 | [] |


## 创建平台通用存储云盘 - CreatePlatformStorageDisk

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **SetType** | string | 集群类型 | [] | [] | **Yes** |
| **DiskSpace** | integer | 磁盘大小 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **DiskID** | string | 磁盘ID | [] |


## 删除平台通用存储云盘 - DeletePlatformStorageDisk

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **DiskID** | string | 磁盘ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 绑定平台通用存储云盘 - AttachPlatformStorageDisk

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **DiskID** | string | 磁盘ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 扩容平台通用存储云盘 - ResizePlatformStorageDisk

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **DiskID** | string | 磁盘ID | [] | [] | **Yes** |
| **DiskSpace** | integer | 磁盘大小 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 查询平台通用存储云盘 - DescribePlatformStorageDisk

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **Offset** | integer | 分页偏移量 | [] | [] | No |
| **Limit** | integer | 分页大小 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer |  | [] |
| **Infos** | array<[*PlatformStorageDiskInfo*](#PlatformStorageDiskInfo)> | 磁盘信息 | [] |

### 数据模型

#### PlatformStorageDiskInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **DiskID** | string | 磁盘ID | [] |
| **DiskSpace** | integer | 磁盘大小 | [] |
| **DiskStatus** | string | 磁盘状态 | [] |
| **SetType** | string | 集群类型 | [] |
| **SetArch** | string | 集群架构 | [] |
