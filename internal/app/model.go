package app

import (
	"bytes"
	"context"
	"errors"
	"strconv"

	"io"
	"net/http"
	"strings"
	"time"

	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
	"github.com/arfadmuzali/restui/internal/help"
	"github.com/arfadmuzali/restui/internal/hint"
	"github.com/arfadmuzali/restui/internal/method"
	"github.com/arfadmuzali/restui/internal/request"
	"github.com/arfadmuzali/restui/internal/response"
	"github.com/arfadmuzali/restui/internal/url"
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
	CancelContext    context.Context
	CancelRequest    context.CancelFunc
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

	ctx := context.Background()

	cancelCtx, cancel := context.WithTimeout(ctx, time.Second*60)
	m.CancelContext = cancelCtx
	m.CancelRequest = cancel

	return m, tea.Batch(func() tea.Msg {
		return response.IsLoadingMsg(true)
	}, m.spinner.Tick)
}

func (m MainModel) HandleHttpRequest() tea.Msg {
	client := &http.Client{
		// move timeout to context
		// Timeout:   60 * time.Second,
		Transport: &http.Transport{},
	}

	headers := map[string]string{}

	for _, header := range m.RequestModel.Headers {
		headers[header.Key] = header.Value
	}

	// Validate Url
	url := strings.TrimSpace(m.UrlModel.UrlInput.Value())

	if url == "" {
		return nil
	}
	if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}

	responseHeader := http.Header{}

	requestBody := bytes.NewReader([]byte(m.RequestModel.TextArea.Value()))

	if _, isContentTypeExist := headers["Content-Type"]; requestBody.Len() > 0 && !isContentTypeExist {
		headers["Content-Type"] = "application/json"
	}

	req, err := http.NewRequestWithContext(m.CancelContext, m.MethodModel.ActiveState.String(), url, requestBody)
	defer m.CancelRequest()

	if err != nil {
		return response.ResultMsg(response.ResultMsg{Data: nil, Error: err, Headers: responseHeader, StatusCode: 0})
	}

	headers["Host"] = req.Host

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	start := time.Now()
	resp, err := client.Do(req)
	responseTime := time.Since(start)

	if err != nil {

		if errors.Is(err, context.Canceled) {
			return response.ResultMsg(response.ResultMsg{Data: nil, Error: errors.New("Request cancelled"), Headers: responseHeader, StatusCode: 0})
		} else if errors.Is(err, context.DeadlineExceeded) {
			return response.ResultMsg(response.ResultMsg{Data: nil, Error: errors.New("Request cancelled because it took so long"), Headers: responseHeader, StatusCode: 0})
		} else {
			return response.ResultMsg(response.ResultMsg{Data: nil, Error: err, Headers: responseHeader, StatusCode: 0})
		}
	}
	responseHeader = resp.Header

	result, err := io.ReadAll(resp.Body)

	if err == io.EOF || err == io.ErrUnexpectedEOF {
		return response.ResultMsg(response.ResultMsg{Data: nil, Error: err, Headers: responseHeader, StatusCode: resp.StatusCode})
	}

	if err != nil {
		return response.ResultMsg(response.ResultMsg{Data: nil, Error: err, Headers: responseHeader, StatusCode: resp.StatusCode})
	}

	return response.ResultMsg(response.ResultMsg{
		Data:         result,
		Error:        nil,
		Headers:      responseHeader,
		StatusCode:   resp.StatusCode,
		ResponseTime: strconv.Itoa(int(responseTime.Abs().Milliseconds())),
	})
}
