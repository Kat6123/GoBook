package main

import "github.com/gorilla/mux"

func setRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", printHandler)

	return router
}
