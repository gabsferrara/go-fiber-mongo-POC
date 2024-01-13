package tasks

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"todo/db"
	"todo/tags"
)

func insertTasks(c *fiber.Ctx) error {
	var task = new(Task)
	err := c.BodyParser(task)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid json")
	}

	id, err := db.Insert("tasks", task)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	task.ID = id

	err = tags.AddTask(task.ID.Hex(), task.Tags)
	if err != nil {
		db.DeleteByID("tasks", task.ID.Hex())
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(http.StatusCreated).JSON(task)
}
