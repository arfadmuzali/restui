package url

import (
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

func (m UrlModel) View() string {
	urlInputRendered := zone.Mark("url", lipgloss.NewStyle().MaxWidth(m.UrlInput.Width).Render(m.UrlInput.View()))
	return urlInputRendered
}
