package response

import (
	"encoding/json"
	"errors"
	"math"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/arfadmuzali/restui/internal/utils"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	zone "github.com/lrstanley/bubblezone/v2"
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

		// Use rounded splits so left + right == total width
		borderH := utils.BoxStyle.GetHorizontalBorderSize()
		leftOuter := int(math.Round(float64(msg.Width)*0.4)) - borderH

		leftOuter = max(leftOuter, 1)
		responseOuter := msg.Width - leftOuter

		// inner viewport width for response excludes the border
		m.ResponseWidth = responseOuter - borderH

		m.ResponseWidth = max(m.ResponseWidth, 1)

		if !m.ViewportReady {
			m.Viewport = viewport.New(viewport.WithWidth(m.ResponseWidth), viewport.WithHeight(m.ResponseHeight))
			m.ViewportReady = true

			fullMessage := lipgloss.JoinVertical(lipgloss.Center,
				"No response yet",
				"Send request to see the response",
			)
			s := lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center).Width(m.Viewport.Width()).Height(m.Viewport.Height()).Render(fullMessage)
			m.Viewport.SetContent(s)
		} else {
			m.Viewport.SetWidth(m.ResponseWidth)
			m.Viewport.SetHeight(m.ResponseHeight)
		}

	case ResultMsg:
		m.IsLoading = false

		m.Result = msg

		var s string

		if m.Result.Error != nil {
			errUrl := &url.Error{}
			var errMessage string
			if errors.As(m.Result.Error, &errUrl) {
				errMessage = errUrl.Err.Error()
			} else {
				errMessage = m.Result.Error.Error()
			}
			textStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(utils.RedColor))
			fullMessage := lipgloss.JoinVertical(lipgloss.Center,
				textStyle.Bold(true).Render("Could not send request"),
				textStyle.Render(errMessage),
			)
			s = lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center).Width(m.Viewport.Width()).Height(m.Viewport.Height()).Render(fullMessage)

		} else {
			contentType := m.Result.Headers.Get("Content-Type")

			if contentType == "" {
				contentType = http.DetectContentType(m.Result.Data)
			}

			if strings.Contains(contentType, "application/json") {
				var temp any
				err := json.Unmarshal(m.Result.Data, &temp)
				if err != nil {
					s = string(m.Result.Data)
					m.Result.Body = wrap.String(s, m.ResponseWidth)
					m.Viewport.SetContent(m.Result.Body)
					return m, nil
				}

				body, err := utils.Formatter.Marshal(temp)
				if err != nil {
					s = err.Error()
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
		if msg.Mouse().Button == tea.MouseLeft {
			if zone.Get("responseBody").InBounds(msg) {
				m.FocusedTab = Body
				m.Hovered = true
				m.Viewport.SetContent(m.Result.Body)
			} else if zone.Get("responseHeaders").InBounds(msg) {
				m.FocusedTab = Headers
				m.Hovered = true

				t := table.NewWriter()
				// compute column widths using rounded splits of the response viewport width
				keyWidth := int(math.Round(float64(m.ResponseWidth) * 0.4))
				keyWidth = max(keyWidth, 1)
				valueWidth := m.ResponseWidth - keyWidth - 2
				valueWidth = max(valueWidth, 1)
				t.AppendHeader(table.Row{
					lipgloss.NewStyle().Width(keyWidth).Render("Key"),
					lipgloss.NewStyle().Width(valueWidth).Render("Value"),
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
					t.AppendRow(table.Row{wrap.String(key, keyWidth), wrap.String(m.Result.Headers.Get(key), valueWidth)})
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
