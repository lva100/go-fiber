package home

import (
	"lva100/go-fiber/internal/vacancy"
	"lva100/go-fiber/pkg/tmpladapter"
	"lva100/go-fiber/views"
	"lva100/go-fiber/views/components"
	"math"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
	repository   *vacancy.VacancyRepository
	store        *session.Store
}

type User struct {
	Id   int
	Name string
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger, repository *vacancy.VacancyRepository, store *session.Store) {
	h := &HomeHandler{
		router:       router,
		customLogger: customLogger,
		repository:   repository,
		store:        store,
	}
	h.router.Get("/", h.home)
	h.router.Get("/login", h.login)
	h.router.Get("/404", h.err)

	h.router.Post("/api/login", h.apiLogin)
	h.router.Get("/api/logout", h.apiLogout)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	PAGE_ITEMS := 3
	page := c.QueryInt("page", 1)
	sess, err := h.store.Get(c)
	if err != nil {
		panic(err)
	}
	userEmail := ""
	if email, ok := sess.Get("email").(string); ok {
		userEmail = email
	}
	c.Locals("email", userEmail)
	count := h.repository.CountAll()
	vacancies, err := h.repository.GetAll(PAGE_ITEMS, (page-1)*PAGE_ITEMS)
	if err != nil {
		h.customLogger.Error().Msg(err.Error())
		return c.SendStatus(500)
	}
	component := views.Main(vacancies, int(math.Ceil(float64(count/PAGE_ITEMS))), page)
	return tmpladapter.Render(c, component, http.StatusOK)
}

func (h *HomeHandler) login(c *fiber.Ctx) error {
	component := views.Login()
	sess, err := h.store.Get(c)
	if err != nil {
		panic(err)
	}
	sess.Set("name", "Vasya Pupkin")
	if err := sess.Save(); err != nil {
		panic(err)
	}
	return tmpladapter.Render(c, component, http.StatusOK)
}

func (h *HomeHandler) apiLogin(c *fiber.Ctx) error {
	form := LoginForm{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}
	if form.Email == "a@a.ru" && form.Password == "1" {
		sess, err := h.store.Get(c)
		if err != nil {
			panic(err)
		}
		sess.Set("email", form.Email)
		if err := sess.Save(); err != nil {
			panic(err)
		}
		c.Response().Header.Add("Hx-Redirect", "/")
		return c.Redirect("/", http.StatusOK)
	}
	component := components.Notification("Неверный логин или пароль", components.NotificationFail)
	return tmpladapter.Render(c, component, http.StatusBadRequest)
}

func (h *HomeHandler) apiLogout(c *fiber.Ctx) error {
	sess, err := h.store.Get(c)
	if err != nil {
		panic(err)
	}
	sess.Delete("email")
	if err := sess.Save(); err != nil {
		panic(err)
	}
	c.Response().Header.Add("Hx-Redirect", "/")
	return c.Redirect("/", http.StatusOK)
}

func (h *HomeHandler) err(c *fiber.Ctx) error {
	h.customLogger.Info().
		Bool("isAdmin", true).
		Str("email", "test@test.ru").
		Int("Id", 10).
		Msg("Инфо")
	return c.SendString("Error page.")
}
