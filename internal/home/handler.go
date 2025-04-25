package home

import (
	"lva100/go-fiber/pkg/tmpladapter"
	"lva100/go-fiber/views"
	"net/http"

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
	h.router.Get("/", h.home)
	h.router.Get("/404", h.err)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	component := views.Main()
	return tmpladapter.Render(c, component, http.StatusOK)
}

func (h *HomeHandler) err(c *fiber.Ctx) error {
	h.customLogger.Info().
		Bool("isAdmin", true).
		Str("email", "test@test.ru").
		Int("Id", 10).
		Msg("Инфо")
	return c.SendString("Error page.")
}
