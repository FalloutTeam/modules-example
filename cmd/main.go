package main

import (
	"database/sql"
	"example/modules/audit"
	"example/modules/auth"
	"log"

	"github.com/go-chi/chi"
)

func main() {
	db, err := sql.Open("postgres", "user=youruser dbname=yourdb sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Инициализация маршрутизатора
	router := chi.NewRouter()

	auth.InitModule(router, db)
	audit.InitModule(router, db)
}
