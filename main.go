package main

import (
	"api_interes_compuesto/controller"
	"api_interes_compuesto/exceptions"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.NotFoundHandler = http.HandlerFunc(exceptions.NotFound)
	router.MethodNotAllowedHandler = http.HandlerFunc(exceptions.NotAllowedHandler)
	// Paths version 1.0.0
	versionPrefix := router.PathPrefix("/1.0.0").Subrouter()
	// Final path API
	finalPath := versionPrefix.PathPrefix("/compound-interest").Subrouter()

	finalPath.HandleFunc("/year", controller.CalculateByYear).Methods("GET")
	finalPath.HandleFunc("/month", controller.CalculateByMonth).Methods("GET")

	srv := &http.Server{
		Handler: router,
		Addr:    ":8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
