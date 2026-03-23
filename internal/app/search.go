package app

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m *model) initSearch() {
	input := textinput.New()
	input.Prompt = ""
	input.Placeholder = "Search title, category, priority..."
	input.Width = 28
	input.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#111827"))
	input.PlaceholderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#94A3B8"))
	m.searchInput = input
}

func (m *model) openSearch() tea.Cmd {
	m.searchMode = true
	m.searchInput.SetValue(m.searchQuery)
	return m.searchInput.Focus()
}

func (m *model) closeSearch(clear bool) {
	m.searchMode = false
	m.searchInput.Blur()
	if clear {
		m.searchInput.SetValue("")
		m.searchQuery = ""
		m.statusMessage = "搜索已清除。"
	}
}
