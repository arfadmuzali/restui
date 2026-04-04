package hint

import tea "charm.land/bubbletea/v2"

func (m HintModel) Init() tea.Cmd {
	return nil
}

func (m HintModel) Update(msg tea.Msg) (HintModel, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Help.SetWidth(msg.Width)
		return m, nil
	}

	return m, nil
}
