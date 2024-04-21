package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	portstr := os.Getenv("PORT")
	if portstr == "" {
		log.Fatal("PORT not found in .env file")
	} else {
		fmt.Println("Port: ", portstr)
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
	v1Router.Get("/err", err)

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
