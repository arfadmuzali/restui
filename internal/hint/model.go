package hint

type HintModel struct {
	shortcuts []string
}

func New() HintModel {
	return HintModel{shortcuts: []string{
		"Quit: ^c",
		"Focus URL: ^l",
		"Choose Method: ^h",
		"Focus Body: ^b",
		"Send Request: alt+enter",
	}}
}
