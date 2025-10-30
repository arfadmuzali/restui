package utils

import (
	"math"
	"strings"

	"github.com/TylerBrock/colorjson"
	"github.com/charmbracelet/lipgloss"
	"github.com/fatih/color"
)

func ThumbHeightFromProgress(height, TotalLine int) int {
	heightSquare := height * height
	return int(math.Max(1, math.Round(float64(heightSquare)/float64(TotalLine))))
}

var BoxStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder())

var Line = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).BorderRight(true).BorderLeft(false).BorderTop(false).BorderBottom(false)
var Separator = lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center)

var BoldStyle = lipgloss.NewStyle().Bold(true)

func RenderSeparator() string {
	return Separator.Render(Line.Render())
}

var Formatter = colorjson.Formatter{
	Indent:      2,
	KeyColor:    color.New(color.FgBlue),
	StringColor: color.New(color.FgGreen),
	NumberColor: color.New(color.FgCyan),
	BoolColor:   color.New(color.FgYellow),
}

func PrintHorizontalBorder(height, totalLineCount int, scrollPercent float64) (leftBorder, rightBorder string) {
	var b strings.Builder

	for i := 0; i < height; i++ {
		b.WriteString("│")
		if i < height-1 {
			b.WriteByte('\n')
		}
	}
	leftBorder = b.String()
	b.Reset()

	thumbHeight := int(math.Max(1, math.Round(float64(height*height)/float64(totalLineCount))))
	thumbTop := int(math.Round(scrollPercent * float64(height-thumbHeight)))

	for i := 0; i < height; i++ {
		if i >= thumbTop && i < thumbTop+thumbHeight {
			b.WriteString("▐")
		} else {
			b.WriteString("│")
		}
		if i < height-1 {
			b.WriteByte('\n')
		}
	}
	rightBorder = b.String()

	return
}

func PrintVerticalBorder(width int) (topBorder, bottomBorder string) {
	for col := 0; col < width; col++ {
		switch col {
		case 0:
			topBorder += "╭"
		case width - 1:
			topBorder += "╮"
		default:
			topBorder += "─"
		}
	}

	for col := 0; col < width; col++ {
		switch col {
		case 0:
			bottomBorder += "╰"
		case width - 1:
			bottomBorder += "╯"
		default:
			bottomBorder += "─"
		}
	}
	return topBorder, bottomBorder
}
