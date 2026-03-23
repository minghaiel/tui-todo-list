package app

import (
	"time"
	"tui-todo-list/internal/domain"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/textinput"
)

const dateLayout = domain.DateLayout

type appMode int

const (
	modeList appMode = iota
	modeForm
)

type statusFilter int

const (
	filterAll  statusFilter = statusFilter(domain.FilterAll)
	filterOpen statusFilter = statusFilter(domain.FilterOpen)
	filterDone statusFilter = statusFilter(domain.FilterDone)
)

type inputField int

const (
	fieldTitle inputField = iota
	fieldCategory
	fieldPriority
	fieldDueDate
)

type todo = domain.Todo

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
	searchInput    textinput.Model
	searchMode     bool
	searchQuery    string
	selected       map[int]struct{}
}

func normalizeTodos(todos []todo) []todo {
	return domain.NormalizeTodos(todos, time.Now())
}
