package response

import (
	"net/http"

	"github.com/charmbracelet/bubbles/viewport"
)

type ResponseTab int

const (
	Body ResponseTab = iota
	Headers
	Cookies
)

func (r ResponseTab) String() string {
	switch r {
	case Body:
		return "Body"
	case Headers:
		return "Headers"
	case Cookies:
		return "Cookies"
	default:
		return "Unknown"
	}
}

type ResultMsg struct {
	Cookies            []*http.Cookie
	StatusCode         int
	Headers            http.Header
	Data               []byte
	Body               string
	Error              error
	ResponseFocusedTab string
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
	FocusedTab    ResponseTab
}

func New() ResponseModel {
	return ResponseModel{FocusedTab: Body}
}
