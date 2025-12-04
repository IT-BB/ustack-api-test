# MYSQL

## 创建MySQL - CreateMySQL

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **ProjectID** | string | 项目组ID | [] | [] | No |
| **TagKeyValuePairs** | array<string> | 标签键值对 | [] | [] | **Yes** |
| **Name** | string | 名字 | [] | [] | **Yes** |
| **Remark** | string | 描述 | [] | [] | No |
| **ChargeType** | string | 计费类型 | [] | [] | **Yes** |
| **Quantity** | integer | 计费时长 | [] | [] | **Yes** |
| **VMType** | string | 虚拟机所在宿主机的集群类型 | [] | [] | **Yes** |
| **Memory** | integer | 内存 | [] | [] | **Yes** |
| **DiskSetType** | string | 存储集群 | [] | [] | **Yes** |
| **DiskSpace** | integer | 存储容量 | [] | [] | **Yes** |
| **VPCID** | string | vpcID | [] | [] | **Yes** |
| **SubnetID** | string | 子网ID | [] | [] | **Yes** |
| **EIPID** | string | eipID | [] | [] | No |
| **WANSGID** | string | 外网安全组 | [] | [] | No |
| **Password** | string | root 密码 | [] | [] | **Yes** |
| **ConfigFileID** | string | 配置文件ID | [] | [] | **Yes** |
| **BackupID** | string | 备份文件ID | [] | [] | No |
| **Version** | string | MySQL 版本，必须为 'MySQL 5.7' 或者'MySQL 8.0' | [] | [] | **Yes** |
| **HighAvailability** | string | MySQL 数据库类型 | [] | [] | No |
| **RestoreTime** | integer | 恢复时间 | [] | [] | No |
| **TerminationPolicy** | integer | 是否开启删除保护, 0:开启; 1:关闭 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **MySQLID** | string | MySQLID | [] |


## 查询MySQL信息 - DescribeMySQL

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **ProjectIDs** | array<string> | 项目组ID | [] | [] | No |
| **Keyword** | string | 搜索关键词 | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **MySQLIDs** | array<string> | mysql id | [] | [] | No |
| **Status** | array<string> | MySQL状态 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数 | [] |
| **Infos** | array<[*MySQLInfo*](#MySQLInfo)> | MySQL 信息 | [] |

### 数据模型

#### MySQLInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **MySQLID** | string | MySQLID | [] |
| **Name** | string | 名字 | [] |
| **Remark** | string | 描述 | [] |
| **Region** | string | 地域 | [] |
| **RegionAlias** | string | 地域别名 | [] |
| **CompanyID** | integer | 租户ID | [] |
| **CompanyName** | string | 租户名 | [] |
| **Email** | string | 租户邮箱 | [] |
| **ProjectID** | string | 项目组ID | [] |
| **ProjectName** | string | 项目组名 | [] |
| **Reason** | string | 失败原因 | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> | 标签 | [] |
| **ExpireTime** | integer | 过期时间 | [] |
| **ChargeType** | string | 计费类型 | [] |
| **Version** | string | MySQL 版本 | [] |
| **VMTypeAlias** | string | 计算集群名称 | [] |
| **VPCName** | string | vpc名称 | [] |
| **SubnetName** | string | 子网名称 | [] |
| **VPCID** | string | vpcID | [] |
| **SubnetID** | string | 子网ID | [] |
| **Endpoints** | array<[*MysqlEndpoint*](#MysqlEndpoint)> | MySQL网络信息 | [] |
| **Status** | string | 状态 | [] |
| **Memory** | integer | 内存 | [] |
| **DiskID** | string | 磁盘ID | [] |
| **DiskSpace** | integer | 磁盘容量 | [] |
| **StorageSetType** | string | 存储集群 | [] |
| **HighAvailability** | string | 高可用 | [] |
| **SlaveInfos** | array<[*MySQLSlaveInfo*](#MySQLSlaveInfo)> | 从库信息 | [] |
| **Role** | string | 角色 | [] |
| **IOPS** | integer | IOPS | [] |
| **Bandwidth** | integer | 带宽 | [] |
| **Arch** | string | 架构 | [] |
| **StorageSetArch** | string | 存储集群架构 | [] |
| **TerminationPolicy** | integer | 是否开启删除保护 0:开启; 1:关闭 | [] |
| **AuditStatus** | string | SQL审计状态，0:关闭 1:开启 | [] |

#### MySQLSlaveInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **MySQLID** | string | MySQLID | [] |
| **Region** | string | 地域 | [] |
| **Name** | string | 名字 | [] |
| **Remark** | string | 描述 | [] |
| **CompanyID** | integer | 租户ID | [] |
| **Email** | string | 租户Email | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **UpdateTime** | integer | 更新时间 | [] |
| **ExpireTime** | integer | 过期时间 | [] |
| **Version** | string | MySQL 版本 | [] |
| **ChargeType** | string | 计费类型 | [] |
| **ProjectID** | string | 项目组ID | [] |
| **ProjectName** | string | 项目组名 | [] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> | 标签 | [] |
| **VMTypeAlias** | string | 计算集群名称 | [] |
| **VPCName** | string | VPC 名称 | [] |
| **SubnetName** | string | 子网名称 | [] |
| **VPCID** | string | vpcID | [] |
| **SubnetID** | string | 子网ID | [] |
| **Endpoints** | array<[*MysqlEndpoint*](#MysqlEndpoint)> | MySQL网络信息 | [] |
| **Status** | string | 状态 | [] |
| **Reason** | string | 失败原因 | [] |
| **Memory** | integer | 内存 | [] |
| **DiskID** | string | 磁盘ID | [] |
| **DiskSpace** | integer | 磁盘容量 | [] |
| **StorageSetType** | string | 存储集群 | [] |
| **Role** | string | 角色 | [] |
| **IOPS** | integer | IOPS | [] |
| **Bandwidth** | integer | 带宽 | [] |
| **Arch** | string | 架构 | [] |
| **TerminationPolicy** | integer | 是否开启删除保护 0:开启; 1:关闭 | [] |

#### MysqlEndpoint

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **IP** | string | IP | [] |
| **IPVersion** | string | IP版本 | [] |
| **NicType** | string | 网卡类型 | [] |
| **Endpoint** | string | 端点 | [] |
| **SGID** | string | 安全组ID | [] |
| **SGName** | string | 安全组名称 | [] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

## 查询MySQL价格 - GetMySQLPrice

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **ChargeType** | string | 计费类型 | [] | [] | **Yes** |
| **VMType** | string | 计算集群 | [] | [] | **Yes** |
| **Memory** | integer | 内存 | [] | [] | **Yes** |
| **DiskSetType** | string | 存储集群 | [] | [] | **Yes** |
| **DiskSpace** | integer | 硬盘大小 | [] | [] | **Yes** |
| **Quantity** | integer | 计费时长 | [] | [] | **Yes** |
| **MySQLID** | string | MySQL实例ID | [] | [] | No |
| **HighAvailability** | string | 数据库类型 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*MysqlPriceInfo*](#MysqlPriceInfo)> | 价格信息 | [] |

### 数据模型

#### MysqlPriceInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ChargeType** | string | 计费类型 | [] |
| **OrderType** | string | 订单类型 | [] |
| **Price** | double | 价格 | [] |

## 升级MySQL - UpgradeMySQL

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **MySQLID** | string | MySQLID | [] | [] | **Yes** |
| **Memory** | integer | 内存 | [] | [] | **Yes** |
| **DiskSpace** | integer | 硬盘容量 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 删除MySQL - DeleteMySQL

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **MySQLID** | string | MySQLID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 修改MySQL密码 - ResetMySQLPassword

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **MySQLID** | string | MySQLID | [] | [] | **Yes** |
| **Password** | string | 密码 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 应用MySQL参数模板 - ApplyMySQLParamTpl

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **TplID** | string | 模版ID | [] | [] | **Yes** |
| **MySQLID** | string | MySQLID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 创建MySQL参数模板 - CreateMySQLParamTpl

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **Name** | string | 名称 | [] | [] | **Yes** |
| **Remark** | string | 描述 | [] | [] | No |
| **Version** | string | 版本 | [] | [] | **Yes** |
| **SrcTplID** | string | 源模版ID | [] | [] | No |
| **Params** | string | 自定义配置项 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TplID** | string | 模版ID | [] |


## 删除MySQL参数模板 - DeleteMySQLParamTpl

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **TplID** | string | 模版ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取MySQL配置参数 - DescribeMySQLConfigParam

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **MySQLID** | string | MySQLID | [] | [] | **Yes** |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数 | [] |
| **Infos** | array<[*MySQLConfigParamInfo*](#MySQLConfigParamInfo)> |  | [] |

### 数据模型

#### MySQLConfigParamInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Description** | string |  参数描述 | [] |
| **Key** | string | 参数名称 | [] |
| **Value** | string | 参数值 | [] |
| **ValueRange** | string | 参数值范围 | [] |
| **ValueType** | string | 参数类型 | [] |
| **NeedRestart** | integer | 是否需要重启，0:不需要；1:需要 | [] |

## 查询MySQL参数模板详细信息 - DescribeMySQLParamTpl

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **TplID** | string | 模版ID | [] | [] | **Yes** |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数 | [] |
| **Infos** | array<[*MySQLConfigParamInfo*](#MySQLConfigParamInfo)> | 模版信息 | [] |

### 数据模型

#### MySQLConfigParamInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Description** | string |  参数描述 | [] |
| **Key** | string | 参数名称 | [] |
| **Value** | string | 参数值 | [] |
| **ValueRange** | string | 参数值范围 | [] |
| **ValueType** | string | 参数类型 | [] |
| **NeedRestart** | integer | 是否需要重启，0:不需要；1:需要 | [] |

## 查询MySQL参数模板列表 - DescribeMySQLParamTpls

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **TplIDs** | array<string> | 模版ID | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数 | [] |
| **Infos** | array<[*MySQLParamTplInfo*](#MySQLParamTplInfo)> | 模版信息 | [] |

### 数据模型

#### MySQLParamTplInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] |
| **CreateTime** | integer | 创建时间 | [] |
| **Email** | string | Email | [] |
| **ID** | string | 模版ID | [] |
| **Name** | string | 名称 | [] |
| **Region** | string | 地域 | [] |
| **Remark** | string | 描述 | [] |
| **Type** | string | 类型 | [] |
| **UpdateTime** | integer | 更新时间 | [] |
| **Version** | string | 版本 | [] |

## 更新MySQL配置参数 - UpdateMySQLConfigParam

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **MySQLID** | string | MySQLID | [] | [] | **Yes** |
| **Params** | string | 配置参数 | [] | [] | **Yes** |
| **AutoRestart** | string | 是否自动重启,0-是；1-否 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*MySQLConfigOpLogInfo*](#MySQLConfigOpLogInfo)> |  | [] |

### 数据模型

#### MySQLConfigOpLogInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ParameterName** | string | 参数名称 | [] |
| **OldValue** | string | 修改前的值 | [] |
| **NewValue** | string | 修改后的值 | [] |
| **UpdateTime** | integer | 更新时间 | [] |

## 更新MySQL参数模板 - UpdateMySQLParamTpl

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **TplID** | string | 模版ID | [] | [] | **Yes** |
| **Params** | string | 配置参数 | [] | [] | No |
| **Name** | string | 名称 | [] | [] | No |
| **Remark** | string | 描述 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 创建MySQL从库 - CreateMySQLSlave

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **TagKeyValuePairs** | array<string> | 标签键值对 | [] | [] | **Yes** |
| **MasterID** | string | 主库ID | [] | [] | **Yes** |
| **Name** | string | 名字 | [] | [] | **Yes** |
| **Remark** | string | 描述 | [] | [] | No |
| **ChargeType** | string | 计费类型 | [] | [] | **Yes** |
| **Quantity** | integer | 计费时长 | [] | [] | **Yes** |
| **Memory** | integer | 内存 | [] | [] | **Yes** |
| **DiskSpace** | integer | 磁盘大小 | [] | [] | **Yes** |
| **VMType** | string | 虚拟机所在宿主机的集群类型 | [] | [] | **Yes** |
| **DiskSetType** | string | 存储集群 | [] | [] | **Yes** |
| **EIPID** | string | 外网 IP ID | [] | [] | No |
| **WANSGID** | string | 安全组 ID | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **MySQLID** | string | MySQLID | [] |


## 升级为高可用版本 - UpgradeMySQLToHA

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **MySQLID** | string | MySQL ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取 PMA URL - DescribePMAURL

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **MySQLID** | string | MySQLID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **PMAURL** | string |  | [] |


## 重启 MySQL 实例 - RestartMySQLInstance

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **MySQLID** | string | MySQLID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 查询 MySQL 慢日志记录 - DescribeMySQLSlowLogRecords

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **BeginTime** | integer | 开始时间 | [] | [] | **Yes** |
| **EndTime** | integer | 结束时间 | [] | [] | **Yes** |
| **MySQLID** | string | MySQL 实例 ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数 | [] |
| **Infos** | array<[*SlowLogInfo*](#SlowLogInfo)> | 慢日志信息 | [] |

### 数据模型

#### SlowLogInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **UserHost** | string | 客户端名称及地址 | [] |
| **QueryTime** | float | 执行时长 | [] |
| **LockTime** | float | 锁定时长 | [] |
| **RowsSent** | integer | 返回行数 | [] |
| **RowsExamined** | integer | 解析行数 | [] |
| **DB** | string | 数据库名称 | [] |
| **LastInsertID** | integer | 最新插入ID | [] |
| **InsertID** | integer | 插入ID | [] |
| **ServerID** | integer | mysql 服务 ID | [] |
| **SQLText** | string | SQL 语句 | [] |
| **ThreadID** | integer | 线程 ID | [] |
| **StartTime** | integer | 开始时间 | [] |

## 查询 MySQL 错误日志 - DescribeMySQLErrorLogs

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **BeginTime** | integer | 开始时间 | [] | [] | **Yes** |
| **EndTime** | integer | 结束时间 | [] | [] | **Yes** |
| **MySQLID** | string | MySQL 实例 ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数 | [] |
| **Infos** | array<[*ErrorLogInfo*](#ErrorLogInfo)> | 慢日志信息 | [] |

### 数据模型

#### ErrorLogInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **CreateTime** | integer | 发生时间 | [] |
| **LogContent** | string | 日志内容 | [] |

## 降级MySQL - DowngradeMySQL

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **MySQLID** | string | MySQLID | [] | [] | **Yes** |
| **Memory** | integer | 内存 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|

