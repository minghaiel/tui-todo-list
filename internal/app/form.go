package app

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m model) updateForm(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, m.keys.Cancel):
		m.mode = modeList
		m.editingIndex = -1
		m.errMessage = ""
		m.statusMessage = "已取消编辑。"
		m.blurForm()
		return m, nil
	case key.Matches(msg, m.keys.NextField):
		m.focusNext()
		return m, m.formInputs[m.formFocus].Focus()
	case key.Matches(msg, m.keys.PrevField):
		m.focusPrev()
		return m, m.formInputs[m.formFocus].Focus()
	case key.Matches(msg, m.keys.Save):
		return m.saveForm()
	case key.Matches(msg, m.keys.FormDelete):
		if m.editingIndex >= 0 && m.editingIndex < len(m.todos) {
			return m.deleteEditingTask()
		}
	case key.Matches(msg, m.keys.Cycle):
		if m.formFocus == int(fieldPriority) {
			m.shiftPriority(1)
			return m, nil
		}
	case key.Matches(msg, m.keys.Left):
		if m.formFocus == int(fieldPriority) {
			m.shiftPriority(-1)
			return m, nil
		}
	case key.Matches(msg, m.keys.Right):
		if m.formFocus == int(fieldPriority) {
			m.shiftPriority(1)
			return m, nil
		}
	case key.Matches(msg, m.keys.Up):
		if m.formFocus == int(fieldPriority) {
			m.shiftPriority(-1)
			return m, nil
		}
	case key.Matches(msg, m.keys.Down):
		if m.formFocus == int(fieldPriority) {
			m.shiftPriority(1)
			return m, nil
		}
	case msg.String() == "enter":
		return m.saveForm()
	}

	if m.formFocus == int(fieldPriority) {
		switch msg.String() {
		case "h", "k":
			m.shiftPriority(-1)
			return m, nil
		case "l", "j":
			m.shiftPriority(1)
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.formInputs[m.formFocus], cmd = m.formInputs[m.formFocus].Update(msg)
	return m, cmd
}

func (m *model) deleteEditingTask() (tea.Model, tea.Cmd) {
	if m.editingIndex < 0 || m.editingIndex >= len(m.todos) {
		return m, nil
	}

	title := m.todos[m.editingIndex].Title
	m.todos = append(m.todos[:m.editingIndex], m.todos[m.editingIndex+1:]...)
	m.mode = modeList
	m.editingIndex = -1
	m.errMessage = ""
	m.statusMessage = "已删除: " + title
	m.blurForm()
	m.clampCursor()
	m.ensureScroll()
	if err := saveTodos(m.storage, m.todos); err != nil {
		m.errMessage = err.Error()
	}
	return m, nil
}

func (m *model) saveForm() (tea.Model, tea.Cmd) {
	title := strings.TrimSpace(m.formInputs[fieldTitle].Value())
	if title == "" {
		m.errMessage = "标题不能为空。"
		return m, nil
	}

	priority, err := normalizePriority(m.formInputs[fieldPriority].Value())
	if err != nil {
		m.errMessage = err.Error()
		return m, nil
	}

	dueDate := strings.TrimSpace(m.formInputs[fieldDueDate].Value())
	if dueDate != "" {
		if _, err := time.Parse(dateLayout, dueDate); err != nil {
			m.errMessage = "截止日期格式必须是 YYYY-MM-DD。"
			return m, nil
		}
	}

	task := todo{
		Title:     title,
		Category:  normalizeCategory(m.formInputs[fieldCategory].Value()),
		Priority:  priority,
		DueDate:   dueDate,
		Completed: false,
	}

	if m.editingIndex >= 0 && m.editingIndex < len(m.todos) {
		task.Completed = m.todos[m.editingIndex].Completed
		m.todos[m.editingIndex] = task
		m.statusMessage = "任务已更新。"
	} else {
		m.todos = append(m.todos, task)
		m.statusMessage = "任务已创建。"
		m.cursor = len(m.filteredIndexes()) - 1
	}

	if err := saveTodos(m.storage, m.todos); err != nil {
		m.errMessage = err.Error()
		return m, nil
	}

	m.errMessage = ""
	m.mode = modeList
	m.editingIndex = -1
	m.blurForm()
	m.clampCursor()
	m.ensureScroll()
	return m, nil
}

func (m *model) initForm() {
	m.formInputs = []textinput.Model{
		newInput("比如：写周报"),
		newInput("比如：工作 / 学习 / 生活"),
		newInput("low / medium / high / urgent"),
		newInput("YYYY-MM-DD，可留空"),
	}
	m.formInputs[fieldTitle].CharLimit = 80
	m.formInputs[fieldCategory].CharLimit = 40
	m.formInputs[fieldPriority].CharLimit = 12
	m.formInputs[fieldDueDate].CharLimit = 10
	m.formInputs[fieldPriority].SetSuggestions([]string{"low", "medium", "high", "urgent"})
	m.formInputs[fieldPriority].ShowSuggestions = true
}

func newInput(placeholder string) textinput.Model {
	input := textinput.New()
	input.Prompt = ""
	input.Placeholder = placeholder
	input.Width = 36
	input.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#111827"))
	input.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#111827"))
	input.PlaceholderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#94A3B8"))
	return input
}

func (m model) renderTextField(input textinput.Model, focused bool) string {
	value := input.Value()
	if value == "" {
		value = input.Placeholder
		if focused {
			value += " |"
		}
		return m.styles.Muted.Render(value)
	}
	if focused {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("#111827")).Render(value + " |")
	}
	return lipgloss.NewStyle().Foreground(lipgloss.Color("#111827")).Render(value)
}

func (m *model) openForm(index int) {
	m.mode = modeForm
	m.editingIndex = index
	m.errMessage = ""
	m.initForm()

	if index >= 0 && index < len(m.todos) {
		item := m.todos[index]
		m.formInputs[fieldTitle].SetValue(item.Title)
		m.formInputs[fieldCategory].SetValue(item.Category)
		m.formInputs[fieldPriority].SetValue(item.Priority)
		m.formInputs[fieldDueDate].SetValue(item.DueDate)
	} else {
		m.formInputs[fieldPriority].SetValue("medium")
	}

	m.formFocus = 0
	for i := range m.formInputs {
		m.formInputs[i].Blur()
	}
	m.formInputs[m.formFocus].Focus()
}

func (m *model) focusNext() {
	m.formInputs[m.formFocus].Blur()
	m.formFocus = (m.formFocus + 1) % len(m.formInputs)
	m.formInputs[m.formFocus].Focus()
}

func (m *model) focusPrev() {
	m.formInputs[m.formFocus].Blur()
	m.formFocus--
	if m.formFocus < 0 {
		m.formFocus = len(m.formInputs) - 1
	}
	m.formInputs[m.formFocus].Focus()
}

func (m *model) blurForm() {
	for i := range m.formInputs {
		m.formInputs[i].Blur()
	}
}

func (m *model) shiftPriority(step int) {
	options := []string{"low", "medium", "high", "urgent"}
	current := normalizePriorityValue(m.formInputs[fieldPriority].Value())
	if current == "" {
		current = "medium"
	}

	index := indexOf(options, current)
	if index == -1 {
		index = 1
	}
	index = (index + step + len(options)) % len(options)
	m.formInputs[fieldPriority].SetValue(options[index])
	m.errMessage = ""
}
