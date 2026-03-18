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
	router := gin.Default()

	router.GET("/welcome", getWithQueryString)

	tests := []struct {
		path 		string
		wantBody	string
	} {
		{ "/welcome?first=John&last=Doe", "Hello John Doe" },
		{ "/welcome?first=&last=", "Hello  " },
		{ "/welcome?first=John", "Hello John Guest" },
	}

	for _, tt := range tests {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", tt.path, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, tt.wantBody, w.Body.String())
	}
}