package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sudachi0114/gin-rest-backend-trial/middleware"
)

func main() {
	engine := gin.Default()

	engine.Use(middleware.RecordUaAndTime)

	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	engine.Run(":3000")
}
