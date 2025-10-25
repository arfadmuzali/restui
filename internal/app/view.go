package app

import (
	"github.com/arfadmuzali/restui/internal/utils"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

func (m MainModel) View() string {
	mainWrapper := lipgloss.NewStyle().
		Align(lipgloss.Center, lipgloss.Top).
		Height(m.WindowHeight).
		Width(m.WindowWidth).
		MaxWidth(m.WindowWidth)

	urlSection := lipgloss.NewStyle().
		Align(lipgloss.Left, lipgloss.Center)

	//TODO: dummy method, move this into its module

	dummyMethod := lipgloss.NewStyle().Width(m.WindowWidth*10/100).Align(lipgloss.Center, lipgloss.Center)

	//TODO: dummy send button, move this into its module

	dummySendButton := lipgloss.NewStyle().
		Width(m.WindowWidth*10/100).
		Align(lipgloss.Center, lipgloss.Center).
		Bold(true).
		Foreground(lipgloss.Color("#1971c2")).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#1971c2"))

	URLAndMethod := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Width(m.WindowWidth*90/100-utils.BoxStyle.GetHorizontalBorderSize()-2).
		Render(dummyMethod.Render("DELETE"), utils.RenderSeparator(), m.UrlModel.View())

	s := lipgloss.Place(
		m.WindowWidth,
		m.WindowHeight,
		lipgloss.Center,
		lipgloss.Center,
		mainWrapper.Render(
			urlSection.Render(lipgloss.JoinHorizontal(lipgloss.Center, URLAndMethod, dummySendButton.Render("SEND"))),
		),
	)

	return zone.Scan(s)
}
