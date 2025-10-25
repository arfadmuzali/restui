package url

import "github.com/charmbracelet/bubbles/textinput"

type UrlModel struct {
	windowWidth int
	UrlInput    textinput.Model
}

func New() UrlModel {
	ti := textinput.New()
	ti.Placeholder = "Enter URL"
	ti.Focus()
	ti.Prompt = ""
	ti.Width = 20
	return UrlModel{UrlInput: ti}
}
