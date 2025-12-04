# VM 模块 E2E 测试设计

本文档定义了 VM（虚拟机）模块的高覆盖率、高质量 E2E 测试场景和测试用例。

## 1. 概述

### 1.1 测试目标

- **API 覆盖率**：100% 覆盖 VM 模块所有 Action
- **参数覆盖率**：覆盖必填参数、常用可选参数、边界值
- **场景覆盖率**：覆盖资源完整生命周期（Create → Query → Update → Delete → Recycle）
- **质量指标**：
  - Happy Path 场景正确性
  - 边界条件和异常场景处理
  - 状态机正确性（Running ↔ Stopped, etc.）
  - 资源隔离和权限验证

### 1.2 关键业务规则

1. **标签 (Tag)**：`TagKeyValuePairs` 不是必填，但若指定需满足 Base64 编码格式 `key:value`
2. **删除生命周期**：
   - 调用 `DeleteVMInstance` → 虚拟机进入"删除中"状态
   - 进入回收站（需通过回收站 API 查询）
   - 在回收站中销毁 → 才算完全删除
3. **镜像选择**：优先选择 CentOS 7.4（OSName="CentOS 7.4"）
4. **集群选择**：优先选择空闲的计算集群（VMType）和存储集群（BootDiskSetType）
5. **密码规则**：长度 6-30 字符，需包含小写字母和特殊字符（非空格）

---

## 2. 测试场景分类

### 2.1 Smoke 测试（基础冒烟）

#### S001: 基础创建和查询

```
用途：验证最小化参数下的创建和查询功能
前置条件：
  - 已获取有效 Region
  - 已获取或创建 VPC 和 Subnet
  - 已获取可用的计算集群（VMType）
  - 已获取可用的存储集群（BootDiskSetType）
  - 已获取有效的外网线路（Segment/OperatorName）
步骤：
  1. 调用 CreateVMInstance
     - 使用最小参数组合：Region, CompanyID, Name, VMType, CPU, Memory, 
       BootDiskSetType, BootDiskSpace, VPCID, SubnetID, OperatorName, 
       IPVersion, Bandwidth, ChargeType="Dynamic", Quantity=1
     - 不指定 ImageID，使用 OSType/OSDistribution/OSVersion（CentOS）
     - 不指定 TagKeyValuePairs
  2. 记录返回的 VMID
  3. 验证 VMID 非空
  4. 调用 DescribeVMInstance 查询该 VM
     - 传入 VMID
  5. 验证响应包含完整的虚拟机信息：Name, Status, CPU, Memory, ImageID, etc.
  6. 验证虚拟机状态为 "Starting" 或 "Running"
期望结果：虚拟机创建成功，可查询到完整信息
```

#### S002: 启动/停止/重启基本功能

```
用途：验证虚拟机基本生命周期操作
前置条件：
  - 已创建一个运行中的虚拟机（使用 S001 结果）
步骤：
  1. 确保虚拟机状态为 "Running"
  2. 调用 StopVMInstance（关闭）
  3. 轮询 DescribeVMInstance 直至状态为 "Stopped"（超时 120s）
  4. 验证虚拟机状态转换：Running → Stopping → Stopped
  5. 调用 StartVMInstance（启动）
  6. 轮询 DescribeVMInstance 直至状态为 "Running"（超时 120s）
  7. 验证虚拟机状态转换：Stopped → Starting → Running
  8. 调用 RestartVMInstance（重启）
  9. 轮询 DescribeVMInstance（预期 Restarting → Running，超时 120s）
期望结果：虚拟机状态按预期转换，无异常
```

---

### 2.2 Functional 测试（功能完整性）

#### F001: 完整创建参数覆盖

```
用途：验证各个创建参数的有效性和组合
前置条件：
  - 已获取有效的资源 ID（VPC, Subnet, SecurityGroup, Image, etc.）
步骤：
  1. 测试参数组合 A：带可选参数的创建
     CreateVMInstance(
       Region, CompanyID, VMType, CPU, Memory, VPCID, SubnetID,
       OperatorName, IPVersion, Bandwidth, ChargeType, Quantity,
       Name, Remark,
       BootSourceType="Image", ImageID="centos74_id",
       BootDiskSetType, BootDiskSpace,
       DataDiskSetType, DataDiskSpace, DataDiskAutoMount=true,
       LANSGID, WANSGID,
       Password, Hostname, DNS,
       CPUMode="host-passthrough",
       HighAvailability="NeverStop",
       BootloaderType="uefi",
       TagKeyValuePairs=["env:test", "app:demo"]  # Base64 编码
     )
  2. 记录返回的 VMID
  3. DescribeVMInstance 验证所有参数均已生效：
     - CPU, Memory 匹配
     - ImageID 匹配
     - Hostname, DNS 匹配
     - Tags 包含指定的键值对
     - BootloaderType 为 "uefi"
     - CPUMode 为 "host-passthrough"
     - HighAvailability 为 "NeverStop"
     - SecurityGroups 包含 LANSGID 和 WANSGID
期望结果：各参数组合均创建成功，参数生效
```

#### F002: 虚拟机查询和筛选

```
用途：验证 DescribeVMInstance 的各种查询能力
前置条件：
  - 已创建多个虚拟机（不同 VPC、Status、Project 等）
步骤：
  1. 不传参数查询（分页）
     DescribeVMInstance(Limit=10, Offset=0)
     验证 TotalCount, Infos 数组非空
  2. 按 VMID 查询
     DescribeVMInstance(VMIDs=[vm_id_1, vm_id_2])
     验证返回的虚拟机数量匹配
  3. 按 VPC 和 Subnet 筛选
     DescribeVMInstance(VPCID=vpc_id, SubnetID=subnet_id)
     验证所有返回虚拟机的 VPCID 和 SubnetID 匹配
  4. 按状态筛选
     DescribeVMInstance(States=["Running", "Stopped"])
     验证返回虚拟机状态均在指定范围内
  5. 按排序字段排序
     DescribeVMInstance(SortBy="CreateTime", Sort="Descending")
     验证返回结果按创建时间降序
期望结果：各种查询方式均返回正确结果
```

#### F003: 密码重置

```
用途：验证虚拟机密码重置功能
前置条件：
  - 已创建一个虚拟机
步骤：
  1. 虚拟机保持 Running 状态
  2. 调用 ResetVMInstancePassword
     ResetVMInstancePassword(VMID, Password="newPass@123")
  3. 轮询 DescribeVMInstance 确保无异常状态
  4. 验证虚拟机状态不变
期望结果：密码重置成功，虚拟机正常运行
```

#### F004: 系统重装

```
用途：验证虚拟机系统重装功能
前置条件：
  - 已创建一个虚拟机
  - 已获取另一个有效的镜像 ID（如 Ubuntu 镜像）
步骤：
  1. 虚拟机为 Stopped 状态
  2. 调用 ReinstallVMInstance
     ReinstallVMInstance(
       VMID,
       ImageID=<new_image_id>,
       Password="newPass@123",
       UserData=<cloud_init_script>
     )
  3. 轮询 DescribeVMInstance 直至状态为 Running（超时 180s）
  4. 验证虚拟机状态：Stopped → Reinstalling → Running
  5. DescribeVMInstance 验证 ImageID 已更新
期望结果：系统重装成功，虚拟机运行新镜像
```

#### F005: 虚拟机配置升降级

```
用途：验证虚拟机配置修改（CPU/Memory/GPU）
前置条件：
  - 已创建一个虚拟机（如 1核 2G）
步骤：
  1. 虚拟机为 Stopped 状态
  2. 调用 GetPaymentOfPremium 获取升级差价
     GetPaymentOfPremium(VMID, CPU=2, Memory=4096)
  3. 验证返回 Price 和 OrderType（应为 "3" 表示升级）
  4. 调用 ResizeVMConfig 执行升级
     ResizeVMConfig(VMID, CPU=2, Memory=4096)
  5. 轮询 DescribeVMInstance 直至配置生效
  6. 验证虚拟机 CPU=2, Memory=4096
期望结果：配置升降级成功，差价计算正确
```

#### F006: 高级选项更新

```
用途：验证虚拟机高级参数的修改
前置条件：
  - 已创建一个虚拟机
步骤：
  1. 调用 UpdateVMAdvancedOptions 修改多个高级参数
     UpdateVMAdvancedOptions(
       VMID,
       DNS="8.8.8.8,8.8.4.4",
       CPUMode="custom",
       CPUModel="default",
       HighAvailability="None",
       DiskCacheMode="writeback",
       UserData=<new_script>
     )
  2. DescribeVMInstance 验证参数已更新
     - DNS 匹配
     - CPUMode 为 "custom"
     - HighAvailability 为 "None"
期望结果：高级参数更新成功
```

#### F007: VNC 和 Spice 信息获取

```
用途：验证远程连接信息获取
前置条件：
  - 已创建一个运行中的虚拟机
步骤：
  1. 调用 GetVMVNCInfo(VMID)
     验证返回：VNCIP, VNCPort, VNCPassword 均非空
  2. 调用 GetVMSpiceInfo(VMID)
     验证返回：SpiceIP, SpicePort, SpicePassword 均非空
期望结果：连接信息获取成功
```

#### F008: 虚拟机价格查询

```
用途：验证虚拟机价格计算
前置条件：
  - 已获取有效的配置参数
步骤：
  1. 调用 GetVMInstancePrice 查询多种配置价格
     GetVMInstancePrice(
       VMType=<cluster>, CPU=2, Memory=4096, Count=1,
       ChargeType="Dynamic", BootDiskSpace=100, BootDiskSetType=<storage>
     )
  2. 验证返回的 Infos 数组包含 Price 和 ChargeType
  3. 验证 Price > 0
期望结果：价格查询成功，返回有效价格
```

---

### 2.3 Boundary 测试（边界条件）

#### B001: 虚拟机名称边界

```
用途：验证虚拟机名称的长度和字符限制
前置条件：
  - 已获取有效的资源信息
步骤：
  1. 最短名称：Name="a" （长度 1）
     验证创建成功
  2. 最长名称：Name="x" * 50 （长度 50）
     验证创建成功
  3. 超长名称：Name="x" * 51 （长度 51）
     验证创建失败，返回参数错误
  4. 特殊字符：Name="test-vm.01_test" （包含 -, ., _）
     验证创建成功
  5. 非法字符：Name="test vm@test" （包含空格和 @）
     验证创建失败
期望结果：名称验证符合规范
```

#### B002: 资源规格边界

```
用途：验证 CPU、Memory、Disk 的规格限制
前置条件：
  - 已获取有效的资源信息
步骤：
  1. 最小 CPU：CPU=1
     验证创建成功
  2. Memory 必须为 1024 倍数
     - Memory=2048：验证成功
     - Memory=2000：验证失败
  3. BootDiskSpace 范围 [20, 2000]
     - BootDiskSpace=20：验证成功
     - BootDiskSpace=2000：验证成功
     - BootDiskSpace=19：验证失败
     - BootDiskSpace=2001：验证失败
期望结果：规格验证符合文档定义
```

#### B003: 密码复杂性边界

```
用途：验证密码的复杂性要求
前置条件：
  - 已获取有效的资源信息
步骤：
  1. 有效密码：Password="abcD@1234" （6-30字符，含小写、特殊字符）
     验证创建成功
  2. 长度过短：Password="a@1" （3 字符）
     验证创建失败
  3. 无特殊字符：Password="abcd1234" 
     验证创建失败
期望结果：密码验证符合复杂性要求
```

---

### 2.4 Failure 测试（异常场景）

#### F401: 无效参数异常

```
用途：验证 API 对无效参数的处理
步骤：
  1. 缺少必填参数
     CreateVMInstance(omit Region)
     验证返回错误，提示 "Region" 为必填
  2. 无效的 Region
     CreateVMInstance(Region="invalid-region")
     验证返回错误
  3. 无效的 VPCID
     CreateVMInstance(VPCID="invalid-vpc")
     验证返回错误
期望结果：所有无效参数均返回合适的错误信息
```

#### F402: 资源不存在异常

```
用途：验证对不存在资源的操作
步骤：
  1. DescribeVMInstance(VMID="vm-nonexistent")
     验证返回空列表或 TotalCount=0
  2. StartVMInstance(VMID="vm-nonexistent")
     验证返回错误 "Resource not found"
期望结果：不存在的资源操作均返回合适错误
```

#### F403: 状态冲突异常

```
用途：验证状态不匹配操作的处理
前置条件：
  - 已创建一个虚拟机并确保其状态
步骤：
  1. 虚拟机为 Running 状态时调用 StartVMInstance
     验证返回错误 "Invalid state"
  2. 虚拟机为 Running 状态时调用 ResizeVMConfig
     验证返回错误（配置修改需要 Stopped）
期望结果：状态冲突操作均返回合适错误
```

---

### 2.5 State Machine 测试（状态机）

#### SM001: 虚拟机完整生命周期状态转换

```
用途：验证虚拟机各状态的正确转换
步骤：
  1. 创建虚拟机
     期望状态转换：(None) → Starting → Running
  2. 关闭虚拟机
     期望状态转换：Running → Stopping → Stopped
  3. 启动虚拟机
     期望状态转换：Stopped → Starting → Running
  4. 重启虚拟机
     期望状态转换：Running → Restarting → Running
  5. 断电虚拟机
     期望状态转换：Running/Stopped → Poweroffing → Stopped
  6. 重装系统（虚拟机需为 Stopped）
     期望状态转换：Stopped → Reinstalling → Running
  7. 删除虚拟机
     期望状态转换：Any → Deleting → (in recycle bin)
期望结果：所有状态转换符合预期
```

---

### 2.6 Integration 测试（资源联动）

#### I001: 虚拟机与磁盘的联动

```
用途：验证虚拟机创建时自动创建和挂载磁盘
前置条件：
  - 已获取有效的存储集群信息
步骤：
  1. 创建虚拟机，指定 BootDiskSetType 和 BootDiskSpace
  2. DescribeVMInstance 验证虚拟机信息
  3. 在返回的 DiskInfos 中验证：
     - 存在 Type="boot" 的系统盘
     - Size 匹配指定的 BootDiskSpace
     - StorageSetType 匹配 BootDiskSetType
期望结果：虚拟机创建时自动创建相关磁盘
```

#### I002: 虚拟机与 IP 的联动

```
用途：验证虚拟机创建时自动分配 IP
前置条件：
  - 已获取有效的 VPC、Subnet、Segment 信息
步骤：
  1. 创建虚拟机，指定 VPCID、SubnetID、OperatorName
  2. DescribeVMInstance 验证虚拟机信息
  3. 在返回的 IPInfos 中验证：
     - 存在内网 IP（Type="private"）
     - 若指定了 InternetIP，则存在外网 IP（Type="public"）
期望结果：虚拟机创建时自动分配 IP
```

#### I003: 虚拟机与安全组的联动

```
用途：验证虚拟机与安全组的关联
前置条件：
  - 已创建安全组
  - 已创建虚拟机并关联安全组
步骤：
  1. 创建虚拟机，指定 LANSGID 和 WANSGID
  2. DescribeVMInstance 验证虚拟机信息
  3. 在返回的 IPInfos 中验证关联的安全组
期望结果：虚拟机与安全组正确关联
```

#### I004: 虚拟机与标签的联动

```
用途：验证虚拟机标签的创建和查询
前置条件：
  - 已获取有效的标签 KV 对（Base64 编码）
步骤：
  1. 创建虚拟机，指定 TagKeyValuePairs=["env:prod", "app:web"]
  2. DescribeVMInstance 验证虚拟机信息
  3. 验证虚拟机 Tags 数组包含所有指定的标签
期望结果：虚拟机标签正确创建和关联
```

---

### 2.7 Delete Lifecycle 测试（删除生命周期）

#### D001: 虚拟机删除和回收站流程

```
用途：验证虚拟机从删除到最终销毁的完整流程
前置条件：
  - 已创建一个虚拟机
步骤：
  1. 调用 DeleteVMInstance(VMID)
  2. 轮询 DescribeVMInstance 验证虚拟机状态转换：
     现有状态 → Deleting → (disappears from normal list)
  3. 确认虚拟机已进入回收站（Recycle API 可查询到）
  4. 从回收站查询该虚拟机
     验证虚拟机状态为 "Deleted" 或类似
  5. 调用回收站销毁 API 彻底删除虚拟机
  6. 确认虚拟机从回收站消失
期望结果：虚拟机删除和销毁流程完整
```

#### D002: 虚拟机回收站恢复

```
用途：验证从回收站恢复虚拟机
前置条件：
  - 虚拟机已被删除并进入回收站
步骤：
  1. 从回收站查询虚拟机
  2. 调用回收站恢复 API 恢复虚拟机
  3. 验证虚拟机重新出现在虚拟机列表中
期望结果：虚拟机从回收站成功恢复
```

---

### 2.8 Advanced Features（高级功能）

#### A001: GPU 虚拟机创建和查询

```
用途：验证 GPU 虚拟机的创建和信息查询
前置条件：
  - 已获取支持 GPU 的计算集群
  - 已获取有效的 GPU 配置
步骤：
  1. 创建虚拟机，指定：
     GPUType="GPU", GPU=2, GPUMdevName="nvidia-2080ti"
  2. DescribeVMInstance 验证虚拟机信息：
     - GPUType 为 "GPU"
     - GPU 颗数为 2
期望结果：GPU 虚拟机创建和查询成功
```

#### A002: 隔离组关联

```
用途：验证虚拟机与隔离组的关联
前置条件：
  - 已创建隔离组
步骤：
  1. 创建虚拟机，指定 IGID=<isolation_group_id>
  2. DescribeVMInstance 验证虚拟机信息：
     - IGIDs 数组包含指定的隔离组 ID
期望结果：隔离组关联成功
```

#### A003: USB 设备挂载

```
用途：验证 USB 设备挂载功能
前置条件：
  - 已获取有效的 USB 设备 ID
步骤：
  1. 创建虚拟机，指定：
     USBDeviceID=<usb_id>, USBAttachType="redir"
  2. DescribeVMInstance 验证虚拟机信息：
     - AttachedUSBIDs 包含指定的 USB 设备
期望结果：USB 设备挂载成功
```

---

## 3. 测试执行规范

### 3.1 环境配置

所有测试应从环境变量读取：
```
USTACK_API_BASE_URL       - API 基础 URL
USTACK_API_PUBLIC_KEY     - 签名公钥
USTACK_API_PRIVATE_KEY    - 签名私钥
USTACK_API_REGION         - 地域 ID
USTACK_API_COMPANY_ID     - 租户 ID
```

### 3.2 Ginkgo 标签和组织

```go
// Smoke 测试
Describe("VM Smoke Tests", Label("Feature:VM", "Tier:Smoke"), func() {
  It("S001: Create and Query VM", func() { ... })
  It("S002: Basic Lifecycle Operations", func() { ... })
})

// Functional 测试
Describe("VM Functional Tests", Label("Feature:VM", "Tier:Functional"), func() {
  It("F001: Complete Creation Parameters", func() { ... })
})
```

### 3.3 测试执行命令

运行所有 VM 测试：
```bash
ginkgo -r -v --label-filter="Feature:VM" test/e2e
```

运行特定 Tier：
```bash
ginkgo -r -v --label-filter="Feature:VM && Tier:Smoke" test/e2e
```

---

## 4. 覆盖率目标

### 4.1 API 覆盖率：目标 100%

- CreateVMInstance ✓
- DescribeVMInstance ✓
- DeleteVMInstance ✓
- StartVMInstance ✓
- StopVMInstance ✓
- RestartVMInstance ✓
- PoweroffVMInstance ✓
- ReinstallVMInstance ✓
- ResetVMInstancePassword ✓
- ResizeVMConfig ✓
- UpdateVMAdvancedOptions ✓
- GetVMVNCInfo ✓
- GetVMSpiceInfo ✓
- GetVMInstancePrice ✓
- GetPaymentOfPremium ✓
- DescribeCIStatus ✓

### 4.2 参数覆盖率：目标 100%

- 所有必填参数 ✓
- 常用可选参数 ✓
- 参数组合 ✓

### 4.3 场景覆盖率：目标 100%

- 生命周期完整性 ✓
- 状态机正确性 ✓
- 资源联动 ✓
- 异常处理 ✓
- 删除生命周期 ✓
- 高级功能 ✓

---

*本文档持续更新，最后修改时间：2024年*
