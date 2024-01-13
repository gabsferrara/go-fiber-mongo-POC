package tags

import "github.com/gofiber/fiber/v2"

func SetRoutes(r fiber.Router) {
	tags := r.Group("/tags")

	tags.Get("/:id", getTagByID)
	tags.Get("/", getTags)
	tags.Delete("/:id", deleteTags)
	tags.Post("/", insertTags)
	tags.Put("/:id", updateTags)
}
