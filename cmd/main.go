package main

import (
	"lva100/go-fiber/config"
	"lva100/go-fiber/internal/home"
	"lva100/go-fiber/pkg/logger"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()
	customLogger := logger.NewLogger(logConfig)
	engine := html.New("./html", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))
	app.Use(recover.New())

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, world!!!")
	// })
	home.NewHandler(app, customLogger)
	app.Listen(":3000")
}
