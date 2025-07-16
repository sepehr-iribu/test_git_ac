package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSendMessageHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/send", strings.NewReader("Test message"))
	w := httptest.NewRecorder()

	http.HandlerFunc(sendMessageHandler).ServeHTTP(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	expected := "Message received: Test message"
	if string(body) != expected {
		t.Errorf("Expected %q, got %q", expected, string(body))
	}
}
