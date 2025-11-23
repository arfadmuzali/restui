package help

import "github.com/charmbracelet/bubbles/viewport"

type HelpModel struct {
	OverlayActive    bool
	helpWindowWidth  int
	helpWindowHeight int
	Viewport         viewport.Model
	ViewportReady    bool
}

func New() HelpModel {

	return HelpModel{OverlayActive: false}
}
