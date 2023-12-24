package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func parsePaginationParams(ctx *fiber.Ctx) (page, limit int, err error) {
	page = ctx.QueryInt("page", 1)
	limit = ctx.QueryInt("limit", 256)

	if page <= 0 {
		err = errors.New("query param page must be greater than 0")
		return
	}

	if limit <= 0 {
		err = errors.New("query limit page must be greater than 0")
		return
	}

	return
}
