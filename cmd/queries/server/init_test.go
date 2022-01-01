package server

import (
	"log"

	cnst "agileful.com/queries/cmd/queries/internal"
	"agileful.com/queries/cmd/queries/service/queries"
	"agileful.com/queries/cmd/queries/store/postgres"
	"github.com/gofiber/fiber/v2"
)

func startTestServer() (*fiber.App, *handlers, postgres.PostgresRepo) {
	var f = fiber.New()

	postgresDatabase, err := postgres.Init()
	if err != nil {
		log.Fatalf(cnst.InitDBError, err)
	}

	postgresRepo := postgres.PostgresRepo{DB: postgresDatabase}

	queriesService := queries.NewQueriesService(postgresRepo)

	h := &handlers{queriesService: queriesService}

	return f, h, postgresRepo
}
