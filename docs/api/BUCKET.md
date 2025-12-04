# BUCKET

## 创建桶 - CreateBucket

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Name** | string | 桶名称，创建后不可修改 | [] | [] | **Yes** |
| **AccessType** | string | 访问类型 | [] | [] | **Yes** |
| **ObjectLockEnabled** | bool | 对象锁定开关 | [] | [] | No |
| **ObjectLockDays** | integer | 对象锁定时间，开启时必填 | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **OSSID** | string | 对象存储ID | [] | [] | No |
| **Versioning** | bool | 多版本开关 | [] | [] | No |
| **BucketCryptEnabled** | bool | 桶加密开关 | [] | [] | No |
| **Quota** | integer | 桶配额，单位GB | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 桶列表 - DescribeBuckets

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Offset** | integer |  | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **OSSID** | string | 对象存储ID | [] | [] | No |
| **Keyword** | string | 搜索关键词 | [] | [] | No |
| **BucketNames** | array<string> | 桶名称 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer |  | [] |
| **Infos** | array<[*BucketInfo*](#BucketInfo)> |  | [] |

### 数据模型

#### BucketInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **UserID** | string | 用户ID | [] |
| **Name** | string | 桶名称 | [] |
| **AccessType** | string | 访问类型 | [] |
| **UpdateTime** | integer | 更新时间 | [] |
| **Region** | string | 地域 | [] |
| **RegionAlias** | string | 地域别名 | [] |
| **Versioning** | bool | 多版本状态，废弃，应该使用 VersionStatus | [] |
| **ObjectLockEnabled** | bool | 对象锁定开关 | [] |
| **ObjectLockDays** | integer | 对象锁定天数 | [] |
| **Endpoints** | array<string> | 访问地址 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **VersionStatus** | string | 多版本状态，状态改为：关闭(Closed)、开启(Enabled)、暂停(Suspended) | [] |
| **Quota** | integer | 桶配额，单位GB | [] |

## 删除桶 - DeleteBucket

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Name** | string | 桶名称 | [] | [] | **Yes** |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **OSSID** | string | 对象存储ID | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 更新桶访问类型 - UpdateBucketAccessType

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Name** | string | 桶名称 | [] | [] | **Yes** |
| **AccessType** | string | 访问类型public,publicead,public-read-write | [] | [] | **Yes** |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **OSSID** | string | 对象存储ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 更新桶的对象锁定开关 - UpdateBucketObjectLock

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Name** | string | 桶名称 | [] | [] | **Yes** |
| **Days** | integer | 锁定天数 | [] | [] | **Yes** |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **OSSID** | string | 对象存储ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 更新桶的多版本开关 - UpdateBucketVersioning

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Name** | string | 桶名称 | [] | [] | **Yes** |
| **Enabled** | bool | 是否启用 | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **OSSID** | string | 对象存储ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取S3登录信息 - DOSLogin

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **Bucket** | string | 桶名称 | [] | [] | **Yes** |
| **OSSID** | string | 对象存储ID | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **AccessKey** | string | 密钥ID | [] |
| **SecretKey** | string | 密钥 | [] |
| **Endpoint** | string | 访问地址 | [] |
| **DOSAuth** | string | S3会话认证信息 | [] |


## 获取令牌列表 - DescribeDOSToken

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | string | 租户ID | [] | [] | No |
| **OSSID** | string | OSSID | [] | [] | **Yes** |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*DOSToken*](#DOSToken)> | 令牌列表 | [] |
| **TotalCount** | integer | 总数 | [] |

### 数据模型

#### DOSToken

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string | 令牌名称 | [] |
| **AccessKey** | string | 公钥 | [] |
| **SecretKey** | string | 私钥 | [] |
| **ExpireTime** | integer | 过期时间 | [] |
| **AllowOps** | array<string> | 允许的操作 | [] |
| **AllowBuckets** | array<string> | 授权的桶 | [] |
| **AllowObjectPrefix** | string | 授权的对象前缀, * 代表所有对象 | [] |
| **WhiteIPs** | array<string> | IP白名单 | [] |
| **BlackIPs** | array<string> | IP黑名单 | [] |

## 创建令牌 - CreateDOSToken

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | string | 租户ID | [] | [] | No |
| **OSSID** | string | OSSID | [] | [] | **Yes** |
| **Name** | string | 令牌名称 | [] | [] | **Yes** |
| **ExpireTime** | integer | 过期时间 | [] | [] | **Yes** |
| **AllowOps** | array<string> | 允许的操作 | [] | [] | **Yes** |
| **AllowBuckets** | array<string> | 授权的桶 | [] | [] | **Yes** |
| **AllowObjectPrefix** | string | 授权的对象前缀, * 代表所有对象 | [] | [] | **Yes** |
| **WhiteIPs** | array<string> | IP白名单 | [] | [] | No |
| **BlackIPs** | array<string> | IP黑名单 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 更新令牌 - UpdateDOSToken

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | string | 租户ID | [] | [] | No |
| **OSSID** | string | OSSID | [] | [] | **Yes** |
| **Name** | string | 令牌名称 | [] | [] | **Yes** |
| **ExpireTime** | integer | 过期时间 | [] | [] | **Yes** |
| **AllowOps** | array<string> | 允许的操作 | [] | [] | **Yes** |
| **AllowBuckets** | array<string> | 授权的桶 | [] | [] | **Yes** |
| **AllowObjectPrefix** | string | 授权的对象前缀, * 代表所有对象 | [] | [] | No |
| **WhiteIPs** | array<string> | IP白名单 | [] | [] | No |
| **BlackIPs** | array<string> | IP黑名单 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 删除令牌 - DeleteDOSToken

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | string | 租户ID | [] | [] | No |
| **OSSID** | string | OSSID | [] | [] | **Yes** |
| **Name** | string | 令牌名称 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 清空桶数据 - FlushBucket

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **OSSID** | string | 对象存储ID | [] | [] | **Yes** |
| **BucketName** | string | 桶名称 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Count** | integer | 被清理的对象数量 | [] |


## 更新存储桶配额 - UpdateBucketQuota

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **BucketName** | string | 桶名称 | [] | [] | **Yes** |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **OSSID** | string | 对象存储ID | [] | [] | **Yes** |
| **Quota** | integer | 桶配额，单位GB, 0 表示不限制，移除配额 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 创建桶的生命周期 - CreateBucketLifecycleRule

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **BucketName** | string | 桶名称 | [] | [] | **Yes** |
| **ObjectKeyPrefix** | string | 对象名前缀 | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ExpirationType** | string | 过期类型（当前版本、非当前版本、碎片） | [] | [] | **Yes** |
| **ExpirationDays** | integer | 过期天数 | [] | [] | **Yes** |
| **OSSID** | string | 对象存储ID | [] | [] | **Yes** |
| **TransitionStorageClass** | string | 归档存储类 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **LifeCycleID** | string | 生命周期ID | [] |


## 桶的生命周期列表 - DescribeBucketLifecycleRules

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Offset** | integer |  | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **BucketName** | string | 桶名称 | [] | [] | **Yes** |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **OSSID** | string | 对象存储ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer |  | [] |
| **BucketName** | string | 桶名称 | [] |
| **Infos** | array<[*BucketLifecycleRule*](#BucketLifecycleRule)> |  | [] |

### 数据模型

#### BucketLifecycleRule

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **LifeCycleID** | string | 生命周期ID | [] |
| **ObjectKeyPrefix** | string | 对象名前缀 | [] |
| **ExpirationType** | string | 过期类型（当前版本、非当前版本、碎片） | [] |
| **ExpirationDays** | integer | 过期天数 | [] |
| **TransitionStorageClass** | string | 归档存储类 | [] |

## 更新桶的生命周期 - UpdateBucketLifecycleRule

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **BucketName** | string | 桶名称 | [] | [] | **Yes** |
| **LifeCycleID** | string | 生命周期规则ID | [] | [] | **Yes** |
| **ObjectKeyPrefix** | string | 对象名前缀 | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ExpirationType** | string | 过期类型（当前版本、非当前版本、碎片） | [] | [] | **Yes** |
| **ExpirationDays** | integer | 过期天数 | [] | [] | **Yes** |
| **OSSID** | string | 对象存储ID | [] | [] | **Yes** |
| **TransitionStorageClass** | string | 归档存储类 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 删除桶的生命周期 - DeleteBucketLifecycleRule

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **BucketName** | string | 桶名称 | [] | [] | **Yes** |
| **LifeCycleIDs** | array<string> | 生命周期规则ID | [] | [] | **Yes** |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **OSSID** | string | 对象存储ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|

