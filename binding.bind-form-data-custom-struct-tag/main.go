package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const (
	customerTag   = "url"
	defaultMemory = 32 << 20
)

type customerBinding struct{}

func (customerBinding) Name() string {
	return "form"
}

func (customerBinding) Bind(req *http.Request, obj any) error {
	if err := req.ParseForm(); err != nil {
		return err
	}
	if err := req.ParseMultipartForm(defaultMemory); err != nil {
		if err != http.ErrNotMultipart {
			return err
		}
	}
	if err := binding.MapFormWithTag(obj, req.Form, customerTag); err != nil {
		return err
	}
	return validate(obj)
}

func validate(obj any) error {
	if binding.Validator == nil {
		return nil
	}
	return binding.Validator.ValidateStruct(obj)
}

// FormA is an external type that we can't modify its tag
type FormA struct {
	FieldA string `url:"field_a"`
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/list", func(c *gin.Context) {
		var urlBinding = customerBinding{}
		var opt FormA
		if err := c.MustBindWith(&opt, urlBinding); err != nil {
			return
		}
		c.JSON(http.StatusOK, gin.H{"field_a": opt.FieldA})
	})

	return router
}

func main() {
	router := gin.Default()
	router.Run(":8080")
}
