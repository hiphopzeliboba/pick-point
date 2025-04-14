package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"pickpoint/internal/api/handler/intake"
	"pickpoint/internal/api/handler/pickpoint"
	"pickpoint/internal/api/handler/user"
)

func NewRouter(
	userHandler *user.UserHandler,
	pickPointHandler *pickpoint.PickPointHandler,
	intakeHandler *intake.IntakeHandler,
) *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	// Публичные маршруты
	r.Post("/register", userHandler.Register)
	r.Post("/login", userHandler.Login)

	// Защищенные маршруты (в будущем добавим middleware для проверки JWT)
	r.Route("/pvz", func(r chi.Router) {
		r.Post("/", pickPointHandler.Create)
		r.Get("/", pickPointHandler.List)
		r.Post("/{pvzId}/close_last_reception", intakeHandler.CloseLastReception)
	})

	r.Route("/receptions", func(r chi.Router) {
		r.Post("/", intakeHandler.Create)
	})

	return r
}
