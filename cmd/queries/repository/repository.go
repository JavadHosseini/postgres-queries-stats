package repository

import "agileful.com/queries/cmd/queries/models"

type QueriesRepository interface {
	PostgresListAllQueries(queryType, sortType string, pageNumber, perPage int) (*[]models.StatsOutput, int64, error)
}
