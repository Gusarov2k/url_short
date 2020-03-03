package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"url_short/connectdb"
	"url_short/repository"
)

func main() {
	var db db.Client
	err := db.Open("host=localhost port=5432 user=ivan dbname=short_link_development password=1234 sslmode=disable")
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}

	log.Println()

	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.POST("/link", func(c *gin.Context) {

			c.JSON(200, gin.H{"data": "link"})
		})

		r.Run(":3002")
	}
}
