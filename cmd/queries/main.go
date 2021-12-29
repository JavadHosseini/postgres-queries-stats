package main

import (
	"fmt"
	"log"
	"os"

	cnst "agileful.com/queries/cmd/queries/internal"
	"agileful.com/queries/cmd/queries/server"
	"agileful.com/queries/cmd/queries/service/queries"
	"agileful.com/queries/cmd/queries/store/postgres"
)

func main() {
	// initialize a postgres connection
	postgresDatabase, err := postgres.Init()
	if err != nil {
		log.Fatalf(cnst.InitDBError, err)
	}

	// Close db connection in defer
	defer func() {
		sqlDB, _ := postgresDatabase.DB()
		if err := sqlDB.Close(); err != nil {
			log.Fatalf(cnst.CloseDBError, err)
		}
	}()

	postgresRepo := postgres.PostgresRepo{DB: postgresDatabase}

	queriesService := queries.NewQueriesService(postgresRepo)

	httpServer := server.NewHttpServer(queriesService)
	var serverPort string
	// read port from os env
	if serverPort = os.Getenv("SERVER_PORT"); len(serverPort) == 0 {
		// set default port value
		serverPort = fmt.Sprintf("%v", cnst.ServerPort)
	}
	// Starting the http server
	log.Fatalf(httpServer.Start(serverPort).Error())
}
