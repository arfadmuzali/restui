package app

import (
	"github.com/arfadmuzali/restui/internal/response"
	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmds []tea.Cmd
	var cmd tea.Cmd

	cmds = append(cmds, cmd)
	m.UrlModel, cmd = m.UrlModel.Update(msg)
	cmds = append(cmds, cmd)

	m.HintModel, cmd = m.HintModel.Update(msg)
	cmds = append(cmds, cmd)

	m.MethodModel, cmd = m.MethodModel.Update(msg)
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
	case tea.WindowSizeMsg:
		m.WindowWidth = msg.Width
		m.WindowHeight = msg.Height

	case response.IsLoadingMsg:
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
		case "ctrl+c":
			return m, tea.Quit
		case "ctrl+h":
			if m.MethodModel.OverlayActive {
				m.MethodModel.OverlayActive = false
			} else {
				m.MethodModel.OverlayActive = true
				m = m.BlurAllInput()
			}
		}
	}

	return m, tea.Batch(cmds...)
}
