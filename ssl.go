package main

import (
	"fmt"
	"net/http"
	"reflect"

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
	err = conn.WriteJSON(map[string]interface{}{"board": jsonBoard(), "messages": messages})
	eh(err)
	var msg webReq
	for {
		err := conn.ReadJSON(&msg)
		if err != nil {
			break
		}
		if !reflect.DeepEqual(jsonBoard(), msg.Board) || !reflect.DeepEqual(messages, msg.Messages) {
			fmt.Println(jsonBoard(), messages, msg)
			conn.WriteJSON(map[string]interface{}{"board": jsonBoard(), "messages": messages})
		} else {
			fmt.Println(true)
		}
	}
}
