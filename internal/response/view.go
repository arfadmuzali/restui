package response

func (m ResponseModel) View() string {
	return string(m.Viewport.View())
}
