package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetWithQueryStringParameters(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/testing?name=appleboy&address=xyz&birthday=1992-03-15", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "appleboy")
}

func TestPostWithFormData(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	form := url.Values{}
	form.Add("name", "appleboy")
	form.Add("address", "xyz")
	form.Add("birthday", "1992-03-15")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/testing", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "appleboy")
}

func TestPostWithJsonBody(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	body := `{"name":"appleboy","address":"xyz","birthday":"1992-03-15T00:00:00Z"}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/testing", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "appleboy")
}

