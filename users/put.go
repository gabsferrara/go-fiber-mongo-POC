package users

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"todo/db"
)

func updateUser(c *fiber.Ctx) error {
	body := new(User)
	if err := c.BodyParser(body); err != nil {
		return c.Status(http.StatusBadRequest).JSON("invalid json")
	}

	var result User

	err := db.UpdateByID("users", c.Params("id"), body, &result)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.JSON(result)
}
