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

func WsConv(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: ", err)
		return
	}
	Clients[conn] = true
}

func WsDealer() {
	for {
		time.Sleep(time.Second)
		for client := range Clients {
			fmt.Println(time.Now())
			Err = client.WriteJSON(map[string]interface{}{"board": HtmlBoard(), "messages": Messages})
			if Err != nil {
				fmt.Println("Websocket error:", Err)
				Err = client.Close()
				Eh(Err)
				delete(Clients, client)
			}
		}
	}
}
