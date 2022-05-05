package service

import "github.com/sudachi0114/gin-rest-backend-trial/model"

type NoteService struct{}

func (NoteService) GetNoteList() []model.Note {
	notes := make([]model.Note, 0)
	err := DbEngine.Distinct("id", "title", "content").Limit(10, 0).Find(&notes)
	if err != nil {
		panic(err)
	}
	return notes
}
