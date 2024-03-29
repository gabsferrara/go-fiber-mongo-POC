package tags

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"todo/db"
)

func getTagByID(c *fiber.Ctx) error {
	var tag = new(Tag)
	err := db.FindByID("tags", c.Params("id"), tag)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(http.StatusOK).JSON(tag)
}

func getTags(c *fiber.Ctx) error {
	var tag = new([]Tag)
	err := db.Find("tags", nil, tag)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(http.StatusOK).JSON(tag)
}
