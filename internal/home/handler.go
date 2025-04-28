package home

import (
	"lva100/go-fiber/internal/vacancy"
	"lva100/go-fiber/pkg/tmpladapter"
	"lva100/go-fiber/views"
	"math"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
	repository   *vacancy.VacancyRepository
}

type User struct {
	Id   int
	Name string
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger, repository *vacancy.VacancyRepository) {
	h := &HomeHandler{
		router:       router,
		customLogger: customLogger,
		repository:   repository,
	}
	h.router.Get("/", h.home)
	h.router.Get("/404", h.err)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	PAGE_ITEMS := 3
	page := c.QueryInt("page", 1)
	count := h.repository.CountAll()
	vacancies, err := h.repository.GetAll(PAGE_ITEMS, (page-1)*PAGE_ITEMS)
	if err != nil {
		h.customLogger.Error().Msg(err.Error())
		return c.SendStatus(500)
	}
	component := views.Main(vacancies, int(math.Ceil(float64(count/PAGE_ITEMS))), page)
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
