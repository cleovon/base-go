package server

import (
	conf "base-go/src/modules/configuration"
	"base-go/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Run() {
	app := fiber.New()
	app.Use(logger.New())
	routes.Register(app)
	app.Listen("0.0.0.0:" + conf.Get("port"))
}
