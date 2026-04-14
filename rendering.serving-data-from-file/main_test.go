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

	tests := []struct {
		name string
		path string
		wantBody string
	} {
		{
			"Serve a file inline (displayed in browser)",
			"/local/file",
			"local file",
		},
		{
			"Serve a file from an http.FileSystem",
			"/fs/file",
			"fs file",
		},
		{
			"Serve a file as a downloadable attachment with a custom filename",
			"/download",
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := setupRouter()
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tt.path, nil)
			router.ServeHTTP(w, req)

			assert.Contains(t, w.Body.String(), tt.wantBody)
		})
	}
}