package main

import (
	"log"
	"note_storage_service/cmd/add"
	"note_storage_service/cmd/migrate"
	"note_storage_service/cmd/run"
	"note_storage_service/version"
	"os"

	cli "github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Usage: "note storage service",
		Commands: []*cli.Command{
			&add.Cmd,
			&migrate.Cmd,
			&run.Cmd,
		},
		Version: version.Version + "(" + version.GitCommit + ")",
		OnUsageError: func(ctx *cli.Context, err error, isSubcommand bool) error {
			return cli.ShowAppHelp(ctx)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}
