package add

import (
	"context"
	"errors"
	"note_storage_service/internal/models"
	"note_storage_service/internal/store/notes"
	"note_storage_service/pkg/postgres"

	"github.com/google/uuid"
	"github.com/urfave/cli/v2"
)

var ErrNoArguments = errors.New("no arguments provided")
var ErrTooManyArguments = errors.New("unexpected number of arguments")

var Cmd = cli.Command{
	Name:  "add",
	Usage: "add note to database:\nthe first argument is the title;\nthe second argument is the content",
	Flags: cmdFlags,
	OnUsageError: func(ctx *cli.Context, err error, isSubcommand bool) error {
		return cli.ShowCommandHelp(ctx, "add")
	},
	Action: add,
}

func add(ctx *cli.Context) error {
	if ctx.NArg() == 0 {
		return ErrNoArguments
	} else if ctx.NArg() > 2 {
		return ErrNoArguments
	}

	title := ctx.Args().Get(0)
	content := ctx.Args().Get(1)

	config := postgres.Config{
		User:       ctx.String("db-user"),
		Password:   ctx.String("db-password"),
		Host:       ctx.String("db-host"),
		DBName:     ctx.String("db-dbname"),
		DisableTLS: true,
	}
	client, err := postgres.NewClient(config)
	if err != nil {
		return err
	}

	store := notes.NewStore(client)
	return store.CreateNote(context.Background(), models.Note{
		ID:      uuid.New().String(),
		Title:   title,
		Content: content,
	})
}
