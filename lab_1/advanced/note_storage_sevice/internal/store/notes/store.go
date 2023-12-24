package notes

import (
	"context"
	"fmt"
	"note_storage_service/internal/models"
	"note_storage_service/pkg/postgres"
)

const noteTable = "notes"

type (
	Store interface {
		CreateNote(ctx context.Context, note models.Note) error
		GetNote(ctx context.Context, noteID string) (*models.Note, error)
		GetNotes(ctx context.Context) ([]models.Note, error)
		UpdateNote(ctx context.Context, noteID string, note models.Note) error
		DeleteNote(ctx context.Context, noteID string) (int, error)
	}

	storeImpl struct {
		postgresClient *postgres.Client
	}
)

func NewStore(postgresClient *postgres.Client) Store {
	return storeImpl{
		postgresClient: postgresClient,
	}
}

func (store storeImpl) CreateNote(ctx context.Context, note models.Note) error {
	tx, err := store.postgresClient.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx,
		fmt.Sprintf(
			"INSERT INTO %s (id,title,content) VALUES ($1,$2,$3)",
			noteTable),
		note.ID, note.Title, note.Content,
	)

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (store storeImpl) GetNote(ctx context.Context, noteID string) (*models.Note, error) {
	var note models.Note

	row := store.postgresClient.QueryRowContext(ctx,
		fmt.Sprintf(
			`SELECT id, title, content FROM %s WHERE id = $1`,
			noteTable),
		noteID,
	)

	err := row.Scan(&note.ID, &note.Title, &note.Content)
	if err != nil {
		return nil, err
	}

	if err = row.Err(); err != nil {
		return nil, err
	}

	return &note, nil
}

func (store storeImpl) GetNotes(ctx context.Context) ([]models.Note, error) {
	var notes []models.Note

	rows, err := store.postgresClient.QueryContext(ctx,
		fmt.Sprintf(
			"SELECT id, title, content FROM %s",
			noteTable,
		),
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var note models.Note
		err := rows.Scan(&note.ID, &note.Title, &note.Content)

		if err != nil {
			return nil, err
		}

		notes = append(notes, note)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return notes, nil
}

func (store storeImpl) UpdateNote(ctx context.Context, noteID string, note models.Note) error {
	tx, err := store.postgresClient.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx,
		fmt.Sprintf(
			"UPDATE %s SET title = $1, content = $2 WHERE id = $3",
			noteTable),
		note.Title, note.Content, noteID,
	)

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (store storeImpl) DeleteNote(ctx context.Context, noteID string) (int, error) {
	tx, err := store.postgresClient.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	result, err := tx.ExecContext(ctx,
		fmt.Sprintf(
			"DELETE FROM %s WHERE id = $1",
			noteTable),
		noteID,
	)

	if err != nil {
		return 0, err
	}

	countRows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(countRows), tx.Commit()
}
