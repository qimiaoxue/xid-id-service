package main

import (
	"id-service/xid"

	"github.com/gin-gonic/gin"
)

func main() {
	guid := xid.New()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": guid.String(),
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
