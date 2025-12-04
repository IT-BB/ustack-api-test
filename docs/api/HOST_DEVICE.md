# HOST_DEVICE

## 创建外置设备 - CreateNodeHostDevice

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **ProjectID** | string | 项目ID | [] | [] | No |
| **HostDeviceName** | string | 名称 | [] | [] | **Yes** |
| **Remark** | string | 备注 | [] | [] | No |
| **NodeID** | string | 物理机节点 ID | [] | [] | **Yes** |
| **Bus** | integer | 物理机节点上的 bus，扫描会返回该数据 | [] | [] | **Yes** |
| **Device** | integer | 物理机节点上的 device | [] | [] | **Yes** |
| **VendorName** | string | 厂商名称 | [] | [] | No |
| **ProductName** | string | 产品名称 | [] | [] | No |
| **HostDeviceType** | string | 设备类型 | [] | [] | No |
| **TargetCompanyID** | integer | 接受分配的租户ID | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 弹出/删除外置设备 - DeleteNodeHostDevice

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **HostDeviceID** | string | 外置设备 ID | [] | [] | **Yes** |
| **HostDeviceType** | string | 外置设备类型 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 扫描外置设备 - DescribeNodeHostDevice

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **NodeID** | string | 物理机 ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*HostDeviceInfo*](#HostDeviceInfo)> |  | [] |

### 数据模型

#### HostDeviceInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **HostDeviceID** | string |  | [] |
| **HostDeviceName** | string |  | [] |
| **Remark** | string |  | [] |
| **CompanyID** | integer |  | [] |
| **CompnayName** | string |  | [] |
| **Email** | string |  | [] |
| **ProjectID** | string |  | [] |
| **ProjectName** | string |  | [] |
| **Reason** | string |  | [] |
| **CreateTime** | integer |  | [] |
| **UpdateTime** | integer |  | [] |
| **HostDeviceType** | string | 外置设备类型，例如：USB | [] |
| **VendorName** | string |  | [] |
| **VendorID** | string |  | [] |
| **ProductName** | string |  | [] |
| **ProductID** | string |  | [] |
| **Serial** | string | 外置设备序列号 | [] |
| **Status** | string | 外置设备状态 | [] |
| **Bus** | integer |  | [] |
| **Device** | integer |  | [] |
| **AttachedVMID** | string | 加载的的虚拟机ID | [] |
| **AttachedType** | string | 加载类型 | [] |

## 分配外置设备给租户 - AllocateNodeHostDevice

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string |  | [] | [] | **Yes** |
| **TargetCompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **HostDeviceID** | string | 外置设备 ID | [] | [] | **Yes** |
| **ProjectID** | string |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|

