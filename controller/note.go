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

func NoteList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ノートのリストを取得して、ここに表示します",
	})
}
