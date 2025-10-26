package utils

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
)

// Position defines the positioning type for overlay elements.
type Position int

const (
	Left Position = iota
	Center
	Right
	Top
	Bottom
)

// Composite merges and flattens the background and foreground views into a single view.
// This implementation is based off of the one used by Superfile -
// https://github.com/yorukot/superfile/blob/main/src/pkg/string_function/overplace.go
func Composite(fg, bg string, xPos, yPos Position, xOff, yOff int) string {
	fgWidth, fgHeight := lipgloss.Size(fg)
	bgWidth, bgHeight := lipgloss.Size(bg)

	if fgWidth >= bgWidth && fgHeight >= bgHeight {
		return fg
	}

	x, y := Offsets(fg, bg, xPos, yPos, xOff, yOff)
	x = Clamp(x, 0, bgWidth-fgWidth)
	y = Clamp(y, 0, bgHeight-fgHeight)

	fgLines := Lines(fg)
	bgLines := Lines(bg)
	var sb strings.Builder

	for i, bgLine := range bgLines {
		if i > 0 {
			sb.WriteByte('\n')
		}
		if i < y || i >= y+fgHeight {
			sb.WriteString(bgLine)
			continue
		}

		pos := 0
		if x > 0 {
			left := ansi.Truncate(bgLine, x, "")
			pos = ansi.StringWidth(left)
			sb.WriteString(left)
			if pos < x {
				sb.WriteString(Whitespace(x - pos))
				pos = x
			}
		}

		fgLine := fgLines[i-y]
		sb.WriteString(fgLine)
		pos += ansi.StringWidth(fgLine)

		right := ansi.TruncateLeft(bgLine, pos, "")
		bgWidth := ansi.StringWidth(bgLine)
		rightWidth := ansi.StringWidth(right)
		if rightWidth <= bgWidth-pos {
			sb.WriteString(Whitespace(bgWidth - rightWidth - pos))
		}
		sb.WriteString(right)
	}
	return sb.String()
}

// Offsets calculates the actual vertical and horizontal offsets used to position the foreground
// relative to the background.
func Offsets(fg, bg string, xPos, yPos Position, xOff, yOff int) (int, int) {
	var x, y int
	switch xPos {
	case Center:
		halfBackgroundWidth := lipgloss.Width(bg) / 2
		halfForegroundWidth := lipgloss.Width(fg) / 2
		x = halfBackgroundWidth - halfForegroundWidth
	case Right:
		x = lipgloss.Width(bg) - lipgloss.Width(fg)
	}

	switch yPos {
	case Center:
		halfBackgroundHeight := lipgloss.Height(bg) / 2
		halfForegroundHeight := lipgloss.Height(fg) / 2
		y = halfBackgroundHeight - halfForegroundHeight
	case Bottom:
		y = lipgloss.Height(bg) - lipgloss.Height(fg)
	}

	Debug(
		"X position: "+strconv.Itoa(int(xPos)),
		"Y position: "+strconv.Itoa(int(yPos)),
		"X offset: "+strconv.Itoa(x+xOff),
		"Y offset: "+strconv.Itoa(y+yOff),
		"Background width: "+strconv.Itoa(lipgloss.Width(bg)),
		"Foreground width: "+strconv.Itoa(lipgloss.Width(fg)),
		"Background height: "+strconv.Itoa(lipgloss.Height(bg)),
		"Foreground height: "+strconv.Itoa(lipgloss.Height(fg)),
	)

	return x + xOff, y + yOff
}

// Clamp calculates the lowest possible number between the given boundaries.
func Clamp(v, lower, upper int) int {
	if upper < lower {
		return min(max(v, upper), lower)
	}
	return min(max(v, lower), upper)
}

// Lines normalises any non standard new lines within a string, and then splits and returns a slice
// of strings split on the new lines.
func Lines(s string) []string {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	return strings.Split(s, "\n")
}

// Whitespace returns a string of whitespace characters of the requested length.
func Whitespace(length int) string {
	return strings.Repeat(" ", length)
}

// Debug is a placeholder function for debug logging.
// Override this function to implement custom debug logging.
var Debug = func(...string) {}
