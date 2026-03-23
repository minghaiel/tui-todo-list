package app

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Up         key.Binding
	Down       key.Binding
	Left       key.Binding
	Right      key.Binding
	Cycle      key.Binding
	New        key.Binding
	Edit       key.Binding
	Toggle     key.Binding
	Delete     key.Binding
	FormDelete key.Binding
	Filter     key.Binding
	Category   key.Binding
	Help       key.Binding
	Quit       key.Binding
	NextField  key.Binding
	PrevField  key.Binding
	Save       key.Binding
	Cancel     key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.New, k.Toggle, k.Delete, k.Filter, k.Category, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Left, k.Right},
		{k.Toggle, k.Delete, k.New, k.Edit},
		{k.Filter, k.Category, k.NextField, k.PrevField},
		{k.Save, k.Cancel, k.Help, k.Quit},
	}
}

var keys = keyMap{
	Up:         key.NewBinding(key.WithKeys("up", "k"), key.WithHelp("↑/k", "up")),
	Down:       key.NewBinding(key.WithKeys("down", "j"), key.WithHelp("↓/j", "down")),
	Left:       key.NewBinding(key.WithKeys("left", "h"), key.WithHelp("←/h", "prev option")),
	Right:      key.NewBinding(key.WithKeys("right", "l"), key.WithHelp("→/l", "next option")),
	New:        key.NewBinding(key.WithKeys("n", "a"), key.WithHelp("n", "new")),
	Edit:       key.NewBinding(key.WithKeys("enter", "e"), key.WithHelp("enter", "edit")),
	Toggle:     key.NewBinding(key.WithKeys("c", " "), key.WithHelp("c", "done")),
	Delete:     key.NewBinding(key.WithKeys("d", "x"), key.WithHelp("d/x", "delete")),
	FormDelete: key.NewBinding(key.WithKeys("ctrl+d"), key.WithHelp("ctrl+d", "delete task")),
	Filter:     key.NewBinding(key.WithKeys("f", "1", "2", "3"), key.WithHelp("1/2/3", "status")),
	Category:   key.NewBinding(key.WithKeys("g", "[", "]"), key.WithHelp("[/]", "category")),
	Help:       key.NewBinding(key.WithKeys("?"), key.WithHelp("?", "help")),
	Quit:       key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q", "quit")),
	NextField:  key.NewBinding(key.WithKeys("tab"), key.WithHelp("tab", "next field")),
	PrevField:  key.NewBinding(key.WithKeys("shift+tab"), key.WithHelp("shift+tab", "prev field")),
	Save:       key.NewBinding(key.WithKeys("ctrl+s"), key.WithHelp("ctrl+s", "save")),
	Cancel:     key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "cancel")),
	Cycle:      key.NewBinding(key.WithKeys("p"), key.WithHelp("p", "cycle priority")),
}
