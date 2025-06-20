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

	// Import driver untuk sqlite
	_ "modernc.org/sqlite"
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

	dbPath := os.Getenv("GOOSE_DBSTRING")
	if dbPath == "" {
		log.Fatal("GOOSE_DBSTRING is not found in env")
	}

	// * Buka koneksi ke SQLite menggunakan driver "sqlite"
	conn, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatal("Can't connect to database: ", err)
	}

	// * Buat instance dari sqlc Queries
	db := sqlc.New(conn)
	apiCfg := apiConfig{DB: db}

	router := chi.NewRouter()

	// * Middleware
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
