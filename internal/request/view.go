package request

import zone "github.com/lrstanley/bubblezone"

func (m RequestModel) View() string {
	switch m.FocusedTab {
	case Body:
		m.Viewport.SetContent(m.TextArea.View())
	case Headers:
		m.Viewport.SetContent(m.CreateHeadersTable().Render())

	}
	return zone.Mark("request", m.Viewport.View())
}
