# 33 Redis

## 33.1 概览

Redis是平台提供的一种数据库服务，支持单机版和主备版两种机型，并提供了备份、监控等功能，满足高读写的性能场景。

## 33.2 创建Redis

平台用户可以通过指定计算集群、存储集群、内存容量、参数模板、VPC、子网、外网IP、外网安全组、Redis名称/备注、密码、项目组等相关基础信息创建Redis，可通过导航栏进入[Redis]资源控制台，通过“**创建**”按钮进入向导页面，如下图所示：

![createRedis](../images/userguide/createRedis.png)

1. 选择并配置Redis的基础配置、网络设置及管理配置信息：

* 计算集群：支持选择计算集群，并展示所选计算集群的CPU和内存使用率；
* 存储集群：支持选择存储集群，并展示所选存储集群的存储用量；
* 机型：支持 单机版和主备版 两种机型；
* 内存容量：选择创建Redis的内存容量，支持1G、2G、4G、8G等；
* 版本：支持Redis 4.0与Redis 7.0两种版本；
* 线程数量：Redis4.0版本不支持多线程；Redis7.0版本支持2、4、6三种线程数量；
* VPC和子网：创建Redis时必须选择 VPC 网络和所属子网，即选择要加入的网络及 IP 网段；
* 外网IP：为Redis提供外网访问服务，Redis支持最多绑定一个外网IP；

2. 选择购买数量和付费方式，确认订单金额并点击“立即购买” 进行Redis的创建：

- 购买数量：默认支持创建 1 个Redis；
- 付费方式：选择Redis的计费方式，支持按月、按年、按时三种方式，可根据需求选择合适的付费方式；
- 合计费用：基于购买数量与付费方式的总费用展示；
- 立即购买：点击立即购买后，会返回Redis资源列表页，在列表页可查看Redis的创建过程。

## 33.3 查看Redis列表

平台支持用户查看Redis列表信息，包括名称、资源ID、状态、容量使用率、机型、版本、线程数量、IP和端口、VPC、子网、安全组、实例容量、计费方式、项目组、创建时间、过期时间、操作。可通过导航栏进入[Redis]资源控制台查看，如下图所示：

![Redislist](../images/userguide/Redislist.png)

* 名称：Redis的名称，可通过点击名称进入Redis详情页；
* 资源ID：Redis的资源ID，作为全局唯一标识；
* 状态：Redis的状态，包括启动中、升级中、删除中等；
* 容量使用率：Redis本身的内存容量使用率；
* 机型：Redis的机型信息，支持单机版和主备版；
* 版本：Redis的版本信息，支持Redis4.0和Redis7.0；
* 线程数量：Redis7.0版本，支持选择线程数量：2、4、6；
* IP和端口：Redis的网络信息，包括内/外网IP和端口；
* 实例容量：Redis创建时选择的内存容量信息，可通过升级内存功能升级内存容量；
* VPC/子网：Redis内网信息，包括VPC/子网的名称和资源ID；
* 安全组：Redis绑定外网IP后支持关联安全组；
* 计费方式：Redis资源的计费方式，支持小时、月、年；
* 项目组：Redis资源所属的项目组，可通过项目组的转入/转出功能修改关联关系；
* 创建时间/过期时间；Redis的创建时间和过期时间；
* 操作：Redis资源的可操作项内容，包括创建从库、删除、续费、重置密码、应用参数模板、升级内存、清除数据、升级至主备版、修改告警模板、绑定外网IP、解绑外网IP、修改安全组等。

## 33.4 查看Redis概览信息

平台支持用户查看Redis资源概览信息，包括基本信息、配置信息、监控信息。可通过点击Redis列表中的名称进入概览页面，如下图所示：

![Redisoverview](../images/userguide/Redisoverview.png)

## 33.5 Redis创建从库

平台支持用户对Redis主库进行创建从库的操作，可选择计算集群、内存容量、外网IP等，从库内存容量不可低于主库，机型只能为单机版。可通过Redis列表中操作项的“**创建从库**”按钮进行操作，每个Redis主库支持最多创建5个从库。如下图所示：

![createRedis3](../images/userguide/createRedis3.png)

## 33.6 Redis续费

平台支持用户对Redis进行续费操作，可通过Redis列表中操作项的“**续费**”按钮进行操作，如下图所示：

![Redisrenew](../images/userguide/Redisrenew.png)

续费支持更改续费方式和续费时长，更改续费方式只支持由短周期改为长周期，比如从“**小时**”更改为“**月**”。

## 33.7 Redis重置密码

平台支持用户对Redis进行设置密码操作，可通过Redis列表中操作项的“**设置密码**”按钮进行操作，如下图所示：

![Redisreset1](../images/userguide/Redisreset1.png)

![Redisreset2](../images/userguide/Redisreset2.png)

## 33.8 参数配置

平台支持用户对Redis进行参数配置相关操作，包括修改实例参数、应用参数模板、导入参数文件、导出为模板、导出参数文件。可点击Redis名称进入详情页，切换到“**参数设置**”页面查看；也可通过Redis列表操作项的“**修改实例参数**”按钮点击跳转，如下图所示：

![Redisconfig2](../images/userguide/Redisconfig2.png)

![Redisconfig1](../images/userguide/Redisconfig1.png)

### 33.8.1 修改实例参数

平台支持用户对Redis进行修改实例参数操作，可进入“**参数设置**”页面，点击“**编辑**”按钮进行操作，如下图所示：

![editRedis](../images/userguide/editRedis.png)

### 33.8.2 应用参数模板

平台支持用户对Redis进行应用参数模板操作，可点击“**应用参数模板**”进行操作，，如下图所示：

![Redistemplate](../images/userguide/Redistemplate.png)

### 33.8.3 导入参数文件

平台支持用户对Redis进行导入参数文件操作，可点击“**导入参数文件**”进行操作，如下图所示：

![Redisfile1](../images/userguide/Redisfile1.png)

### 33.8.4 导出为模板

平台支持用户对Redis参数配置进行导出为模板操作，可点击“**导出为模板**”进行操作，如下图所示：

![Redistemplate3](../images/userguide/Redistemplate3.png)

导出为模板操作成功后，参数模板列表新增一条模板数据。

### 33.8.5 导出参数文件

平台支持用户对Redis参数配置进行导出参数文件操作，可点击“**导出参数文件**”进行操作，下载参数配置文件至本地。

![Redistemplate4](../images/userguide/Redistemplate4.png)

## 33.9 修改配置

平台支持用户对Redis进行在线升级内存操作，可点击Redis列表中操作项的“**修改配置**”按钮进行操作，如下图所示：

![Redisupgrade](../images/userguide/Redisupgrade.png)

平台支持 Redis 实例离线配置升降级，实例关机后，可执行内存规格升级或降级操作。点击列表中操作项的“**修改配置**”按钮进行操作，如下图所示：

![Redisupgrade2](../images/userguide/Redisupgrade2.png)

注意事项：
* Redis ARM版本不支持在线升级内存容量
* 从库配置要大于等于主库，Redis主从库升级需要先升级从库后升级主库，降级需要先降级主库后降级从库

## 33.10 升级主备版

平台支持用户对单机版Redis进行升级至主备版操作，其余配置不可修改。可点击Redis列表中操作项的“**升级至主备版**”按钮进行操作，如下图所示：

![upgradeRedismodel](../images/userguide/upgradeRedismodel.png)

## 33.11 修改告警模板

平台支持用户对Redis进行修改告警模板操作。可点击Redis列表中操作项的“**修改告警模板**”按钮进行操作，如下图所示：

![Redisalarm](../images/userguide/Redisalarm.png)

## 33.12 网络

平台支持用户查看Redis的网络信息，包括基本信息和IP列表。可点击Redis名称进入详情页，切换到“**网络**”页面进行查看，如下图所示：

![Redisnetwork](../images/userguide/Redisnetwork.png)

### 33.12.1 查看网络列表

平台支持用户查看Redis的网络列表信息，包括IP、IP ID、IP版本、状态、网络类型、所属网络、是否VIP、带宽、绑定资源、MAC地址、操作，如下图所示：

![Redisnetworklist](../images/userguide/Redisnetworklist.png)

- IP：IP地址；
- IP ID：IP的ID，内网IP没有具体ID，外网IP会展示ID；
- IP版本：Redis目前只支持IPv4；
- 网络类型：有内网和外网两周网络类型；
- 所属网络：IP所属的网络；
- 绑定资源：IP所绑定的资源名和资源ID。

### 33.12.2 绑定外网IP

平台支持用户对Redis进行绑定外网IP操作，可通过“**绑定**”按钮进行操作，也可通过Redis列表中操作项的“**绑定外网IP**”进行操作，每个Redis支持最多绑定一个外网IP。如下图所示：

![Redisbindeip](../images/userguide/Redisbindeip.png)

外网IP绑定成功后，在Redis网络列表中新增一条IP数据。

### 33.12.3 解绑外网IP

平台支持用户对已绑定外网IP的Redis进行解绑外网IP操作，可通过Redis网络详情页的“**解绑**”按钮进行操作，也可通过Redis列表中操作项的“**解绑外网IP**”进行操作，如下图所示：

![Redisunbindeip](../images/userguide/Redisunbindeip.png)

## 33.13 修改安全组

平台支持用户对已绑定外网IP的Redis进行修改安全组操作，可通过Redis列表中操作项的“**修改安全组**”进行操作，如下图所示：

![editRedissg](../images/userguide/editRedissg.png)

## 33.14 清理数据

平台支持用户对Redis进行清理数据操作，可通过Redis列表中操作项的“**清理数据**”进行操作，如下图所示：

![cleanRedis](../images/userguide/cleanRedis.png)

## 33.15 备份管理

### 33.15.1 查看备份管理列表

平台支持用户查看备份管理列表信息，包括资源ID、状态、所属存储池、来源备份计划、保留时间(天)、保存路径、创建时间、更新时间、到期时间、操作。可点击Redis名称进入详情页，切换到“**备份管理**”页面进行查看，如下图所示：

![Redisdbs](../images/userguide/Redisdbs.png)

### 33.15.2 删除备份

平台支持用户对备份数据进行删除操作，可点击备份列表中操作项的”**删除**“按钮进行操作，也可通过备份列表的“**批量删除**”按钮进行操作，如下图所示：

![deleteRedisdbs](../images/userguide/deleteRedisdbs.png)

## 33.16 查看操作日志

平台支持用户查看Redis的操作日志，并可根据操作结果和操作周期进行筛选，可点击Redis名称进入详情页，切换到“**操作日志**”页面进行查看，如下图所示：

![Redislog](../images/userguide/Redislog.png)

## 33.17 查看事件

平台支持用户查看Redis的事件，并可根据事件周期进行筛选，可点击Redis名称进入详情页，切换到“**事件**”页面进行查看，如下图所示：

![Redisevent](../images/userguide/Redisevent.png)

## 33.18 删除Redis

平台支持用户对Redis进行删除操作，删除主库前需先将从库删除。可点击Redis列表中操作项的“**删除**”按钮进行操作，如下图所示：

![Redisdelete](../images/userguide/Redisdelete.png)

## 33.19 创建参数模板

平台支持用户指定创建方式创建参数模板，包括复制现有模板和导入模板文件，如下图所示：

![createRedistemplate1](../images/userguide/createRedistemplate1.png)

## 33.20 删除参数模板

平台支持用户对自定义参数模板进行删除操作，可点击参数模板列中操作项的“**删除**”按钮进行操作，如下图所示：

![deleteRedistemplate](../images/userguide/deleteRedistemplate.png)

## 33.21 修改IP

平台支持用户修改Redis的内网IP(VIP)地址。

![updatefsip](../images/userguide/updatefsip.png)

## 33.22 从备份创建

平台支持用户从备份创建Redis，基础信息会展示备份的ID，如下图所示：

![createredisfrombak](../images/userguide/createredisfrombak.png)

## 32.23 查看慢日志信息

慢日志展示执行时长超过一定阈值的 Redis 命令信息，支持用户查看 Redis 慢日志信息，如下图所示：

![redisslowlog](../images/userguide/redisslowlog.png)

## 33.24 删除保护

支持用户开启删除保护功能，用于防止Redis产品被误删除释放，在创建时增加「删除保护」选项，如下图所示：

![redisview4](../images/userguide/redisview4.png)

删除资源时需要将此选项关闭才可以进行删除，如下图所示：

![redisview5](../images/userguide/redisview5.png)

## 33.25 连接信息

支持展示Redis产品客户端连接的网络信息，内网访问支持展示请求地址机器名称，外网仅支持展示请求地址，如下图所示：

![redisview1](../images/userguide/redisview1.png)

## 33.26 参数修改记录

Redis 部分参数新增重启生效字段，如果后续新增了一些不支持动态更改的参数，支持「立刻重启」、「稍后手动重启」两个功能。修改参数后立刻重启后，Redis页面和后台的数据都会更改；修改参数后稍后手动重启，Redis页面数据会更改，后台数据不会更改，如下图所示：

![redisview2](../images/userguide/redisview2.png)

Redis 支持展示参数修改记录，如下图所示：

![redisview3](../images/userguide/redisview3.png)

33.27 Redis 开机关机

Redis 支持实例开关机操作，便于用户做离线配置升降级等操作，如下图所示：

![redisview6](../images/userguide/redisview6.png)

![redisview7](../images/userguide/redisview7.png)

注意事项：
* Redis的备份任务非可用状态，Redis无法关机
* Redis的备份任务是可用状态，Redis关机，定时备份计划会自动暂停，开机后会自动恢复
* Redis的数据传输服务在运行中，Redis无法关机
* Redis主库关机，从库会跟随关机
* Redis主库开机，从库不会跟随开机

