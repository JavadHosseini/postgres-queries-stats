package postgres

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPostgresRepo_PostgresListAllQueries(t *testing.T) {
	db, err := Init()

	if err != nil {
		t.Error(err.Error())
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	repo := PostgresRepo{DB: db}

	Convey("tests", t, func() {
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
			{
				QueryType:  "delete",
				SortType:   "desc",
				PageNumber: 2,
				PerPage:    5,
				HasError:   false,
				ErrMsg:     "",
				status:     http.StatusOK,
			},
			{
				QueryType:  "delete",
				SortType:   "desc",
				PageNumber: 2,
				PerPage:    5,
				HasError:   false,
				ErrMsg:     "",
				status:     http.StatusOK,
			}, {
				QueryType:  "delete",
				SortType:   "desc",
				PageNumber: 0,
				PerPage:    0,
				HasError:   false,
				ErrMsg:     "",
				status:     http.StatusOK,
			},
		}
		for _, t := range tests {
			queries, total, err := repo.PostgresListAllQueries(t.QueryType, t.SortType, t.PageNumber, t.PerPage)

			if t.HasError {
				So(err, ShouldNotEqual, nil)
				So(total, ShouldEqual, 0)
			} else {
				So(err, ShouldEqual, nil)
				So(queries, ShouldNotEqual, nil)
			}
		}
	})
}
