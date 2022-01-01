package queries

import (
	"net/http"
	"testing"

	"agileful.com/queries/cmd/queries/store/postgres"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_handlers_Get(t *testing.T) {

	db, err := postgres.Init()
	if err != nil {
		t.Error(err.Error())
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	repo := postgres.PostgresRepo{DB: db}

	Convey("test", t, func() {
		queriesService := NewQueriesService(repo)

		tests := []struct {
			QueryType  string
			SortType   string
			PageNumber int
			PerPage    int
			HasError   bool
			ErrMsg     string
			status     int
		}{
			{
				QueryType:  "select",
				SortType:   "asc",
				PageNumber: 2,
				PerPage:    5,
				HasError:   false,
				ErrMsg:     "",
				status:     http.StatusOK,
			},
		}
		for _, t := range tests {
			queries, _, status, err := queriesService.ListAllQueries(t.QueryType, t.SortType, t.PageNumber, t.PerPage)
			if t.HasError {
				So(err, ShouldNotEqual, nil)
			} else {
				So(err, ShouldEqual, nil)
				So(queries, ShouldNotEqual, nil)

			}
			So(status, ShouldEqual, t.status)
		}
	})
}
