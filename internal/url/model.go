package url

import "github.com/charmbracelet/bubbles/textinput"

type UrlModel struct {
	UrlInput textinput.Model
}

func New() UrlModel {
	ti := textinput.New()
	ti.Placeholder = "Enter URL"
	ti.Focus()
	ti.Prompt = ""
	return UrlModel{UrlInput: ti}
}
