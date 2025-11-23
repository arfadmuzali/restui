package help

import (
	"github.com/arfadmuzali/restui/internal/utils"
	"github.com/charmbracelet/lipgloss"
)

func (m HelpModel) View() string {

	left, right := utils.PrintHorizontalBorder(
		m.helpWindowHeight,
		m.Viewport.TotalLineCount(),
		m.Viewport.ScrollPercent(),
	)
	top, bottom := utils.PrintVerticalBorder(m.helpWindowWidth)

	content := lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.NewStyle().Render(top),
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			lipgloss.NewStyle().Render(left),
			m.Viewport.View(),
			lipgloss.NewStyle().Render(right),
		),
		lipgloss.NewStyle().Render(bottom),
	)

	return content
}
