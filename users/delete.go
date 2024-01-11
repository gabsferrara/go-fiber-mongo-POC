package users

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"todo/db"
)

func deleteUser(c *fiber.Ctx) error {
	err := db.DeleteByID("users", c.Params("id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(http.StatusNoContent).SendString("")
}
