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
			"HTML rendering 1 ",
			"/index",
			"Main Website",
		},
		{
			"HTML rendering 2",
			"/posts/index",
			"Using posts/index.tmpl",
		},
			{
			"HTML rendering 3",
			"/users/index",
			"Using users/index.tmpl",
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