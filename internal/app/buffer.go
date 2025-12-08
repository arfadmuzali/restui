package app

import (
	"github.com/arfadmuzali/restui/internal/help"
	"github.com/arfadmuzali/restui/internal/hint"
	"github.com/arfadmuzali/restui/internal/method"
	"github.com/arfadmuzali/restui/internal/request"
	"github.com/arfadmuzali/restui/internal/response"
	"github.com/arfadmuzali/restui/internal/url"
	"github.com/arfadmuzali/restui/internal/utils"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/google/uuid"
)

type Buffer struct {
	Id    string
	Model Model
}

type BufferModalModel struct {
	OverlayActive bool
	Viewport      viewport.Model
	ViewportReady bool
	BufferHovered string
}

func (m MainModel) BufferNavigation(msg tea.Msg) (MainModel, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		index, ok := m.IndexBuffers[m.BufferModalModel.BufferHovered]
		switch msg.String() {
		case "ctrl+d":
			var cmd tea.Cmd
			if ok {
				m, cmd = m.DeleteBuffer(m.BufferModalModel.BufferHovered)
				m.BufferModalModel.BufferHovered = m.ActiveBufferId
			}
			return m, cmd
		case "up", "k":
			if index >= 0 {
				if index == 0 {
					index = len(m.Buffers)
				}
				index = index - 1
			}
			m.BufferModalModel.BufferHovered = m.Buffers[index].Id
		case "down", "j":
			if index <= len(m.Buffers)-1 {
				if index == len(m.Buffers)-1 {
					index = -1
				}
				index = index + 1
			}
			m.BufferModalModel.BufferHovered = m.Buffers[index].Id
		case "enter":
			if ok {
				m = m.ChangeBuffer(m.BufferModalModel.BufferHovered)
			}
			m.BufferModalModel.OverlayActive = false
			return m, nil
		}
	}
	return m, nil
}

// save old buffer and change the active buffer into new buffer
func (m MainModel) ChangeBuffer(id string) MainModel {
	oldIndex, ok := m.IndexBuffers[m.ActiveBufferId]
	if ok && oldIndex < len(m.Buffers) {
		oldBuffer := Buffer{
			Id: m.ActiveBufferId,
			Model: Model{
				UrlModel:      m.UrlModel,
				HintModel:     m.HintModel,
				MethodModel:   m.MethodModel,
				ResponseModel: m.ResponseModel,
				RequestModel:  m.RequestModel,
				HelpModel:     m.HelpModel,
				spinner:       m.spinner,
			},
		}
		m.Buffers[oldIndex] = oldBuffer
	}

	m.ActiveBufferId = id
	m.Model = m.Buffers[m.IndexBuffers[id]].Model

	return m
}

func (m MainModel) DeleteBuffer(id string) (MainModel, tea.Cmd) {
	index, ok := m.IndexBuffers[id]
	if !ok {
		return m, nil
	}

	deletingActive := (id == m.ActiveBufferId)

	m.Buffers = append(m.Buffers[:index], m.Buffers[index+1:]...)

	delete(m.IndexBuffers, id)
	for i := index; i < len(m.Buffers); i++ {
		m.IndexBuffers[m.Buffers[i].Id] = i
	}

	if len(m.Buffers) == 0 {
		return m, tea.Quit
	}

	var newActive string
	if deletingActive {
		if index == 0 {
			newActive = m.Buffers[0].Id
			// Becarefull touching this code
		} else if index >= len(m.Buffers) {
			newActive = m.Buffers[len(m.Buffers)-1].Id
		} else {
			newActive = m.Buffers[index].Id
		}
	} else {
		newActive = m.ActiveBufferId
	}

	m = m.ChangeBuffer(newActive)

	return m, nil
}

func CreateNewBuffer() Buffer {
	s := spinner.New()
	s.Spinner = spinner.Points
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color(utils.BlueColor))

	id := uuid.New()

	model := Buffer{
		Id: id.String(),
		Model: Model{
			UrlModel:      url.New(),
			HintModel:     hint.New(),
			MethodModel:   method.New(),
			ResponseModel: response.New(),
			RequestModel:  request.New(),
			HelpModel:     help.New(),
			spinner:       s,
		},
	}

	return model
}
