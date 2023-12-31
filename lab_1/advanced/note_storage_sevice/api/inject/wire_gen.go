// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package inject

import (
	"github.com/urfave/cli/v2"
	"note_storage_service/api"
	notes2 "note_storage_service/internal/services/notes"
	"note_storage_service/internal/store/notes"
)

// Injectors from wire.go:

func InitializeApplication(ctx *cli.Context) (api.Container, error) {
	server, err := provideFiber()
	if err != nil {
		return api.Container{}, err
	}
	config := providePosgresConfig(ctx)
	client, err := providePosgresClient(config)
	if err != nil {
		return api.Container{}, err
	}
	store := notes.NewStore(client)
	service := notes2.NewService(store)
	resolver := provideRestResolver(server, service)
	container := api.NewContainer(server, resolver, client)
	return container, nil
}
