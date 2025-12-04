# USB

## 获取USB设备信息 - ListUSBs

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **HostIP** | string | 宿主机IP | [] | [] | No |
| **VMID** | string | 虚拟机ID | [] | [] | No |
| **USBDeviceIDs** | array<string> | USB设备ID | [] | [] | No |
| **FilterAttached** | bool | 是否过滤已加载的USB设备 | [] | [] | No |
| **ProjectIDs** | array<string> | 项目组ID | [] | [] | No |
| **SetType** | string | 计算集群 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer |  | [] |
| **Infos** | array<[*USBInfo*](#USBInfo)> |  | [] |

### 数据模型

#### USBInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string | 名称 | [] |
| **Remark** | string |  | [] |
| **Region** | string |  | [] |
| **RegionAlias** | string |  | [] |
| **CompanyID** | integer | 租户ID | [] |
| **CompanyName** | string |  | [] |
| **Email** | string | 租户账号 | [] |
| **Reason** | string | 加载失败原因 | [] |
| **CreateTime** | integer |  | [] |
| **ProjectID** | string |  | [] |
| **ProjectName** | string |  | [] |
| **USBDeviceID** | string | USB设备ID | [] |
| **Serial** | string | 序列号 | [] |
| **Version** | string | USB协议版本 | [] |
| **Vendor** | string | 厂商 | [] |
| **Class** | string | 产品类型 | [] |
| **Product** | string | 产品名称 | [] |
| **HostIP** | string | 宿主机IP | [] |
| **Status** | string | 加载状态 | [] |
| **AttachedVMID** | string | 加载的的虚拟机ID | [] |
| **AttachedType** | string | 加载类型 | [] |
| **ComputeSetID** | string | 宿主机所在的计算集群 ID | [] |

## 分配USB设备 - AllocateUSB

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string |  | [] | [] | **Yes** |
| **TargetCompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **USBDeviceID** | string | USBID | [] | [] | **Yes** |
| **ProjectID** | string |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 加载USB设备 - AttachUSB

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string |  | [] | [] | **Yes** |
| **CompanyID** | integer |  | [] | [] | No |
| **USBDeviceID** | string | USBID | [] | [] | **Yes** |
| **VMID** | string | 虚拟机ID | [] | [] | **Yes** |
| **AttachType** | string | 绑定类型 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 卸载USB设备 - DetachUSB

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string |  | [] | [] | **Yes** |
| **CompanyID** | integer |  | [] | [] | No |
| **USBDeviceID** | string | USBID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|

