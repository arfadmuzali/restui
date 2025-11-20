package restui

import (
	"github.com/arfadmuzali/restui/internal/app"
	"github.com/arfadmuzali/restui/internal/config"
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

func Execute() error {
	err := config.ConfigInitialization()
	if err != nil {
		return err
	}

	zone.NewGlobal()
	defer zone.Close()
	p := app.InitModel()
	_, err = tea.NewProgram(p, tea.WithAltScreen(), tea.WithMouseCellMotion()).Run()

	return err
}
