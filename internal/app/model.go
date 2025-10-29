package app

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/arfadmuzali/restui/internal/hint"
	"github.com/arfadmuzali/restui/internal/method"
	"github.com/arfadmuzali/restui/internal/response"
	"github.com/arfadmuzali/restui/internal/url"
	"github.com/arfadmuzali/restui/internal/utils"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MainModel struct {
	WindowWidth  int
	WindowHeight int
	spinner      spinner.Model

	UrlModel      url.UrlModel
	HintModel     hint.HintModel
	MethodModel   method.MethodModel
	ResponseModel response.ResponseModel
}

func InitModel() MainModel {

	s := spinner.New()
	s.Spinner = spinner.Points
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color(utils.BlueColor))

	model := MainModel{
		UrlModel:      url.New(),
		HintModel:     hint.New(),
		MethodModel:   method.New(),
		ResponseModel: response.New(),
		spinner:       s,
	}

	return model
}

func (m MainModel) BlurAllInput() MainModel {
	m.UrlModel.UrlInput.Blur()
	return m
}

func (m MainModel) StartRequest() (MainModel, tea.Cmd) {

	if m.UrlModel.UrlInput.Value() == "" {
		return m, nil
	}
	m.ResponseModel.IsLoading = true

	return m, tea.Batch(func() tea.Msg {
		return response.IsLoadingMsg(true)
	}, m.spinner.Tick)
}

func (m MainModel) HandleHttpRequest() tea.Msg {
	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	// TODO: start dummy headers
	headers := map[string]string{"User-Agent": "RESTUI/0.0.1", "Accept": "*/*"}

	url := m.UrlModel.UrlInput.Value()

	if url == "" {
		return nil
	}
	if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}

	req, err := http.NewRequest(m.MethodModel.ActiveState.String(), url, bytes.NewReader([]byte(`{"nama":"arfad"}`)))
	headers["Host"] = req.Host
	if m.MethodModel.ActiveState != method.GET {
		headers["Content-Type"] = "application/json"
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// TODO: end dummy headers

	if err != nil {
		return response.ResultMsg(response.ResultMsg{Data: nil, Error: err})
	}

	resp, err := client.Do(req)
	if err != nil {
		return response.ResultMsg(response.ResultMsg{Data: nil, Error: err})
	}

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return response.ResultMsg(response.ResultMsg{Data: nil, Error: err})
	}

	return response.ResultMsg(response.ResultMsg{Data: result, Error: nil})
}
