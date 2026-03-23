package app

import (
	"strings"

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
