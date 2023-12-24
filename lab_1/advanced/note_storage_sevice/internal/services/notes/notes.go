package notes

import (
	"context"
	"database/sql"
	"errors"
	"note_storage_service/internal/models"
	storage "note_storage_service/internal/store/notes"

	"github.com/google/uuid"
)

var ErrNoChange = errors.New("no notes were deleted")
var ErrIncorrectDeletion = errors.New("the wrong number of notes were deleted")

type (
	Service interface {
		CreateNote(ctx context.Context, note models.Note) (*models.Note, error)
		GetNote(ctx context.Context, noteID string) (*models.Note, error)
		GetNotes(ctx context.Context) ([]models.Note, error)
		UpdateNote(ctx context.Context, noteID string, note models.Note) (*models.Note, error)
		DeleteNote(ctx context.Context, noteID string) error
	}

	serviceImpl struct {
		store storage.Store
	}
)

func NewService(store storage.Store) Service {
	return serviceImpl{
		store: store,
	}
}

func (service serviceImpl) CreateNote(ctx context.Context, note models.Note) (*models.Note, error) {
	noteID := uuid.New().String()
	note.ID = noteID

	err := service.store.CreateNote(ctx, note)
	if err != nil {
		return nil, err
	}

	return &models.Note{
		ID:      noteID,
		Title:   note.Title,
		Content: note.Content,
	}, nil
}

func (service serviceImpl) GetNote(ctx context.Context, noteID string) (*models.Note, error) {
	note, err := service.store.GetNote(ctx, noteID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return note, nil
}

func (service serviceImpl) GetNotes(ctx context.Context) ([]models.Note, error) {
	notes, err := service.store.GetNotes(ctx)
	if err != nil {
		return nil, err
	}

	return notes, nil
}

func (service serviceImpl) UpdateNote(ctx context.Context, noteID string, note models.Note) (*models.Note, error) {
	_, err := service.store.GetNote(ctx, noteID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = service.store.CreateNote(ctx, note)
			if err != nil {
				return nil, err
			}
			return &note, nil
		}
		return nil, err
	}

	err = service.store.UpdateNote(ctx, noteID, note)
	if err != nil {
		return nil, err
	}
	return &note, nil
}

func (service serviceImpl) DeleteNote(ctx context.Context, noteID string) error {
	countRows, err := service.store.DeleteNote(ctx, noteID)
	if err != nil {
		return err
	}

	if countRows == 0 {
		return ErrNoChange
	} else if countRows != 1 {
		return ErrIncorrectDeletion
	}

	return nil
}
