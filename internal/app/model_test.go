package app

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/arfadmuzali/restui/internal/method"
	"github.com/arfadmuzali/restui/internal/response"
)

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
	ctx := context.Background()

	cancelContext, cancelFunc := context.WithTimeout(ctx, 60*time.Second)

	m.CancelContext = cancelContext
	m.CancelRequest = cancelFunc

	msg := m.HandleHttpRequest()
	_ = msg
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
