# SECURITY_COMMON

## 查询安全组件规格信息 - GetSecurityProductSpec

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ResourceType** | string | 资源类型 | [] | [] | **Yes** |
| **VMType** | string | 计算集群类型 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*SecurityProductSpecInfos*](#SecurityProductSpecInfos)> |  | [] |

### 数据模型

#### ExtraAsset

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string |  | [] |
| **Value** | integer |  | [] |

#### SecurityProductSpecInfos

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string | 规格的key 用于API传值 | [] |
| **VersionName** | string | 规格 多语言下显示为各自翻译 | [] |
| **Rank** | integer | 规格等级 | [] |
| **CPU** | integer | cpu | [] |
| **Memory** | integer | 内存 | [] |
| **BootDiskSpace** | integer | 系统盘大小 | [] |
| **ExtraAssets** | array<[*ExtraAsset*](#ExtraAsset)> | 可管理资产数 | [] |

## 获取安全组件各规格价格 - GetSecurityProductSpecPrice

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **Quantity** | integer | 计费数量 | [] | [] | **Yes** |
| **VersionNameEN** | string | 规格英文 | [] | [] | **Yes** |
| **VMType** | string | 计算集群 | [] | [] | **Yes** |
| **DiskSetType** | string | 存储集群 | [] | [] | **Yes** |
| **ResourceType** | string | 资源类型 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Price** | double | 价格 | [] |


## 获取升级配置后的差价 - GetSecurityProductUpgradedPrice

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **ResourceID** | string | 安全组件ID | [] | [] | **Yes** |
| **VersionNameEN** | string | 规格英文 | [] | [] | **Yes** |
| **VMType** | string | 计算集群 | [] | [] | **Yes** |
| **DiskSetType** | string | 存储集群 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Price** | double | 价格 | [] |


## 获取安全组件单点登录URL - GetSecurityProductSSOURL

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **ResourceID** | string | 资源ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **URL** | string | 单点登录URL | [] |

