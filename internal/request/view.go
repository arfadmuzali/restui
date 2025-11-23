package request

import (
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
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
				lipgloss.Top,
				zone.Mark("keyInputHeader",
					lipgloss.NewStyle().Width(m.RequestWidth*50/100-2).Border(lipgloss.RoundedBorder()).Render(m.KeyInput.View()),
				),
				zone.Mark("valueInputHeader",
					lipgloss.NewStyle().Width(m.RequestWidth*50/100-2).Border(lipgloss.RoundedBorder()).Render(m.ValueInput.View()),
				),
			),
		)
	}
	return zone.Mark("request", s)
}
