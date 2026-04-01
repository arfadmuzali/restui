package restui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/arfadmuzali/restui/internal/app"
	"github.com/arfadmuzali/restui/internal/config"
	zone "github.com/lrstanley/bubblezone/v2"
)

func Execute() error {
	err := config.ConfigInitialization()
	if err != nil {
		return err
	}

	zone.NewGlobal()
	defer zone.Close()
	p := app.InitModel()
	_, err = tea.NewProgram(p).Run()

	return err
}
