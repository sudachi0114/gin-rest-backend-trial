package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NoteTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
