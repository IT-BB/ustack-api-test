# COMPANY

## 删除租户 - DeleteCompany

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [CreateUser DescribeUser DescribeMember] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 创建租户 - CreateUser

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Email** | string | 租户邮箱，租户注册时提供的用于接收通知和联系的邮箱 | [example@ucloud.cn] | [] | **Yes** |
| **UserEmail** | string | 租户邮箱，可用于系统告警通知 | [example@ucloud.cn] | [] | **Yes** |
| **PassWord** | string | 租户密码，密码须包含有小写字母、特殊字符(除空格)，。 | [mnSslwW6bkAaldaPUROcuA==] | [] | **Yes** |
| **Phone** | string | 租户手机号 | [138****0000] | [] | No |
| **Name** | string | 租户名称 | [user_name] | [] | **Yes** |
| **UserName** | string | 租户用户名 | [user_nick_name] | [] | **Yes** |
| **Audit** | string | 是否开启资源审批，取值范围：1、 1（是） 2、2（否） | [1] | [] | **Yes** |
| **AutoAudit** | string | 是否开启自动审批，取值范围：1、 1（是） 2、2（否） | [1] | [] | **Yes** |
| **ResetPassword** | string | 首次登录是否强制修改密码，取值：1, Y, Yes, YES, True, true 表示是；其他值表示否。用于增强账号安全性。 | [1 Y No] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **UserID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] |


## 冻结租户 - FreezeUser

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [CreateUser DescribeUser DescribeMember] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 解冻租户 - UnFreezeUser

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [CreateUser DescribeUser DescribeMember] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取租户列表 - DescribeUser

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Keyword** | string | 搜索关键词，支持租户名称、邮箱、手机号 | [ucloud] | [] | No |
| **Limit** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。 | [20] | [] | No |
| **Offset** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 | [0] | [] | No |
| **UserIDs** | array<int32> | 用户ID列表，用于指定查询的用户。查询多个用户用,分隔 | [200000246, 200000247] | [] | No |
| **CompanyIDs** | array<int32> | 租户ID列表，用于指定查询的租户。查询多个租户用,分隔。 | [200000233, 200000234] | [CreateUser] | No |
| **Region** | string | 所属的地域ID，一批可共享的物理资源使用集合。 | [develop] | [DescribeRegion] | No |
| **ProductType** | string | 产品类型 | [DISK] | [ListProductResources] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*TenantTenantInfoRes*](#TenantTenantInfoRes)> |  | [] |
| **TotalCount** | integer |  | [] |

### 数据模型

#### TenantTenantInfoRes

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] |
| **UserID** | integer | 用户ID，与租户ID相同，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] |
| **MemberID** | integer | 账号成员ID，用于标识资源所属的成员，实现多成员环境下的资源隔离。 | [200000246] |
| **Email** | string | 用户邮箱 | [example@ucloud.cn] |
| **Phone** | string | 用户手机号 | [138****0000] |
| **ManagerEmail** | string | 用户邮箱 | [example@ucloud.cn] |
| **Name** | string | 用户名称 | [user_name] |
| **Audit** | integer | 是否开启资源审批('1':是;'0':否) | [0] |
| **AutoAudit** | integer | 是否开启自动审批('1':是;'0':否) | [0] |
| **Status** | string | 租户状态，取值COMPANY_STATUS_NONE, AVAILABLE, FREEZE | [AVAILABLE] |
| **MemberStatus** | string | 账号状态，取值MEMBER_STATUS_NONE, AVAILABLE, FREEZE, LOCKED | [AVAILABLE] |
| **CreateTime** | integer | 租户创建时间 | [1694502400] |
| **UpdateTime** | integer | 租户更新时间 | [1694502400] |
| **Amount** | double | 租户余额 | [1000000] |
| **AmountFree** | double | 租户免费额度 | [1000000] |
| **AmountCredit** | double | 租户信用额度 | [1000000] |
| **PublicKey** | string | 账号API公钥，用于API调用时的身份验证和签名验证，采用RSA算法生成。 | [MIIBIjANBgkqhkiG9w0BAQEFAAOCA******BCgKCAQEA] |
| **PrivateKey** | string | 账号API私钥，用于API调用时的请求签名，采用RSA算法生成。请妥善保管，避免泄露。 | [MIIEvQIBADANBgkqhkiG9w0BAQEF******cwggSjAgEAAoIBAQC] |

## 更新租户名称 - UpdateCompanyName

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [CreateUser DescribeUser DescribeMember] | **Yes** |
| **Name** | string | 租户名称，长度为1-50个字符，名称只能包含中英文、数字、点（.），下划线（_）和中划线（-） | [new_company_name] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 修改租户邮箱 - UpdateCompanyEmail

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Email** | string | 租户邮箱，租户注册时提供的用于接收通知和联系的邮箱 | [example@ucloud.cn] | [DescribeUser DescribeMember] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [CreateUser DescribeUser DescribeMember] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 设置用户登录IP白名单 - UpdateLoginWhitelist

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [CreateUser DescribeUser DescribeMember] | **Yes** |
| **Whitelist** | string | 白名单，即配置的 IP 地址/段客户端才可正常登录控制台或访问 API 。为空则不限制登录的IP地址，多个IP或网段用逗号分隔 | [192.168.1.1, 10.0.0.0/24] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取用户登录IP白名单 - DescribeLoginWhitelist

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [CreateUser DescribeUser DescribeMember] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*IAMLoginWhitelist*](#IAMLoginWhitelist)> | 登录白名单详情 | [] |

### 数据模型

#### IAMLoginWhitelist

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [] |
| **Whitelist** | string | 白名单，配置的 IP 地址/段客户端才可正常登录控制台或访问 API 。为空则不限制登录的IP地址。 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |
| **WhitelistID** | string | 白名单ID，白名单的唯一标识。 | [] |

## 获取租户资源 - DescribeTenantResources

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [CreateUser DescribeUser DescribeMember] | **Yes** |
| **Limit** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。 | [20] | [] | No |
| **Offset** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 | [0] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总记录数 | [] |
| **Infos** | array<[*TenantResourceInfo*](#TenantResourceInfo)> | 租户关联资源列表 | [] |

### 数据模型

#### TenantResourceInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ResourceID** | string | 资源ID，要挂载ISO镜像的目标资源的唯一标识符。如虚拟机ID，格式为对应资源类型前缀加随机字符串。 | [vm-2aa5tms6u9duc5] |
| **ResourceType** | string | 资源类型 | [VM] |
| **Status** | string | 资源状态 | [Available] |
| **Name** | string | 资源名称 | [MyVM] |
| **IsRecycled** | bool | 该资源是否支持回收站 | [true] |
