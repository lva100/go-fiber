package main

import (
	"lva100/go-fiber/config"
	"lva100/go-fiber/internal/home"
	"lva100/go-fiber/internal/sitemap"
	"lva100/go-fiber/internal/vacancy"
	"lva100/go-fiber/pkg/database"
	"lva100/go-fiber/pkg/logger"
	"lva100/go-fiber/pkg/middleware"
	"time"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres/v3"
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
	app.Static("/robots.txt", "./public/robots.txt")
	dbPool := database.CreateDbPool(dbConfig, customLogger)
	defer dbPool.Close()
	storage := postgres.New(postgres.Config{
		DB:         dbPool,
		Table:      "sessions",
		Reset:      false,
		GCInterval: 10 * time.Second,
	})
	store := session.New(session.Config{
		Storage: storage,
	})
	app.Use(middleware.AuthMiddleware(store))
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, world!!!")
	// })
	//Repositories
	vacancyRepo := vacancy.NewVacancyRepository(dbPool, customLogger)
	//Handlers
	home.NewHandler(app, customLogger, vacancyRepo, store)
	vacancy.NewHandler(app, customLogger, vacancyRepo)
	sitemap.NewHandler(app)
	app.Listen(":3000")
}
