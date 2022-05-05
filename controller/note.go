package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sudachi0114/gin-rest-backend-trial/model"
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
func NoteAdd(c *gin.Context) {
	note := model.Note{}
	err := c.Bind(&note)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}

	noteService := service.NoteService{}
	err = noteService.SetNote(&note)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
