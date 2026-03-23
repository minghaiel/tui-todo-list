# tui-todo-list

`tui-todo-list` 是一个基于 Go 和 Bubble Tea 的终端待办事项应用。

它面向键盘操作，支持任务的创建、编辑、搜索、筛选、批量处理和本地持久化，适合在终端里快速管理个人待办。

## Screenshot

![tui-todo-list screenshot](./assets/image.png)

## Features

- 任务新增、编辑、删除
- 单任务完成状态切换
- 批量选择、批量完成、批量删除
- 状态筛选：`All / Open / Done`
- 分类筛选：按已有分类循环切换
- 搜索：按标题、分类、优先级、截止日期匹配
- 优先级：`low / medium / high / urgent`
- 二级排序：先按优先级，再按截止日期
- 截止日期与逾期识别
- 本地 JSON 持久化
- storage 和领域筛选/排序逻辑单元测试

## Stack

- Go `1.24.2`
- `github.com/charmbracelet/bubbletea`
- `github.com/charmbracelet/bubbles`
- `github.com/charmbracelet/lipgloss`

## Run

在项目目录执行：

```bash
go run .
```

构建：

```bash
go build ./...
```

运行测试：

```bash
go test ./...
```

## Data File

任务默认保存到：

```text
~/.todo-tui.json
```

首次启动如果没有本地数据，会生成示例任务。

## Interaction

### List View

- `n` / `a`：新建任务
- `enter` / `e`：编辑当前任务
- `↑/k`：上移
- `↓/j`：下移
- `c` / `space`：切换完成状态
- `d` / `x`：删除当前任务
- `/`：进入搜索
- `esc`：退出搜索
- `v`：选中或取消选中当前任务
- `u`：清空已选任务
- `C`：批量切换已选任务完成状态
- `X`：批量删除已选任务
- `1`：筛选 `All`
- `2`：筛选 `Open`
- `3`：筛选 `Done`
- `[` / `]`：切换分类筛选
- `?`：显示或收起帮助
- `q`：退出

### Form View

- `tab`：下一个字段
- `shift+tab`：上一个字段
- `enter`：确认当前字段，在最后一个字段保存
- `ctrl+s`：直接保存
- `esc`：取消编辑
- `ctrl+d`：删除当前正在编辑的任务

优先级字段支持：

- `←/h`
- `→/l`
- `↑/k`
- `↓/j`
- `p`

## Task Model

每个任务包含以下字段：

- `title`
- `category`
- `priority`
- `due_date`
- `completed`

分类为空时会归一化为 `inbox`，优先级会归一化为 `low / medium / high / urgent`。

## Sorting And Filtering

当前列表查询逻辑由领域层统一处理，规则如下：

- 先按状态筛选
- 再按分类筛选
- 再按搜索关键字筛选
- 排序优先级为：`urgent > high > medium > low`
- 同优先级下，截止日期更近的任务靠前
- 有截止日期的任务排在无截止日期任务前
- 未完成任务排在已完成任务前

## Project Structure

```text
.
├── main.go
├── README.md
├── go.mod
├── go.sum
├── tmp
│   └── image.png
└── internal
    ├── app
    │   ├── domain.go
    │   ├── form.go
    │   ├── keys.go
    │   ├── model.go
    │   ├── run.go
    │   ├── search.go
    │   ├── storage.go
    │   ├── storage_test.go
    │   ├── styles.go
    │   ├── types.go
    │   ├── update_list.go
    │   ├── util.go
    │   └── view.go
    └── domain
        ├── todo.go
        └── todo_test.go
```

## Architecture

当前代码按两层拆分：

- `internal/app`
  负责 Bubble Tea 的状态管理、交互更新、表单处理、视图渲染和持久化接入。
- `internal/domain`
  负责任务模型、字段归一化、分类选项、搜索、筛选和排序规则。

这种拆分的目标是：

- 让 UI 层和业务规则分开
- 让排序、筛选、归一化逻辑可单独测试
- 降低单文件复杂度，避免所有逻辑堆在 `main.go`

## Tests

当前已有两类单元测试：

- [internal/app/storage_test.go](/Users/starhming/owner/starhming/go-tui-demo/internal/app/storage_test.go#L1)
  覆盖本地保存和读取
- [internal/domain/todo_test.go](/Users/starhming/owner/starhming/go-tui-demo/internal/domain/todo_test.go#L1)
  覆盖筛选、搜索、分类选项和排序

## Limitations

- 当前只支持单个本地 JSON 文件存储，没有同步能力
- 批量操作以快捷键为主，没有单独的操作面板
- 排序策略目前固定，还不支持运行时自定义
- 还没有归档、标签、重复任务等更完整的任务管理能力

## Next

- 增加导入、导出和备份能力
- 增加标签和归档
- 扩展更多排序模式
- 增加更多领域和交互测试
