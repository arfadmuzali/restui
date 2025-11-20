package url

import (
	"fmt"

	"github.com/arfadmuzali/restui/internal/config"
	"github.com/charmbracelet/bubbles/textinput"
)

type UrlModel struct {
	UrlInput    textinput.Model
	Suggestions []string
}

func New() UrlModel {
	ti := textinput.New()
	ti.Placeholder = "Enter URL"
	ti.Focus()
	ti.Prompt = ""
	ti.ShowSuggestions = true

	suggestions, err := config.GetSuggestions()
	if err != nil {
		fmt.Println("RESTUI error: ", err.Error())
	}
	ti.SetSuggestions(suggestions)

	return UrlModel{UrlInput: ti}
}
