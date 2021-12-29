package repository

import "agileful.com/queries/cmd/queries/models"

type QueriesRepository interface {
	PostgresListAllQueries(queryType, sortType string) (*[]models.StatsOutput, error)
}
