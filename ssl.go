package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader websocket.Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return r.Header.Get("Origin") != ""
	},
}

func wsConv(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: ", err)
		return
	}
	clients[conn] = true
}

func wsDealer() {
	var msg string
	for {
		time.Sleep(time.Second)
		//newMsg := <-msgChan
		//newMove := <-moveChan
		for client := range clients {
			fmt.Println(msg)
			err = client.WriteJSON(map[string]interface{}{"board": htmlBoard(), "messages": messages})
			if err != nil {
				fmt.Println("Websocket error:", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
