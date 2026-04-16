package main

import (
	"io"
	"log/slog"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const REQUEST_ID_HEADER_NAME = "X-Request-Id"

func RequestIDMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        requestID := c.GetHeader(REQUEST_ID_HEADER_NAME)
        if requestID == "" {
            requestID = uuid.New().String()
        }
        c.Set("request_id", requestID)
        c.Header(REQUEST_ID_HEADER_NAME, requestID)
        c.Next()
    }
}

func SlogMiddleware(logger *slog.Logger) gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        path := c.Request.URL.Path
        query := c.Request.URL.RawQuery

        c.Next()

        logger.Info("request",
            slog.String("method", c.Request.Method),
            slog.String("path", path),
            slog.String("query", query),
            slog.Int("status", c.Writer.Status()),
            slog.Duration("latency", time.Since(start)),
            slog.String("client_ip", c.ClientIP()),
            slog.Int("body_size", c.Writer.Size()),   
        )

        if len(c.Errors) > 0 {
            for _, err := range c.Errors {
                logger.Error("request error", slog.String("error", err.Error()))
            }
        }
    }
}


func setupRouter() (*gin.Engine, *os.File) {
    gin.DisableConsoleColor()

    f, _ := os.Create("gin.log")
    multiWriter := io.MultiWriter(f, os.Stdout)

    logger := slog.New(slog.NewJSONHandler(multiWriter, nil))

    router := gin.New()
    router.Use(RequestIDMiddleware())
    router.Use(SlogMiddleware(logger))
    router.Use(gin.Recovery())
    
    router.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
    })

    return router, f
}

func main() {
    router, _ := setupRouter()
    router.Run(":3000")
}