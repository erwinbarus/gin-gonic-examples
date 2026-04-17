package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestStreamEndpointCancelation(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := setupRouter()

    ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, "GET", "/api/stream", nil)
	w := httptest.NewRecorder()

	done := make(chan bool)

	go func() {
		router.ServeHTTP(w, req)
		close(done)
	}()

	// Let it run for a short time to accumulate events
	time.Sleep(2500 * time.Millisecond)

	cancel()

	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("handler did not stop after client disconnect")
	}

	if w.Body.Len() == 0 {
		t.Fatal("expected some SSE output, got empty response")
	}
}

func TestStreamSendsEvents(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := setupRouter()

	req, _ := http.NewRequest("GET", "/api/stream", nil)
	w := httptest.NewRecorder()

	go func() {
		router.ServeHTTP(w, req)
	}()

	// wait for at least one event cycle
	time.Sleep(1100 * time.Millisecond)

	body := w.Body.String()

	if body == "" {
		t.Fatal("expected SSE data but got empty body")
	}
}