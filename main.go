package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	//"cellulario/structs"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true;
	},
}

//var initial  = true

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Panic(err)
	}

	msgs := make(chan string)

	go read(msgs, conn)

	for {
		foo := <-msgs
		log.Println(foo)
		conn.WriteMessage(1, []byte(foo))
	}
}

func read(messages chan string, conn *websocket.Conn) {
	for {
		//fmt.Println(conn)
		_, foo, _ := conn.ReadMessage()
		messages <- string(foo)
	}
}
