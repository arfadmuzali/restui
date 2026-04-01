package help

import (
	"charm.land/lipgloss/v2"
	"github.com/arfadmuzali/restui/internal/utils"
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
