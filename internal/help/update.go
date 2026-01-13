package help

import (
	"github.com/arfadmuzali/restui/internal/utils"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
)

func (m HelpModel) Init() tea.Cmd {
	return nil
}

func (m HelpModel) Update(msg tea.Msg) (HelpModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.helpWindowWidth = msg.Width * 70 / 100
		m.helpWindowHeight = msg.Height * 80 / 100
		if !m.ViewportReady {
			m.Viewport = viewport.New(
				m.helpWindowWidth-utils.BoxStyle.GetHorizontalBorderSize(),
				m.helpWindowHeight-utils.BoxStyle.GetVerticalBorderSize(),
			)

			m.ViewportReady = true

			r, glamourErr := glamour.NewTermRenderer(
				glamour.WithStylePath("dark"),
				glamour.WithWordWrap(m.helpWindowWidth-utils.BoxStyle.GetHorizontalBorderSize()),
			)

			if glamourErr != nil {
				m.Viewport.SetContent("Failed to render guide " + glamourErr.Error())
				return m, nil
			}

			guide, err := r.Render(Guide)

			if err != nil {
				m.Viewport.SetContent("Failed to render guide " + err.Error())
			} else {
				m.Viewport.SetContent(string(guide))
			}

		} else {
			m.Viewport.Width = m.helpWindowWidth - utils.BoxStyle.GetHorizontalBorderSize()
			m.Viewport.Height = m.helpWindowHeight - utils.BoxStyle.GetVerticalBorderSize()
		}
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.OverlayActive = false
		}

	}
	var cmd tea.Cmd
	if m.OverlayActive && m.ViewportReady {
		m.Viewport, cmd = m.Viewport.Update(msg)
	}

	return m, cmd
}
