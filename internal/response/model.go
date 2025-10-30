package response

import (
	"net/http"

	"github.com/charmbracelet/bubbles/viewport"
)

type ResultMsg struct {
	Header http.Header
	Data   []byte
	Error  error
}

type IsLoadingMsg bool

type ResponseModel struct {
	ResponseWidth  int
	ResponseHeight int
	Result         ResultMsg
	IsLoading      bool

	Viewport      viewport.Model
	ViewportReady bool
	Hovered       bool
}

func New() ResponseModel {
	return ResponseModel{}
}
