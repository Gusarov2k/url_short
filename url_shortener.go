package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"url_short/connectdb"
)

func main() {
	var db db.Client
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.POST("/link", func(c *gin.Context) {

			c.JSON(200, gin.H{"data": "link"})
		})

		r.Run(":3002")
	}
}
