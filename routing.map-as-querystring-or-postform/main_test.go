package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	form := url.Values{}
	form.Add("names[x]", "foo")
	form.Add("names[y]", "bar")

	req, err := http.NewRequest(
		http.MethodPost,
		"/post?ids[a]=123&ids[b]=456",
		bytes.NewBufferString(form.Encode()),
	)
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "123", response["ids"]["a"])
	assert.Equal(t, "456", response["ids"]["b"])

	assert.Equal(t, "foo", response["names"]["x"])
	assert.Equal(t, "bar", response["names"]["y"])
}