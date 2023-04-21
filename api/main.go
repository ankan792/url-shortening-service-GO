package main

import (
	"log"
	"os"

	"github.com/ankan792/url-shortening-service-GO/api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Post("/api/v1", routes.ShortenURL)
	app.Get(":id", routes.ResolveURL)
	//serve static
	app.Static("/", "./static")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	app := fiber.New()

	app.Use(logger.New())

	setupRoutes(app)

	log.Fatalln(app.Listen(os.Getenv("APP_PORT")))
}
