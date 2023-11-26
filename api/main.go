package main

import (
	"github.com/gofiber/fiber/v2"
	"url-shortener/api/routes"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)

}

func main() {
	err := godotenv.Load()
	app := fiber.New()

}
