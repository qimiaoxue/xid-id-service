package main

import (
	"github.com/gin-gonic/gin"
	"github.com/qimiaoxue/xid-id-service/uuid"
	"github.com/qimiaoxue/xid-id-service/xid"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		guid := xid.New()
		uid := uuid.New()
		c.JSON(200, gin.H{
			"xid_message":  guid.String(),
			"Machine":      guid.Machine(),
			"Pid":          guid.Pid(),
			"Time":         guid.Time(),
			"Counter":      guid.Counter(),
			"uuid_message": uid,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
