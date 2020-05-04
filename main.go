package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var router = mux.NewRouter()

func display(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test server page")
}

func setupRoutes() {
	router.HandleFunc("/", display).Methods("POST")
}

func main() {
	log.Info("Starting test API")
	setupRoutes()
	http.ListenAndServe(":8000", router)
}
