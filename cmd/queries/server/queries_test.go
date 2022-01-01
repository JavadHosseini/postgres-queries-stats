package server

import (
	"net/http"
	"testing"

	cnst "agileful.com/queries/cmd/queries/internal"
	"github.com/gofiber/fiber/v2"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/valyala/fasthttp"
)

func Test_handlers_Get(t *testing.T) {
	f, h, repo := startTestServer()
	sqlDB, _ := repo.DB.DB()
	defer sqlDB.Close()

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	Convey("tests", t, func() {
		tests := []struct {
			QueryType  string
			SortType   string
			PageNumber string
			PerPage    string
			HasError   bool
			ErrMsg     string
			Status     int
		}{
			{
				QueryType:  "select",
				SortType:   "asc",
				PageNumber: "aa",
				PerPage:    "5",
				HasError:   true,
				ErrMsg:     cnst.ErrPageNotValid,
				Status:     http.StatusBadRequest,
			},
			{
				QueryType:  "delete",
				SortType:   "desc",
				PageNumber: "1",
				PerPage:    "aa",
				HasError:   true,
				ErrMsg:     cnst.ErrPerPageNotValid,
				Status:     http.StatusBadRequest,
			}, {
				QueryType:  "delete",
				SortType:   "desc",
				PageNumber: "aa",
				PerPage:    "aa",
				HasError:   true,
				ErrMsg:     cnst.ErrPageNotValid,
				Status:     http.StatusBadRequest,
			}, {
				QueryType:  "delete",
				SortType:   "desc",
				PageNumber: "0",
				PerPage:    "10",
				HasError:   true,
				ErrMsg:     cnst.ErrPageNotValid,
				Status:     http.StatusBadRequest,
			}, {
				QueryType:  "delete",
				SortType:   "desc",
				PageNumber: "2",
				PerPage:    "0",
				HasError:   true,
				ErrMsg:     cnst.ErrPerPageNotValid,
				Status:     http.StatusBadRequest,
			},
			{
				QueryType:  "delete",
				SortType:   "desc",
				PageNumber: "2",
				PerPage:    "10",
				HasError:   false,
				ErrMsg:     "",
				Status:     http.StatusOK,
			},
		}
		for _, t := range tests {
			req.URI().QueryArgs().Set("page", t.PageNumber)
			req.URI().QueryArgs().Set("per_page", t.PerPage)
			req.URI().QueryArgs().Set("type", t.QueryType)
			req.URI().QueryArgs().Set("sort", t.SortType)
			reqCTX := &fasthttp.RequestCtx{Request: *req, Response: *resp}

			err := h.Get(f.AcquireCtx(reqCTX))

			if t.HasError {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldContainSubstring, t.ErrMsg)
				So(err.(*fiber.Error).Code, ShouldEqual, t.Status)

			} else {
				So(err, ShouldEqual, nil)
			}
		}

	})
}
