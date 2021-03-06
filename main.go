package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/sudachi0114/gin-rest-backend-trial/controller"
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

	noteEngine := engine.Group("/note")
	{
		v1 := noteEngine.Group("/v1")
		{
			v1.GET("/test", controller.NoteTest)
			v1.GET("/list", controller.NoteList)
			v1.POST("/add", controller.NoteAdd)
		}
	}

	engine.Run(":3000")
}
