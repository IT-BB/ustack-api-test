# DISK

## 绑定磁盘 - AttachDisk

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **ResourceType** | string | 资源类型，指定要挂载磁盘的类型 | [VM] | [] | No |
| **DiskID** | string | 磁盘ID，要挂载的磁盘的唯一标识符。 | [disk-6hvtncafjwmv4n] | [DescribeDisk] | No |
| **ResourceID** | string | 资源ID，可以通过 DescribeVMInstance 方法查看存在的虚拟机资源列表 | [vm-23om12646vbly3] | [DescribeVMInstance] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 克隆硬盘 - CloneDisk

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **ProjectID** | string | 项目组ID，用于资源分组管理 | [project-5kyt09jzwcqh6j] | [ListProjects] | No |
| **TagKeyValuePairs** | array<string> | 标签键值对，用于资源标记和分类管理。格式为 key:value 的字符串，传入​​ Base64 编码的字符串。 | [dGVzdA==:dGVzdC0x] | [] | **Yes** |
| **Name** | string | 磁盘名称，长度为1-50个字符，名称只能包含中英文、数字、点（.），下划线（_）和中划线（-） | [testDiskName] | [] | **Yes** |
| **Remark** | string | 磁盘描述，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [用于测试的磁盘] | [] | No |
| **ChargeType** | string | 计费类型，用于指定磁盘的计费模式。取值范围：1、Dynamic（按小时计费）2、Month（按月计费）3、Year（按年计费） | [Month] | [] | **Yes** |
| **Quantity** | integer | 计费数量，指定计费周期的数量。按月/年计费时表示创建 Quantity 个月/年的磁盘，按小时计费时强制为1。取值范围：正整数，最小值为1。 | [1] | [] | **Yes** |
| **DiskSpace** | integer | 磁盘大小，单位（GiB），指定克隆磁盘的存储容量。格式要求：正整数 | [10] | [] | No |
| **SetType** | string | 存储集群，存储集群的唯一标识符。 | [StorageSetCeph] | [DescribeStorageSet] | No |
| **SrcID** | string | 源磁盘ID，要克隆的源磁盘的唯一标识符。 | [disk-2aa5tms6u9duc5] | [DescribeDisk] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **DiskID** | string | 磁盘ID | [disk-2aa5tms6u9duc5] |


## 创建数据盘 - CreateDisk

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **ProjectID** | string | 项目组ID，用于资源分组管理 | [project-5kyt09jzwcqh6j] | [ListProjects] | No |
| **TagKeyValuePairs** | array<string> | 标签键值对，用于资源标记和分类管理。格式为 key:value 的字符串，传入​​ Base64 编码的字符串。 | [dGVzdA==:dGVzdC0x] | [] | **Yes** |
| **Name** | string | 磁盘名称，长度为1-50个字符，名称只能包含中英文、数字、点（.），下划线（_）和中划线（-） | [testDiskName] | [] | **Yes** |
| **Remark** | string | 磁盘描述，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [用于测试的磁盘] | [] | No |
| **ChargeType** | string | 计费类型，用于指定磁盘的计费模式。取值范围：1、Dynamic（按小时计费）2、Month（按月计费）3、Year（按年计费） | [Month] | [] | **Yes** |
| **Quantity** | integer | 计费数量，指定计费周期的数量。按月/年计费时表示创建 Quantity 个月/年的磁盘，按小时计费时强制为1。取值范围：正整数，最小值为1。 | [1] | [] | **Yes** |
| **ApplicationName** | string | 审批名称，用于审批流程的标题描述。当启用审批流程时使用。 | [磁盘资源申请] | [] | No |
| **ApplicationReason** | string | 审批理由，用于说明申请创建虚拟机的原因。当启用审批流程时使用。 | [用于新项目测试环境搭建] | [] | No |
| **SetType** | string | 集群类型，存储集群的唯一标识符。 | [StorageSetCeph] | [DescribeStorageSet] | **Yes** |
| **DiskSpace** | integer | 磁盘大小，单位（GiB），最小值为10，步长为1的整数， | [10] | [] | No |
| **Secret** | string | 加密密钥，密钥须包含有大写字母、小写字母、数字、特殊符号(除空格)中的两种或以上 | [Yfg41.?] | [] | No |
| **ShareAble** | bool | 是否为共享盘，指定磁盘是否支持同时挂载到多个虚拟机。true 表示共享盘，false 表示普通盘。默认为 false。 | [false] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **DiskID** | string | 磁盘ID | [disk-2aa5tms6u9duc5] |


## 从快照创建数据盘 - CreateDiskFromSnapshot

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **ProjectID** | string | 项目组ID，用于资源分组管理 | [project-5kyt09jzwcqh6j] | [ListProjects] | No |
| **TagKeyValuePairs** | array<string> | 标签键值对，用于资源标记和分类管理。格式为 key:value 的字符串，传入​​ Base64 编码的字符串。 | [dGVzdA==:dGVzdC0x] | [] | **Yes** |
| **Name** | string | 磁盘名称，长度为1-50个字符，名称只能包含中英文、数字、点（.），下划线（_）和中划线（-） | [testDiskName] | [] | **Yes** |
| **Remark** | string | 磁盘描述，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [用于测试的磁盘] | [] | No |
| **ChargeType** | string | 计费类型，用于指定磁盘的计费模式。取值范围：1、Dynamic（按小时计费）2、Month（按月计费）3、Year（按年计费） | [Month] | [] | **Yes** |
| **Quantity** | integer | 计费数量，按月/年计费，就是创建Quantity月/年的磁盘，按小时计费强制为1 | [1] | [] | **Yes** |
| **SrcSnapshotID** | string | 源快照ID | [disksnapshot-v2lcgob2bide53] | [DescribeSnapshot] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **DiskID** | string | 所创建的磁盘ID | [disk-2aa5tms6u9duc5] |


## 删除磁盘 - DeleteDisk

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **DiskID** | string | 磁盘ID，要删除的磁盘的唯一标识符。 | [disk-2aa5tms6u9duc5] | [DescribeDisk] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 查询磁盘信息 - DescribeDisk

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **Offset** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 默认值：0 | [0] | [] | No |
| **Limit** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。默认值：10 | [20] | [] | No |
| **Keyword** | string | 搜索关键词 | [example_key] | [] | No |
| **ProjectIDs** | array<string> | 项目组ID列表，用于筛选指定项目组下的磁盘资源，多个元素之间用,分隔。 | [project-5kyt09jzwcqh6j] | [] | No |
| **DiskType** | string | 硬盘类型，已废弃的参数，请使用 DiskTypes 替代。当 DiskTypes 有值时，将忽略本参数。 | [] | [] | No |
| **DiskIDs** | array<string> | 硬盘ID列表，用于查询指定磁盘的详细信息，多个元素之间用,分隔。 | [disk-2aa5tms6u9duc5] | [DescribeDiskRequest] | No |
| **AttachResourceID** | string | 挂载的资源ID，用于筛选已挂载到指定资源的磁盘。 | [vm-f123thdfgessa] | [DescribeVMInstance] | No |
| **ShareAbleFilter** | string | 筛选共享盘，用于过滤磁盘类型。true 仅返回共享盘，false 仅返回普通盘，空值 表示返回所有类型。 | [true] | [] | No |
| **DiskTypes** | array<string> | 硬盘类型列表，替换已废弃的 DiskType 参数，支持更多的硬盘类型筛选。可指定多种磁盘类型进行过滤。取值范围：1、Data 2、Boot | [Data] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 查询到的总记录数 | [] |
| **Infos** | array<[*DiskInfo*](#DiskInfo)> | 查询到的磁盘信息 | [] |

### 数据模型

#### AttachShareBlock

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **AttachResourceID** | string | 已挂载资源ID | [vm-rfg34tffdggr] |
| **AttachResourceName** | string | 已挂载资源名称 | [test-vm] |
| **AttachResourceType** | string | 已挂载资源类型 | [VM] |
| **AttachStatus** | string | 挂载状态 | [Bound] |

#### DiskInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string | 磁盘名称 | [test] |
| **Remark** | string | 备注，用于进行说明和注释，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [这是一个备注] |
| **Region** | string | 磁盘所属的地域 | [region-1] |
| **RegionAlias** | string | 地域别名 | [test] |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [2034365] |
| **CompanyName** | string | 租户名称 | [ucloud] |
| **Email** | string | 租户邮箱，租户注册时提供的用于接收通知和联系的邮箱 | [example@ucloud.cn] |
| **ProjectID** | string | 项目组ID，用于资源分组管理 | [project-5kyt09jzwcqh6j] |
| **ProjectName** | string | 项目组名称 | [project01] |
| **Reason** | string | 创建失败原因 | [存储集群空间不足] |
| **CreateTime** | integer | 创建时间，Unix 时间戳（秒级） | [1754644992] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> | 标签 | [cluster-resouce.system:offline-node] |
| **ExpireTime** | integer | 过期时间，Unix 时间戳（秒级） | [1754644992] |
| **ChargeType** | string | 计费类型，用于指定计费模式。取值范围：1、Dynamic（按小时计费）2、Month（按月计费）3、Year（按年计费） | [Month] |
| **DiskID** | string | 磁盘ID | [disk-213rgt3tyhyh3] |
| **DiskType** | string | 磁盘类型 | [Data] |
| **Size** | integer | 磁盘大小，单位：GiB | [10] |
| **Drive** | string | 设备名 | [vda] |
| **DiskStatus** | string | 磁盘状态，取值范围：（Unknown:未知状态, Creating:创建中, Detached:可用就绪/已卸载, Attaching:装载中, Attached:已装载, Detaching:卸载中, Deleting:删除中, Deleted:已删除, Rollbacking:回滚中, Cloning:克隆中, Shareabling:共享中, Shareabled:已共享, Blockcopying:复制中, Blockcopyed:复制完成） | [Attached] |
| **DisableQoS** | bool | 是否禁用QoS | [true] |
| **IOPS** | integer | 每秒输入输出操作次数(IOPS)，表示磁盘性能指标 | [5000] |
| **Bandwidth** | integer | 磁盘带宽(MB/s)，表示磁盘数据传输速率 | [200] |
| **SetType** | string | 存储集群类型 | [StorageSetNPVK] |
| **SetAlias** | string | 存储集群别名 | [StorageSetNPVK] |
| **SetArch** | string | 集群架构 | [HDD SSD] |
| **SetProvider** | string | 存储集群制备器 | [rbd ustorbs cjfs] |
| **AttachResourceID** | string | 已挂载资源ID | [vm-j40jm6vb8lo67c] |
| **AttachResourceName** | string | 已挂载资源名称 | [vm-test] |
| **AttachResourceType** | string | 已挂载资源类型 | [VM] |
| **Encrypted** | bool | 是否加密 | [true] |
| **ShareAble** | bool | 是否共享 | [true] |
| **AttachShareBlockInfos** | array<[*AttachShareBlock*](#AttachShareBlock)> | 共享信息 | [] |
| **Percent** | float | 迁移进度百分比(0.00~100.00) | [99.99] |
| **SetID** | string | 存储集群ID | [storage-set-rwio] |
| **SnapshotCount** | integer | 快照数量 | [2] |
| **AttachResourceStatus** | string | 已挂载资源状态 | [Poweroffing] |
| **AttachResourceSnapCount** | integer | 已挂载资源整机快照数量 | [1] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

## 解绑磁盘 - DetachDisk

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **ResourceType** | string | 资源类型 | [VM] | [] | No |
| **DiskID** | string | 磁盘ID，要分离的磁盘的唯一标识符。 | [disk-2aa5tms6u9duc5] | [DescribeDisk] | **Yes** |
| **ResourceID** | string | 资源ID | [vm-23om12646vbly3] | [DescribeVMInstance] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取数据盘的价格 - GetDiskPrice

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **ChargeType** | string | 计费类型，用于指定磁盘的计费模式。取值范围：1、Dynamic（按小时计费）2、Month（按月计费）3、Year（按年计费） | [Month] | [] | **Yes** |
| **Quantity** | integer | 计费数量，指定计费周期的数量。按月/年计费时表示创建 Quantity 个月/年的磁盘，按小时计费时强制为1。取值范围：正整数，最小值为1。 | [1] | [] | **Yes** |
| **Count** | integer | 磁盘数量，指定要创建的磁盘数量。用于批量创建磁盘时的价格计算。取值范围：正整数，最小值为1。 | [1] | [] | **Yes** |
| **SetType** | string | 集群类型，存储集群的唯一标识符。 | [DescribeStorageSet] | [] | **Yes** |
| **DiskSpace** | integer | 磁盘大小，单位：GiB，指定磁盘的存储容量。最小值为10的正整数。 | [10] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*DiskPriceInfo*](#DiskPriceInfo)> | 磁盘价格信息 | [] |

### 数据模型

#### DiskPriceInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Price** | double | 价格，磁盘的计费价格。单位根据计费类型而定，如按小时计费则为每小时价格。 | [] |
| **ChargeType** | string | 计费类型，用于指定磁盘的计费模式。取值范围：1、Dynamic（按小时计费）2、Month（按月计费）3、Year（按年计费） | [Month] |
| **PurchaseValue** | string | 购买价值 | [] |

## 获取创建硬盘价格 - GetCreateDiskPrice

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **ChargeType** | string | 计费类型，用于指定磁盘的计费模式。取值范围：1、Dynamic（按小时计费）2、Month（按月计费）3、Year（按年计费） | [Month] | [] | **Yes** |
| **Quantity** | integer | 计费数量，指定计费周期的数量。按月/年计费时表示创建 Quantity 个月/年的磁盘，按小时计费时强制为1。取值范围：正整数，最小值为1。 | [1] | [] | **Yes** |
| **Count** | integer | 磁盘数量，指定要创建的磁盘数量。用于批量创建磁盘时的价格计算。取值范围：正整数，最小值为1。 | [1] | [] | No |
| **SetType** | string | 存储集群类型，存储集群的唯一标识符。 | [StorageSetCeph] | [DescribeStorageSet] | **Yes** |
| **DiskSpace** | integer | 磁盘大小，单位：GiB，指定磁盘的存储容量。最小值为10，步长为1的整数。用于计算不同容量磁盘的价格。 | [10] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*DiskPriceInfo*](#DiskPriceInfo)> | 磁盘价格信息 | [] |

### 数据模型

#### DiskPriceInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Price** | double | 价格，磁盘的计费价格。单位根据计费类型而定，如按小时计费则为每小时价格。 | [] |
| **ChargeType** | string | 计费类型，用于指定磁盘的计费模式。取值范围：1、Dynamic（按小时计费）2、Month（按月计费）3、Year（按年计费） | [Month] |
| **PurchaseValue** | string | 购买价值 | [] |

## 获取升级虚拟硬盘的差价 - GetUpgradeDiskPrice

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **DiskID** | string | 磁盘ID，要升级的磁盘的唯一标识符。 | [disk-2aa5tms6u9duc5] | [DescribeDisk] | **Yes** |
| **DiskSpace** | integer | 磁盘大小，单位：GiB，指定升级后的磁盘存储容量。必须大于当前磁盘大小，最小值为10的正整数。 | [10] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Price** | double | 需要补的差价 | [] |
| **Infos** | array<[*DiskPriceInfo*](#DiskPriceInfo)> | 磁盘价格信息 | [] |

### 数据模型

#### DiskPriceInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Price** | double | 价格，磁盘的计费价格。单位根据计费类型而定，如按小时计费则为每小时价格。 | [] |
| **ChargeType** | string | 计费类型，用于指定磁盘的计费模式。取值范围：1、Dynamic（按小时计费）2、Month（按月计费）3、Year（按年计费） | [Month] |
| **PurchaseValue** | string | 购买价值 | [] |

## 设置硬盘QoS - UpdateDiskQoS

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **DiskID** | string | 磁盘ID，要更新QoS配置的磁盘的唯一标识符。 | [disk-2aa5tms6u9duc5] | [DescribeDisk] | **Yes** |
| **DiskIOPS** | integer | 硬盘IOPS限制，指定磁盘的每秒输入/输出操作数限制。取值范围：0-50000，0表示不限制。用于控制磁盘性能。 | [] | [] | No |
| **DiskBandwidth** | integer | 硬盘带宽限制，单位: MBps，指定磁盘的数据传输带宽限制。取值范围：0-1000，0表示不限制。用于控制磁盘吞吐量。 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 升级虚拟硬盘 - UpgradeDisk

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **ApplicationName** | string | 审批名称，用于审批流程的标题描述。当启用审批流程时使用。 | [磁盘资源申请] | [] | No |
| **ApplicationReason** | string | 审批理由，用于说明申请创建虚拟机的原因。当启用审批流程时使用。 | [用于新项目测试环境搭建] | [] | No |
| **DiskID** | string | 磁盘ID，要升级的磁盘的唯一标识符。 | [disk-2aa5tms6u9duc5] | [DescribeDisk] | **Yes** |
| **DiskSpace** | integer | 磁盘大小，单位：GiB，指定升级后的磁盘存储容量。取值范围：必须大于当前磁盘大小，正整数。 | [10] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 绑定iso - AttachISO

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **ResourceID** | string | 资源ID，要挂载ISO镜像的目标资源的唯一标识符。 | [vm-2aa5tms6u9duc5] | [DescribeVMInstance] | **Yes** |
| **ImageID** | string | ISO镜像ID，要挂载的ISO镜像文件的唯一标识符。 | [image-windows-virtio-iso] | [DescribeSnapshot] | **Yes** |
| **SetType** | string | 存储集群类型，指定ISO镜像存储的集群类型。如果不指定，默认使用系统盘所在的存储集群。 | [StorageSetCeph] | [DescribeStorageSet] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 解绑iso - DetachISO

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **ResourceID** | string | 资源ID，要卸载ISO镜像的目标资源的唯一标识符。 | [vm-2aa5tms6u9duc5] | [DescribeVMInstance] | **Yes** |
| **ImageID** | string | ISO镜像ID，要卸载的ISO镜像文件的唯一标识符。 | [image-windows-virtio-iso] | [DescribeSnapshot] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 查询iso信息 - DescribeVMISO

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **AttachResourceID** | string | 挂载的资源ID，要查询ISO镜像挂载信息的目标资源的唯一标识符。 | [vm-2aa5tms6u9duc5] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 查询到的总记录数 | [] |
| **Infos** | array<[*ISOInfo*](#ISOInfo)> | iso磁盘信息 | [] |

### 数据模型

#### ISOInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Region** | string | ISO所属地域 | [cn] |
| **Type** | string | ISO磁盘类型 | [cdrom] |
| **Size** | integer | 磁盘大小，单位：GiB | [1] |
| **Status** | string | ISO磁盘状态 | [Bound] |
| **SetType** | string | 集群类型 | [StorageSetNPVK] |
| **SetArch** | string | 集群架构 | [SSD] |
| **ImageID** | string | iso镜像ID | [image-windows-virtio-iso] |
| **Boot** | bool | 是否从此盘引导 | [true] |
| **ImageName** | string | iso镜像名字 | [image-windows-virtio-iso] |
