package audit

import (
	"database/sql"
	"example/modules/audit/service/audit"
	"example/modules/auth/handler"
	"example/modules/auth/repository"

	"github.com/go-chi/chi"
)

func InitModule(router *chi.Mux, db *sql.DB) {
	auditRepo := repository.NewRepository(db)

	auditService := audit.NewAuditService(auditRepo)

	handlers := handler.Handlers{
		AuditHandler: *auditHandler.NewAuthHandler(auditService),
	}

	// Регистрация ручек
	router.Route("/audit", func(r chi.Router) {
		r.Post("/do-some", handlers.AuditHandler.DoSome)
		// Другие маршруты
	})
}
