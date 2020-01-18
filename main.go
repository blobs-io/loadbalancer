package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/y21/loadbalancer/structures"
	"github.com/y21/loadbalancer/ws"
)

var config structures.Config
var nodes []structures.Node = make([]structures.Node, 16)
var wsUpgrader = websocket.Upgrader{
	EnableCompression: true,
	WriteBufferSize:   1024,
	ReadBufferSize:    1024,
}

func main() {
	file, err := ioutil.ReadFile("./configs/config.json")
	if err != nil {
		fmt.Printf("could not read config file: %s", err)
		os.Exit(1)
	}
	json.Unmarshal(file, &config)
	file, err = ioutil.ReadFile("./configs/nodes.json")
	if err != nil {
		fmt.Printf("could not read nodes file: %s", err)
		os.Exit(1)
	}
	json.Unmarshal(file, &nodes)

	go structures.PingAllNodes(&nodes)

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		node := structures.GetOptimalNode(&nodes, config)
		http.Redirect(w, r, node.Host, 302)
	})

	router.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		var body bytes.Buffer
		body.WriteString("Total Nodes: " + strconv.Itoa(len(nodes)) + "\n-------------\n")
		for _, node := range nodes {
			body.WriteString(node.ToString(config) + "\n")
		}
		fmt.Fprintf(w, body.String())
	})

	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		connection, err := wsUpgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Printf("ws connection error: %v\n", err)
			return
		}

		connection.WriteJSON(ws.Message{
			Op:   ws.OpcodeHello,
			Data: strconv.Itoa((int)(config.PingInterval)),
		})

		for {
			var msg ws.Message
			err := connection.ReadJSON(&msg)
			if err != nil {
				fmt.Printf("ws message error: %v\n", err)
				break
			}

			if msg.Type == "stats" {
				var data ws.UpdateData
				json.Unmarshal([]byte(msg.Data), &data)
				for i := range nodes {
					node := &(nodes[i])
					if node.AccessToken == data.AccessToken {
						node.CPU = data.CPU
						node.Mem = data.Mem
						node.Available = true
					}
				}
			}
		}
	})

	fmt.Printf("Webserver running on port %d\n", config.Port)
	http.ListenAndServe(":"+strconv.Itoa((int)(config.Port)), router)
}
