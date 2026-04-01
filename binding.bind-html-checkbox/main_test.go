package main

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const CONTENT_TYPE = "Content-Type"

func TestGetIndex(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "<p>Check some colors</p>")
}

func TestSelectAllThreeColors(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()
	w := httptest.NewRecorder()

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	_ = writer.WriteField("colors", "red")
	_ = writer.WriteField("colors", "green")
	_ = writer.WriteField("colors", "blue")

	writer.Close()

	req, _ := http.NewRequest(http.MethodPost, "/", &body)
	req.Header.Set(CONTENT_TYPE, writer.FormDataContentType())
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "red")
	assert.Contains(t, w.Body.String(), "green")
	assert.Contains(t, w.Body.String(), "blue")
}

func TestSelectOnlyOneColor(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()
	w := httptest.NewRecorder()

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	_ = writer.WriteField("colors", "red")

	writer.Close()

	req, _ := http.NewRequest(http.MethodPost, "/", &body)
	req.Header.Set(CONTENT_TYPE, writer.FormDataContentType())
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "red")
	assert.NotContains(t, w.Body.String(), "green")
	assert.NotContains(t, w.Body.String(), "blue")
}

func TestNoColorsSelected(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()
	w := httptest.NewRecorder()

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	writer.Close()

	req, _ := http.NewRequest(http.MethodPost, "/", &body)
	req.Header.Set(CONTENT_TYPE, writer.FormDataContentType())
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotContains(t, w.Body.String(), "red")
	assert.NotContains(t, w.Body.String(), "green")
	assert.NotContains(t, w.Body.String(), "blue")
}
