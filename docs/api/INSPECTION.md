# INSPECTION

## 创建一键巡检报告 - CreateOnSiteInspection

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Name** | string |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **InspectionID** | string |  | [] |


## 删除一键巡检报告 - DeleteOnSiteInspection

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **InspectionID** | string |  | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取巡检报告列表 - ListOnSiteInspections

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **Keyword** | string |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer |  | [] |
| **Infos** | array<[*InspectionInspectionInfo*](#InspectionInspectionInfo)> |  | [] |

### 数据模型

#### InspectionInspectionInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **InspectionID** | string |  | [] |
| **CompanyID** | integer |  | [] |
| **Name** | string |  | [] |
| **Status** | string |  | [] |
| **Spend** | integer |  | [] |
| **CreateTime** | integer |  | [] |
| **UpdateTime** | integer |  | [] |

## 获取巡检报告列表 - GetOnSiteInspection

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **InspectionID** | string |  | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Info** | object |  | [] |

### 数据模型

#### InspectionReport

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **InspectionID** | string |  | [] |
| **Status** | string |  | [] |
| **Name** | string |  | [] |
| **Score** | integer |  | [] |
| **StartTime** | integer |  | [] |
| **EndTime** | integer |  | [] |
| **Spend** | integer |  | [] |
| **TotalCount** | integer |  | [] |
| **NormalCount** | integer |  | [] |
| **AbnormalCount** | integer |  | [] |
| **RegionReports** | array<[*MenuOne*](#MenuOne)> |  | [] |

#### ItemReport

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ItemNo** | integer |  | [] |
| **ItemName** | string |  | [] |
| **ItemDesc** | string |  | [] |
| **Suggest** | string |  | [] |
| **Results** | array<[*Result*](#Result)> |  | [] |

#### MenuOne

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Title** | string |  | [] |
| **Key** | string |  | [] |
| **Children** | array<[*MenuTwo*](#MenuTwo)> |  | [] |

#### MenuTwo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Title** | string |  | [] |
| **Key** | string |  | [] |
| **Children** | array<[*ItemReport*](#ItemReport)> |  | [] |

#### Result

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string |  | [] |
| **Result** | string |  | [] |
| **Present** | string |  | [] |
| **Score** | integer |  | [] |
| **Top** | array<[*TopMap*](#TopMap)> |  | [] |

#### TopMap

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string |  | [] |
| **Pid** | string |  | [] |
| **Use** | string |  | [] |
| **Status** | string |  | [] |
| **Size** | string |  | [] |

## 创建资源使用情况报告 - CreateResourceUsage

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Name** | string |  | [] | [] | **Yes** |
| **Region** | string |  | [] | [] | **Yes** |
| **CompanyIDs** | array<int32> |  | [] | [] | No |
| **ProjectIDs** | array<string> |  | [] | [] | No |
| **ResourceTypes** | array<string> |  | [] | [] | No |
| **BeginTime** | integer |  | [] | [] | No |
| **EndTime** | integer |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ResourceUsageID** | string |  | [] |


## 删除资源使用情况报告 - DeleteResourceUsage

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **ResourceUsageID** | string |  | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取资源使用情况列表 - ListResourceUsages

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **Keyword** | string |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer |  | [] |
| **Infos** | array<[*ResourceUsageBaseInfo*](#ResourceUsageBaseInfo)> |  | [] |

### 数据模型

#### InspectionResourceUsageCompanyInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] |
| **Name** | string | 租户名称 | [] |
| **Email** | string | 邮箱 | [] |

#### InspectionResourceUsageProjectInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ProjectID** | string | 项目组ID | [] |
| **ProjectName** | string | 项目组名称 | [] |

#### InspectionResourceUsageQueryCriteria

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Region** | string |  | [] |
| **CompanyInfos** | array<[*InspectionResourceUsageCompanyInfo*](#InspectionResourceUsageCompanyInfo)> |  | [] |
| **ProjectInfos** | array<[*InspectionResourceUsageProjectInfo*](#InspectionResourceUsageProjectInfo)> |  | [] |
| **ResourceTypes** | array<string> |  | [] |
| **BeginTime** | integer |  | [] |
| **EndTime** | integer |  | [] |
| **RegionAlias** | string |  | [] |

#### ResourceUsageBaseInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ResourceUsageID** | string |  | [] |
| **Name** | string |  | [] |
| **Status** | string |  | [] |
| **CreateTime** | integer |  | [] |
| **QueryCriteria** | object |  | [] |

## 获取资源使用情况详细信息 - GetResourceUsage

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **ResourceUsageID** | string |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ResourceUsageID** | string |  | [] |
| **QueryCriteria** | object |  | [] |
| **TotalCount** | integer |  | [] |
| **Infos** | array<[*InspectionResourceUsageInfo*](#InspectionResourceUsageInfo)> |  | [] |

### 数据模型

#### InspectionResourceUsageCompanyInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] |
| **Name** | string | 租户名称 | [] |
| **Email** | string | 邮箱 | [] |

#### InspectionResourceUsageDetail

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ResourceID** | string |  | [] |
| **ResourceName** | string |  | [] |
| **TotalCount** | integer |  | [] |
| **Infos** | array<[*InspectionResourceUsageForm*](#InspectionResourceUsageForm)> |  | [] |

#### InspectionResourceUsageForm

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |
| **Unit** | string |  | [] |

#### InspectionResourceUsageInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Region** | string |  | [] |
| **CompanyID** | integer |  | [] |
| **ProjectID** | string |  | [] |
| **ResourceType** | string |  | [] |
| **TotalCount** | integer |  | [] |
| **Infos** | array<[*InspectionResourceUsageDetail*](#InspectionResourceUsageDetail)> |  | [] |

#### InspectionResourceUsageProjectInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ProjectID** | string | 项目组ID | [] |
| **ProjectName** | string | 项目组名称 | [] |

#### InspectionResourceUsageQueryCriteria

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Region** | string |  | [] |
| **CompanyInfos** | array<[*InspectionResourceUsageCompanyInfo*](#InspectionResourceUsageCompanyInfo)> |  | [] |
| **ProjectInfos** | array<[*InspectionResourceUsageProjectInfo*](#InspectionResourceUsageProjectInfo)> |  | [] |
| **ResourceTypes** | array<string> |  | [] |
| **BeginTime** | integer |  | [] |
| **EndTime** | integer |  | [] |
| **RegionAlias** | string |  | [] |

## 获取资源事件过程 - DescribeResourceEvent

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Region** | string | 地域 | [] | [] | No |
| **ResourceIDs** | array<string> | 资源ID | [] | [] | No |
| **ResourceTypes** | array<string> | 资源类型 | [] | [] | No |
| **EventTypes** | array<string> | 事件类型 | [] | [] | No |
| **Levels** | array<string> | 事件等级 | [] | [] | No |
| **BeginTime** | integer | 开始时间 | [] | [] | No |
| **EndTime** | integer | 结束时间 | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **ProjectIDs** | array<string> | 项目组 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*ResourceEventInfo*](#ResourceEventInfo)> | 资源事件信息 | [] |
| **TotalCount** | integer |  | [] |

### 数据模型

#### ResourceEventInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ResourceEventID** | string | 资源事件ID | [] |
| **CompanyID** | integer | 租户ID | [] |
| **Region** | string | 地域 | [] |
| **ResourceID** | string | 资源ID | [] |
| **ResourceType** | string | 资源类型 | [] |
| **Type** | string | 事件类型 | [] |
| **Level** | string | 事件等级 | [] |
| **Content** | string | 事件内容 | [] |
| **StartTime** | integer | 事件开始时间 | [] |
| **UpdateTime** | integer |  | [] |
| **Count** | integer | 次数 | [] |
| **ResourceName** | string | 资源名称 | [] |
| **CompanyEmail** | string | 租户名称 | [] |

## 获取资源事件状态 - DescribeResourceCondition

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | No |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **ResourceIDs** | array<string> | 资源ID | [] | [] | No |
| **ResourceTypes** | array<string> | 资源类型 | [] | [] | No |
| **ProjectIDs** | array<string> | 项目组 | [] | [] | No |
| **Status** | string | 状态 | [] | [] | No |
| **SortBy** | string | 排序指标 | [] | [] | No |
| **Sort** | string | 排序方向 | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*ConditionInfo*](#ConditionInfo)> | 资源事件信息 | [] |
| **TotalCount** | integer |  | [] |

### 数据模型

#### ConditionInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Region** | string | 地域 | [] |
| **CompanyID** | integer | 租户ID | [] |
| **CompanyEmail** | string | 租户名称 | [] |
| **ResourceID** | string | 资源ID | [] |
| **ResourceName** | string | 资源名称 | [] |
| **ResourceType** | string | 资源类型 | [] |
| **Type** | string | 类型 | [] |
| **TypeDesc** | string | 类型翻译 | [] |
| **Status** | string | 状态 | [] |
| **Content** | string | 状态详细内容 | [] |
| **LastTransitionTime** | integer | 最近一次状态变更时间 | [] |

## 获取网络拓扑信息 - DescribeNetworkTopology

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **ProjectIDs** | array<string> | 项目组ID | [] | [] | No |
| **ResourceTypes** | array<string> | 资源类型 | [] | [] | No |
| **SubnetIDs** | array<string> | 子网ID | [] | [] | No |
| **SegmentIDs** | array<string> | 外网线路ID | [] | [] | No |
| **ResourceIDs** | array<string> | 资源ID | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **FlatNetworkIDs** | array<string> | 扁平网络ID | [] | [] | No |
| **DirectConnectIDs** | array<string> | 专线 ID | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer |  | [] |
| **Infos** | array<[*InspectionNetwork*](#InspectionNetwork)> |  | [] |
| **SummaryInfos** | array<[*ResourceUsageInfo*](#ResourceUsageInfo)> |  | [] |

### 数据模型

#### InspectionNetwork

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Region** | string |  | [] |
| **CompanyID** | integer |  | [] |
| **Email** | string |  | [] |
| **ResourceName** | string |  | [] |
| **VpcName** | string |  | [] |
| **SubnetName** | string |  | [] |
| **SegmentName** | string |  | [] |
| **SegmentNetwork** | string |  | [] |
| **ResourceID** | string |  | [] |
| **ResourceType** | string |  | [] |
| **VpcID** | string |  | [] |
| **SubnetID** | string |  | [] |
| **SegmentID** | string |  | [] |
| **IpID** | string |  | [] |
| **IpAddress** | string |  | [] |
| **IpVersion** | string |  | [] |
| **NetworkType** | string |  | [] |
| **MacAddresses** | array<string> |  | [] |
| **IsVip** | bool |  | [] |
| **Bandwidth** | integer |  | [] |
| **Status** | string |  | [] |
| **FlatNetworkName** | string |  | [] |
| **FlatNetworkID** | string |  | [] |
| **IsPrimary** | bool | 是否主节点 | [] |

#### ResourceUsageInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ResourceType** | string | 资源类型 | [] |
| **ResourceUsageCount** | integer | 资源使用数量 | [] |

## 获取资源用量图表 - DescribeResourceChart

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | No |
| **ResourceType** | string | 资源类型 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*ChartInfo*](#ChartInfo)> | 图表内容 | [] |

### 数据模型

#### ChartInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Title** | string | 图表名 | [] |
| **Labels** | array<[*ChartLabel*](#ChartLabel)> | 额外的标签，比如集群图表需要知道集群的resource_id，监控大屏要把图标放置在不同的tab栏 | [] |
| **Items** | array<[*ChartItem*](#ChartItem)> | 图表元素 | [] |
| **ID** | string | 图表唯一标识，区分cpu，memory，storage | [] |

#### ChartItem

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string | 图表名称 | [] |
| **Value** | double | 图表值 | [] |
| **DisplayValue** | double | 显示值 | [] |
| **Unit** | string | 单位 | [] |
| **ItemID** | string |  | [] |
| **SubItems** | array<[*ChartSubItem*](#ChartSubItem)> | 图表元素 | [] |

#### ChartLabel

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string | 标签Key | [] |
| **Value** | string | 标签值 | [] |

#### ChartSubItem

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string | 图表名称 | [] |
| **Value** | double | 图表值 | [] |
| **DisplayValue** | double | 显示值 | [] |
| **Unit** | string | 单位 | [] |
