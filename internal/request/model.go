package request

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

type RequestTab int

const (
	Body RequestTab = iota
	Headers
)

func (r RequestTab) String() string {
	switch r {
	case Body:
		return "Body"
	case Headers:
		return "Headers"
	default:
		return "Unknown"

	}
}

type RequestModel struct {
	Hovered       bool
	Viewport      viewport.Model
	ViewportReady bool
	TextArea      textarea.Model

	FocusedTab RequestTab

	RequestHeight int
	RequestWidth  int
}

func New() RequestModel {
	ta := textarea.New()
	ta.Placeholder = "Enter your request body, please define your Content-Type on Headers section"
	ta.ShowLineNumbers = false
	ta.KeyMap.WordBackward.SetKeys("alt+left", "ctrl+left")
	ta.KeyMap.WordForward.SetKeys("alt+right", "ctrl+right")
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()
	ta.Prompt = ""
	ta.SetWidth(20)

	return RequestModel{FocusedTab: Body, TextArea: ta}
}
