package url

import (
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

func (m UrlModel) View() string {
	// m.UrlInput.Width = m.windowWidth
	if m.UrlInput.Focused() {
		m.UrlInput.PromptStyle.BorderStyle(lipgloss.RoundedBorder())
	}
	urlInputRendered := zone.Mark("url", m.UrlInput.View())
	return urlInputRendered
}
