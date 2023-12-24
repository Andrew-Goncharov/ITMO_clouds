package schemes

import "note_storage_service/internal/models"

const (
	StatusOk = "OK"
)

type GetNoteResponse struct {
	Data   models.Note `json:"data"`
	Status string      `json:"status"`
}

type GetNotesResponse struct {
	Data   []models.Note `json:"data"`
	Status string        `json:"status"`
}

type CreateNoteResponse struct {
	Data   models.Note `json:"data"`
	Status string      `json:"status"`
}

type UpdateNoteResponse struct {
	Data   models.Note `json:"data"`
	Status string      `json:"status"`
}

type DeleteNoteResponse struct {
	Status string `json:"status"`
}
