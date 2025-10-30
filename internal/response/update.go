package response

import (
	"github.com/arfadmuzali/restui/internal/utils"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

func (m ResponseModel) Init() tea.Cmd {
	return nil
}

func (m ResponseModel) Update(msg tea.Msg) (ResponseModel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// minus 1 for text header
		m.ResponseHeight = msg.Height*90/100 - utils.BoxStyle.GetVerticalBorderSize() - 1

		var addon int
		if msg.Width%10 != 0 {
			addon = 1
		}
		m.ResponseWidth = msg.Width*60/100 - utils.BoxStyle.GetHorizontalBorderSize() + addon

		if !m.ViewportReady {
			m.Viewport = viewport.New(m.ResponseWidth, m.ResponseHeight)
			m.ViewportReady = true
		} else {
			m.Viewport.Width = m.ResponseWidth
			m.Viewport.Height = m.ResponseHeight
		}
	case tea.MouseMsg:
		m.Hovered = zone.Get("response").InBounds(msg)
	}
	if m.Hovered {
		m.Viewport, cmd = m.Viewport.Update(msg)
	}

	return m, cmd
}
