package home

import (
	"bytes"
	"text/template"

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
	// return c.SendString("Home page.")
	// return fiber.ErrBadRequest
	// return fiber.NewError(400, "Custom error")
	// "{{.Count}} - число пользователей"
	tmpl, err := template.New("test").Parse("{{.Count}} - число пользователей")
	data := struct{ Count int }{Count: 1}
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Template error")
	}
	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Template compile error")
	}
	return c.Send(tpl.Bytes())
}

func (h *HomeHandler) err(c *fiber.Ctx) error {
	h.customLogger.Info().
		Bool("isAdmin", true).
		Str("email", "test@test.ru").
		Int("Id", 10).
		Msg("Инфо")
	// logger := zerolog.New(os.Stderr).With().Timestamp().Logger().Level(1)
	// logger.Info().Msg("Logger 2")
	return c.SendString("Error page.")
}
