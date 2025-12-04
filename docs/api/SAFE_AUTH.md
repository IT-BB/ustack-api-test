# SAFE_AUTH

## 查询用户安全设置 - DescribeSafeAuthFlag

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **MemberID** | integer | 账号ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **MemberID** | integer | 账号ID | [] |
| **Email** | string | 租户邮箱 | [] |
| **IsLoginVerificationEnable** | bool | 登录保护是否启用 | [] |
| **IsActionVerificationEnable** | bool | 敏感操作保护是否启用 | [] |
| **LoginFlag** | object | 登录保护设置 | [] |
| **ActionFlag** | object | 敏感操作保护设置 | [] |

### 数据模型

#### LoginActionFlag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Phone** | integer | 安全手机校验, 0: 非安全手机校验, 1: 安全手机校验 | [] |
| **Token** | integer | Token 校验, 0: 非 Token 校验 1: Token 校验 | [] |
| **Email** | integer | 安全邮箱校验, 0:非邮箱校验, 1:邮箱校验 | [] |
