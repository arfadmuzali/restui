package app

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/arfadmuzali/restui/internal/method"
	"github.com/arfadmuzali/restui/internal/request"
	"github.com/arfadmuzali/restui/internal/response"
	"github.com/arfadmuzali/restui/internal/utils"
	zone "github.com/lrstanley/bubblezone/v2"
)

func (m MainModel) View() tea.View {

	minWindowWidth := 62
	minWindowHeight := 31

	var v tea.View
	v.AltScreen = true
	v.MouseMode = tea.MouseModeCellMotion

	if m.WindowWidth < minWindowWidth || m.WindowHeight < minWindowHeight {
		wrapper := lipgloss.NewStyle().
			Align(lipgloss.Center, lipgloss.Center).
			Height(m.WindowHeight).
			Width(m.WindowWidth)
		v.SetContent(wrapper.Render(lipgloss.JoinVertical(
			lipgloss.Center,
			"Terminal size is too small",
			fmt.Sprintf("%v < %d x %v < %d",
				lipgloss.NewStyle().Foreground(lipgloss.Color(utils.RedColor)).Render(strconv.Itoa(m.WindowWidth)),
				minWindowWidth,
				lipgloss.NewStyle().Foreground(lipgloss.Color(utils.RedColor)).Render(strconv.Itoa(m.WindowHeight)),
				minWindowHeight,
			),
		)))
		return v
	}

	mainWrapper := lipgloss.NewStyle().
		Align(lipgloss.Center, lipgloss.Top).
		Height(m.WindowHeight).
		Width(m.WindowWidth)

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
		v.SetContent(zone.Scan(Render(layout, m.MethodModel.View())))
		return v
	} else if m.HelpModel.OverlayActive {
		v.SetContent(zone.Scan(Render(layout, m.HelpModel.View())))
		return v
	} else if m.BufferModalModel.OverlayActive {
		OverlayModal := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Render(m.BufferModalModel.Viewport.View())
		v.SetContent(zone.Scan(Render(layout, OverlayModal)))
		return v
	}

	v.SetContent(zone.Scan(layout))

	return v
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

	// Request Section (left section)
	var requestTabs []string

	var requestHoveredColor string

	if m.RequestModel.Hovered {
		requestHoveredColor = utils.BlueColor
	}

	for i := 0; i < 2; i++ {
		focusedTabStyle := lipgloss.NewStyle().Padding(0, 1).Foreground(lipgloss.Color(utils.BlueColor))
		switch i {
		case 0:
			if m.RequestModel.FocusedTab == request.Body {
				requestTabs = append(requestTabs, zone.Mark("requestBody", focusedTabStyle.Render("Body")))
			} else {
				requestTabs = append(requestTabs, zone.Mark("requestBody", lipgloss.NewStyle().Padding(0, 1).Render("Body")))
			}
		case 1:
			if m.RequestModel.FocusedTab == request.Headers {
				requestTabs = append(requestTabs, zone.Mark("requestHeaders", focusedTabStyle.Render("Headers")))
			} else {
				requestTabs = append(requestTabs, zone.Mark("requestHeaders", lipgloss.NewStyle().Padding(0, 1).Render("Headers")))
			}
		}
	}

	// Use widths that models compute so inner view sizes match their content
	borderH := utils.BoxStyle.GetHorizontalBorderSize()

	// request outer width as computed in request.Update
	requestOuter := m.RequestModel.RequestWidth
	if requestOuter <= 0 {
		requestOuter = int(math.Round(float64(bodyWidth)*0.4)) - borderH
		requestOuter = max(requestOuter, 1)
	}

	// response inner width (viewport) as computed in response.Update
	responseInner := m.ResponseModel.ResponseWidth
	if responseInner <= 0 {
		// fallback: remaining width after request and border
		responseInner = bodyWidth - requestOuter - borderH
		responseInner = max(responseInner, 1)
	}

	requestSection := lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.NewStyle().Align(lipgloss.Left).Render(strings.Join(requestTabs, "|")),
		lipgloss.NewStyle().
			Height(bodyHeight-utils.BoxStyle.GetHorizontalBorderSize()-1).
			Width(requestOuter).
			BorderForeground(lipgloss.Color(requestHoveredColor)).
			Border(lipgloss.RoundedBorder()).Render(m.RequestModel.View()),
	)
	// requestSection = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Width(bodyWidth*40/100 - 3).Height(bodyHeight - 3).Render("")

	// Response Section (right section)

	var responseHoveredColor string
	if m.ResponseModel.Hovered {
		responseHoveredColor = utils.BlueColor
	}

	// minus 1 for the tabs
	left, right := utils.PrintHorizontalBorder(bodyHeight-utils.BoxStyle.GetHorizontalBorderSize()-1, m.ResponseModel.Viewport.TotalLineCount(), m.ResponseModel.Viewport.ScrollPercent())
	// compute outer width for response (inner viewport width + border size)
	responseOuter := responseInner + borderH
	top, bottom := utils.PrintVerticalBorder(responseOuter)

	content := lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.NewStyle().Foreground(lipgloss.Color(responseHoveredColor)).Render(top),
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			lipgloss.NewStyle().Foreground(lipgloss.Color(responseHoveredColor)).Render(left),
			m.ResponseModel.View(),
			lipgloss.NewStyle().Foreground(lipgloss.Color(responseHoveredColor)).Render(right),
		),
		lipgloss.NewStyle().Foreground(lipgloss.Color(responseHoveredColor)).Render(bottom),
	)

	var responseContent string
	if m.ResponseModel.IsLoading {
		responseContent = lipgloss.NewStyle().
			// minus 1 for the tabs
			Height(bodyHeight-1).
			Width(responseInner).
			Align(lipgloss.Center, lipgloss.Center).
			BorderForeground(lipgloss.Color(responseHoveredColor)).
			Border(lipgloss.RoundedBorder()).Render(m.spinner.View())
	} else {
		responseContent = lipgloss.NewStyle().Width(responseOuter).Render(zone.Mark("response", content))
	}

	var responseTabs []string
	for i := 0; i < 2; i++ {
		focusedStyle := lipgloss.NewStyle().Padding(0, 1).Foreground(lipgloss.Color(utils.BlueColor))

		switch i {
		case 0:
			if m.ResponseModel.FocusedTab == response.Body {
				responseTabs = append(responseTabs, zone.Mark("responseBody", focusedStyle.Render("Response Body")))
			} else {
				responseTabs = append(responseTabs, zone.Mark("responseBody", lipgloss.NewStyle().Padding(0, 1).Render("Response Body")))
			}
		case 1:
			if m.ResponseModel.FocusedTab == response.Headers {
				responseTabs = append(responseTabs, zone.Mark("responseHeaders", focusedStyle.Render("Headers")))
			} else {
				responseTabs = append(responseTabs, zone.Mark("responseHeaders", lipgloss.NewStyle().Padding(0, 1).Render("Headers")))
			}
			// case 2:
			// 	if m.ResponseModel.FocusedTab == response.Cookies {
			// 		responseTabs = append(responseTabs, zone.Mark("responseCookies", focusedStyle.Render("Cookies")))
			// 	} else {
			// 		responseTabs = append(responseTabs, zone.Mark("responseCookies", lipgloss.NewStyle().Padding(0, 1).Render("Cookies")))
			// 	}
		}
	}

	responseTabsSection := strings.Join(responseTabs, "|")
	responseSection := lipgloss.JoinVertical(
		lipgloss.Left,

		lipgloss.JoinHorizontal(
			lipgloss.Left,
			lipgloss.NewStyle().Align(lipgloss.Left, lipgloss.Center).Render(responseTabsSection),
			// lipgloss.NewStyle().Width(bodyWidth*35/100).Align(lipgloss.Right, lipgloss.Center).Background(lipgloss.Color(utils.BlueColor)).Render(
			// 	lipgloss.NewStyle().Bold(true).Padding(0, 1).Render(copyButton, responseStatusCode),
			// ),
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
	// compute send button and method widths using rounding
	sendBtnWidth := int(math.Round(float64(m.WindowWidth) * 0.1))
	sendBtnWidth = max(sendBtnWidth, 1)

	method := lipgloss.NewStyle().Width(sendBtnWidth+1).Align(lipgloss.Center, lipgloss.Center).Foreground(lipgloss.Color(color))

	sendButton := lipgloss.NewStyle().
		Width(sendBtnWidth).
		Align(lipgloss.Center, lipgloss.Top).
		Foreground(lipgloss.Color(utils.BlueColor)).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(utils.BlueColor))

	separator := utils.Separator.Render(utils.Line.Foreground(lipgloss.Color(utils.WhiteColor)).Render())

	if m.UrlModel.UrlInput.Focused() {
		separator = utils.Separator.Render(utils.Line.BorderForeground(lipgloss.Color(utils.BlueColor)).Render())
	}
	URLAndMethod := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Width(m.WindowWidth-sendBtnWidth).
		Render(
			zone.Mark("method", method.Render(utils.BoldStyle.Render(m.MethodModel.ActiveState.String()))),
			separator,
			m.UrlModel.View(),
		)

	return urlSection.MaxHeight(3).Render(lipgloss.JoinHorizontal(lipgloss.Left, URLAndMethod, zone.Mark("send", sendButton.Render("SEND"))))
}
