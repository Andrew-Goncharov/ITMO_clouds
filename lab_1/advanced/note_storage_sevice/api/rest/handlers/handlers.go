package handlers

import (
	"errors"
	"log"
	"net/http"
	"note_storage_service/api/rest/handlers/schemes"
	"note_storage_service/internal/models"
	"note_storage_service/internal/services/notes"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (r *Resolver) GetNote(ctx *fiber.Ctx) error {
	noteID := ctx.Params("note_id", "")
	if noteID == "" {
		log.Print("node_id is empty string")
		statusCode := http.StatusBadRequest
		return ctx.Status(statusCode).JSON(schemes.NewSimpleWebError(statusCode, "node_id is empty string"))
	}

	if _, err := uuid.Parse(noteID); err != nil {
		log.Print("node_id is invalid uuid")
		statusCode := http.StatusBadRequest
		return ctx.Status(statusCode).JSON(schemes.NewSimpleWebError(statusCode, "node_id is invalid uuid"))
	}

	note, err := r.service.GetNote(ctx.UserContext(), noteID)
	if err != nil {
		log.Print("note not found")
		statusCode := http.StatusBadRequest
		return ctx.Status(statusCode).JSON(schemes.NewSimpleWebError(statusCode, "error with getting note"))
	}

	if note == nil {
		log.Print("note not found")
		statusCode := http.StatusNotFound
		return ctx.Status(statusCode).JSON(schemes.NewSimpleWebError(statusCode, "note is not found"))
	}

	return ctx.Status(http.StatusOK).JSON(
		schemes.CreateNoteResponse{
			Status: schemes.StatusOk,
			Data:   *note,
		},
	)
}

func (r *Resolver) GetNotes(ctx *fiber.Ctx) error {
	notes, err := r.service.GetNotes(ctx.UserContext())
	if err != nil {
		log.Print(err.Error())
		statusCode := http.StatusInternalServerError
		return ctx.Status(statusCode).JSON(schemes.NewSimpleWebError(statusCode, "error with getting notes"))
	}

	if notes == nil {
		return ctx.Status(http.StatusOK).JSON(
			schemes.GetNotesResponse{
				Status: schemes.StatusOk,
				Data:   []models.Note{},
			},
		)
	}

	return ctx.Status(http.StatusOK).JSON(
		schemes.GetNotesResponse{
			Status: schemes.StatusOk,
			Data:   notes,
		},
	)
}

func (r *Resolver) addNote(ctx *fiber.Ctx) error {
	var request models.NoteData
	if err := ctx.BodyParser(&request); err != nil {
		log.Printf("Parse body error: %s\n", err.Error())
		statusCode := http.StatusInternalServerError
		return ctx.Status(statusCode).JSON(schemes.NewSimpleWebError(statusCode, "error with parsing request body"))
	}

	if request.Title == "" {
		log.Print("note title is missing")
		statusCode := http.StatusNotFound
		return ctx.Status(statusCode).JSON(schemes.NewSimpleWebError(statusCode, "note title is missing"))
	}

	note, err := r.service.CreateNote(
		ctx.UserContext(),
		models.Note{
			Title:   request.Title,
			Content: request.Content,
		},
	)
	if err != nil {
		log.Println(err.Error())
		statusCode := http.StatusInternalServerError
		return ctx.Status(statusCode).JSON(schemes.NewSimpleWebError(statusCode, "error with creating note"))
	}

	return ctx.Status(http.StatusCreated).JSON(
		schemes.CreateNoteResponse{
			Status: schemes.StatusOk,
			Data:   *note,
		},
	)
}

func (r *Resolver) updateNote(ctx *fiber.Ctx) error {
	var request models.NoteData
	if err := ctx.BodyParser(&request); err != nil {
		log.Printf("Parse body error: %s\n", err.Error())
		statusCode := http.StatusInternalServerError
		return ctx.Status(statusCode).JSON(schemes.NewSimpleWebError(statusCode, "error with parsing request body"))
	}

	noteID := ctx.Params("note_id", "")
	if noteID == "" {
		log.Print("node_id is empty string")
		statusCode := http.StatusBadRequest
		return ctx.Status(statusCode).JSON(schemes.NewSimpleWebError(statusCode, "node_id is empty string"))
	}

	if _, err := uuid.Parse(noteID); err != nil {
		log.Print("node_id is invalid uuid")
		statusCode := http.StatusBadRequest
		return ctx.Status(statusCode).JSON(schemes.NewSimpleWebError(statusCode, "node_id is invalid uuid"))
	}

	note, err := r.service.UpdateNote(
		ctx.UserContext(),
		noteID,
		models.Note{
			ID:      noteID,
			Title:   request.Title,
			Content: request.Content,
		},
	)
	if err != nil {
		log.Println(err.Error())
		statusCode := http.StatusInternalServerError
		return ctx.Status(statusCode).JSON(schemes.NewSimpleWebError(statusCode, "error with updating note"))
	}

	return ctx.Status(http.StatusOK).JSON(
		schemes.UpdateNoteResponse{
			Status: schemes.StatusOk,
			Data:   *note,
		},
	)
}

func (r *Resolver) deleteNote(ctx *fiber.Ctx) error {
	noteID := ctx.Params("note_id", "")
	if noteID == "" {
		log.Print("node_id is empty string")
		statusCode := http.StatusBadRequest
		return ctx.Status(statusCode).JSON(schemes.NewSimpleWebError(statusCode, "node_id is empty string"))
	}

	if _, err := uuid.Parse(noteID); err != nil {
		log.Print("node_id is invalid uuid")
		statusCode := http.StatusBadRequest
		return ctx.Status(statusCode).JSON(schemes.NewSimpleWebError(statusCode, "node_id is invalid uuid"))
	}

	err := r.service.DeleteNote(ctx.UserContext(), noteID)
	if err != nil {
		log.Print(err.Error())
		var statusCode int

		if errors.Is(err, notes.ErrNoChange) {
			statusCode = http.StatusNotFound
			return ctx.Status(statusCode).JSON(schemes.NewSimpleWebError(statusCode, "note is not found"))
		}

		statusCode = http.StatusInternalServerError
		return ctx.Status(statusCode).JSON(schemes.NewSimpleWebError(statusCode, "error with deleting note"))
	}

	return ctx.Status(http.StatusOK).JSON(
		schemes.DeleteNoteResponse{
			Status: schemes.StatusOk,
		},
	)
}
