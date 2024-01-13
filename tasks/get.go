package tasks

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"todo/db"
)

func getAll(c *fiber.Ctx) error {
	var task = new([]Tasks)
	err := db.Find("tasks", task)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.JSON(task)
}

func getByID(c *fiber.Ctx) error {
	var task = new(Tasks)
	err := db.FindByID("tasks", c.Params("id"), task)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err)
	}
	return c.JSON(task)
}
