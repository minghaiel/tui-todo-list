package app

import (
	"errors"
	"sort"
	"strings"
	"time"
)

func (m model) filteredIndexes() []int {
	var indexes []int
	for i, item := range m.todos {
		if !m.matchesStatus(item) {
			continue
		}
		if m.categoryFilter != "all" && normalizeCategory(item.Category) != m.categoryFilter {
			continue
		}
		indexes = append(indexes, i)
	}
	return indexes
}

func (m model) matchesStatus(item todo) bool {
	switch m.statusFilter {
	case filterOpen:
		return !item.Completed
	case filterDone:
		return item.Completed
	default:
		return true
	}
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
	set := map[string]struct{}{"all": {}}
	for _, item := range m.todos {
		set[normalizeCategory(item.Category)] = struct{}{}
	}

	var options []string
	for category := range set {
		options = append(options, category)
	}
	sort.Strings(options)
	if idx := indexOf(options, "all"); idx > 0 {
		options[0], options[idx] = options[idx], options[0]
	}
	return options
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
	value = strings.TrimSpace(strings.ToLower(value))
	if value == "" {
		return "inbox"
	}
	return value
}

func normalizePriority(input string) (string, error) {
	priority := normalizePriorityValue(input)
	if priority == "" {
		return "", errors.New("优先级必须是 low、medium、high 或 urgent。")
	}
	return priority, nil
}

func normalizePriorityValue(input string) string {
	value := strings.TrimSpace(strings.ToLower(input))
	switch value {
	case "", "med":
		return "medium"
	case "low", "medium", "high", "urgent":
		return value
	default:
		return ""
	}
}

func isOverdue(item todo, now time.Time) bool {
	if item.Completed || trimSpace(item.DueDate) == "" {
		return false
	}
	date, err := time.Parse(dateLayout, item.DueDate)
	if err != nil {
		return false
	}
	ny, nm, nd := now.Date()
	today := time.Date(ny, nm, nd, 0, 0, 0, 0, now.Location())
	return date.Before(today)
}

func sameDay(a, b time.Time) bool {
	ay, am, ad := a.Date()
	by, bm, bd := b.Date()
	return ay == by && am == bm && ad == bd
}
