package tasks

import (
	"context"
	"log"
	"server/lib"
	"server/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetTasks(c *fiber.Ctx) error {
	var TaskList []models.Task

	// Database client
	client := lib.MongoDBClient
	if client == nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{"msg": "Database not connected"})
	}

	collection := client.Database("task_assignment").Collection("tasks")
	cursor, err := collection.Find(context.TODO(), bson.D{{Key: "isdone", Value: true}})

	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{"msg": "Error getting data"})
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			log.Println("Error decoding task:", err)
			return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{"msg": "Error decoding task"})
		}
		TaskList = append(TaskList, task)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"msg": "Task list", "tasks": TaskList})
}
