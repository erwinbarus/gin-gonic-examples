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
			"Valid UUID -- binding succeeds",
			"/thinkerou/987fbc97-4bed-5078-9f07-9141ba07c9f3",
			http.StatusOK,
			"987fbc97-4bed-5078-9f07-9141ba07c9f3",
		},
		{
			"Invalid UUID -- binding fails with validation error",
			"/thinkerou/not-uuid",
			http.StatusBadRequest,
			"error",
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
