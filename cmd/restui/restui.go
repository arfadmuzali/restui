package restui

import (
	"log"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/arfadmuzali/restui/internal/app"
	"github.com/arfadmuzali/restui/internal/config"
	zone "github.com/lrstanley/bubblezone/v2"
)

func Execute() error {
	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			log.Fatal("fatal:", err)
		}
		defer f.Close()
	}
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
