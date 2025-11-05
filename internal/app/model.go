package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/arfadmuzali/restui/internal/hint"
	"github.com/arfadmuzali/restui/internal/method"
	"github.com/arfadmuzali/restui/internal/request"
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
	RequestModel  request.RequestModel
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
		RequestModel:  request.New(),
		spinner:       s,
	}

	return model
}

func (m MainModel) BlurAllInput(exeptions ...string) MainModel {
	exs := make(map[string]bool, len(exeptions))
	for _, value := range exeptions {
		exs[value] = true
	}
	if !exs["url"] {
		m.UrlModel.UrlInput.Blur()
	} else if !exs["requestBody"] {
		m.RequestModel.TextArea.Blur()
	}
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

	responseHeader := http.Header{}

	requestBody := bytes.NewReader([]byte(m.RequestModel.TextArea.Value()))

	if _, isBodyexists := headers["Content-Type"]; requestBody.Len() > 0 && !isBodyexists {
		headers["Content-Type"] = "application/json"
	}

	var js any
	testerror := json.Unmarshal([]byte(m.RequestModel.TextArea.Value()), &js)

	// if !json.Valid([]byte(m.RequestModel.TextArea.Value())) &&
	if testerror != nil &&
		headers["Content-Type"] == "application/json" &&
		m.MethodModel.ActiveState != method.GET {
		return response.ResultMsg(response.ResultMsg{Data: nil, Error: fmt.Errorf("Something wrong with your request body\n%s", testerror.Error()), Headers: responseHeader, StatusCode: 400})
	}

	req, err := http.NewRequest(m.MethodModel.ActiveState.String(), url, requestBody)
	if err != nil {
		return response.ResultMsg(response.ResultMsg{Data: nil, Error: err, Headers: responseHeader, StatusCode: 404})
	}

	headers["Host"] = req.Host

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// TODO: end dummy headers

	resp, err := client.Do(req)
	if err != nil {
		return response.ResultMsg(response.ResultMsg{Data: nil, Error: err, Headers: responseHeader, StatusCode: 0})
	}
	responseHeader = resp.Header

	if resp.StatusCode >= 400 {
		return response.ResultMsg(response.ResultMsg{Data: nil, Error: fmt.Errorf("Unexpected status code: %v", resp.StatusCode), Headers: responseHeader, StatusCode: resp.StatusCode})
	}

	result, err := io.ReadAll(resp.Body)

	if err != nil {
		return response.ResultMsg(response.ResultMsg{Data: nil, Error: err, Headers: responseHeader, StatusCode: resp.StatusCode})
	}

	return response.ResultMsg(response.ResultMsg{Data: result, Error: nil, Headers: responseHeader, StatusCode: resp.StatusCode})
}
