package app

import (
	tea "github.com/charmbracelet/bubbletea"
)

func Run() error {
	storage, err := storagePath()
	if err != nil {
		return err
	}

	todos, err := loadTodos(storage)
	if err != nil {
		return err
	}

	p := tea.NewProgram(newModel(storage, todos), tea.WithAltScreen())
	_, err = p.Run()
	return err
}

func newModel(storage string, todos []todo) model {
	h := helpModel()
	m := model{
		todos:          normalizeTodos(todos),
		mode:           modeList,
		statusFilter:   filterAll,
		categoryFilter: "all",
		help:           h,
		keys:           keys,
		styles:         newStyles(),
		storage:        storage,
		statusMessage:  "按 n 新增任务，Enter 编辑，1/2/3 切换状态筛选。",
		editingIndex:   -1,
		selected:       map[int]struct{}{},
	}
	m.initForm()
	m.initSearch()
	m.clampCursor()
	return m
}
