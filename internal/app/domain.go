package app

import (
	"time"
	"tui-todo-list/internal/domain"
)

func (m model) filteredIndexes() []int {
	return domain.FilterAndSortIndexes(m.todos, domain.Query{
		Status:   domain.StatusFilter(m.statusFilter),
		Category: m.categoryFilter,
		Search:   m.searchQuery,
		Sort:     domain.SortPriorityDue,
	}, time.Now())
}

func (m model) matchesStatus(item todo) bool {
	indexes := domain.FilterAndSortIndexes([]todo{item}, domain.Query{
		Status: domain.StatusFilter(m.statusFilter),
		Sort:   domain.SortPriorityDue,
	}, time.Now())
	return len(indexes) == 1
}

func (m *model) clampCursor() {
	filtered := m.filteredIndexes()
	if len(filtered) == 0 {
		m.cursor = 0
		return
	}
	if m.cursor < 0 {
		m.cursor = 0
	}
	if m.cursor >= len(filtered) {
		m.cursor = len(filtered) - 1
	}
}

func (m *model) ensureScroll() {
	visibleHeight := max(5, m.height-16)
	if m.cursor < m.scrollOffset {
		m.scrollOffset = m.cursor
	}
	if m.cursor >= m.scrollOffset+visibleHeight {
		m.scrollOffset = m.cursor - visibleHeight + 1
	}
	if m.scrollOffset < 0 {
		m.scrollOffset = 0
	}
}

func (m model) nextCategory() string {
	categories := m.categoryOptions()
	pos := indexOf(categories, m.categoryFilter)
	if pos == -1 {
		return "all"
	}
	return categories[(pos+1)%len(categories)]
}

func (m model) prevCategory() string {
	categories := m.categoryOptions()
	pos := indexOf(categories, m.categoryFilter)
	if pos == -1 {
		return "all"
	}
	pos--
	if pos < 0 {
		pos = len(categories) - 1
	}
	return categories[pos]
}

func (m model) categoryOptions() []string {
	return domain.CategoryOptions(m.todos)
}

func (m model) statusFilterLabel() string {
	switch m.statusFilter {
	case filterOpen:
		return "open"
	case filterDone:
		return "done"
	default:
		return "all"
	}
}

func categoryColors(category string) (string, string) {
	palette := map[string][2]string{
		"inbox":  {"#F1F5F9", "#334155"},
		"work":   {"#E0F2FE", "#0369A1"},
		"study":  {"#FEF3C7", "#92400E"},
		"life":   {"#DCFCE7", "#166534"},
		"health": {"#FCE7F3", "#9D174D"},
		"home":   {"#F3E8FF", "#7E22CE"},
	}
	if colors, ok := palette[category]; ok {
		return colors[0], colors[1]
	}
	return "#E2E8F0", "#334155"
}

func normalizeCategory(value string) string {
	return domain.NormalizeCategory(value)
}

func normalizePriority(input string) (string, error) {
	return domain.NormalizePriority(input)
}

func normalizePriorityValue(input string) string {
	return domain.NormalizePriorityValue(input)
}

func priorityOrder(priority string) int {
	return domain.PriorityOrder(priority)
}

func isOverdue(item todo, now time.Time) bool {
	return domain.IsOverdue(item, now)
}

func sameDay(a, b time.Time) bool {
	return domain.SameDay(a, b)
}
