package main

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const CONTENT_TYPE = "Content-Type"

func TestMultipartForm(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()
	w := httptest.NewRecorder()

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	_ = writer.WriteField("user", "user")
	_ = writer.WriteField("password", "password")

	writer.Close()

	req, _ := http.NewRequest(http.MethodPost, "/login", &body)
	req.Header.Set(CONTENT_TYPE, writer.FormDataContentType())
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "you are logged in", w.Body.String())
}

func TestUrlEncodedForm(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()
	w := httptest.NewRecorder()

	data := url.Values{}

	data.Add("user", "user")
	data.Add("password", "password")

	req, _ := http.NewRequest(http.MethodPost, "/login", strings.NewReader(data.Encode()))
	req.Header.Set(CONTENT_TYPE, "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "you are logged in", w.Body.String())
}

func TestWrongCredential(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()
	w := httptest.NewRecorder()

	data := url.Values{}

	data.Add("user", "incorrect")
	data.Add("password", "incorrect")

	req, _ := http.NewRequest(http.MethodPost, "/login", strings.NewReader(data.Encode()))
	req.Header.Set(CONTENT_TYPE, "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, "unauthorized", w.Body.String())
}

func TestMissingRequiredField(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()
	w := httptest.NewRecorder()

	data := url.Values{}

	data.Add("user", "user")

	req, _ := http.NewRequest(http.MethodPost, "/login", strings.NewReader(data.Encode()))
	req.Header.Set(CONTENT_TYPE, "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "error")
}
