package tasks

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"todo/db"
)

func uptadeTasks(c *fiber.Ctx) error {
	var tasks = new(Tasks)
	err := c.BodyParser(tasks)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("invalid json")
	}
	var result = new(Tasks)
	err = db.UpdateByID("tasks", c.Params("id"), tasks, result)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.JSON(result)
}
