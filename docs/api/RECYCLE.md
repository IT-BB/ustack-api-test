# RECYCLE

## 获取回收站资源 - DescribeRecycledResource

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [DescribeUser] | No |
| **Region** | string | 所属的地域 | [develop] | [DescribeRegion] | **Yes** |
| **ResourceIDs** | array<string> | 资源IDs，该地域下租户所拥有资源的资源IDs，多个资源ID之间用逗号分隔 | [disk-xwbm****s0hxy0,vm-xwbm****s0hab0] | [DescribeTenantResources] | **Yes** |
| **ProjectIDs** | array<string> | 项目ID，用于标识资源所属的项目，多个项目ID之间用逗号分隔 | [project-51chy****squbp] | [ListProjects] | No |
| **Keyword** | string | 搜索关键词，支持租户名称、邮箱、手机号 | [ucloud] | [] | No |
| **Status** | string | 状态，无需填写 | [] | [] | No |
| **Offset** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 | [0] | [] | No |
| **Limit** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。 | [20] | [] | No |
| **ParentResourceID** | string | 父资源 ID，可根据外网 ID 获取处于回收站中的 EIP | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*RecycleResourceInfo*](#RecycleResourceInfo)> | 回收资源信息 | [] |
| **TotalCount** | integer | 总回收资源数量 | [] |

### 数据模型

#### RecycleResourceInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string | 资源名称，用于标识和显示资源的自定义名称 | [my-vm-001] |
| **Remark** | string | 资源备注信息，用于记录资源的详细描述或用途说明 | [测试环境虚拟机] |
| **Region** | string | 资源所在地域的标识符，表示资源的物理位置 | [pro2133] |
| **RegionAlias** | string | 地域的中文别名，便于用户理解的地域显示名称 | [测试地域] |
| **CompanyID** | integer | 租户ID，标识资源所属的租户账户 | [200000233] |
| **CompanyName** | string | 租户名称，资源所属租户的显示名称 | [租户名称] |
| **Email** | string | 租户邮箱地址，用于联系和通知的邮箱 | [example@ucloud.cn] |
| **ProjectID** | string | 项目ID，用于标识资源所属的项目，可以通过 ListProjects 获取 | [project-8mz5l****9k9lp] |
| **ProjectName** | string | 项目名称，资源所属项目的显示名称 | [默认项目] |
| **Reason** | string | 资源操作原因，记录资源状态变更或操作的原因说明 | [用户主动删除] |
| **CreateTime** | integer | 资源创建时间，Unix时间戳格式，表示资源首次创建的时间 | [1640995200] |
| **UpdateTime** | integer | 资源最后更新时间，Unix时间戳格式，记录资源最近一次修改的时间 | [1640995800] |
| **DeleteTime** | integer | 资源删除时间，Unix时间戳格式，表示资源被删除进入回收站的时间 | [1641000000] |
| **ResourceID** | string | 资源唯一标识符，全局唯一的资源ID | [uhost-abc123def456] |
| **ResourceType** | string | 资源类型，标识资源的种类，如VM、DISK、EIP等 | [VM] |
| **Description** | string | 资源描述信息，详细说明资源的配置和特性 | [2核4G内存的云主机] |
| **State** | string | 资源业务状态的字符串表示，用于前端显示，如Available、Deleting等 | [DELETED] |
| **Status** | integer | 资源状态码，数值表示的资源状态：0-初始化中，1-可用，3-失败，97-删除中，98-终止中，99-已删除，100-已终止 | [99] |
| **ExpireTime** | integer | 资源过期时间，Unix时间戳格式，表示资源的到期时间，0表示永不过期 | [1672531200] |
| **WillTerminateTime** | integer | 资源预计销毁时间，Unix时间戳格式，表示回收站中的资源将被自动销毁的时间 | [1641604800] |
| **IsAutoTerminated** | bool | 是否自动销毁标识，true表示资源会在到期后自动从回收站销毁，false表示需要手动操作 | [true] |
| **IsFailed** | bool | 是否失败状态标识，true表示资源处于失败状态，可能需要特殊处理 | [false] |

## 恢复资源 - RollbackResource

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [DescribeUser] | No |
| **Region** | string | 所属的地域 | [develop] | [DescribeRegion] | **Yes** |
| **ResourceID** | string | 资源ID，该地域下租户所拥有资源的资源ID | [disk-xwbm****s0hxy0] | [DescribeTenantResources] | **Yes** |
| **ProjectID** | string | 项目ID，用于标识资源所属的项目 | [project-8mz5l****9k9lp] | [ListProjects] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 销毁资源 - TerminateResource

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [DescribeUser] | No |
| **Region** | string | 所属的地域 | [develop] | [DescribeRegion] | **Yes** |
| **ResourceID** | string | 资源ID，该地域下租户所拥有资源的资源ID | [disk-xwbm****s0hxy0] | [DescribeTenantResources] | **Yes** |
| **ProjectID** | string | 项目ID，用于标识资源所属的项目 | [project-8mz5l****9k9lp] | [ListProjects] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|

