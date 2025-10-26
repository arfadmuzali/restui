package app

import (
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// case tea.MouseMsg:
	// 	if msg.Action != tea.MouseActionRelease || msg.Button != tea.MouseButtonLeft {
	// 		if zone.Get("method").InBounds(msg) {
	// 			m.OverlayActive = true
	// 		}
	// 	}
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.WindowWidth = msg.Width
		m.WindowHeight = msg.Height

	case tea.MouseMsg:
		if msg.Action != tea.MouseActionRelease || msg.Button != tea.MouseButtonLeft {

			if zone.Get("method").InBounds(msg) {
				if m.MethodModel.OverlayActive == true {
					m.MethodModel.OverlayActive = false
				} else {
					m.MethodModel.OverlayActive = true
				}
			}
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "ctrl+h":
			if m.MethodModel.OverlayActive {
				m.MethodModel.OverlayActive = false
			} else {
				m.MethodModel.OverlayActive = true
				m = m.BlurAllInput()
			}
		}
	}

	var cmds []tea.Cmd
	var cmd tea.Cmd

	m.UrlModel, cmd = m.UrlModel.Update(msg)
	cmds = append(cmds, cmd)

	m.HintModel, cmd = m.HintModel.Update(msg)
	cmds = append(cmds, cmd)

	m.MethodModel, cmd = m.MethodModel.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
