package api

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var cells []cell
var foodItems []food

func SpawnFood(w http.ResponseWriter, r *http.Request) {
	foodItems = []food{}

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

		var pos []int
		for i := 0; i < 2; i++ {
			s := rand.NewSource(time.Now().UnixNano())
			r := rand.New(s)
			pos = append(pos, r.Intn(1000))
		}

		var newItem food
		value := values[r1.Intn(99)]
		newItem = food{len(foodItems), pos, value}

		foodItems = append(foodItems, newItem)
	}

	log.Println(foodItems)

	_ = json.NewEncoder(w).Encode(foodItems)
}

func GetFood(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(foodItems)
}


func GetCells(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(cells)
}
