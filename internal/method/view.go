package method

import (
	"github.com/arfadmuzali/restui/internal/utils"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

func (m MethodModel) View() string {

	componentWidth := m.windowWidth * 20 / 100

	var getComponent = lipgloss.NewStyle().
		Bold(true).
		Padding(1, 1).
		Foreground(lipgloss.Color(utils.GreenColor))
	var patchComponent = lipgloss.
		NewStyle().
		Bold(true).
		Padding(1, 1).
		Foreground(lipgloss.Color(utils.PurpleColor))
	var postComponent = lipgloss.
		NewStyle().
		Bold(true).
		Padding(1, 1).
		Foreground(lipgloss.Color(utils.OrangeColor))
	var putComponent = lipgloss.
		NewStyle().
		Bold(true).
		Padding(1, 1).
		Foreground(lipgloss.Color(utils.BlueColor))
	var deleteComponent = lipgloss.
		NewStyle().
		Bold(true).
		Padding(1, 1).
		Foreground(lipgloss.Color(utils.RedColor))

	switch m.ActiveState.String() {
	case "GET":
		getComponent = lipgloss.NewStyle().
			Bold(true).
			Padding(1, 1).
			Foreground(lipgloss.Color(utils.WhiteColor))

	case "POST":
		postComponent = lipgloss.NewStyle().
			Bold(true).
			Padding(1, 1).
			Foreground(lipgloss.Color(utils.WhiteColor))
	case "PUT":
		putComponent = lipgloss.NewStyle().
			Bold(true).
			Padding(1, 1).
			Foreground(lipgloss.Color(utils.WhiteColor))
	case "PATCH":
		patchComponent = lipgloss.NewStyle().
			Bold(true).
			Padding(1, 1).
			Foreground(lipgloss.Color(utils.WhiteColor))
	case "DELETE":
		deleteComponent = lipgloss.NewStyle().
			Bold(true).
			Padding(1, 1).
			Foreground(lipgloss.Color(utils.WhiteColor))
	}

	layout := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Align(lipgloss.Center, lipgloss.Top)
	components := lipgloss.JoinVertical(lipgloss.Center,
		"Select Method",
		zone.Mark("GET",
			lipgloss.JoinHorizontal(lipgloss.Center, getComponent.Width(componentWidth).Render("GET"), getComponent.Render("(G)")),
		),
		zone.Mark("POST",
			lipgloss.JoinHorizontal(lipgloss.Center, postComponent.Width(componentWidth).Render("POST"), postComponent.Render("(P)")),
		),
		zone.Mark("PUT",
			lipgloss.JoinHorizontal(lipgloss.Center, putComponent.Width(componentWidth).Render("PUT"), putComponent.Render("(U)")),
		),
		zone.Mark("PATCH",
			lipgloss.JoinHorizontal(lipgloss.Center, patchComponent.Width(componentWidth).Render("PATCH"), patchComponent.Render("(A)")),
		),
		zone.Mark("DELETE",
			lipgloss.JoinHorizontal(lipgloss.Center, deleteComponent.Width(componentWidth).Render("DELETE"), deleteComponent.Render("(D)")),
		),
		lipgloss.NewStyle().Render("Close: ESC"),
	)

	return layout.Render(components)
}
