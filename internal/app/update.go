package app

import (
	"github.com/arfadmuzali/restui/internal/config"
	"github.com/arfadmuzali/restui/internal/request"
	"github.com/arfadmuzali/restui/internal/response"
	"github.com/arfadmuzali/restui/internal/utils"
	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

func (m MainModel) Init() tea.Cmd {
	return nil
}

// global key msg does't affected by anything
func globalKeyMsg(m MainModel, msg tea.Msg) (MainModel, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.WindowWidth = msg.Width
		m.WindowHeight = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "ctrl+o":
			m.MethodModel.OverlayActive = !m.MethodModel.OverlayActive
			m = m.BlurAll()
			return m, nil
		case "f1":
			m.HelpModel.OverlayActive = !m.HelpModel.OverlayActive
			m = m.BlurAll()
			return m, nil
		}
	}
	return m, nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd
	var cmds []tea.Cmd

	m, cmd = globalKeyMsg(m, msg)
	cmds = append(cmds, cmd)

	m.MethodModel, cmd = m.MethodModel.Update(msg)
	cmds = append(cmds, cmd)

	m.HelpModel, cmd = m.HelpModel.Update(msg)
	cmds = append(cmds, cmd)

	if m.MethodModel.OverlayActive || m.HelpModel.OverlayActive {
		return m, tea.Batch(cmds...)
	}

	m.HintModel, cmd = m.HintModel.Update(msg)
	cmds = append(cmds, cmd)

	m.ResponseModel, cmd = m.ResponseModel.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case spinner.TickMsg:
		var cmd tea.Cmd
		if m.ResponseModel.IsLoading == true {
			m.spinner, cmd = m.spinner.Update(msg)
		}
		return m, cmd

	case response.IsLoadingMsg:
		// Add suggestion
		config.AddSuggestion(m.UrlModel.UrlInput.Value())
		suggestions, err := config.GetSuggestions()
		if err == nil {
			m.UrlModel.UrlInput.SetSuggestions(suggestions)
		}

		return m, m.HandleHttpRequest

	case tea.MouseMsg:
		if msg.Button == tea.MouseButtonLeft && msg.Action == tea.MouseActionRelease {
			if zone.Get("method").InBounds(msg) {
				m.MethodModel.OverlayActive = !m.MethodModel.OverlayActive
			} else if zone.Get("send").InBounds(msg) {
				return m.StartRequest()
			} else if zone.Get("copyResponseBody").InBounds(msg) {
				clipboard.WriteAll(string(m.ResponseModel.Result.Data))
				return m, nil
			}
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.UrlModel.UrlInput.Focused() {
				return m.StartRequest()
			}
		case "alt+enter":
			return m.StartRequest()
		// WARN: maybe this shortcut will cause bugs in the future
		case "ctrl+l":
			m = m.BlurAll()
			m.UrlModel.UrlInput.Focus()
			return m, nil
		case "ctrl+b":
			m = m.BlurAll()
			m.RequestModel.FocusedTab = request.Body
			m.RequestModel.Hovered = true
			m.RequestModel.TextArea.Focus()
			m.RequestModel.Viewport.SetContent(m.RequestModel.TextArea.View())
			m.RequestModel.Viewport.Height = m.RequestModel.RequestHeight
		case "ctrl+h":
			m = m.BlurAll()
			m.RequestModel.FocusedTab = request.Headers
			m.RequestModel.Hovered = true
			m.RequestModel.Viewport.Height = m.RequestModel.RequestHeight - utils.BoxStyle.GetVerticalBorderSize() - 1
			m.RequestModel.Viewport.SetContent(
				m.RequestModel.TableHeaders.View(),
			)
		}
	}

	m.RequestModel, cmd = m.RequestModel.Update(msg)
	cmds = append(cmds, cmd)

	m.UrlModel, cmd = m.UrlModel.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
