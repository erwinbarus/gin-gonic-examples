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
		name       string
		path       string
		wantStatus int
		wantBody   string
	}{
		{
			"Valid bind tags by csv: comma-separated values",
			"/search?tags=go,web,api",
			http.StatusOK,
			"go",
		},
		{
			"Valid bind labels by multi (default): repeated keys or comma-separated values",
			"/search?labels=bug&labels=helpwanted",
			http.StatusOK,
			"helpwanted",
		},
		{
			"Valid bind ids by ssv: space-separated values",
			"/search?ids_ssv=1 2 3",
			http.StatusOK,
			"1",
		},
		{
			"Valid bind levels by pipes: pipe-separated values",
			"/search?levels=1|2|3",
			http.StatusOK,
			"3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tt.path, nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.wantBody)
		})
	}
}
