package response

import "github.com/charmbracelet/lipgloss"

func (m ResponseModel) View() string {

	s := lipgloss.JoinVertical(lipgloss.Left,
		"Response",
	)
	return s
}
