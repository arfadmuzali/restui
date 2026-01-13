package response

import (
	"encoding/json"
	"net/http"
	"sort"
	"strings"

	"github.com/arfadmuzali/restui/internal/utils"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	zone "github.com/lrstanley/bubblezone"
	"github.com/muesli/reflow/wrap"
)

func (m ResponseModel) Init() tea.Cmd {
	return nil
}

func (m ResponseModel) Update(msg tea.Msg) (ResponseModel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// minus 1 for text header and for RequestTime
		m.ResponseHeight = msg.Height*90/100 - utils.BoxStyle.GetVerticalBorderSize() - 1 - 1

		var addon int
		if msg.Width%10 != 0 {
			addon = 1
		}

		m.ResponseWidth = msg.Width*60/100 - utils.BoxStyle.GetHorizontalBorderSize() + addon
		if !m.ViewportReady {
			m.Viewport = viewport.New(m.ResponseWidth, m.ResponseHeight)
			m.ViewportReady = true
		} else {
			m.Viewport.Width = m.ResponseWidth
			m.Viewport.Height = m.ResponseHeight
		}

	case ResultMsg:
		m.IsLoading = false

		m.Result = msg

		var s string

		if m.Result.Error != nil {
			s = m.Result.Error.Error()
		} else {
			contentType := m.Result.Headers.Get("Content-Type")

			if contentType == "" {
				contentType = http.DetectContentType(m.Result.Data)
			}

			if strings.HasPrefix(contentType, "application/json") {
				var temp any
				err := json.Unmarshal(msg.Data, &temp)
				if err != nil {
					s = m.Result.Error.Error()
					m.Result.Body = wrap.String(s, m.ResponseWidth)
					m.Viewport.SetContent(m.Result.Body)
					return m, nil
				}

				body, err := utils.Formatter.Marshal(temp)
				if err != nil {
					s = m.Result.Error.Error()
					m.Result.Body = wrap.String(s, m.ResponseWidth)
					m.Viewport.SetContent(m.Result.Body)
					return m, nil
				}

				s = string(body)
			} else {
				s = string(m.Result.Data)
			}

		}
		m.ResponseTime = msg.ResponseTime + " ms"
		m.FocusedTab = Body
		m.Result.Body = wrap.String(s, m.ResponseWidth)
		m.Viewport.SetContent(m.Result.Body)
		return m, nil
	case tea.MouseMsg:
		m.Hovered = zone.Get("response").InBounds(msg)
		if msg.Action == tea.MouseActionRelease && msg.Button == tea.MouseButtonLeft {
			if zone.Get("responseBody").InBounds(msg) {
				m.FocusedTab = Body
				m.Hovered = true
				m.Viewport.SetContent(m.Result.Body)
			} else if zone.Get("responseHeaders").InBounds(msg) {
				m.FocusedTab = Headers
				m.Hovered = true

				t := table.NewWriter()
				t.AppendHeader(table.Row{
					lipgloss.NewStyle().Width(m.ResponseWidth * 40 / 100).Render("Key"),
					lipgloss.NewStyle().Width(m.ResponseWidth * 60 / 100).Render("Value"),
				})
				t.Style().Size.WidthMin = m.ResponseWidth
				t.Style().Box.UnfinishedRow = ""
				t.Style().Color.RowAlternate = text.Colors{text.BgBlack}
				t.Style().Box = table.BoxStyle{
					MiddleHorizontal: "─",
					PaddingLeft:      " ",
					PaddingRight:     " ",
				}
				t.Style().Options = table.Options{
					DrawBorder:      false,
					SeparateColumns: false,
					SeparateHeader:  true,
					SeparateRows:    false,
				}
				headers := make([]string, 0)

				for key := range m.Result.Headers {
					headers = append(headers, key)
				}
				sort.Strings(headers)
				for _, key := range headers {
					t.AppendRow(table.Row{wrap.String(key, m.ResponseWidth*40/100), wrap.String(m.Result.Headers.Get(key), m.ResponseWidth*60/100-2)})
				}
				m.Viewport.SetContent(t.Render())
			} else if zone.Get("responseCookies").InBounds(msg) {
				m.Viewport.SetContent("Cookies section is comming soon")
				m.FocusedTab = Cookies
				m.Hovered = true
			}
		}
	}
	if m.Hovered {
		m.Viewport, cmd = m.Viewport.Update(msg)
	}

	return m, cmd
}
