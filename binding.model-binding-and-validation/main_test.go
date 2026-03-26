package main

import (
	"bytes"
	"testing"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestJsonLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	tests := []struct {
		name       string
		body       string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "valid credentials",
			body:       `{"user":"manu","password":"123"}`,
			wantStatus: http.StatusOK,
			wantBody:   "you are logged in",
		},
		{
			name:       "wrong password",
			body:       `{"user":"manu","password":"wrong"}`,
			wantStatus: http.StatusUnauthorized,
			wantBody:   "unauthorized",
		},
		{
			name:       "missing fields",
			body:       `{"user":"manu"}`,
			wantStatus: http.StatusBadRequest,
			wantBody:   "missing field",
		},
		{
			name:       "invalid json",
			body:       `invalid-json`,
			wantStatus: http.StatusBadRequest,
			wantBody:   "error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/loginJSON", bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.wantBody)
		})
	}
}