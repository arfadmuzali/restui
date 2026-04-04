package hint

func (m HintModel) View() string {
	s := m.Help.View(m.Keys)
	return s
}
