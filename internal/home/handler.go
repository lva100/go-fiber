package home

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
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
	log.Trace("Trace")
	log.Debug("Debug")
	log.Info("Info")
	log.Warn("Warn")
	log.Error("Error")
	log.Panic("Panic")
	return c.SendString("Error page.")
}
