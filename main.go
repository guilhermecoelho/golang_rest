package main

import (
	"gorest/configurations"
	"gorest/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	go configurations.InitDatabase()

	//go configurations.InitDatabaseGorm()

	routes := mux.NewRouter().StrictSlash(true)

	getRouter := routes.Methods(http.MethodGet).Subrouter()
	putRouter := routes.Methods(http.MethodPut).Subrouter()
	postRouter := routes.Methods(http.MethodPost).Subrouter()
	deleteRouter := routes.Methods(http.MethodDelete).Subrouter()

	getRouter.HandleFunc("/moviment", handlers.GetMoviment)
	getRouter.HandleFunc("/moviment/{id}", handlers.GetMovimentById)
	putRouter.HandleFunc("/moviment", handlers.PutMoviment)
	postRouter.HandleFunc("/moviment", handlers.PostMoviment)
	deleteRouter.HandleFunc("/moviment", handlers.DeleteMoviment)

	log.Fatal(http.ListenAndServe(":8080", routes))

}
