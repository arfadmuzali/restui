package app

import (
	"github.com/arfadmuzali/restui/internal/hint"
	"github.com/arfadmuzali/restui/internal/method"
	"github.com/arfadmuzali/restui/internal/url"
)

type MainModel struct {
	WindowWidth  int
	WindowHeight int

	UrlModel    url.UrlModel
	HintModel   hint.HintModel
	MethodModel method.MethodModel
}

func InitModel() MainModel {

	model := MainModel{
		UrlModel:    url.New(),
		HintModel:   hint.New(),
		MethodModel: method.New(),
	}

	return model
}

func (m MainModel) BlurAllInput() MainModel {
	m.UrlModel.UrlInput.Blur()
	return m
}
