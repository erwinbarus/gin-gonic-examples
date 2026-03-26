package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
	"testing"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSingleUpload(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := SetupRouter()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", "test.txt")
	assert.NoError(t, err)

	_, err = io.WriteString(part, "this is a test file")
	assert.NoError(t, err)

	writer.Close()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/upload/single", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "uploaded!")

	_, err = os.Stat("./files/test.txt")
	assert.NoError(t, err)

	_ = os.Remove("./files/test.txt")
}

func TestUploadMultipleFiles(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := SetupRouter()

	// Ensure directory exists
	err := os.MkdirAll("./files", os.ModePerm)
	assert.NoError(t, err)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create multiple files
	files := map[string]string{
		"file1.txt": "hello world",
		"file2.txt": "golang test",
	}

	for filename, content := range files {
		part, err := writer.CreateFormFile("files", filename)
		assert.NoError(t, err)

		_, err = io.WriteString(part, content)
		assert.NoError(t, err)
	}

	writer.Close()

	req, err := http.NewRequest(http.MethodPost, "/upload/multiple", body)
	assert.NoError(t, err)

	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "2 files uploaded!")

	// Check files exist
	for filename := range files {
		path := "./files/" + filename
		_, err := os.Stat(path)
		assert.NoError(t, err)

		// Cleanup
		_ = os.Remove(path)
	}

	_ = os.Remove("./files")
}

func TestUploadSingleFileTooLarge(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := SetupRouter()

	// Create a buffer and a multipart writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create a "large" file exceeding MaxUploadSize
	part, err := writer.CreateFormFile("file", "largefile.txt")
	assert.NoError(t, err)

	// Fill the file with MaxUploadSize + 1 bytes
	largeContent := strings.Repeat("a", MaxUploadSize+1)
	_, err = io.WriteString(part, largeContent)
	assert.NoError(t, err)

	writer.Close()

	// Build the request
	req, err := http.NewRequest(http.MethodPost, "/upload/single", body)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Record response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusRequestEntityTooLarge, w.Code)
	assert.Contains(t, w.Body.String(), "file too large")
}