package utils

import "github.com/charmbracelet/lipgloss"

var BoxStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder())

var Line = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).BorderRight(true).BorderLeft(false).BorderTop(false).BorderBottom(false)
var Separator = lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center)

func BodyHeight(windowHeight int) int {
	return windowHeight*90/100 - BoxStyle.GetVerticalBorderSize()

}
func BodyWidth(windowWidth int) int {
	return windowWidth - BoxStyle.GetHorizontalBorderSize()
}

var BoldStyle = lipgloss.NewStyle().Bold(true)

func RenderSeparator() string {
	return Separator.Render(Line.Render())
}
