package handlers

import (
	"note_storage_service/internal/services/notes"
	"note_storage_service/pkg/rest"
)

const (
	APIPrefix  = "/notes_service/api"
	APIVersion = "v1"
	pathPrefix = APIPrefix + "/" + APIVersion
)

type Resolver struct {
	server  *rest.Server
	service notes.Service
}

func NewResolver(server *rest.Server, service notes.Service) *Resolver {
	resolver := &Resolver{
		server:  server,
		service: service,
	}

	resolver.initRoutes()

	return resolver
}

func (r *Resolver) initRoutes() {
	r.server.App().Get(pathPrefix+"/get-notes/:note_id", r.GetNote)
	r.server.App().Get(pathPrefix+"/get-notes", r.GetNotes)
	r.server.App().Post(pathPrefix+"/create-notes", r.addNote)
	r.server.App().Put(pathPrefix+"/update-notes/:note_id", r.updateNote)
	r.server.App().Delete(pathPrefix+"/delete-notes/:note_id", r.deleteNote)
}
