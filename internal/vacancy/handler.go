package vacancy

import (
	"lva100/go-fiber/pkg/tmpladapter"
	"lva100/go-fiber/views/components"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type VacancyHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &VacancyHandler{
		router:       router,
		customLogger: customLogger,
	}
	vacancyGroup := h.router.Group("/vacancy")
	vacancyGroup.Post("/", h.createVacancy)
}

func (h *VacancyHandler) createVacancy(c *fiber.Ctx) error {
	email := c.FormValue("email")
	var component templ.Component
	if email == "" {
		component = components.Notification("Не заполнено поле email.", components.NotificationFail)
		return tmpladapter.Render(c, component)
	}
	component = components.Notification("Вакансия успешно создана.", components.NotificationSuccess)
	return tmpladapter.Render(c, component)
}
