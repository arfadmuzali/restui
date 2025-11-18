package restui

import (
	"github.com/arfadmuzali/restui/internal/app"
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

func Execute() error {
	zone.NewGlobal()
	defer zone.Close()
	p := app.InitModel()
	_, err := tea.NewProgram(p, tea.WithAltScreen(), tea.WithMouseCellMotion()).Run()

	return err
}
