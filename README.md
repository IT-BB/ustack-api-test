## UStack API 自动化测试项目

本项目旨在利用 **LLM（大语言模型）** 的推理能力，基于 API 文档自动生成高覆盖率、高质量的 **Go/Ginkgo** 测试代码，构建一套类似 **Kubernetes E2E** 的标准化自动化测试体系。

### 核心工作流

1.  **输入（Input）**：`docs/api/*.md`（API 定义） + `docs/userguide/*.md`（产品使用手册） + 领域知识（虚拟化/私有云常识）。
2.  **推理（Reasoning）**：由 LLM 扮演资深测试开发工程师，分析文档，设计覆盖 Happy Path、边界条件、异常分支、状态机流转的测试场景。
3.  **生成（Generation）**：LLM 直接产出基于 `Ginkgo` 框架的 Go 代码，代码即用例。
4.  **执行（Execution）**：使用 `ginkgo` 运行测试，输出 JUnit/Allure 报告，并在 CI/CD 中持续回归。

### 目录结构 (Ref: Kubernetes E2E)

```
ustack-api-test/
├── docs/api/                 # API 原始文档（测试源）
├── test/
│   ├── e2e/                  # E2E 测试代码根目录
│   │   ├── framework/        # 测试框架封装（ClientSet, Config, Waiter, Matcher）
│   │   ├── vm/               # VM 模块测试套件 (ginkgo specs)
│   │   ├── vpc/              # VPC 模块测试套件
│   │   ├── ...               # 其他模块
│   │   └── e2e_test.go       # 测试入口 (TestMain)
│   └── utils/                # 通用工具库
├── go.mod                    # Go 依赖定义
└── README.md                 # 项目说明
```

### 技术栈

-   **语言**：Go 1.25+
-   **测试框架**：[Ginkgo v2](https://github.com/onsi/ginkgo) (BDD 风格) + [Gomega](https://github.com/onsi/gomega) (断言库)
-   **HTTP 客户端**：原生 net/http 封装
-   **设计模式**：Service Client 模式（仿 OpenStack Tempest / K8s ClientSet）

### 快速开始

#### 1. 环境准备

确保已安装 Go 1.25+。设置必要的环境变量（用于连接被测环境）：

```bash
export USTACK_API_BASE_URL="http://192.168.170.180/api"
export USTACK_API_PUBLIC_KEY="gj-hCNWsxZ1v1Hvwk-MoNUmVts74m0ZsXMZujggSV5DTc7cq-kvW7-Jg"
export USTACK_API_PRIVATE_KEY="8WgjtUKBTNf4NVJX0a1SALfUH2luhojziN_LgmCR7VFasBCbP0Ugpg2tZxw75zA6"
export USTACK_API_REGION="region-zhu"
```

#### 2. 初始化依赖

```bash
go mod tidy
go install github.com/onsi/ginkgo/v2/ginkgo
```

#### 3. 运行测试

运行所有 E2E 测试：

```bash
ginkgo -r -v test/e2e
```

运行特定模块（通过 Label 过滤）：

```bash
# 只运行 VM 相关的测试
ginkgo -r -v --label-filter="Feature:VM" test/e2e

# 只运行冒烟测试
ginkgo -r -v --label-filter="Tier:Smoke" test/e2e
```

### 编写/生成测试指南

**不要手动从零编写测试！** 请遵循“LLM 辅助生成”流程：

1.  **提供上下文**：将 `docs/api/TargetModule.md` 和 `test/e2e/framework` 下的基础代码提供给 LLM。
2.  **下达指令**：
    > "作为测试专家，请阅读 VM 模块文档，基于我们现有的 framework 包，生成 Ginkgo 测试代码。要求覆盖创建、查询、生命周期管理（Start/Stop/Reboot）、异常参数校验等场景。请使用 `framework.NewClient` 获取客户端，使用 `gomega` 做断言。"
3.  **代码审查与入库**：将生成的代码保存到 `test/e2e/{module}/{module}_test.go`，运行测试验证无误后提交。

### 覆盖率目标

-   **API 覆盖率**：100% 覆盖文档中列出的所有 Action。
-   **参数覆盖率**：覆盖必填参数、常用可选参数组合、边界值。
-   **场景覆盖率**：覆盖资源完整生命周期（Create -> Update -> State Change -> Delete）。

---
*本项目采用文档驱动开发（Doc-Driven Testing），API 文档是测试用例的唯一真理来源。*
