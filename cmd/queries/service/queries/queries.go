package queries

import (
	"net/http"

	"agileful.com/queries/cmd/queries/models"
	"agileful.com/queries/cmd/queries/repository"
)

type IQueriesService interface {
	ListAllQueries(queryType, sortType string) (*[]models.StatsOutput, int, error)
}

func (s Service) ListAllQueries(queryType, sortType string) (*[]models.StatsOutput, int, error) {
	queries, err := s.queriesRepo.PostgresListAllQueries(queryType, sortType)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return queries, http.StatusOK, nil
}

type Service struct {
	queriesRepo repository.QueriesRepository
}

func NewQueriesService(queriesRepo repository.QueriesRepository) IQueriesService {
	return &Service{
		queriesRepo: queriesRepo}
}
