package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		headers    testHeader
		wantStatus int
	}{
		{
			"Pass custom headers",
			testHeader{300, "music"},
			http.StatusOK,
		},
		{
			"Missing headers -- zero values are used",
			testHeader{0, ""},
			http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			router := setupRouter()
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/", nil)
			if tt.headers.Rate != 0 {
				req.Header.Add("Rate", strconv.Itoa(tt.headers.Rate))
			}
			if tt.headers.Domain != "" {
				req.Header.Add("Domain", tt.headers.Domain)
			}
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}
