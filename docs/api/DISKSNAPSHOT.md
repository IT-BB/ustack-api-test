# DISKSNAPSHOT

## 创建快照 - CreateSnapshot

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **ProjectID** | string | 项目组ID | [] | [] | No |
| **TagKeyValuePairs** | array<string> | 标签键值对 | [] | [] | **Yes** |
| **Name** | string | 名称 | [] | [] | **Yes** |
| **Remark** | string | 备注 | [] | [] | No |
| **Type** | string | 快照类型 | [] | [] | No |
| **DiskID** | string | 磁盘ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **SnapshotID** | string | 快照ID | [] |


## 删除快照 - DeleteSnapshot

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **SnapshotID** | string | 快照ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 查询快照 - DescribeSnapshot

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Offset** | integer | 分页偏移量 | [] | [] | No |
| **Limit** | integer | 分页大小 | [] | [] | No |
| **Keyword** | string | 关键词 | [] | [] | No |
| **ProjectIDs** | array<string> | 项目组ID | [] | [] | No |
| **DiskID** | string | 磁盘ID | [] | [] | No |
| **SnapshotIDs** | array<string> | 快照ID | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数 | [] |
| **Infos** | array<[*SnapshotInfo*](#SnapshotInfo)> | 详情 | [] |

### 数据模型

#### SnapshotInfo

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
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> | 标签 | [] |
| **DiskAttachResourceType** | string | 磁盘关联资源类型 | [] |
| **Type** | string | 快照类型 | [] |
| **SnapshotID** | string | 快照ID | [] |
| **SnapshotStatus** | string | 快照状态 | [] |
| **SnapshotSize** | integer | 快照大小 | [] |
| **DiskID** | string | 磁盘ID | [] |
| **DiskType** | string | 磁盘类型 | [] |
| **DiskAttachResourceID** | string | 磁盘关联资源ID | [] |
| **Encrypted** | bool | 是否加密 | [] |
| **DiskName** | string | 磁盘名称 | [] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

## 快照回滚 - RollbackSnapshot

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **DiskID** | string | 磁盘ID | [] | [] | **Yes** |
| **SnapshotID** | string | 快照ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|

