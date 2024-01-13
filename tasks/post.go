package tasks

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"todo/db"
)

func insertTasks(c *fiber.Ctx) error {
	var task = new(Tasks)
	err := c.BodyParser(task)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid json")
	}

	id, err := db.Insert("tasks", task)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	task.ID = id
	return c.Status(http.StatusCreated).JSON(task)
}
