package api

import (
	"encoding/json"
	"net/http"
)

var cells []cell

func GetCells(w http.ResponseWriter, r *http.Request) {
	json.Encoder{}
}