package main

import (
	"api_interes_compuesto/controller"
	"api_interes_compuesto/exceptions"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No se encontro archivo .env, se usaran variables del sistema")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := mux.NewRouter()

	router.NotFoundHandler = http.HandlerFunc(exceptions.NotFound)
	router.MethodNotAllowedHandler = http.HandlerFunc(exceptions.NotAllowedHandler)
	// Paths version 1.0.0
	versionPrefix := router.PathPrefix("/v1.0.0").Subrouter()
	// Final path API
	finalPath := versionPrefix.PathPrefix("/compound-interest").Subrouter()

	finalPath.HandleFunc("/year", controller.CalculateByYear).Methods("GET")
	finalPath.HandleFunc("/month", controller.CalculateByMonth).Methods("GET")

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Servidor corriendo en puerto %s", port)

	log.Fatal(srv.ListenAndServe())
}
