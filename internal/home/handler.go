package home

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type HomeHandler struct {
	router fiber.Router
}

func NewHandler(router fiber.Router) {
	h := &HomeHandler{
		router: router,
	}
	// h.router.Get("/", h.home)
	api := h.router.Group("api")
	api.Get("/", h.home)
	api.Get("/error", h.err)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	// return c.SendString("Home page.")
	// return fiber.ErrBadRequest
	// return fiber.NewError(400, "Custom error")
	return fiber.NewError(fiber.StatusBadRequest, "Custom error")
}

func (h *HomeHandler) err(c *fiber.Ctx) error {
	log.Info().
		Bool("isAdmin", true).
		Str("email", "test@test.ru").
		Int("Id", 10).
		Msg("Инфо")
	// logger := zerolog.New(os.Stderr).With().Timestamp().Logger().Level(1)
	// logger.Info().Msg("Logger 2")
	return c.SendString("Error page.")
}
