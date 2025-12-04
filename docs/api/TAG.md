# TAG

## 创建标签 - CreateTag

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [DescribeUser] | No |
| **Key** | string | Tag的键，用于标识标签的名称，必须唯一。最大长度127个字符 | [example_key] | [] | **Yes** |
| **Values** | array<string> | Tag的值，用于标识标签的具体取值，最大长度127个字符 | [example_value] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 删除标签 - DeleteTag

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [DescribeUser] | No |
| **Key** | string | Tag的键，用于标识标签的名称，必须唯一。最大长度127个字符 | [example_key] | [] | **Yes** |
| **Values** | array<string> | Tag的值，用于标识标签的具体取值，最大长度127个字符，为空则删除key | [example_value] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 查询标签 - DescribeTag

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [DescribeUser] | No |
| **Key** | string | Tag的键，用于标识标签的名称，必须唯一。最大长度127个字符 | [example_key] | [] | No |
| **Values** | array<string> | Tag的值，用于标识标签的具体取值，最大长度127个字符，为空则删除key | [example_value] | [] | No |
| **Keyword** | string | 搜索关键词 | [example_key] | [] | No |
| **Offset** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 | [0] | [] | No |
| **Limit** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。 | [20] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*TagInfo*](#TagInfo)> | 标签信息 | [] |
| **TotalCount** | integer | 总记录数 | [] |

### 数据模型

#### TagInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，标识标签所属的租户账户 | [200000233] |
| **CompanyName** | string | 租户名称，标签所属租户的显示名称 | [UCloud] |
| **CompanyEmail** | string | 租户邮箱地址，用于联系和通知的邮箱 | [example@ucloud.cn] |
| **Key** | string | 标签键名，用于标识标签的名称 | [example_key] |
| **Value** | string | 标签值，与键名配对使用的标签值 | [example_value] |
| **ResourceCount** | integer | 资源数量，使用该标签的资源总数 | [1] |

## 绑定标签 - BindTag

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 所属的地域，用于唯一标识地域 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [DescribeUser] | No |
| **KeyValuePairs** | array<string> | 键值对，租户下唯一，形如key(Base64):value(Base64)，其中key和value为Base64编码后的值 | [] | [] | **Yes** |
| **ResourceIDs** | array<string> | 待绑定资源ID | [disk-lckc3f****aof2] | [ListProductResources] | **Yes** |
| **ResourceType** | string | 资源类型 | [DISK] | [ListProductResources] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 标签解绑 - UnBindTag

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 所属的地域 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [DescribeUser] | No |
| **Keys** | array<string> | Tag的键，用于标识标签的名称，多个keys用逗号分隔 | [example_key1,example_key2] | [] | **Yes** |
| **ResourceIDs** | array<string> | 待解绑资源ID | [disk-lckc3f****aof2, disk-lfac3f****aof2] | [ListProductResources] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 查询标签资源 - DescribeTagResource

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [DescribeUser] | No |
| **Key** | string | Tag的键，用于标识标签的名称 | [example_key1] | [] | No |
| **Values** | array<string> | Tag的值，用于标识标签的具体取值，最大长度127个字符 | [example_value1, example_value2] | [] | **Yes** |
| **BindResource** | string | 绑定的资源ID | [disk-lckc3f****aof2] | [] | No |
| **BindResourceType** | string | 绑定的资源类型 | [DISK] | [] | No |
| **Offset** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 | [0] | [] | No |
| **Limit** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。 | [20] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*TagResourceInfo*](#TagResourceInfo)> | 标签资源信息 | [] |
| **TotalCount** | integer | 总记录数 | [] |

### 数据模型

#### Tag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string | 标签键名，用于标识标签的名称 | [example_key] |
| **Value** | string | 标签值，与键名配对使用的标签值 | [example_value] |

#### TagResourceInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Region** | string | 地域标识符，表示资源所在的地域 | [cn] |
| **ResourceID** | string | 资源唯一标识符，全局唯一的资源ID | [disk-lckc3fgga3aof2] |
| **ResourceName** | string | 资源名称，用户自定义的资源显示名称 | [example_disk] |
| **ResourceType** | string | 资源类型，可以通过ListProductResources获取，如VM、DISK、EIP等 | [DISK] |
| **Tags** | array<[*Tag*](#Tag)> | 标签列表，资源绑定的所有标签信息 | [] |
| **TagCount** | integer | 标签数量，该资源绑定的标签总数 | [1] |
| **RegionAlias** | string | 地域别名，便于用户理解的地域显示名称 | [cn] |

## 查询可绑定标签的资源 - DescribeBindableTagResource

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [DescribeUser] | No |
| **Region** | string | 所属的地域ID，用于表示资源所在的地域 | [develop] | [DescribeRegion] | **Yes** |
| **Key** | string | 键，用于标识标签的名称 | [example_key] | [] | **Yes** |
| **ResourceType** | string | 资源类型，如VM、DISK、EIP等 | [DISK] | [ListProductResources] | **Yes** |
| **Offset** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 | [0] | [] | No |
| **Limit** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。 | [20] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*BindableTagResourceInfo*](#BindableTagResourceInfo)> | 标签资源信息 | [] |
| **TotalCount** | integer | 总记录数 | [] |

### 数据模型

#### BindableTagResourceInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Region** | string | 地域标识符，表示资源所在的地域 | [cn] |
| **ResourceID** | string | 资源唯一标识符，全局唯一的资源ID | [disk-lckc3fgga3aof2] |
| **ResourceName** | string | 资源名称，用户自定义的资源显示名称 | [example_disk] |
| **ResourceType** | string | 资源类型，标识资源的种类，如VM、DISK、EIP等 | [DISK] |
| **CompanyID** | integer | 租户ID，标识资源所属的租户账户 | [200000236] |
| **CompanyName** | string | 租户名称，资源所属租户的显示名称 | [ucloud] |
| **CompanyEmail** | string | 租户邮箱地址，用于联系和通知的邮箱 | [ucloud@example.cn] |
| **TagValue** | string | 标签值，如果资源已经绑定过当前标签键才会展示该值 | [example_value] |
| **CreateTime** | integer | 资源创建时间，Unix时间戳格式，表示资源首次创建的时间 | [1640995200] |

## 设置资源最终绑定的所有标签 - SetResourceTags

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 所属的地域，用于唯一标识地域 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [DescribeUser] | No |
| **KeyValuePairs** | array<string> | 键值对，租户下唯一，形如key:value，其中key和value为Base64编码后的值 | [dGVzdA==:dGVzdC0x] | [] | No |
| **ResourceID** | string | 待绑定资源ID | [disk-lckc3f****aof2] | [ListProductResources] | **Yes** |
| **ResourceType** | string | 资源类型 | [DISK] | [ListProductResources] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|

