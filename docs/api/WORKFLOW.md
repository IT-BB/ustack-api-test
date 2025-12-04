# WORKFLOW

## 查询审批列表 - DescribeApplicationNode

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [CreateUser DescribeUser DescribeMember] | No |
| **MemberID** | integer | 子账号ID，用于表示子账号的唯一标识，用于标识资源所属的成员，实现多成员环境下的资源隔离 | [200000234] | [CreateSubMember DescribeMember GetMemberInfo] | No |
| **States** | array<string> | 审核节点状态，已办，传Approved, Rejected。待办，传Handling | [Approved,Rejected Handling] | [] | No |
| **Keyword** | string | 搜索关键字 | [] | [] | No |
| **Limit** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。 | [20] | [] | No |
| **Offset** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 | [0] | [] | No |
| **Sort** | string | 排序方式，降序排列(Descending)或升序排列(Ascending) | [Descending Ascending] | [] | No |
| **SortBy** | string | 排序字段，根据指定字段进行排序。 | [CreateTime] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*WorkflowApplicationNode*](#WorkflowApplicationNode)> | 审核节点信息 | [] |
| **TotalCount** | integer | 总数 | [] |

### 数据模型

#### ObserverInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ObserverMemberID** | integer | 观察人ID | [200000237] |
| **ObserverEmail** | string | 观察人邮箱 | [admin@example.com] |

#### WorkflowApplicationNode

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ApplicationMemberID** | integer | 申请者账户ID，子账号申请，则为该子账号ID | [200000247] |
| **ApplicationEmail** | string | 申请者账户邮箱 | [ucloud.sub@example.com] |
| **ApplicationID** | string | 审批工单ID | [apply-utkk9iljh7lpb1] |
| **ApplicationName** | string | 审批工单名称 | [new_application_name] |
| **ResourceType** | string | 审批资源类型，可以通过ListProductResources获取，可以通过ListProductResources获取，可以通过ListProductResources获取 | [DISK] |
| **APIName** | string | 审批资源对应API名称，例如创建磁盘操作，提交审批，APIName为CreateDisk | [CreateDisk] |
| **NodeID** | string | 审批节点ID | [1] |
| **NodeName** | string | 审批节点名称 | [节点1] |
| **State** | string | 审批节点状态，取值为Undo Committed Rejected Aborted Approved Handling Succeed Failed | [Handling] |
| **Remark** | string | 审批节点备注 | [这是审批备注] |
| **NodeOperator** | string | 节点审核人邮箱 | [example@ucloud.cn] |
| **NodeOperatorID** | integer | 节点审核人账号ID | [200000239] |
| **NodeOperatorCompanyID** | integer | 节点审核人租户ID | [200000233] |
| **CreateTime** | integer | 创建时间 | [1754968392] |
| **UpdateTime** | integer | 更新时间 | [1754968392] |
| **ObserverInfos** | array<[*ObserverInfo*](#ObserverInfo)> | 节点观察人信息 | [] |
| **AutoAudit** | bool | 节点是否自动审批 | [true] |

## 查询审批记录 - DescribeApplication

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [CreateUser DescribeUser DescribeMember] | No |
| **MemberID** | integer | 子账号ID，用于表示子账号的唯一标识，用于标识资源所属的成员，实现多成员环境下的资源隔离 | [200000234] | [CreateSubMember DescribeMember GetMemberInfo] | No |
| **ApplicationID** | string | 审批ID，子账号提交审批后，得到的审批ID，唯一标识一个审批流程 | [apply-utkk9iljh7lpb1] | [] | No |
| **ResourceType** | string | 产品类型 | [DISK] | [ListProductResources] | No |
| **Keyword** | string | 搜索关键字 | [] | [] | No |
| **Limit** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。 | [20] | [] | No |
| **Offset** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 | [0] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*WorkflowApplication*](#WorkflowApplication)> | 审批资源内容 | [] |
| **TotalCount** | integer | 资源数量 | [] |

### 数据模型

#### ObserverInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ObserverMemberID** | integer | 观察人ID | [200000237] |
| **ObserverEmail** | string | 观察人邮箱 | [admin@example.com] |

#### WorkflowApplication

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ApplicationID** | string | 审批工单ID | [apply-utkk9iljh7lpb1] |
| **WorkflowID** | string | 自定义流程ID | [workflow-utkk9iljh7lpb1] |
| **Name** | string | 审批名称 | [new_application_name] |
| **Reason** | string | 审批原因 | [这是审批理由] |
| **ResourceType** | string | 资源类型，可以通过ListProductResources获取 | [DISK] |
| **APIName** | string | 审批资源对应API名称，例如创建磁盘操作，提交审批，APIName为CreateDisk | [CreateDisk] |
| **Config** | string | API请求对应的参数 | [{"Action":["CreateDisk"],"ApplicationName":["application_name"],"ApplicationReason":["这是审批理由"],"ChargeType":["Month"],"Count":["1"],"DiskSpace":["10"],"Name":["application_disk"],"ProjectID":["project-ac6so62so01gry"],"Quantity":["1"],"Region":["pro2133"],"Remark":["备注"],"SetType":["StorageSetOUWI"],"ShareAble":["false"],"StorageType":["LocalDisk"],"_timestamp":["1754968392169"]}] |
| **Email** | string | 租户邮箱 | [example@ucloud.cn] |
| **MemberID** | integer | 申请人子账户ID | [200000247] |
| **CompanyID** | integer | 申请人租户ID | [200000233] |
| **Region** | string | 地域 | [pro2133] |
| **State** | string | 审批工单当前审核状态，Rejected：已拒绝，Handling：处理中，Succeed：已成功，Failed：已失败 | [Handling] |
| **ResourceIDs** | array<string> | 审批操作资源ID | [] |
| **NodeID** | string | 审批当前节点ID | [1] |
| **NodeName** | string | 审批当前节点名称 | [节点1] |
| **NodeOperator** | string | 审批当前节点审批人子账户邮箱 | [] |
| **NodeOperatorID** | integer | 审批当前节点审批人子账户ID | [200000247] |
| **Status** | integer | 审批工单状态,标记是否删除 | [1] |
| **CreateTime** | integer | 创建时间 | [1754968392169] |
| **UpdateTime** | integer | 更新时间 | [1754968392169] |
| **ApplicationNodes** | array<[*WorkflowApplicationNode*](#WorkflowApplicationNode)> | 审批工单所有节点 | [] |

#### WorkflowApplicationNode

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ApplicationMemberID** | integer | 申请者账户ID，子账号申请，则为该子账号ID | [200000247] |
| **ApplicationEmail** | string | 申请者账户邮箱 | [ucloud.sub@example.com] |
| **ApplicationID** | string | 审批工单ID | [apply-utkk9iljh7lpb1] |
| **ApplicationName** | string | 审批工单名称 | [new_application_name] |
| **ResourceType** | string | 审批资源类型，可以通过ListProductResources获取，可以通过ListProductResources获取，可以通过ListProductResources获取 | [DISK] |
| **APIName** | string | 审批资源对应API名称，例如创建磁盘操作，提交审批，APIName为CreateDisk | [CreateDisk] |
| **NodeID** | string | 审批节点ID | [1] |
| **NodeName** | string | 审批节点名称 | [节点1] |
| **State** | string | 审批节点状态，取值为Undo Committed Rejected Aborted Approved Handling Succeed Failed | [Handling] |
| **Remark** | string | 审批节点备注 | [这是审批备注] |
| **NodeOperator** | string | 节点审核人邮箱 | [example@ucloud.cn] |
| **NodeOperatorID** | integer | 节点审核人账号ID | [200000239] |
| **NodeOperatorCompanyID** | integer | 节点审核人租户ID | [200000233] |
| **CreateTime** | integer | 创建时间 | [1754968392] |
| **UpdateTime** | integer | 更新时间 | [1754968392] |
| **ObserverInfos** | array<[*ObserverInfo*](#ObserverInfo)> | 节点观察人信息 | [] |
| **AutoAudit** | bool | 节点是否自动审批 | [true] |

## 更新审批节点信息 - UpdateApplicationNode

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **ApplicationID** | string | 审批工单ID | [apply-utkk9iljh7lpb1] | [DescribeApplication] | **Yes** |
| **NodeID** | string | 审批节点ID | [1] | [DescribeApplication] | **Yes** |
| **NodeState** | string | 审核节点状态，已办，传Approved, Rejected。待办，传Handling | [Approved,Rejected Handling] | [] | **Yes** |
| **Remark** | string | 备注，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [new_application_remark] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 创建自定义流程 - CreateWorkflow

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [CreateUser DescribeUser DescribeMember] | **Yes** |
| **Name** | string | 自定义流程名称，长度为1-50个字符，名称只能包含中英文、数字、点（.），下划线（_）和中划线（-） | [new_custom_workflow_name] | [] | **Yes** |
| **Remark** | string | 备注，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [new_custom_workflow_remark] | [] | No |
| **ResourceTypes** | array<string> | 资源类型 | [DISK VM FS] | [ListProductResources] | **Yes** |
| **NodeInfos** | array<string> | 审核节点信息，是用|分割的字符串，格式为审核节点名称|审批人MemberID|自动审批true or false|观察人1 MemberID_1|观察人2 MemberID_2|...，其中审批人和观察人ID可以通过ListAdmin, DescribeUser和DescribeMember获取 | [name1|200000239|false|200000239|200000247|200000237 name2|200000237|true|200000242] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **WorkflowID** | string | 流程ID | [workflow-zh4o5juur1zki7] |


## 查询自定义流程 - DescribeWorkflow

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [CreateUser DescribeUser DescribeMember] | No |
| **Offset** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 | [0] | [] | No |
| **Limit** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。 | [20] | [] | No |
| **WorkflowIDs** | array<string> | 流程ID | [workflow-zh4o5juur1zki7] | [] | No |
| **Keyword** | string | 搜索关键字 | [] | [] | No |
| **Origin** | string | 来源，取值为Admin|Company, Admin代表来源自管理员, Compnay代表来源自租户 | [Admin] | [] | No |
| **ResourceType** | string | 资源类型，用于检索某一类资源类型的流程 | [DISK] | [ListProductResources] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 流程数 | [] |
| **Infos** | array<[*Workflow*](#Workflow)> | 流程信息 | [] |

### 数据模型

#### Workflow

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string | 自定义流程名称 | [new_custom_workflow_name] |
| **Remark** | string | 备注，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [new_custom_workflow_remark] |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] |
| **Email** | string | 租户邮箱 | [example@ucloud.cn] |
| **WorkflowID** | string | 流程ID | [workflow-zh4o5juur1zki7] |
| **ResourceTypes** | array<string> | 资源类型，可以通过ListProductResources获取，表明流程申请的资源类型 | [DISK VM FS] |
| **Nodes** | array<[*WorkflowNode*](#WorkflowNode)> | 审核节点 | [] |
| **Origin** | string | 来源 枚举值Admin|Company, Admin代表来源自管理员, Compnay代表来源自租户 | [Admin] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |

#### WorkflowNode

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string | 审核节点名称 | [name1] |
| **AuditMemberInfo** | object | 审核人信息 | [] |
| **ObserverMemberInfos** | array<[*WorkflowNodeMemberInfo*](#WorkflowNodeMemberInfo)> | 观察人信息 | [] |
| **AutoAudit** | bool | 是否自动审批 | [true] |

#### WorkflowNodeMemberInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **MemberID** | integer | 审核人或观察人ID | [200000239] |
| **Email** | string | 审核人或观察人邮箱 | [example@ucloud.cn] |

## 更新自定义流程 - UpdateWorkflow

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **WorkflowID** | string | 流程ID | [workflow-zh4o5juur1zki7] | [DescribeWorkflow] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [CreateUser DescribeUser DescribeMember] | **Yes** |
| **ResourceTypes** | array<string> | 资源类型，用于检索某一类资源类型的流程 | [DISK, FS, VM] | [ListProductResources] | **Yes** |
| **NodeInfos** | array<string> | 审核节点信息，是用|分割的字符串，格式为审核节点名称|审批人MemberID|自动审批true or false|观察人1 MemberID_1|观察人2 MemberID_2|...，其中审批人和观察人ID可以通过ListAdmin, DescribeUser和DescribeMember获取 | [name1|200000239|false|200000239|200000247|200000237 name2|200000237|true|200000242] | [] | **Yes** |
| **Name** | string | 自定义流程名称 | [new_custom_workflow_name] | [] | **Yes** |
| **Remark** | string | 备注，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [new_custom_workflow_remark] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 删除自定义流程 - DeleteWorkflow

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [CreateUser DescribeUser DescribeMember] | **Yes** |
| **WorkflowID** | string | 流程ID，表明需要删除的流程 | [workflow-zh4o5juur1zki7] | [DescribeWorkflow] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|

