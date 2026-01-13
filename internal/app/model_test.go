package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arfadmuzali/restui/internal/method"
	"github.com/arfadmuzali/restui/internal/response"
)

// Test that HandleHttpRequest returns an error ResultMsg when JSON body is invalid
func TestHandleHttpRequest_InvalidJSONBody(t *testing.T) {
	m := InitModel()

	// set method to POST so body validation runs
	m.MethodModel.ActiveState = method.POST

	// set a body that's invalid JSON
	m.RequestModel.TextArea.SetValue("{invalid json}")

	// add Content-Type header so validation is applied
	m.RequestModel.Headers = append(m.RequestModel.Headers, struct{ Key, Value string }{Key: "Content-Type", Value: "application/json"})

	msg := m.HandleHttpRequest()
	res, ok := msg.(response.ResultMsg)
	if !ok {
		t.Fatalf("expected response.ResultMsg, got %T", msg)
	}

	if res.Error == nil {
		t.Fatalf("expected an error for invalid JSON body, got nil")
	}
	if res.StatusCode != 400 {
		t.Fatalf("expected status code 400 for invalid JSON body, got %d", res.StatusCode)
	}
}

// Test a successful GET against an httptest server
func TestHandleHttpRequest_SuccessGET(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("hello"))
	}))
	defer ts.Close()

	m := InitModel()
	m.UrlModel.UrlInput.SetValue(ts.URL)
	m.MethodModel.ActiveState = method.GET
	// ensure no body
	m.RequestModel.TextArea.SetValue("")

	msg := m.HandleHttpRequest()
	res, ok := msg.(response.ResultMsg)
	if !ok {
		t.Fatalf("expected response.ResultMsg, got %T", msg)
	}

	if res.Error != nil {
		t.Fatalf("unexpected error from request: %v", res.Error)
	}
	if res.StatusCode != 200 {
		t.Fatalf("expected 200 OK, got %d", res.StatusCode)
	}
	if string(res.Data) != "hello" {
		t.Fatalf("unexpected response body: %q", string(res.Data))
	}
}
