package app

import "github.com/arfadmuzali/restui/internal/url"

type MainModel struct {
	WindowWidth  int
	WindowHeight int

	UrlModel url.UrlModel
}

func InitModel() MainModel {
	return MainModel{UrlModel: url.New()}
}
