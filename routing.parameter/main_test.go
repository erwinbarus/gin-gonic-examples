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
	
	router.GET("/user/:name", getWithParam)
	router.GET("/user/:name/*action", getWithParams)

	tests := []struct {
		method 		string
		path 		string
		wantStatus	int
		wantBody	string
	} {
		{ "GET", "/user/John", 200, "Hello, John!" },
		{ "GET", "/user/John/send", 200, "John is /send" },
		{ "GET", "/user/John/send/", 200, "John is /send/" },
		{ "GET", "/user/John/send/other/email", 200, "John is /send/other/email" },
		{ "GET", "/user/John/", 200, "John is /" },
		{ "GET", "/", 404, "404 page not found" },
	}

	for _, tt := range tests {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(tt.method, tt.path,  nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, tt.wantStatus, w.Code)
		assert.Equal(t, tt.wantBody, w.Body.String())
	}
}