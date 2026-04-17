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
        wantStatus int
        wantBody string
    } {
        {
            "Health Check",
            200,
            "ok",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            router := setupRouter()
            w := httptest.NewRecorder()
            q, _ := http.NewRequest(http.MethodGet, "/healthz", nil)
            router.ServeHTTP(w, q)

            assert.Equal(t, tt.wantStatus, w.Code)
            assert.Contains(t, w.Body.String(), tt.wantBody)
        })
    }
}