# BILL

## 获取订单信息 - DescribeOrder

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **UserAccountID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [CreateUser DescribeUser DescribeMember] | No |
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | No |
| **BeginTime** | integer | 开始时间 | [1751958000] | [] | **Yes** |
| **EndTime** | integer | 结束时间 | [1754550000] | [] | **Yes** |
| **Limit** | integer | 分页大小，指定返回结果的数量。用于分页查询。 | [25] | [] | No |
| **Offset** | integer | 分页偏移量，指定返回结果的起始位置。用于分页查询，默认为0。 | [0] | [] | No |
| **ProjectIDs** | array<string> | 项目组ID列表，查看多个项目组的订单信息，不填写默认查看所有项目组的订单 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*Order*](#Order)> | 返回开始时间到结束时间租户所有的订单信息列表 | [] |
| **TotalCount** | integer | 返回开始时间到结束时间租户所有的订单数量 | [] |

### 数据模型

#### Order

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **OrderNO** | string | 订单编号，订单信息的唯一标识 | [order-xf3sxw5u4rm4sg] |
| **CompanyID** | integer | 租户ID，订单所属的租户ID | [200000233] |
| **CompanyName** | string | 租户名称，订单所属的租户名称，与CompanyID对应 | [ucloud] |
| **Email** | string | 租户邮箱，订单所属的租户邮箱，与CompanyID对应 | [ucloud@example.cn] |
| **OrderType** | string | 订单类型，订单的类型，例如：充值BUY、续费RENEW、退款REFUND等 | [BUY] |
| **ChargeType** | string | 计费类型, 自动为按小时计费，取值范围：hour, month, year | [hour month year] |
| **AmountTotal** | double | 总金额，订单的总金额，单位为元 | [100.0] |
| **AmountReal** | double | 外部账户金额，订单的外部账户金额，单位为元 | [100.0] |
| **AmountFree** | double | 内部账户金额，订单的内部账户金额，单位为元 | [100.0] |
| **Region** | string | 所属的地域，订单资源所属的地域 | [pro2133] |
| **RegionAlias** | string | 地域别名，订单资源所属的地域别名 | [测试地域] |
| **ResourceID** | string | 订单对应所创建资源的资源ID，例如该订单是创建磁盘，该字段为磁盘ID | [disk-12saf3****sf456] |
| **CreateTime** | integer | 订单的创建时间 | [1753867542] |
| **OrderDetails** | array<[*OrderDetail*](#OrderDetail)> | 订单详情 | [] |
| **ResourceName** | string | 订单对应所创建资源的资源名称 | [测试磁盘] |

#### OrderDetail

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ResourceID** | string | 订单对应的资源ID，例如该订单是创建磁盘，该字段为磁盘ID | [disk-12saf3****sf456] |
| **Factor** | integer | 因子 | [] |
| **Quantity** | integer | 操作资源的数量，例如该订单是创建磁盘，该字段为创建的磁盘数量 | [1] |
| **Property** | string | 类型，例如该订单是关于磁盘的操作，该字段为磁盘的集群类型 | [StorageSetOUWI] |
| **ProductName** | string | 产品名称，例如该订单是关于磁盘的操作，该字段为磁盘的产品名称 | [磁盘] |
| **Unit** | string | 资源单位，例如该订单是关于磁盘的操作，该字段为磁盘的单位 | [GB] |

## 获取价格信息 - DescribePrice

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [CreateUser DescribeUser DescribeMember] | No |
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **ProductIDs** | array<string> | 产品ID列表，查询指定产品的价格信息，不填写默认查询所有产品 | [SECURITY_WAF VM DISK] | [ListProductResources] | No |
| **SetTypes** | array<string> | 集群类型列表，查询指定集群类型的价格信息，不填写默认查询所有集群类型 | [ComputeSetPECR StorageSetOUWI] | [DescribeStorageSet] | No |
| **Keyword** | string | 关键字搜索，支持按产品名称、集群类型等进行模糊搜索 | [Web应用防火墙] | [] | No |
| **Limit** | integer | 分页大小，每页返回的记录数量 | [25] | [] | No |
| **Offset** | integer | 分页偏移，从第几条记录开始返回 | [0] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*Price*](#Price)> | 价格信息列表，包含查询条件范围内的所有价格信息 | [] |
| **TotalCount** | integer | 价格信息总数量，用于分页计算 | [] |

### 数据模型

#### BillPriceInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **LowerMultiple** | integer | 最小范围，价格阶梯的最小数量 | [0] |
| **UpperMultiple** | integer | 最大范围，价格阶梯的最大数量 | [0] |
| **Price** | double | 价格，该阶梯范围内的单价，单位为元 | [1000] |

#### Price

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ProductID** | string | 产品ID，产品的唯一标识 | [SECURITY_WAF] |
| **Product** | string | 产品名称，产品的显示名称 | [Web应用防火墙可管理资产数] |
| **Region** | string | 地域，价格所属的地域 | [pro2133] |
| **RegionAlias** | string | 地域别名，地域的显示名称 | [pro2133] |
| **SetType** | string | 集群类型，资源所属的集群类型 | [ComputeSetPECR] |
| **SetTypeAlias** | string | 集群别名，集群类型的显示名称 | [ComputeSetPECR] |
| **ChargeType** | string | 计费类型，资源的计费方式，取值范围：Hour, Month, Year | [Month] |
| **ChargeRule** | string | 计费规则，详细的计费规则说明 | [GRADIENT] |
| **Infos** | array<[*BillPriceInfo*](#BillPriceInfo)> | 账单价格信息，包含不同数量阶梯的价格信息 | [] |
| **DiscountInfos** | array<[*BillPriceInfo*](#BillPriceInfo)> | 折扣信息，包含不同数量阶梯的折扣价格信息 | [] |
| **Discount** | double | 折扣，折扣比例，100表示无折扣 | [100] |
| **Unit** | string | 单位，价格的计量单位 | [个] |
| **CreateTime** | integer | 创建时间，价格信息的创建时间戳 | [1753790301] |
| **UpdateTime** | integer | 更新时间，价格信息的最后更新时间戳 | [1753790301] |

## 获取充值信息 - DescribeRecharge

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **RechargeType** | string | 充值类型，充值的类型分类，内部充值或外部充值，1标识外部充值，2标识内部充值 | [2] | [] | No |
| **FromType** | string | 来源类型，充值的来源渠道，取值范围：FROM_TYPE_NONE, ALIPAY, OFFLINE, SINPAY, INNER, WECHAT_PAY | [ALIPAY] | [] | No |
| **BeginTime** | integer | 开始时间，查询充值记录的开始时间戳 | [1751958000] | [] | **Yes** |
| **EndTime** | integer | 结束时间，查询充值记录的结束时间戳 | [1754550000] | [] | **Yes** |
| **CompanyID** | integer | 租户ID，查询指定租户的充值记录 | [200000233] | [CreateUser DescribeUser DescribeMember] | No |
| **UserAccountID** | integer | 用户账号ID，和租户ID相同，查询指定用户账号的充值记录 | [200000233] | [DescribeUser, DescribeMember] | No |
| **Limit** | integer | 分页大小，每页返回的记录数量 | [25] | [] | No |
| **Offset** | integer | 分页偏移，从第几条记录开始返回 | [0] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*Recharge*](#Recharge)> | 充值信息列表，包含查询条件范围内的所有充值记录 | [] |
| **TotalCount** | integer | 充值记录总数量，用于分页计算 | [] |

### 数据模型

#### Recharge

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TransactonNO** | string | 交易号，充值交易的唯一标识 | [trade-pr1pv7knuiz7zj] |
| **CompanyID** | integer | 租户ID，充值所属的租户ID | [200000233] |
| **CompanyName** | string | 租户名称，充值所属的租户名称，与CompanyID对应 | [testName] |
| **Email** | string | 租户邮箱，充值所属的租户邮箱，与CompanyID对应 | [ucloud@example.cn] |
| **AccountID** | integer | 账号ID，充值操作的账号ID | [200000233] |
| **AccountName** | string | 账号名称，充值操作的账号名称，与AccountID对应 | [test_name] |
| **RechargeType** | string | 充值类型，充值的类型分类 | [FREE] |
| **TradeAmount** | double | 充值金额，本次充值的金额，单位为元 | [5000] |
| **FromType** | string | 来源类型，充值的来源渠道，例如：ALIPAY, WECHAT_PAY, INNER等 | [INNER] |
| **SerialNo** | string | 序列号，充值的序列号或凭证号 | [serial-5exc2****vgn0d] |
| **CreateTime** | integer | 创建时间，充值记录的创建时间戳 | [1753841388] |
| **UpdateTime** | integer | 更新时间，充值记录的最后更新时间戳 | [1753841388] |

## 获取交易记录 - DescribeTransaction

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | No |
| **TransactionType** | string | 交易类型，0:无交易类型，1:扣费Charge，2:充值Recharge，3:退款Refund，4:提现Withdraw，5:免费充值FreeRecharge，6:免费提现FreeWithdraw | [2] | [] | No |
| **BeginTime** | integer | 开始时间，查询交易记录的开始时间戳 | [1751958000] | [] | **Yes** |
| **EndTime** | integer | 结束时间，查询交易记录的结束时间戳 | [1754550000] | [] | **Yes** |
| **UserAccountID** | integer | 用户账号ID，和租户ID相同，查询指定用户账号的交易记录 | [200000238] | [] | No |
| **Limit** | integer | 分页大小，每页返回的记录数量 | [10] | [] | No |
| **Offset** | integer | 分页偏移，从第几条记录开始返回 | [0] | [] | No |
| **ProjectIDs** | array<string> | 项目组ID列表，查询指定项目组的交易记录，不填写默认查询所有项目组 | [project-123 project-456] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*Transaction*](#Transaction)> | 交易信息列表，包含查询条件范围内的所有交易记录 | [] |
| **TotalCount** | integer | 交易记录总数量，用于分页计算 | [] |

### 数据模型

#### Transaction

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TransactionNO** | string | 交易号，交易记录的唯一标识 | [trade-4fnns****et31] |
| **TradeNO** | string | 交易号，交易的业务流水号 | [order-p5tmw****tokl] |
| **Region** | string | 地域，交易所属的地域 | [pro2133] |
| **RegionAlias** | string | 地域别名，地域的显示名称 | [pro2133] |
| **CompanyID** | integer | 租户ID，交易所属的租户ID | [200000233] |
| **AccountID** | integer | 账号ID，交易操作的账号ID | [200000233] |
| **CompanyName** | string | 租户名称，交易所属的租户名称，与CompanyID对应 | [company_name] |
| **AccountName** | string | 账号名称，交易操作的账号名称，与AccountID对应 | [account_name] |
| **Email** | string | 租户邮箱，交易所属的租户邮箱，与CompanyID对应 | [ucloud@example.cn] |
| **TransactionType** | integer | 交易类型，交易的类型编码 | [1] |
| **ExpenseAmount** | double | 交易金额，支出的金额，单位为元 | [0] |
| **IncomeAmount** | double | 收入金额，收入的金额，单位为元 | [0] |
| **BalanceSurplus** | double | 余额，交易后的账户余额，单位为元 | [0] |
| **FreeBalanceSurplus** | double | 内部账号余额，交易后的内部账号余额，单位为元 | [500000] |
| **State** | integer | 状态，交易的状态编码 | [0] |
| **CreateTime** | integer | 创建时间，交易记录的创建时间戳 | [1754043643] |

## 充值 - Recharge

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，指定租户的充值 | [200000233] | [CreateUser DescribeUser DescribeMember] | **Yes** |
| **RechargeType** | integer | 交易类型，0:无交易类型，1:扣费Charge，2:充值Recharge，3:退款Refund，4:提现Withdraw，5:免费充值FreeRecharge，6:免费提现FreeWithdraw | [2] | [] | **Yes** |
| **FromType** | string | 来源类型，充值的来源渠道，取值范围：FROM_TYPE_NONE, ALIPAY, OFFLINE, SINPAY, INNER, WECHAT_PAY | [ALIPAY] | [] | No |
| **SerialNo** | string | 序列号，充值的序列号或凭证号 | [serial-5exc2****vgn0d] | [DescribeRecharge] | No |
| **Amount** | double | 充值金额，要充值的金额，单位为元 | [100.0] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TransactonNo** | string | 交易号，充值成功后返回的交易号 | [trade-q42yd****cugql] |


## 续费 - RenewResource

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于唯一标识租户，便于租户管理 | [200000233] | [CreateCompany GetCompany] | **Yes** |
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **ResourceID** | string | 资源ID，要续费的资源唯一标识 | [vm-12saf3****sf456] | [] | **Yes** |
| **Quantity** | integer | 数量，续费的数量 | [1] | [] | No |
| **ChargeType** | string | 计费类型，用于指定计费模式。取值范围：hour（按小时计费）、month（按月计费）、year（按年计费）。 | [hour] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TransactonNo** | string | 交易号，续费操作生成的交易号 | [trade-pr1pv7knuiz7zj] |


## 更新价格 - UpdatePrice

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **ProductID** | string | 产品ID，用于唯一标识产品，便于产品管理 | [DISK] | [ListProductResources] | **Yes** |
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **ChargeType** | string | 计费类型，用于指定计费模式。取值范围：hour（按小时计费）、month（按月计费）、year（按年计费）。 | [hour] | [] | **Yes** |
| **SetType** | string | 集群类型 | [StorageSetOUWI] | [DescribeStorageSet] | **Yes** |
| **PriceRuleInfo** | string | 价格规则信息，标明上下限和价格 | [{"LowerMultiple":0,"UpperMultiple":0,"Price":0.6667}] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 更新折扣 - UpdateDiscount

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于唯一标识租户，便于租户管理 | [200000233] | [CreateCompany GetCompany] | **Yes** |
| **ProductID** | string | 产品ID，用于唯一标识产品，便于产品管理 | [DISK] | [ListProductResources] | **Yes** |
| **Region** | string | 地域ID，地域，一批可共享的物理资源使用集合 | [develop] | [DescribeRegion] | **Yes** |
| **ChargeType** | string | 计费类型，用于指定计费模式。取值范围：hour（按小时计费）、month（按月计费）、year（按年计费）。 | [hour] | [] | **Yes** |
| **SetType** | string | 集群类型 | [StorageSetOUWI] | [DescribeStorageSet] | **Yes** |
| **Discount** | double | 折扣 | [] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|


## 获取账单总览 - DescribeBillOverView

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **Dimension** | string | 合并维度，统计数据的合并维度(product,region,project,order,charge) | [product] | [] | **Yes** |
| **CycleUnit** | string | 统计周期单位，时间统计的周期单位(hour,day,week,month) | [month] | [] | **Yes** |
| **CycleCount** | integer | 统计周期数量，要统计的周期数量 | [6] | [] | **Yes** |
| **Regions** | array<string> | 地域列表，查询指定地域的账单概览 | [pro2133 cn-bj2] | [DescribeRegion] | No |
| **CompanyIDs** | array<int32> | 租户ID列表，查询指定租户的账单概览 | [200000233 200000234] | [DescribeUser] | No |
| **ProjectIDs** | array<string> | 项目组ID列表，查询指定项目组的账单概览 | [project-123 project-456] | [ListProjects] | No |
| **OrderTypes** | array<string> | 订单类型列表，查询指定订单类型的账单详情(BUY,RENEW,UPGRADE,DOWNGRADE,REFUND,ORDER_TYPE_NONE) | [BUY RENEW] | [] | No |
| **ProductTypes** | array<string> | 产品类型列表，查询指定产品类型的账单概览 | [VM DISK] | [ListProductResources] | No |
| **ChargeTypes** | array<string> | 计费类型列表，查询指定计费类型的账单概览，取值范围：hour（按小时计费）、month（按月计费）、year（按年计费）。 | [Hour Month] | [] | No |
| **BeginTime** | integer | 开始时间，查询的开始时间戳 | [1751958000] | [] | No |
| **EndTime** | integer | 结束时间，查询的结束时间戳 | [1754550000] | [] | No |
| **Sort** | string | 排序方式，升序或降序 | [Ascending] | [] | No |
| **SortBy** | string | 排序字段，按哪个字段排序 | [Amount] | [] | No |
| **Limit** | integer | 分页大小，每页返回的记录数量 | [10] | [] | No |
| **Offset** | integer | 分页偏移，从第几条记录开始返回 | [0] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*BillOverView*](#BillOverView)> | 详情，账单概览信息列表 | [] |
| **TotalCount** | integer | 总数量，符合条件的记录总数 | [50] |
| **TotalAmount** | double | 总金额，所有记录的总金额，单位为元 | [15000.75] |
| **TotalAmountFree** | double | 内部账号总金额，所有记录的内部账号总金额，单位为元 | [5000.25] |
| **TotalAmountReal** | double | 外部账号总金额，所有记录的外部账号总金额，单位为元 | [10000.50] |

### 数据模型

#### BillOverView

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Time** | integer | 时间，统计周期的时间戳 | [1753867542] |
| **CycleUnit** | string | 统计周期单位，时间周期的单位，如hour、day、month等 | [day] |
| **CycleAmount** | double | 统计周期金额，该周期内的总金额，单位为元 | [2500.75] |
| **CycleAmountFree** | double | 统计周期内部账号金额，该周期内的内部账号金额，单位为元 | [800.25] |
| **CycleAmountReal** | double | 统计周期外部账号金额，该周期内的外部账号金额，单位为元 | [1700.50] |
| **CycleCount** | integer | 统计周期数量，该周期内的记录数量 | [15] |
| **CycleDetails** | array<[*BillOverViewDetail*](#BillOverViewDetail)> | 详情，该周期内各维度的详细统计信息 | [] |

#### BillOverViewDetail

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Dimension** | string | 维度，统计的维度类型，如product、region等 | [product] |
| **Name** | string | 名称，维度对应的名称 | [云主机] |
| **Amount** | double | 金额，该维度的总金额，单位为元 | [1500.50] |
| **AmountFree** | double | 内部账号金额，该维度的内部账号金额，单位为元 | [500.00] |
| **AmountReal** | double | 外部账号金额，该维度的外部账号金额，单位为元 | [1000.50] |

## 获取账单详情 - DescribeBillDetail

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **BeginTime** | integer | 开始时间，查询的开始时间戳 | [1751958000] | [] | **Yes** |
| **EndTime** | integer | 结束时间，查询的结束时间戳 | [1754550000] | [] | **Yes** |
| **Regions** | array<string> | 地域列表，查询指定地域的账单概览 | [pro2133 cn-bj2] | [DescribeRegion] | No |
| **CompanyIDs** | array<int32> | 租户ID列表，查询指定租户的账单概览 | [200000233 200000234] | [DescribeUser] | No |
| **ProjectIDs** | array<string> | 项目组ID列表，查询指定项目组的账单概览 | [project-123 project-456] | [ListProjects] | No |
| **OrderTypes** | array<string> | 订单类型列表，查询指定订单类型的账单详情(BUY,RENEW,UPGRADE,DOWNGRADE,REFUND,ORDER_TYPE_NONE) | [BUY RENEW] | [] | No |
| **ProductTypes** | array<string> | 产品类型列表，查询指定产品类型的账单概览 | [VM DISK] | [ListProductResources] | No |
| **ChargeTypes** | array<string> | 计费类型列表，查询指定计费类型的账单概览，取值范围：hour（按小时计费）、month（按月计费）、year（按年计费）。 | [Hour Month] | [] | No |
| **Sort** | string | 升序排序 | [Ascending] | [] | No |
| **SortBy** | string | 排序指标，按时间或价格排序(TIME,PRICE) | [TIME] | [] | No |
| **Limit** | integer | 分页大小，每页返回的记录数量 | [10] | [] | No |
| **Offset** | integer | 分页偏移，从第几条记录开始返回 | [0] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Infos** | array<[*BillDetailInfo*](#BillDetailInfo)> | 详情，账单详情信息列表 | [] |
| **TotalCount** | integer | 总数量，符合条件的记录总数 | [200] |

### 数据模型

#### BillDetailInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Region** | string | 地域，账单详情所属的地域 | [pro2133] |
| **RegionAlias** | string | 地域别名，地域的显示名称 | [测试地域] |
| **ProjectID** | string | 项目组ID，账单详情所属的项目组ID | [project-123] |
| **ProjectName** | string | 项目组名称，项目组的显示名称 | [测试项目] |
| **CompanyID** | integer | 租户ID，账单详情所属的租户ID | [200000233] |
| **CompanyName** | string | 租户名称，租户的显示名称 | [ucloud] |
| **CompanyEmail** | string | 租户邮箱，租户的邮箱地址 | [ucloud@example.cn] |
| **OrderNO** | string | 订单号，关联的订单编号 | [order-xf3sxw5u4rm4sg] |
| **OrderType** | string | 订单类型，订单的类型 | [BUY] |
| **TransactionNO** | string | 交易号，关联的交易编号 | [trade-4fnnsmyn8egt31] |
| **TransactionType** | string | 交易类型，交易的类型 | [REFUND] |
| **ProductType** | string | 产品类型，产品的类型 | [VM] |
| **ChargeType** | string | 计费类型，计费的方式 | [Hour] |
| **ResourceID** | string | 资源ID，资源的唯一标识 | [vm-12saf3****sf456] |
| **StartTime** | integer | 开始时间，计费的开始时间戳 | [1753867542] |
| **EndTime** | integer | 结束时间，计费的结束时间戳 | [1753953942] |
| **CreateTime** | integer | 创建时间，账单详情的创建时间戳 | [1753867542] |
| **Amount** | double | 金额，账单详情的总金额，单位为元 | [100.50] |
| **AmountReal** | double | 外部账号金额，外部账号支付的金额，单位为元 | [60.30] |
| **AmountFree** | double | 内部账号金额，内部账号支付的金额，单位为元 | [40.20] |
| **ResourceName** | string | 资源名称，资源的显示名称 | [测试云主机] |

## 获取资源账单详情 - DescribeBillResource

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **BeginTime** | integer | 开始时间，查询的开始时间戳 | [1751958000] | [] | **Yes** |
| **EndTime** | integer | 结束时间，查询的结束时间戳 | [1754550000] | [] | **Yes** |
| **CycleUnit** | string | 统计周期单位，时间统计的周期单位(hour,day,week,month) | [month] | [] | No |
| **CycleCount** | integer | 统计周期数量，要统计的周期数量 | [1] | [] | No |
| **Regions** | array<string> | 地域列表，查询指定地域的账单概览 | [pro2133 cn-bj2] | [DescribeRegion] | No |
| **CompanyIDs** | array<int32> | 租户ID列表，查询指定租户的账单概览 | [200000233 200000234] | [DescribeUser] | No |
| **ProjectIDs** | array<string> | 项目组ID列表，查询指定项目组的账单概览 | [project-123 project-456] | [ListProjects] | No |
| **OrderTypes** | array<string> | 订单类型列表，查询指定订单类型的账单详情(BUY,RENEW,UPGRADE,DOWNGRADE,REFUND,ORDER_TYPE_NONE) | [BUY RENEW] | [] | No |
| **ProductTypes** | array<string> | 产品类型列表，查询指定产品类型的账单概览 | [VM DISK] | [ListProductResources] | No |
| **ChargeTypes** | array<string> | 计费类型列表，查询指定计费类型的账单概览，取值范围：hour（按小时计费）、month（按月计费）、year（按年计费）。 | [Hour Month] | [] | No |
| **Sort** | string | 排序方式，升序或降序 | [Ascending] | [] | No |
| **SortBy** | string | 排序字段，按哪个字段排序 | [Amount] | [] | No |
| **Limit** | integer | 分页大小，每页返回的记录数量 | [10] | [] | No |
| **Offset** | integer | 分页偏移，从第几条记录开始返回 | [0] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 总数量，符合条件的记录总数 | [100] |
| **Infos** | array<[*BillResourceInfo*](#BillResourceInfo)> | 详情，资源账单信息列表 | [] |

### 数据模型

#### BillResourceInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **ResourceID** | string | 资源ID，资源的唯一标识 | [vm-12saf3****sf456] |
| **Amount** | double | 金额，该资源的总金额，单位为元 | [500.50] |
| **AmountFree** | double | 内部账号金额，该资源的内部账号金额，单位为元 | [200.00] |
| **AmountReal** | double | 外部账号金额，该资源的外部账号金额，单位为元 | [300.50] |
| **Region** | string | 地域，资源所属的地域 | [pro2133] |
| **RegionAlias** | string | 地域别名，地域的显示名称 | [测试地域] |
| **ProjectID** | string | 项目组ID，资源所属的项目组ID | [project-123] |
| **ProjectName** | string | 项目组名称，项目组的显示名称 | [测试项目] |
| **CompanyID** | integer | 租户ID，资源所属的租户ID | [200000233] |
| **CompanyName** | string | 租户名称，租户的显示名称 | [ucloud] |
| **CompanyEmail** | string | 租户邮箱，租户的邮箱地址 | [ucloud@example.cn] |
| **ChargeType** | string | 计费类型，资源的计费方式 | [Hour] |
| **ProductType** | string | 产品类型，资源的产品类型 | [VM] |
| **StartTime** | integer | 开始时间，资源计费的开始时间戳 | [1753867542] |
| **ResourceName** | string | 资源名称，资源的显示名称 | [测试云主机] |

## 获取续费价格 - GetRenewPrice

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，查询续费价格的租户ID | [200000233] | [DescribeUser] | No |
| **Region** | string | 地域，您可以调用 DescribeRegion 方法查看最新的地域列表 | [pro2133] | [DescribeRegion] | **Yes** |
| **ResourceID** | string | 资源ID，要查询续费价格的资源唯一标识 | [vm-12saf3****sf456] | [] | **Yes** |
| **ChargeType** | string | 计费类型，续费的计费方式，取值范围：hour（按小时计费）、month（按月计费）、year（按年计费）。 | [Month] | [] | No |
| **Quantity** | integer | 数量，续费的数量 | [1] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Price** | double | 续费的价格，单位为元 | [150.50] |


## 获取账户可提现金额等信息 - GetWithdrawableAmount

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，查询指定租户的可提现金额信息 | [200000233] | [DescribeUser] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **Amount** | double | 租户账户中的现金余额 | [1000.50] |
| **AmountFree** | double | 租户账户中的赠金余额，单位为元 | [500.00] |
| **WithdrawableAmount** | double | 可提现余额，单位为元 | [800.50] |
| **WithdrawableAmountFree** | double | 可提现赠金余额，单位为元 | [300.00] |


## 获取提现流水列表 - DescribeWithdraw

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [DescribeUser] | No |
| **BeginTime** | integer | 开始时间 | [1751958000] | [] | **Yes** |
| **EndTime** | integer | 结束时间 | [1754550000] | [] | **Yes** |
| **Limit** | integer | 分页大小 | [10] | [] | No |
| **Offset** | integer | 分页偏移 | [0] | [] | No |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TotalCount** | integer | 提现记录总数量 | [25] |
| **Infos** | array<[*WithdrawTradeInfo*](#WithdrawTradeInfo)> | 详情 | [] |

### 数据模型

#### WithdrawTradeInfo

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID | [200000233] |
| **CompanyName** | string | 租户名称 | [testName] |
| **Email** | string | 租户邮箱 | [ucloud@example.cn] |
| **AccountID** | integer | 账号ID | [200000233] |
| **AccountName** | string | 账号名称 | [test_name] |
| **SerialNo** | string | 提现单号 | [serial-5exc2****vgn0d] |
| **TradeNo** | string | 交易单号 | [trade-pr1pv7knuiz7zj] |
| **WithdrawAmount** | double | 提现金额，单位为元 | [500.00] |
| **SrcAccountType** | string | 源账户类型(内部/外部) | [FREE] |
| **CreateTime** | integer | 创建时间 | [1753841388] |

## 申请提现 - Withdraw

### 请求参数

| 字段名 | 类型 | 描述信息 | 示例 | 参数来源 | 必填 |
|:---|:---|:---|:---|:---|:---|
| **CompanyID** | integer | 租户ID，用于标识资源所属的租户，实现多租户环境下的资源隔离 | [200000233] | [DescribeUser] | **Yes** |
| **SrcAccountType** | string | 源账户类型，内部/外部 | [FREE] | [] | **Yes** |
| **WithdrawAmount** | double | 提现金额，单位为元 | [500.00] | [] | **Yes** |

### 响应字段

| 字段名 | 类型 | 描述信息 | 示例 |
|:---|:---|:---|:---|
| **TransactonNo** | string | 交易号 | [trade-1tk013akzzsgtf] |

