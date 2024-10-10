package tasks

import (
	"context"
	"server/lib"
	"server/models"

	"github.com/gofiber/fiber/v2"
)

func AddTask(c *fiber.Ctx) error {
	var newTask models.Task
	bodyErr := c.BodyParser(&newTask)

	if bodyErr != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"msg": "Body is required"})
	}
	if newTask.Title == "" {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"msg": "Task title required"})
	}

	// Database client
	client := lib.MongoDBClient
	if client == nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{"msg": "Database not connected"})
	}

	collection := client.Database("task_assignment").Collection("tasks")
	result, insertErr := collection.InsertOne(context.TODO(), newTask)

	if insertErr != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{"msg": "Error inserting data"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"msg": "Task added", "id": result.InsertedID})
}
