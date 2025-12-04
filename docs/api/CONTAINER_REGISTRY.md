# CONTAINER_REGISTRY

## 创建镜像仓库 - CreateContainerImageRepository

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Name** | string | 名称 | [] | [] | **Yes** |
| **Remark** | string | 备注 | [] | [] | No |
| **Region** | string | 地域 | [] | [] | **Yes** |
| **ProjectID** | string | 项目组ID | [] | [] | No |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **TagKeyValuePairs** | array<string> | 标签键值对 | [] | [] | **Yes** |
| **Public** | bool | 是否是公有仓库 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 更新镜像仓库 - UpdateContainerImageRepository

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **ContainerImageRepositoryID** | string | 名称 | [] | [] | **Yes** |
| **Public** | bool | 是否是公有仓库 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 删除镜像仓库 - DeleteContainerImageRepository

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **ContainerImageRepositoryID** | string | 名称 | [] | [] | **Yes** |
| **DeleteAllImages** | bool | 是否删除所有镜像 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 查询镜像仓库 - DescribeContainerImageRepository

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **ProjectIDs** | array<string> | 项目组ID | [] | [] | No |
| **Public** | bool | 是否是公有仓库 | [] | [] | No |
| **Keyword** | string | 关键词 | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数 | [] |
| **Infos** | array<[*ContainerImageRepositoryInfo*](#ContainerImageRepositoryInfo)> | 详情 | [] |

### 数据模型

#### ContainerImageRepositoryInfo

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
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> | 标签 | [] |
| **ContainerImageRepositoryID** | string | 镜像仓库Id | [] |
| **Public** | bool | 是否为公有仓库 | [] |
| **ContainerImageNumbers** | integer | 镜像数量 | [] |
| **Status** | string | 镜像仓库状态 | [] |
| **ImagePullNumbers** | integer | 镜像拉取次数 | [] |
| **RegistryAddress** | string | 镜像仓库地址 | [] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

## 查询容器镜像 - DescribeContainerImage

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Keyword** | string | 关键词 | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **ContainerImageRepositoryID** | string | 镜像仓库ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数 | [] |
| **Infos** | array<[*ContainerImageInfo*](#ContainerImageInfo)> | 详情 | [] |
| **ContainerImageRepositoryAddress** | string | 仓库地址 | [] |

### 数据模型

#### ContainerImageInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string | 名称 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **ContainerImagePullNumbers** | integer | 镜像拉取次数 | [] |
| **LatestTag** | string | 最新tag名称 | [] |
| **ContainerImageRepositoryID** | string | 镜像仓库ID | [] |
| **CompanyID** | integer | 租户ID | [] |
| **FullName** | string | 镜像完整名字 | [] |

## 删除容器镜像 - DeleteContainerImage

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **ContainerImageRepositoryID** | string | 镜像仓库ID | [] | [] | **Yes** |
| **ContainerImageName** | string | 镜像名 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 查询容器镜像tags - DescribeContainerImageTag

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **ContainerImageRepositoryID** | string | 镜像仓库ID | [] | [] | **Yes** |
| **ContainerImageName** | string | 镜像名 | [] | [] | **Yes** |
| **Keyword** | string |  | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数 | [] |
| **Infos** | array<[*ContainerImageTagInfo*](#ContainerImageTagInfo)> | 详情 | [] |

### 数据模型

#### ContainerImageTagInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TagName** | string | Tag名称 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **ContainerImageTagDigest** | string | TagDigest | [] |
| **ContainerImageTagPullNumbers** | integer | 镜像Tag拉取次数 | [] |
| **Size** | integer | 镜像Size | [] |
| **ContainerImageRepositoryID** | string | 镜像仓库ID | [] |
| **ContainerImage** | string | 镜像名 | [] |
| **CompanyID** | integer | 租户ID | [] |
| **FullName** | string | 镜像完整名字 | [] |

## 删除容器镜像tag - DeleteContainerImageTag

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **ContainerImageRepositoryID** | string | 镜像仓库ID | [] | [] | **Yes** |
| **ContainerImageName** | string | 镜像名 | [] | [] | **Yes** |
| **Tag** | string | Tag号 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|

