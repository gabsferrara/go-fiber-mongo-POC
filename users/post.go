package users

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"todo/db"
)

func addUser(c *fiber.Ctx) error {
	body := new(User)
	if err := c.BodyParser(body); err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid json")
	}

	id, err := db.Insert("users", body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	body.ID = id
	return c.Status(http.StatusCreated).JSON(body)

}
