package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/arfadmuzali/restui/internal/help"
	"github.com/arfadmuzali/restui/internal/hint"
	"github.com/arfadmuzali/restui/internal/method"
	"github.com/arfadmuzali/restui/internal/request"
	"github.com/arfadmuzali/restui/internal/response"
	"github.com/arfadmuzali/restui/internal/url"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	spinner       spinner.Model
	UrlModel      url.UrlModel
	HintModel     hint.HintModel
	MethodModel   method.MethodModel
	ResponseModel response.ResponseModel
	RequestModel  request.RequestModel
	HelpModel     help.HelpModel
}

type MainModel struct {
	WindowWidth  int
	WindowHeight int

	Model
	ActiveBufferId   string
	Buffers          []Buffer
	IndexBuffers     map[string]int
	BufferModalModel BufferModalModel
}

func InitModel() MainModel {
	buffer := CreateNewBuffer()
	buffer.Id = "first"

	bufferModalModel := BufferModalModel{
		OverlayActive: false,
	}

	mainModel := MainModel{
		ActiveBufferId: buffer.Id,
		Model:          buffer.Model,
		Buffers: []Buffer{
			buffer,
		},
		IndexBuffers: map[string]int{
			buffer.Id: 0,
		},
		BufferModalModel: bufferModalModel,
	}

	return mainModel
}

func (m MainModel) BlurAll() MainModel {
	m.UrlModel.UrlInput.Blur()
	m.RequestModel.TextArea.Blur()
	m.RequestModel.ValueInput.Blur()
	m.RequestModel.KeyInput.Blur()
	m.RequestModel.Hovered = false
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
		Timeout:   60 * time.Second,
		Transport: &http.Transport{},
	}

	headers := map[string]string{}

	for _, header := range m.RequestModel.Headers {
		headers[header.Key] = header.Value
	}

	url := strings.TrimSpace(m.UrlModel.UrlInput.Value())

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

	if testerror != nil &&
		headers["Content-Type"] == "application/json" &&
		m.MethodModel.ActiveState != method.GET {
		return response.ResultMsg(response.ResultMsg{Data: nil, Error: fmt.Errorf("Something wrong with your request body\n %s", testerror.Error()), Headers: responseHeader, StatusCode: 400})
	}

	req, err := http.NewRequest(m.MethodModel.ActiveState.String(), url, requestBody)
	if err != nil {
		return response.ResultMsg(response.ResultMsg{Data: nil, Error: err, Headers: responseHeader, StatusCode: 404})
	}

	headers["Host"] = req.Host

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return response.ResultMsg(response.ResultMsg{Data: nil, Error: err, Headers: responseHeader, StatusCode: 0})
	}
	responseHeader = resp.Header

	result, err := io.ReadAll(resp.Body)

	if err == io.EOF || err == io.ErrUnexpectedEOF {
		return response.ResultMsg(response.ResultMsg{Data: nil, Error: err, Headers: responseHeader, StatusCode: resp.StatusCode})
	}

	if err != nil {
		return response.ResultMsg(response.ResultMsg{Data: nil, Error: err, Headers: responseHeader, StatusCode: resp.StatusCode})
	}

	return response.ResultMsg(response.ResultMsg{Data: result, Error: nil, Headers: responseHeader, StatusCode: resp.StatusCode})
}
