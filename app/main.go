// File: app/main.go

package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Rizz404/go-clean-arch-for-covid-api/internal/repository/sqlc"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *sqlc.Queries
}

func main() {
	godotenv.Load()

	addr := os.Getenv("ADDR")
	if addr == "" {
		log.Fatal("ADDR is not found in env")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not found in env")
	}
	log.Println("Attempting to connect to CockroachDB...")

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database: ", err)
	}

	db := sqlc.New(conn)
	apiCfg := apiConfig{DB: db}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Post("/covids", apiCfg.createCovidHandler)
	v1Router.Get("/covids", apiCfg.getCovidsHandler)
	v1Router.Get("/covids/{id}", apiCfg.getCovidByIdHandler)
	// ! masih gak work urlencoded buat patch
	v1Router.Patch("/covids/{id}", apiCfg.updateCovidHandler)
	v1Router.Delete("/covids/{id}", apiCfg.deleteCovidHandler)

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    addr,
	}

	log.Printf("Server running on http://localhost%s", addr)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
