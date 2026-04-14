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
		wantBody string
	} {
		{
			"Standard JSON -- escapes HTML characters",
			"/json",
			`\u003cb\u003eHello, world!\u003c/b\u003e`,
		},
		{
			"PureJSON -- serves literal characters",
			"/pure-json",
			"<b>Hello, world!</b>",
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