package rest

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app *fiber.App
}

func NewServer() (*Server, error) {
	var appConfig = fiber.Config{
		BodyLimit:             1024 * 1024 * 1024,
		DisableStartupMessage: true,
		CaseSensitive:         true,
		StrictRouting:         true,
	}

	server := &Server{
		app: fiber.New(appConfig),
	}

	return server, nil
}

func (server *Server) Run() {
	log.Print("REST server was started")
	if err := server.app.Listen("0.0.0.0:8080"); err != nil {
		log.Printf("the REST server was stopped: %s", err.Error())
	}
}

func (server *Server) App() *fiber.App {
	return server.app
}

func (server *Server) Shutdown() error {
	return server.app.Shutdown()
}
