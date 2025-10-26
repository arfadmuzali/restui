package method

import (
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

func (m MethodModel) Init() tea.Cmd {
	return nil
}

func (m MethodModel) Update(msg tea.Msg) (MethodModel, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.MouseMsg:
		if zone.Get("GET").InBounds(msg) {
			m.ActiveState = GET
			m.OverlayActive = false
		} else if zone.Get("POST").InBounds(msg) {
			m.ActiveState = POST
			m.OverlayActive = false
		} else if zone.Get("PUT").InBounds(msg) {
			m.ActiveState = PUT
			m.OverlayActive = false
		} else if zone.Get("PATCH").InBounds(msg) {
			m.ActiveState = PATCH
			m.OverlayActive = false
		} else if zone.Get("DELETE").InBounds(msg) {
			m.ActiveState = DELETE
			m.OverlayActive = false
		}
	case tea.WindowSizeMsg:
		m.windowHeight = msg.Height
		m.windowWidth = msg.Width
	case tea.KeyMsg:

		if m.OverlayActive == true {
			switch msg.String() {
			case "up", "k":
				if m.ActiveState > 0 {
					m.ActiveState = m.ActiveState - 1
				}
			case "down", "j":
				if m.ActiveState < 4 {
					m.ActiveState = m.ActiveState + 1
				}
			case "esc":
				m.OverlayActive = false
			case "enter":
				m.OverlayActive = false
			case "g", "G":
				m.ActiveState = GET
				m.OverlayActive = false
			case "p", "P":
				m.ActiveState = POST
				m.OverlayActive = false
			case "u", "U":
				m.ActiveState = PUT
				m.OverlayActive = false
			case "a", "A":
				m.ActiveState = PATCH
				m.OverlayActive = false
			case "d", "D":
				m.ActiveState = DELETE
				m.OverlayActive = false
			}
		}

	}

	return m, nil
}
