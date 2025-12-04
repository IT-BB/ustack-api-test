# RESOURCE_TEMPLATE

## 创建资源模板 - CreateResourceTemplate

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **TemplateType** | string | 模版类型 | [] | [] | **Yes** |
| **Content** | string | 模板内容, json字符串, 等同于相应创建接口的入参 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TemplateID** | string | 模板ID | [] |


## 更新资源模板 - UpdateResourceTemplate

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **TemplateID** | string | 模板ID | [] | [] | **Yes** |
| **Content** | string | 模板内容, json字符串, 等同于相应创建接口的入参 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 查询资源模板 - DescribeResourceTemplate

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **TemplateIDs** | array<string> | 模板ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | No |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **ProjectIDs** | array<string> | 项目组ID | [] | [] | No |
| **Limit** | integer | 分页参数 | [] | [] | No |
| **Offset** | integer | 分页参数 | [] | [] | No |
| **Keyword** | string | 查询关键字 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数 | [] |
| **Infos** | array<[*ResourceTemplateInfo*](#ResourceTemplateInfo)> | 模板详情 | [] |

### 数据模型

#### ResourceTemplateInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string | 名称 | [] |
| **Remark** | string | 备注 | [] |
| **Region** | string | 地域 | [] |
| **RegionAlias** | string | 地域别名 | [] |
| **CompanyID** | integer | 租户ID | [] |
| **CompanyName** | string | 租户名称 | [] |
| **Email** | string | 邮箱 | [] |
| **ProjectID** | string | 项目组ID | [] |
| **ProjectName** | string | 项目组名称 | [] |
| **Reason** | string | 原因 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **DeleteTime** | integer | 删除时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> | 标签 | [] |
| **Status** | string | 状态 | [] |
| **TemplateID** | string | 模板ID | [] |
| **TemplateType** | string | 模板类型 | [] |
| **Content** | string | 模板内容 | [] |
| **BindAsGroupIDs** | array<string> | 资源模板绑定了弹性伸缩的ID | [] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

## 删除资源模版 - DeleteResourceTemplate

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **TemplateID** | string | 模版ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 通过模板创建资源 - CreateResourceFromTemplate

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **TemplateID** | string | 模版ID | [] | [] | **Yes** |
| **Name** | string | 资源名称 | [] | [] | No |
| **Password** | string | 资源密码 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ResourceID** | string | 资源ID | [] |

