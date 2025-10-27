package app

import (
	"fmt"

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

	bodyHeight := m.WindowHeight * 90 / 100
	bodyWidth := m.WindowWidth

	s := lipgloss.NewStyle().
		Height(bodyHeight).
		Width(bodyWidth).
		// Border(lipgloss.RoundedBorder()).
		Align(lipgloss.Left, lipgloss.Top)
	var xs []string

	for i := 0; i < (bodyHeight)*(bodyWidth); i++ {
		if i%2 == 0 {
			xs = append(xs, "o")
		} else {
			xs = append(xs, "x")
		}
	}

	return s.Render(fmt.Sprintf("%v %v, input w %v", bodyHeight, bodyWidth, m.UrlModel.UrlInput.Width))
}

func header(m MainModel) string {
	urlSection := lipgloss.NewStyle().
		Align(lipgloss.Center, lipgloss.Center).Width(m.WindowWidth)

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
	// width add 1 (one) cause i use separator
	method := lipgloss.NewStyle().Width(m.WindowWidth*10/100+1).Align(lipgloss.Center, lipgloss.Center).Foreground(lipgloss.Color(color))

	sendButton := lipgloss.NewStyle().
		Width(m.WindowWidth*10/100-utils.BoxStyle.GetHorizontalBorderSize()).
		Align(lipgloss.Center, lipgloss.Top).
		Foreground(lipgloss.Color("#1971c2")).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#1971c2"))

	// dunno why but i have to add the widht by 1
	if m.WindowWidth%10 != 0 {
		sendButton = lipgloss.NewStyle().
			Width(m.WindowWidth*10/100-utils.BoxStyle.GetHorizontalBorderSize()+1).
			Align(lipgloss.Center, lipgloss.Top).
			Foreground(lipgloss.Color("#1971c2")).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#1971c2"))
	}

	separator := utils.Separator.Render(utils.Line.Foreground(lipgloss.Color(utils.WhiteColor)).Render())

	if m.UrlModel.UrlInput.Focused() {
		separator = utils.Separator.Render(utils.Line.BorderForeground(lipgloss.Color(utils.BlueColor)).Render())
	}
	URLAndMethod := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Width(m.WindowWidth*90/100-utils.BoxStyle.GetHorizontalBorderSize()).
		Render(
			zone.Mark("method", method.Render(utils.BoldStyle.Render(m.MethodModel.ActiveState.String()))),
			separator,
			m.UrlModel.View(),
		)

	return urlSection.Render(lipgloss.JoinHorizontal(lipgloss.Left, URLAndMethod, sendButton.Render("SEND")))
}
