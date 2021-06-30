package main

import (
	"gorest/configurations"
	"gorest/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	configurations.InitDatabase()

	routes := mux.NewRouter().StrictSlash(true)

	getRouter := routes.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/moviment", handlers.GetMoviment)

	log.Fatal(http.ListenAndServe(":8080", routes))

}
