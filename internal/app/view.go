package app

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/arfadmuzali/restui/internal/method"
	"github.com/arfadmuzali/restui/internal/response"
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
		Align(lipgloss.Left, lipgloss.Top)

	requestSection := lipgloss.JoinVertical(
		lipgloss.Left,
		// TODO: this is dummy header
		" Headers | Body",
		lipgloss.NewStyle().
			Height(bodyHeight-utils.BoxStyle.GetHorizontalBorderSize()-1).
			Width(bodyWidth*40/100-utils.BoxStyle.GetHorizontalBorderSize()).
			Border(lipgloss.RoundedBorder()).Render(fmt.Sprintf("%.2f, responseheight: %v, lencontent: %v",
			m.ResponseModel.Viewport.ScrollPercent(),
			m.ResponseModel.ResponseHeight,
			m.ResponseModel.Viewport.TotalLineCount(),
		)),
	)

	responseView := m.ResponseModel.View()

	var addon int
	if m.WindowWidth%10 != 0 {
		addon = 1
	}

	var hoveredColor string
	if m.ResponseModel.Hovered {
		hoveredColor = utils.BlueColor
	}

	left, right := utils.PrintHorizontalBorder(bodyHeight-utils.BoxStyle.GetHorizontalBorderSize()-1, m.ResponseModel.Viewport.TotalLineCount(), m.ResponseModel.Viewport.ScrollPercent())
	top, bottom := utils.PrintVerticalBorder(bodyWidth*60/100 + addon)

	content := lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.NewStyle().Foreground(lipgloss.Color(hoveredColor)).Render(top),
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			lipgloss.NewStyle().Foreground(lipgloss.Color(hoveredColor)).Render(left),
			responseView,
			lipgloss.NewStyle().Foreground(lipgloss.Color(hoveredColor)).Render(right),
		),
		lipgloss.NewStyle().Foreground(lipgloss.Color(hoveredColor)).Render(bottom),
	)

	var responseContent string
	if m.ResponseModel.IsLoading {
		responseContent = lipgloss.NewStyle().
			Height(bodyHeight-utils.BoxStyle.GetHorizontalBorderSize()-1).
			Width(bodyWidth*60/100-utils.BoxStyle.GetHorizontalBorderSize()+addon).
			Align(lipgloss.Center, lipgloss.Center).
			BorderForeground(lipgloss.Color(hoveredColor)).
			Border(lipgloss.RoundedBorder()).Render(m.spinner.View())
	} else {
		responseContent = lipgloss.NewStyle().Render(zone.Mark("response", content))
	}
	var responseTabs []string

	for i := 0; i < 3; i++ {
		focusedStyle := lipgloss.NewStyle().Padding(0, 1).Foreground(lipgloss.Color(utils.BlueColor))

		switch i {
		case 0:
			if m.ResponseModel.FocusedTab == response.Body {
				responseTabs = append(responseTabs, zone.Mark("responseBody", focusedStyle.Render("Body")))
			} else {
				responseTabs = append(responseTabs, zone.Mark("responseBody", lipgloss.NewStyle().Padding(0, 1).Render("Body")))
			}
		case 1:
			if m.ResponseModel.FocusedTab == response.Headers {
				responseTabs = append(responseTabs, zone.Mark("responseHeaders", focusedStyle.Render("Headers")))
			} else {
				responseTabs = append(responseTabs, zone.Mark("responseHeaders", lipgloss.NewStyle().Padding(0, 1).Render("Headers")))
			}
		case 2:
			if m.ResponseModel.FocusedTab == response.Cookies {
				responseTabs = append(responseTabs, zone.Mark("responseCookies", focusedStyle.Render("Cookies")))
			} else {
				responseTabs = append(responseTabs, zone.Mark("responseCookies", lipgloss.NewStyle().Padding(0, 1).Render("Cookies")))
			}
		}
	}

	var responseStatusCode string
	if m.ResponseModel.Result.StatusCode == 0 {
		responseStatusCode = ""
	} else if m.ResponseModel.Result.StatusCode <= 300 {
		responseStatusCode = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(utils.GreenColor)).Render(strconv.Itoa(m.ResponseModel.Result.StatusCode))
	} else if m.ResponseModel.Result.StatusCode <= 400 {

		responseStatusCode = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(utils.OrangeColor)).Render(strconv.Itoa(m.ResponseModel.Result.StatusCode))
	} else if m.ResponseModel.Result.StatusCode <= 500 {

		responseStatusCode = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(utils.RedColor)).Render(strconv.Itoa(m.ResponseModel.Result.StatusCode))
	}

	var copyButton string
	if m.ResponseModel.FocusedTab == response.Body && m.ResponseModel.Result.Data != nil {
		copyButton = zone.Mark("copyResponseBody", lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(utils.OrangeColor)).Render("Copy-Response"))
	}

	responseSection := lipgloss.JoinVertical(
		lipgloss.Left,

		lipgloss.JoinHorizontal(
			lipgloss.Left,
			lipgloss.NewStyle().Width(bodyWidth*30/100).Align(lipgloss.Left, lipgloss.Center).Render(
				strings.Join(responseTabs, "|"),
			),
			lipgloss.NewStyle().Width(bodyWidth*30/100).Align(lipgloss.Right, lipgloss.Center).Render(
				lipgloss.NewStyle().Bold(true).Padding(0, 1).Render(copyButton, responseStatusCode),
			),
		),

		responseContent,
	)
	return s.Render(
		lipgloss.JoinHorizontal(lipgloss.Center, requestSection, responseSection),
	)
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
		Foreground(lipgloss.Color(utils.BlueColor)).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(utils.BlueColor))

	// XXX: dunno why but i have to add the widht by 1
	if m.WindowWidth%10 != 0 {
		sendButton = lipgloss.NewStyle().
			Width(m.WindowWidth*10/100-utils.BoxStyle.GetHorizontalBorderSize()+1).
			Align(lipgloss.Center, lipgloss.Top).
			Foreground(lipgloss.Color(utils.BlueColor)).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(utils.BlueColor))
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

	return urlSection.Render(lipgloss.JoinHorizontal(lipgloss.Left, URLAndMethod, zone.Mark("send", sendButton.Render("SEND"))))
}
