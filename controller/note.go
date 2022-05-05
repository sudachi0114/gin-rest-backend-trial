package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sudachi0114/gin-rest-backend-trial/service"
)

func NoteTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func NoteList(c *gin.Context) {
	noteService := service.NoteService{}
	NotesList := noteService.GetNoteList()
	c.JSON(http.StatusOK, gin.H{
		"data": NotesList,
	})
}
