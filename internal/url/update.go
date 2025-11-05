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
		m.UrlInput.Width = msg.Width * 75 / 100
	case tea.MouseMsg:
		if msg.Action == tea.MouseActionRelease && msg.Button == tea.MouseButtonLeft {
			if zone.Get("url").InBounds(msg) {
				m.UrlInput.Focus()
				if m.UrlInput.Focused() {
					cursorPosition := msg.X - zone.Get("url").StartX
					// TODO: FIX this bug: the bug is when the cursor > urlInputWIdth the click cursor doesnt work correctly
					if m.UrlInput.Width < m.UrlInput.Position() {
						m.UrlInput.SetCursor(m.UrlInput.Position() - m.UrlInput.Width + cursorPosition)
					} else {
						m.UrlInput.SetCursor(cursorPosition)
					}
				}
			} else {
				m.UrlInput.Blur()
			}
		}

	}
	var cmd tea.Cmd
	m.UrlInput, cmd = m.UrlInput.Update(msg)
	return m, cmd
}
