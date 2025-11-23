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

			rawGuide := `# Common Keybinds
- **ctrl+c**: Exit
- **ESC**: Close modal/popup
- **alt+enter**: Send request
- **F1**: Help
# Request Keybinds
- **ctrl+l**: Go to url
- **ctrl+b**: Go to request body
- **ctrl+o**: Toggle method modal/popup
    - **k/up**: Navigate up in method modal/popup
    - **j/down**: Navigate down  in method modal/popup
    - **g**: Switch to GET method
    - **p**: Switch to POST method
    - **u**: Switch to PUT method
    - **a**: Switch to PATCH method
    - **d**: Switch to DELETE method
    - **ESC/enter**: Close the method modal/popup
- **ctrl+h**: Go to request header
    - **tab**: Toggle between header key and header value
        - **enter**: Add header
    - **ctrl+k/up**: Navigate header table up
    - **ctrl+j/down**: Navigate header table down
    - **ctrl+d**: Delete header
# Tips & Tricks
- if you want to select and copy text, use copy-mode/selection-mode in your terminal level. For an example Kitty Terminal Emulator has ctrl+shift+h to enter copy-mode`

			r, glamourErr := glamour.NewTermRenderer(
				glamour.WithWordWrap(m.helpWindowWidth - utils.BoxStyle.GetHorizontalBorderSize()),
			)

			if glamourErr != nil {
				m.Viewport.SetContent("Failed to render guide " + glamourErr.Error())
				return m, nil
			}

			guide, err := r.Render(rawGuide)

			if err != nil {
				m.Viewport.SetContent("Failed to render guide " + err.Error())
			} else {
				m.Viewport.SetContent(string(guide))
			}

		} else {
			m.Viewport.Width = m.helpWindowWidth - utils.BoxStyle.GetHorizontalBorderSize()
			m.Viewport.Height = m.helpWindowHeight - utils.BoldStyle.GetVerticalBorderSize()
		}
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.OverlayActive = false
		}

	}
	var cmd tea.Cmd
	m.Viewport, cmd = m.Viewport.Update(msg)

	return m, cmd
}
