package app

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Up          key.Binding
	Down        key.Binding
	Left        key.Binding
	Right       key.Binding
	Cycle       key.Binding
	New         key.Binding
	Edit        key.Binding
	Search      key.Binding
	SearchExit  key.Binding
	Toggle      key.Binding
	Delete      key.Binding
	FormDelete  key.Binding
	Select      key.Binding
	ClearBatch  key.Binding
	BatchDone   key.Binding
	BatchDelete key.Binding
	Filter      key.Binding
	Category    key.Binding
	Help        key.Binding
	Quit        key.Binding
	NextField   key.Binding
	PrevField   key.Binding
	Save        key.Binding
	Cancel      key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.New, k.Search, k.Toggle, k.Select, k.BatchDelete, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Left, k.Right},
		{k.Toggle, k.Delete, k.Select, k.ClearBatch},
		{k.BatchDone, k.BatchDelete, k.New, k.Edit},
		{k.Search, k.SearchExit, k.Filter, k.Category},
		{k.Save, k.Cancel, k.Help, k.Quit},
	}
}

var keys = keyMap{
	Up:          key.NewBinding(key.WithKeys("up", "k"), key.WithHelp("↑/k", "up")),
	Down:        key.NewBinding(key.WithKeys("down", "j"), key.WithHelp("↓/j", "down")),
	Left:        key.NewBinding(key.WithKeys("left", "h"), key.WithHelp("←/h", "prev option")),
	Right:       key.NewBinding(key.WithKeys("right", "l"), key.WithHelp("→/l", "next option")),
	New:         key.NewBinding(key.WithKeys("n", "a"), key.WithHelp("n", "new")),
	Edit:        key.NewBinding(key.WithKeys("enter", "e"), key.WithHelp("enter", "edit")),
	Search:      key.NewBinding(key.WithKeys("/"), key.WithHelp("/", "search")),
	SearchExit:  key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "close search")),
	Toggle:      key.NewBinding(key.WithKeys("c", " "), key.WithHelp("c", "done")),
	Delete:      key.NewBinding(key.WithKeys("d", "x"), key.WithHelp("d/x", "delete")),
	FormDelete:  key.NewBinding(key.WithKeys("ctrl+d"), key.WithHelp("ctrl+d", "delete task")),
	Select:      key.NewBinding(key.WithKeys("v"), key.WithHelp("v", "select")),
	ClearBatch:  key.NewBinding(key.WithKeys("u"), key.WithHelp("u", "clear picks")),
	BatchDone:   key.NewBinding(key.WithKeys("C"), key.WithHelp("C", "toggle picks")),
	BatchDelete: key.NewBinding(key.WithKeys("X"), key.WithHelp("X", "delete picks")),
	Filter:      key.NewBinding(key.WithKeys("f", "1", "2", "3"), key.WithHelp("1/2/3", "status")),
	Category:    key.NewBinding(key.WithKeys("g", "[", "]"), key.WithHelp("[/]", "category")),
	Help:        key.NewBinding(key.WithKeys("?"), key.WithHelp("?", "help")),
	Quit:        key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q", "quit")),
	NextField:   key.NewBinding(key.WithKeys("tab"), key.WithHelp("tab", "next field")),
	PrevField:   key.NewBinding(key.WithKeys("shift+tab"), key.WithHelp("shift+tab", "prev field")),
	Save:        key.NewBinding(key.WithKeys("ctrl+s"), key.WithHelp("ctrl+s", "save")),
	Cancel:      key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "cancel")),
	Cycle:       key.NewBinding(key.WithKeys("p"), key.WithHelp("p", "cycle priority")),
}
