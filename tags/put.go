package tags

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"todo/db"
)

func updateTags(c *fiber.Ctx) error {
	var tag = new(Tag)

	err := c.BodyParser(tag)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid json")
	}

	var tagAtt = new(Tag)

	err = db.UpdateByID("tags", c.Params("id"), tag, tagAtt)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	return c.Status(http.StatusOK).JSON(tagAtt)
}
