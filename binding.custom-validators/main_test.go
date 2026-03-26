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
	router := setupRouter()

	tests := []struct {
		name		string
		path       	string
		wantStatus 	int
		wantBody   	string
	} {
		{
			name:       "valid date",
			path:		"/bookable?check_in=2118-04-16&check_out=2118-04-17",
			wantStatus: http.StatusOK,
			wantBody:   "Booking dates are valid!",
		},
		{
			name:       "valid date",
			path:		"/bookable?check_in=2118-04-17&check_out=2118-04-16",
			wantStatus: http.StatusBadRequest,
			wantBody:   "Error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tt.path, nil)
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.wantBody)
		})
	}
}