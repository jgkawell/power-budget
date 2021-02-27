package handler

import (
	"github.com/gin-gonic/gin"
)

// CreateRestHandler returns a new gin rest handler
func CreateRestHandler() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
