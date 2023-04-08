package routes

import (
	"base-go/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	app.Get("/healthcheck", controllers.GetHealthCheck)
}
