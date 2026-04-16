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
        username string
        password string
        wantStatus int
        wantBody string
    } {
        {
            "Successful authentication -- with email",
            "foo",
            "bar",
            http.StatusOK,
            "foo@bar.com",
        },
        {
            "Successful authentication -- no email",
            "manu",
            "4321",
            http.StatusOK,
            "NO SECRET",
        },
        {
            "Wrong password -- returns 401 Unauthorized",
            "foo",
            "wrongpassword",
            http.StatusUnauthorized,
            "",
        },
        {
            "No credentials -- returns 401 Unauthorized",
            "",
            "",
            http.StatusUnauthorized,
            "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            router := setupRouter()
            w := httptest.NewRecorder()
            req, _ := http.NewRequest(http.MethodGet, "/admin/secrets", nil)
            if tt.username != "" && tt.password != "" {
                req.SetBasicAuth(tt.username, tt.password)
            }
            router.ServeHTTP(w, req)

            assert.Equal(t, tt.wantStatus, w.Code)
            assert.Contains(t, w.Body.String(), tt.wantBody)
        })
    }
}