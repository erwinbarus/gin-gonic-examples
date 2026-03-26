package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	tests := []struct {
		method string
		path string
		wantedStatus int
		wantedBody string
	} {
		{ http.MethodGet, "/old", http.StatusMovedPermanently, "Moved Permanently" },
		{ http.MethodPost, "/submit", http.StatusFound, "" },
		{ http.MethodGet, "/test", http.StatusOK, "hello" },
		{ http.MethodGet, "/final", http.StatusOK, "hello" },
		{ http.MethodGet, "/result", http.StatusOK, "Redirected here!" },
	}

	for _, tt := range tests {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(tt.method, tt.path, nil)
		router.ServeHTTP(w, req)
		
		assert.Equal(t, tt.wantedStatus, w.Code)
		assert.Contains(t, w.Body.String(), tt.wantedBody)
	}
}