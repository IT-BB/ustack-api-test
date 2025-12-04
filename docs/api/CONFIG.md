# CONFIG

## 获取地域指定配置 - GetRegionConfig

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **ConfigKey** | string | 配置键 | [] | [] | **Yes** |
| **Region** | string | 地域 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Info** | object | 详情 | [] |

### 数据模型

#### RegionConfigBaseInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ConfigKey** | string | 配置键 | [] |
| **ConfigValue** | string | 配置值 | [] |

## 按照类型和地域获取全局配置 - ListGlobalConfigs

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **ConfigType** | string | 配置类型 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*CategoryInfo*](#CategoryInfo)> | 详情 | [] |
| **TotalCount** | integer | 总数 | [] |

### 数据模型

#### CategoryInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Category** | string | 类别 | [] |
| **CategoryName** | string | 类别名 | [] |
| **CategoryPosition** | integer | 类别排序 | [] |
| **ConfigInfos** | array<[*ConfigInfo*](#ConfigInfo)> | 详情 | [] |

#### ConfigInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ConfigKey** | string | 配置键 | [] |
| **ConfigValue** | string | 配置值 | [] |
| **Description** | string | 配置描述 | [] |
| **ValueType** | string | 配置值类型 | [] |
| **Unit** | string | 配置值单位 | [] |
| **ValueRange** | string | 配置值范围 | [] |
| **Category** | string | 类别 | [] |
| **CategoryName** | string | 类别名 | [] |
| **Editable** | string | 是否可编辑 | [] |
| **Hint** | string | 配置提示 | [] |
| **ConfigPosition** | integer | 配置排序 | [] |
| **ConfigType** | string | 配置类型 | [] |

## 按照类型和地域获取地域配置 - ListRegionConfigs

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ConfigType** | string | 配置类型 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*RegionConfigCategoryInfo*](#RegionConfigCategoryInfo)> | 详情 | [] |
| **TotalCount** | integer | 总数 | [] |

### 数据模型

#### RegionConfigCategoryInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Category** | string | 类别 | [] |
| **CategoryName** | string | 类别名 | [] |
| **CategoryPosition** | integer | 类别排序 | [] |
| **ConfigInfos** | array<[*RegionConfigInfo*](#RegionConfigInfo)> | 详情 | [] |

#### RegionConfigInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ConfigKey** | string | 配置键 | [] |
| **ConfigValue** | string | 配置值 | [] |
| **Description** | string | 配置描述 | [] |
| **ValueType** | string | 配置值类型 | [] |
| **Unit** | string | 配置值单位 | [] |
| **ValueRange** | string | 配置值范围 | [] |
| **Category** | string | 类别 | [] |
| **CategoryName** | string | 类别名 | [] |
| **Editable** | string | 是否可编辑 | [] |
| **Hint** | string | 配置提示 | [] |
| **ConfigPosition** | integer | 配置排序 | [] |
| **ConfigType** | string | 配置类型 | [] |

## 更新地域配置 - UpdateRegionConfig

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ConfigKey** | string | 配置键 | [] | [] | **Yes** |
| **ConfigValue** | string | 配置值 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 更新产品规格 - UpdateProductSpecification

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **SpecificationID** | string | 规格ID | [] | [] | **Yes** |
| **Value** | string | 规格值 | [] | [] | **Yes** |
| **Region** | string | 地域 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 更新配置 - UpdateConfig

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **ConfigKey** | string | 配置键 | [] | [] | **Yes** |
| **ConfigValue** | string | 配置值 | [] | [] | No |
| **FileBytes** | string | 图片Base64二进制内容 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取指定配置 - GetConfig

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **ConfigKey** | string | 配置键 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Info** | object | 详情 | [] |

### 数据模型

#### ConfigBaseInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ConfigKey** | string | 配置键 | [] |
| **ConfigValue** | string | 配置值 | [] |
| **FileBytes** | string | 文件Base64二进制内容 | [] |

## 获取sso配置信息 - GetSSOConfig

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Info** | object | 详情 | [] |

### 数据模型

#### SSOConfigInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **SystemActiveStatus** | string | 系统活跃状态 | [] |
| **Registerable** | string | 是否可注册 | [] |
| **SSORegisterPasswordLength** | integer | 登录密码最低长度 | [] |
| **SSORegisterPasswordComplexity** | string | 登录密码复杂度 | [] |
| **SSOPasswordInvalidChange** | bool | 登录密码非法是否要求修改 | [] |

## 查询配额资源用量列表 - DescribeQuotaUsage

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | No |
| **CompanyID** | integer | 租户ID | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*QuotaUsage*](#QuotaUsage)> | 详情 | [] |
| **TotalCount** | integer | 总数 | [] |

### 数据模型

#### QuotaUsage

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Region** | string | 地域 | [] |
| **Quota** | integer | 配额 | [] |
| **ProductType** | string | 产品类型 | [] |
| **ResourceType** | string | 资源类型 | [] |
| **FactorType** | string | 配额因子 | [] |
| **Usage** | integer | 已使用 | [] |
| **Unit** | string | 单位 | [] |
| **RegionAlias** | string | 地域别名 | [] |

## 查询集群资源用量列表 - DescribeSetAllocateUsage

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 不传就查所有Region，排行返回以Region为key。传了就查单个Region，排行以Region下的集群为单位。 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **VCPUUsageInfos** | array<[*VCPUUsageInfo*](#VCPUUsageInfo)> | vCPU使用情况 | [] |
| **MemoryUsageInfos** | array<[*MemoryUsageInfo*](#MemoryUsageInfo)> | 内存使用情况 | [] |
| **StorageUsageInfos** | array<[*StorageUsageInfo*](#StorageUsageInfo)> | 存储使用情况 | [] |

### 数据模型

#### MemoryUsageInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Region** | string | 地域 | [] |
| **SetID** | string | 集群ID | [] |
| **RegionAlias** | string | 地域别名 | [] |
| **MemoryCap** | integer | 内存总量 | [] |
| **MemoryUsed** | integer | 内存已使用 | [] |
| **MemoryAllocable** | integer | 内存可分配 | [] |
| **MemoryPhysicUsed** | integer | 物理内存已使用 | [] |
| **MemoryReserved** | integer | 内存预留量 | [] |
| **MemoryUsageRate** | double | 内存使用率 | [] |

#### StorageUsageInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Region** | string | 地域 | [] |
| **SetID** | string | 集群ID | [] |
| **RegionAlias** | string | 地域别名 | [] |
| **StorageCap** | integer | 存储总量 | [] |
| **StorageUsed** | integer | 存储已使用 | [] |
| **StoragePhysicUsed** | integer | 物理存储已使用 | [] |
| **StorageUsageRate** | double | 存储使用率 | [] |
| **StoragePhysicUsageRate** | double | 存储实际使用率 | [] |

#### VCPUUsageInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Region** | string | 地域 | [] |
| **SetID** | string | 集群ID | [] |
| **RegionAlias** | string | 地域别名 | [] |
| **CPUCount** | integer | CPU总量 | [] |
| **CPUPhysicCount** | integer | 物理CPU总量 | [] |
| **CPUUsed** | integer | CPU已使用 | [] |
| **CPUAllocable** | integer | CPU可分配 | [] |
| **CPUUsageRate** | double | CPU使用率 | [] |

## 设置资源配额 - SetAccountQuota

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **ResourceType** | string | 资源类型 | [] | [] | **Yes** |
| **FactorType** | string | 配额因子 | [] | [] | No |
| **Quota** | integer | 配额 | [] | [] | **Yes** |
| **Allocation** | string | 计算/存储集群分配的配额，json字符串 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 查询配额列表 - DescribeQuota

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | No |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **ProductTypes** | array<string> |  | [] | [] | No |
| **FactorTypes** | array<string> |  | [] | [] | No |
| **Limit** | integer | 分页大小 | [] | [] | No |
| **Offset** | integer | 分页偏移量 | [] | [] | No |
| **Keyword** | string | 搜索关键词· | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数 | [] |
| **Infos** | array<[*Quota*](#Quota)> | 详情 | [] |

### 数据模型

#### Quota

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] |
| **ResourceType** | string | 资源类型 | [] |
| **ProductType** | string | 产品类型 | [] |
| **FactorType** | string | 配额因子 | [] |
| **Region** | string | 地域 | [] |
| **RegionAlias** | string | 地域别名 | [] |
| **Quota** | integer | 配额 | [] |
| **Allocation** | string | 集群分配配额 | [] |
| **Status** | string | 状态 | [] |
| **Unit** | string | 单位 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |

## 验证邮箱服务器可用性 - VerifyEmailAvailability

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Receiver** | string | 收件人 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 查询产品规格 - DescribeProductSpecification

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **SpecificationIDs** | array<string> | 规格ID列表 | [] | [] | No |
| **ProductTypes** | array<string> | 产品类型 | [] | [] | No |
| **SetTypes** | array<string> | 集群类型 | [] | [] | No |
| **ResourceType** | string | 资源类型 | [] | [] | No |
| **SpecificationName** | string | 规格名称 | [] | [] | No |
| **Limit** | integer | 分页大小 | [] | [] | No |
| **Offset** | integer | 分页偏移量 | [] | [] | No |
| **Keyword** | string | 关键词 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*Specification*](#Specification)> | 详情 | [] |
| **TotalCount** | integer | 总数 | [] |

### 数据模型

#### Dimension

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **DimensionValue** | string | 维度值 | [] |
| **DimensionName** | string | 维度名称 | [] |

#### Specification

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Region** | string | 地域 | [] |
| **SetType** | string | 集群类型 | [] |
| **SpecificationID** | string | 规格ID | [] |
| **ProductType** | string | 产品类型 | [] |
| **ResourceType** | string | 资源类型 | [] |
| **SpecificationName** | string | 规格名称 | [] |
| **Value** | string | 规格值 | [] |
| **Default** | string | 默认值 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |
| **Dimension** | object | 维度 | [] |

## 获取规格/价格/配额分类筛选关键字 - GetFilterKeywords

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **QuotaInfo** | object | 配额 | [] |
| **SpecificationInfo** | object | 规格 | [] |
| **BillPriceGroupInfo** | object | 价格 | [] |

### 数据模型

#### BillPriceGroupInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ProductIDs** | array<string> |  | [] |
| **SetTypes** | array<string> |  | [] |

#### QuotaInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ProductTypes** | array<string> |  | [] |
| **FactorTypes** | array<string> |  | [] |

#### SpecificationInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ProductTypes** | array<string> |  | [] |
| **SetTypes** | array<string> |  | [] |

## 查询产品规格模板 - DescribeProductSpecificationTemplate

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **ProductType** | string | 产品类型 | [] | [] | No |
| **ResourceType** | string | 资源类型 | [] | [] | No |
| **SpecificationName** | string | 规格名称 | [] | [] | No |
| **Limit** | integer | 分页大小 | [] | [] | No |
| **Offset** | integer | 分页偏移量 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*SpecificationTemplate*](#SpecificationTemplate)> | 详情 | [] |
| **TotalCount** | integer | 总数 | [] |

### 数据模型

#### SpecificationTemplate

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TemplateID** | string | 模板ID | [] |
| **ProductType** | string | 产品类型 | [] |
| **ResourceType** | string | 资源类型 | [] |
| **SpecificationName** | string | 规格名称 | [] |
| **Description** | string | 规格描述 | [] |
| **ValueType** | string | 规格值类型 | [] |
| **ValueFormat** | string | 规格值格式 | [] |
| **AllowedValues** | string | 规格允许值 | [] |
| **SetClassification** | string | 集群分类 | [] |
| **Unit** | string | 单位 | [] |

## 创建产品规格 - CreateProductSpecification

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **SetType** | string | 集群类型 | [] | [] | **Yes** |
| **ResourceType** | string | 资源类型 | [] | [] | **Yes** |
| **SpecificationName** | string | 规格名称 | [] | [] | **Yes** |
| **Value** | string | 规格值 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **SpecificationID** | string | 规格ID | [] |


## 删除产品规格 - DeleteProductSpecification

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **SpecificationID** | string | 规格ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **SpecificationID** | string | 规格ID | [] |


## 获取资源信息 - DescribeResourceInfo

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **ResourceIDs** | array<string> |  | [] | [] | No |
| **Region** | string |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ResourceInfos** | array<[*ResourceInfos*](#ResourceInfos)> |  | [] |

### 数据模型

#### ResourceInfos

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Region** | string | 地域 | [] |
| **ResourceID** | string | 资源ID | [] |
| **Name** | string | 资源名称 | [] |
| **CompanyID** | string | 租户ID | [] |
| **ResourceType** | string | 资源类型 | [] |
| **CreateTime** | integer | 资源创建时间 | [] |
