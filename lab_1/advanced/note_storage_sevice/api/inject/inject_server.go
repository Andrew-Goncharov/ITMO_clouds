package inject

import (
	"note_storage_service/api/rest/handlers"
	"note_storage_service/internal/services/notes"
	"note_storage_service/pkg/rest"

	"github.com/google/wire"
)

var serverSet = wire.NewSet(
	provideFiber,
	provideRestResolver,
)

func provideFiber() (*rest.Server, error) {
	return rest.NewServer()
}

func provideRestResolver(server *rest.Server, service notes.Service) *handlers.Resolver {
	return handlers.NewResolver(server, service)
}
