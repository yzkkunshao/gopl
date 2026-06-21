# 修改 CLAUDE.md：增加标准库解释规则

## TL;DR

> **目标**：在 CLAUDE.md 中增加指令，让 agent 在讲解 Go 代码时对标准库给出充分的分层解释
> 
> **交付物**：
> - 修改后的 CLAUDE.md（包含学习者画像、标准库解释规则、修改后的交互规则、标准库速查表）
> 
> **预计工作量**：Quick
> **并行执行**：NO - 单文件顺序修改

---

## Context

### 原始需求
用户对 Go 标准库不熟悉，希望 CLAUDE.md 中增加规则让 agent 在讲解代码时给出充分的标准库解释，而非一笔带过。

### 修改方案

共 4 处修改，全部在 `CLAUDE.md` 一个文件中：

#### 修改1：在 `## Role: Go 编程导师` 之后、`## 交互规则` 之前插入新段落

插入位置：第 29 行（`## Role` 段落）之后，第 31 行（`## 交互规则`）之前

插入内容：

```markdown
## 学习者画像

- **Go 经验水平**：初学者，对 Go 标准库不熟悉
- **学习需求**：不仅要知道代码怎么写，还要理解为什么这样写、标准库提供了什么、还有哪些替代方案

## 标准库解释规则（MANDATORY）

每当代码中出现 Go 标准库的包、函数、类型或方法时，**必须**提供以下三层解释：

### 第一层：这是什么？
- 用一句话说明该标准库元素的用途和定位
- 示例：`fmt.Fprintf(w io.Writer, format string, args ...interface{})` → "将格式化字符串写入任意 io.Writer，而不仅是标准输出"

### 第二层：为什么用它？
- 说明选择这个函数/类型而非其他替代方案的原因
- 说明它解决了什么问题
- 示例：为什么用 `bufio.Scanner` 而非 `bufio.ReadLine`？→ "Scanner 提供更简洁的逐行读取接口，自动处理行尾，而 ReadLine 需要手动管理缓冲区"

### 第三层：还有哪些替代？
- 列出同包或相关包中的替代方案
- 简要说明何时用替代方案
- 示例：`fmt.Fprintf` 的替代 → `fmt.Sprintf`（只需字符串不需写入）/ `fmt.Printf`（固定写标准输出）

### 解释强度分级

| 标准库出现频率 | 解释要求 |
|---|---|
| 首次出现 | 完整三层解释（是什么 + 为什么 + 替代） |
| 同一包内不同函数 | 两层解释（是什么 + 与同包其他函数对比） |
| 重复出现（同包同函数） | 简要提示 → "（详见 chX 的 XXX 解释）" |
| 常见/易混淆函数 | 强调区别 → "注意：`strings.Trim` 移除字符集，`strings.TrimSuffix` 移除指定后缀" |
```

#### 修改2：扩展"章节讲解"步骤3

**原文**（第 37 行）：
```
3. 每个示例要说明：**这段代码要解决什么问题、关键语法是什么、运行会输出什么**
```

**替换为**：
```
3. 每个示例要说明：
   - **这段代码要解决什么问题**
   - **关键语法是什么**
   - **使用了哪些标准库元素**（按"标准库解释规则"分层解释，首次出现做完整三层）
   - **运行会输出什么**
```

#### 修改3：修改"提问回答"规则

**原文**（第 51-53 行）：
```
- 结合本章代码示例回答
- 指出对应的代码文件和行号
- 必要时可以对比前后章节的差异
```

**替换为**：
```
- 结合本章代码示例回答
- 指出对应的代码文件和行号
- 必要时可以对比前后章节的差异
- **遇到标准库函数/类型时，按"标准库解释规则"分层解释，不得省略**
```

#### 修改4：在文档末尾（常用命令之后）追加标准库速查表

追加内容：

```markdown

## 标准库包速查（本书涉及）

> 以下列出本书各章涉及的核心标准库包，供快速定位。详细解释见各章节讲解。

| 包 | 首次出现 | 一句话定位 |
|---|---|---|
| `fmt` | ch1 | 格式化 I/O：Print/Scan/Fprintf 等函数族 |
| `os` | ch1 | 操作系统交互：文件、环境变量、命令行参数 |
| `bufio` | ch1 | 缓冲 I/O：Scanner、Writer、Reader |
| `io` | ch7 | I/O 原语：Reader、Writer、Copy 等接口 |
| `strings` | ch3 | 字符串操作：查找、替换、分割、Trim 族 |
| `strconv` | ch3 | 字符串与基本类型互转：Atoi、Itoa、FormatFloat |
| `unicode` | ch3 | Unicode 支持：IsDigit、IsLetter、IsSpace |
| `net/http` | ch1/ch7 | HTTP 客户端与服务端 |
| `encoding/json` | ch4 | JSON 编解码：Marshal、Unmarshal |
| `html/template` | ch4 | HTML 模板引擎 |
| `sync` | ch9 | 同步原语：Mutex、RWMutex、Once、WaitGroup |
| `reflect` | ch12 | 运行时反射：Type、Value |
| `unsafe` | ch13 | 底层编程：Sizeof、Alignof、Offsetof |
| `time` | ch8 | 时间操作：定时器、Ticker、Sleep |
| `context` | ch8 | 请求生命周期管理与取消信号传播 |
| `math` | ch3 | 数学函数：Pi、Sqrt、Abs、Cos 等 |
| `image` | ch1 | 图像处理：RGBA、Pix、解码编码 |
| `path/filepath` | ch8 | 跨平台路径操作 |
| `sort` | ch7 | 排序接口与工具函数 |
| `errors` | ch7 | 错误创建与处理 |
| `testing` | ch11 | 测试框架：Test、Benchmark、Example |
```

---

## TODOs

- [x] 1. 在 `## Role: Go 编程导师` 之后插入"学习者画像"和"标准库解释规则"段落

  **What to do**:
  - 在 CLAUDE.md 的第 29 行（`## Role` 段落结束）之后、第 31 行（`## 交互规则`）之前，插入"学习者画像"和"标准库解释规则（MANDATORY）"两个新段落
  - 内容如上方"修改1"所述

  **Must NOT do**:
  - 不改动已有的 Role 描述内容

  **Recommended Agent Profile**:
  - **Category**: `quick`
  - **Skills**: []

  **Parallelization**:
  - **Can Run In Parallel**: NO
  - **Parallel Group**: Sequential
  - **Blocks**: Task 2, 3, 4
  - **Blocked By**: None

  **References**:
  - `CLAUDE.md:27-30` - 当前 Role 段落，在其后插入
  - `CLAUDE.md:31` - `## 交互规则` 标题，插入点在此行之前

  **Acceptance Criteria**:
  - [ ] CLAUDE.md 中 `## Role` 和 `## 交互规则` 之间出现"学习者画像"和"标准库解释规则"两个段落
  - [ ] 标准库解释规则包含三层解释（是什么/为什么/替代）和解释强度分级表

  **QA Scenarios**:
  ```
  Scenario: 验证学习者画像段落存在
    Tool: Bash (grep)
    Steps:
      1. grep -c "学习者画像" CLAUDE.md
      2. Assert count >= 1
    Expected Result: 找到学习者画像段落
    Evidence: .sisyphus/evidence/task-1-learner-profile.txt
  
  Scenario: 验证标准库解释规则存在
    Tool: Bash (grep)
    Steps:
      1. grep -c "标准库解释规则" CLAUDE.md
      2. Assert count >= 1
    Expected Result: 找到标准库解释规则段落
    Evidence: .sisyphus/evidence/task-1-stdlib-rule.txt
  ```

  **Commit**: YES (group with Task 2-4)
  - Message: `docs: add stdlib explanation rules to CLAUDE.md`
  - Files: `CLAUDE.md`

- [x] 2. 扩展"章节讲解"步骤3

  **What to do**:
  - 将第 37 行的单行步骤3替换为四点列表格式
  - 新增"使用了哪些标准库元素"条目，引用"标准库解释规则"

  **Must NOT do**:
  - 不修改步骤1、2、4的内容

  **Recommended Agent Profile**:
  - **Category**: `quick`
  - **Skills**: []

  **Parallelization**:
  - **Can Run In Parallel**: NO
  - **Parallel Group**: Sequential (after Task 1)
  - **Blocks**: None
  - **Blocked By**: Task 1

  **References**:
  - `CLAUDE.md:37` - 当前步骤3内容：`3. 每个示例要说明：**这段代码要解决什么问题、关键语法是什么、运行会输出什么**`

  **Acceptance Criteria**:
  - [ ] 步骤3变为四点列表
  - [ ] 包含"使用了哪些标准库元素"条目

  **QA Scenarios**:
  ```
  Scenario: 验证步骤3格式
    Tool: Bash (grep)
    Steps:
      1. grep "使用了哪些标准库元素" CLAUDE.md
    Expected Result: 找到新条目
    Evidence: .sisyphus/evidence/task-2-step3.txt
  ```

  **Commit**: YES (group with Task 1, 3, 4)

- [x] 3. 修改"提问回答"规则

  **What to do**:
  - 在提问回答规则第 53 行之后追加一条规则：遇到标准库函数/类型时按规则分层解释

  **Must NOT do**:
  - 不修改前三条规则内容

  **Recommended Agent Profile**:
  - **Category**: `quick`
  - **Skills**: []

  **Parallelization**:
  - **Can Run In Parallel**: NO
  - **Parallel Group**: Sequential (after Task 2)
  - **Blocks**: None
  - **Blocked By**: Task 2

  **References**:
  - `CLAUDE.md:51-53` - 当前提问回答规则三条

  **Acceptance Criteria**:
  - [ ] 提问回答规则有四条，第四条包含"标准库解释规则"引用

  **QA Scenarios**:
  ```
  Scenario: 验证第四条规则
    Tool: Bash (grep)
    Steps:
      1. grep "标准库解释规则" CLAUDE.md | wc -l
    Expected Result: 至少2行匹配（规则定义处 + 引用处）
    Evidence: .sisyphus/evidence/task-3-qa-rule.txt
  ```

  **Commit**: YES (group with Task 1, 2, 4)

- [x] 4. 在文档末尾追加标准库包速查表

  **What to do**:
  - 在 `## 常用命令` 代码块之后，追加 `## 标准库包速查（本书涉及）` 段落
  - 包含表格，列出本书涉及的核心标准库包及其首次出现章节和一句话定位

  **Must NOT do**:
  - 不修改已有内容
  - 不遗漏任何在项目中实际使用的包

  **Recommended Agent Profile**:
  - **Category**: `quick`
  - **Skills**: []

  **Parallelization**:
  - **Can Run In Parallel**: NO
  - **Parallel Group**: Sequential (after Task 3)
  - **Blocks**: None
  - **Blocked By**: Task 3

  **References**:
  - `CLAUDE.md:73-87` - 当前常用命令段落（文档末尾）

  **Acceptance Criteria**:
  - [ ] 文档末尾出现"标准库包速查"段落
  - [ ] 表格包含至少 15 个包
  - [ ] 每个包有首次出现章节和一句话定位

  **QA Scenarios**:
  ```
  Scenario: 验证速查表存在且完整
    Tool: Bash (grep)
    Steps:
      1. grep "标准库包速查" CLAUDE.md
      2. grep -c "| \`" CLAUDE.md（统计表格行数）
    Expected Result: 速查标题存在，表格行数 >= 15
    Evidence: .sisyphus/evidence/task-4-ref-table.txt
  ```

  **Commit**: YES (group with Task 1-3)
  - Message: `docs: add stdlib explanation rules to CLAUDE.md`
  - Files: `CLAUDE.md`

---

## Final Verification Wave

- [x] F1. **Plan Compliance Audit** — `oracle`
  读取 CLAUDE.md，验证：学习者画像段落存在、标准库解释规则含三层+分级表、步骤3含标准库条目、提问回答含第四条、速查表含 15+ 包。

- [x] F2. **Content Quality Review** — `unspecified-high`
  逐条检查新增内容：三层解释是否清晰、分级表是否合理、速查表是否准确匹配各章内容。

- [x] F3. **Format Integrity** — `quick`
  确认 CLAUDE.md Markdown 格式正确、无错位标题、无断行、表格渲染正常。

- [x] F4. **Scope Fidelity Check** — `deep`
  验证只修改了 CLAUDE.md 一个文件，且仅新增了 4 处内容，未改动原有内容（除了步骤3的格式变更）。

---

## Commit Strategy

- **1**: `docs: add stdlib explanation rules to CLAUDE.md` - CLAUDE.md

---

## Success Criteria

### Verification Commands
```bash
grep -c "学习者画像" CLAUDE.md       # Expected: >= 1
grep -c "标准库解释规则" CLAUDE.md    # Expected: >= 2 (定义处 + 引用处)
grep "使用了哪些标准库元素" CLAUDE.md  # Expected: found
grep "标准库包速查" CLAUDE.md         # Expected: found
grep -c "| \`" CLAUDE.md              # Expected: >= 15 (速查表行)
```

### Final Checklist
- [x] 学习者画像段落存在
- [x] 标准库解释规则含三层解释 + 强度分级表
- [x] 章节讲解步骤3包含标准库元素条目
- [x] 提问回答规则含第四条（分层解释）
- [x] 速查表包含 15+ 个标准库包
- [x] 所有原有内容未被破坏