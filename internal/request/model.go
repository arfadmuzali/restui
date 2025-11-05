package request

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/muesli/reflow/wrap"
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
	TextArea      textarea.Model
	Headers       []Header

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
	headers := []Header{}
	headers = append(headers, Header{Key: "User-Agent", Value: "RESTUI/0.0.1"})
	headers = append(headers, Header{Key: "Accept", Value: "*/*"})

	return RequestModel{FocusedTab: Body, TextArea: ta, Headers: headers}
}

func (m RequestModel) CreateHeadersTable() table.Writer {

	t := table.NewWriter()
	t.Style().Size.WidthMin = m.RequestWidth
	t.Style().Box.UnfinishedRow = ""
	t.Style().Color.RowAlternate = text.Colors{text.BgBlack}
	t.Style().Box = table.BoxStyle{
		PaddingLeft:  " ",
		PaddingRight: " ",
	}
	t.Style().Options = table.Options{
		DrawBorder:      false,
		SeparateColumns: false,
		SeparateHeader:  true,
		SeparateRows:    false,
	}
	for _, header := range m.Headers {
		t.AppendRow(table.Row{wrap.String(header.Key, m.RequestWidth*40/100), wrap.String(header.Value, m.RequestWidth*60/100)})
	}
	return t
}
