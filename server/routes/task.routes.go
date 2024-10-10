package routes

import (
	"server/controller/tasks"

	"github.com/gofiber/fiber/v2"
)

func TaskRoute(app *fiber.App) {

	taskRoute := app.Group("/api/task")

	taskRoute.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello from task route")
	})

	taskRoute.Get("/", tasks.GetTasks)
	taskRoute.Put("/", tasks.AddTask)

}
