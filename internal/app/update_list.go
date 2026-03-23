package app

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) updateList(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	filtered := m.filteredIndexes()
	keyStr := msg.String()

	switch {
	case key.Matches(msg, m.keys.Quit):
		return m, tea.Quit
	case key.Matches(msg, m.keys.Help):
		m.showHelp = !m.showHelp
		m.help.ShowAll = m.showHelp
	case key.Matches(msg, m.keys.Up):
		if m.cursor > 0 {
			m.cursor--
		}
	case key.Matches(msg, m.keys.Down):
		if m.cursor < len(filtered)-1 {
			m.cursor++
		}
	case key.Matches(msg, m.keys.New):
		m.openForm(-1)
		return m, m.formInputs[0].Focus()
	case key.Matches(msg, m.keys.Edit):
		if len(filtered) == 0 {
			break
		}
		m.openForm(filtered[m.cursor])
		return m, m.formInputs[0].Focus()
	case key.Matches(msg, m.keys.Toggle):
		if len(filtered) == 0 {
			break
		}
		m.toggleTodo(filtered[m.cursor])
	case key.Matches(msg, m.keys.Delete):
		if len(filtered) == 0 {
			break
		}
		m.deleteTodo(filtered[m.cursor])
	case key.Matches(msg, m.keys.Filter):
		m.applyStatusFilter(keyStr)
	case key.Matches(msg, m.keys.Category):
		if keyStr == "[" {
			m.categoryFilter = m.prevCategory()
		} else {
			m.categoryFilter = m.nextCategory()
		}
		m.cursor = 0
		m.scrollOffset = 0
		m.statusMessage = "分类筛选已切换。"
	}

	m.clampCursor()
	m.ensureScroll()
	return m, nil
}

func (m *model) toggleTodo(index int) {
	m.todos[index].Completed = !m.todos[index].Completed
	m.statusMessage = "任务状态已更新。"
	m.errMessage = ""
	if err := saveTodos(m.storage, m.todos); err != nil {
		m.errMessage = err.Error()
	}
}

func (m *model) deleteTodo(index int) {
	title := m.todos[index].Title
	m.todos = append(m.todos[:index], m.todos[index+1:]...)
	m.statusMessage = "已删除: " + title
	m.errMessage = ""
	if err := saveTodos(m.storage, m.todos); err != nil {
		m.errMessage = err.Error()
	}
}

func (m *model) applyStatusFilter(keyStr string) {
	switch keyStr {
	case "1":
		m.statusFilter = filterAll
	case "2":
		m.statusFilter = filterOpen
	case "3":
		m.statusFilter = filterDone
	default:
		m.statusFilter = (m.statusFilter + 1) % 3
	}
	m.cursor = 0
	m.scrollOffset = 0
	m.statusMessage = "状态筛选已切换。"
}
