package response

import tea "github.com/charmbracelet/bubbletea"

func (m ResponseModel) Init() tea.Cmd {
	return nil
}

func (m ResponseModel) Update(msg tea.Msg) (ResponseModel, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.ResponseHeight = msg.Height * 90 / 100
		m.ResponseWidth = msg.Width * 60 / 100
	}

	return m, nil
}
