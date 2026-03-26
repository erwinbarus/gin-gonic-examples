package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
)

func TestBindDefaultValue(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/person", strings.NewReader(""))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedJSON := `{
		"Name":"William",
		"Age":10,
		"Friends":["Will","Bill"],
		"Addresses":["foo","bar"],
		"LapTimes":[1,2,3]
	}`
	assert.JSONEq(t, expectedJSON, w.Body.String())
}