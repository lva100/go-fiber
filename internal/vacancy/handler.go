package vacancy

import (
	"lva100/go-fiber/pkg/tmpladapter"
	"lva100/go-fiber/pkg/validator"
	"lva100/go-fiber/views/components"
	"time"

	"github.com/a-h/templ"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
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
	form := VacancyCreateForm{
		Email: c.FormValue("email"),
	}
	errors := validate.Validate(
		&validators.EmailIsPresent{Name: "Email", Field: form.Email, Message: "Email не задан или неверный"},
	)
	time.Sleep(time.Second * 2)
	var component templ.Component
	if len(errors.Errors) > 0 {
		component = components.Notification(validator.FormatErrors(errors), components.NotificationFail)
		return tmpladapter.Render(c, component)
	}
	component = components.Notification("Вакансия успешно создана.", components.NotificationSuccess)
	return tmpladapter.Render(c, component)
}
