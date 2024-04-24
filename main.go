package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/yash-sudo-web/rssagg/internal/database"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load()

	portstr := os.Getenv("PORT")
	if portstr == "" {
		log.Fatal("PORT not found in .env file")
	} else {
		fmt.Println("Port: ", portstr)
	}

	dbstr := os.Getenv("DB_URL")
	if dbstr == "" {
		log.Fatal("DB_URL not found in .env file")
	}

	con, err1 := sql.Open("postgres", dbstr)
	if err1 != nil {
		log.Fatal("Cannot connect to the DB.", err1)
	}

	apicfg := apiConfig{
		DB: database.New(con),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/health", health)
	v1Router.Get("/err", handle_err)
	v1Router.Post("/createuser", apicfg.handleCreateUser)
	v1Router.Get("/getuser", apicfg.middlewareAuth(apicfg.handleGetUser))
	v1Router.Post("/createfeed", apicfg.middlewareAuth(apicfg.handleCreateFeed))
	v1Router.Get("/getfeed", apicfg.handleGetFeed)
	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portstr,
	}

	fmt.Printf("Server is running on http://localhost:%s\n", portstr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
