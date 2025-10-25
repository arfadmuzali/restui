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

	s := lipgloss.Place(
		m.WindowWidth,
		m.WindowHeight,
		lipgloss.Center,
		lipgloss.Center,
		mainWrapper.Render(
			lipgloss.JoinVertical(
				lipgloss.Center,
				header(m),
				body(m),
				footer(m)),
		),
	)

	return zone.Scan(s)
}

func footer(m MainModel) string {

	bodyHeight := 1
	bodyWidth := m.WindowWidth
	s := lipgloss.NewStyle().Align(lipgloss.Left, lipgloss.Bottom).Height(bodyHeight).Width(bodyWidth)

	return s.Render(m.HintModel.View())
}

func body(m MainModel) string {
	bodyHeight := m.WindowHeight*90/100 - utils.BoxStyle.GetVerticalBorderSize()
	bodyWidth := m.WindowWidth - utils.BoxStyle.GetHorizontalBorderSize()
	s := lipgloss.NewStyle().Height(bodyHeight).Width(bodyWidth).Border(lipgloss.HiddenBorder()).Align(lipgloss.Center, lipgloss.Center)
	return s.Render("2")
}

func header(m MainModel) string {
	urlSection := lipgloss.NewStyle().
		Align(lipgloss.Center, lipgloss.Center).Width(m.WindowWidth).Height(m.WindowHeight * 5 / 100)

	//TODO: dummy method, move this into its module

	dummyMethod := lipgloss.NewStyle().Width(m.WindowWidth*10/100).Align(lipgloss.Center, lipgloss.Center)

	//TODO: dummy send button, move this into its module

	dummySendButton := lipgloss.NewStyle().
		Width(m.WindowWidth*10/100-utils.BoxStyle.GetHorizontalBorderSize()).
		Align(lipgloss.Center, lipgloss.Center).
		Foreground(lipgloss.Color("#1971c2")).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#1971c2"))

	URLAndMethod := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Width(m.WindowWidth*90/100-utils.BoxStyle.GetHorizontalBorderSize()).
		Render(dummyMethod.Render("DELET"),
			utils.RenderSeparator(),
			m.UrlModel.View(),
		)

	return urlSection.Render(lipgloss.JoinHorizontal(lipgloss.Top, URLAndMethod, dummySendButton.Render("SEND")))
}
