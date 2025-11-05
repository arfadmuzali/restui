package request

import zone "github.com/lrstanley/bubblezone"

func (m RequestModel) View() string {
	m.Viewport.SetContent(m.TextArea.View())
	return zone.Mark("request", m.Viewport.View())
}
