package postgres

import (
	"fmt"

	cnst "agileful.com/queries/cmd/queries/internal"
	"agileful.com/queries/cmd/queries/models"
)

func (p PostgresRepo) PostgresListAllQueries(queryType, sortType string, pageNumber, perPage int) (*[]models.StatsOutput, int64, error) {

	queriesList := &[]models.StatsOutput{}
	command := "select * from pg_stat_statements"

	if queryType != cnst.Empty {
		command = fmt.Sprintf("%s where LOWER(query) like '%s%%'", command, queryType)
	}

	if sortType != cnst.Empty && (sortType == "asc" || sortType == "desc") {
		command = fmt.Sprintf("%s order by mean_time %s", command, sortType)
	}

	var items_count int64
	DB.Raw(command).Count(&items_count)

	command = fmt.Sprintf("%s LIMIT %d OFFSET %d", command, perPage, (pageNumber-1)*perPage)

	if err := DB.Raw(command).Scan(&queriesList).Error; err != nil {
		return nil, 0, err
	}

	return queriesList, items_count, nil
}
