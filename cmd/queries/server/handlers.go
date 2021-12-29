package server

import (
	"agileful.com/queries/cmd/queries/service/queries"
)

type handlers struct {
	queriesService queries.IQueriesService
}

func NewHandlers(queriesService queries.IQueriesService,
) *handlers {
	return &handlers{
		queriesService: queriesService,
	}
}
