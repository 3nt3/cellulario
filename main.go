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
	r.HandleFunc("/getFood", api.GetFood).Methods("GET")
	r.HandleFunc("/initCell", api.InitCell).Methods("POST")
	r.HandleFunc("/eat", api.Eat).Methods("POST")
	r.HandleFunc("/delall", api.Dellall).Methods("GET")
	r.HandleFunc("/updateSize/{cellId}", api.UpdateSize).Methods("POST")

	// In production use :80
	go log.Fatal(http.ListenAndServe(":8000", r))
}
