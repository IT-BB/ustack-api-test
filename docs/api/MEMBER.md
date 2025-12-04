# MEMBER

## 密码登录 - LoginByPassword

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Email** | string | 租户邮箱，租户注册时提供的用于接收通知和联系的邮箱 | [example@ucloud.cn] | [DescribeUser] | **Yes** |
| **UserEmail** | string | 电子邮箱，用作登录账号，和Email一致 | [example@ucloud.cn] | [DescribeUser] | **Yes** |
| **Password** | string | 用户登录的密码，采用Base64编码方式传入，同时需要在请求头上添加一个自定义请求头来指定认证方法为 plain，例如X-Auth-Method: plain | [mnSslwW6bkAaldaPUROcuA==] | [] | **Yes** |
| **ClientID** | string | OAuth2 ClientID，OAuth2客户端标识符，用于标识第三方应用程序的唯一ID，需要预先在系统中注册，配合ClientSecret用于客户端身份验证，遵循OAuth2.0授权码模式流程 | [] | [] | No |
| **ResponseType** | string | OAuth2响应类型，ClientID存在时必填，指定授权服务器返回的授权凭证类型， | [] | [] | No |
| **RedirectURI** | string | OAuth2重定向URI，ClientID存在时必填，授权完成后用户代理重定向的回调地址，必须与客户端注册时配置的重定向URI完全匹配或为其子路径，用于接收授权码 | [] | [] | No |
| **State** | string | OAuth2状态参数，ClientID存在时必填，客户端生成的不透明随机字符串，用于防止跨站请求伪造(CSRF)攻击，授权服务器会原样返回此参数供客户端验证请求的完整性 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **CompanyID** | integer | 租户唯一标识ID，标识用户所属的租户组织，用于实现多租户环境下的资源隔离和权限控制，系统根据此ID确定用户的资源访问范围 | [200000233] |
| **MemberID** | integer | 用户账号唯一标识ID，标识登录用户的唯一身份标识符，用于后续API调用时的身份验证和权限判断 | [200000234] |
| **Email** | string | 用户邮箱地址，用户注册时使用的邮箱账号，作为用户身份的重要标识，用于密码重置、通知发送等功能 | [ucloud@example.com] |
| **SSOToken** | string | 单点登录令牌，用于在多个系统间实现免密登录的安全令牌，包含用户身份信息和权限范围，具有时效性 | [eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...] |
| **Code** | string | OAuth2授权码，用于OAuth2授权流程中的临时授权凭证，客户端使用此码换取访问令牌，具有短期有效性 | [abc123def456] |
| **State** | string | OAuth2状态参数，用于防止CSRF攻击的安全参数，客户端在授权请求中传入，授权服务器原样返回用于验证 | [random_state_string_xyz789] |


## 创建用户 - CreateSubMember

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [CreateUser DescribeUser] | **Yes** |
| **MemberEmail** | string | 邮箱地址，作为子账号的唯一标识和登录凭证。必须是有效的邮箱格式，且在当前租户下唯一。 | [subucloud@example.com] | [] | **Yes** |
| **Password** | string | 密码，密码长度限6-30个字符,密码须包含有小写字母、特殊字符(除空格)。 | [new_password_12345] | [] | **Yes** |
| **MemberName** | string | 子账号名称，长度为1-50个字符，名称只能包含中英文、数字、点（.），下划线（_）和中划线（-） | [张三 John Doe] | [] | **Yes** |
| **Phone** | string | 手机号码，用于接收短信验证码、安全通知等。格式需符合国际手机号标准。 | [138****8000] | [] | No |
| **ResetPassword** | string | 是否首次登录修改密码，是否要求子账号首次登录时强制修改密码。默认为No，取值：1, Y, Yes, YES, True, true 表示是；其他值表示否。用于增强账号安全性。 | [1 Y No] | [] | No |
| **OAuth2UniqueID** | string | OAuth2唯一标识ID，用于关联外部身份提供商（如企业SSO、LDAP等）的用户身份。可选字段，仅在启用第三方认证时使用。 | [oauth2_unique_id_12345 ldap_user_dn] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **MemberID** | integer | 新建的子账号ID，用于唯一标识子账号 | [200000234] |


## 冻结子账号 - FreezeSubMember

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **MemberID** | integer | 子账号ID，用于表示子账号的唯一标识 | [200000234] | [CreateSubMember DescribeMember GetMemberInfo] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 解冻子账号 - UnFreezeSubMember

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **MemberID** | integer | 子账号ID，用于表示子账号的唯一标识 | [200000234] | [CreateSubMember DescribeMember GetMemberInfo] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 删除用户 - DeleteMember

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **MemberID** | integer | 账号ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


##  - LogoutToken

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 由管理员为子账号更改密码 - ChangeMemberPassword

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **SpecMemberID** | integer | 账号ID，用于表示子账号的唯一标识 | [200000234] | [CreateSubMember DescribeMember GetMemberInfo] | **Yes** |
| **Password** | string | 密码，密码长度限6-30个字符,密码须包含有小写字母、特殊字符(除空格)。 | [new_password_12345] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 创建管理员 - CreateAdmin

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **MemberEmail** | string | 邮箱地址，作为子账号的唯一标识和登录凭证。必须是有效的邮箱格式，且在当前租户下唯一。 | [subucloud@example.com] | [] | **Yes** |
| **Password** | string | 密码，密码长度限6-30个字符,密码须包含有小写字母、特殊字符(除空格)。 | [new_password_12345] | [] | **Yes** |
| **MemberName** | string | 账号名称，长度为1-50个字符，名称只能包含中英文、数字、点（.），下划线（_）和中划线（-） | [张三 John Doe] | [] | **Yes** |
| **Phone** | string | 账号绑定的手机号码，用于接收短信验证码、安全通知等。格式需符合国际手机号标准。 | [138****8000] | [] | No |
| **ResetPassword** | string | 登录修改密码，是否要求账号首次登录时强制修改密码。默认为No，取值：1, Y, Yes, YES, True, true 表示是；其他值表示否。用于增强账号安全性。 | [1 Y No] | [] | No |
| **ReadOnly** | string | 是否为只读权限管理员，当RoleID为空时必填。取值：'1', 'Y', 'Yes', 'YES', 'True', 'true' 表示只读；其他值表示可读写。 | [1 false] | [] | No |
| **GrantLevel** | string | 管理员授权等级，决定管理员的管理范围。取值：'System'（系统级管理员，可管理所有租户），'Region'（区域级管理员，可管理指定区域资源）。 | [System] | [] | **Yes** |
| **RoleIDs** | array<string> | 角色ID列表，分配给管理员的角色，定义具体的权限范围。当ReadOnly为空时必填。可指定多个角色ID以组合权限。只能指定单一权限范围，不能即指定系统权限，又指定地域权限 | [system_read region_admin] | [ListRoles] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **MemberID** | integer | 管理员账号ID，用于唯一标识账号。 | [] |


## 删除管理员 - DeleteAdmin

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **MemberID** | integer | 管理员账号ID，用于表示账号的唯一标识，可以通过ListAdmin获取。 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 列出管理员 - ListAdmin

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Keyword** | string | 关键词，用于搜索账号名称或邮箱，为空则查询全部 | [admin] | [] | No |
| **Limit** | integer | 分页大小，指定返回结果的数量。用于分页查询。 | [20] | [] | No |
| **Offset** | integer | 分页偏移量，指定返回结果的起始位置。用于分页查询，默认为0。 | [0] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*AdminInfo*](#AdminInfo)> | 管理员信息列表，包含符合查询条件的所有管理员账号的详细信息。 | [] |
| **TotalCount** | integer | 管理员总数量，表示符合查询条件的管理员账号总数，用于分页计算。 | [10] |

### 数据模型

#### AdminInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **MemberID** | integer | 管理员账号唯一标识ID，系统自动生成的数字ID，用于内部标识和关联管理员数据。 | [200000234] |
| **MemberName** | string | 管理员显示名称，用于在界面中展示的管理员名称。支持中英文、数字和常见符号，长度1-50个字符。 | [系统管理员 System Admin] |
| **MemberEmail** | string | 管理员邮箱地址，作为管理员账号的唯一标识和登录凭证，同时用于接收系统通知邮件。 | [admin@example.com] |
| **Phone** | string | 管理员绑定的手机号码，用于接收短信验证码、安全通知等。支持国际手机号格式，可能出于隐私保护进行脱敏显示。 | [138****8000] |
| **Status** | string | 管理员账号状态，表示账号当前的可用性。取值：'Available'（可用）、'Freeze'（冻结）、'Locked'（锁定）、'Deleted'（已删除）。 | [Available Freeze] |
| **ReadOnly** | bool | 是否为只读权限管理员，决定管理员是否只能查看而不能修改系统配置。取值：true（只读权限）、false（读写权限）。 | [true false] |
| **Privilege** | string | 管理员权限级别，表示管理员的管理权限范围。取值：'Admin'（超级管理员）、'RegionAdmin'（区域管理员）、'SystemAdmin'（系统管理员）。 | [Admin RegionAdmin] |
| **GrantLevel** | string | 管理员授权等级，决定管理员的管理范围。取值：'System'（系统级管理员，可管理所有租户）、'Region'（区域级管理员，可管理指定区域资源）。 | [System Region] |
| **CreateTime** | integer | 管理员账号创建时间，Unix时间戳格式（秒级），记录管理员账号首次创建的时间。 | [1640995200] |
| **UpdateTime** | integer | 管理员账号最后更新时间，Unix时间戳格式（秒级），记录管理员账号信息最后一次修改的时间。 | [1640995200] |

## 获取用户访问控制信息 - GetMemberInfo

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Info** | object | 当前用户的所有信息 | [] |

### 数据模型

#### UserAccessInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **MemberID** | integer | 账号唯一标识ID，系统自动生成的数字ID，用于内部标识和关联用户数据。 | [200000234] |
| **CompanyID** | integer | 租户唯一标识ID，标识账号所属的租户组织，实现多租户环境下的资源隔离 | [200000233] |
| **Email** | string | 账号邮箱地址，作为账号的唯一标识和登录凭证，同时用于接收系统通知邮件。 | [ucloud@example.com] |
| **Phone** | string | 账号绑定的手机号码，用于接收短信验证码、安全通知等。支持国际手机号格式。 | [138****8000] |
| **UserEmail** | string | 用户邮箱地址（兼容字段），与Email字段功能相同。 | [ucloud@example.com] |
| **MemberName** | string | 账号显示名称，用于在界面中展示的用户名称。支持中英文、数字和常见符号，长度1-50个字符。 | [张三 John Doe] |
| **UserName** | string | 用户名称（兼容字段），与MemberName字段功能相同，保持向后兼容性。 | [张三 John Doe] |
| **Status** | string | 账号状态，表示账号当前的可用性。取值：'Available'（可用）、'Freeze'（冻结）、'Locked'（锁定）、'Deleted'（已删除）。 | [Available Freeze] |
| **CompanyStatus** | string | 租户状态，表示账号所属租户的当前状态。取值：'Available'（可用）、'Freeze'（冻结）、'Locked'（锁定）。 | [Available Freeze] |
| **CreateTime** | integer | 账号创建时间，Unix时间戳格式（秒级），记录账号首次创建的时间。 | [1640995200] |
| **UpdateTime** | integer | 账号最后更新时间，Unix时间戳格式（秒级），记录账号信息最后一次修改的时间。 | [1640995200] |
| **PublicKey** | string | 账号API公钥，用于API调用时的身份验证和签名验证，采用RSA算法生成。 | [MIIBIjANBgkqhkiG9w0BAQEFAAOCA******BCgKCAQEA] |
| **PrivateKey** | string | 账号API私钥，用于API调用时的请求签名，采用RSA算法生成。请妥善保管，避免泄露。 | [MIIEvQIBADANBgkqhkiG9w0BAQEF******cwggSjAgEAAoIBAQC] |
| **Privileges** | string | 账号权限级别，表示账号的管理权限范围。取值：'Admin'（管理员）、'Member'（普通成员）。 | [Admin Member] |
| **Audit** | integer | 租户审批开关状态，表示是否启用资源审批流程。取值：1（启用审批）、0（关闭审批）。 | [1 0] |
| **Amount** | double | 账号总余额，包含内部账号余额和外部账号余额的总和，单位为元，保留2位小数。 | [1000.50] |
| **AmountFree** | double | 内部账号余额，用于内部资源消费的余额，单位为元，保留2位小数。 | [500.25] |
| **AmountCredit** | double | 外部账号余额，用于外部服务消费的余额，单位为元，保留2位小数。 | [500.25] |
| **GrantLevels** | array<string> | 账号授权等级列表，表示账号拥有的管理权限范围。可包含多个等级：'System'（系统级）、'Region'（区域级）、'Project'（项目级）。 | [System Region] |
| **OAuth2UniqueID** | string | OAuth2第三方认证的唯一标识ID，用于关联外部身份提供商（如企业SSO、LDAP等）的用户身份。仅在启用第三方认证时有值。 | [oauth2_unique_id_12345 ldap_user_dn] |

## 获取用户访问控制接口信息 - DescribePermission

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Info** | array<[*PermissionInfo*](#PermissionInfo)> | 详情 | [] |

### 数据模型

#### PermissionInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Action** | string | API的action | [] |
| **Permission** | bool | 是否有该action权限 | [] |

## 修改账号邮箱 - UpdateMemberEmail

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [CreateUser DescribeUser DescribeMember GetMemberInfo] | No |
| **MemberID** | integer | 子账号ID，用于表示子账号的唯一标识 | [200000234] | [CreateSubMember DescribeMember GetMemberInfo] | **Yes** |
| **UserEmail** | string | 账号邮箱，修改后的子账号邮箱 | [subucloud@example.com] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 修改账号OAuth2唯一标识ID - UpdateMemberOAuth2UniqueID

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **MemberID** | integer | 子账号ID，用于表示子账号的唯一标识 | [200000234] | [CreateSubMember DescribeMember GetMemberInfo] | **Yes** |
| **OAuth2UniqueID** | string | OAuth2标识ID，修改后的子账号OAuth2唯一标识ID | [new_oauth2_unique_id_12345] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 修改账号安全手机 - UpdateMemberPhone

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **MemberID** | integer | 子账号ID，用于表示子账号的唯一标识 | [200000234] | [CreateSubMember DescribeMember GetMemberInfo] | **Yes** |
| **PhoneNum** | string | 新手机号码 | [138****8000] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 修改账号名称 - UpdateMemberName

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **MemberID** | integer | 账号ID，用于表示子账号的唯一标识 | [200000234] | [CreateSubMember DescribeMember GetMemberInfo] | **Yes** |
| **Name** | string | 账号名称，长度为1-50个字符，名称只能包含中英文、数字、点（.），下划线（_）和中划线（-） | [new_sub_member_name] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取账号列表 - DescribeMember

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **MemberIDs** | array<int32> | 用户ID列表，用于指定查询的用户。查询多个用户用,分隔 | [200000246, 200000247] | [DescribeUser] | No |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [CreateUser DescribeUser] | No |
| **Email** | string | 账号邮箱地址，用于按邮箱筛选账号，支持精确匹配和模糊匹配。 | [ucloud@example.com] | [DescribeUser] | No |
| **MemberPubKey** | string | 账号API公钥，用于按公钥筛选账号，支持精确匹配。采用RSA算法生成的Base64编码字符串。 | [MIIBIjANBgkqhkiG9w0BAQEFAAOCA******BCgKCAQEA] | [] | No |
| **Limit** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。 | [20] | [] | No |
| **Offset** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 | [0] | [] | No |
| **Keyword** | string | 关键词，用于模糊搜索账号名、邮箱等字段，支持中英文搜索。长度限制：1-50字符。 | [admin 测试账号] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*MemberInfo*](#MemberInfo)> | 账号详情列表，包含查询到的所有账号信息。数组长度受Limit参数限制。 | [] |
| **TotalCount** | integer | 总记录数，表示满足查询条件的账号总数（不受分页限制），用于前端计算总页数。 | [156] |

### 数据模型

#### MemberInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **MemberID** | integer | 账号唯一标识ID，系统自动生成的数字ID，用于内部标识和关联用户数据。 | [200000234] |
| **CompanyID** | integer | 租户唯一标识ID，标识账号所属的租户组织，实现多租户环境下的资源隔离 | [200000233] |
| **Email** | string | 账号邮箱地址，作为账号的唯一标识和登录凭证，同时用于接收系统通知邮件。 | [admin@company.com] |
| **Password** | string | 账号登录密码（已加密存储），出于安全考虑返回时通常为空或掩码显示。 | [Mypw@1234] |
| **Phone** | string | 账号绑定的手机号码，用于接收短信验证码、安全通知等。支持国际手机号格式。 | [138****8000] |
| **MemberName** | string | 账号显示名称，用于在界面中展示的用户名称。支持中英文、数字和常见符号，长度1-50个字符。 | [张三 John Doe test_user] |
| **PublicKey** | string | 账号API公钥，用于API调用时的身份验证和签名验证，采用RSA算法生成。 | [MIIBIjANBgkqhkiG9w0BAQEFAAOCA******BCgKCAQEA] |
| **PrivateKey** | string | 账号API私钥，用于API调用时的请求签名，采用RSA算法生成。请妥善保管，避免泄露。 | [MIIEvQIBADANBgkqhkiG9w0BAQEF******cwggSjAgEAAoIBAQC] |
| **Privileges** | string | 账号权限级别，表示账号的管理权限范围。取值：'Admin'（管理员）、'Member'（普通成员）。 | [Admin Member] |
| **Admin** | integer | 管理员标识，表示账号的管理员级别。取值：0（普通用户）、1（管理员）、2（超级管理员）。 | [1 0] |
| **Permission** | array<int32> | 权限列表，具体的权限ID数组，每个ID对应系统中的特定权限项。 | [[1, 3, 5, 7]] |
| **CreateTime** | integer | 账号创建时间，Unix时间戳格式（秒级），记录账号首次创建的时间。 | [1640995200] |
| **UpdateTime** | integer | 账号最后更新时间，Unix时间戳格式（秒级），记录账号信息最后一次修改的时间。 | [1640995200] |
| **PwdUpdateTime** | integer | 密码最后更新时间，Unix时间戳格式（秒级），记录密码最后一次修改的时间，用于密码过期策略。 | [1640995200] |
| **Status** | string | 账号状态，表示账号当前的可用性。取值：'active'（正常）、'frozen'（冻结）、'disabled'（禁用）、'pending'（待审核）。 | [active frozen] |
| **Audit** | integer | 审批状态，表示账号的审批情况。取值：0（待审批）、1（已通过）、2（已拒绝）、3（审批中）。 | [1 0] |
| **Amount** | double | 账号总余额，包含内部账号余额和外部账号余额的总和，单位为元，保留2位小数。 | [1000.50] |
| **AmountFree** | double | 内部账号余额，用于内部资源消费的余额，单位为元，保留2位小数。 | [500.25] |
| **AmountCredit** | double | 外部账号余额，用于外部服务消费的余额，单位为元，保留2位小数。 | [500.25] |
| **OAuth2UniqueID** | string | OAuth2第三方认证的唯一标识ID，用于关联外部身份提供商（如企业SSO、LDAP等）的用户身份。仅在启用第三方认证时有值。 | [oauth2_unique_id_12345 ldap_user_dn] |

## 更新数字证书 - UpdateDigitalCert

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CertIDs** | string | 证书ID列表 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|

