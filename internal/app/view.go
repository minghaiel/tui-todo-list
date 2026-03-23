package app

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)

func (m model) renderHeader() string {
	innerWidth := max(20, m.listBlockWidth()-4)
	headerFill := lipgloss.NewStyle().
		Width(innerWidth).
		Background(lipgloss.Color("#0E3A46"))

	title := headerFill.Render(m.styles.HeaderTitle.Render("Tui-Todo-List"))
	subtitle := headerFill.Render(m.styles.HeaderSubtitle.Render("Priority-first task tracking"))
	stats := headerFill.Render(m.renderStats())

	return m.styles.Header.Width(m.listBlockWidth()).Render(lipJoinVertical(title, subtitle, stats))
}

func (m model) renderStats() string {
	total, done, overdue := len(m.todos), 0, 0
	today := time.Now()
	for _, item := range m.todos {
		if item.Completed {
			done++
		}
		if isOverdue(item, today) {
			overdue++
		}
	}

	open := total - done
	return lipgloss.JoinHorizontal(lipgloss.Left,
		m.badge("All", fmt.Sprintf("%d", total), "#164E63", "#E6FFFB"),
		m.badge("Open", fmt.Sprintf("%d", open), "#164E63", "#FFD8A8"),
		m.badge("Done", fmt.Sprintf("%d", done), "#164E63", "#BBF7D0"),
		m.badge("Overdue", fmt.Sprintf("%d", overdue), "#164E63", "#FECACA"),
	)
}

func (m model) badge(label string, value string, bg string, fg string) string {
	return lipgloss.NewStyle().
		Bold(true).
		Padding(0, 1).
		MarginRight(1).
		Background(lipgloss.Color(bg)).
		Foreground(lipgloss.Color(fg)).
		Render(label + " " + value)
}

func (m model) renderList() string {
	filters := m.renderFilters()
	filtered := m.filteredIndexes()
	if len(filtered) == 0 {
		empty := m.styles.Empty.Width(m.listBlockWidth()).Render("当前筛选条件下没有任务。\n按 n 新建一个任务，或按 1/2/3、[/] 调整筛选。")
		return lipJoinVertical(filters, empty)
	}

	var rows []string
	availableHeight := max(5, m.height-16)
	start := m.scrollOffset
	end := min(len(filtered), start+availableHeight)
	for i := start; i < end; i++ {
		rows = append(rows, m.renderTaskCard(filtered[i], m.todos[filtered[i]], i == m.cursor))
	}
	return lipJoinVertical(filters, lipJoinVertical(rows...))
}

func (m model) renderTaskCard(index int, item todo, selected bool) string {
	cardStyle := m.styles.Card
	if selected {
		cardStyle = m.styles.SelectedCard
	}

	check := "○"
	if item.Completed {
		check = "●"
	}

	titleStyle := m.styles.TitleInput
	if item.Completed {
		titleStyle = m.styles.Completed
	}

	titleText := truncateRunes(item.Title, max(18, m.listBlockWidth()-18))
	pick := " "
	if _, ok := m.selected[index]; ok {
		pick = "•"
	}
	title := titleStyle.Render(pick + " " + check + " " + titleText)
	meta := lipgloss.JoinHorizontal(lipgloss.Left,
		m.priorityBadge(item.Priority),
		m.categoryBadge(item.Category),
		m.dueBadge(item),
	)
	body := lipJoinVertical(title, m.styles.Meta.Render(meta))
	return cardStyle.Width(m.listBlockWidth()).Render(body)
}

func (m model) renderFilters() string {
	statuses := []struct {
		label string
		mode  statusFilter
		key   string
	}{
		{"all", filterAll, "1"},
		{"open", filterOpen, "2"},
		{"done", filterDone, "3"},
	}

	var statusParts []string
	for _, entry := range statuses {
		style := m.styles.FilterInactive
		if m.statusFilter == entry.mode {
			style = m.styles.FilterActive
		}
		statusParts = append(statusParts, style.Render(entry.key+":"+entry.label))
	}

	categoryStyle := m.styles.FilterInactive
	if m.categoryFilter != "all" {
		categoryStyle = m.styles.FilterActive
	}
	category := categoryStyle.Render("cat:" + m.categoryFilter + " [←/→]")
	searchLabel := m.styles.FilterInactive.Render("search: " + searchPreview(m))
	if m.searchMode {
		searchLabel = m.styles.FilterActive.Render("search: " + m.searchInput.Value() + "|")
	}
	if !m.searchMode && m.searchQuery != "" {
		searchLabel = m.styles.FilterActive.Render("search: " + m.searchQuery)
	}
	filterLine := lipgloss.JoinHorizontal(
		lipgloss.Left,
		append(statusParts, "  ", searchLabel, "  ", category)...,
	)
	return m.styles.FilterBar.Width(m.listBlockWidth()).Render(filterLine)
}

func (m model) renderForm() string {
	title := "New Task"
	if m.editingIndex >= 0 {
		title = "Edit Task"
	}

	labels := []string{"标题", "分类", "优先级", "截止日期"}
	hints := []string{
		"必填",
		"留空时默认 inbox",
		"可选值：low / medium / high / urgent，可用方向键或 p 切换",
		"格式：YYYY-MM-DD",
	}

	var lines []string
	lines = append(lines, m.styles.PanelTitle.Render(title))
	lines = append(lines, m.styles.FormHint.Render("Tab 切换字段；优先级支持方向键和 p；Enter 或 Ctrl+S 保存；编辑已有任务时 Ctrl+D 删除。"))
	lines = append(lines, "")

	for i, input := range m.formInputs {
		label := m.styles.FormLabel.Render(labels[i])
		fieldStyle := m.styles.Field
		if i == m.formFocus {
			fieldStyle = m.styles.FocusedField
		}

		fieldView := m.renderTextField(input, i == m.formFocus)
		if i == int(fieldPriority) {
			fieldView = m.renderPriorityPicker()
		}

		lines = append(lines, lipJoinVertical(label, fieldStyle.Width(44).Render(fieldView)))
		lines = append(lines, m.styles.FormHint.Render(hints[i]))
	}

	if m.errMessage != "" {
		lines = append(lines, "", m.styles.Error.Render(m.errMessage))
	}

	panelWidth := min(max(56, m.width-8), 84)
	return m.styles.Overlay.Width(panelWidth).Render(lipJoinVertical(lines...))
}

func (m model) renderFooter() string {
	message := m.statusMessage
	if m.errMessage != "" && m.mode == modeList {
		message = m.styles.Error.Render(m.errMessage)
	}
	shortHelp := m.help.ShortHelpView(m.keys.ShortHelp())
	statusLine := fmt.Sprintf("Filter: %s / %s  •  %s", m.statusFilterLabel(), m.categoryFilter, message)
	if len(m.selected) > 0 {
		statusLine = fmt.Sprintf("%s  •  selected:%d", statusLine, len(m.selected))
	}
	return m.styles.Footer.Render(lipJoinVertical(
		statusLine,
		shortHelp,
	))
}

func (m model) priorityBadge(priority string) string {
	p := normalizePriorityValue(priority)
	var bg, fg string
	switch p {
	case "low":
		bg, fg = "#E7F7ED", "#166534"
	case "high":
		bg, fg = "#FFF0DA", "#9A3412"
	case "urgent":
		bg, fg = "#FDEBEC", "#B42318"
	default:
		bg, fg = "#E6EEF8", "#1D4ED8"
		p = "medium"
	}
	return lipgloss.NewStyle().Bold(true).Padding(0, 1).MarginRight(1).Background(lipgloss.Color(bg)).Foreground(lipgloss.Color(fg)).Render(strings.ToUpper(p))
}

func (m model) renderPriorityPicker() string {
	options := []string{"low", "medium", "high", "urgent"}
	current := normalizePriorityValue(m.formInputs[fieldPriority].Value())
	if current == "" {
		current = "medium"
	}

	parts := make([]string, 0, len(options))
	for _, option := range options {
		if option == current {
			parts = append(parts, lipgloss.NewStyle().
				Bold(true).
				Padding(0, 1).
				MarginRight(1).
				Background(lipgloss.Color("#113946")).
				Foreground(lipgloss.Color("#FFF7E6")).
				Render("["+strings.ToUpper(option)+"]"))
			continue
		}
		parts = append(parts, lipgloss.NewStyle().
			Padding(0, 1).
			MarginRight(1).
			Foreground(lipgloss.Color("#64748B")).
			Render(strings.ToUpper(option)))
	}
	return lipgloss.JoinHorizontal(lipgloss.Left, parts...)
}

func (m model) categoryBadge(category string) string {
	c := normalizeCategory(category)
	bg, fg := categoryColors(c)
	return lipgloss.NewStyle().Padding(0, 1).MarginRight(1).Background(lipgloss.Color(bg)).Foreground(lipgloss.Color(fg)).Render("#" + c)
}

func (m model) dueBadge(item todo) string {
	if trimSpace(item.DueDate) == "" {
		return lipgloss.NewStyle().Padding(0, 1).Foreground(lipgloss.Color("#64748B")).Render("No due date")
	}

	date, err := time.Parse(dateLayout, item.DueDate)
	if err != nil {
		return lipgloss.NewStyle().Padding(0, 1).Foreground(lipgloss.Color("#B42318")).Render("Invalid due date")
	}

	now := time.Now()
	label := "Due " + item.DueDate
	style := lipgloss.NewStyle().Padding(0, 1).Background(lipgloss.Color("#EEF2FF")).Foreground(lipgloss.Color("#4338CA"))
	if !item.Completed && sameDay(date, now) {
		style = lipgloss.NewStyle().Padding(0, 1).Background(lipgloss.Color("#FFF0DA")).Foreground(lipgloss.Color("#9A3412"))
		label = "Due today"
	}
	if isOverdue(item, now) {
		style = lipgloss.NewStyle().Padding(0, 1).Background(lipgloss.Color("#FDEBEC")).Foreground(lipgloss.Color("#B42318"))
		label = "Overdue " + item.DueDate
	}
	return style.Render(label)
}
