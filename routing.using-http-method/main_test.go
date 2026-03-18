package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAllMethod(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("", getting)
	router.POST("/", posting)
	router.PUT("/", putting)
	router.DELETE("/", deleting)
	router.OPTIONS("/", optioning)
	router.HEAD("/", heading)

	tests := []struct {
		method string
		wantStatus int
		wantBody string
	} {
		{ "GET", 200, `{"method":"GET"}` },
		{ "POST", 201, `{"method":"POST"}` },
		{ "PUT", 200, `{"method":"PUT"}` },
		{ "DELETE", 200, `{"method":"DELETE"}` },
		{ "OPTIONS", 200, `{"method":"OPTIONS"}` },
		{ "HEAD", 200, `{"method":"HEAD"}` },
	}

	for _, tt := range tests {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(tt.method, "/", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, tt.wantStatus, w.Code)
		assert.Equal(t, tt.wantBody, w.Body.String())
	}
}