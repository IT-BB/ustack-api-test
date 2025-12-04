# ALERTING

## 创建告警模版 - CreateAlertTemplate

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 所属的地域，一批可共享的物理资源使用集合。 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离。 | [200000233] | [DescribeUser] | No |
| **TemplateType** | string | 告警模版类型，表明创建的是哪一个资源的告警模板 | [VM] | [DescribeMetric] | **Yes** |
| **Name** | string | 告警模板名称，名称只能包含中英文、数字、点（.），下划线（_）和中划线（-），且必须以字母开头和结尾 | [vm-alert-name] | [] | **Yes** |
| **Remark** | string | 备注，用于对告警模板进行说明和注释，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [这是一个告警模板] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TemplateID** | string | 告警模板ID，用于唯一标识告警模板 | [tpl-qbw3th****eka9] |


## 更新告警模版 - UpdateAlertTemplate

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **TemplateID** | string | 告警模板ID，用于唯一标识告警模板 | [tpl-qbw3th****eka9] | [DescribeAlertTemplate CreateAlertTemplate] | **Yes** |
| **Name** | string | 告警模板名称，名称只能包含中英文、数字、点（.），下划线（_）和中划线（-），且必须以字母开头和结尾 | [vm-alert-name] | [] | **Yes** |
| **Remark** | string | 备注，用于对告警模板进行说明和注释，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [这是一个告警模板] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取告警模版 - DescribeAlertTemplate

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **TemplateID** | string | 告警模板ID，用于唯一标识告警模板 | [tpl-qbw3th****eka9] | [CreateAlertTemplate] | No |
| **Region** | string | 所属的地域，一批可共享的物理资源使用集合。 | [develop] | [DescribeRegion] | No |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离。 | [200000233] | [DescribeUser] | No |
| **TemplateTypes** | array<string> | 告警模版类型，表明创建的是哪一个资源的告警模板，多个类型用,分隔 | [VM, FS] | [DescribeMetric] | No |
| **Limit** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。 | [20] | [] | No |
| **Offset** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 | [0] | [] | No |
| **Keyword** | string | 搜索关键字，用于根据名称或备注进行模糊搜索。 | [vm] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数 | [] |
| **Infos** | array<[*AlertTemplateInfo*](#AlertTemplateInfo)> | 告警模板内容 | [] |

### 数据模型

#### AlertTemplateInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TemplateID** | string | 告警模板ID，用于唯一标识告警模板 | [tpl-qbw3th****eka9] |
| **Region** | string | 所属的地域ID | [develop] |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离。 | [200000233] |
| **Email** | string | 租户邮箱，用于登录与表明租户信息 | [example@ucloud.cn] |
| **TemplateType** | string | 告警模板类型，表明创建的是哪一个资源的告警模板 | [VM] |
| **Name** | string | 告警模板名称 | [vm-alert-name] |
| **Remark** | string | 告警模板备注 | [这是一个告警模板] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |
| **BoundTargetCount** | integer | 绑定资源数 | [] |

## 删除告警模版 - DeleteAlertTemplate

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **TemplateID** | string | 告警模板ID，用于唯一标识告警模板 | [tpl-qbw3th****eka9] | [DescribeAlertTemplate] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 创建告警通知组 - CreateAlertNotifyGroup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 所属的地域，一批可共享的物理资源使用集合。 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离。 | [200000233] | [DescribeUser] | No |
| **Name** | string | 告警通知组名称，名称只能包含中英文、数字、点（.），下划线（_）和中划线（-） | [vm-alert-group] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **NotifyGroupID** | string | 告警通知组ID，用于唯一标识告警通知组 | [ng-qbw3th****eka9] |


## 更新告警通知组 - UpdateAlertNotifyGroup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **NotifyGroupID** | string | 告警通知组ID，用于唯一标识告警通知组 | [ng-qbw3th****eka9] | [DescribeAlertNotifyGroup CreateAlertNotifyGroup] | **Yes** |
| **Name** | string | 告警通知组名称，名称只能包含中英文、数字、点（.），下划线（_）和中划线（-） | [vm-alert-group] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 删除告警通知组 - DeleteAlertNotifyGroup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **NotifyGroupID** | string | 告警通知组ID，用于唯一标识告警通知组 | [ng-qbw3th****eka9] | [DescribeAlertNotifyGroup CreateAlertNotifyGroup] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取告警通知组 - DescribeAlertNotifyGroup

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **NotifyGroupID** | string | 告警通知组ID，用于唯一标识告警通知组 | [ng-qbw3th****eka9] | [CreateAlertNotifyGroup] | No |
| **Region** | string | 所属的地域，一批可共享的物理资源使用集合。 | [develop] | [DescribeRegion] | No |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离。 | [200000233] | [DescribeUser] | No |
| **Limit** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。 | [20] | [] | No |
| **Offset** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 | [0] | [] | No |
| **Keyword** | string | 搜索关键字，用于根据名称或备注进行模糊搜索。 | [vm] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 告警通知组总数 | [] |
| **Infos** | array<[*AlertNotifyGroupInfo*](#AlertNotifyGroupInfo)> | 通知组信息 | [] |

### 数据模型

#### AlertNotifyGroupInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **NotifyGroupID** | string | 通知组ID，用于唯一标识告警通知组 | [ng-qbw3th****eka9] |
| **Region** | string | 所属的地域 | [develop] |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离。 | [200000233] |
| **Email** | string | 租户邮箱，用于登录与表明租户信息 | [example@ucloud.cn] |
| **Name** | string | 通知组名称 | [vm-alert-group] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |

## 创建告警通知人 - CreateAlertNotifyReceiver

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **NotifyGroupID** | string | 告警通知组ID，用于唯一标识告警通知组 | [ng-qbw3th****eka9] | [DescribeAlertNotifyGroup CreateAlertNotifyGroup] | **Yes** |
| **Name** | string | 名称，用于标识告警通知接收人 | [ucloud] | [] | **Yes** |
| **Email** | string | 邮件，用于告警通知的发送 | [example@ucloud.cn] | [] | **Yes** |
| **Phone** | string | 电话，用于告警通知的发送 | [138****0000] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **NotifyReceiverID** | string | 通知接收人ID，用于唯一标识告警通知接收人 | [nr-qbw3th****eka9] |


## 更新告警通知人 - UpdateAlertNotifyReceiver

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **NotifyReceiverID** | string | 通知接收人ID，用于唯一标识告警通知接收人 | [nr-qbw3th****eka9] | [DescribeAlertNotifyReceiver CreateAlertNotifyReceiver] | **Yes** |
| **Name** | string | 名称，用于标识告警通知接收人 | [ucloud] | [] | **Yes** |
| **Email** | string | 邮件，用于告警通知的发送 | [example@ucloud.cn] | [] | **Yes** |
| **Phone** | string | 电话，用于告警通知的发送 | [138****0000] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 删除告警通知人 - DeleteAlertNotifyReceiver

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **NotifyReceiverID** | string | 通知接收人ID，用于唯一标识告警通知接收人 | [nr-qbw3th****eka9] | [DescribeAlertNotifyReceiver CreateAlertNotifyReceiver] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取告警通知人 - DescribeAlertNotifyReceiver

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **NotifyGroupID** | string | 通知组ID，用于唯一标识告警通知组 | [ng-qbw3th****eka9] | [DescribeAlertNotifyGroup CreateAlertNotifyGroup] | No |
| **Limit** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。 | [20] | [] | No |
| **Offset** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 | [0] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 资源总数 | [] |
| **Infos** | array<[*AlertNotifyReceiverInfo*](#AlertNotifyReceiverInfo)> | 通知接收人信息 | [] |

### 数据模型

#### AlertNotifyReceiverInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **NotifyGroupID** | string | 通知组ID，用于唯一标识告警通知组，也表明通知接收人所属通知组 | [ng-qbw3th****eka9] |
| **NotifyReceiverID** | string | 通知接收人ID，用于唯一标识告警通知接收人 | [nr-qbw3th****eka9] |
| **Name** | string | 通知接收人名称 | [ucloud] |
| **Email** | string | 通知接收人邮件 | [example@ucloud.cn] |
| **Phone** | string | 通知接收人电话 | [138****0000] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |

## 创建告警模版规则 - CreateAlertTemplateRule

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **TemplateID** | string | 告警模板ID，用于唯一标识告警模板 | [tql-qbw3th****eka9] | [DescribeAlertTemplate] | **Yes** |
| **NotifyGroupID** | string | 通知组ID，用于唯一标识告警通知组 | [ng-qbw3th****eka9] | [DescribeAlertNotifyGroup CreateAlertNotifyGroup] | **Yes** |
| **Metric** | string | 告警规则名称 | [vm_cpu_utilization] | [] | **Yes** |
| **Query** | string | 表达式 | [示例] | [] | **Yes** |
| **ConditionType** | string | 条件类型，取值范围是GE（大于等于）,LE（小于等于）,EQ（等于）,NE（不等于）,LT（小于）,GT（大于） | [GE LE EQ NE] | [] | **Yes** |
| **Thresholds** | double | 指标阈值 | [80] | [] | **Yes** |
| **ForSeconds** | integer | 告警持续时间，单位：秒 | [60] | [] | **Yes** |
| **Summary** | string | 告警规则描述 | [CPU使用率>=80%] | [] | **Yes** |
| **Severity** | string | 告警级别，取值范围：1、warning（一般告警）2、 critical（危险告警）3、error（重要告警） | [warning] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **RuleID** | string | 告警规则ID，用于唯一标识告警规则 | [rule-e7gv****xz0anv] |


## 更新告警通知规则 - UpdateAlertTemplateRule

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **RuleID** | string | 告警规则ID，用于唯一标识告警规则 | [rule-e7gv****xz0anv] | [DescribeAlertTemplateRule CreateAlertTemplateRule] | **Yes** |
| **NotifyGroupID** | string | 通知组ID，用于唯一标识告警通知组 | [ng-qbw3th****eka9] | [DescribeAlertNotifyGroup CreateAlertNotifyGroup] | **Yes** |
| **Metric** | string | 告警规则名称 | [vm_cpu_utilization] | [] | **Yes** |
| **Query** | string | 表达式 | [示例] | [] | **Yes** |
| **ConditionType** | string | 条件类型，取值范围是GE（大于等于）,LE（小于等于）,EQ（等于）,NE（不等于）,LT（小于）,GT（大于） | [GE LE EQ NE] | [] | **Yes** |
| **Thresholds** | double | 指标阈值 | [50] | [] | **Yes** |
| **ForSeconds** | integer | 告警持续时间，单位：秒 | [60] | [] | **Yes** |
| **Summary** | string | 告警规则描述 | [CPU使用率>=80%] | [] | **Yes** |
| **Severity** | string | 告警级别，取值范围：1、warning（一般告警）2、 critical（危险告警）3、error（重要告警） | [warning] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 删除告警模版规则 - DeleteAlertTemplateRule

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **RuleID** | string | 规则ID，指定删除的告警规则 | [rule-e7gv****xz0anv] | [DescribeAlertTemplateRule CreateAlertTemplateRule] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取告警模版规则 - DescribeAlertTemplateRule

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **TemplateID** | string | 告警模板ID，用于唯一标识告警模板 | [tql-qbw3th****eka9] | [CreateAlertTemplateRule] | No |
| **Limit** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。 | [20] | [] | No |
| **Offset** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 | [0] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 资源总数 | [] |
| **Infos** | array<[*AlertTemplateRuleInfo*](#AlertTemplateRuleInfo)> | 告警规则信息 | [] |

### 数据模型

#### AlertTemplateRuleInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **RuleID** | string | 告警规则ID，用于唯一标识告警规则 | [] |
| **TemplateID** | string | 告警模板ID，用于唯一标识告警模板 | [] |
| **NotifyGroupID** | string | 通知组ID，用于唯一标识告警通知组 | [] |
| **NotifyGroupName** | string | 通知组名称，用于标识告警通知组的名称 | [] |
| **Metric** | string | 告警指标，用于标识告警规则的指标 | [] |
| **Query** | string | 告警表达式，用于定义告警规则的触发条件 | [] |
| **ConditionType** | string | 条件类型，取值范围是GE（大于等于）,LE（小于等于）,EQ（等于）,NE（不等于）,LT（小于）,GT（大于） | [] |
| **Thresholds** | double | 告警阈值，用于定义告警规则的触发条件 | [] |
| **ForSeconds** | integer | 持续时间，单位秒，用于定义告警规则的触发持续时间 | [] |
| **Summary** | string | 规则描述 | [CPU使用率>=50%] |
| **Severity** | string | 告警级别，取值范围：1、warning（一般告警）2、 critical（危险告警）3、error（重要告警） | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |

## 创建操作日志通知规则 - CreateOPLogNotifyRule

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离。 | [26765446] | [DescribeUser] | No |
| **NotifyGroupID** | string | 通知组ID，用于标识通知规则所属的通知组 | [ng-ar5em6l8n2oyzg] | [DescribeAlertNotifyGroup] | **Yes** |
| **MonitorRegion** | string | 监视地域，一批可共享的物理资源使用集合 | [cn] | [DescribeRegion] | **Yes** |
| **MonitorLevel** | array<string> | 监控级别，监控操作成功（success），操作失败（failure），多个值用,分隔 | [success] | [] | **Yes** |
| **MonitorModule** | array<string> | 监控模块，表明需要监控那些模块的操作，多个值用,分隔 | [VM, DISK, IMAGE] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **RuleID** | string | 规则ID，操作日志通知规则ID | [] |


## 更新操作日志通知规则 - UpdateOPLogNotifyRule

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **RuleID** | string | 规则ID，操作日志通知规则ID | [rule-yd8lnm53ytv3sw] | [CreateOPLogNotifyRule DescribeOPLogNotifyRule] | **Yes** |
| **NotifyGroupID** | string | 通知组ID，用于标识通知规则所属的通知组 | [ng-ar5em6l8n2oyzg] | [DescribeAlertNotifyGroup] | **Yes** |
| **MonitorRegion** | string | 监视地域，一批可共享的物理资源使用集合 | [cn] | [DescribeRegion] | **Yes** |
| **MonitorLevel** | array<string> | 监控级别，监控操作成功（success），操作失败（failure），多个值用,分隔 | [success] | [] | **Yes** |
| **MonitorModule** | array<string> | 监控模块，表明需要监控那些模块的操作，多个值用,分隔 | [VM, DISK, IMAGE] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 删除操作日志通知规则 - DeleteOPLogNotifyRule

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **RuleID** | string | 规则ID，操作日志通知规则ID | [rule-yd8lnm53ytv3sw] | [CreateOPLogNotifyRule DescribeOPLogNotifyRule] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取操作日志通知规则 - DescribeOPLogNotifyRule

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离。 | [26765446] | [DescribeUser] | No |
| **Limit** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。 | [20] | [] | No |
| **Offset** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 | [0] | [] | No |
| **Keyword** | string | 搜索关键词 | [example_key] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 资源总数 | [] |
| **Infos** | array<[*OPLogNotifyRuleInfo*](#OPLogNotifyRuleInfo)> | 操作日志通知规则详情 | [] |

### 数据模型

#### OPLogNotifyRuleInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户 | [] |
| **Email** | string | 租户邮箱 | [] |
| **RuleID** | string | 规则ID | [] |
| **NotifyGroupID** | string | 通知组ID | [] |
| **NotifyGroupName** | string | 通知组名称 | [] |
| **MonitorRegion** | string | 监控地域 | [] |
| **MonitorLevel** | array<string> | 监控等级 | [] |
| **MonitorModule** | array<string> | 监控模块 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |
| **MonitorRegionAlias** | string | 监控地域别名 | [] |

## 创建资源事件通知规则 - CreateResourceEventNotifyRule

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离。 | [26765446] | [DescribeUser] | No |
| **NotifyGroupID** | string | 通知组ID，用于标识资源事件通知规则所属的通知组 | [ng-ar5em6l8n2oyzg] | [DescribeAlertNotifyGroup] | **Yes** |
| **MonitorRegion** | string | 监视地域，一批可共享的物理资源使用集合 | [cn] | [DescribeRegion] | **Yes** |
| **MonitorLevel** | array<string> | 监控级别，取值为正常（Info），警告（Warning），错误（Error），多种监控等级用，分隔 | [Info, Warning] | [] | **Yes** |
| **MonitorModule** | array<string> | 监控模块，表明需要监控那些模块的操作，多个值用,分隔 | [VM, DISK, IMAGE] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **RuleID** | string | 规则ID，用于标识资源事件通知规则 | [rule-6co6bhc1pu3ca5] |


## 更新资源事件通知规则 - UpdateResourceEventNotifyRule

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **RuleID** | string | 规则ID，用于标识资源事件通知规则 | [] | [DescribeResourceEventNotifyRule CreateResourceEventNotifyRule] | **Yes** |
| **NotifyGroupID** | string | 通知组ID，用于标识资源事件通知规则所属的通知组 | [ng-ar5em6l8n2oyzg] | [DescribeAlertNotifyGroup] | **Yes** |
| **MonitorRegion** | string | 监视地域，一批可共享的物理资源使用集合 | [cn] | [DescribeRegion] | **Yes** |
| **MonitorLevel** | array<string> | 监控级别，取值为正常（Info），警告（Warning），错误（Error），多种监控等级用，分隔 | [Info, Warning] | [] | **Yes** |
| **MonitorModule** | array<string> | 监控模块，表明需要监控那些模块的操作，多个值用,分隔 | [VM, DISK, IMAGE] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 删除资源事件通知规则 - DeleteResourceEventNotifyRule

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **RuleID** | string | 规则ID，用于标识资源事件通知规则 | [] | [DescribeResourceEventNotifyRule CreateResourceEventNotifyRule] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取资源事件通知规则 - DescribeResourceEventNotifyRule

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离。 | [26765446] | [DescribeUser] | No |
| **Limit** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。 | [20] | [] | No |
| **Offset** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 | [0] | [] | No |
| **Keyword** | string | 搜索关键词 | [example_key] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 资源总数 | [] |
| **Infos** | array<[*ResourceEventNotifyRuleInfo*](#ResourceEventNotifyRuleInfo)> | 资源事件通知规则详情 | [] |

### 数据模型

#### ResourceEventNotifyRuleInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离，可以通过 DescribeUser 获取。 | [200000240] |
| **Email** | string | 租户邮箱地址，作为租户的联系邮箱，用于接收系统通知和告警信息。 | [example@ucloud.cn] |
| **RuleID** | string | 规则ID，资源事件通知规则的唯一标识符，格式为 rule- 前缀加随机字符串。 | [rule-6co6bhc1pu3ca5] |
| **NotifyGroupID** | string | 通知组ID，用于标识通知组的唯一标识符，格式为 ng- 前缀加随机字符串。 | [ng-ar5em6l8n2oyzg] |
| **NotifyGroupName** | string | 通知组名称，用于在界面中展示的通知组名称。支持中英文、数字和常见符号，长度1-50个字符。 | [test] |
| **MonitorRegion** | string | 监控地域，指定要监控的地域标识符。您可以调用 DescribeRegion 方法查看最新的地域列表。 | [cn] |
| **MonitorLevel** | array<string> | 监控级别列表，指定要监控的告警级别。取值范围：Info（信息级别）、Warning（警告级别）、Error（错误级别）、Critical（严重级别）。 | [info, warning] |
| **MonitorModule** | array<string> | 监控模块列表，指定要监控的资源模块类型。可以通过 ListProductResources 获取支持的模块类型。 | [VM, FS, REDIS] |
| **CreateTime** | integer | 创建时间，Unix时间戳格式（秒级），记录规则首次创建的时间。 | [1755488912] |
| **UpdateTime** | integer | 更新时间，Unix时间戳格式（秒级），记录规则最后一次修改的时间。 | [1755488912] |
| **MonitorRegionAlias** | string | 监控地域别名，用于在界面中展示的地域友好名称，便于用户识别和管理。 | [cn] |

## 绑定告警模版 - BindAlertTemplate

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **TemplateID** | string | 告警模板ID，用于唯一标识告警模板 | [tpl-qbw3thu38ieka9] | [CreateAlertTemplate DescribeAlertTemplate] | **Yes** |
| **TargetID** | string | 目标资源ID，即为哪个资源绑定告警模板 | [vm-m3wz2711abiug8] | [] | **Yes** |
| **Region** | string | 地域，可以通过DescribeRegions接口获取 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 解绑告警模版 - UnbindAlertTemplate

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **TargetID** | string | 目标资源ID，即为哪个资源取消绑定告警模板 | [vm-m3wz2711abiug8] | [] | **Yes** |
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [cn] | [DescribeRegion] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取告警模版绑定目标 - DescribeAlertTemplateTarget

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **TemplateID** | string | 模板ID，用于唯一标识告警模板 | [tpl-qbw3thu38ieka9] | [CreateAlertTemplate] | No |
| **TargetID** | string | 绑定资源ID，即为哪个资源绑定了告警模板 | [vm-m3wz2711abiug8] | [] | No |
| **Limit** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。 | [20] | [] | No |
| **Offset** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 | [0] | [] | No |
| **Region** | string | 所属的地域，一批可共享的物理资源使用集合。 | [develop] | [DescribeRegion] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 资源总数 | [] |
| **Infos** | array<[*AlertTemplateTargetInfo*](#AlertTemplateTargetInfo)> | 绑定关系信息 | [] |

### 数据模型

#### AlertTemplateTargetInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TemplateID** | string | 告警模板ID，用于唯一标识告警模板 | [tpl-qbw3thu38ieka9] |
| **TemplateName** | string | 告警模板名称 | [] |
| **TargetID** | string | 绑定资源ID，即为哪个资源绑定了告警模板 | [vm-m3wz2711abiug8] |
| **TargetName** | string | 绑定资源名称 | [] |

## 查询告警 - DescribeAlert

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 所属的地域，一批可共享的物理资源使用集合。 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离。 | [200000233] | [DescribeUser] | No |
| **Severities** | array<string> | 告警级别，取值范围：1、warning（一般告警）2、 critical（危险告警）3、error（重要告警） | [warning] | [] | No |
| **States** | array<string> | 告警状态，取值范围是：pending（待触发）, firing（触发中），inactive（未触发） | [firing] | [] | No |
| **Limit** | integer | 分页参数，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。 | [20] | [] | No |
| **Offset** | integer | 分页参数，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 | [0] | [] | No |
| **TemplateTypes** | array<string> | 告警模板类型，表明创建的是哪一个资源的告警模板，可以通过DescribeMetric获取，多个模板类型用,分隔 | [VM, FS] | [] | No |
| **ProjectIDs** | array<string> | 项目组，用于标识资源所属的项目组，实现项目组级别的资源隔离，可以通过 ListProjects 获取。 | [project-51chy****squbp] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*AlertInfo*](#AlertInfo)> | 告警详情 | [] |
| **TotalCount** | integer | 告警总数 | [] |

### 数据模型

#### AlertInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ActiveAt** | integer | 活跃时间，Unix时间戳格式（秒级），表示告警首次触发或最后一次活跃的时间 | [1755065626] |
| **State** | string | 告警状态，表示当前告警的触发状态。取值：inactive（未触发）、pending（待触发）、firing（触发中） | [firing] |
| **Thresholds** | double | 阈值，告警规则中设置的触发阈值，当监控指标达到此值时触发告警 | [1] |
| **Value** | double | 当前值，监控指标的实时数值，用于与阈值进行比较判断是否触发告警 | [127.343199] |
| **TargetID** | string | 资源ID，触发告警的目标资源的唯一标识符，用于定位具体的监控对象 | [storage-set-rwio] |
| **TargetName** | string | 资源名称，触发告警的目标资源的显示名称，便于用户识别和管理 | [] |
| **Metric** | string | 告警指标，触发告警的监控指标名称，如CPU使用率、内存使用率、存储使用率等 | [] |
| **CompanyID** | integer | 租户ID，告警所属租户的唯一标识，用于实现多租户环境下的告警隔离，可以通过 DescribeUser 获取 | [200000231] |
| **Email** | string | 租户邮箱，告警所属租户的邮箱地址，用于发送告警通知邮件 | [admin@cloud.cn] |
| **Summary** | string | 描述，告警的详细描述信息，包含告警的具体内容和触发条件说明 | [存储分配率 >= 1 %] |
| **Severity** | string | 告警级别，表示告警的严重程度。取值：info（信息）、warning（警告）、error（错误）、critical（严重） | [error] |
| **TemplateType** | string | 模板类型，告警规则使用的模板类型，用于分类管理不同类型的告警规则 | [STORAGE_SET] |
| **LabelSet** | array<[*AlertLabelSet*](#AlertLabelSet)> | 告警标签，包含告警相关的标签信息，用于告警的分类、过滤和路由 | [] |
| **RegionName** | string | 地域名称，告警所属的地域标识，表示告警发生的物理位置。您可以调用 DescribeRegion 方法查看最新的地域列表 | [cn] |
| **Region** | string | 地域 | [] |

#### AlertLabelSet

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string | 标签名 | [] |
| **Value** | string | 标签值 | [] |

## 创建告警回调接口 - CreateAlertNotifyWebhook

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **NotifyGroupID** | string | 通知组ID，用于标识通知规则所属的通知组 | [ng-ar5em6l8n2oyzg] | [DescribeAlertNotifyGroup CreateAlertNotifyGroup] | **Yes** |
| **Name** | string | 名称，Webhook通知规则的名称，用于标识和区分不同的Webhook配置，长度为1-50个字符，名称只能包含中英文、数字、点（.），下划线（_）和中划线（-） | [webhook_name] | [] | **Yes** |
| **Url** | string | URL链接，Webhook的目标URL链接，当告警触发时系统将向该地址发送HTTP请求 | [https://example.com/webhook] | [] | **Yes** |
| **Method** | string | HTTP请求方法类型，支持POST和GET两种方式，用于指定向Webhook URL发送请求时使用的HTTP方法 | [POST] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **NotifyWebhookID** | string | 通知WebhookID，用于标识通知规则的唯一标识符，格式为 webhook- 前缀加随机字符串。 | [webhook-6co6bhc1pu3ca5] |


## 更新告警回调接口 - UpdateAlertNotifyWebhook

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **NotifyWebhookID** | string | Webhook通知规则的ID，用于标识要更新的Webhook配置 | [webhook-6co6bhc1pu3ca5] | [CreateAlertNotifyWebhook DescribeAlertNotifyWebhook] | **Yes** |
| **Name** | string | 名称，Webhook通知规则的名称，用于标识和区分不同的Webhook配置，长度为1-50个字符，名称只能包含中英文、数字、点（.），下划线（_）和中划线（-） | [webhook_name] | [] | **Yes** |
| **Url** | string | URL链接，Webhook的目标URL链接，当告警触发时系统将向该地址发送HTTP请求 | [https://example.com/webhook] | [] | **Yes** |
| **Method** | string | HTTP请求方法类型，支持POST和GET两种方式，用于指定向Webhook URL发送请求时使用的HTTP方法 | [POST] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 删除告警回调接口 - DeleteAlertNotifyWebhook

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **NotifyWebhookID** | string | 通知WebhookID，用于标识要删除的Webhook配置 | [webhook-6co6bhc1pu3ca5] | [CreateAlertNotifyWebhook DescribeAlertNotifyWebhook] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取告警回调接口 - DescribeAlertNotifyWebhook

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **NotifyGroupID** | string | 通知组ID，用于标识通知规则所属的通知组 | [ng-ar5em6l8n2oyzg] | [DescribeAlertNotifyGroup CreateAlertNotifyGroup] | **Yes** |
| **Limit** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。 | [20] | [] | No |
| **Offset** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 | [0] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数 | [] |
| **Infos** | array<[*AlertNotifyWebhookInfo*](#AlertNotifyWebhookInfo)> | 详情 | [] |

### 数据模型

#### AlertNotifyWebhookInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **NotifyGroupID** | string | 通知组ID | [] |
| **NotifyWebhookID** | string | 通知组WebhookID | [] |
| **Name** | string | 名称 | [] |
| **Url** | string | URL | [] |
| **Method** | string | HTTP请求的方法 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |

## 获取监控指标 - DescribeMetric

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **TargetTypes** | array<string> | 资源类型，监控对象的类型，例如：CLUSTER、NODE、FS、OSD、BLOCK、OBJECT等 | [VM] | [] | **Yes** |
| **TargetID** | string | 监控资源ID，监控对象的唯一标识符，根据监控类型的不同传入对应的监控资源ID等 | [vm-12safafsadf3] | [] | **Yes** |
| **Labels** | array<string> | 标签，用于过滤监控指标，弹性伸缩监控在不过滤的情况下会返回所有伸缩类型的监控，Labels的作用是用于过滤 | [horizontal_vm] | [] | No |
| **Region** | string | 地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Metrics** | array<[*Metric*](#Metric)> | 监控指标列表，包含系统中所有可用的监控指标信息 | [] |

### 数据模型

#### Enums

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Title** | string | 枚举项标题，用于在界面上显示的枚举选项名称 | [] |
| **Value** | string | 枚举项的实际值，用于系统内部处理和存储的枚举值 | [] |

#### Metric

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TargetType** | string | 资源类型，标识该监控指标所监控的资源类型，如虚拟机、磁盘、网络等 | [] |
| **IsAlerting** | bool | 是否支持告警功能，标识该指标是否可以用于配置告警规则 | [] |
| **Infos** | array<[*MetricInfo*](#MetricInfo)> | 监控指标详细信息列表，包含该资源类型下所有可用的具体监控指标 | [] |
| **Authority** | array<string> | 权限列表，定义哪些角色或用户可以访问和使用该监控指标 | [] |
| **UseYCallback** | bool | 是否需要单位转换回调，标识在显示指标数据时是否需要进行单位转换处理 | [] |

#### MetricFilter

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Description** | string | 过滤器的描述信息，说明该过滤器的作用和用途 | [] |
| **Condition** | object | 过滤器的具体条件配置，定义过滤的键值对和操作方式 | [] |

#### MetricFilterCondition

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string | 过滤条件的键名，指定要过滤的字段或标签名称 | [] |
| **Operation** | string | 过滤操作类型，定义过滤时使用的比较操作，如等于、包含、正则匹配等 | [] |
| **Values** | array<string> | 过滤条件的值列表，指定过滤时要匹配的具体值 | [] |

#### MetricInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **MetricID** | string | 监控指标的唯一标识符，用于在系统中唯一标识一个监控指标 | [] |
| **IsAlerting** | bool | 是否支持告警功能，标识该具体指标是否可以用于配置告警规则 | [] |
| **Name** | string | 指标的中文名称和含义描述，用于在界面上显示给用户理解 | [] |
| **Unit** | string | 指标数据的计量单位，如百分比、字节、次数等 | [] |
| **Metric** | string | Prometheus查询语句，用于从监控系统中获取该指标的实际数据 | [] |
| **RangeMin** | string | 指标数据的最小有效值，用于数据验证和界面显示范围限制 | [] |
| **RangeMax** | string | 指标数据的最大有效值，用于数据验证和界面显示范围限制 | [] |
| **CompareType** | string | 告警阈值对比类型，定义告警触发时的比较方式，如大于、小于、等于等 | [] |
| **Enums** | array<[*Enums*](#Enums)> | 枚举值列表，当指标为枚举类型时提供可选的枚举项 | [] |
| **Filters** | array<[*MetricFilter*](#MetricFilter)> | 指标过滤器列表，用于对监控数据进行筛选和过滤的条件配置 | [] |
| **Labels** | array<string> | 指标标签列表，用于对监控指标进行分类和标记的标签信息 | [] |
| **UseYCallback** | bool | 是否需要单位转换回调处理，标识在显示该指标数据时是否需要进行单位转换 | [] |
| **Categories** | string | 指标分组类别，用于在界面上对监控指标进行分类显示和管理 | [] |

## 获取即时查询监控数据 - PrometheusQuery

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域标识，指定Prometheus服务所在的地域，用于获取对应地域的Prometheus客户端 | [cn] | [DescribeRegion] | **Yes** |
| **Query** | string | 查询表达式，Prometheus查询表达式，支持PromQL语法，用于查询监控指标数据 | [up{job="prometheus"}] | [] | **Yes** |
| **Time** | string | 查询时间点，支持Unix时间戳或RFC3339格式，为空时使用当前时间 | [1640995200] | [] | No |
| **Timeout** | string | 查询超时时间，支持Go duration格式（如30s、1m），用于控制查询执行时间 | [30s] | [] | No |
| **Limit** | integer | 查询结果数量，用于控制返回的数据点数量，避免结果集过大 | [1000] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取范围查询监控数据 - PrometheusQueryRange

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域标识，指定Prometheus服务所在的地域，用于获取对应地域的Prometheus客户端 | [cn] | [DescribeRegion] | **Yes** |
| **Query** | string | 查询表达式，Prometheus查询表达式，支持PromQL语法，用于查询监控指标数据 | [up{job="prometheus"}] | [] | **Yes** |
| **Start** | string | 查询开始时间，支持Unix时间戳或RFC3339格式，定义时间范围查询的起始点 | [1640995200] | [] | **Yes** |
| **End** | string | 查询结束时间，支持Unix时间戳或RFC3339格式，定义时间范围查询的结束点 | [1640998800] | [] | **Yes** |
| **Step** | string | 查询步长，支持数字（秒）或Go duration格式，定义数据点之间的时间间隔 | [60] | [] | **Yes** |
| **Timeout** | string | 查询超时时间，支持Go duration格式（如30s、1m），用于控制查询执行时间 | [30s] | [] | No |
| **Limit** | integer | 查询结果数量，用于控制返回的时间序列数据点数量，避免结果集过大 | [1000] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|

