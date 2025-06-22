package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/Jasperino64/goserver/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	fileServerHits atomic.Int32
	dbQueries      *database.Queries
	platform	   string
}

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	platform := os.Getenv("PLATFORM")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable not set")
	}
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()
	dbQueries := database.New(db)
	
	const filepathRoot = "."
	const port = "8080"
	config := &apiConfig{
		dbQueries: dbQueries,
		platform: platform,
	}
	mux := http.NewServeMux()

	
	mux.Handle("/app/", http.StripPrefix("/app", config.middlewareMetricsInc(http.FileServer(http.Dir(filepathRoot)))))
	mux.HandleFunc("GET /api/healthz", handlerReadiness)

	mux.HandleFunc("GET /api/metrics", config.handlerMetrics)
	mux.HandleFunc("POST /admin/reset", config.handlerReset)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	mux.HandleFunc("GET /admin/metrics", config.handlerAdminMetrics)
	mux.HandleFunc("POST /api/validate_chirp", handlerChirpsValidate)

	mux.HandleFunc("POST /api/users", config.handlerCreateUser)
	mux.HandleFunc("POST /api/chirps", config.handlerCreateChirp)
	mux.HandleFunc("GET /api/chirps", config.handlerGetChirps)
	mux.HandleFunc("GET /api/chirps/{chirp_id}", config.handlerGetChirp)

	mux.HandleFunc("POST /api/login", config.handlerLogin)

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
