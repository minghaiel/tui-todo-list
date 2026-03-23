package app

import (
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/textinput"
)

const dateLayout = "2006-01-02"

type appMode int

const (
	modeList appMode = iota
	modeForm
)

type statusFilter int

const (
	filterAll statusFilter = iota
	filterOpen
	filterDone
)

type inputField int

const (
	fieldTitle inputField = iota
	fieldCategory
	fieldPriority
	fieldDueDate
)

type todo struct {
	Title     string `json:"title"`
	Category  string `json:"category,omitempty"`
	Priority  string `json:"priority,omitempty"`
	DueDate   string `json:"due_date,omitempty"`
	Completed bool   `json:"completed"`
}

type model struct {
	todos          []todo
	cursor         int
	scrollOffset   int
	width          int
	height         int
	mode           appMode
	statusFilter   statusFilter
	categoryFilter string
	help           help.Model
	keys           keyMap
	styles         styles
	storage        string
	statusMessage  string
	errMessage     string
	showHelp       bool
	formInputs     []textinput.Model
	formFocus      int
	editingIndex   int
}

func normalizeTodos(todos []todo) []todo {
	if len(todos) == 0 {
		return []todo{
			{Title: "升级这个 TUI", Category: "work", Priority: "high", DueDate: time.Now().AddDate(0, 0, 2).Format(dateLayout)},
			{Title: "买点水果", Category: "life", Priority: "medium", DueDate: time.Now().AddDate(0, 0, 1).Format(dateLayout)},
		}
	}

	for i := range todos {
		todos[i].Category = normalizeCategory(todos[i].Category)
		todos[i].Priority = normalizePriorityValue(todos[i].Priority)
		todos[i].DueDate = trimSpace(todos[i].DueDate)
	}
	return todos
}
