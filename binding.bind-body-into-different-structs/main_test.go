package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestBodyMatchesFormA(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	formTest := formA{
		Foo: "hello",
	}
	fromJson, _ := json.Marshal(formTest)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/bind", strings.NewReader(string(fromJson)))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedJSON := `{
		"foo":"hello",
		"message":"matched formA"
	}`
	assert.JSONEq(t, expectedJSON, w.Body.String())
}

func TestBodyMatchesFormB(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	formTest := formB{
		Bar: "world",
	}
	fromJson, _ := json.Marshal(formTest)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/bind", strings.NewReader(string(fromJson)))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedJSON := `{
		"bar":"world",
		"message":"matched formB"
	}`
	assert.JSONEq(t, expectedJSON, w.Body.String())
}

func TestBodyDoesNotMatchAnyForm(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/bind", strings.NewReader(""))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "error")
}
