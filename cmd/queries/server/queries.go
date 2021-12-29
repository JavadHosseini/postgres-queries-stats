package server

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (h *handlers) Get(c *fiber.Ctx) error {
	queryType := strings.ToLower(c.Query("type"))
	sortType := strings.ToLower(c.Query("sort"))
	// if queryType == cnst.Empty {
	// 	c.JSON(fiber.Map{
	// 		"message": "query type is not valid",
	// 	})
	// }

	queries, _, _ := h.queriesService.ListAllQueries(queryType, sortType)

	return c.JSON(queries)
}
