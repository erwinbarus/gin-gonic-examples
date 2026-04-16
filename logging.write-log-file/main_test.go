package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var LOG_FILE = "gin.log"

func TestSetupRouterCreateLogFile(t *testing.T) {
    gin.SetMode(gin.TestMode)

    router, f := setupRouter()

    req, _ := http.NewRequest("GET", "/ping", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    f.Close()

    if _, err := os.Stat(LOG_FILE); os.IsNotExist(err) {
        t.Fatalf("expected gin.log to be created")
    }

    _ = os.Remove(LOG_FILE)
}