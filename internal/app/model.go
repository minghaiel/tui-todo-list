package app

import tea "github.com/charmbracelet/bubbletea"

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.help.Width = max(30, msg.Width-8)
		m.ensureScroll()
		return m, nil
	case tea.KeyMsg:
		if m.mode == modeForm {
			return m.updateForm(msg)
		}
		return m.updateList(msg)
	}

	if m.mode == modeForm {
		var cmd tea.Cmd
		m.formInputs[m.formFocus], cmd = m.formInputs[m.formFocus].Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m model) View() string {
	header := m.renderHeader()
	body := m.renderList()
	if m.mode == modeForm {
		body = lipJoinVertical(body, "", m.renderForm())
	}
	footer := m.renderFooter()
	return m.styles.App.Render(lipJoinVertical(header, "", body, "", footer))
}
