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
			"Serves an entire directory",
			"/assets/style.css",
			"background-color",
		},
		{
			"Serve files from an embedded filesystem.",
			"/public/script.js",
			"console.log",
		},
		{
			"Serves a single file.",
			"/robots.txt",
			"search engine indexing whitelist",
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