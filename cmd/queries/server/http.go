package server

import (
	cnst "agileful.com/queries/cmd/queries/internal"
	"agileful.com/queries/cmd/queries/service/queries"
	"github.com/gofiber/fiber/v2"
)

// define new fiber
var f = fiber.New()

type httpServer struct {
	handlers *handlers
	v1       fiber.Router
}

func NewHttpServer(queriesService queries.IQueriesService,
) IServer {
	return &httpServer{
		handlers: NewHandlers(queriesService),
		v1:       f.Group(cnst.PathRouteGroupV1),
	}
}

func (h *httpServer) Start(port string) error {
	routes(h.v1, h.handlers)
	return f.Listen(":" + port)
}

func routes(g fiber.Router, h *handlers) {
	g.Get("/list", h.Get)

}
