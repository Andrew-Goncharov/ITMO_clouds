package api

import (
	"note_storage_service/api/rest/handlers"
	"note_storage_service/pkg/postgres"
	"note_storage_service/pkg/rest"
)

type Container struct {
	Server         *rest.Server
	Resolver       *handlers.Resolver
	PostgresClient *postgres.Client
}

func NewContainer(
	server *rest.Server,
	resolver *handlers.Resolver,
	postgresClient *postgres.Client,
) Container {
	return Container{
		Resolver:       resolver,
		Server:         server,
		PostgresClient: postgresClient,
	}
}
