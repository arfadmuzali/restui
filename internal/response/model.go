package response

import "github.com/charmbracelet/bubbles/viewport"

type ResultMsg struct {
	Data  []byte
	Error error
}

type IsLoadingMsg bool

type ResponseModel struct {
	ResponseWidth  int
	ResponseHeight int
	Result         ResultMsg
	IsLoading      bool
	Viewport       viewport.Model
	ViewportReady  bool
}

func New() ResponseModel {
	return ResponseModel{}
}
