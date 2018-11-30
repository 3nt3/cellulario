package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var cells []cell
var foodItems []food

func SpawnFood(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	foodItems = []food{}

	valuesSrc := []int{5, 10, 20}
	rarities := []int{50, 40, 10}

	var values []int
	for i := 0; i < 10; i++ {
		if len(values) < rarities[0] {
			values = append(values, valuesSrc[0])
		} else if len(values) < rarities[1]+rarities[0] {
			values = append(values, valuesSrc[1])
		} else {
			values = append(values, valuesSrc[2])
		}
	}
	for i := 0; i < 10; i++ {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)

		var pos []int
		for i := 0; i < 2; i++ {
			s := rand.NewSource(time.Now().UnixNano())
			r := rand.New(s)
			pos = append(pos, r.Intn(1000))
		}

		var newItem food
		value := values[r1.Intn(19)]
		newItem = food{len(foodItems), pos, value}

		foodItems = append(foodItems, newItem)
	}

	log.Println(foodItems)

	_ = json.NewEncoder(w).Encode(foodItems)
}

func InitCell(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var name string
	_ = json.NewDecoder(r.Body).Decode(&name)

	fmt.Printf("New Cell: %s\n", name)

	var pos []int
	for i := 0; i < 2; i++ {
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)
		pos = append(pos, r.Intn(1000))
	}

	newCell := cell{len(cells), name, true, 10, 0, []cell{}, pos}
	cells = append(cells, newCell)

	_ = json.NewEncoder(w).Encode(newCell.Id)
}

func UpdateSize(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var size int
	_ = json.NewDecoder(r.Body).Decode(&size)

	id, _ := strconv.Atoi(mux.Vars(r)["cellId"])

	cells[id].Size = size
}

// Just for testing
func Dellall(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	cells = []cell{}
}

func Eat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var cellId int
	_ = json.NewDecoder(r.Body).Decode(&cellId)

	cells[cellId].Alive = false

}

func GetFood(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	_ = json.NewEncoder(w).Encode(foodItems)
}

func GetCells(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	_ = json.NewEncoder(w).Encode(cells)
}
