package hint

import "strings"

func (m HintModel) View() string {
	return strings.Join(m.shortcuts, " | ")
}
