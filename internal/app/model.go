package app

import (
	"github.com/arfadmuzali/restui/internal/hint"
	"github.com/arfadmuzali/restui/internal/url"
)

type MainModel struct {
	WindowWidth  int
	WindowHeight int

	UrlModel  url.UrlModel
	HintModel hint.HintModel
}

func InitModel() MainModel {
	return MainModel{UrlModel: url.New(), HintModel: hint.New()}
}
