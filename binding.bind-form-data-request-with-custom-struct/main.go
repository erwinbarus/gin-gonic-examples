package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StructA struct {
	FieldA string `form:"field_a"`
}

type StructB struct {
	NestedStruct StructA
	FieldB       string `form:"field_b"`
}

type StructC struct {
	NestedStructPointer *StructA
	FieldC              string `form:"field_c"`
}

type StructD struct {
	NestedAnonyStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/getb", func(c *gin.Context) {
		var b StructB
		c.Bind(&b)
		c.JSON(http.StatusOK, gin.H{
			"a": b.NestedStruct,
			"b": b.FieldB,
		})
	})

	router.GET("/getc", func(c *gin.Context) {
		var b StructC
		c.Bind(&b)
		c.JSON(http.StatusOK, gin.H{
			"a": b.NestedStructPointer,
			"c": b.FieldC,
		})
	})

	router.GET("/getd", func(c *gin.Context) {
		var b StructD
		c.Bind(&b)
		c.JSON(http.StatusOK, gin.H{
			"x": b.NestedAnonyStruct,
			"d": b.FieldD,
		})
	})

	return router
}

func main() {
	router := setupRouter()
	router.Run(":3000")
}
