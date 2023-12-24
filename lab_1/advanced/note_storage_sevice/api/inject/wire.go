//go:build wireinject
// +build wireinject

package inject

import (
	"note_storage_service/api"

	"github.com/google/wire"
	"github.com/urfave/cli/v2"
)

func InitializeApplication(ctx *cli.Context) (api.Container, error) {
	wire.Build(
		serviceSet,
		serverSet,
		api.NewContainer,
	)

	return api.Container{}, nil
}
