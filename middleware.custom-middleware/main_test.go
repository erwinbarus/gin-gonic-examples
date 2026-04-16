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

    tests := []struct {
        name string
        path string
        wantBody string
    } {
        {
            "Custom Logger Middleware",
            "/test",
            "12345",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            router := setupRouter()
            w := httptest.NewRecorder()
            req, _ := http.NewRequest(http.MethodGet, tt.path, nil)
            router.ServeHTTP(w, req)

            assert.Equal(t, tt.wantBody, w.Body.String())
        })
    }
}