package main

import (
	"bytes"
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
		httpMethod	   string
		path	   string
		body       string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "GET with query parameters",
			httpMethod:	http.MethodGet,
			path:       "/testing?name=appleboy&address=xyz",
			body:       "",
			wantStatus: http.StatusOK,
			wantBody:   `{"address":"xyz","name":"appleboy"}`,
		},
		{
			name:       "POST with query parameters -- body is ignored, only query is bound",
			httpMethod: http.MethodPost,
			path:       "/testing?name=appleboy&address=xyz",
			body:       `{"name":"ignored","address":"ignored"}`,
			wantStatus: http.StatusOK,
			wantBody:   `{"address":"xyz","name":"appleboy"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(tt.httpMethod, tt.path, bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			assert.Equal(t, tt.wantBody, w.Body.String())
		})
	}
}