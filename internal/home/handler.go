package home

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
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
	data := struct {
		Count   int
		IsAdmin bool
		CanUse  bool
	}{Count: 1, IsAdmin: true, CanUse: true}
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
