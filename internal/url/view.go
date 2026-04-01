package url

import (
	"charm.land/lipgloss/v2"
	zone "github.com/lrstanley/bubblezone/v2"
)

func (m UrlModel) View() string {
	urlInputRendered := zone.Mark("url", lipgloss.NewStyle().MaxWidth(m.UrlInput.Width()).Render(m.UrlInput.View()))
	return urlInputRendered
}
