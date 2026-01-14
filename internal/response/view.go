package response

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/arfadmuzali/restui/internal/utils"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

func (m ResponseModel) View() string {
	// to prevent layout breakage. 45 is the minimum width when responseStatusCode, responseStatusText, and responseTime are combined
	responseStatusText := http.StatusText(m.Result.StatusCode)
	if m.ResponseWidth < 45 {
		responseStatusText = ""
	}

	var responseStatusCode string
	if m.Result.StatusCode == 0 {
		responseStatusCode = ""
	} else if m.Result.StatusCode < 300 {
		responseStatusCode = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(utils.GreenColor)).
			Render(strconv.Itoa(m.Result.StatusCode), responseStatusText)
	} else if m.Result.StatusCode < 400 {
		responseStatusCode = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(utils.OrangeColor)).
			Render(strconv.Itoa(m.Result.StatusCode), responseStatusText)
	} else {
		responseStatusCode = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(utils.RedColor)).
			Render(strconv.Itoa(m.Result.StatusCode), responseStatusText)
	}

	var copyButton string
	if m.FocusedTab == Body && m.Result.Data != nil {
		copyButton = zone.Mark("copyResponseBody", lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(utils.OrangeColor)).Render("Copy"))
	}
	var components string
	if responseStatusCode != "" && m.ResponseTime != "" {
		components = fmt.Sprintf("%v · %v · %v", responseStatusCode, m.ResponseTime, copyButton)
	}

	layout := lipgloss.JoinVertical(
		lipgloss.Right,
		m.Viewport.View(),
		lipgloss.NewStyle().Height(1).AlignHorizontal(lipgloss.Center).Render(components),
	)
	return layout
}
