package main

import (
	"net/http/httptest"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	tests := []struct {
		path		string
		wantStatus	int
		wantedBody	string
	} {
		{ "/public/health", 200, "Healthy" },
		{ "/private/profile", 200, "View profile here!" },
		{ "/private/setting", 200, "View setting here!" },
	}

	for _, tt := range tests {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, tt.path, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, tt.wantStatus, w.Code)
		assert.Equal(t, tt.wantedBody, w.Body.String())
	}
}