package hint

import (
	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
)

type KeyMap struct {
	Send   key.Binding
	URL    key.Binding
	Body   key.Binding
	Header key.Binding
	Method key.Binding
	Buffer key.Binding
	Help   key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Send, k.URL, k.Body, k.Header, k.Method, k.Buffer, k.Help}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Send, k.URL, k.Body, k.Header, k.Method, k.Buffer},
		{k.Help},
	}
}

type HintModel struct {
	Keys KeyMap
	Help help.Model
}

func New() HintModel {
	keys := KeyMap{
		Send: key.NewBinding(
			key.WithKeys("ctrl+enter", "alt+enter"),
			key.WithHelp("ctrl+enter", "Send"),
		),
		URL: key.NewBinding(
			key.WithKeys("ctrl+l"),
			key.WithHelp("ctrl+l", "URL"),
		),
		Body: key.NewBinding(
			key.WithKeys("ctrl+b"),
			key.WithHelp("ctrl+b", "Body"),
		),
		Header: key.NewBinding(
			key.WithKeys("ctrl+h"),
			key.WithHelp("ctrl+h", "Header"),
		),
		Method: key.NewBinding(
			key.WithKeys("ctrl+m", "ctrl+o"),
			key.WithHelp("ctrl+m", "Method"),
		),
		Buffer: key.NewBinding(
			key.WithKeys("ctrl+n"),
			key.WithHelp("ctrl+n", "New"),
		),
		Help: key.NewBinding(
			key.WithKeys("f1"),
			key.WithHelp("f1", "Help"),
		),
	}

	return HintModel{
		Keys: keys,
		Help: help.New(),
	}
}
