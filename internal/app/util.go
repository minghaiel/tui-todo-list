package app

import (
	"strings"
	"unicode/utf8"

	"github.com/charmbracelet/lipgloss"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func indexOf(values []string, target string) int {
	for i, value := range values {
		if value == target {
			return i
		}
	}
	return -1
}

func trimSpace(value string) string {
	return strings.TrimSpace(value)
}

func lipJoinVertical(lines ...string) string {
	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}

func truncateRunes(value string, limit int) string {
	if limit <= 0 {
		return ""
	}
	if utf8.RuneCountInString(value) <= limit {
		return value
	}
	runes := []rune(value)
	if limit <= 1 {
		return string(runes[:limit])
	}
	return string(runes[:limit-1]) + "…"
}

func (m model) listBlockWidth() int {
	return min(max(48, m.width-10), 76)
}
