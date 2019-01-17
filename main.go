package main

import (
	"cellulario/funcs"
	"cellulario/structs"
	"cellulario/vars"
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Panic(err)
	}

	// establish channel
	msgs := make(chan structs.WsMsg)

	// goroutines
	go read(msgs, conn)
	go checkState(conn)

	// main loop
	for {
		cresp := &structs.WsMsg{}
		allDead := false
		for _, item := range vars.State.Food {
			if item.Alive {
				break
			} else {
				allDead = true
			}
		}

		if len(vars.State.Food) == 0 {
			vars.State.Food = funcs.SpawnFood()
		} else if allDead {
			// check if the food needs to be respawned
			vars.State.Food = funcs.SpawnFood()
		} else {
			// wait for user input
			creq := <-msgs
			cresp.Type = creq.Type
			if creq.Type == "" {
				continue
			} else {
				switch creq.Type {
				case "changePos":
					var pos []int
					for _, item := range creq.Data["pos"].([]interface{}) {
						pos = append(pos, int(item.(float64)))
					}
					id := int(creq.Data["id"].(float64))
					cresp.Data = toInterface(funcs.ChangePos(id, pos))

				case "eat":
					mealType := creq.Data["type"].(string)
					mealId := int(creq.Data["mealId"].(float64))
					id := int(creq.Data["id"].(float64))
					cresp.Data = toInterface(funcs.Eat(id, mealId, mealType))

				case "initCell":
					name := creq.Data["name"].(string)
					cresp.Data = toInterface(funcs.InitCell(name))

				case "delall":
					funcs.Delall()
				}
			}
		}
		conn.WriteJSON(cresp)
	}
}

func read(messages chan structs.WsMsg, conn *websocket.Conn) {
	for {
		creq := &structs.WsMsg{}
		_ = conn.ReadJSON(creq)
		messages <- *creq
	}
}

func checkState(conn *websocket.Conn) {
	var oldState structs.GameState
	for {
		if !cmp.Equal(vars.State, oldState) {
			cresp := structs.WsMsg{"state", toInterface(vars.State)}
			conn.WriteJSON(cresp)
		}
		oldState = vars.State
	}
}

func toInterface(a interface{}) map[string]interface{} {
	foo, _ := json.Marshal(a)
	var bar interface{}
	_ = json.Unmarshal(foo, &bar)
	return bar.(map[string]interface{})
}
