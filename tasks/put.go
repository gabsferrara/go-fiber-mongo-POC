package tasks

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"todo/db"
	"todo/tags"
)

func uptadeTasks(c *fiber.Ctx) error {
	var tasks = new(Task)
	err := c.BodyParser(tasks)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("invalid json")
	}

	var prev Task
	err = db.FindByID("tasks", c.Params("id"), &prev)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	var result = new(Task)
	err = db.UpdateByID("tasks", c.Params("id"), tasks, result)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	err = updateTagsTask(c.Params("id"), prev.Tags, result.Tags)
	if err != nil {
		err = tags.RemoveTask(c.Params("id"))
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(err.Error())
		}
		err = db.DeleteByID("tasks", c.Params("id"))
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(err.Error())
		}

		_, err = db.Insert("tasks", &result)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(err.Error())
		}

		err = tags.AddTask(result.ID.Hex(), result.Tags)
		if err != nil {
			db.DeleteByID("tasks", c.Params("id"))
			return c.Status(http.StatusInternalServerError).JSON(err.Error())
		}
	}

	return c.JSON(result)
}

func updateTagsTask(id string, ot, nt []string) error {
	mot := make(map[string]int, len(ot))

	for k, v := range ot {
		mot[v] = k
	}
	var diff []string
	for _, v := range nt {
		if _, key := mot[v]; !key {
			diff = append(diff, v)
		} else {
			delete(mot, v)
		}
	}

	if len(diff) > 0 {
		err := tags.AddTask(id, diff)
		if err != nil {
			return err
		}
	}

	if len(mot) > 0 {
		dt := make([]string, 0, len(mot))
		for k := range mot {
			dt = append(dt, k)
		}
		err := tags.RemoveTask(id, dt...)
		if err != nil {
			return err
		}
	}

	return nil
}
