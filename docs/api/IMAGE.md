# IMAGE

## 自制镜像复制成基础镜像 - CloneCustomImageToBaseImage

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **DestImageName** | string | 目标镜像名，指定复制后的镜像名称，长度为1-50个字符，名称只能包含中英文、数字、点（.）下划线（_）和中划线（-） | [clone-img] | [] | **Yes** |
| **SourceImageID** | string | 源镜像ID，可以通过DescribeImage来获取 | [image-ngod7vm5ow2cxf] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **DestImageID** | string | 克隆出来的目标镜像ID | [image-77rnt1gt5ovkp0] |


## 制作虚拟机镜像 - CreateCustomImage

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **ImageName** | string | 镜像名称，长度为1-50个字符，名称只能包含中英文、数字、点（.）下划线（_）和中划线（-） | [test_CreateCustomImage] | [] | **Yes** |
| **ImageDescription** | string | 备注，用于进行说明和注释，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [这是一个备注] | [] | No |
| **ProjectID** | string | 项目组ID，用于资源分组管理 | [project-5kyt09jzwcqh6j] | [ListProjects] | No |
| **TagKeyValuePairs** | array<string> | 标签键值对，用于资源标记和分类管理。格式为 key:value 的字符串，传入​​ Base64 编码的字符串。 | [dGVzdA==:dGVzdC0x] | [] | **Yes** |
| **VMID** | string | 源虚拟机ID | [vm-7rli7wr5jpk8h5] | [DescribeVMInstance] | **Yes** |
| **StorageSetType** | string | 存储集群，镜像存储的目标存储集群 | [StorageSetNPVK] | [DescribeStorageSet] | No |
| **DiskID** | string | 磁盘ID，指定制作镜像的源存储磁盘ID | [disk-2frgtg4354] | [DescribeDisk] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ImageID** | string | 镜像 ID，创建成功的镜像ID | [image-3r7855j9jjjml5] |


## 取消制作虚拟机镜像 - AbortCustomImage

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **ImageID** | string | 镜像ID，目的镜像ID。 | [image-3r7855j9jjjml5] | [DescribeImage] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 删除基础镜像 - DeleteBaseImage

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **ImageID** | string | 镜像ID，目的镜像ID。 | [image-3r7855j9jjjml5] | [DescribeImage] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 删除主机镜像 - DeleteCustomImage

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **ImageID** | string | 镜像ID，目的镜像ID。 | [image-3r7855j9jjjml5] | [DescribeImage] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取镜像 - DescribeImage

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **ProjectIDs** | array<string> | 项目组ID列表，用于资源分组管理，多个元素之间用,分隔。 | [project-5kyt09jzwcqh6j] | [ListProjects] | No |
| **Keyword** | string | 搜索关键词 | [example_key] | [] | No |
| **Limit** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 默认值：0 | [0] | [] | No |
| **Offset** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。默认值：10 | [20] | [] | No |
| **ImageIDs** | array<string> | 镜像ID列表，用于查询指定的镜像，多个元素之间用,分隔。 | [image-3r7855j9jjjml5] | [DescribeImage] | No |
| **ImageType** | string | 镜像类型，取值范围：1、Base（基础镜像）2、Custom（自制镜像） | [Base] | [] | No |
| **ImageFormat** | string | 镜像文件格式，取值范围：1、qcow2（QEMU镜像格式）2、iso（光盘镜像格式） | [qcow2] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 查询到的总记录数 | [] |
| **Infos** | array<[*ImageInfo*](#ImageInfo)> | 镜像详情 | [] |

### 数据模型

#### ImageInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Name** | string | 镜像名称，长度为1-50个字符，名称只能包含中英文、数字、点（.）下划线（_）和中划线（-） | [test] |
| **Remark** | string | 备注，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [备注一] |
| **Region** | string | 地域 | [cn] |
| **RegionAlias** | string | 地域别名 | [cn] |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] |
| **CompanyName** | string | 租户名称 | [jack.ma] |
| **Email** | string | 租户邮箱 | [jack.ma@ucloud.cn] |
| **ProjectID** | string | 项目组ID，用于资源分组管理 | [project-5kyt09jzwcqh6j] |
| **ProjectName** | string | 项目组名称 | [project-01] |
| **Reason** | string | 失败原因 | [存储集群空间不足] |
| **CreateTime** | integer | 创建时间，Unix 时间戳（秒级） | [1754644992] |
| **UpdateTime** | integer | 更新时间，Unix 时间戳（秒级） | [1754644992] |
| **Tags** | array<[*UnifiedTag*](#UnifiedTag)> | 标签 | [cluster-resouce.system:offline-node] |
| **ImageName** | string | 镜像名称，长度为1-50个字符，名称只能包含中英文、数字、点（.）下划线（_）和中划线（-） | [test] |
| **ImageDescription** | string | 镜像描述 | [备注一] |
| **Remarkable** | bool | 是否可以修改备注 | [true] |
| **ImageID** | string | 镜像ID | [image-3r7855j9jjjml5] |
| **ImageType** | string | 镜像类型，取值范围：1、Base（基础镜像）2、Custom（自制镜像） | [Custom] |
| **ImageFormat** | string | 镜像文件格式（qcow2:QEMU镜像格式，iso:光盘镜像格式） | [qcow2] |
| **ImageStatus** | string | 镜像状态 | [WaitMaking] |
| **ImageSize** | integer | 镜像大小，单位：GiB | [100] |
| **OSType** | string | 操作系统类型 | [Linux] |
| **OSName** | string | 操作系统名称 | [OpenEuler 22.03_SP4 x86_64] |
| **OSDistribution** | string | 操作系统发行版 | [OpenEuler] |
| **SetArch** | string | 集群架构 | [x86_64] |
| **CanVGPU** | bool | 是否支持VGPU | [false] |
| **SupportHotplug** | bool | 是否支持热插拔 | [false] |
| **PreparePrecent** | integer | 自制镜像进度 or 远程上传进度 | [100] |
| **SupportQGA** | bool | 是否支持QGA | [true] |
| **SupportCloudInit** | bool | 是否支持CloudInit | [true] |
| **Encrypted** | bool | 是否加密 | [true] |
| **ImportType** | string | 上传类型，取值范围：1、Remote（远程 URL） 2、Local（本地上传） | [Remote] |
| **BootloaderType** | string | 引导类型，取值范围：1、uefi 2、bios | [uefi] |
| **BoundVMSetIDs** | array<string> | 当前镜像绑定的计算集群ID列表，多个元素之间用,分隔。 | [compute-set-p40,compute-set-hygon] |
| **BoundVMIDs** | array<string> | ISO镜像绑定的虚拟机列表，多个元素之间用,分隔。 | [vm-qdfg545tf] |
| **ImageFrom** | string | 镜像来源，取值范围：1、vm-xx-boot（虚拟机制作） 2、image-xx（镜像复制），3、https://xx Local Upload（上传） 4、System（平台内置） | [System] |

#### UnifiedTag

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Key** | string |  | [] |
| **Value** | string |  | [] |

## 获取基础镜像权限信息 - DescribeBaseImage

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **ImageIDs** | array<string> | 镜像ID列表，用于查询指定的镜像，多个元素之间用,分隔。 | [image-3r7855j9jjjml5] | [DescribeImage] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 查询到的总记录数 | [] |
| **Infos** | array<[*BaseImageInfo*](#BaseImageInfo)> | 基础镜像详情信息 | [] |

### 数据模型

#### BaseImageInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Region** | string | 地域 | [cn] |
| **ImageID** | string | 镜像ID | [image-77rnt1gt5ovkp0] |
| **Name** | string | 名称 | [test] |
| **Description** | string | 描述 | [用于测试的基础镜像] |
| **CreateTime** | integer | 创建时间，Unix 时间戳（秒级） | [1754644992] |
| **UpdateTime** | integer | 更新时间，Unix 时间戳（秒级） | [1754644992] |
| **Permission** | string | 权限，哪些租户可以访问 | [all] |

## 上传镜像 - ImportImage

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | **Yes** |
| **ProjectID** | string | 项目组ID，用于资源分组管理 | [project-5kyt09jzwcqh6j] | [ListProjects] | No |
| **TagKeyValuePairs** | array<string> | 标签键值对，用于资源标记和分类管理。格式为 key:value 的字符串，传入​​ Base64 编码的字符串。 | [dGVzdA==:dGVzdC0x] | [] | **Yes** |
| **ImageName** | string | 镜像名称，长度为1-50个字符，名称只能包含中英文、数字、点（.）下划线（_）和中划线（-） | [image-test] | [] | **Yes** |
| **ImageFormat** | string | 镜像文件格式，取值范围：1、qcow2（QEMU镜像格式）2、iso（光盘镜像格式） | [iso] | [] | **Yes** |
| **ImageDescription** | string | 镜像描述，用于进行说明和注释，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [镜像测试] | [] | No |
| **LoadURL** | string | 来源URL，仅通过 URL 上传时需要指定 | [http://192.168.178.xxx/s3/download?dfe] | [] | No |
| **OSType** | string | 操作系统类型，取值范围：当前支持的操作系统类型 | [Linux] | [DescribeImageOSVersions] | **Yes** |
| **SetArch** | string | 集群架构，取值范围：当前支持的集群架构 | [x86_64] | [DescribeImageOSVersions] | **Yes** |
| **OSDistribution** | string | 操作系统发行版，取值范围：当前支持的操作系统发行版 | [Ubuntu TencentOS] | [DescribeImageOSVersions] | **Yes** |
| **OSVersion** | string | 操作系统版本，取值范围：当前支持的操作系统版本 | [2.4] | [DescribeImageOSVersions] | **Yes** |
| **ImageSize** | integer | 镜像大小，单位GiB | [100] | [] | No |
| **ImportType** | string | 上传类型，取值范围：1、Remote（远程 URL）2、Local（本地） | [Remote] | [] | **Yes** |
| **SupportQGA** | bool | 是否支持qemu-ga | [true] | [] | No |
| **SupportCloudInit** | bool | 是否支持cloud-init | [true] | [] | No |
| **Secret** | string | 镜像密钥 | [dregf9u9g23] | [] | No |
| **BootloaderType** | string | 虚拟机引导方式，取值范围：1、bios 2、uefi | [bios] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ImageID** | string | 镜像ID | [image-3r7855j9jjjml5] |
| **UploadID** | string | 上传ID，仅本地上传时返回 | [uploadid-fqfeger2] |
| **Bucket** | string | 存入的桶 | [mirror] |


## 合并本地上传镜像分片 - CompleteImageMultipartUpload

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **ImageID** | string | 镜像ID，上传镜像后返回目标镜像ID | [image-ef8ta3pyt3upwv] | [ImportImage] | **Yes** |
| **UploadID** | string | 上传ID，本地上传时返回的上传任务编号 | [ODRiYWMzNWUtZDNjMi00OTUxLWJlZTUtMDYxZDNiOWRmMmJmLjk5M2FiNzc3LWNlZjQtNDIwZS05ODkwLThkODQzMjhmNjBjNngxNzU1NTg3Nzc3MTMzNjQ3MDA3] | [ImportImage] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 取消本地上传镜像 - AbortImageMultipartUpload

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | **Yes** |
| **UploadID** | string | 上传ID，本地上传时返回的上传任务编号 | [ODRiYWMzNWUtZDNjMi00OTUxLWJlZTUtMDYxZDNiOWRmMmJmLjk5M2FiNzc3LWNlZjQtNDIwZS05ODkwLThkODQzMjhmNjBjNngxNzU1NTg3Nzc3MTMzNjQ3MDA3] | [ImportImage] | No |
| **ImageID** | string | 镜像ID，上传镜像后返回目标镜像ID | [image-ef8ta3pyt3upwv] | [ImportImage] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取镜像下载地址 - GetImageDownloadURL

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **ImageID** | string | 镜像ID，目的镜像ID。 | [image-lagwg0dgmdv13q] | [DescribeImage] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Info** | object | 镜像 URL 详情信息 | [] |

### 数据模型

#### ImageDownloadInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **DownloadQuery** | string | 下载 Query，需要拼接当前地域 s3 的地址获取完整的下载链接。例如：http://xxx.xx.xxx.xx/s3/download? + DownloadQuery | [] |
| **ExpiredTime** | integer | 过期时间 | [] |

## 查询镜像系统规格 - DescribeImageOSVersions

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | No |
| **ImageOSIDs** | array<string> | 镜像的系统ID列表，用于查询指定的系统，多个元素之间用,分隔。 | [] | [DescribeImageOSVersions] | No |
| **Limit** | integer | 分页偏移量，指定跳过的记录数，用于实现分页查询。取值范围：≥0。 默认值：0 | [0] | [] | No |
| **Offset** | integer | 分页大小，指定每页返回的记录数，用于控制返回数据量。取值范围：1-100。默认值：10 | [20] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 查询到的总记录数 | [] |
| **Infos** | array<[*ImageOSInfo*](#ImageOSInfo)> | 镜像的系统信息详情 | [] |

### 数据模型

#### ImageOSInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ImageOSID** | string | 镜像操作系统ID | [] |
| **OSType** | string | 操作系统类型 | [] |
| **SetArch** | string | 系统架构 | [x86_64] |
| **OSDistribution** | string | 系统平台 | [CentOS] |
| **OSVersion** | string | 系统版本，centos为6、7；ubuntu为14、16；windows可选；支持用户自定义 | [9.2] |
| **NetworkMode** | string | 网络模版，取值范围：1、NetworkScripts 2、Interfaces 3、Netplan 4、None | [NetworkScripts] |
| **Remark** | string | 备注，长度为 0-100 个英文或中文字符，不能使用的http://或https://等非法字符。可为空 | [windows系统] |
| **Status** | string | 状态，值：Ready（就绪） | [Ready] |

## 修改镜像属性 - UpdateImage

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000236] | [DescribeUser] | **Yes** |
| **ImageID** | string | 镜像ID，目的镜像ID。 | [image-lagwg0dgmdv13q] | [DescribeImage] | **Yes** |
| **OSType** | string | 操作系统类型，取值范围：当前支持的操作系统类型 | [Linux] | [DescribeImageOSVersions] | **Yes** |
| **SetArch** | string | 集群架构，取值范围：当前支持的集群架构 | [x86_64] | [DescribeImageOSVersions] | **Yes** |
| **OSDistribution** | string | 操作系统发行版，取值范围：当前支持的操作系统发行版 | [Ubuntu TencentOS] | [DescribeImageOSVersions] | **Yes** |
| **OSVersion** | string | 操作系统版本，取值范围：当前支持的操作系统版本 | [2.4] | [DescribeImageOSVersions] | **Yes** |
| **SupportQGA** | bool | 是否支持qemu-ga | [true] | [] | No |
| **SupportCloudInit** | bool | 是否支持cloud-init | [true] | [] | No |
| **BootloaderType** | string | 虚拟机引导方式，取值范围：1、bios 2、uefi | [bios] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|

