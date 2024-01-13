package tasks

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"todo/db"
	"todo/tags"
)

func deleteTasks(c *fiber.Ctx) error {
	err := tags.RemoveTask(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	err = db.DeleteByID("tasks", c.Params("id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(http.StatusNoContent).JSON("")
}
