package main

import (
	"id-service/xid"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		guid := xid.New()
		c.JSON(200, gin.H{
			"message": guid.String(),
			"Machine": guid.Machine(),
			"Pid":     guid.Pid(),
			"Time":    guid.Time(),
			"Counter": guid.Counter(),
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
