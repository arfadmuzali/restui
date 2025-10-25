package main

import (
	"github.com/arfadmuzali/restui/internal/app"
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
	"log"
)

func main() {
	zone.NewGlobal()
	defer zone.Close()
	p := app.InitModel()
	_, err := tea.NewProgram(p, tea.WithAltScreen(), tea.WithMouseCellMotion()).Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}
