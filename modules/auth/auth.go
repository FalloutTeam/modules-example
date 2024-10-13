package auth

import (
	"database/sql"
	"example/modules/auth/handler"
	authHandler "example/modules/auth/handler/auth"
	"example/modules/auth/repository"
	authService "example/modules/auth/service/auth"

	"github.com/go-chi/chi"
)

func InitModule(router *chi.Mux, db *sql.DB) {
	authRepo := repository.NewRepository(db)

	authService := authService.NewService(authRepo)

	handlers := handler.Handlers{
		AuthHandler: *authHandler.NewAuthHandler(authService),
	}

	// Регистрация ручек
	router.Route("/auth", func(r chi.Router) {
		r.Post("/login", handlers.AuthHandler.Login)
		// Другие маршруты аутентификации
	})
}
