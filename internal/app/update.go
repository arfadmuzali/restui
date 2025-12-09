package app

import (
	"github.com/arfadmuzali/restui/internal/config"
	methodModel "github.com/arfadmuzali/restui/internal/method"
	"github.com/arfadmuzali/restui/internal/request"
	"github.com/arfadmuzali/restui/internal/response"
	"github.com/arfadmuzali/restui/internal/utils"
	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
	zone "github.com/lrstanley/bubblezone"
)

func (m MainModel) Init() tea.Cmd {

	return tea.Batch(
		m.RequestModel.Init(),
		m.ResponseModel.Init(),
		m.HelpModel.Init(),
		m.MethodModel.Init(),
		m.UrlModel.Init(),
		m.HintModel.Init(),
	)
}

// global key msg does't affected by anything
func globalKeyMsg(m MainModel, msg tea.Msg) (MainModel, tea.Cmd) {
	var cmds []tea.Cmd

	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.WindowWidth = msg.Width
		m.WindowHeight = msg.Height

		if !m.BufferModalModel.ViewportReady {
			m.BufferModalModel.ViewportReady = true
			m.BufferModalModel.Viewport = viewport.New(
				m.WindowWidth*50/100-utils.BoxStyle.GetHorizontalBorderSize(),
				m.WindowHeight*85/100-utils.BoxStyle.GetHorizontalBorderSize(),
			)
		} else {
			m.BufferModalModel.Viewport.Width = m.WindowWidth*50/100 - utils.BoxStyle.GetHorizontalBorderSize()
			m.BufferModalModel.Viewport.Height = m.WindowHeight*85/100 - utils.BoxStyle.GetVerticalBorderSize()
		}
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+n":
			buffer := CreateNewBuffer()
			if m.BufferModalModel.OverlayActive {
				return m, tea.WindowSize()
			}

			m.IndexBuffers[buffer.Id] = len(m.Buffers)
			m.Buffers = append(m.Buffers, buffer)

			m = m.ChangeBuffer(buffer.Id)

			return m, tea.WindowSize()
		case "ctrl+pgup":
			index := m.IndexBuffers[m.ActiveBufferId]
			if index < len(m.Buffers)-1 {
				m = m.ChangeBuffer(m.Buffers[index+1].Id)
			}

			if index == len(m.Buffers)-1 {
				m = m.ChangeBuffer(m.Buffers[0].Id)
			}
			cmds = append(cmds, tea.WindowSize())
		case "ctrl+pgdown":
			index := m.IndexBuffers[m.ActiveBufferId]
			if index > 0 {
				m = m.ChangeBuffer(m.Buffers[index-1].Id)
			}
			if index == 0 {
				m = m.ChangeBuffer(m.Buffers[len(m.Buffers)-1].Id)
			}

			cmds = append(cmds, tea.WindowSize())
		case "ctrl+x":
			m, cmd = m.DeleteBuffer(m.ActiveBufferId)

			cmds = append(cmds, cmd)

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
		case "esc":
			m.BufferModalModel.OverlayActive = false
		case "ctrl+t":
			m.BufferModalModel.OverlayActive = !m.BufferModalModel.OverlayActive
			m.BufferModalModel.BufferHovered = m.ActiveBufferId
			m = m.BlurAll()
		}
	}

	if m.BufferModalModel.OverlayActive && m.BufferModalModel.ViewportReady {
		m, cmd = m.BufferNavigation(msg)
		cmds = append(cmds, cmd)

		buffersComponent := []string{}

		bufferStyle := lipgloss.NewStyle().
			MaxWidth(m.WindowWidth*50/100 - utils.BoxStyle.GetHorizontalBorderSize()).
			Width(m.WindowWidth*50/100 - utils.BoxStyle.GetHorizontalBorderSize())

		for _, buffer := range m.Buffers {
			var (
				isActive    = buffer.Id == m.ActiveBufferId
				isHovered   = buffer.Id == m.BufferModalModel.BufferHovered
				tempStyle   = bufferStyle
				method      string
				url         string
				methodStyle lipgloss.Style
			)

			if isActive {
				method = m.MethodModel.ActiveState.String()
				url = m.Model.UrlModel.UrlInput.Value()
			} else {
				method = buffer.Model.MethodModel.ActiveState.String()
				url = buffer.Model.UrlModel.UrlInput.Value()
			}

			switch method {
			case "GET":
				methodStyle = methodModel.GetStyle
			case "POST":
				methodStyle = methodModel.PostStyle
			case "PUT":
				methodStyle = methodModel.PutStyle
			case "PATCH":
				methodStyle = methodModel.PatchStyle
			case "DELETE":
				methodStyle = methodModel.DeleteStyle
			}

			if isHovered {
				tempStyle = tempStyle.Background(lipgloss.Color(utils.GrayColor))
				methodStyle = methodStyle.Background(lipgloss.Color(utils.GrayColor))
			}

			buffersComponent = append(buffersComponent, zone.Mark(buffer.Id, lipgloss.JoinHorizontal(lipgloss.Left,
				methodStyle.Render(method+" "),
				tempStyle.Render(ansi.Truncate(url, m.WindowWidth*48/100, "")),
			)),
			)
		}

		m.BufferModalModel.Viewport.SetContent(
			lipgloss.JoinVertical(lipgloss.Left, buffersComponent...),
		)
		m.BufferModalModel.Viewport, cmd = m.BufferModalModel.Viewport.Update(msg)
	}

	return m, tea.Batch(cmds...)
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

	if m.MethodModel.OverlayActive || m.HelpModel.OverlayActive || m.BufferModalModel.OverlayActive {
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
