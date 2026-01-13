package help

import "github.com/charmbracelet/bubbles/viewport"

type HelpModel struct {
	OverlayActive    bool
	helpWindowWidth  int
	helpWindowHeight int
	Viewport         viewport.Model
	ViewportReady    bool
}

var Guide = `
# Common Keybinds
- **ctrl+c**: Exit
- **ESC**: Close modal/popup
- **alt+enter**: Send request
- **F1**: Help

# Request Keybinds
- **ctrl+l**: Go to URL
- **ctrl+b**: Go to request body
- **ctrl+f**: Format request body (JSON only)
- **ctrl+o**: Toggle method modal/popup
    - **k/up**: Navigate up in method modal/popup
    - **j/down**: Navigate down in method modal/popup
    - **g**: Switch to GET method
    - **p**: Switch to POST method
    - **u**: Switch to PUT method
    - **a**: Switch to PATCH method
    - **d**: Switch to DELETE method
    - **ESC/enter**: Close the method modal/popup
- **ctrl+h**: Go to request header
    - **tab**: Toggle between header key and header value
        - **enter**: Add header
    - **ctrl+k/up**: Navigate header table up
    - **ctrl+j/down**: Navigate header table down
    - **ctrl+d**: Delete header

# Buffer Management
- **ctrl+n**: New buffer
- **ctrl+x**: Delete active buffer
- **ctrl+pgup**: Move to next buffer
- **ctrl+pgdown**: Move to previous buffer
- **ctrl+t**: Toggle buffer modal/popup
    - **k/up**: Navigate up in buffer modal/popup
    - **j/down**: Navigate down in buffer modal/popup
    - **ctrl+d**: Delete selected buffer
    - **enter**: Open selected buffer
    - **ESC**: Close buffer modal/popup

# Tips & Tricks
- If you want to select and copy text, use copy-mode/selection-mode in your terminal level. For example, Kitty Terminal Emulator has ctrl+shift+h to enter copy-mode.`

func New() HelpModel {

	return HelpModel{OverlayActive: false}
}
