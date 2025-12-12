package main

import (
	"context"
	"database/sql"
	"embed"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/google/uuid"
	"github.com/joho/godotenv"

	"github.com/jeronimoLa/eagle/internal/database"
	"github.com/jeronimoLa/eagle/internal/edgar"

	_ "github.com/lib/pq"
)

//go:embed static/*
var staticFiles embed.FS

type apiConfig struct {
	db    *database.Queries
	edgar *edgar.EdgarConfig
}

func main() {
	apiCfg := &apiConfig{}
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Println("Database url is not set")

	} else {
		db, err := sql.Open("postgres", dbURL)
		if err != nil {
			log.Println("unable to connect to postgres")
		}
		dbQueries := database.New(db)
		apiCfg.db = dbQueries
		log.Println("Successfully connected to db")
	}

	apiCfg.edgar = edgar.New()

	type ticker struct {
		ticker string
		cik    string
	}

	listOfTickers := []ticker{
		{ticker: "TSLA", cik: "CIK0001318605"},
		{ticker: "AAPL", cik: "CIK0000320193"},
	}

	for _, mapper := range listOfTickers {
		_, err := apiCfg.db.InsertTickerCik(context.Background(), database.InsertTickerCikParams{
			ID:     uuid.New(),
			Ticker: mapper.ticker,
			Cik:    mapper.cik,
		})
		if err != nil {
			log.Println(err)
		}
	}

	mainRouter := chi.NewRouter()
	fileServer := http.FileServer(http.Dir("static/"))
	mainRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	mainRouter.Handle("/*", http.StripPrefix("/", fileServer))

	// Web Interface
	mainRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		f, err := staticFiles.Open("static/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()
		if _, err := io.Copy(w, f); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	// API
	v1 := chi.NewRouter()
	v1.Get("/ticker", apiCfg.handlerF4Filings)
	mainRouter.Mount("/v1", v1)

	http.ListenAndServe(":3000", mainRouter)

}
