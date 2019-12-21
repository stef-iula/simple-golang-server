package main

import (
	"handlers"
	"log"
	"net/http"
	"os"
	"utils"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	BuildEnv := utils.GetEnv("BUILD_ENV", "local")
	log.Printf("Build env: %s", BuildEnv)
	log.Println("Iinitalizing router..")

	// initialize router and handlers
	r := mux.NewRouter()

	// healthcheck route
	r.HandleFunc("/", handlers.Healthy).Methods(http.MethodGet, http.MethodOptions)

	// api route
	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/hello-world", handlers.HelloWorldHandler).Methods(http.MethodGet, http.MethodPost, http.MethodOptions)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	routerHandler := c.Handler(r)

	loggedRouter := ghandlers.LoggingHandler(os.Stdout, routerHandler)
	log.Print("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", loggedRouter))
}
