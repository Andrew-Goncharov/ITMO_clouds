package migrate

import (
	"fmt"
	"note_storage_service/internal/migrate"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/urfave/cli/v2"
)

var Cmd = cli.Command{
	Name:  "migrate",
	Usage: "set actual version migration",
	Flags: cmdFlags,
	OnUsageError: func(ctx *cli.Context, err error, isSubcommand bool) error {
		return cli.ShowCommandHelp(ctx, "migrate")
	},
	Action: run,
}

func run(ctx *cli.Context) error {
	migrate.Up(
		fmt.Sprintf(
			"postgres://%s:%s@%s/%s?sslmode=disable",
			ctx.String("db-user"),
			ctx.String("db-password"),
			ctx.String("db-host"),
			ctx.String("db-dbname"),
		),
		"file://../migrations",
	)

	return nil
}
