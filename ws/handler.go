package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/y21/loadbalancer/node"
	"github.com/y21/loadbalancer/structures"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{
	EnableCompression: true,
	WriteBufferSize:   1024,
	ReadBufferSize:    1024,
}

// Handle handles the websocket server
func Handle(router *mux.Router, config *structures.Config, nodes *[]node.Node) {
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		connection, err := wsUpgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Printf("ws connection error: %v\n", err)
			return
		}

		connection.WriteJSON(Message{
			Op:   OpcodeHello,
			Data: strconv.Itoa((int)(config.PingInterval)),
		})

		for {
			var msg Message
			err := connection.ReadJSON(&msg)
			if err != nil {
				fmt.Printf("ws message error: %v\n", err)
				break
			}

			if msg.Type == "stats" {
				var data UpdateData
				err := json.Unmarshal([]byte(msg.Data), &data)
				if err != nil {
					fmt.Printf("ws data parsing error: %v\n", err)
					break
				}

				for i := range *nodes {
					node := &(*nodes)[i]
					if node.AccessToken == data.AccessToken {
						node.CPU = data.CPU
						node.Mem = data.Mem
						node.Available = true
					}
				}
			}
		}
	})
}
