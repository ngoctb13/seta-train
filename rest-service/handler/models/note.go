package models

import useCaseModel "github.com/ngoctb13/seta-train/rest-service/internal/domains/models"

type Note struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type CreateNoteRequest struct {
	Notes []Note `json:"notes"`
}

type UpdateNoteRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func ToNoteUseCaseModel(notes []Note) []useCaseModel.Note {
	useCaseNotes := make([]useCaseModel.Note, len(notes))
	for _, note := range notes {
		useCaseNotes = append(useCaseNotes, useCaseModel.Note{
			Title: note.Title,
			Body:  note.Body,
		})
	}
	return useCaseNotes
}
