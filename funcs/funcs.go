package funcs

import (
	"cellulario/structs"
	"cellulario/vars"
	"log"
	"math/rand"
	"time"
)

func SpawnFood() {
	valuesSrc := []int{1, 3, 5}
	rarities := []int{5, 4, 1}

	vars.State.Food = []structs.Food{}

	var values []int
	items := 50

	var foodItems []structs.Food

	for i := 0; i < items; i++ {
		if len(values) < rarities[0] {
			values = append(values, valuesSrc[0])
		} else if len(values) < rarities[1]+rarities[0] {
			values = append(values, valuesSrc[1])
		} else {
			values = append(values, valuesSrc[2])
		}
	}
	for i := 0; i < len(values); i++ {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)

		var pos []int
		for i := 0; i < 2; i++ {
			s := rand.NewSource(time.Now().UnixNano())
			r := rand.New(s)
			pos = append(pos, r.Intn(2000)-1000)
		}

		var newItem structs.Food
		value := values[r1.Intn(len(values))]
		newItem = structs.Food{len(vars.State.Food), pos, value, true}

		foodItems = append(foodItems, newItem)
	}

	vars.State.Food = foodItems
}

func InitCell(name string) structs.Cell {
	var NewCell structs.Cell
	var pos []int

	for i := 0; i < 2; i++ {
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)
		pos = append(pos, r.Intn(2000)-1000)
	}

	NewCell = structs.Cell{len(vars.State.Cells), name, true, 30, 0, []structs.Cell{}, pos}

	vars.State.Cells = append(vars.State.Cells, NewCell)
	log.Printf("New Cell with id %d called  %s at %v", NewCell.Id, NewCell.Name, NewCell.Pos)

	return NewCell
}

func ChangeSize(id int, size int) []structs.Cell {
	vars.State.Cells[id].Size += size
	log.Printf("The size of id %d ('%s') was changed by %d. The current size is %d", id, vars.State.Cells[id].Name, size, vars.State.Cells[id].Size)
	return vars.State.Cells
}

func Delall() {
	vars.State.Cells = []structs.Cell{}
	vars.State.Food = []structs.Food{}
}

func Eat(id int, mealId int, mealType string) []structs.Cell {
	if mealType == "food" {
		meal := &vars.State.Food[id]
		player := &vars.State.Cells[id]

		size := player.Size
		ChangeSize(id, meal.Value)
		meal.Alive = false

		log.Printf("The player with id %d (%s) of size %d has eaten foodItem %d of valu e %d. The new size of player %d is now %d", player.Id, player.Name, size, meal.Id, meal.Value, player.Id, player.Size)
	} else {
		meal := &vars.State.Cells[mealId]
		player := &vars.State.Cells[id]

		sizes := []int{player.Size, meal.Size}
		ChangeSize(id, meal.Size)
		meal.Alive = false
		player.Meals = append(vars.State.Cells[id].Meals, *meal)
		player.Kills += 1

		log.Printf("The player with id %d (%s) of size %d has eaten player %d (%s) of size %d. The new size of player %d is now %d", player.Id, player.Name, sizes[0], meal.Id, meal.Name, sizes[1], player.Id, player.Size)
	}
	return vars.State.Cells
}

func ChangePos(id int, pos []int) {
	vars.State.Cells[id].Pos = pos
}
