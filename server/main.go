package main

import (
	"log"
	"server/config"
	"server/lib"
	"server/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// loading env file
	envErr := config.LoadEnv()
	if envErr != nil {
		log.Fatal("Process ending with .env.local missing")

	}

	// creating new fiber app
	app := fiber.New()

	// registering routes to fiber app
	routes.TaskRoute(app)

	// connecting to database
	DatabaseError := lib.ConnectDB()
	defer lib.DisconnectDB()
	if DatabaseError != nil {
		log.Fatal(DatabaseError)
	}

	// listening on PORT
	PORT := config.GetEnv("PORT", ":3000")
	err := app.Listen(PORT)
	if err != nil {
		log.Fatal(err)
	}
}
