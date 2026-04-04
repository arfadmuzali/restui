package request

import (
	"charm.land/lipgloss/v2"
	zone "github.com/lrstanley/bubblezone/v2"
)

func (m RequestModel) View() string {
	var s string
	switch m.FocusedTab {
	case Body:
		m.Viewport.SetContent(m.TextArea.View())
		s = m.Viewport.View()
	case Headers:
		m.Viewport.SetContent(m.TableHeaders.View())
		s = lipgloss.JoinVertical(
			lipgloss.Left,
			m.Viewport.View(),
			lipgloss.JoinHorizontal(
				lipgloss.Center,
				zone.Mark("keyInputHeader",
					//BUG: i dont know why i have to -3
					lipgloss.NewStyle().Width(m.KeyInput.Width()+3).Border(lipgloss.RoundedBorder()).Render(m.KeyInput.View()),
				),
				zone.Mark("valueInputHeader",
					//BUG: i dont know why i have to -3
					lipgloss.NewStyle().Width(m.ValueInput.Width()+3).Border(lipgloss.RoundedBorder()).Render(m.ValueInput.View()),
				),
			),
		)
	}
	return zone.Mark("request", s)
}
