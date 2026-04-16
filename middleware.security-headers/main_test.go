package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
    gin.SetMode(gin.TestMode)

    tests := []struct {
        name string
        path string
        wantHeader string
    } {
        {
            "Security Headers Custom 1",
            "/ping",
            "Content-Security-Policy",
        },
        {
            "Security Headers Custom 2",
            "/ping",
            "Referrer-Policy",
        },
        {
            "Security Headers Custom 3",
            "/ping",
            "Permissions-Policy",
        },
        {
            "Security Headers from ginhelmet",
            "/ping",
            "Strict-Transport-Security",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            router := setupRouter()
            w := httptest.NewRecorder()
            req, _ := http.NewRequest(http.MethodGet, tt.path, nil)
            router.ServeHTTP(w, req)

            assert.Contains(t, w.Header(), tt.wantHeader)
        })
    }
}