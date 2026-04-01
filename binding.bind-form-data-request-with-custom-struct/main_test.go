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
		name       string
		path       string
		wantStatus int
		wantBody   string
	}{
		{
			"Nested struct -- fields from StructA are bound alongside StructB's own fields",
			"/getb?field_a=hello&field_b=world",
			200,
			"FieldA",
		},
		{
			"Nested struct pointer -- works the same way, Gin allocates the pointer automatically",
			"/getc?field_a=hello&field_c=world",
			200,
			"FieldA",
		},
		{
			"Nested struct pointer -- works the same way, Gin allocates the pointer automatically",
			"/getd?field_x=hello&field_d=world",
			200,
			"FieldX",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := setupRouter()
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tt.path, nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.wantBody)
		})
	}
}
