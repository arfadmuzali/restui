package request

import (
	"github.com/arfadmuzali/restui/internal/utils"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

func (m RequestModel) Init() tea.Cmd {
	return textarea.Blink
}

func (m RequestModel) Update(msg tea.Msg) (RequestModel, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd
	m.TextArea, cmd = m.TextArea.Update(msg)
	cmds = append(cmds, cmd)
	m.Viewport, cmd = m.Viewport.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// minus 1 for the tabs
		m.RequestHeight = msg.Height*90/100 - utils.BoxStyle.GetVerticalBorderSize() - 1

		m.RequestWidth = msg.Width*40/100 - utils.BoxStyle.GetHorizontalBorderSize()

		m.TextArea.SetWidth(m.RequestWidth)
		m.TextArea.SetHeight(m.RequestHeight)
		m.TextArea.MaxWidth = m.RequestWidth

		if !m.ViewportReady {
			m.Viewport = viewport.New(m.RequestWidth, m.RequestHeight)
			m.ViewportReady = true
			m.Viewport.SetContent(m.TextArea.View())
		} else {
			m.Viewport.Width = m.RequestWidth
			m.Viewport.Height = m.RequestHeight
		}
	case tea.MouseMsg:
		m.Hovered = zone.Get("request").InBounds(msg)

		if m.Hovered && m.FocusedTab == Body {
			m.TextArea.Focus()
		} else {
			m.TextArea.Blur()
		}

		if msg.Action == tea.MouseActionRelease && msg.Button == tea.MouseButtonLeft {

			if zone.Get("requestBody").InBounds(msg) {
				m.FocusedTab = Body
				m.TextArea.Focus()
			} else if zone.Get("requestHeaders").InBounds(msg) {
				m.FocusedTab = Headers

			}
		}
	}
	return m, tea.Batch(cmds...)
}
