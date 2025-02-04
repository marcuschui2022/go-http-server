package main

import (
	"context"
	"database/sql"
	"example.com/marcus/go-http-server/internal/database"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

// Chirpy start!
func main() {
	err := godotenv.Load()
	if err != nil {
		return
	}
	dbURL := os.Getenv("DB_URL")
	fmt.Println(dbURL)
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)
	user, err := dbQueries.CreateUser(context.Background(), "marcus@marcus.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)

	const filepathRoot = "."
	const apiPrefix = "/api/"
	const port = "8080"
	apiCfg := apiConfig{
		fileserverHits: atomic.Int32{},
	}

	mux := http.NewServeMux()
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app/", http.FileServer(http.Dir(filepathRoot)))))

	mux.HandleFunc("GET "+apiPrefix+"healthz", handlerReadiness)

	mux.HandleFunc("POST /admin/reset", apiCfg.handlerReset)
	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	mux.HandleFunc("POST /api/validate_chirp", handlerChirpsValidate)

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
