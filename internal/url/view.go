package url

import (
	zone "github.com/lrstanley/bubblezone"
)

func (m UrlModel) View() string {

	urlInputRendered := zone.Mark("url", m.UrlInput.View())
	return urlInputRendered
}
