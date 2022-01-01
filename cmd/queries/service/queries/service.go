package queries

import (
	"net/http"

	"agileful.com/queries/cmd/queries/models"
	"agileful.com/queries/cmd/queries/repository"
)

type IQueriesService interface {
	ListAllQueries(queryType, sortType string, pageNumber, perPage int) (*[]models.StatsOutput, int64, int, error)
}

func (s Service) ListAllQueries(queryType, sortType string, pageNumber, perPage int) (*[]models.StatsOutput, int64, int, error) {
	queries, items_count, err := s.queriesRepo.PostgresListAllQueries(queryType, sortType, pageNumber, perPage)
	if err != nil {
		return nil, 0, http.StatusBadRequest, err
	}

	return queries, items_count, http.StatusOK, nil
}

type Service struct {
	queriesRepo repository.QueriesRepository
}

func NewQueriesService(queriesRepo repository.QueriesRepository) IQueriesService {
	return &Service{
		queriesRepo: queriesRepo}
}
