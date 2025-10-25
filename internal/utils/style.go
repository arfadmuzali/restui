package utils

import "github.com/charmbracelet/lipgloss"

var BoxStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder())

var Line = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).BorderRight(true).BorderLeft(false).BorderTop(false).BorderBottom(false)
var Separator = lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center).Width(1)

func RenderSeparator() string {
	return Separator.Render(Line.Render())
}
