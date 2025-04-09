package main

import (
	"lva100/go-fiber/config"
	"lva100/go-fiber/internal/home"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.Init()
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()
	log.SetLevel(log.Level(logConfig.Level))
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, world!!!")
	// })
	home.NewHandler(app)
	app.Listen(":3000")
}
