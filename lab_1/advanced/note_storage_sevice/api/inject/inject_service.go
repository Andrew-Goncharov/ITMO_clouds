package inject

import (
	"note_storage_service/pkg/postgres"

	"note_storage_service/internal/services/notes"
	storage "note_storage_service/internal/store/notes"

	"github.com/google/wire"
	"github.com/urfave/cli/v2"
)

var serviceSet = wire.NewSet(
	providePosgresConfig,
	providePosgresClient,
	storage.NewStore,
	notes.NewService,
)

func providePosgresConfig(c *cli.Context) postgres.Config {
	return postgres.Config{
		User:       c.String("db-user"),
		Password:   c.String("db-password"),
		Host:       c.String("db-host"),
		DBName:     c.String("db-dbname"),
		DisableTLS: true,
	}
}

func providePosgresClient(cfg postgres.Config) (*postgres.Client, error) {
	return postgres.NewClient(cfg)
}
