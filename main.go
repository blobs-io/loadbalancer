package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/y21/loadbalancer/node"
	"github.com/y21/loadbalancer/structures"
	"github.com/y21/loadbalancer/utils"
	"github.com/y21/loadbalancer/webserver"
	"github.com/y21/loadbalancer/ws"
)

func main() {
	var config structures.Config
	var nodes = make([]node.Node, 16)
	var router = mux.NewRouter()

	// Configs
	utils.ParseConfig("./configs/config.json", &config)
	fmt.Println("[LOG] Config loaded")
	utils.ParseNodes("./configs/nodes.json", &nodes)
	fmt.Printf("[LOG] Loaded %d nodes\n", len(nodes))

	// Webserver & WebSocket server
	webserver.Run(router, &config, &nodes)
	ws.Handle(router, &config, &nodes)

	// Ping nodes
	go node.PingAllNodes(&nodes)

	err := http.ListenAndServe(":"+strconv.Itoa((int)(config.Port)), router)
	if err != nil {
		fmt.Println(err)
	}
}