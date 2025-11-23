package request

import (
	"github.com/arfadmuzali/restui/internal/utils"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
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

type Header struct {
	Key   string
	Value string
}

type RequestModel struct {
	Hovered       bool
	Viewport      viewport.Model
	ViewportReady bool

	TextArea     textarea.Model
	TableHeaders table.Model
	KeyInput     textinput.Model
	ValueInput   textinput.Model

	Headers    []Header
	FocusedTab RequestTab

	RequestHeight int
	RequestWidth  int
}

func New() RequestModel {
	ta := textarea.New()
	ta.Placeholder = "Enter your request body, please define your Content-Type on Headers section"
	ta.ShowLineNumbers = true
	ta.KeyMap.WordBackward.SetKeys("alt+left", "ctrl+left")
	ta.KeyMap.WordForward.SetKeys("alt+right", "ctrl+right")
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()
	ta.Prompt = ""
	ta.SetWidth(20)

	keyInput := textinput.New()
	keyInput.Prompt = ""
	keyInput.Placeholder = "Enter key"
	keyInput.ShowSuggestions = true
	keyInput.SetSuggestions([]string{
		"Accept",
		"Accept-Language",
		"Accept-Encoding",
		"User-Agent",
		"Host",
		"Connection",
		"Authorization",
		"Cookie",
		"X-CSRF-Token",
		"X-API-Key",
		"Content-Type",
		"Content-Length",
		"Origin",
		"Referer",
		"Access-Control-Request-Method",
		"Access-Control-Request-Headers",
		"If-Modified-Since",
		"If-None-Match",
		"Range",
		"X-Requested-With",
	})

	valueInput := textinput.New()
	valueInput.Prompt = ""
	valueInput.Placeholder = "Enter value"

	headers := []Header{
		{Key: "User-Agent", Value: "RESTUI/0.0.1"},
		{Key: "Accept", Value: "*/*"},
	}

	tableHeadersValue := make([]table.Row, 0, len(headers))

	for _, h := range headers {
		tableHeadersValue = append(tableHeadersValue, table.Row{h.Key, h.Value})
	}

	tableHeaders := table.New(
		table.WithFocused(true),
		table.WithColumns([]table.Column{{Title: "Key"}, {Title: "Value"}}),
		table.WithRows(tableHeadersValue),
		table.WithKeyMap(table.KeyMap{
			LineUp: key.NewBinding(
				key.WithKeys("up", "ctrl+k"),
				key.WithHelp("↑/k", "up"),
			),
			LineDown: key.NewBinding(
				key.WithKeys("down", "ctrl+j"),
				key.WithHelp("↓", "down"),
			),
		}),
	)

	tableStyle := table.DefaultStyles()
	tableStyle.Header = tableStyle.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color(utils.WhiteColor)).
		BorderBottom(true).
		Bold(true)

	tableStyle.Selected = tableStyle.Selected.
		Background(lipgloss.Color(utils.GrayColor)).
		Foreground(lipgloss.Color(utils.BlueColor)).
		Bold(false)
	tableHeaders.SetStyles(tableStyle)

	return RequestModel{FocusedTab: Body, TextArea: ta, Headers: headers, KeyInput: keyInput, ValueInput: valueInput, TableHeaders: tableHeaders}
}
