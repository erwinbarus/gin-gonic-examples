package main

import (
	"testing"
	"net/http"
	"net/url"
	"net/http/httptest"
	"strings"
	"bytes"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const PATH = "/form-post"
const CONTENT_TYPE = "Content-Type"

func TestUrlEncodedForm(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	data := url.Values{}
	data.Add("message", "hello")
	data.Add("nick", "world")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, PATH, strings.NewReader(data.Encode()))
	req.Header.Set(CONTENT_TYPE, "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message":"hello","nick":"world","status":"posted"}`, w.Body.String())
}

func TestUrlEncodedFormWithMissingNick(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	data := url.Values{}
	data.Add("message", "hello")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, PATH, strings.NewReader(data.Encode()))
	req.Header.Set(CONTENT_TYPE, "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message":"hello","nick":"anonymous","status":"posted"}`, w.Body.String())
}

func TestMultipartForm(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	_ = writer.WriteField("message", "hello")
	_ = writer.WriteField("nick", "world")

	writer.Close()

		w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, PATH, &body)
	req.Header.Set(CONTENT_TYPE, writer.FormDataContentType())
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message":"hello","nick":"world","status":"posted"}`, w.Body.String())
}