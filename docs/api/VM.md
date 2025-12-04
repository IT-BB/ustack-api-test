# VM

## 创建虚拟机 - CreateVMInstance

### 请求参数

| 字段名                     | 类型          | 描述信息                                                                                                                                                                                                                                                                                                                                          | 示例                               | 参数来源                  | 必填    |
| :------------------------- | :------------ | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | :--------------------------------- | :------------------------ | :------ |
| **Region**                 | string        | 地域ID，地域，一批可共享的物理资源使用集合                                                                                                                                                                                                                                                                                                        | [develop]                          | [DescribeRegion]          | **Yes** |
| **ApplicationName**        | string        | 审批名称，用于审批流程的标题描述。当启用审批流程时使用。                                                                                                                                                                                                                                                                                          | [虚拟机资源申请]                   | []                        | No      |
| **ApplicationReason**      | string        | 审批理由，用于说明申请创建虚拟机的原因。当启用审批流程时使用。                                                                                                                                                                                                                                                                                    | [用于新项目测试环境搭建]           | []                        | No      |
| **CompanyID**              | integer       | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离                                                                                                                                                                                                                                                                                        | [200000236]                        | [DescribeUser]            | **Yes** |
| **Email**                  | string        | 租户邮箱，租户注册时提供的用于接收通知和联系的邮箱                                                                                                                                                                                                                                                                                                | [example@ucloud.cn]                | [DescribeUser]            | No      |
| **ProjectID**              | string        | 项目组ID，用于资源分组管理                                                                                                                                                                                                                                                                                                                        | [project-5kyt09jzwcqh6j]           | [ListProjects]            | No      |
| **TagKeyValuePairs**       | array<string> | 标签键值对，用于资源标记和分类管理。格式为 key:value 的字符串，传入​​ Base64 编码的字符串。                                                                                                                                                                                                                                                       | [dGVzdA==:dGVzdC0x]                | []                        | **Yes** |
| **Name**                   | string        | 虚拟机名称，长度为1-50个字符，名称只能包含中英文、数字、点（.），下划线（_）和中划线（-）                                                                                                                                                                                                                                                         | [vm-test]                          | []                        | **Yes** |
| **Remark**                 | string        | 备注，用于进行说明和注释，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空                                                                                                                                                                                                                                            | [这是一个备注]                     | []                        | No      |
| **ChargeType**             | string        | 计费类型，用于指定计费模式。取值范围：1、Dynamic（按小时计费）2、Month（按月计费）3、Year（按年计费）                                                                                                                                                                                                                                             | []                                 | []                        | **Yes** |
| **Quantity**               | integer       | 计费数量，指定计费周期的数量。按月/年计费时表示创建 Quantity 个月/年的磁盘，按小时计费时强制为1。取值范围：正整数                                                                                                                                                                                                                                 | []                                 | []                        | **Yes** |
| **VMType**                 | string        | 计算集群类型，创建虚拟机时可指定计算集群，虚拟机所运行物理机的集群类型，代表不同架构、不同型号的 CPU 或硬件特征。                                                                                                                                                                                                                                 | [ComputeSetTTPZ]                   | [DescribeVMSet]           | **Yes** |
| **BootSourceType**         | string        | 启动源类型，会影响虚拟机系统盘的创建方式。Image (默认)或 Disk                                                                                                                                                                                                                                                                                     | [Image]                            | []                        | No      |
| **ImageID**                | string        | 镜像ID，通过该 ID 获取虚拟机的镜像信息                                                                                                                                                                                                                                                                                                            | [image-k8s-node]                   | [DescribeImage]           | No      |
| **OSDistribution**         | string        | 系统类型，如 windows、centos，ImageID为空时必填                                                                                                                                                                                                                                                                                                   | [windows]                          | [DescribeImage]           | No      |
| **OSVersion**              | string        | 系统版本，如 xp 74，ImageID为空时必填                                                                                                                                                                                                                                                                                                             | [xp]                               | [DescribeImage]           | No      |
| **OSType**                 | string        | 操作系统类型，如 Linux，ImageID为空时必填                                                                                                                                                                                                                                                                                                         | [Linux]                            | [DescribeImage]           | No      |
| **CPU**                    | integer       | CPU核数                                                                                                                                                                                                                                                                                                                                           | [1]                                | []                        | **Yes** |
| **Memory**                 | integer       | 内存大小，单位：MiB，1024 的倍数                                                                                                                                                                                                                                                                                                                  | [1024]                             | []                        | **Yes** |
| **BootDiskID**             | string        | 系统盘ID，启动虚拟机所用磁盘                                                                                                                                                                                                                                                                                                                      | [vm-j40jm6vb8l213-boot]            | [DescribeDisk]            | No      |
| **BootDiskSetType**        | string        | 系统盘集群ID，未指定系统盘 ID 时必填，系统盘ID和系统盘集群ID至少需要指定一项                                                                                                                                                                                                                                                                      | []                                 | [DescribeStorageSet]      | No      |
| **BootDiskSpace**          | integer       | 系统盘大小，单位：GiB，未指定系统盘 ID 时必填，范围 20-2000                                                                                                                                                                                                                                                                                       | [100]                              | []                        | No      |
| **DataDiskSetType**        | string        | 数据盘集群ID，如果未指定数据盘 ID，用来创建数据盘                                                                                                                                                                                                                                                                                                 | [StorageSetNPVK]                   | [DescribeStorageSet]      | No      |
| **DataDiskSpace**          | integer       | 数据盘大小，如果未指定数据盘 ID，用来创建数据盘。最小值为 10                                                                                                                                                                                                                                                                                      | []                                 | []                        | No      |
| **DataDiskID**             | string        | 数据盘ID，指定已有的数据盘                                                                                                                                                                                                                                                                                                                        | [disk-me3aurdiccjf46]              | [DescribeDisk]            | No      |
| **DataDiskAutoMount**      | bool          | 数据盘是否自动挂载，若为 true，将挂载到 /data                                                                                                                                                                                                                                                                                                     | [true]                             | []                        | No      |
| **LANSGID**                | string        | 内网安全组ID，如果需要使用，需提前创建安全组，得到安全组 ID                                                                                                                                                                                                                                                                                       | [sg-v55mgokwcc11c5]                | [DescribeSecurityGroup]   | No      |
| **WANSGID**                | string        | 外网安全组ID，如果需要使用，需提前创建安全组，得到安全组 ID                                                                                                                                                                                                                                                                                       | [sg-v55mgokwcc11c5]                | [DescribeSecurityGroup]   | No      |
| **VPCID**                  | string        | VPC ID，创建虚拟机时必须选择 VPC 网络和所属子网，即选择虚拟要加入的网络及 IP 网段                                                                                                                                                                                                                                                                 | [vpc-el7miluuodq]                  | [DescribeVPC]             | **Yes** |
| **SubnetID**               | string        | 子网ID，创建虚拟机时必须选择 VPC 网络和所属子网，即选择虚拟要加入的网络及 IP 网段                                                                                                                                                                                                                                                                 | [subnet-lq10np1d231]               | [DescribeSubnet]          | **Yes** |
| **InternalIP**             | string        | 内网IP地址，支持 IPv4 或 IPv6，留空则自动分配                                                                                                                                                                                                                                                                                                     | [10.0.0.1]                         | []                        | No      |
| **OperatorName**           | string        | 外网线路ID                                                                                                                                                                                                                                                                                                                                        | [segment-35dehi0g8zdqqq]           | [DescribeSegment]         | **Yes** |
| **IPVersion**              | string        | IP版本，支持 IPv4 或 IPv6                                                                                                                                                                                                                                                                                                                         | [IPv4]                             | []                        | **Yes** |
| **Bandwidth**              | integer       | IP带宽，单位:Mb。当前 IP 地址的带宽上限，扁平网络可以不设置带宽                                                                                                                                                                                                                                                                                   | [1]                                | []                        | **Yes** |
| **InternetIP**             | string        | 外网IP地址，支持 IPv4 或 IPv6，留空则自动分配                                                                                                                                                                                                                                                                                                     | [192.31.31.1]                      | []                        | No      |
| **Password**               | string        | 管理员密码，要求：1、密码长度限6-30个字符 2、需要同时包含小写字母、特殊字符(除空格) 3、不能包含非法字符                                                                                                                                                                                                                                           | [f34f*2defrefe]                    | []                        | No      |
| **Hostname**               | string        | 主机名，要求：1、Windows 系统，长度为2~15个字符。允许使用大小写字母、数字或连接符（-）。不能以连字符（-）开头或结尾，不能连续使用连字符（-），也不能仅使用数字。2、Linux 系统，长度为2~63个字符。允许使用大小写字母、数字、点号（.）或连接符（-）。不能以点号（.）或连字符（-）开头或结尾，不能连续使用点号（.）或连字符（-），也不能仅使用数字。 | [windows-test]                     | []                        | No      |
| **USBDeviceID**            | string        | USB设备ID，用于挂载指定的 USB 设备                                                                                                                                                                                                                                                                                                                | [usb-test-eqweqw]                  | [ListUSBs]                | No      |
| **USBAttachType**          | string        | USB 挂载类型，取值范围：redir（重定向）、pass-through（直通）                                                                                                                                                                                                                                                                                     | [redir]                            | [ListUSBs]                | No      |
| **DNS**                    | string        | DNS                                                                                                                                                                                                                                                                                                                                               | [114.114.114.114]                  | []                        | No      |
| **CPUMode**                | string        | CPU 启动模式，取值范围：1、host-passthrough（直通模式）2、custom（自定义）                                                                                                                                                                                                                                                                        | [host-passthrough]                 | []                        | No      |
| **HighAvailability**       | string        | 虚拟机高可用模式，取值范围：1、NeverStop（默认）2、None                                                                                                                                                                                                                                                                                           | [NeverStop]                        | []                        | No      |
| **GPUType**                | string        | GPU 类型，取值范围：1、GPU  2、VGPU                                                                                                                                                                                                                                                                                                               | [GPU]                              | []                        | No      |
| **GPUMdevName**            | string        | GPU规格名称，GPUType=GPU时指定                                                                                                                                                                                                                                                                                                                    | [nvidia-2080ti]                    | [ListGPUs]                | No      |
| **GPU**                    | integer       | GPU的颗数，GPUType=GPU时指定                                                                                                                                                                                                                                                                                                                      | [2]                                | []                        | No      |
| **MdevName**               | string        | vGPU的规格名称，GPUType=VGPU时指定                                                                                                                                                                                                                                                                                                                | [nvidia-2080ti]                    | [ListVGPUs]               | No      |
| **IGID**                   | string        | 隔离组id，隔离组是一种针对虚拟机资源的简单编排策略                                                                                                                                                                                                                                                                                                | [isolationgroup-77go7p0prdwq]      | [DescribeIsolationGroups] | No      |
| **UserData**               | string        | cloud-init脚本，需要镜像支持，支持自定义初始化配置                                                                                                                                                                                                                                                                                                | [#cloud-config\nhostname: my-vm]   | []                        | No      |
| **BootDiskSecret**         | string        | 启动盘加密密钥，仅支持加密的存储可用                                                                                                                                                                                                                                                                                                              | [21fe921]                          | []                        | No      |
| **DataDiskSecret**         | string        | 数据盘加密密钥，仅支持加密的存储可用                                                                                                                                                                                                                                                                                                              | [21fe921]                          | []                        | No      |
| **BootloaderType**         | string        | 虚拟机引导方式，引导方式不匹配可能导致虚拟机无法正常工作，取值范围：bios、uefi                                                                                                                                                                                                                                                                    | [uefi]                             | []                        | No      |
| **SupportQGA**             | bool          | 镜像是否支持qemu-ga                                                                                                                                                                                                                                                                                                                               | [true]                             | []                        | No      |
| **SupportCloudInit**       | bool          | 镜像是否支持cloudinit                                                                                                                                                                                                                                                                                                                             | [true]                             | []                        | No      |
| **SupportHotplug**         | bool          | 镜像是否支持热升级                                                                                                                                                                                                                                                                                                                                | [true]                             | []                        | No      |
| **FlatNetworkSGID**        | string        | 扁平网络安全组ID                                                                                                                                                                                                                                                                                                                                  | [sg-flatnet-123456]                | [DescribeSecurityGroup]   | No      |
| **FlatNetworkID**          | string        | 扁平网络ID，带宽不为空时必传                                                                                                                                                                                                                                                                                                                      | [flatnet-abc123]                   | [DescribeFlatIP]          | No      |
| **CPUModel**               | string        | CPU 模型类型，仅在 CPUMode 为 custom 时生效                                                                                                                                                                                                                                                                                                       | [default Cascadelake-Server-noTSX] | [DescribeVMSet]           | No      |
| **LANMAC**                 | string        | 内网/扁平网络mac，不指定时自动生成                                                                                                                                                                                                                                                                                                                | [00:1A:2B:3C:4D:5E]                | []                        | No      |
| **WANMAC**                 | string        | 外网mac，不指定时自动生成                                                                                                                                                                                                                                                                                                                         | [00:1A:2B:3C:4D:5E]                | []                        | No      |
| **PFCode**                 | string        | 物理网卡型号的标准编号                                                                                                                                                                                                                                                                                                                            | [intel-x710]                       | [DescribeNode]            | No      |
| **LANInAverageBandwidth**  | integer       | 入向平均带宽，第一张网卡入向平均带宽（单位：Mbps），0 表示不限制                                                                                                                                                                                                                                                                                  | [1000]                             | []                        | No      |
| **LANOutAverageBandwidth** | integer       | 出向平均带宽，第一张网卡出向平均带宽（单位：Mbps），0 表示不限制                                                                                                                                                                                                                                                                                  | [1000]                             | []                        | No      |
| **DiskCacheMode**          | string        | 磁盘缓存模式，虚拟机重启、创建或热插拔时生效，取值范围：1、directsync（默认）2、none 3、writeback                                                                                                                                                                                                                                                 | [directsync]                       | []                        | No      |
| **VCPUBindingNode**        | string        | VCPU绑定的物理节点                                                                                                                                                                                                                                                                                                                                | [node-123456]                      | [DescribeNode]            | No      |
| **VCPUBindingInfos**       | array<string> | VCPU绑定信息                                                                                                                                                                                                                                                                                                                                      | []                                 | []                        | No      |
| **VCPUBindingDegrandable** | bool          | 是否允许降级，VCPU绑定在宕机迁移时是否允许降级                                                                                                                                                                                                                                                                                                    | [true]                             | []                        | No      |
| **GPUBindingEnable**       | bool          | 是否GPU绑定优化，指定节点需已开启vcpu绑定                                                                                                                                                                                                                                                                                                         | [true]                             | [DescribeNode]            | No      |
| **InternalIPVersion**      | string        | 内网协议，空值为所选VPC/外网默认网络的网络协议, 枚举支持 IPv4 IPv6 ALL和空值                                                                                                                                                                                                                                                                      | []                                 | []                        | No      |
| **InternalExpandIP**       | string        | 内网拓展IP地址                                                                                                                                                                                                                                                                                                                                    | []                                 | []                        | No      |

### 响应字段

| 字段名       | 类型   | 描述信息   | 示例                     |
| :----------- | :----- | :--------- | :----------------------- |
| **VMID**     | string | 虚拟机ID   | [vm-j40jm6vb8lo321]      |
| **DiskID**   | string | 磁盘ID     | [vm-j40jm6vb8lo321-boot] |
| **EIPID**    | string | 外网资源ID | [eip-789def]             |
| **FlatIPID** | string | 扁平网络ID | [flatip-456ghi]          |


## 获取虚拟机信息 - DescribeVMInstance

### 请求参数

| 字段名         | 类型          | 描述信息                                                                                                                                                                                                                              | 示例                          | 参数来源                  | 必填    |
| :------------- | :------------ | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | :---------------------------- | :------------------------ | :------ |
| **Region**     | string        | 地域ID，地域，一批可共享的物理资源使用集合                                                                                                                                                                                            | [develop]                     | [DescribeRegion]          | **Yes** |
| **ProjectIDs** | array<string> | 项目组ID，用于资源分组管理                                                                                                                                                                                                            | [project-5kyt09jzwcqh6j]      | [ListProjects]            | No      |
| **Keyword**    | string        | 搜索关键词                                                                                                                                                                                                                            | [example_key]                 | []                        | No      |
| **CompanyID**  | integer       | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离                                                                                                                                                                            | [200000236]                   | [DescribeUser]            | No      |
| **Limit**      | integer       | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。默认值：10                                                                                                                                                       | [20]                          | []                        | No      |
| **Offset**     | integer       | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 默认值：0                                                                                                                                                              | [0]                           | []                        | No      |
| **SortBy**     | string        | 排序指标，取值范围：1、CreateTime 2、CPUUtilization 3、MemUsage 4、SpaceUsage                                                                                                                                                         | [CreateTime]                  | []                        | No      |
| **Sort**       | string        | 排序方向，取值范围：1、Ascending 2、Descending                                                                                                                                                                                        | [Descending]                  | []                        | No      |
| **States**     | array<string> | 状态，过滤虚拟机的运行状态，取值范围：1、Starting 2、Restarting 3、Running 4、Stopping 5、Stopped 6、Poweroffing 7、Reinstalling 8、Deleting 9、Failed 10、RFailed 11、Migrating 12、MigrateDisking 13、Saving 14、Saved 15、Reverted | [Failed,RFailed]              | []                        | No      |
| **VMIDs**      | array<string> | 虚拟机ID列表，用于查询指定的虚拟机                                                                                                                                                                                                    | [vm-dqwg3r431rg1]             | [DescribeVMInstance]      | No      |
| **VPCID**      | string        | vpcID，通过指定vpcID进行筛选                                                                                                                                                                                                          | [vpc-el7miluuodq]             | [DescribeVPC]             | No      |
| **SubnetID**   | string        | 子网ID，通过指定子网ID进行筛选                                                                                                                                                                                                        | [subnet-lq10np1d231]          | [DescribeSubnet]          | No      |
| **IGID**       | string        | 隔离组ID，通过指定隔离组ID进行筛选                                                                                                                                                                                                    | [isolationgroup-77go7p0prdwq] | [DescribeIsolationGroups] | No      |
| **SetID**      | string        | 计算集群ID，通过指定集群ID进行筛选                                                                                                                                                                                                    | [set-abc123]                  | [DescribeVMSet]           | No      |
| **HostIP**     | string        | 物理机IP，通过指定宿主机IP进行筛选                                                                                                                                                                                                    | [192.168.0.12]                | [DescribeNode]            | No      |

### 响应字段

| 字段名         | 类型                       | 描述信息               | 示例 |
| :------------- | :------------------------- | :--------------------- | :--- |
| **TotalCount** | integer                    | 查询得到的总数         | []   |
| **Infos**      | array<[*VMInfo*](#VMInfo)> | 所查询的虚拟机信息详情 | []   |

### 数据模型

#### IPInfo

| 字段名            | 类型    | 描述信息                                                            | 示例                     |
| :---------------- | :------ | :------------------------------------------------------------------ | :----------------------- |
| **IPID**          | string  | ipID                                                                | [vm-z7un40imvhraji-lan0] |
| **SGName**        | string  | 绑定的安全组名称                                                    | [test]                   |
| **IPVersion**     | string  | IP版本                                                              | [IPv4 IPv6]              |
| **ISDefaultGW**   | integer | 是否是默认网关（是/否:1/0）                                         | [1]                      |
| **IsElastic**     | string  | 是否弹性IP                                                          | [Y N]                    |
| **IsVIP**         | string  | 是否是VIP                                                           | [Y N]                    |
| **Type**          | string  | 网络类型                                                            | [public private]         |
| **SGID**          | string  | 安全组ID                                                            | [sg-9mdqfo7pvin9wz]      |
| **IP**            | string  | IP地址                                                              | [172.16.0.7]             |
| **NicType**       | string  | 网卡类型，取值范围：1、LAN（内网）2、WAN（外网）3、Flat（扁平网络） | [LAN]                    |
| **InterfaceID**   | string  | 网卡ID                                                              | [vm-ds7z7z6qz6at9g-lan]  |
| **MAC**           | string  | MAC地址                                                             | [52:54:00:91:89:95]      |
| **Mode**          | string  | 绑定模式                                                            | [Direct NAT]             |
| **NicPFCode**     | string  | 物理网卡型号的标准编号                                              | [intel-x710]             |
| **NicPFVendor**   | string  | 物理网卡厂商                                                        | [Intel]                  |
| **NicPFProduct**  | string  | 物理网卡型号                                                        | [X710-DA4]               |
| **InterfaceName** | string  | 网卡名称                                                            | [eth0]                   |

#### NumaGPUInfo

| 字段名          | 类型   | 描述信息            | 示例           |
| :-------------- | :----- | :------------------ | :------------- |
| **VenderID**    | string | GPU厂商ID（16进制） | [10de]         |
| **ProductID**   | string | GPU设备ID（16进制） | [13f8]         |
| **HostPCIPath** | string | 主机PCI设备路径     | [0000:3b:00.0] |

#### PendingChange

| 字段名      | 类型   | 描述信息                 | 示例             |
| :---------- | :----- | :----------------------- | :--------------- |
| **Field**   | string | 待生效的配置字段名称     | [BootloaderType] |
| **Restart** | bool   | 是否需要电源重置才能生效 | [true]           |

#### UnifiedTag

| 字段名    | 类型   | 描述信息 | 示例 |
| :-------- | :----- | :------- | :--- |
| **Key**   | string |          | []   |
| **Value** | string |          | []   |

#### VCPUBinding

| 字段名   | 类型    | 描述信息                | 示例    |
| :------- | :------ | :---------------------- | :------ |
| **VCPU** | integer | 虚拟CPU编号             | [0]     |
| **PCPU** | string  | 绑定的物理CPU编号或范围 | [2-4,8] |

#### VMInfo

| 字段名                            | 类型                                     | 描述信息                                                                                                                             | 示例                                   |
| :-------------------------------- | :--------------------------------------- | :----------------------------------------------------------------------------------------------------------------------------------- | :------------------------------------- |
| **Name**                          | string                                   | 虚拟机名称                                                                                                                           | [test-vm]                              |
| **Remark**                        | string                                   | 虚拟机备注                                                                                                                           | [用于测试的虚拟机]                     |
| **Region**                        | string                                   | 地域                                                                                                                                 | [region-A]                             |
| **RegionAlias**                   | string                                   | 地域别名                                                                                                                             | [地域 A]                               |
| **CompanyID**                     | integer                                  | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离                                                                           | [200000236]                            |
| **CompanyName**                   | string                                   | 租户名称                                                                                                                             | [ucloud]                               |
| **Email**                         | string                                   | 租户邮箱，租户注册时提供的用于接收通知和联系的邮箱                                                                                   | [example@ucloud.cn]                    |
| **ProjectID**                     | string                                   | 项目组ID，用于资源分组管理                                                                                                           | [project-5kyt09jzwcqh6j]               |
| **ProjectName**                   | string                                   | 项目组名称                                                                                                                           | [project01]                            |
| **Reason**                        | string                                   | 创建失败原因                                                                                                                         | [磁盘创建失败]                         |
| **CreateTime**                    | integer                                  | 创建时间，Unix 时间戳（秒级）                                                                                                        | [1754644992]                           |
| **UpdateTime**                    | integer                                  | 更新时间，Unix 时间戳（秒级）                                                                                                        | [1754644992]                           |
| **Tags**                          | array<[*UnifiedTag*](#UnifiedTag)>       | 标签                                                                                                                                 | [cluster-resouce.system:offline-node]  |
| **ChargeType**                    | string                                   | 计费类型，取值范围：1、Dynamic（按小时计费）2、Month（按月计费）3、Year（按年计费）                                                  | [Month]                                |
| **ExpireTime**                    | integer                                  | 过期时间，Unix 时间戳（秒级）                                                                                                        | [1754644992]                           |
| **VMTypeAlias**                   | string                                   | 计算集群别名                                                                                                                         | [ComputeSetarm]                        |
| **BasicImageName**                | string                                   | 基础镜像名称                                                                                                                         | [e2e-arm]                              |
| **VPCName**                       | string                                   | vpc名称                                                                                                                              | [小子网]                               |
| **SubnetName**                    | string                                   | 子网名称                                                                                                                             | [subnet-01]                            |
| **CPUUtilization**                | float                                    | CPU利用率，记录在 resource annotations 当中, 10分钟平均CPU利用率, 非实时, 间隔3分钟检查更新一次, 仅在虚拟机 Running 状态下有值       | [0.6218442]                            |
| **MemUsage**                      | float                                    | 内存利用率，记录在 resource annotations 当中, 10分钟平均Mem利用率, 非实时, 间隔3分钟检查更新一次, 仅在虚拟机 Running 状态下有值      | [0]                                    |
| **SpaceUsage**                    | float                                    | 空间利用率，记录在 resource annotations 当中, 最新数据当中的空间利用率, 非实时, 间隔3分钟检查更新一次, 仅在虚拟机 Running 状态下有值 | [0]                                    |
| **VMID**                          | string                                   | 虚拟机ID                                                                                                                             | [vm-718j5rgm2wwz5d]                    |
| **ImageID**                       | string                                   | 镜像ID                                                                                                                               | [image-rcbliayh05vn06]                 |
| **OSType**                        | string                                   | 系统类型                                                                                                                             | [Linux]                                |
| **OSName**                        | string                                   | 系统名称                                                                                                                             | [OpenEuler 22.03_SP4 aarch64]          |
| **VMType**                        | string                                   | 计算集群                                                                                                                             | [ComputeSetarm]                        |
| **SetID**                         | string                                   | 集群ID                                                                                                                               | [compute-set-arm]                      |
| **SetArch**                       | string                                   | 集群架构                                                                                                                             | [aarch64]                              |
| **HostIP**                        | string                                   | 宿主机IP                                                                                                                             | [10.10.1.18]                           |
| **CPU**                           | integer                                  | CPU核数                                                                                                                              | [1]                                    |
| **Memory**                        | integer                                  | 内存大小，单位：MiB                                                                                                                  | [1024]                                 |
| **VPCID**                         | string                                   | vpcID                                                                                                                                | [vpc-0hesw42f1nvtld]                   |
| **SubnetID**                      | string                                   | 子网ID                                                                                                                               | [subnet-hdufcmhizdc95w]                |
| **DiskInfos**                     | array<[*VmDiskInfo*](#VmDiskInfo)>       | 磁盘信息，结构见DVmDiskInfo                                                                                                          | []                                     |
| **IPInfos**                       | array<[*IPInfo*](#IPInfo)>               | IP信息，结构见IPInfo                                                                                                                 | []                                     |
| **IGIDs**                         | array<string>                            | 隔离组ID列表，多个元素之间用,分隔。                                                                                                  | [ig-qw1f4321]                          |
| **State**                         | string                                   | 虚拟机运行状态                                                                                                                       | [Running]                              |
| **CanLogin**                      | bool                                     | 是否可以登录                                                                                                                         | [true]                                 |
| **StatusProcessProgress**         | double                                   | 状态处理进度，虚拟机创建中时表示镜像准备进度, 虚拟机迁移中时表示迁移进度                                                             | [0]                                    |
| **SupportHotplug**                | bool                                     | 是否支持热插拔                                                                                                                       | [true]                                 |
| **Hostname**                      | string                                   | 主机名                                                                                                                               | [wky-3]                                |
| **DNS**                           | string                                   | DNS                                                                                                                                  | [114.114.114.114,114.114.115.115]      |
| **CPUMode**                       | string                                   | CPU启动模式，取值范围：host-passthrough（直通模式）、custom（自定义）                                                                | [custom]                               |
| **HighAvailability**              | string                                   | 高可用模式，取值范围：1、NeverStop（默认）2、None                                                                                    | [NeverStop]                            |
| **GPUType**                       | string                                   | GPU的类型，枚举值：1、GPU  2、VGPU                                                                                                   | [GPU]                                  |
| **GPUMdevName**                   | string                                   | GPU规格名称，GPUType=GPU时指定                                                                                                       | [nvidia-2080ti]                        |
| **GPU**                           | integer                                  | GPU的颗数，GPUType=GPU时指定                                                                                                         | [2]                                    |
| **VGPUID**                        | string                                   | vGPUID，GPUType=VGPU时                                                                                                               | [vgpu-21verf31]                        |
| **MdevName**                      | string                                   | vGPU的规格名称，GPUType=VGPU时                                                                                                       | [nvidia-2080ti]                        |
| **QGAEnabled**                    | bool                                     | 是否启用 qemu-agent                                                                                                                  | [true]                                 |
| **QGARunning**                    | bool                                     | qemu-agent 是否在运行                                                                                                                | [true]                                 |
| **CloudInitEnabled**              | bool                                     | 是否启用 cloud-init                                                                                                                  | [true]                                 |
| **CloudInitRunning**              | bool                                     | cloud-init 是否在运行                                                                                                                | [true]                                 |
| **UserData**                      | string                                   | cloud-init 脚本                                                                                                                      | [#cloud-config\nchpasswd:\n  list:     | \n    root:$6$rounds=5000$ustacksalt$vIwDyL8r188dP2Po3AbH3ue.7RzxRFI2H7riDwXEnOS4NiCxrgkGtJkR9hgfhcf3shIG6l3HSxfwN1WQA8G9I0\n  expire: False\n] |
| **IsFromISO**                     | bool                                     | 是否由ISO镜像创建而来                                                                                                                | [false]                                |
| **BootloaderType**                | string                                   | 虚拟机引导方式，引导方式不匹配可能导致虚拟机无法正常工作，取值范围：1、bios 2、uefi                                                  | [bios]                                 |
| **FlatNetworkID**                 | string                                   | 扁平网络ID                                                                                                                           | [flatnet-abc123]                       |
| **FlatNetworkName**               | string                                   | 扁平网络名称                                                                                                                         | [faltenet-name]                        |
| **FlatNetwork**                   | string                                   | 扁平网络网段                                                                                                                         | [192.168.100.0/24]                     |
| **IPv4OutputName**                | string                                   | ipv4 出口名                                                                                                                          | [eip-6jq05uutty43dd]                   |
| **IPv6OutputName**                | string                                   | ipv6 出口名                                                                                                                          | [eip-4125uutty43dd]                    |
| **CPUModel**                      | string                                   | CPUModel，CPU 模型类型，仅在 CPUMode 为 custom 时生效                                                                                | [Cascadelake-Server-noTSX]             |
| **CPUModelInSetIntersection**     | bool                                     | CPUModel是否为集群通用                                                                                                               | [true]                                 |
| **CPUModelSpec**                  | string                                   | 期望的CPUModel                                                                                                                       | [Cascadelake-Server-noTSX]             |
| **CPUModelInSetIntersectionSpec** | bool                                     | 期望的CPUModel是否为集群通用                                                                                                         | [false]                                |
| **HostUUID**                      | string                                   | 宿主机UUID                                                                                                                           | [4c4c4544-005a-5210-8050-b1c04f4d3633] |
| **UUID**                          | string                                   | 实例UUID                                                                                                                             | [20afa183-5cb6-ef9f-a1b5-cda5f7d8db97] |
| **VMSPInfos**                     | array<[*VMSPInfo*](#VMSPInfo)>           | 快照信息详情信息                                                                                                                     | []                                     |
| **AttachedUSBIDs**                | array<string>                            | 绑定usbID列表，多个元素之间用,分隔。                                                                                                 | [usb-dqdqwfe213]                       |
| **ISOTotalSpec**                  | integer                                  | 当前设置的ISO插槽数                                                                                                                  | [1]                                    |
| **ISOTotalStatus**                | integer                                  | 当前真实的ISO插槽数                                                                                                                  | [1]                                    |
| **DiskCacheMode**                 | string                                   | 磁盘缓存模式，虚拟机重启、创建或热插拔时生效，取值范围：1、directsync（默认）2、none 3、writeback                                    | [directsync]                           |
| **CanMigrateAbort**               | bool                                     | 是否可以取消迁移                                                                                                                     | [false]                                |
| **VCPUBindingInfos**              | array<[*VCPUBinding*](#VCPUBinding)>     | VCPU绑定信息                                                                                                                         | []                                     |
| **VCPUBindingDegrandable**        | bool                                     | VCPU绑定在宕机迁移时是否允许降级                                                                                                     | [false]                                |
| **VCPUBindingNode**               | string                                   | VCPU绑定的物理节点                                                                                                                   | [node-123456]                          |
| **GPUBindingEnabled**             | bool                                     | 是否已开启gpu绑定                                                                                                                    | [false]                                |
| **VMNumaInfos**                   | array<[*VMNumaInfo*](#VMNumaInfo)>       | 虚拟机Numa信息                                                                                                                       | []                                     |
| **PendingChanges**                | array<[*PendingChange*](#PendingChange)> | 待生效的配置                                                                                                                         | []                                     |

#### VMNumaInfo

| 字段名         | 类型                                 | 描述信息                              | 示例      |
| :------------- | :----------------------------------- | :------------------------------------ | :-------- |
| **NumaID**     | integer                              | 虚拟机内部的NUMA节点ID                | [0]       |
| **HostNumaID** | integer                              | 该虚拟NUMA节点对应的宿主机NUMA节点ID  | [1]       |
| **Memory**     | integer                              | 分配给该NUMA节点的内存大小(单位：KiB) | [4194304] |
| **VCPUSet**    | string                               | 绑定的CPU集合                         | [0-1]     |
| **GPUInfos**   | array<[*NumaGPUInfo*](#NumaGPUInfo)> | 该NUMA节点绑定的GPU信息列表           | []        |

#### VMSPInfo

| 字段名             | 类型    | 描述信息                      | 示例                  |
| :----------------- | :------ | :---------------------------- | :-------------------- |
| **SPName**         | string  | 快照名称                      | [test-disk]           |
| **SPID**           | string  | 快照ID                        | [vmsp-o332raaryvia5l] |
| **Remark**         | string  | 备注                          | [备注一]              |
| **SPStatus**       | string  | 快照状态                      | [ReadyVMSnapshot]     |
| **SPSize**         | integer | 快照大小                      | []                    |
| **CreateTime**     | integer | 创建时间，Unix 时间戳（秒级） | [1754644992]          |
| **IsCurrent**      | bool    | 激活状态                      | [true]                |
| **WithMemory**     | bool    | 是否包含内存数据              | [true]                |
| **WithDisk**       | bool    | 是否包含磁盘                  | [true]                |
| **CPU**            | integer | CPU核数                       | [1]                   |
| **Memory**         | integer | 内存大小，单位：MiB           | [1024]                |
| **ComputeSetType** | string  | 计算集群类型                  | [ComputeSetTTPZ]      |

#### VmDiskInfo

| 字段名                 | 类型    | 描述信息                                 | 示例                     |
| :--------------------- | :------ | :--------------------------------------- | :----------------------- |
| **DiskID**             | string  | 虚拟机绑定的磁盘ID                       | [vm-0319mftrkgiytq-boot] |
| **Name**               | string  | 名称                                     | [disk-vip4]              |
| **Drive**              | string  | 设备名                                   | [vda]                    |
| **IsElastic**          | string  | 是否为弹性设备                           | [N Y]                    |
| **Size**               | integer | 磁盘大小，单位：GiB                      | [40]                     |
| **StorageSetType**     | string  | 存储集群id                               | [StorageSetceph]         |
| **Type**               | string  | 磁盘类型：1、启动盘(boot)2、数据盘(data) | [boot]                   |
| **ShareAble**          | bool    | 是否可共享                               | [false]                  |
| **IsSharedblock**      | bool    | 是否是外置存储盘                         | [true]                   |
| **StorageSetArch**     | string  |                                          | [HDD SSD]                |
| **StorageSetProvider** | string  | 存储集群制备器                           | [rbd ustorbs cjfs]       |
| **SnapshotCount**      | integer | 快照数量                                 | [1]                      |
| **Encrypted**          | bool    | 是否加密                                 | [false]                  |

## 删除虚拟机 - DeleteVMInstance

### 请求参数

| 字段名              | 类型    | 描述信息                                                   | 示例               | 参数来源             | 必填    |
| :------------------ | :------ | :--------------------------------------------------------- | :----------------- | :------------------- | :------ |
| **Region**          | string  | 地域ID，地域，一批可共享的物理资源使用集合                 | [develop]          | [DescribeRegion]     | **Yes** |
| **CompanyID**       | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236]        | [DescribeUser]       | No      |
| **DelBindResource** | integer | 是否删除虚拟机绑定资源，0表示否（默认），1表示是           | [1]                | []                   | No      |
| **VMID**            | string  | 虚拟机ID                                                   | [vm-fe32freg35h2h] | [DescribeVMInstance] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
| :----- | :--- | :------- | :--- |


## 查询CI状态 - DescribeCIStatus

### 请求参数

| 字段名           | 类型          | 描述信息                                                    | 示例                     | 参数来源             | 必填    |
| :--------------- | :------------ | :---------------------------------------------------------- | :----------------------- | :------------------- | :------ |
| **Region**       | string        | 地域ID，地域，一批可共享的物理资源使用集合                  | [develop]                | [DescribeRegion]     | **Yes** |
| **CompanyID**    | integer       | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离  | [200000236]              | [DescribeUser]       | No      |
| **ResourceType** | string        | 资源类型                                                    | [VM]                     | []                   | **Yes** |
| **ProjectIDs**   | array<string> | 项目组ID列表，用于资源分组管理，多个元素之间用,分隔。       | [project-5kyt09jzwcqh6j] | [ListProjects]       | No      |
| **States**       | array<string> | 状态                                                        | [Availabile]             | []                   | No      |
| **ResourceIDs**  | array<string> | 虚拟机资源ID列表，用于查询指定的资源，多个元素之间用,分隔。 | [vm-fe32freg35h2h]       | [DescribeVMInstance] | **Yes** |

### 响应字段

| 字段名         | 类型                                   | 描述信息         | 示例 |
| :------------- | :------------------------------------- | :--------------- | :--- |
| **TotalCount** | integer                                | 查询到的总记录数 | []   |
| **Infos**      | array<[*CIStatusInfo*](#CIStatusInfo)> | CI状态详情       | []   |

### 数据模型

#### CIStatusInfo

| 字段名         | 类型   | 描述信息     | 示例               |
| :------------- | :----- | :----------- | :----------------- |
| **ResourceID** | string | 虚拟机资源ID | [vm-fe32freg35h2h] |
| **CIID**       | string | CIID         | [vm-fe32freg35h2h] |
| **CIStatus**   | string | CI状态       | [Running]          |

## 获取修改配置后的差价 - GetPaymentOfPremium

### 请求参数

| 字段名          | 类型    | 描述信息                                                   | 示例                | 参数来源             | 必填    |
| :-------------- | :------ | :--------------------------------------------------------- | :------------------ | :------------------- | :------ |
| **Region**      | string  | 地域ID，地域，一批可共享的物理资源使用集合                 | [develop]           | [DescribeRegion]     | **Yes** |
| **CompanyID**   | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236]         | [DescribeUser]       | No      |
| **VMID**        | string  | 虚拟机ID                                                   | [vm-29by1ohfdf2f41] | [DescribeVMInstance] | **Yes** |
| **CPU**         | integer | CPU核数                                                    | [1]                 | []                   | **Yes** |
| **Memory**      | integer | 内存大小，单位：MiB，必须是 1024 的倍数                    | [1024]              | []                   | **Yes** |
| **GPUType**     | string  | GPU 类型，取值范围：1、GPU  2、VGPU                        | [GPU]               | []                   | No      |
| **GPUMdevName** | string  | GPU规格名称，GPUType=GPU时指定                             | [nvidia-2080ti]     | [ListGPUs]           | No      |
| **GPU**         | integer | GPU的颗数，GPUType=GPU时指定                               | [2]                 | []                   | No      |
| **MdevName**    | string  | vGPU的规格名称，GPUType=VGPU时指定                         | [nvidia-2080ti]     | [ListVGPUs]          | No      |

### 响应字段

| 字段名        | 类型   | 描述信息                                    | 示例     |
| :------------ | :----- | :------------------------------------------ | :------- |
| **Price**     | double | 价格，单位为元                              | [199.99] |
| **OrderType** | string | 订单类型，0表示无变更，3表示升级，7表示降级 | [3]      |


## 获取虚拟机价格 - GetVMInstancePrice

### 请求参数

| 字段名              | 类型    | 描述信息                                                                                                          | 示例             | 参数来源             | 必填    |
| :------------------ | :------ | :---------------------------------------------------------------------------------------------------------------- | :--------------- | :------------------- | :------ |
| **Region**          | string  | 地域ID，地域，一批可共享的物理资源使用集合                                                                        | [develop]        | [DescribeRegion]     | **Yes** |
| **CompanyID**       | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离                                                        | [200000236]      | [DescribeUser]       | No      |
| **ChargeType**      | string  | 计费类型，用于指定计费模式。取值范围：1、Dynamic（按小时计费）2、Month（按月计费）3、Year（按年计费）             | []               | []                   | **Yes** |
| **Quantity**        | integer | 计费数量，指定计费周期的数量。按月/年计费时表示创建 Quantity 个月/年的磁盘，按小时计费时强制为1。取值范围：正整数 | []               | []                   | **Yes** |
| **Count**           | integer | 购买数量，需为大于1的正整数                                                                                       | [2]              | []                   | **Yes** |
| **ImageID**         | string  | 镜像ID，指定的镜像ID                                                                                              | [image-k8s-node] | [DescribeImage]      | No      |
| **VMType**          | string  | 计算集群ID                                                                                                        | [ComputeSetTTPZ] | [DescribeVMSet]      | **Yes** |
| **BootDiskSpace**   | integer | 系统盘大小，单位:GB                                                                                               | [100]            | []                   | No      |
| **BootDiskSetType** | string  | 系统盘集群ID，系统盘大小不为空时必填                                                                              | [cluster-01]     | [DescribeStorageSet] | No      |
| **StorageType**     | string  | 存储类型                                                                                                          | [LocalDisk]      | []                   | No      |
| **CPU**             | integer | CPU核数                                                                                                           | [1]              | []                   | **Yes** |
| **Memory**          | integer | 内存大小，单位：MiB，1024 的倍数                                                                                  | [1024]           | []                   | **Yes** |
| **GPUType**         | string  | GPU 类型，取值范围：1、GPU  2、VGPU                                                                               | [GPU]            | []                   | No      |
| **GPUMdevName**     | string  | GPU规格名称，GPUType=GPU时指定                                                                                    | [nvidia-2080ti]  | [ListGPUs]           | No      |
| **GPU**             | integer | GPU的颗数，GPUType=GPU时指定                                                                                      | [2]              | []                   | No      |
| **MdevName**        | string  | vGPU的规格名称，GPUType=VGPU时指定                                                                                | [nvidia-2080ti]  | [ListVGPUs]          | No      |

### 响应字段

| 字段名    | 类型                                 | 描述信息           | 示例 |
| :-------- | :----------------------------------- | :----------------- | :--- |
| **Infos** | array<[*VmPriceInfo*](#VmPriceInfo)> | 虚拟机实例价格详情 | []   |

### 数据模型

#### VmPriceInfo

| 字段名         | 类型   | 描述信息                                                                            | 示例      |
| :------------- | :----- | :---------------------------------------------------------------------------------- | :-------- |
| **Price**      | double | 价格，单位：元                                                                      | [199.99]  |
| **ChargeType** | string | 计费类型，取值范围：1、Dynamic（按小时计费）2、Month（按月计费）3、Year（按年计费） | [Dynamic] |

## 获取Spice信息 - GetVMSpiceInfo

### 请求参数

| 字段名        | 类型    | 描述信息                                                   | 示例                | 参数来源             | 必填    |
| :------------ | :------ | :--------------------------------------------------------- | :------------------ | :------------------- | :------ |
| **Region**    | string  | 地域ID，地域，一批可共享的物理资源使用集合                 | [develop]           | [DescribeRegion]     | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236]         | [DescribeUser]       | No      |
| **VMID**      | string  | 虚拟机ID                                                   | [vm-29by1ohfdf2f41] | [DescribeVMInstance] | **Yes** |

### 响应字段

| 字段名   | 类型   | 描述信息  | 示例 |
| :------- | :----- | :-------- | :--- |
| **Info** | object | Spice信息 | []   |

### 数据模型

#### SpiceInfo

| 字段名            | 类型    | 描述信息  | 示例                |
| :---------------- | :------ | :-------- | :------------------ |
| **VMID**          | string  | 虚拟机ID  | [vm-29by1ohfdf2f41] |
| **SpiceIP**       | string  | SpiceIP   | [192.168.178.41]    |
| **SpicePort**     | integer | Spice端口 | [31241]             |
| **SpicePassword** | string  | Spice密码 | [eqe1f31]           |

## 获取VNC信息 - GetVMVNCInfo

### 请求参数

| 字段名        | 类型    | 描述信息                                                   | 示例                | 参数来源             | 必填    |
| :------------ | :------ | :--------------------------------------------------------- | :------------------ | :------------------- | :------ |
| **Region**    | string  | 地域ID，地域，一批可共享的物理资源使用集合                 | [develop]           | [DescribeRegion]     | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236]         | [DescribeUser]       | No      |
| **VMID**      | string  | 虚拟机ID                                                   | [vm-29by1ohfdf2f41] | [DescribeVMInstance] | **Yes** |

### 响应字段

| 字段名   | 类型   | 描述信息 | 示例 |
| :------- | :----- | :------- | :--- |
| **Info** | object | VNC信息  | []   |

### 数据模型

#### VNCInfo

| 字段名          | 类型    | 描述信息 | 示例                |
| :-------------- | :------ | :------- | :------------------ |
| **VMID**        | string  | 虚拟机ID | [vm-29by1ohfdf2f41] |
| **VNCIP**       | string  | VNCIP    | [192.168.178.41]    |
| **VNCPort**     | integer | VNC端口  | [31241]             |
| **VNCPassword** | string  | VNC密码  | [eqe1f31]           |

## 断电主机 - PoweroffVMInstance

### 请求参数

| 字段名        | 类型    | 描述信息                                                   | 示例                | 参数来源             | 必填    |
| :------------ | :------ | :--------------------------------------------------------- | :------------------ | :------------------- | :------ |
| **Region**    | string  | 地域ID，地域，一批可共享的物理资源使用集合                 | [develop]           | [DescribeRegion]     | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236]         | [DescribeUser]       | No      |
| **VMID**      | string  | 虚拟机ID                                                   | [vm-29by1ohfdf2f41] | [DescribeVMInstance] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
| :----- | :--- | :------- | :--- |


## 重装系统 - ReinstallVMInstance

### 请求参数

| 字段名        | 类型    | 描述信息                                                                                                                               | 示例                             | 参数来源             | 必填    |
| :------------ | :------ | :------------------------------------------------------------------------------------------------------------------------------------- | :------------------------------- | :------------------- | :------ |
| **Region**    | string  | 地域ID，地域，一批可共享的物理资源使用集合                                                                                             | [develop]                        | [DescribeRegion]     | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离                                                                             | [200000236]                      | [DescribeUser]       | No      |
| **VMID**      | string  | 虚拟机ID                                                                                                                               | [vm-29by1ohfdf2f41]              | [DescribeVMInstance] | **Yes** |
| **ImageID**   | string  | 镜像ID                                                                                                                                 | [img-dq123fasd33]                | [DescribeImage]      | **Yes** |
| **Password**  | string  | 新的管理员密码，要求：1、密码长度限6-30个字符 2、需要同时包含小写字母、特殊字符(除空格) 3、不能包含非法字符                            | [f34f*2defrefe]                  | []                   | No      |
| **UserData**  | string  | cloud-init脚本，需要镜像支持，支持自定义初始化配置。自定义初始化脚本，经过base64编码，最大1M，可在初次启动和每次开机/重装/重启时执行。 | [#cloud-config\nhostname: my-vm] | []                   | No      |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
| :----- | :--- | :------- | :--- |


## 重置主机密码 - ResetVMInstancePassword

### 请求参数

| 字段名        | 类型    | 描述信息                                                                                                | 示例                | 参数来源             | 必填    |
| :------------ | :------ | :------------------------------------------------------------------------------------------------------ | :------------------ | :------------------- | :------ |
| **Region**    | string  | 地域ID，地域，一批可共享的物理资源使用集合                                                              | [develop]           | [DescribeRegion]     | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离                                              | [200000236]         | [DescribeUser]       | No      |
| **VMID**      | string  | 虚拟机ID                                                                                                | [vm-29by1ohfdf2f41] | [DescribeVMInstance] | **Yes** |
| **Password**  | string  | 管理员密码，要求：1、密码长度限6-30个字符 2、需要同时包含小写字母、特殊字符(除空格) 3、不能包含非法字符 | [f34f*2defrefe]     | []                   | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
| :----- | :--- | :------- | :--- |


## 修改虚拟机配置 - ResizeVMConfig

### 请求参数

| 字段名                | 类型    | 描述信息                                                       | 示例                     | 参数来源             | 必填    |
| :-------------------- | :------ | :------------------------------------------------------------- | :----------------------- | :------------------- | :------ |
| **Region**            | string  | 地域ID，地域，一批可共享的物理资源使用集合                     | [develop]                | [DescribeRegion]     | **Yes** |
| **CompanyID**         | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离     | [200000236]              | [DescribeUser]       | No      |
| **ApplicationName**   | string  | 审批名称，用于审批流程的标题描述。当启用审批流程时使用。       | [虚拟机资源申请]         | []                   | No      |
| **ApplicationReason** | string  | 审批理由，用于说明申请创建虚拟机的原因。当启用审批流程时使用。 | [用于新项目测试环境搭建] | []                   | No      |
| **VMID**              | string  | 虚拟机ID                                                       | [vm-29by1ohfdf2f41]      | [DescribeVMInstance] | **Yes** |
| **CPU**               | integer | CPU核数                                                        | [1]                      | []                   | **Yes** |
| **Memory**            | integer | 内存大小，单位：MiB，1024 的倍数                               | []                       | []                   | **Yes** |
| **GPUType**           | string  | GPU 类型，取值范围：1、GPU  2、VGPU                            | [GPU]                    | []                   | No      |
| **GPUMdevName**       | string  | GPU规格名称，GPUType=GPU时指定                                 | [nvidia-2080ti]          | [ListGPUs]           | No      |
| **GPU**               | integer | GPU的颗数，GPUType=GPU时指定                                   | [2]                      | []                   | No      |
| **MdevName**          | string  | vGPU的规格名称，GPUType=VGPU时指定                             | [nvidia-2080ti]          | [ListVGPUs]          | No      |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
| :----- | :--- | :------- | :--- |


## 重启主机 - RestartVMInstance

### 请求参数

| 字段名        | 类型    | 描述信息                                                   | 示例                | 参数来源             | 必填    |
| :------------ | :------ | :--------------------------------------------------------- | :------------------ | :------------------- | :------ |
| **Region**    | string  | 地域ID，地域，一批可共享的物理资源使用集合                 | [develop]           | [DescribeRegion]     | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236]         | [DescribeUser]       | No      |
| **VMID**      | string  | 虚拟机ID                                                   | [vm-29by1ohfdf2f41] | [DescribeVMInstance] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
| :----- | :--- | :------- | :--- |


## 启动主机 - StartVMInstance

### 请求参数

| 字段名        | 类型    | 描述信息                                                   | 示例                | 参数来源             | 必填    |
| :------------ | :------ | :--------------------------------------------------------- | :------------------ | :------------------- | :------ |
| **Region**    | string  | 地域ID，地域，一批可共享的物理资源使用集合                 | [develop]           | [DescribeRegion]     | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236]         | [DescribeUser]       | No      |
| **VMID**      | string  | 虚拟机ID                                                   | [vm-29by1ohfdf2f41] | [DescribeVMInstance] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
| :----- | :--- | :------- | :--- |


## 关闭主机 - StopVMInstance

### 请求参数

| 字段名        | 类型    | 描述信息                                                   | 示例                | 参数来源             | 必填    |
| :------------ | :------ | :--------------------------------------------------------- | :------------------ | :------------------- | :------ |
| **Region**    | string  | 地域ID，地域，一批可共享的物理资源使用集合                 | [develop]           | [DescribeRegion]     | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236]         | [DescribeUser]       | No      |
| **VMID**      | string  | 虚拟机ID                                                   | [vm-29by1ohfdf2f41] | [DescribeVMInstance] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
| :----- | :--- | :------- | :--- |


## 设置虚拟机出口 - UpdateVMDefaultGW

### 请求参数

| 字段名        | 类型    | 描述信息                                                   | 示例                | 参数来源                               | 必填    |
| :------------ | :------ | :--------------------------------------------------------- | :------------------ | :------------------------------------- | :------ |
| **Region**    | string  | 地域ID，地域，一批可共享的物理资源使用集合                 | [develop]           | [DescribeRegion]                       | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236]         | [DescribeUser]                         | No      |
| **VMID**      | string  | 虚拟机ID                                                   | [vm-29by1ohfdf2f41] | [DescribeVMInstance]                   | **Yes** |
| **EIPID**     | string  | 外网资源ID，eipID 和 ipID 指定一个即可                     | [eip-789def]        | [DescribeEIP]                          | No      |
| **IPID**      | string  | ip资源ID，eipID 和 ipID 指定一个即可                       | [ip-789def]         | [DescribeEIP ListAllocatedIPsInSubnet] | No      |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
| :----- | :--- | :------- | :--- |


## 申请VNC会话 - AllocateVMVNCSession

### 请求参数

| 字段名     | 类型   | 描述信息                                   | 示例                | 参数来源             | 必填    |
| :--------- | :----- | :----------------------------------------- | :------------------ | :------------------- | :------ |
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop]           | [DescribeRegion]     | **Yes** |
| **VMID**   | string | 虚拟机ID                                   | [vm-29by1ohfdf2f41] | [DescribeVMInstance] | **Yes** |

### 响应字段

| 字段名        | 类型   | 描述信息       | 示例                   |
| :------------ | :----- | :------------- | :--------------------- |
| **SessionID** | string | VNC Session ID | [tssid-8goqdy5l1l3t8p] |


## 设置虚拟机高级参数 - UpdateVMAdvancedOptions

### 请求参数

| 字段名               | 类型    | 描述信息                                                                                                                                                                                                 | 示例                             | 参数来源             | 必填    |
| :------------------- | :------ | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | :------------------------------- | :------------------- | :------ |
| **Region**           | string  | 地域ID，地域，一批可共享的物理资源使用集合                                                                                                                                                               | [develop]                        | [DescribeRegion]     | **Yes** |
| **VMID**             | string  | 虚拟机ID                                                                                                                                                                                                 | [vm-29by1ohfdf2f41]              | [DescribeVMInstance] | **Yes** |
| **DNS**              | string  | DNS                                                                                                                                                                                                      | [114.114.114.114]                | []                   | No      |
| **HighAvailability** | string  | 虚拟机高可用模式，取值范围：1、NeverStop（默认）2、None                                                                                                                                                  | [NeverStop]                      | []                   | No      |
| **CPUMode**          | string  | CPU 启动模式，取值范围：1、host-passthrough（直通模式）2、custom（自定义）                                                                                                                               | [host-passthrough]               | []                   | No      |
| **UserData**         | string  | cloud-init脚本，需要镜像支持，支持自定义初始化配置                                                                                                                                                       | [#cloud-config\nhostname: my-vm] | []                   | No      |
| **UninstallISO**     | bool    | 是否卸载ISO镜像，true表示卸载，false表示保留                                                                                                                                                             | [true]                           | []                   | No      |
| **BootloaderType**   | string  | 虚拟机引导方式，引导方式不匹配可能导致虚拟机无法正常工作，取值范围：bios、uefi                                                                                                                           | [uefi]                           | []                   | No      |
| **CPUModel**         | string  | CPU 模型类型，仅在 CPUMode 为 custom 时生效，取值范围：default、general、other                                                                                                                           | [default]                        | []                   | No      |
| **ISOTotal**         | integer | ISO插槽数量，修改成功后需关机/断电再重新开机才能生效。                                                                                                                                                   | [1]                              | []                   | No      |
| **DiskCacheMode**    | string  | 磁盘缓存模式，取值范围：1、writeback 2、none 3、directsync。writeback性能高但安全性中低，适合虚拟机和SSD存储；none性能与安全性均衡，适合NVMe和I/O密集型应用；directsync安全性最高，适合数据库和金融交易. | [writeback]                      | []                   | No      |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
| :----- | :--- | :------- | :--- |


## 修改网卡的MAC - UpdateVMMAC

### 请求参数

| 字段名      | 类型   | 描述信息                                                            | 示例                | 参数来源         | 必填    |
| :---------- | :----- | :------------------------------------------------------------------ | :------------------ | :--------------- | :------ |
| **Region**  | string | 地域ID，地域，一批可共享的物理资源使用集合                          | [develop]           | [DescribeRegion] | **Yes** |
| **VMID**    | string | 虚拟机ID                                                            | [vm-j40jm6vb8lo321] | []               | **Yes** |
| **NICType** | string | 网卡类型，取值范围：1、LAN(局域网) 2、WAN(广域网) 3、Flat(扁平网络) | [LAN]               | []               | **Yes** |
| **MAC**     | string | MAC地址，格式为6组2位十六进制数，以冒号或连字符分隔                 | [00:1A:2B:3C:4D:5E] | []               | No      |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
| :----- | :--- | :------- | :--- |


## 设置虚拟机从Cdrom启动 - SetBootFromCdrom

### 请求参数

| 字段名        | 类型    | 描述信息                                                   | 示例                       | 参数来源             | 必填    |
| :------------ | :------ | :--------------------------------------------------------- | :------------------------- | :------------------- | :------ |
| **Region**    | string  | 地域ID，地域，一批可共享的物理资源使用集合                 | [develop]                  | [DescribeRegion]     | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236]                | [DescribeUser]       | No      |
| **VMID**      | string  | 虚拟机ID                                                   | [vm-29by1ohfdf2f41]        | [DescribeVMInstance] | **Yes** |
| **ImageID**   | string  | 镜像ID，用于指定镜像                                       | [image-windows-virtio-iso] | [DescribeImage]      | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
| :----- | :--- | :------- | :--- |


## 设置虚拟机不从Cdrom启动 - UnSetBootFromCdrom

### 请求参数

| 字段名        | 类型    | 描述信息                                                   | 示例                       | 参数来源             | 必填    |
| :------------ | :------ | :--------------------------------------------------------- | :------------------------- | :------------------- | :------ |
| **Region**    | string  | 地域ID，地域，一批可共享的物理资源使用集合                 | [develop]                  | [DescribeRegion]     | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236]                | [DescribeUser]       | No      |
| **VMID**      | string  | 虚拟机ID                                                   | [vm-29by1ohfdf2f41]        | [DescribeVMInstance] | **Yes** |
| **ImageID**   | string  | 镜像ID，用于指定镜像                                       | [image-windows-virtio-iso] | [DescribeImage]      | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
| :----- | :--- | :------- | :--- |


## 虚拟机热存储迁移 - MigrateVMStorage

### 请求参数

| 字段名         | 类型          | 描述信息                                                                                                          | 示例                | 参数来源                                                                                                           | 必填                 |
| :------------- | :------------ | :---------------------------------------------------------------------------------------------------------------- | :------------------ | :----------------------------------------------------------------------------------------------------------------- | :------------------- |
| **Region**     | string        | 地域ID，地域，一批可共享的物理资源使用集合                                                                        | [develop]           | [DescribeRegion]                                                                                                   | **Yes**              |
| **CompanyID**  | integer       | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离                                                        | [200000236]         | [DescribeUser]                                                                                                     | **Yes**              |
| **VMID**       | string        | 虚拟机ID                                                                                                          | [vm-j40jm6vb8lo321] | []                                                                                                                 | **Yes**              |
| **Infos**      | array<string> | 磁盘迁移信息，规则: 原盘id                                                                                        | 目标集群id          | 目标盘id；若为 ceph -> ceph 不需要指定目标盘id；若为 外置盘 -> ceph 可以不指定目标盘id，会自动创建相同容量的ceph盘 | [disk-zsgvd1z7oy8yws | sharedblock-ox58hkcxn4av6a | disk-x4oerm9r9n06ws disk-zsgvd1z7oy8yws | sharedblock-ox58hkcxn4av6a | ] | [] | **Yes** |
| **ChargeType** | string        | 计费类型，用于指定计费模式。取值范围：1、Dynamic（按小时计费）2、Month（按月计费）3、Year（按年计费）             | []                  | []                                                                                                                 | No                   |
| **Quantity**   | integer       | 计费数量，指定计费周期的数量。按月/年计费时表示创建 Quantity 个月/年的磁盘，按小时计费时强制为1。取值范围：正整数 | []                  | []                                                                                                                 | No                   |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
| :----- | :--- | :------- | :--- |


## 虚拟机整机快照 - SaveVMInstance

### 请求参数

| 字段名            | 类型    | 描述信息                                                                                               | 示例                | 参数来源             | 必填    |
| :---------------- | :------ | :----------------------------------------------------------------------------------------------------- | :------------------ | :------------------- | :------ |
| **Region**        | string  | 地域ID，地域，一批可共享的物理资源使用集合                                                             | [develop]           | [DescribeRegion]     | **Yes** |
| **CompanyID**     | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离                                             | [200000236]         | [DescribeUser]       | No      |
| **SPName**        | string  | 虚拟机快照名称，长度为1-50个字符，名称只能包含中英文、数字、点（.），下划线（_）和中划线（-）          | [vm-snapshot-01]    | []                   | No      |
| **Remark**        | string  | 备注，用于进行说明和注释，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [这是一个备注]      | []                   | No      |
| **VMID**          | string  | 虚拟机ID                                                                                               | [vm-29by1ohfdf2f41] | [DescribeVMInstance] | **Yes** |
| **WithoutMemory** | bool    | 是否不包含内存快照，暂存虚拟机时，WithoutMemory 为 false                                               | [false]             | []                   | No      |
| **WithoutDisk**   | bool    | 是否不包含磁盘，WithoutDisk 为 true 时：暂存虚拟机。WithoutDisk 为 false 时：虚拟机快照                | [false]             | []                   | No      |

### 响应字段

| 字段名   | 类型   | 描述信息 | 示例                  |
| :------- | :----- | :------- | :-------------------- |
| **SPID** | string | 快照ID   | [vmsp-yhzksg7tg2zjnq] |


## 虚拟机恢复快照 - RestoreVMInstance

### 请求参数

| 字段名             | 类型    | 描述信息                                                   | 示例                  | 参数来源             | 必填    |
| :----------------- | :------ | :--------------------------------------------------------- | :-------------------- | :------------------- | :------ |
| **Region**         | string  | 地域ID，地域，一批可共享的物理资源使用集合                 | [develop]             | [DescribeRegion]     | **Yes** |
| **CompanyID**      | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236]           | [DescribeUser]       | No      |
| **VMID**           | string  | 虚拟机ID                                                   | [vm-29by1ohfdf2f41]   | [DescribeVMInstance] | **Yes** |
| **SPID**           | string  | 快照ID，从快照恢复时传入，从暂存恢复时不需要传入。         | [vmsp-yhzksg7tg2zjnq] | [DescribeSnapshot]   | No      |
| **AllowOtherHost** | bool    | 是否允许调度到其他节点，调度到快照时节点以外的节点         | [false]               | []                   | No      |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
| :----- | :--- | :------- | :--- |


## 虚拟机删除快照 - DeleteVMSnapshot

### 请求参数

| 字段名        | 类型    | 描述信息                                                   | 示例                  | 参数来源             | 必填    |
| :------------ | :------ | :--------------------------------------------------------- | :-------------------- | :------------------- | :------ |
| **Region**    | string  | 地域ID，地域，一批可共享的物理资源使用集合                 | [develop]             | [DescribeRegion]     | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236]           | [DescribeUser]       | No      |
| **VMID**      | string  | 虚拟机ID                                                   | [vm-29by1ohfdf2f41]   | [DescribeVMInstance] | **Yes** |
| **SPID**      | string  | 快照ID，当不清空时必须提供，用于删除指定的快照             | [vmsp-48e2d2avkck3mc] | [DescribeSnapshot]   | No      |
| **IsClear**   | bool    | 是否清空快照                                               | [false]               | []                   | No      |

### 响应字段

| 字段名    | 类型          | 描述信息             | 示例                  |
| :-------- | :------------ | :------------------- | :-------------------- |
| **SPIDs** | array<string> | 被删除的快照 ID 列表 | [vmsp-48e2d2avkck3mc] |


## 创建 vmware 虚拟机控制台凭证 - GenerateVMWareConsoleTicket

### 请求参数

| 字段名         | 类型   | 描述信息        | 示例                | 参数来源             | 必填    |
| :------------- | :----- | :-------------- | :------------------ | :------------------- | :------ |
| **ConfigName** | string | vmware 配置名称 | [config-test]       | []                   | **Yes** |
| **VMName**     | string | 虚拟机名称      | [vm-test]           | [DescribeVMInstance] | **Yes** |
| **Datacenter** | string | 数据中心名称    | [Datacenter-test]   | []                   | **Yes** |
| **VMID**       | string | 虚拟机ID        | [vm-j40jm6vb8lo321] | [DescribeVMInstance] | **Yes** |

### 响应字段

| 字段名   | 类型   | 描述信息             | 示例                                                            |
| :------- | :----- | :------------------- | :-------------------------------------------------------------- |
| **Link** | string | VMware控制台访问链接 | [https://vcenter.example.com/console?vmId=vm-123&ticket=abc123] |


## 获取 vmware 计算集群的信息 - GetVMWareClusterDatastore

### 请求参数

| 字段名          | 类型   | 描述信息        | 示例                      | 参数来源        | 必填    |
| :-------------- | :----- | :-------------- | :------------------------ | :-------------- | :------ |
| **ConfigName**  | string | vmware 配置名称 | [config-test]             | []              | **Yes** |
| **Datacenter**  | string | 数据中心名称    | [config-test]             | []              | **Yes** |
| **ClusterName** | string | 集群集群名称    | [cluster-test]            | [DescribeVMSet] | **Yes** |
| **ClusterID**   | string | 集群集群ID      | [cluster-eqwgr1rt54345r3] | [DescribeVMSet] | **Yes** |

### 响应字段

| 字段名    | 类型                                         | 描述信息                 | 示例 |
| :-------- | :------------------------------------------- | :----------------------- | :--- |
| **Infos** | array<[*VMwareDatastore*](#VMwareDatastore)> | VMwareDatastore 详情信息 | []   |

### 数据模型

#### VMwareDatastore

| 字段名          | 类型   | 描述信息       | 示例            |
| :-------------- | :----- | :------------- | :-------------- |
| **Name**        | string | 存储名称       | [datastore1]    |
| **DatastoreID** | string | 存储唯一标识符 | [datastore-123] |

## 获取 vmware 虚拟机信息 - DescribeVMWareVMs

### 请求参数

| 字段名         | 类型   | 描述信息                       | 示例                    | 参数来源        | 必填    |
| :------------- | :----- | :----------------------------- | :---------------------- | :-------------- | :------ |
| **ConfigName** | string | vmware 配置名称                | [vmware-config-name]    | []              | **Yes** |
| **Datacenter** | string | 数据中心ID                     | [datacenter-123efesdfr] | []              | **Yes** |
| **Cluster**    | string | 计算集群ID，虚拟机所在计算集群 | [ComputeSetTTPZ]        | [DescribeVMSet] | No      |

### 响应字段

| 字段名         | 类型                           | 描述信息         | 示例 |
| :------------- | :----------------------------- | :--------------- | :--- |
| **TotalCount** | integer                        | 查询到的总记录数 | []   |
| **Infos**      | array<[*VMWareVM*](#VMWareVM)> | VMWareVM 信息    | []   |

### 数据模型

#### VMWareVM

| 字段名            | 类型    | 描述信息        | 示例                                |
| :---------------- | :------ | :-------------- | :---------------------------------- |
| **VMID**          | string  | 虚拟机ID        | [vm-29by1ohfdf2f41]                 |
| **Name**          | string  | 虚拟机名称      | [vm-test]                           |
| **InventoryPath** | string  | 存储库路径      | [/Datacenter/vm/Production/vm-test] |
| **PowerState**    | string  | 电源状态        | [poweredOn]                         |
| **CPU**           | integer | 虚拟机CPU核心数 | [4]                                 |
| **Memory**        | integer | 内存大小(MB)    | [8192]                              |

## 整机克隆 - CloneVMInstance

### 请求参数

| 字段名                     | 类型    | 描述信息                                                                                                          | 示例                     | 参数来源                | 必填    |
| :------------------------- | :------ | :---------------------------------------------------------------------------------------------------------------- | :----------------------- | :---------------------- | :------ |
| **Region**                 | string  | 地域ID，地域，一批可共享的物理资源使用集合                                                                        | [develop]                | [DescribeRegion]        | **Yes** |
| **ApplicationName**        | string  | 审批名称，用于审批流程的标题描述。当启用审批流程时使用。                                                          | [虚拟机资源申请]         | []                      | No      |
| **ApplicationReason**      | string  | 审批理由，用于说明申请创建虚拟机的原因。当启用审批流程时使用。                                                    | [用于新项目测试环境搭建] | []                      | No      |
| **CompanyID**              | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离                                                        | [200000236]              | [DescribeUser]          | **Yes** |
| **Email**                  | string  | 租户邮箱，租户注册时提供的用于接收通知和联系的邮箱                                                                | [example@ucloud.cn]      | [DescribeUser]          | No      |
| **ProjectID**              | string  | 项目组ID，用于资源分组管理                                                                                        | [project-5kyt09jzwcqh6j] | [ListProjects]          | No      |
| **Name**                   | string  | 虚拟机名称，长度为1-50个字符，名称只能包含中英文、数字、点（.），下划线（_）和中划线（-）                         | [vm-test]                | []                      | **Yes** |
| **Remark**                 | string  | 备注，用于进行说明和注释，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空            | [这是一个备注]           | []                      | No      |
| **CPU**                    | integer | CPU核数                                                                                                           | [1]                      | []                      | **Yes** |
| **Memory**                 | integer | 内存大小，单位：MiB，必须是 1024 的倍数                                                                           | [1024]                   | []                      | **Yes** |
| **VMID**                   | string  | 源虚拟机ID                                                                                                        | [vm-rioehgoih923]        | [DescribeVMInstance]    | **Yes** |
| **VMSPID**                 | string  | 快照ID，克隆所需要的整机快照ID                                                                                    | [vmsp-uunpn7ish05arm]    | [DescribeSnapshot]      | **Yes** |
| **VPCID**                  | string  | vpcID，创建虚拟机时必须选择 VPC 网络和所属子网，即选择虚拟要加入的网络及 IP 网段                                  | [vpc-el7miluuodq]        | [DescribeVPC]           | **Yes** |
| **SubnetID**               | string  | 子网ID，创建虚拟机时必须选择 VPC 网络和所属子网，即选择虚拟要加入的网络及 IP 网段                                 | [subnet-lq10np1d231]     | [DescribeSubnet]        | **Yes** |
| **InternalIP**             | string  | 内网IP地址，支持 IPv4 或 IPv6，留空则自动分配                                                                     | [10.0.0.1]               | []                      | No      |
| **LANSGID**                | string  | 内网安全组ID，如果需要使用，需提前创建安全组，得到安全组 ID                                                       | [sg-v55mgokwcc11c5]      | [DescribeSecurityGroup] | No      |
| **LANMAC**                 | string  | 内网/扁平网络mac，不指定时自动生成                                                                                | [00:1A:2B:3C:4D:5E]      | []                      | No      |
| **IPVersion**              | string  | IP版本，支持 IPv4 或 IPv6                                                                                         | []                       | []                      | **Yes** |
| **OperatorName**           | string  | 外网线路ID，带宽不为空时必传                                                                                      | [segment-35dehi0g8zdqqq] | []                      | No      |
| **InternetIP**             | string  | 外网 IP 地址，支持 IPv4 或 IPv6，留空则自动分配                                                                   | [192.31.31.1]            | []                      | No      |
| **Bandwidth**              | integer | IP 带宽，单位:Mb。当前 IP 地址的带宽上限，扁平网络可以不设置带宽                                                  | [1]                      | []                      | **Yes** |
| **WANSGID**                | string  | 外网安全组ID，如果需要使用，需提前创建安全组，得到安全组 ID                                                       | [sg-v55mgokwcc11c5 ]     | [DescribeSecurityGroup] | No      |
| **WANMAC**                 | string  | 外网mac，不指定时自动生成                                                                                         | [00:1A:2B:3C:4D:5E]      | []                      | No      |
| **FlatNetworkID**          | string  | 扁平网络ID， 带宽不为空时必传                                                                                     | [faltenid-dqdqg4341]     | [DescribeFlatIP]        | No      |
| **FlatNetworkSGID**        | string  | 扁平网络安全组ID                                                                                                  | [sg-flatnet-123456]      | [DescribeSecurityGroup] | No      |
| **LANInAverageBandwidth**  | integer | 入向平均带宽，第一张网卡入向平均带宽（单位：Mbps），0 表示不限制                                                  | [1000]                   | []                      | No      |
| **LANOutAverageBandwidth** | integer | 出向平均带宽，第一张网卡出向平均带宽（单位：Mbps），0 表示不限制                                                  | [1000]                   | []                      | No      |
| **ChargeType**             | string  | 计费类型，用于指定计费模式。取值范围：1、Dynamic（按小时计费）2、Month（按月计费）3、Year（按年计费）             | []                       | []                      | **Yes** |
| **Quantity**               | integer | 计费数量，指定计费周期的数量。按月/年计费时表示创建 Quantity 个月/年的磁盘，按小时计费时强制为1。取值范围：正整数 | []                       | []                      | **Yes** |
| **InternalIPVersion**      | string  | 内网协议，空值为所选VPC/外网默认网络的网络协议, 枚举支持 IPv4 IPv6 ALL 和空值                                     | []                       | []                      | No      |
| **InternalExpandIP**       | string  | 内网拓展IP地址                                                                                                    | []                       | []                      | No      |

### 响应字段

| 字段名    | 类型   | 描述信息                                                                                         | 示例                 |
| :-------- | :----- | :----------------------------------------------------------------------------------------------- | :------------------- |
| **VMID**  | string | 所克隆出来的虚拟机ID                                                                             | [vm-98d0o148l3vylg]  |
| **VMCID** | string | VMC资源ID，虚拟机克隆任务的元数据，方便系统知道“这台 VM 是从哪里、什么时候、用哪些磁盘克隆出来的 | [vmc-14d49mfu25vioz] |


## 获取 VMC 信息 - DescribeVMC

### 请求参数

| 字段名        | 类型    | 描述信息                                                                        | 示例                  | 参数来源         | 必填    |
| :------------ | :------ | :------------------------------------------------------------------------------ | :-------------------- | :--------------- | :------ |
| **Region**    | string  | 地域ID，地域，一批可共享的物理资源使用集合                                      | [develop]             | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离                      | [200000236]           | [DescribeUser]   | No      |
| **VMSPID**    | string  | 整机快照ID，查询由某个快照克隆出的所有虚拟机                                    | [vmsp-uunpn7ish05arm] | []               | **Yes** |
| **Limit**     | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 默认值：0        | [0]                   | []               | No      |
| **Offset**    | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。默认值：10 | [20]                  | []               | No      |

### 响应字段

| 字段名        | 类型    | 描述信息                                                                        | 示例                  |
| :------------ | :------ | :------------------------------------------------------------------------------ | :-------------------- |
| **Region**    | string  | 地域ID，地域，一批可共享的物理资源使用集合                                      | [develop]             |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离                      | [200000236]           |
| **VMSPID**    | string  | 整机快照ID，查询由某个快照克隆出的所有虚拟机                                    | [vmsp-uunpn7ish05arm] |
| **Limit**     | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 默认值：0        | [0]                   |
| **Offset**    | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。默认值：10 | [20]                  |


## 删除 VMC 资源 - DeleteVMC

### 请求参数

| 字段名        | 类型    | 描述信息                                                                                         | 示例                 | 参数来源         | 必填    |
| :------------ | :------ | :----------------------------------------------------------------------------------------------- | :------------------- | :--------------- | :------ |
| **Region**    | string  | 地域ID，地域，一批可共享的物理资源使用集合                                                       | [develop]            | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离                                       | [200000236]          | [DescribeUser]   | No      |
| **VMCID**     | string  | VMC资源ID，虚拟机克隆任务的元数据，方便系统知道“这台 VM 是从哪里、什么时候、用哪些资源克隆出来的 | [vmc-14d49mfu25vioz] | [DescribeVMC]    | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
| :----- | :--- | :------- | :--- |


## 取消整机克隆 - CancelCloneVMInstance

### 请求参数

| 字段名         | 类型    | 描述信息                                                                                         | 示例                 | 参数来源             | 必填    |
| :------------- | :------ | :----------------------------------------------------------------------------------------------- | :------------------- | :------------------- | :------ |
| **Region**     | string  | 地域ID，地域，一批可共享的物理资源使用集合                                                       | [develop]            | [DescribeRegion]     | **Yes** |
| **CompanyID**  | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离                                       | [200000236]          | [DescribeUser]       | No      |
| **VMCID**      | string  | VMC资源ID，虚拟机克隆任务的元数据，方便系统知道“这台 VM 是从哪里、什么时候、用哪些资源克隆出来的 | [vmc-14d49mfu25vioz] | [DescribeVMC]        | **Yes** |
| **VMID**       | string  | 克隆虚拟机ID                                                                                     | [vm-dfe3412f123]     | [DescribeVMInstance] | No      |
| **SourceVMID** | string  | 源虚拟机ID                                                                                       | [vm-sou2fwef12321]   | [DescribeVMInstance] | No      |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
| :----- | :--- | :------- | :--- |


## 虚拟机网络参数重置 - ResetVMNetConfig

### 请求参数

| 字段名        | 类型    | 描述信息                                                   | 示例                | 参数来源             | 必填    |
| :------------ | :------ | :--------------------------------------------------------- | :------------------ | :------------------- | :------ |
| **Region**    | string  | 地域ID，地域，一批可共享的物理资源使用集合                 | [develop]           | [DescribeRegion]     | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236]         | [DescribeUser]       | No      |
| **VMID**      | string  | 虚拟机ID                                                   | [vm-29by1ohfdf2f41] | [DescribeVMInstance] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
| :----- | :--- | :------- | :--- |


## 虚拟机更新VCPU绑定 - UpdateVMVCPUBinding

### 请求参数

| 字段名                     | 类型          | 描述信息                                                         | 示例                | 参数来源             | 必填    |
| :------------------------- | :------------ | :--------------------------------------------------------------- | :------------------ | :------------------- | :------ |
| **Region**                 | string        | 目的地域ID，您可以调用 DescribeRegion 方法，查看最新的地域列表。 | [develop]           | []                   | **Yes** |
| **VMID**                   | string        | 虚拟机ID                                                         | [vm-29by1ohfdf2f41] | [DescribeVMInstance] | **Yes** |
| **VCPUBindingNode**        | string        | VCPU绑定的物理节点                                               | [node-123456]       | [DescribeNode]       | No      |
| **VCPUBindingInfos**       | array<string> | VCPU绑定信息                                                     | []                  | []                   | No      |
| **VCPUBindingDegrandable** | bool          | VCPU绑定在宕机迁移时是否允许降级                                 | [true]              | []                   | No      |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
| :----- | :--- | :------- | :--- |

