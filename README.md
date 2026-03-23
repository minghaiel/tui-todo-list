# tui-todo-list

一个用 Go 编写的终端待办事项应用，基于 `Bubble Tea`、`Bubbles` 和 `Lip Gloss` 实现。

## Screenshot

将运行截图放到 `assets/tui-todo-list.png` 后，这里会直接显示：

![tui-todo-list screenshot](assets/tui-todo-list.png)

支持：

- 任务新增、编辑、删除
- 完成状态切换
- 分类筛选
- 状态筛选：`All / Open / Done`
- 优先级：`low / medium / high / urgent`
- 截止日期
- 本地 JSON 持久化

## Tech Stack

- Go `1.24.2`
- `github.com/charmbracelet/bubbletea`
- `github.com/charmbracelet/bubbles`
- `github.com/charmbracelet/lipgloss`

## Run

在项目目录执行：

```bash
go run .
```

编译：

```bash
go build ./...
```

## Data Storage

任务数据默认保存在：

```text
~/.todo-tui.json
```

## Features

每个任务包含这些字段：

- `title`
- `category`
- `priority`
- `due_date`
- `completed`

应用内置：

- 列表视图
- 表单视图
- 彩色状态徽标
- 深浅色终端适配的基础样式

## Keybindings

### List View

- `n` / `a`: 新建任务
- `enter` / `e`: 编辑当前任务
- `↑/k`: 上移
- `↓/j`: 下移
- `c` / `space`: 切换完成状态
- `d` / `x`: 删除当前任务
- `1`: 筛选 `All`
- `2`: 筛选 `Open`
- `3`: 筛选 `Done`
- `[` / `]`: 切换分类筛选
- `?`: 展开帮助
- `q`: 退出

### Form View

- `tab`: 下一个字段
- `shift+tab`: 上一个字段
- `enter`: 保存
- `ctrl+s`: 保存
- `esc`: 取消
- `ctrl+d`: 删除当前正在编辑的任务

优先级字段支持：

- `←/h`
- `→/l`
- `↑/k`
- `↓/j`
- `p`

## Project Structure

```text
.
├── main.go
├── go.mod
├── go.sum
└── internal
    └── app
        ├── domain.go
        ├── form.go
        ├── keys.go
        ├── model.go
        ├── run.go
        ├── storage.go
        ├── styles.go
        ├── types.go
        ├── update_list.go
        ├── util.go
        └── view.go
```

## Design Notes

这次重构的目标是避免把所有逻辑堆在一个文件里，并尽量符合常见设计原则：

- 单一职责：视图、更新、存储、领域规则分离
- 高内聚：表单逻辑集中在 `form.go`
- 低耦合：入口只依赖 `app.Run()`
- 可维护性：按 Bubble Tea 的职责拆分 `Update` / `View` / state

当前职责划分：

- `run.go`: 应用启动和装配
- `types.go`: 基础类型和应用状态
- `keys.go`: 按键映射
- `styles.go`: UI 样式
- `model.go`: Bubble Tea 生命周期入口
- `update_list.go`: 列表页交互
- `form.go`: 表单页交互与保存
- `view.go`: 所有视图渲染
- `domain.go`: 过滤、分类、优先级、日期等规则
- `storage.go`: 本地持久化
- `util.go`: 通用辅助方法

## Limitations

当前版本还有一些可以继续优化的点：

- 表单 `Enter` 直接保存，交互还可以更细化
- 缺少自动化测试
- 长标题的布局处理还可以继续加强
- 任务排序和搜索还没有实现

## Next Steps

如果继续演进，建议优先做：

1. 为 `storage` 和筛选逻辑补单元测试
2. 给表单改成“逐字段确认”的交互
3. 抽离更明确的领域层，例如 `internal/domain`
4. 增加搜索、排序、批量操作
