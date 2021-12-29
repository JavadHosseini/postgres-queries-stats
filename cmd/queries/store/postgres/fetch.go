package postgres

import (
	"fmt"

	"agileful.com/queries/cmd/queries/models"
)

func (p PostgresRepo) PostgresListAllQueries(queryType, sortType string) (*[]models.StatsOutput, error) {

	models := &[]models.StatsOutput{}
	err := DB.Raw(fmt.Sprintf("select * from pg_stat_statements where LOWER(query) like '%s%%' order by mean_time %s;", queryType, sortType)).Scan(&models).Error
	if err != nil {
		return nil, err
	}

	return models, nil
}
