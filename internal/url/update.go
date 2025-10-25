package url

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

func (m UrlModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m UrlModel) Update(msg tea.Msg) (UrlModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		urlSize := msg.Width * 75 / 100

		m.UrlInput.Width = urlSize
	case tea.MouseMsg:
		if msg.Action != tea.MouseActionRelease || msg.Button != tea.MouseButtonLeft {
			if zone.Get("url").InBounds(msg) {
				m.UrlInput.Focus()
			} else {
				m.UrlInput.Blur()
			}
		}
	case tea.KeyMsg:
		switch msg.String() {
		// WARN: maybe this shortcut will be bug in the future
		case "ctrl+l":
			m.UrlInput.Focus()
		}

	}
	var cmd tea.Cmd
	m.UrlInput, cmd = m.UrlInput.Update(msg)
	return m, cmd
}
