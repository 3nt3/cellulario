package main

import (
	"cellulario/funcs"
	"cellulario/structs"
	"cellulario/vars"
	"github.com/google/go-cmp/cmp"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true;
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
	msgs := make(chan string)
	var oldState structs.GameState
	initial := true

	go read(msgs, conn)

	for {
		// check if it is the first iteration
		if initial {
			if len(vars.State.Food) == 0 {
				vars.State.Food = funcs.SpawnFood()
			}
			oldState = vars.State
			initial = false
		} else {
			// check if all the food is eaten
			allDead := false

			for _, item := range vars.State.Food {
				if item.Alive {
					break
				} else {
					allDead = true
				}
			}

			if allDead {
				vars.State.Food = funcs.SpawnFood()
			}

			if !cmp.Equal(vars.State, oldState) {
				conn.WriteJSON(vars.State)
			}
		}
	}
}

func read(messages chan string, conn *websocket.Conn) {
	for {
		//fmt.Println(conn)
		_, foo, _ := conn.ReadMessage()
		messages <- string(foo)
	}
}
