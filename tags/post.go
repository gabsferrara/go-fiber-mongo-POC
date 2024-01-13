package tags

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"todo/db"
)

func insertTags(c *fiber.Ctx) error {
	var tag = new(Tag)
	err := c.BodyParser(tag)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	id, err := db.Insert("tags", tag)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	tag.ID = id
	return c.Status(http.StatusCreated).JSON(tag)
}
