package home

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

type User struct {
	Id   int
	Name string
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &HomeHandler{
		router:       router,
		customLogger: customLogger,
	}
	// h.router.Get("/", h.home)
	api := h.router.Group("api")
	api.Get("/", h.home)
	api.Get("/error", h.err)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	users := []User{
		{Id: 1, Name: "User1"},
		{Id: 2, Name: "User2"},
		{Id: 3, Name: "User3"},
	}
	names := []string{"User1", "User2", "User3"}
	data := struct {
		Users []User
		Names []string
	}{
		Users: users,
		Names: names,
	}
	return c.Render("page", data)
	// return c.Render("page", fiber.Map{
	// 	"Count": 3,
	// })
}

func (h *HomeHandler) err(c *fiber.Ctx) error {
	h.customLogger.Info().
		Bool("isAdmin", true).
		Str("email", "test@test.ru").
		Int("Id", 10).
		Msg("Инфо")
	return c.SendString("Error page.")
}
