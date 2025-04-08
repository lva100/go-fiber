package main

import (
	"log"
	"lva100/go-fiber/config"
	"lva100/go-fiber/internal/home"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.Init()
	app := fiber.New()
	app.Use(recover.New())
	dbConf := config.NewDatabaseConfig()
	log.Println(dbConf)
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, world!!!")
	// })
	home.NewHandler(app)
	app.Listen(":3000")
}
