package server

import (
	"math"
	"net/http"
	"strconv"
	"strings"

	cnst "agileful.com/queries/cmd/queries/internal"
	"github.com/gofiber/fiber/v2"
)

func (h *handlers) Get(c *fiber.Ctx) error {
	queryType := strings.ToLower(c.Query("type"))
	sortType := strings.ToLower(c.Query("sort"))
	pageNumber, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || pageNumber == 0 {
		return fiber.NewError(http.StatusBadRequest, cnst.ErrPageNotValid)
	}

	perPage, err := strconv.Atoi(c.Query("per_page", "10"))
	if err != nil || perPage == 0 {
		return fiber.NewError(http.StatusBadRequest, cnst.ErrPerPageNotValid)
	}

	queries, items_count, _, _ := h.queriesService.ListAllQueries(queryType, sortType, pageNumber, perPage)

	return c.JSON(fiber.Map{
		"queries":     queries,
		"page_number": pageNumber,
		"total":       items_count,
		"last_page":   math.Ceil(float64(items_count) / float64(perPage)),
	})
}
