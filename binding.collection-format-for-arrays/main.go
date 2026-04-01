package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Filters struct {
	Tags   []string `form:"tags" collection_format:"csv"`     // /search?tags=go,web,api
	Labels []string `form:"labels" collection_format:"multi"` // /search?labels=bug&labels=helpwanted
	IdsSSV []int    `form:"ids_ssv" collection_format:"ssv"`  // /search?ids_ssv=1 2 3
	IdsTSV []int    `form:"ids_tsv" collection_format:"tsv"`  // /search?ids_tsv=1\t2\t3
	Levels []int    `form:"levels" collection_format:"pipes"` // /search?levels=1|2|3
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/search", func(c *gin.Context) {
		var f Filters
		if err := c.ShouldBind(&f); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, f)
	})
	return router
}

func main() {
	router := setupRouter()
	router.Run(":3000")
}
