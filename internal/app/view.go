package app

import (
	"strings"

	"github.com/arfadmuzali/restui/internal/method"
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

	layout := lipgloss.Place(
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

	if m.MethodModel.OverlayActive {
		return zone.Scan(Render(layout, m.MethodModel.View()))
	}

	return zone.Scan(layout)
}

func Render(background string, foreground string) string {

	overlayStack := utils.Composite(
		foreground,
		background,
		utils.Center,
		utils.Center,
		0,
		0,
	)
	return overlayStack
}

func footer(m MainModel) string {

	footerHeight := 1
	footerWidth := m.WindowWidth
	s := lipgloss.NewStyle().Align(lipgloss.Left, lipgloss.Bottom).Height(footerHeight).Width(footerWidth)

	return s.Render(m.HintModel.View())
}

func body(m MainModel) string {

	s := lipgloss.NewStyle().
		Height(utils.BodyHeight(m.WindowHeight)).
		Width(utils.BodyWidth(m.WindowWidth)).
		Border(lipgloss.RoundedBorder()).
		Align(lipgloss.Center, lipgloss.Center)
	var xs []string

	for i := 0; i < utils.BodyWidth(m.WindowWidth)*utils.BodyHeight(m.WindowHeight); i++ {
		if i%2 == 0 {
			xs = append(xs, "o")
		} else {
			xs = append(xs, "x")
		}
	}

	return s.Render(strings.Join(xs, ""))
}

func header(m MainModel) string {
	urlSection := lipgloss.NewStyle().
		Align(lipgloss.Center, lipgloss.Center).Width(m.WindowWidth).Height(m.WindowHeight * 5 / 100)

	var color string

	switch m.MethodModel.ActiveState {
	case method.GET:
		color = utils.GreenColor
	case method.POST:
		color = utils.OrangeColor
	case method.PUT:
		color = utils.BlueColor
	case method.PATCH:
		color = utils.PurpleColor
	case method.DELETE:
		color = utils.RedColor
	}

	method := lipgloss.NewStyle().Width(m.WindowWidth*8/100).Align(lipgloss.Center, lipgloss.Center).Foreground(lipgloss.Color(color))

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
		Render(
			zone.Mark("method", method.Render(utils.BoldStyle.Render(m.MethodModel.ActiveState.String()))),
			utils.RenderSeparator(),
			m.UrlModel.View(),
		)

	return urlSection.Render(lipgloss.JoinHorizontal(lipgloss.Top, URLAndMethod, dummySendButton.Render("SEND")))
}
