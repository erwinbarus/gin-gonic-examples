package main

import (
	"testing"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name string
		path string
		wantContentType string
	} {
		{
			"json rendering",
			"/json",
			"application/json",
		},
		{
			"xml rendering",
			"/xml",
			"application/xml",
		},
		{
			"yaml rendering",
			"/yaml",
			"application/yaml",
		},
		{
			"toml rendering",
			"/toml",
			"application/toml",
		},
		{
			"protobuff rendering",
			"/proto",
			"application/x-protobuf",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := setupRouter()
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tt.path, nil)
			router.ServeHTTP(w, req)

			assert.Contains(t, w.HeaderMap.Get("Content-Type"), tt.wantContentType)
		})
	}
}