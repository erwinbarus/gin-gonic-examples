package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSessionFlow(t *testing.T) {
    gin.SetMode(gin.TestMode)

    router := setupRouter()

    // 1. Login request
    w1 := httptest.NewRecorder()
    req1, _ := http.NewRequest(http.MethodGet, "/login", nil)
    router.ServeHTTP(w1, req1)

    assert.Equal(t, http.StatusOK, w1.Code)
    assert.Contains(t, w1.Body.String(), "logged in")

    // Extract cookie
    cookies := w1.Result().Cookies()

    // 2. Profile request with cookie
    w2 := httptest.NewRecorder()
    req2, _ := http.NewRequest(http.MethodGet, "/profile", nil)

    for _, c := range cookies {
        req2.AddCookie(c)
    }

    router.ServeHTTP(w2, req2)

    assert.Equal(t, http.StatusOK, w2.Code)
    assert.Contains(t, w2.Body.String(), "john")

    // 3. Logout request
    w3 := httptest.NewRecorder()
    req3, _ := http.NewRequest(http.MethodGet, "/logout", nil)
    router.ServeHTTP(w3, req3)

    assert.Equal(t, http.StatusOK, w3.Code)
    assert.Contains(t, w3.Body.String(), "logged out")
}