# 15 高可用VIP

## 15.1 高可用VIP 简介

### 15.1.1 概述

高可用VIP（ High available Virtual IP Address ，简称 HAVIP ），高可用虚拟IP地址，是归属于VPC内某个子网内的可漂移内网IP，用户可将HAVIP与高可用服务结合，以便在服务出现故障时进行服务入口的漂移，以实现服务的高可用。

### 15.1.2 工作原理

HaVIP作为一个不绑定特定设备的浮动IP，通常和高可用软件（keepalived、heartbeat、Failover Cluster）配合使用，用于搭建高可用主备集群，比如HA负载均衡、主备版数据库。这里以keepalived为例介绍下HaVIP的工作原理

![havip](../images/userguide/havip.png)

* master和slave均安装 keepalived，配置 从控制台申请出来的HaVIP 为 VRRP VIP，分别设置优先级（priority值）；
* Keepalived 中的 VRRP 协议通过对比两台虚拟机的初始优先级大小，选举出 Master 服务器；
* Master 服务器向外发送 ARP 报文，宣告 VIP，实现 VIP 和 MAC 的地址映射更新（arp缓存)；
* 此时真正对外提供服务的服务器为 Master 服务器，通信的内网 IP 为 HaVIP ；
* Master 服务器周期性发送 VRRP 报文给 slave 服务器。如果 Master 服务器异常，Backup 服务器在一定时间内没有收到 VRRP 报文，则会将自己设置为 Master，并对外发送 ARP更新（GARP），报文携带自己的 MAC 地址；
* 此时 slave 服务器 将作为Master服务器对外提供通信服务，外部访问的报文将转发至 slave 处理，直至实现了realserver的切换；

## 15.2 申请高可用VIP

云平台用户可通过 API 接口或控制台创建高可用VIP，用于服务的高可用，创建高可用VIP前需保证账户至少拥有一个 VPC 网络和子网。通过导航栏进入虚拟机控制台，切换至【VIP】管理页面，点击“新建”按钮进入高可用VIP创建向导弹窗，如图所示：

![createVIP](../images/userguide/createVIP2.png)

* 名称/备注：申请高可用VIP 的名称和备注，申请时必须指定名称。
* VIP类型：内网和外网
* 网络设置：高可用VIP的所属网络，创建时必须指定。
* IP 版本：内网 VIP 支持 IPv4，外网 VIP 支持 IPv4 和 IPv6。
* IP 地址：用户手动指定 IP 地址申请HAVIP，指定的 IP 地址必须在所选网段的 IP 范围内。若手动指定的 IP 地址已被使用，则会弹出占用提示。
* 关联虚拟机：用户可以选择与VIP同vpc同子网下的虚拟机，并绑定HAVIP。
  * 单VIP可绑定虚拟机默认不超过3台，数量限制可以联系运维人员动态的调整。
  * 一台虚拟机只能绑定5个vip
* 确认创建：点击确认后，会返回 HAVIP 资源列表页，在列表页可查看 HAVIP 的资源状态，通常创建成功后会显示“可用”的状态，如果因为某些原因没有创建成功会显示”失败“的状态。

## 15.3 查看高可用VIP

通过导航栏进入虚拟机控制台，切换至VIP管理页面可查看高可用VIP资源的列表及相关信息，包括高可用VIP的名称备注、资源 ID、状态、vpc子网、IP 地址、关联资源、项目组、创建时间及操作项，如下图所示：

![VIPList](../images/userguide/VIPList.png)

* 名称/备注：申请高可用VIP 的名称和描述。
* 资源ID：高可用VIP的全局唯一标识符。
* vpc子网/外网网络：高可用VIP的所属网络
* IP 地址与带宽：用户手动指定 IP 地址申请HAVIP，指定的 IP 地址必须在所选网段的 IP 范围内。
* 关联资源：用户可以选择与VIP同vpc同子网下的虚拟机，并绑定HAVIP。
* 创建时间：当前网卡的创建时间。
* 状态：高可用VIP当前的状态，包括可用、失败、删除中等状态

列表上的操作项是指对单个HAVIP的操作，包括创建、更新、删除等，可通过搜索框对HAVIP列表进行搜索和筛选，支持模糊搜索。

为方便租户对HAVIP资源的统计及维护，平台支持下载当前用户所拥有的所有HAVIP资源列表信息为 Excel 表格；同时支持对HAVIP进行批量删除操作。

## 15.4 更新高可用VIP

修改HAVIP的名称和备注，在任何状态下均可进行操作。可通过VIP列表页面每个HAVIP名称右侧的“编辑”按钮进行修改；

![updateVIP1](../images/userguide/updateVIP1.png)

更新关联虚拟机，可以替换、删除、新增虚拟机，关联虚拟机数量默认是3台，如图所示：

![updateVIP2](../images/userguide/updateVIP2.png)
VIP 绑定虚拟机数量的限制可以联系运维人员动态的调整。

## 15.5 删除高可用VIP

支持用户删除高可用VIP资源，可支持删除【可用】【失败】状态的高可用VIP。删除网卡后，会自动解绑与之关联的虚拟机。用户可通过VIP列表进行高可用VIP的删除操作，支持批量删除。

![deleteVIP](../images/userguide/deleteVIP.png)

## 15.6 使用VIP

### 15.6.1 Linux 虚拟机使用VIP

1、查看 keepalived 软件包版本号是否符合要求。

```
yum list keepalived
```

2、使用 yum 方式安装软件包。

```
yum install -y keepalived
```

3、配置 keepalived，绑定高可用 VIP 到主备云服务器，登录主节点云服务器 HAVIP-01，执行 `vim /etc/keepalived/keepalived.conf` ，修改相关配置。

```
! Configuration File for keepalived
global_defs {
    script_user root
    enable_script_security
}
# 注意主备参数选择
vrrp_instance VI_VPN {
    state MASTER # 设置状态均为“主“，备节点为 BACKUP
    nopreempt # 设置非抢占模式
    advert_int 1
    priority 110
    interface ens5
    virtual_router_id 111 
    authentication {
        auth_type PASS
        auth_pass bless
    }
    unicast_src_ip 10.0.6.9 # 设置本机内网 IP 地址
    unicast_peer {
        10.0.6.10 # 对端设备的 IP 地址
    }
    virtual_ipaddress {
        10.0.6.19/24 dev ens5 # 设置高可用虚拟 VIP
    }
    virtual_routes { # 自定义路由规则
    }
}
```

4、退出编辑状态，输入 `:wq!`保存并退出，重启 keepalived 进程使配置生效。

```
systemctl restart keepalived
```

### 15.6.2 Windows 虚拟机使用VIP

1、配置主机域环境（关联VIP的虚拟机都需要配置）

登录windows虚拟机，打开“此电脑-系统-高级系统设置”；点击“更改”——设置“计算机名”——点击“其他”——设置“主 DNS 后缀”——点击“确定”，如下：

![windowsvip1](../images/userguide/windowsvip1.png)

操作重启虚拟机后生效，如下：

![windowsvip2](../images/userguide/windowsvip2.png)

![windowsvip3](../images/userguide/windowsvip3.png)

2、安装故障转移群集（关联VIP的虚拟机都需要操作）

操作服务器管理器-添加角色和功能，如下：

![windowsvip4](../images/userguide/windowsvip4.png)

直接点击下一步，到“功能”项，选择“故障转移群集”，如下：

![windowsvip5](../images/userguide/windowsvip5.png)

![windowsvip6](../images/userguide/windowsvip6.png)

点击“安装”，如下：

![windowsvip7](../images/userguide/windowsvip7.png)

3、配置DNS（选择关联VIP的其中一台虚拟机）

打开hosts文件，如下：

![windowsvip8](../images/userguide/windowsvip8.png)

将加入vip的其他windows虚拟机的ip和计算机全名（例如：vip.test.com）信息配置到文件中，如下：

![windowsvip9](../images/userguide/windowsvip9.png)

4、配置故障转移群集（操作的虚拟机与 3 保持一致）

点击“创建群集”，如下：

![windowsvip10](../images/userguide/windowsvip10.png)

点击下一步，如下：

![windowsvip11](../images/userguide/windowsvip11.png)

![windowsvip12](../images/userguide/windowsvip12.png)

已加入hosts的主机可以通过域名添加，本机可以通过主机名添加。

添加完成后，一直点击“下一步”，完成“验证告警”，如下：

![windowsvip13](../images/userguide/windowsvip13.png)

![windowsvip14](../images/userguide/windowsvip14.png)

“验证告警”完成后，回到“创建群集”流程；输入“群集名称”，并在地址栏输入VIP的内网地址，如下：

![windowsvip15](../images/userguide/windowsvip15.png)

继续点击“下一步”，等待创建群集完成，如下：

![windowsvip16](../images/userguide/windowsvip16.png)

![windowsvip17](../images/userguide/windowsvip17.png)

5、验证

可以看到新建的群集，如下：

![windowsvip18](../images/userguide/windowsvip18.png)

登录主服务器对应的测试机，查看网络信息，可以看到VIP的ip，如下：

![windowsvip19](../images/userguide/windowsvip19.png)

登录未关联VIP的测试机，“ping VIP的ip”可以ping通，如下：

![windowsvip20](../images/userguide/windowsvip20.png)

切换为其他节点（即主服务切换为其他主机），如下：

![windowsvip21](../images/userguide/windowsvip21.png)

![windowsvip22](../images/userguide/windowsvip22.png)

登录新选择的节点，可以查看到VIP的ip,如下：

![windowsvip23](../images/userguide/windowsvip23.png)

登录切换前的节点，查看不到VIP的ip，如下：

![windowsvip24](../images/userguide/windowsvip24.png)

新节点选择后，等待服务器状态为联机后，测试机可以正常ping通，如下：

![windowsvip25](../images/userguide/windowsvip25.png)

![windowsvip26](../images/userguide/windowsvip26.png)
