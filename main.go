package main

import (
	"cellulario/api"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	// Funx
	r.HandleFunc("/getCells", api.GetCells).Methods("GET")
	r.HandleFunc("/spawnFood", api.SpawnFood).Methods("GET")

	// In production use :80
	go log.Fatal(http.ListenAndServe(":8000", r))
}
