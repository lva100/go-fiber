package main

import (
	"lva100/go-fiber/config"
	"lva100/go-fiber/internal/home"
	"lva100/go-fiber/internal/vacancy"
	"lva100/go-fiber/pkg/database"
	"lva100/go-fiber/pkg/logger"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()
	dbConfig := config.NewDatabaseConfig()
	customLogger := logger.NewLogger(logConfig)
	app := fiber.New()

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))
	app.Use(recover.New())
	app.Static("/public", "./public")
	dbPool := database.CreateDbPool(dbConfig, customLogger)
	defer dbPool.Close()
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, world!!!")
	// })
	//Repositories
	vacancyRepo := vacancy.NewVacancyRepository(dbPool, customLogger)
	//Handlers
	home.NewHandler(app, customLogger)
	vacancy.NewHandler(app, customLogger, vacancyRepo)
	app.Listen(":3000")
}
