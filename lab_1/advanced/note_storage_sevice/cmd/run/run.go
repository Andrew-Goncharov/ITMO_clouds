package run

import (
	"context"
	"log"
	"note_storage_service/api/inject"
	"os"
	"os/signal"
	"syscall"

	cli "github.com/urfave/cli/v2"
)

var Cmd = cli.Command{
	Name:  "run",
	Usage: "Start Server",
	Flags: cmdFlags,
	OnUsageError: func(ctx *cli.Context, err error, isSubCommand bool) error {
		return cli.ShowCommandHelp(ctx, "run")
	},
	Action: run,
}

func run(ctx *cli.Context) error {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	signal.Notify(sig, syscall.SIGTERM)

	c, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		select {
		case <-c.Done():
			return
		case s := <-sig:
			log.Printf("signal %s received", s.String())
			cancel()
		}
	}()

	app, err := inject.InitializeApplication(ctx)
	if err != nil {
		log.Fatalf("main: cannot initialize server: %s", err.Error())
	}

	go func() {
		app.Server.Run()
	}()

	<-c.Done()

	_ = app.Server.Shutdown()
	log.Print("context end received")

	return nil
}
