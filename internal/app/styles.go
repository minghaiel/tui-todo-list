package app

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/lipgloss"
)

type styles struct {
	App            lipgloss.Style
	Header         lipgloss.Style
	HeaderEyebrow  lipgloss.Style
	HeaderTitle    lipgloss.Style
	HeaderSubtitle lipgloss.Style
	SectionTitle   lipgloss.Style
	Card           lipgloss.Style
	SelectedCard   lipgloss.Style
	Muted          lipgloss.Style
	Completed      lipgloss.Style
	Meta           lipgloss.Style
	Footer         lipgloss.Style
	Error          lipgloss.Style
	Empty          lipgloss.Style
	TitleInput     lipgloss.Style
	PanelTitle     lipgloss.Style
	FilterActive   lipgloss.Style
	FilterInactive lipgloss.Style
	FilterBar      lipgloss.Style
	FormLabel      lipgloss.Style
	FormHint       lipgloss.Style
	Overlay        lipgloss.Style
	Field          lipgloss.Style
	FocusedField   lipgloss.Style
}

func newStyles() styles {
	return styles{
		App:            lipgloss.NewStyle().Padding(1, 2),
		Header:         lipgloss.NewStyle().Padding(0, 2).MarginBottom(0).Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#0F5B6E")).Background(lipgloss.Color("#0E3A46")).Foreground(lipgloss.Color("#F7F3E9")),
		HeaderEyebrow:  lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#7DD3C7")),
		HeaderTitle:    lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFF7E6")),
		HeaderSubtitle: lipgloss.NewStyle().Foreground(lipgloss.Color("#C7DED7")),
		SectionTitle:   lipgloss.NewStyle().Bold(true).Foreground(lipgloss.AdaptiveColor{Light: "#334155", Dark: "#E5E7EB"}),
		Card:           lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.AdaptiveColor{Light: "#D6DEE8", Dark: "#6B7280"}).Background(lipgloss.AdaptiveColor{Light: "#FFFCF8", Dark: "#1F2937"}).Padding(0, 1).MarginBottom(1),
		SelectedCard:   lipgloss.NewStyle().Border(lipgloss.ThickBorder()).BorderForeground(lipgloss.Color("#D26838")).Background(lipgloss.AdaptiveColor{Light: "#FFF7ED", Dark: "#10212B"}).Padding(0, 1).MarginBottom(1),
		Muted:          lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#6B7280", Dark: "#CBD5E1"}),
		Completed:      lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#6B8A7A", Dark: "#A7F3D0"}).Strikethrough(true),
		Meta:           lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#475569", Dark: "#E2E8F0"}),
		Footer:         lipgloss.NewStyle().PaddingTop(1).Foreground(lipgloss.Color("#64748B")),
		Error:          lipgloss.NewStyle().Foreground(lipgloss.Color("#B42318")).Bold(true),
		Empty:          lipgloss.NewStyle().Border(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("#CBD5E1")).Padding(1, 2).Foreground(lipgloss.Color("#64748B")),
		TitleInput:     lipgloss.NewStyle().Bold(true).Foreground(lipgloss.AdaptiveColor{Light: "#14213D", Dark: "#F8FAFC"}),
		PanelTitle:     lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#0F172A")),
		FilterActive:   lipgloss.NewStyle().Bold(true).Background(lipgloss.Color("#D26838")).Foreground(lipgloss.Color("#FFF7E6")).Padding(0, 1).MarginRight(1),
		FilterInactive: lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#64748B", Dark: "#D1D5DB"}).Background(lipgloss.AdaptiveColor{Light: "#FFFFFF", Dark: "#1F2937"}).Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.AdaptiveColor{Light: "#CBD5E1", Dark: "#4B5563"}).Padding(0, 1).MarginRight(1),
		FilterBar:      lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.AdaptiveColor{Light: "#D6DEE8", Dark: "#4B5563"}).Background(lipgloss.AdaptiveColor{Light: "#F8FAFC", Dark: "#111827"}).Padding(0, 1).MarginBottom(0),
		FormLabel:      lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#334155")),
		FormHint:       lipgloss.NewStyle().Foreground(lipgloss.Color("#64748B")),
		Overlay:        lipgloss.NewStyle().Border(lipgloss.DoubleBorder()).BorderForeground(lipgloss.Color("#113946")).Padding(1, 2).Background(lipgloss.Color("#F8FAFC")),
		Field:          lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#CBD5E1")).Padding(0, 1).Background(lipgloss.Color("#FFFFFF")).Foreground(lipgloss.Color("#111827")),
		FocusedField:   lipgloss.NewStyle().Border(lipgloss.ThickBorder()).BorderForeground(lipgloss.Color("#D26838")).Padding(0, 1).Background(lipgloss.Color("#FFFDF8")).Foreground(lipgloss.Color("#111827")),
	}
}

func helpModel() help.Model {
	h := help.New()
	h.ShowAll = false
	h.Width = 80
	return h
}
