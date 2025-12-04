# 32 MySQL

## 32.1 概览

MySQL是平台提供的一种数据库服务，支持单机版和主备版两种机型，并提供了备份、升级机型、监控等功能。

## 32.2 创建MySQL

平台用户可以通过指定计算集群、存储集群、内存容量、数据盘容量、机型、版本、参数模板、VPC、子网、外网IP、外网安全组、MySQL名称、密码等相关基础信息创建MySQL，可通过导航栏进入"**MySQL**"资源控制台，通过“**创建**”按钮进入向导页面，如下图所示：

![createmysql](../images/userguide/createmysql.png)

1. 选择并配置MySQL的基础配置、网络设置及管理配置信息：

* 计算集群：创建MySQL的计算集群信息，支持选择X86和Aarch集群，并展示所选计算集群的CPU和内存使用率；
* 存储集群：创建MySQL的存储集群信息，可展示所选存储集群的存储用量；
* 容量：MySQL的数据盘容量，可选容量10-64000GB；
* 内存容量：选择创建MySQL的内存容量，支持2G、4G、6G、8G等;
* VPC和子网：创建MySQL时必须选择 VPC 网络和所属子网，即选择要加入的网络及 IP 网段；
* 外网IP：外网 IP 为MySQL提供外网访问服务，MySQL支持最多绑定一个外网IP；
* 机型：支持单机版与主备版；
* 版本：X86类型的计算集群支持MySQL 5.7与MySQL 8.0两个版本；Aarch类型的计算集群只支持MySQL 8.0版本。

2. 选择购买数量和付费方式，确认订单金额并点击“立即购买” 进行MySQL的创建：

- 购买数量：默认支持创建 1 个MySQL；
- 付费方式：选择MySQL的计费方式，支持按月、按年、按时三种方式，可根据需求选择合适的付费方式；
- 合计费用：用户选择MySQL资源按照付费方式的总费用展示；
- 立即购买：点击立即购买后，会返回MySQL资源列表页，在列表页可查看MySQL的创建过程。

## 32.3 查看MySQL列表

平台支持用户查看MySQL列表信息，包括名称、资源ID、状态、机型、集群、存储类型、版本、IP、内存容量、数据盘容量、VPC、子网、安全组、计费方式、项目组、创建时间、过期时间、操作。可通过导航栏进入 "**MySQL**" 资源控制台查看，如下图所示：

![mysqllist](../images/userguide/mysqllist.png)

* 名称：MySQL的名称，可点击进入MySQL详情页；
* 资源ID：MySQL的资源ID，作为全局唯一标识；
* 状态：MySQL的状态，包括启动中、更改配置中、删除中等；
* 机型：MySQL的机型信息，支持单机版和主备版；
* 集群：MySQL所属的计算集群；
* 存储类型：MySQL的存储集群类型；
* 版本：MySQL的版本信息，目前支持MySQL 5.7和MySQL 8.0；
* IP：MySQL的网络信息，包括内网IP和外网IP；
* 内存容量：MySQL创建时选择的内存容量信息，可通过配置升级功能升级内存容量；
* 数据盘容量：MySQL创建时选择的数据盘容量，可通过配置升级功能升级数据盘容量；
* VPC/子网：MySQL内网信息，包括VPC/子网的名称和资源ID；
* 安全组：MySQL绑定外网IP后支持关联安全组；
* 计费方式：MySQL资源的计费方式，支持小时、月、年；
* 项目组：MySQL资源所属的项目组，可通过项目组的转入/转出功能修改关联关系；
* 创建时间/过期时间：MySQL的创建时间和过期时间；
* 操作：MySQL资源的可操作项内容，包括控制台登录、创建从库、删除、续费、重置密码、应用参数模板、配置升级、升级至主备版、修改告警模板、绑定外网IP、解绑外网IP、修改安全组，修改IP等。

## 32.4 查看MySQL概览信息

平台支持用户查看MySQL资源概览信息，包括基本信息、配置信息、监控信息。可通过点击MySQL列表中的名称进入概览页面，如下图所示：

![createmysql](../images/userguide/mysqloverview.png)

## 32.5 MySQL控制台登录

平台支持用户，通过页面登录管理页面对数据库进行管理操作，如下图所示：

![mysqlPMA](../images/userguide/mysqlPMA.png)

## 32.6 MySQL创建从库

平台支持用户对MySQL主库进行创建从库的操作，每个MySQL主库支持最多创建5个从库。可选择计算集群、内存容量、存储集群、数据盘容量、外网IP，可通过MySQL列表中操作项的“**创建从库**”按钮进行操作，如下图所示：

![createmysql](../images/userguide/createmysql3.png)

## 32.7 MySQL续费

平台支持用户对MySQl进行续费操作，可通过MySQL列表中操作项的“**续费**”按钮进行操作，如下图所示：

![createmysql](../images/userguide/mysqlrenew.png)

续费支持更改续费方式和续费时长，更改续费方式只支持由短周期改为长周期，比如从“**小时**”更改为“**月**”。

## 32.8 MySQL重置密码

平台支持用户对MySQl进行重置密码操作，可通过MySQL列表中操作项的“**重置密码**”按钮进行操作，如下图所示：

![createmysql](../images/userguide/mysqlreset.png)

## 32.9 参数配置

平台支持用户对MySQL进行参数配置相关操作，包括修改实例参数、应用参数模板、导入参数文件、导出为模板、导出参数文件。可点击MySQL名称进入详情页，切换到“**参数设置**”页面查看，如下图所示：

![createmysql](../images/userguide/mysqlconfig.png)

### 32.9.1 修改实例参数

平台支持用户对MySQL进行修改实例参数操作，可点击“**修改实例参数**”按钮进行操作，如下图所示：

![createmysql](../images/userguide/editmysql.png)

### 32.9.2 应用参数模板

平台支持用户对MySQL进行应用参数模板操作，可点击“**应用参数模板**”进行操作，也可通过MySQL列表操作项的“**应用参数模板**”进行操作，如下图所示：

![createmysql](../images/userguide/mysqltemplate.png)

### 32.9.3 导入参数文件

平台支持用户对MySQL进行导入参数文件操作，可点击“**导入参数文件**”进行操作，如下图所示：

![createmysql](../images/userguide/mysqlfile1.png)

### 32.9.4 导出为模板

平台支持用户对MySQL参数配置进行导出为模板操作，可点击“**导出为模板**”进行操作，如下图所示：

![createmysql](../images/userguide/mysqltemplate3.png)

导出为模板操作成功后，参数模板列表新增一条模板数据。

### 32.9.5 导出参数文件

平台支持用户对MySQL参数配置进行导出参数文件操作，可点击“**导出参数文件**”进行操作，下载参数配置文件。

![createmysql](../images/userguide/mysqltemplate4.png)

## 32.10 修改配置

平台支持用户对MySQL进行在线配置升级操作，包括内存容量和数据盘容量更改。平台仅支持MySQL容量的扩容，不支持MySQL容量的缩容。可点击MySQL列表中操作项的“**修改配置**”按钮进行操作，如下图所示：

![createmysql](../images/userguide/mysqlupgrade.png)

平台支持 MySQL 实例离线配置升降级，可在实例关机后，同步完成内存容量与数据盘容量的升级，也可单独执行内存容量降级。点击列表中操作项的“**修改配置**”按钮进行操作，如下图所示：

![createmysql](../images/userguide/mysqlupgrade2.png)

注意事项：
* MySQL ARM版本不支持在线升级内存容量
* 从库配置要大于等于主库，MySQL主从库升级需要先升级从库后升级主库，降级需要先降级主库后降级从库

## 32.11 升级主备版

平台支持用户对单机版MySQL进行升级至主备版操作，计算集群、内存容量、存储集群、数据盘容量不可修改。可点击MySQL列表中操作项的“**升级至主备版**”按钮进行操作，如下图所示：

![createmysql](../images/userguide/upgrademysqlmodel.png)

## 32.12 修改告警模板

平台支持用户对MySQL进行修改告警模板操作。可点击MySQL列表中操作项的“**修改告警模板**”按钮进行操作，如下图所示：

![createmysql](../images/userguide/mysqlalarm.png)

## 32.13 网络

平台支持用户查看MySQL的网络信息，包括基本信息和IP列表。可点击MySQL名称进入详情页，切换到“**网络**”页面进行查看，如下图所示：

![createmysql](../images/userguide/mysqlnetwork.png)

### 32.13.1 查看网络列表

平台支持用户查看MySQL的网络列表信息，包括IP、IP ID、IP版本、状态、网络类型、所属网络、是否VIP、带宽、绑定资源、MAC地址、操作，如下图所示：

![createmysql](../images/userguide/mysqlnetworklist.png)

### 32.13.2 绑定外网IP

平台支持用户对MySQL进行绑定外网IP操作，可通过“**绑定**”按钮进行操作，也可通过MySQL列表中操作项的“**绑定外网IP**”进行操作，每个MySQL支持最多绑定一个外网IP。如下图所示：

![createmysql](../images/userguide/mysqlbindeip.png)

外网IP绑定成功后，在MySQL网络列表中新增一条IP数据。

### 32.13.3 解绑外网IP

平台支持用户对已绑定外网IP的MySQL进行解绑外网IP操作，可通过“**解绑**”按钮进行操作，也可通过MySQL列表中操作项的“**解绑外网IP**”进行操作，如下图所示：

![createmysql](../images/userguide/mysqlunbindeip.png)

## 32.14 修改安全组

平台支持用户对已绑定外网IP的MySQL进行修改安全组操作，可通过MySQL列表中操作项的“**修改安全组**”进行操作，如下图所示：

![createmysql](../images/userguide/editmysqlsg.png)

## 32.15 备份管理

### 32.15.1 查看备份管理列表

平台支持用户查看备份管理列表信息，包括资源ID、状态、所属存储池、来源备份计划、保留时间(天)、保存路径、创建时间、更新时间、到期时间、操作。可点击MySQL名称进入详情页，切换到“**备份管理**”页面进行查看，如下图所示：

- 逻辑备份

![mysqlbakview](../images/userguide/mysqlbakview1.png)

- 物理备份

![mysqlbakview](../images/userguide/mysqlbakview2.png)

- 快照备份

![mysqlbakview](../images/userguide/mysqlbakview3.png)

- 增量备份

![mysqlbakview](../images/userguide/mysqlbakview4.png)

### 32.15.2 删除备份

平台支持用户对备份数据进行删除操作，可点击备份列表中操作项的”**删除**“按钮进行操作，也可通过备份列表的“**批量删除**”按钮进行操作，如下图所示：

![createmysql](../images/userguide/deletemysqldbs.png)

### 32.15.3 从备份创建

平台支持用户从备份创建mysql。

![createmysqlfrombak](../images/userguide/createmysqlfrombak.png)


## 32.16 查看操作日志

平台支持用户查看MySQL的操作日志，并可根据操作结果和操作周期进行筛选，可点击MySQL名称进入详情页，切换到“**操作日志**”页面进行查看，如下图所示：

![createmysql](../images/userguide/mysqllog.png)

## 32.17 查看事件

平台支持用户查看MySQL的事件，并可根据事件周期进行筛选，可点击MySQL名称进入详情页，切换到“**事件**”页面进行查看，如下图所示：

![createmysql](../images/userguide/mysqlevent.png)

## 32.18 删除MySQL

平台支持用户对MySQL进行删除操作，删除主库前需先将从库删除。可点击MySQL列表中操作项的“**删除**”按钮进行操作，如下图所示：

![createmysql](../images/userguide/mysqldelete.png)

## 32.19 创建参数模板

平台支持用户创建参数模板，点击”**参数模版**“页面的”**创建**“按钮，可以通过制定参数模版的版本和创建方式，包括MySQL 5.7和MySQL 8.0两种版本，以及“复制现有模板”和“导入模板文件”两种创建方式，如下图所示：

![createmysql](../images/userguide/createmysqltemplate1.png)

## 32.20 应用到实例

平台支持用户对参数模板进行应用到实例操作，可选择要应用到的MySQL资源，点击参数模板列中操作项的“**应用到实例**”按钮进行操作，如下图所示：

![createmysql](../images/userguide/usemysqltemplate.png)

## 32.21 下载参数模板

平台支持用户进行下载参数模板操作，可点击参数模板列中操作项的“**下载**”按钮进行操作。

![createmysql](../images/userguide/usemysqltemplate2.png)

## 32.22 删除参数模板

平台支持用户对自定义参数模板进行删除操作，可点击参数模板列中操作项的“**删除**”按钮进行操作，如下图所示：

![createmysql](../images/userguide/deletemysqltemplate.png)

## 32.23 修改IP

支持用户修改mysql的内网IP(VIP)地址，如下图所示：

![updatefsip](../images/userguide/updatefsip.png)


## 32.24 查看错误日志信息

支持用户查看 MySQL 错误日志信息，如下图所示：

![mysqlerrorlogs](../images/userguide/mysqlerrorlogs.png)

## 32.25 查看慢日志信息

慢日志展示执行时长超过一定阈值的 MySQL 命令信息，支持用户查看 MySQL 慢日志信息，如下图所示：

![mysqlslowlogrecords](../images/userguide/mysqlslowlogrecords.png)

## 32.26 删除保护

支持用户开启删除保护功能，用于防止MySQL产品被误删除释放，在创建时增加「删除保护」选项，如下图所示：

![mysqlremoveprotection](../images/userguide/mysqlremoveprotection.png)

删除资源时需要将此选项关闭才可以进行删除，如下图所示：

![mysqlupgraderemoveprotection](../images/userguide/mysqlupgraderemoveprotection.png)

## 32.27 连接信息

支持展示MySQL产品客户端连接的网络信息，内网访问支持展示请求地址机器名称，外网仅支持展示请求地址，如下图所示：

![mysqlview1](../images/userguide/mysqlview1.png)

## 32.28 参数修改记录

MySQL 修改部分不支持动态更改的参数，支持「立刻重启」、「稍后手动重启」两个功能。修改参数后立刻重启后，MySQL页面和后台的数据都会更改；修改参数后稍后手动重启，MySQL页面数据会更改，后台数据不会更改，如下图所示：

![mysqlview2](../images/userguide/mysqlview2.png)

MySQL 支持展示参数修改记录，如下图所示：

![mysqlview3](../images/userguide/mysqlview3.png)

## 32.29 SQL审计

MySQL 支持数据库审计能力，记录对数据库的访问及 SQL 语句执行情况，帮助企业进行风险控制，提高数据安全等级，支持审计日志事件的选择，如下图所示：

* QUERY：普通 SQL 查询（不区分 DDL/DML/DCL，所有 SQL 都可能被记录）
* QUERY_DDL：结构变更类语句（CREATE TABLE、ALTER TABLE等）
* QUERY_DML：数据修改类语句（INSERT INTO、UPDATE等）
* QUERY_DCL：权限控制类语句（CREATE USER、DROP USER等）

![mysqlview4](../images/userguide/mysqlview4.png)

![mysqlview5](../images/userguide/mysqlview5.png)

MySQL 支持过滤掉不需要展示的语句，如下图所示：

![mysqlview6](../images/userguide/mysqlview6.png)

关于 SQL审计日志大小、SQL审计日志轮转数量支持在管理员页面编辑修改，如下图所示：

![mysqlview7](../images/userguide/mysqlview7.png)

32.30 MySQL 开机关机

MySQL 支持实例开关机操作，便于用户做离线配置升降级等操作，如下图所示：

![mysqlview8](../images/userguide/mysqlview8.png)

![mysqlview9](../images/userguide/mysqlview9.png)

注意事项：
* MySQL的备份任务非可用状态，MySQL无法关机
* MySQL的备份任务是可用状态，MySQL关机，定时备份计划会自动暂停，MySQL开机后会自动恢复
* MySQL的数据传输服务在运行中，MySQL无法关机
* MySQL主库关机，从库会跟随关机
* MySQL主库开机，从库不会跟随开机
* MySQL从库处于关机状态，MySQL相关联的增量备份状态变为暂停，主从状态开机后，增量备份自动恢复



