package api

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

var rooms []room
var cells []cell
var foodItems []food

func SpawnFood(w http.ResponseWriter, r *http.Request) {
	valuesSrc := []int{5, 10, 20}
	rarities := []int{50, 40, 10}

	var values []int
	for i := 0; i < 100; i++ {
		if len(values) < rarities[0] {
			values = append(values, valuesSrc[0])
		} else if len(values) < rarities[1]+rarities[0] {
			values = append(values, valuesSrc[1])
		} else {
			values = append(values, valuesSrc[2])
		}
	}
	for i := 0; i < 100; i++ {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)

		var newItem food
		value := values[r1.Intn(100)]

		foodItems = append(foodItems, newItem)
	}

	newItem := food{}
	foodItems = append(foodItems, newItem)
}

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(1)
}

func GetCells(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(cells)
}
