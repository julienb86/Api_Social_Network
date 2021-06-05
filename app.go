package main

import (
	"github.com/gofiber/fiber/v2"
	routes "github.com/julienb86/Api_Social_Network/Routes"
)

func main() {
	app := fiber.New()

	SetUpRoutes(app)

	// Listen on server 8000 and catch error if any
	err := app.Listen(":8000")

	if err != nil {
		panic(err)
	}

}

func SetUpRoutes(app *fiber.App) {
	api := app.Group("/api")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the endpoint 😉",
		})
	})
	routes.PostRoutes(api.Group("/posts"))
}
