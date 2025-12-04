# SMC

## 创建SMC任务 - CreateSMC

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **ProjectID** | string | 项目ID | [] | [] | No |
| **TagKeyValuePairs** | array<string> | 标签键值对 | [] | [] | **Yes** |
| **Name** | string | 名称 | [] | [] | **Yes** |
| **Remark** | string | 备注 | [] | [] | No |
| **AgentVersion** | string | VPC网段 | [] | [] | **Yes** |
| **OS** | string |  | [] | [] | No |
| **Arch** | string |  | [] | [] | No |
| **Distribution** | string |  | [] | [] | No |
| **DistributionVersion** | string |  | [] | [] | No |
| **Kernel** | string |  | [] | [] | No |
| **Cores** | integer |  | [] | [] | No |
| **Memory** | integer |  | [] | [] | No |
| **Hostname** | string |  | [] | [] | No |
| **IP** | string |  | [] | [] | No |
| **BootloaderType** | string | 引导方式 | [] | [] | No |
| **Disks** | array<Disk> |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **SMCID** | string | SMCID | [] |


## 设置SMC任务 - SetupSMC

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **SMCID** | string | SMCID | [] | [] | **Yes** |
| **VMID** | string | 虚拟机ID | [] | [] | **Yes** |
| **Host** | string | 传输IP | [] | [] | **Yes** |
| **Bandwidth** | integer | 限速 | [] | [] | No |
| **User** | string | 用户名 | [] | [] | No |
| **Password** | string | 密码 | [] | [] | No |
| **DiskMappings** | array<DiskMapping> | 磁盘映射 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 删除SMC任务 - DeleteSMC

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **SMCID** | string | SMCID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 开始迁移 - StartSMC

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **SMCID** | string | SMCID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 停止迁移 - StopSMC

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **SMCID** | string | SMCID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 完成迁移 - CompleteSMC

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **SMCID** | string | SMCID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## smc心跳 - SMCHeartbeat

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer |  | [] | [] | No |
| **SMCID** | string | 迁移ID | [] | [] | **Yes** |
| **DoneSubtask** | integer | 已完成任务数 | [] | [] | No |
| **TotalSubtask** | integer | 总任务数 | [] | [] | No |
| **Subtasks** | array<SMCSubtask> | 任务 | [] | [] | No |
| **State** | string | 状态 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **SMCID** | string | 迁移ID | [] |
| **State** | string | 状态 | [] |
| **CompressLevel** | bool | 是否压缩 | [] |
| **Checksum** | bool | 是否校验和 | [] |
| **Bandwidth** | integer | 限速 | [] |
| **WholeFile** | bool | 全量文件 | [] |
| **VMID** | string | 传输地址 | [] |
| **Host** | string | 传输地址 | [] |
| **Port** | integer | 传输端口 | [] |
| **User** | string | 传输用户 | [] |
| **Password** | string |  | [] |
| **DiskMappings** | array<[*DiskMapping*](#DiskMapping)> | 磁盘映射 | [] |

### 数据模型

#### DiskMapping

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **FsType** | string |  | [] |
| **Size** | double |  | [] |
| **Avail** | double |  | [] |
| **Device** | string |  | [] |
| **MountPoint** | string |  | [] |
| **IsRoot** | bool |  | [] |
| **MappingDiskID** | string |  | [] |

## 获取SMC信息 - DescribeSMC

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **SMCIDs** | array<string> |  | [] | [] | No |
| **ProjectIDs** | array<string> |  | [] | [] | No |
| **Keyword** | string |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **States** | array<string> | 状态 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*SMCInfo*](#SMCInfo)> |  | [] |
| **TotalCount** | integer |  | [] |

### 数据模型

#### DiskMapping

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **FsType** | string |  | [] |
| **Size** | double |  | [] |
| **Avail** | double |  | [] |
| **Device** | string |  | [] |
| **MountPoint** | string |  | [] |
| **IsRoot** | bool |  | [] |
| **MappingDiskID** | string |  | [] |

#### SMCInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **SMCID** | string | 迁移ID | [] |
| **Name** | string | 名称 | [] |
| **Remark** | string | 备注 | [] |
| **Region** | string | 地域 | [] |
| **RegionAlias** | string | 地域别名 | [] |
| **CompanyID** | integer | 租户ID | [] |
| **CompanyName** | string | 租户名称 | [] |
| **Email** | string | 租户邮箱 | [] |
| **ProjectID** | string | 项目组ID | [] |
| **ProjectName** | string | 项目组名称 | [] |
| **Reason** | string |  | [] |
| **CreateTime** | integer |  | [] |
| **UpdateTime** | integer |  | [] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> |  | [] |
| **OS** | string |  | [] |
| **Arch** | string |  | [] |
| **Distribution** | string |  | [] |
| **DistributionVersion** | string |  | [] |
| **Kernel** | string |  | [] |
| **Cores** | integer |  | [] |
| **Memory** | integer |  | [] |
| **Hostname** | string |  | [] |
| **IP** | string |  | [] |
| **BootloaderType** | string |  | [] |
| **DoneSubtask** | integer | 已完成任务数 | [] |
| **TotalSubtask** | integer | 总任务数 | [] |
| **Subtasks** | array<[*SMCSubtask*](#SMCSubtask)> | 任务 | [] |
| **State** | string | 状态 | [] |
| **VMID** | string | 虚拟机ID | [] |
| **Host** | string | 传输IP | [] |
| **Bandwidth** | integer | 限速 | [] |
| **Locked** | bool | 锁定 | [] |
| **LastSyncedTime** | integer | 最后同步时间 | [] |
| **User** | string | 用户名 | [] |
| **DiskMappings** | array<[*DiskMapping*](#DiskMapping)> | 磁盘映射 | [] |

#### SMCSubtask

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string | 名称 | [] |
| **Progress** | double | 进度 | [] |
| **Speed** | string | 速率 | [] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |
