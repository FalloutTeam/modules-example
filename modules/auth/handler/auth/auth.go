package auth

import (
	"encoding/json"
	"example/modules/auth/service/auth"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AuthHandler struct {
	service auth.AuthService
}

func NewAuthHandler(service auth.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	token, err := h.service.Login(creds.Username, creds.Password)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func RegisterRoutes(router *chi.Mux, service auth.AuthService) {
	handler := NewAuthHandler(service)
	router.Route("/auth", func(r chi.Router) {
		r.Post("/login", handler.Login)
		// Другие маршруты аутентификации
	})
}
