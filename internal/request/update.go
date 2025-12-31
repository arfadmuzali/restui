package request

import (
	"strings"

	"github.com/arfadmuzali/restui/internal/utils"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

func (m RequestModel) Init() tea.Cmd {
	return textarea.Blink
}

func (m RequestModel) updateTextArea(msg tea.Msg) (textarea.Model, tea.Cmd) {
	var cmd tea.Cmd

	if !m.TextArea.Focused() {
		return m.TextArea, cmd
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case `"`:
			m.TextArea.InsertString(`""`)
			m.TextArea.SetCursor(m.TextArea.LineInfo().CharOffset - 1)
		case `'`:
			m.TextArea.InsertString(`'`)
			m.TextArea.SetCursor(m.TextArea.LineInfo().CharOffset - 1)
		case "`":
			m.TextArea.InsertString("`")
			m.TextArea.SetCursor(m.TextArea.LineInfo().CharOffset - 1)
		case `(`:
			m.TextArea.InsertString(`()`)
			m.TextArea.SetCursor(m.TextArea.LineInfo().CharOffset - 1)
		case `[`:
			m.TextArea.InsertString(`[]`)
			m.TextArea.SetCursor(m.TextArea.LineInfo().CharOffset - 1)
		case `{`:
			m.TextArea.InsertString(`{}`)
			m.TextArea.SetCursor(m.TextArea.LineInfo().CharOffset - 1)
		case "backspace":
			line := strings.Split(m.TextArea.Value(), "\n")[m.TextArea.Line()]

			lineInfo := m.TextArea.LineInfo()
			// col is cursor position
			col := lineInfo.CharOffset

			if col == 0 {
				m.TextArea, cmd = m.TextArea.Update(msg)
				break
			}

			// current is character before cursor
			current := line[col-1]

			isPairChar := current == '`' ||
				current == '[' ||
				current == '{' ||
				current == '(' ||
				current == '"' ||
				current == '\''

			if !isPairChar {
				m.TextArea, cmd = m.TextArea.Update(msg)
				break
			}

			if col == lineInfo.CharWidth-1 {
				m.TextArea, cmd = m.TextArea.Update(msg)
				break
			}

			if current == line[col-1] && current == line[col] {
				m.TextArea, cmd = m.TextArea.Update(
					tea.KeyMsg{Type: tea.KeyBackspace},
				)
				m.TextArea, cmd = m.TextArea.Update(
					tea.KeyMsg{Type: tea.KeyDelete},
				)
				break
			}

			m.TextArea, cmd = m.TextArea.Update(msg)
		default:
			m.TextArea, cmd = m.TextArea.Update(msg)
		}
	}

	return m.TextArea, cmd
}

func (m RequestModel) Update(msg tea.Msg) (RequestModel, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	m.TextArea, cmd = m.updateTextArea(msg)
	cmds = append(cmds, cmd)
	if m.Hovered {
		m.TableHeaders, cmd = m.TableHeaders.Update(msg)
		cmds = append(cmds, cmd)
	}

	if m.FocusedTab == Headers && m.Hovered {
		m.KeyInput, cmd = m.KeyInput.Update(msg)
		cmds = append(cmds, cmd)

		if strings.EqualFold(m.KeyInput.Value(), "Content-Type") {
			m.ValueInput.SetSuggestions([]string{
				"text/plain",
				"text/html",
				"text/css",
				"text/javascript",
				"text/csv",
				"text/xml",

				"application/json",
				"application/xml",
				"application/x-www-form-urlencoded",
				"application/octet-stream",
				"application/pdf",
				"application/zip",
				"application/gzip",
				"application/vnd.ms-excel",
				"application/vnd.openxmlformats-officedocument.wordprocessingml.document",
				"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
				"application/graphql",
				"application/ld+json",

				"image/png",
				"image/jpeg",
				"image/gif",
				"image/webp",
				"image/svg+xml",
				"image/avif",

				"audio/mpeg",
				"audio/ogg",
				"audio/wav",
				"audio/webm",

				"video/mp4",
				"video/webm",
				"video/ogg",

				"multipart/form-data",
				"multipart/mixed",
				"multipart/alternative",
				"multipart/related",

				"font/ttf",
				"font/otf",
				"font/woff",
				"font/woff2",
			})
		} else {
			// m.ValueInput.SetSuggestions([]string{})
		}

		m.ValueInput, cmd = m.ValueInput.Update(msg)
		cmds = append(cmds, cmd)

	}

	m.Viewport, cmd = m.Viewport.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:

		//BUG: i dont know why but if i add 1 to response section when window width it wont error
		bugAddon := 0
		if msg.Width%10 == 5 {
			bugAddon = 1
		}
		// minus 1 for the tabs
		m.RequestHeight = msg.Height*90/100 - utils.BoxStyle.GetVerticalBorderSize() - 1

		m.RequestWidth = msg.Width*40/100 - utils.BoxStyle.GetHorizontalBorderSize() - bugAddon

		m.TextArea.SetWidth(m.RequestWidth)
		m.TextArea.SetHeight(m.RequestHeight)

		m.TableHeaders.SetColumns([]table.Column{
			{Title: "Key", Width: m.RequestWidth * 50 / 100},
			{Title: "Value", Width: m.RequestWidth * 50 / 100},
		})
		m.TableHeaders.SetHeight(m.RequestHeight - utils.BoxStyle.GetVerticalBorderSize() - 1)

		m.ValueInput.Width = m.RequestWidth*50/100 - utils.BoxStyle.GetHorizontalBorderSize() - 1

		m.KeyInput.Width = m.RequestWidth*50/100 - utils.BoxStyle.GetHorizontalBorderSize() - 1
		if !m.ViewportReady {
			m.Viewport = viewport.New(m.RequestWidth, m.RequestHeight)
			m.ViewportReady = true
			m.Viewport.SetContent(m.TextArea.View())
		} else {
			m.Viewport.Width = m.RequestWidth
			m.Viewport.Height = m.RequestHeight
		}
		switch m.FocusedTab {
		case Body:
			m.Viewport.Height = m.RequestHeight
		case Headers:
			// minus border vertical + 1 for input header
			m.Viewport.Height = m.RequestHeight - utils.BoxStyle.GetVerticalBorderSize() - 1
		}

	case tea.KeyMsg:

		switch msg.String() {
		case "tab":
			if m.FocusedTab == Headers && m.Hovered {
				if m.ValueInput.Focused() {
					m.ValueInput.Blur()
					m.KeyInput.Focus()
				} else {
					m.ValueInput.Focus()
					m.KeyInput.Blur()
				}
			}
			return m, nil
		case "ctrl+d":
			if m.Hovered && m.FocusedTab == Headers && len(m.Headers) > 0 {

				tempHeader := make([]Header, 0, len(m.Headers))

				idxSelectedRowKey := m.TableHeaders.Cursor()
				for _, h := range m.Headers {
					selectedRowKey := m.TableHeaders.Rows()[idxSelectedRowKey][0]
					if selectedRowKey != h.Key {
						tempHeader = append(tempHeader, h)
					}
				}

				m.Headers = tempHeader
				tableHeadersValue := make([]table.Row, 0, len(m.Headers))

				for _, h := range m.Headers {
					tableHeadersValue = append(tableHeadersValue, table.Row{h.Key, h.Value})
				}
				// check if its the end of headers
				if idxSelectedRowKey == len(m.Headers) {
					m.TableHeaders.SetCursor(0)
				}

				m.TableHeaders.SetRows(tableHeadersValue)
				m.KeyInput.SetValue("")
				m.ValueInput.SetValue("")

				return m, nil
			}
		case "enter":
			if m.KeyInput.Focused() || m.ValueInput.Focused() {

				m.KeyInput.SetValue(strings.TrimSpace(m.KeyInput.Value()))

				if len(m.KeyInput.Value()) > 0 &&
					len(m.ValueInput.Value()) > 0 {

					// make sure that keys are not doubled
					tempHeader := make([]Header, 0, len(m.Headers))
					var isDoubled bool

					for _, h := range m.Headers {
						if m.KeyInput.Value() == h.Key {
							tempHeader = append(tempHeader, Header{Key: h.Key, Value: m.ValueInput.Value()})
							isDoubled = true
						} else {
							tempHeader = append(tempHeader, h)
						}
					}

					if isDoubled {
						m.Headers = tempHeader

						tableHeadersValue := make([]table.Row, 0, len(m.Headers))

						for _, h := range m.Headers {
							tableHeadersValue = append(tableHeadersValue, table.Row{h.Key, h.Value})
						}
						m.TableHeaders.SetRows(tableHeadersValue)
						m.KeyInput.SetValue("")
						m.ValueInput.SetValue("")
						return m, nil
					}

					m.Headers = append(m.Headers, Header{Value: m.ValueInput.Value(), Key: m.KeyInput.Value()})

					tableHeadersValue := make([]table.Row, 0, len(m.Headers))

					for _, h := range m.Headers {
						tableHeadersValue = append(tableHeadersValue, table.Row{h.Key, h.Value})
					}

					m.TableHeaders.SetRows(tableHeadersValue)
				}
			}
			m.KeyInput.SetValue("")
			m.ValueInput.SetValue("")
			return m, nil
		}

	case tea.MouseMsg:
		m.Hovered = zone.Get("request").InBounds(msg)

		if m.Hovered && m.FocusedTab == Body {
			m.TextArea.Focus()
		} else {
			m.TextArea.Blur()
		}

		if m.Hovered && m.FocusedTab == Body {
			m.Viewport.SetContent(m.TextArea.View())
		} else if m.Hovered && m.FocusedTab == Headers {
			m.Viewport.SetContent(
				m.TableHeaders.View(),
			)
		}

		if msg.Action == tea.MouseActionRelease && msg.Button == tea.MouseButtonRight {
			if m.Hovered {
				temp := []Header{}
				for _, header := range m.Headers {
					if !zone.Get(header.Key).InBounds(msg) {
						temp = append(temp, header)
					}
				}
				m.Headers = temp
			}
		}
		if msg.Action == tea.MouseActionRelease && msg.Button == tea.MouseButtonLeft {

			if zone.Get("keyInputHeader").InBounds(msg) {
				m.KeyInput.Focus()
			} else {
				m.KeyInput.Blur()
			}

			if zone.Get("valueInputHeader").InBounds(msg) {
				m.ValueInput.Focus()
			} else {
				m.ValueInput.Blur()
			}

			if zone.Get("requestBody").InBounds(msg) {
				m.FocusedTab = Body
				m.Hovered = true
				m.TextArea.Focus()
				m.Viewport.SetContent(m.TextArea.View())
				m.Viewport.Height = m.RequestHeight
			} else if zone.Get("requestHeaders").InBounds(msg) {
				m.FocusedTab = Headers
				m.Hovered = true
				m.Viewport.Height = m.RequestHeight - utils.BoxStyle.GetVerticalBorderSize() - 1
				m.Viewport.SetContent(
					m.TableHeaders.View(),
				)

			}
		}
	}

	return m, tea.Batch(cmds...)
}
