# PM

## 添加物理机 - CreatePM

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string |  | [] | [] | No |
| **CompanyID** | integer |  | [] | [] | No |
| **ProjectID** | string |  | [] | [] | No |
| **Name** | string |  | [] | [] | No |
| **Remark** | string |  | [] | [] | No |
| **IPMIIP** | string |  | [] | [] | No |
| **IPMIUser** | string |  | [] | [] | No |
| **IPMIPassword** | string |  | [] | [] | No |
| **RackLocation** | string |  | [] | [] | No |
| **Label** | string |  | [] | [] | No |
| **DriveType** | string |  | [] | [] | No |
| **CustomMetricsPath** | string |  | [] | [] | No |
| **TargetCompanyID** | integer | 接受分配的租户ID | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **PMID** | string |  | [] |


## 删除物理机 - DeletePM

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string |  | [] | [] | No |
| **PMID** | string |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 开机 - PowerONPM

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string |  | [] | [] | No |
| **PMID** | string |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 关机 - PowerSoftPM

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string |  | [] | [] | No |
| **PMID** | string |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 断电 - PowerOFFPM

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string |  | [] | [] | No |
| **PMID** | string |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 重启 - PowerResetPM

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string |  | [] | [] | No |
| **PMID** | string |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 关机并重新开机 - PowerCyclePM

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string |  | [] | [] | No |
| **PMID** | string |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 准备控制台 - PreparePMKVM

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string |  | [] | [] | No |
| **PMID** | string |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 释放控制台 - CleanPMKVM

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string |  | [] | [] | No |
| **PMID** | string |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取物理机信息 - DescribePM

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **Keyword** | string |  | [] | [] | No |
| **CompanyID** | integer |  | [] | [] | No |
| **PMIDs** | array<string> |  | [] | [] | No |
| **ProjectIDs** | array<string> |  | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*PMInfo*](#PMInfo)> |  | [] |
| **TotalCount** | integer |  | [] |

### 数据模型

#### BMCDetails

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **BMCType** | string |  | [] |
| **Address** | string |  | [] |
| **CredentialsName** | string |  | [] |
| **DisableCertificateVerification** | bool |  | [] |

#### CPUInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ID** | string |  | [] |
| **Present** | bool |  | [] |
| **Model** | string |  | [] |
| **CoreNumber** | integer |  | [] |
| **InUsedCoreNumber** | integer |  | [] |
| **RatedPower** | integer |  | [] |
| **L1Cache** | integer |  | [] |
| **L2Cache** | integer |  | [] |
| **L3Cache** | integer |  | [] |

#### IntrospectionCPU

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ModelName** | string |  | [] |
| **Frequency** | string |  | [] |
| **Count** | integer |  | [] |
| **Architecture** | string |  | [] |

#### IntrospectionDisk

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string |  | [] |
| **Capacity** | integer |  | [] |
| **ByPath** | string |  | [] |

#### IntrospectionNIC

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string |  | [] |
| **Mac** | string |  | [] |
| **State** | string |  | [] |

#### KVMStatus

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Phase** | string |  | [] |
| **URL** | string |  | [] |
| **FailureMessage** | string |  | [] |
| **SupportKVM** | bool |  | [] |

#### MemoryInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ID** | string |  | [] |
| **Present** | bool |  | [] |
| **Manufacture** | string |  | [] |
| **Capacity** | integer |  | [] |
| **Speed** | integer |  | [] |
| **SN** | string |  | [] |

#### PCIEInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Present** | bool |  | [] |
| **VendorID** | string |  | [] |
| **DeviceID** | string |  | [] |
| **Type** | string |  | [] |
| **Width** | string |  | [] |
| **Speed** | string |  | [] |

#### PMIPInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **IPID** | string |  | [] |
| **Type** | string |  | [] |
| **VPCID** | string |  | [] |
| **SubnetID** | string |  | [] |
| **Address** | string |  | [] |
| **Protocol** | string |  | [] |

#### PMInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string |  | [] |
| **Remark** | string |  | [] |
| **Region** | string |  | [] |
| **RegionAlias** | string |  | [] |
| **CompanyID** | integer |  | [] |
| **CompanyName** | string |  | [] |
| **Email** | string |  | [] |
| **ProjectID** | string |  | [] |
| **ProjectName** | string |  | [] |
| **Reason** | string |  | [] |
| **CreateTime** | integer |  | [] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> |  | [] |
| **ExpireTime** | integer |  | [] |
| **PMID** | string |  | [] |
| **IPMIIP** | string |  | [] |
| **IPMIUser** | string |  | [] |
| **IPMIPassword** | string |  | [] |
| **IPMIIsConnected** | bool |  | [] |
| **RackLocation** | string |  | [] |
| **Label** | string |  | [] |
| **Status** | string |  | [] |
| **PowerStatus** | string |  | [] |
| **SN** | string |  | [] |
| **ProductName** | string |  | [] |
| **KVMStatus** | object |  | [] |
| **CPUsInfo** | array<[*CPUInfo*](#CPUInfo)> |  | [] |
| **MemoriesInfo** | array<[*MemoryInfo*](#MemoryInfo)> |  | [] |
| **NICsInfo** | array<[*PmNICInfo*](#PmNICInfo)> |  | [] |
| **PCIEsInfo** | array<[*PCIEInfo*](#PCIEInfo)> |  | [] |
| **CustomMetricsPath** | string |  | [] |
| **ImageID** | string |  | [] |
| **Hostname** | string |  | [] |
| **BootDevice** | string |  | [] |
| **IntrospectionNICsInfo** | array<[*IntrospectionNIC*](#IntrospectionNIC)> |  | [] |
| **IntrospectionDisksInfo** | array<[*IntrospectionDisk*](#IntrospectionDisk)> |  | [] |
| **IntrospectionArch** | string |  | [] |
| **IntrospectionCPU** | object |  | [] |
| **IntrospectionMemory** | integer |  | [] |
| **IPInfos** | array<[*PMIPInfo*](#PMIPInfo)> |  | [] |
| **BMCDetails** | object |  | [] |
| **FailureReason** | string |  | [] |

#### PmNICInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ID** | string |  | [] |
| **Present** | bool |  | [] |
| **Manufacture** | string |  | [] |
| **Device** | string |  | [] |
| **MAC** | string |  | [] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

## 更新物理机信息 - UpdatePM

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string |  | [] | [] | No |
| **CompanyID** | integer |  | [] | [] | No |
| **Name** | string |  | [] | [] | No |
| **Remark** | string |  | [] | [] | No |
| **PMID** | string |  | [] | [] | No |
| **IPMIUser** | string |  | [] | [] | No |
| **IPMIPassword** | string |  | [] | [] | No |
| **RackLocation** | string |  | [] | [] | No |
| **Label** | string |  | [] | [] | No |
| **IPMIIP** | string |  | [] | [] | No |
| **CustomMetricsPath** | string |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 申请物理机远程控制KVMVNC会话 - AllocatePMKVMVNCSession

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **PMID** | string | 裸金属ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **SessionID** | string |  | [] |


## 申请物理机远程控制VNC会话 - AllocatePMVNCSession

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **PMID** | string | 裸金属ID | [] | [] | **Yes** |
| **VNCPort** | integer | 裸金属VNC端口 | [] | [] | **Yes** |
| **VNCPassword** | string | 裸金属VNC密码 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **SessionID** | string |  | [] |


## 上架物理机 - UpperPM

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **PMID** | string | 裸金属ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 物理机装载GuestOS - InstallGuestOS

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **PMID** | string | 裸金属ID | [] | [] | **Yes** |
| **ImageID** | string | 装机镜像 | [] | [] | **Yes** |
| **UserName** | string | 用户名 | [] | [] | **Yes** |
| **Password** | string | 虚拟机密码 | [] | [] | **Yes** |
| **Hostname** | string |  | [] | [] | No |
| **BootDevice** | string |  | [] | [] | No |
| **PhysicsIPAddress** | string |  | [] | [] | No |
| **PhysicsNICs** | array<string> |  | [] | [] | No |
| **OperatorName** | string |  | [] | [] | No |
| **IPVersion** | string |  | [] | [] | No |
| **InternetIP** | string |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 物理机重装GuestOS - ReinstallGuestOS

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **PMID** | string | 裸金属ID | [] | [] | **Yes** |
| **ImageID** | string | 装机镜像 | [] | [] | **Yes** |
| **UserName** | string | 装机镜像 | [] | [] | **Yes** |
| **Password** | string | 虚拟机密码 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 分配节点给租户 - AllocatePM

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **PMID** | string | 裸金属ID | [] | [] | **Yes** |
| **ProjectID** | string |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 将节点从租户回收 - RecyclePM

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 租户ID | [] | [] | **Yes** |
| **PMID** | string | 裸金属ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取物理网络 - GetPhysicsVPCSubnet

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Info** | object |  | [] |

### 数据模型

#### PhysicsSubnet

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **SubnetID** | string |  | [] |
| **Network** | string |  | [] |

#### PhysicsVPCSubnetInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **VPCID** | string |  | [] |
| **Network** | string |  | [] |
| **PhysicsSubnets** | array<[*PhysicsSubnet*](#PhysicsSubnet)> |  | [] |

## 获取带内操作开关 - GetIntrospectionSwitch

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **IsOpen** | bool |  | [] |


## 设置装机网络 - SetPhysicsVPCSubnet

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **VLANID** | integer | VLANID | [] | [] | No |
| **Network** | string | 网段 | [] | [] | No |
| **StartIP** | string | 开始 IP | [] | [] | No |
| **EndIP** | string | 结束 IP | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取物理机列表V2 - ListPMV2

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer |  | [] | [] | No |
| **ProjectIDs** | array<string> |  | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **SortBy** | string |  | [] | [] | No |
| **SortOrder** | string |  | [] | [] | No |
| **PMIDs** | array<string> |  | [] | [] | No |
| **SNs** | array<string> |  | [] | [] | No |
| **Status** | string |  | [] | [] | No |
| **Hostname** | string |  | [] | [] | No |
| **SN** | string |  | [] | [] | No |
| **IP** | string |  | [] | [] | No |
| **IPMIIP** | string |  | [] | [] | No |
| **Tag** | string |  | [] | [] | No |
| **Keyword** | string |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*PMInfoV2*](#PMInfoV2)> |  | [] |
| **TotalCount** | integer |  | [] |

### 数据模型

#### PMInfoV2

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Region** | string |  | [] |
| **CompanyID** | integer |  | [] |
| **CompanyName** | string |  | [] |
| **ProjectID** | string |  | [] |
| **ProjectName** | string |  | [] |
| **Name** | string |  | [] |
| **Remark** | string |  | [] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> |  | [] |
| **CreateTime** | integer |  | [] |
| **UpdateTime** | integer |  | [] |
| **Email** | string |  | [] |
| **Reason** | string |  | [] |
| **ExpireTime** | integer |  | [] |
| **PMID** | string |  | [] |
| **SN** | string |  | [] |
| **Hostname** | string |  | [] |
| **Status** | string |  | [] |
| **IPMIIP** | string |  | [] |
| **IPMIUsername** | string |  | [] |
| **IPMIPassword** | string |  | [] |
| **BMCTypeName** | string |  | [] |
| **BMCInfo** | string |  | [] |
| **IP** | string |  | [] |
| **Netmask** | string |  | [] |
| **Gateway** | string |  | [] |
| **InstallMode** | string |  | [] |
| **ValidityCheckResults** | array<string> |  | [] |
| **PMTags** | array<string> |  | [] |
| **Description** | string |  | [] |
| **HardwareInfo** | string |  | [] |
| **RackLocation** | string |  | [] |
| **CustomMetricsPath** | string |  | [] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

## 创建物理机V2 - CreatePMV2

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer |  | [] | [] | No |
| **ProjectID** | string |  | [] | [] | No |
| **Name** | string |  | [] | [] | No |
| **Remark** | string |  | [] | [] | No |
| **Tags** | array<UnifiedTag> |  | [] | [] | No |
| **Reason** | string |  | [] | [] | No |
| **ExpireTime** | integer |  | [] | [] | No |
| **IPMIIP** | string |  | [] | [] | No |
| **IPMIUsername** | string |  | [] | [] | No |
| **IPMIPassword** | string |  | [] | [] | No |
| **BMCTypeName** | string |  | [] | [] | No |
| **Hostname** | string |  | [] | [] | No |
| **IP** | string |  | [] | [] | No |
| **Netmask** | string |  | [] | [] | No |
| **Gateway** | string |  | [] | [] | No |
| **PMTags** | array<string> |  | [] | [] | No |
| **Description** | string |  | [] | [] | No |
| **RackLocation** | string |  | [] | [] | No |
| **CustomMetricsPath** | string |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **PMID** | string |  | [] |
| **SN** | string |  | [] |
| **Message** | string |  | [] |


## 更新物理机V2 - UpdatePMV2

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **PMID** | string | 物理机ID | [] | [] | **Yes** |
| **Name** | string |  | [] | [] | No |
| **Remark** | string |  | [] | [] | No |
| **IPMIIP** | string |  | [] | [] | No |
| **IPMIUsername** | string |  | [] | [] | No |
| **IPMIPassword** | string |  | [] | [] | No |
| **BMCTypeName** | string |  | [] | [] | No |
| **RackLocation** | string |  | [] | [] | No |
| **CustomMetricsPath** | string |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Message** | string |  | [] |


## 删除物理机V2 - DeletePMV2

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **PMIDs** | array<string> | 物理机ID列表 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Message** | string |  | [] |


## 电源控制V2 - PowerControlPMV2

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **PMID** | string | 物理机ID | [] | [] | **Yes** |
| **PowerAction** | string | 电源操作 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Message** | string |  | [] |


## 获取电源状态V2 - GetPMPowerStatusV2

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **PMID** | string | 物理机ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **PowerStatus** | string |  | [] |
| **LastChecked** | string |  | [] |
| **Message** | string |  | [] |


## 测试IPMI连接V2 - TestPMIPMIV2

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **IPMIIP** | string | IPMI 管理IP | [] | [] | **Yes** |
| **IPMIUsername** | string | IPMI 用户名 | [] | [] | **Yes** |
| **IPMIPassword** | string | IPMI 密码 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Success** | bool |  | [] |
| **SerialNumber** | string |  | [] |
| **BMCInfo** | string |  | [] |
| **Message** | string |  | [] |


## 检测BMC类型V2 - DetectBMCTypeV2

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **IPMIIP** | string | IPMI 管理IP | [] | [] | **Yes** |
| **IPMIUsername** | string | IPMI 用户名 | [] | [] | **Yes** |
| **IPMIPassword** | string | IPMI 密码 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Success** | bool |  | [] |
| **BMCTypeName** | string |  | [] |
| **BMCInfo** | string |  | [] |
| **SerialNumber** | string |  | [] |
| **Message** | string |  | [] |


## 硬件发现V2 - DiscoverPMHardwareV2

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **PMID** | string | 物理机ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Status** | string |  | [] |
| **Message** | string |  | [] |


## 获取硬件信息V2 - GetPMHardwareV2

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **PMID** | string | 物理机ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **HardwareInfo** | string |  | [] |
| **LastUpdateTime** | integer |  | [] |


## 获取OS媒体列表V2 - ListOSMediaV2

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer |  | [] | [] | No |
| **ProjectIDs** | array<string> |  | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **MediaType** | string |  | [] | [] | No |
| **OSName** | string | 操作系统完整名称（如：CentOS 7.4 x86_64） | [] | [] | No |
| **OSDistribution** | string | 操作系统发行版（如：CentOS、OpenEuler） | [] | [] | No |
| **Keyword** | string |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*OSMediaInfoV2*](#OSMediaInfoV2)> |  | [] |
| **TotalCount** | integer |  | [] |

### 数据模型

#### OSMediaInfoV2

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **MediaID** | string |  | [] |
| **Name** | string |  | [] |
| **Type** | string |  | [] |
| **URL** | string |  | [] |
| **OSName** | string | 操作系统完整名称（如：CentOS 7.4 x86_64） | [] |
| **OSDistribution** | string | 操作系统发行版（如：CentOS、OpenEuler） | [] |
| **FileSize** | integer |  | [] |
| **SHA256** | string |  | [] |
| **Description** | string |  | [] |
| **IsDefault** | bool |  | [] |
| **Status** | string |  | [] |
| **ErrorMessage** | string |  | [] |
| **CreatedAt** | integer |  | [] |
| **UpdatedAt** | integer |  | [] |
| **Region** | string |  | [] |
| **CompanyID** | integer |  | [] |
| **ProjectID** | string |  | [] |
| **ImageID** | string |  | [] |

## 创建OS媒体V2 - CreateOSMediaV2

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer | 云管平台租户ID | [] | [] | **Yes** |
| **Name** | string | 媒体名称 | [] | [] | **Yes** |
| **Type** | string | 媒体类型 | [] | [] | **Yes** |
| **ImageID** | string | 镜像ID | [] | [] | **Yes** |
| **OSName** | string | 操作系统完整名称（如：CentOS 7.4 x86_64） | [] | [] | No |
| **OSDistribution** | string | 操作系统发行版（如：CentOS、OpenEuler） | [] | [] | No |
| **Remark** | string |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **MediaID** | string |  | [] |
| **ImageID** | string |  | [] |
| **Message** | string |  | [] |


## 获取OS媒体详情V2 - GetOSMediaV2

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **MediaID** | string | 媒体ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Info** | object |  | [] |

### 数据模型

#### OSMediaInfoV2

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **MediaID** | string |  | [] |
| **Name** | string |  | [] |
| **Type** | string |  | [] |
| **URL** | string |  | [] |
| **OSName** | string | 操作系统完整名称（如：CentOS 7.4 x86_64） | [] |
| **OSDistribution** | string | 操作系统发行版（如：CentOS、OpenEuler） | [] |
| **FileSize** | integer |  | [] |
| **SHA256** | string |  | [] |
| **Description** | string |  | [] |
| **IsDefault** | bool |  | [] |
| **Status** | string |  | [] |
| **ErrorMessage** | string |  | [] |
| **CreatedAt** | integer |  | [] |
| **UpdatedAt** | integer |  | [] |
| **Region** | string |  | [] |
| **CompanyID** | integer |  | [] |
| **ProjectID** | string |  | [] |
| **ImageID** | string |  | [] |

## 更新OS媒体V2 - UpdateOSMediaV2

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **MediaID** | string | 媒体ID | [] | [] | **Yes** |
| **Name** | string |  | [] | [] | No |
| **OSName** | string | 操作系统完整名称（如：CentOS 7.4 x86_64） | [] | [] | No |
| **OSDistribution** | string | 操作系统发行版（如：CentOS、OpenEuler） | [] | [] | No |
| **Description** | string |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Message** | string |  | [] |


## 删除OS媒体V2 - DeleteOSMediaV2

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **MediaID** | string | 媒体ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Message** | string |  | [] |


## 批量删除OS媒体V2 - BatchDeleteOSMediaV2

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **MediaIDs** | array<string> | 媒体ID列表 | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **SuccessCount** | integer |  | [] |
| **FailedCount** | integer |  | [] |
| **FailedMediaIDs** | array<string> |  | [] |
| **Message** | string |  | [] |


## 创建装机任务V2 - CreateInstallTaskV2

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **PMID** | string | 物理机ID | [] | [] | **Yes** |
| **Mode** | string | 安装模式 | [] | [] | **Yes** |
| **ImageURL** | string | 镜像URL (clone模式必填) | [] | [] | No |
| **OSMediaID** | string | 操作系统介质ID (kickstart模式必填) | [] | [] | No |
| **TemplateID** | string | 模板ID (kickstart模式必填) | [] | [] | No |
| **TemplateVars** | string | 模板变量 (JSON格式) | [] | [] | No |
| **PartitionTemplateID** | string | 分区模板ID | [] | [] | No |
| **Hostname** | string | 主机名 | [] | [] | **Yes** |
| **Username** | string | 用户名 | [] | [] | **Yes** |
| **Password** | string | 密码 | [] | [] | **Yes** |
| **SSHKeyIDs** | array<string> | SSH密钥ID列表 | [] | [] | No |
| **SSHPublicKeys** | array<string> | SSH公钥列表 | [] | [] | No |
| **DiskSelector** | object | 磁盘选择器 | [] | [] | **Yes** |
| **Networks** | array<NetworkConfigV2> | 网络配置 | [] | [] | **Yes** |
| **PostInstallScript** | string | 安装后脚本 | [] | [] | No |
| **OSName** | string | 操作系统完整名称（如：CentOS 7.4 x86_64） | [] | [] | **Yes** |
| **OSDistribution** | string | 操作系统发行版（如：CentOS、OpenEuler） | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TaskID** | string |  | [] |
| **Message** | string |  | [] |


## 取消装机任务V2 - CancelInstallTaskV2

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **TaskID** | string | 任务ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Message** | string |  | [] |


## 获取装机日志V2 - GetInstallLogsV2

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **TaskID** | string | 任务ID | [] | [] | **Yes** |
| **Lines** | integer |  | [] | [] | No |
| **Since** | string |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Logs** | array<[*InstallLogEntryV2*](#InstallLogEntryV2)> |  | [] |
| **Lines** | integer |  | [] |

### 数据模型

#### InstallLogEntryV2

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Time** | string |  | [] |
| **Level** | string |  | [] |
| **Message** | string |  | [] |

## 获取装机任务列表V2 - ListInstallTasksV2

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer |  | [] | [] | No |
| **ProjectIDs** | array<string> |  | [] | [] | No |
| **Limit** | integer |  | [] | [] | No |
| **Offset** | integer |  | [] | [] | No |
| **SortBy** | string |  | [] | [] | No |
| **SortOrder** | string |  | [] | [] | No |
| **Status** | string |  | [] | [] | No |
| **MachineSN** | string |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*InstallTaskInfoV2*](#InstallTaskInfoV2)> |  | [] |
| **TotalCount** | integer |  | [] |

### 数据模型

#### InstallTaskInfoV2

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TaskID** | string |  | [] |
| **PMID** | string |  | [] |
| **SN** | string |  | [] |
| **Hostname** | string |  | [] |
| **Status** | string |  | [] |
| **Progress** | integer |  | [] |
| **InstallMode** | string |  | [] |
| **CurrentStage** | string |  | [] |
| **ErrorMessage** | string |  | [] |
| **CreatedAt** | integer |  | [] |
| **UpdatedAt** | integer |  | [] |
| **StartTime** | integer |  | [] |
| **EndTime** | integer |  | [] |
| **EstimatedCompletionTime** | integer |  | [] |
| **Region** | string |  | [] |
| **CompanyID** | integer |  | [] |
| **ProjectID** | string |  | [] |
| **TargetDisk** | string |  | [] |
| **NetworkConfig** | string |  | [] |
| **ConfigSnapshot** | string |  | [] |
| **InstallRetries** | integer |  | [] |
| **MediaName** | string |  | [] |
| **MediaType** | string |  | [] |
| **ImageURL** | string |  | [] |

## 获取装机任务详情V2 - GetInstallTaskV2

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **TaskID** | string | 任务ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Task** | object |  | [] |

### 数据模型

#### InstallTaskInfoV2

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TaskID** | string |  | [] |
| **PMID** | string |  | [] |
| **SN** | string |  | [] |
| **Hostname** | string |  | [] |
| **Status** | string |  | [] |
| **Progress** | integer |  | [] |
| **InstallMode** | string |  | [] |
| **CurrentStage** | string |  | [] |
| **ErrorMessage** | string |  | [] |
| **CreatedAt** | integer |  | [] |
| **UpdatedAt** | integer |  | [] |
| **StartTime** | integer |  | [] |
| **EndTime** | integer |  | [] |
| **EstimatedCompletionTime** | integer |  | [] |
| **Region** | string |  | [] |
| **CompanyID** | integer |  | [] |
| **ProjectID** | string |  | [] |
| **TargetDisk** | string |  | [] |
| **NetworkConfig** | string |  | [] |
| **ConfigSnapshot** | string |  | [] |
| **InstallRetries** | integer |  | [] |
| **MediaName** | string |  | [] |
| **MediaType** | string |  | [] |
| **ImageURL** | string |  | [] |

## 重试装机任务V2 - RetryInstallTaskV2

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **TaskID** | string | 任务ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TaskID** | string |  | [] |
| **Status** | string |  | [] |
| **Message** | string |  | [] |


## 删除装机任务V2 - DeleteInstallTaskV2

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **TaskID** | string | 任务ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Message** | string |  | [] |


## 按任务ID获取装机状态V2 - GetInstallStatusByTaskIDV2

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **TaskID** | string | 任务ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Task** | object |  | [] |

### 数据模型

#### InstallTaskInfoV2

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TaskID** | string |  | [] |
| **PMID** | string |  | [] |
| **SN** | string |  | [] |
| **Hostname** | string |  | [] |
| **Status** | string |  | [] |
| **Progress** | integer |  | [] |
| **InstallMode** | string |  | [] |
| **CurrentStage** | string |  | [] |
| **ErrorMessage** | string |  | [] |
| **CreatedAt** | integer |  | [] |
| **UpdatedAt** | integer |  | [] |
| **StartTime** | integer |  | [] |
| **EndTime** | integer |  | [] |
| **EstimatedCompletionTime** | integer |  | [] |
| **Region** | string |  | [] |
| **CompanyID** | integer |  | [] |
| **ProjectID** | string |  | [] |
| **TargetDisk** | string |  | [] |
| **NetworkConfig** | string |  | [] |
| **ConfigSnapshot** | string |  | [] |
| **InstallRetries** | integer |  | [] |
| **MediaName** | string |  | [] |
| **MediaType** | string |  | [] |
| **ImageURL** | string |  | [] |

## 创建KVM会话V2 - CreateKVMSessionV2

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **PMID** | string | 物理机ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **SessionID** | string |  | [] |
| **URL** | string |  | [] |
| **Token** | string |  | [] |
| **Password** | string |  | [] |
| **ExpiresAt** | integer |  | [] |


## 关闭KVM会话V2 - CloseKVMSessionV2

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **PMID** | string | 物理机ID | [] | [] | **Yes** |
| **SessionID** | string | 会话ID | [] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Message** | string |  | [] |


## 获取KVM会话列表V2 - ListKVMSessionsV2

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域 | [] | [] | **Yes** |
| **CompanyID** | integer |  | [] | [] | No |
| **PMID** | string |  | [] | [] | No |
| **Status** | string |  | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*KVMSessionInfoV2*](#KVMSessionInfoV2)> |  | [] |
| **TotalCount** | integer |  | [] |

### 数据模型

#### KVMSessionInfoV2

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **SessionID** | string |  | [] |
| **PMID** | string |  | [] |
| **SN** | string |  | [] |
| **InstanceName** | string |  | [] |
| **ServerIP** | string |  | [] |
| **VNCPort** | integer |  | [] |
| **URL** | string |  | [] |
| **Token** | string |  | [] |
| **Status** | string |  | [] |
| **CreatedAt** | integer |  | [] |
| **ExpiresAt** | integer |  | [] |
| **Region** | string |  | [] |
| **CompanyID** | integer |  | [] |
