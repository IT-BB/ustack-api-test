# PROJECT

## 添加角色授权 - CreateMemberTag

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **MemberID** | integer | 账号ID，用于表示主（子）账号的唯一标识 | [200000234] | [CreateSubMember DescribeMember GetMemberInfo] | **Yes** |
| **RoleIDs** | array<string> | 角色ID，用于唯一标识角色，便于角色管理，多个角色ID用,分隔 | [role-0qov****ci8m42] | [CreateRole GetRole] | **Yes** |
| **ProjectIDs** | array<string> | 项目组ID列表，多个项目组用,分隔 | [project-51chy****squbp] | [GetProject ListProjects] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 删除角色授权 - DeleteMemberTag

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **MemberID** | integer | 账号ID，用于表示主（子）账号的唯一标识 | [200000234] | [CreateSubMember DescribeMember GetMemberInfo] | **Yes** |
| **RoleID** | string | 角色ID，用于唯一标识角色，便于角色管理 | [role-0qov****ci8m42] | [CreateRole GetRole] | **Yes** |
| **ProjectID** | string | 项目组ID，用于唯一标识项目组，便于项目组管理 | [project-51chy****squbp] | [GetProject ListProjects] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 查询角色授权列表 - ListMemberTags

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **MemberID** | integer | 账号ID，用于表示主（子）账号的唯一标识 | [200000234] | [CreateSubMember DescribeMember GetMemberInfo] | No |
| **RoleID** | string | 角色ID，用于唯一标识角色，便于角色管理 | [role-0qov****ci8m42] | [CreateRole GetRole] | No |
| **Offset** | integer | 分页偏移量，指定返回结果的起始位置。用于分页查询，默认为0。 | [0] | [] | No |
| **Limit** | integer | 分页大小，指定每页返回的记录数量。用于控制返回结果的数量，建议设置合理的值以提高查询效率。 | [25] | [] | No |
| **Keyword** | string | 关键词，用于搜索角色名称或备注 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*AuthorizationInfo*](#AuthorizationInfo)> | 角色授权信息详情 | [] |
| **TotalCount** | integer | 总数 | [] |

### 数据模型

#### AuthorizationInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **MemberID** | integer | 账号ID | [200000237] |
| **Email** | string | 账号对应的邮箱地址 | [example@ucloud.cn] |
| **RoleID** | string | 角色ID | [role-51chy****squbp] |
| **RoleName** | string | 角色名称 | [] |
| **ProjectID** | string | 项目组ID | [] |
| **ProjectName** | string | 项目组名称 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |
| **RoleType** | string | 角色类型，表示系统内置或自定义 | [System Custom] |

## 创建租户级自定义角色 - CreateRole

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Name** | string | 角色名称，长度为1-50个字符，名称只能包含中英文、数字、点（.），下划线（_）和中划线（-） | [new_role_name] | [] | **Yes** |
| **Remark** | string | 角色备注，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [new_role_remark] | [] | No |
| **Permission** | string | 角色权限，定义相关的权限信息，action-1表示操作有权限，允许该操作，action-0表示操作无权限，不填默认操作有权限。多个权限用,分隔 | [DeleteCompany-0,DescribeRecharge-0,GetWithdrawableAmount-0] | [ListProductPermissions] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **RoleID** | string | 角色ID，用于唯一标识角色，便于角色管理 | [role-0qov****ci8m42] |


## 删除租户级自定义角色 - DeleteRole

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **RoleID** | string | 角色ID，用于唯一标识角色，便于角色管理 | [role-0qov****ci8m42] | [CreateRole ListRoles GetRole] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 查询角色详情 - GetRole

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **RoleID** | string | 角色ID，用于唯一标识角色，便于角色管理 | [role-0qov****ci8m42] | [CreateRole ListRoles] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Info** | object | 角色信息详情 | [] |

### 数据模型

#### Permission

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string | 权限的名称 | [DescribeRecharge] |
| **Remark** | string | 权限的备注 | [获取充值信息] |
| **Permission** | integer | 权限的标识值，默认为0 | [0] |
| **IsPermission** | integer | 权限表示，表明是否有权限，0：没有权限，1：有权限 | [1] |
| **ActionType** | string | 操作类型，表示DESCRIBE, CREATE, UPDATE, DELETE, PROXY, NONE | [ACTION_TYPE_DESCRIBE ACTION_TYPE_NONE] |

#### ProductPermission

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ProductType** | string | 产品类型 | [BILL] |
| **Name** | string | 产品名称 | [BILL] |
| **Permissions** | array<[*Permission*](#Permission)> | 产品所有的权限列表 | [] |

#### RoleDetailInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **RoleID** | string | 角色ID | [role-0qov****ci8m42] |
| **CompanyID** | integer | 租户ID | [200000233] |
| **Name** | string | 角色名称 | [new_role_name] |
| **Remark** | string | 角色的备注 | [new_role_remark] |
| **Type** | string | 角色类型，取值范围：全部，系统内置（System），自定义（Custom）,Type为空则查询全部 | [System Custom] |
| **RolePermissions** | array<[*ProductPermission*](#ProductPermission)> | 角色的权限列表 | [] |
| **CreateTime** | integer | 角色信息的创建时间 | [1694502400] |
| **UpdateTime** | integer | 角色信息的更新时间 | [1694502400] |

## 查询租户级角色列表 - ListRoles

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **RoleIDs** | array<string> | 角色ID，用于唯一标识角色，便于角色管理，多个角色ID用,分隔 | [role-0qov****ci8m42] | [CreateRole GetRole] | No |
| **Type** | string | 角色类型，取值范围：全部，系统内置（System），自定义（Custom）,Type为空则查询全部 | [System Custom] | [] | No |
| **Offset** | integer | 分页偏移量，指定返回结果的起始位置。用于分页查询，默认为0。 | [0] | [] | No |
| **Limit** | integer | 分页大小，指定每页返回的记录数量。用于控制返回结果的数量，建议设置合理的值以提高查询效率。 | [25] | [] | No |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [CreateUser DescribeUser DescribeMember] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*RoleInfo*](#RoleInfo)> | 角色详情列表 | [] |
| **TotalCount** | integer | 角色详情列表的总数 | [] |

### 数据模型

#### RoleInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **RoleID** | string | 角色ID | [role-0qov****ci8m42] |
| **CompanyID** | integer | 租户ID | [200000233] |
| **Name** | string | 角色名称 | [new_role_name] |
| **Remark** | string | 角色备注，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [new_role_remark] |
| **Type** | string | 角色类型，取值范围：全部，系统内置（System），自定义（Custom）,Type为空则查询全部 | [System Custom] |
| **CreateTime** | integer | 角色信息的创建时间 | [1694502400] |
| **UpdateTime** | integer | 角色信息的更新时间 | [1694502400] |

## 修改角色权限 - UpdateRolePermission

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **RoleID** | string | 角色ID，用于唯一标识角色，便于角色管理，多个角色ID用,分隔 | [role-0qov****ci8m42] | [CreateRole GetRole] | **Yes** |
| **Permission** | string | 角色权限，定义相关的权限信息 | [DeleteCompany,DescribeRecharge,GetWithdrawableAmount] | [ListProductPermissions] | **Yes** |
| **IsPermission** | integer | 是否开启权限，0表示无权限，1表示有权限 | [0] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 重命名角色名称备注，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 - RenameRole

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **RoleID** | string | 角色ID，用于唯一标识角色，便于角色管理 | [role-0qov****ci8m42] | [CreateRole ListRoles GetRole] | **Yes** |
| **Name** | string | 角色名称，长度为1-50个字符，名称只能包含中英文、数字、点（.），下划线（_）和中划线（-） | [角色名称] | [] | **Yes** |
| **Remark** | string | 角色备注，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [new_role_remark] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取租户可用接口 - ListProductPermissions

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*ProductPermission*](#ProductPermission)> | 产品权限信息列表 | [] |
| **TotalCount** | integer | 产品权限信息列表的总数 | [] |

### 数据模型

#### Permission

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string | 权限的名称 | [DescribeRecharge] |
| **Remark** | string | 权限的备注 | [获取充值信息] |
| **Permission** | integer | 权限的标识值，默认为0 | [0] |
| **IsPermission** | integer | 权限表示，表明是否有权限，0：没有权限，1：有权限 | [1] |
| **ActionType** | string | 操作类型，表示DESCRIBE, CREATE, UPDATE, DELETE, PROXY, NONE | [ACTION_TYPE_DESCRIBE ACTION_TYPE_NONE] |

#### ProductPermission

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ProductType** | string | 产品类型 | [BILL] |
| **Name** | string | 产品名称 | [BILL] |
| **Permissions** | array<[*Permission*](#Permission)> | 产品所有的权限列表 | [] |

## 修改资源所在的项目 - MoveProjectResource

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于唯一标识租户，便于租户管理 | [200000233] | [CreateCompany GetCompany] | **Yes** |
| **ProjectID** | string | 项目组ID，用于唯一标识项目组，便于项目组管理 | [project-51chy****squbp] | [GetProject ListProjects] | **Yes** |
| **ResourceIDs** | array<string> | 资源ID列表，多个资源ID用,分隔 | [vm-m3wz2****biug8] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 创建项目 - CreateProject

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于唯一标识租户，便于租户管理 | [200000233] | [CreateCompany GetCompany] | **Yes** |
| **Name** | string | 项目名称，长度为1-50个字符，名称只能包含中英文、数字、点（.），下划线（_）和中划线（-） | [项目组名称] | [] | **Yes** |
| **Remark** | string | 备注，用于进行说明和注释，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [项目组的备注信息] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ProjectID** | string | 项目组ID，用于唯一标识项目组，便于项目组管理 | [] |


## 删除项目 - DeleteProject

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **ProjectID** | string | 项目组ID，用于唯一标识项目组，便于项目组管理 | [project-51chy****squbp] | [GetProject ListProjects] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取项目详情 - GetProject

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **ProjectID** | string | 项目组ID，用于唯一标识项目组，便于项目组管理 | [project-51chy****squbp] | [GetProject ListProjects] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Info** | object | 项目详情 | [] |

### 数据模型

#### ProjectInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ProjectID** | string | 项目组ID | [project-51chy****squbp] |
| **Name** | string | 名称 | [项目组名称] |
| **Remark** | string | 备注 | [项目组的备注信息] |
| **CreateTime** | integer | 创建时间 | [1694502400] |
| **UpdateTime** | integer | 更新时间 | [1694502400] |
| **CompanyID** | integer | 租户ID | [200000233] |
| **CompanyEmail** | string | 租户邮箱 | [example@ucloud.cn] |
| **CompanyName** | string | 租户名称 | [UCloud] |
| **IsDefault** | bool | 是否默认 | [true] |

## 查询项目列表 - ListProjects

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **ProjectIDs** | array<string> | 项目组ID列表，用于查看多个项目组的详情，为空则获取所有项目组 | [project-8mz5****a9k9ln] | [GetProject CreateProject] | No |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [CreateUser DescribeUser DescribeMember] | No |
| **Offset** | integer | 分页偏移量，指定返回结果的起始位置。用于分页查询，默认为0。 | [0] | [] | No |
| **Limit** | integer | 分页大小，指定每页返回的记录数量。用于控制返回结果的数量，建议设置合理的值以提高查询效率。 | [10] | [] | No |
| **Keyword** | string | 关键词，用于搜索项目组名称或备注 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*ProjectInfo*](#ProjectInfo)> | 项目组的相关详情 | [] |
| **TotalCount** | integer | 总数 | [] |

### 数据模型

#### ProjectInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ProjectID** | string | 项目组ID | [project-51chy****squbp] |
| **Name** | string | 名称 | [项目组名称] |
| **Remark** | string | 备注 | [项目组的备注信息] |
| **CreateTime** | integer | 创建时间 | [1694502400] |
| **UpdateTime** | integer | 更新时间 | [1694502400] |
| **CompanyID** | integer | 租户ID | [200000233] |
| **CompanyEmail** | string | 租户邮箱 | [example@ucloud.cn] |
| **CompanyName** | string | 租户名称 | [UCloud] |
| **IsDefault** | bool | 是否默认 | [true] |

## 重命名项目名称备注 - RenameProject

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **ProjectID** | string | 项目组ID，用于唯一标识项目组，便于项目组管理 | [project-51chy****squbp] | [GetProject CreateProject ListProjects] | **Yes** |
| **Name** | string | 项目组名称，长度为1-50个字符，名称只能包含中英文、数字、点（.），下划线（_）和中划线（-） | [项目组名称] | [] | **Yes** |
| **Remark** | string | 项目组备注，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [new_project_remark] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取某产品已授权的租户信息 - ListProductTypeCompanys

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **ProductType** | string | 产品类型 | [DISK] | [ListProductResources] | **Yes** |
| **Enabled** | string | 是否启用，如果不弃用填 'N'，其他值表示启用。 | [N Yes true] | [] | **Yes** |
| **Offset** | integer | 分页偏移量，指定返回结果的起始位置。用于分页查询，默认为0。 | [0] | [] | No |
| **Limit** | integer | 分页大小，指定每页返回的记录数量。用于控制返回结果的数量，建议设置合理的值以提高查询效率。 | [25] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*CompanyInfo*](#CompanyInfo)> | 获取到的租户详情 | [] |
| **TotalCount** | integer | 总数 | [] |

### 数据模型

#### CompanyInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [200000233] |
| **CompanyName** | string | 租户名称 | [UCloud] |
| **Email** | string | 邮箱 | [example@ucloud.cn] |

## 获取产品类型 - DescribeProduct

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | No |
| **ProductTypes** | array<string> | 产品类型 | [DISK] | [ListProductResources] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*ProductTypeInfo*](#ProductTypeInfo)> | 产品列表 | [] |
| **TotalCount** | integer | 总条数 | [] |

### 数据模型

#### ProductTypeInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Region** | string | 所属的地域，您可以调用 DescribeRegion 方法 查看最新的地域列表。 | [pro2133] |
| **Category** | string | 产品分类 | [Basic] |
| **CategoryName** | string | 产品分类名称 | [基础产服务] |
| **ProductType** | string | 产品类型 | [DISK] |
| **ProductKey** | string | 产品ID | [DISK] |
| **Name** | string | 产品名字 | [DISK] |
| **Remark** | string | 产品描述 | [云硬盘提供用于虚拟机的持久性数据块级存储服务] |
| **Authorized** | string | 是否授权 | [Authorized] |
| **Enable** | string | 是否启用 | [Enable] |
| **EnableCompanyCount** | integer | 已启用租户数量 | [6] |
| **CompanyCount** | integer | 租户总数 | [6] |
| **Services** | array<[*Service*](#Service)> | 子服务列表 | [] |
| **Invisible** | bool | 是否不可见 | [false] |

#### Service

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ServiceProductType** | string | 子服务产品类型 | [] |
| **ServiceProductKey** | string | 子服务产品ID | [] |
| **Name** | string | 子服务产品名字 | [] |
| **Authorized** | string | 是否授权 | [] |

## 租户关闭服务 - DisableCompanyProductType

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyIDs** | array<int32> | 租户ID列表，用于标识资源所属的租户，实现多租户环境下的资源隔离，多个租户用,分隔。 | [200000233, 200000234] | [CreateUser DescribeUser DescribeMember] | **Yes** |
| **ProductType** | string | 产品类型 | [DISK] | [ListProductResources] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 租户启用服务 - EnableCompanyProductType

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyIDs** | array<int32> | 租户ID列表，用于标识资源所属的租户，实现多租户环境下的资源隔离，多个租户用,分隔。 | [200000233, 200000234] | [CreateUser DescribeUser DescribeMember] | **Yes** |
| **ProductType** | string | 产品类型 | [DISK] | [ListProductResources] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取产品资源关系 - ListProductResources

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*ProductResource*](#ProductResource)> | 产品详情 | [] |
| **TotalCount** | integer | 总数 | [] |

### 数据模型

#### ProductResource

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ProductType** | string | 产品类型 | [] |
| **ResourceType** | string | 资源类型 | [] |
