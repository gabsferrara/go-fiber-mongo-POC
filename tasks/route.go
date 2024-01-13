package tasks

import "github.com/gofiber/fiber/v2"

func SetRoutes(r fiber.Router) {
	tasks := r.Group("/tasks")

	tasks.Post("/", insertTasks)
	tasks.Get("/", getAll)
	tasks.Get("/:id", getByID)
	tasks.Put("/:id", uptadeTasks)
	tasks.Delete("/:id", deleteTasks)

}
