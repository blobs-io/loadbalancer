package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/y21/loadbalancer/structures"
)

var config structures.Config
var nodes []structures.Node = make([]structures.Node, 16)

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

	http.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		var body bytes.Buffer
		for _, node := range nodes {
			body.WriteString(node.ToString() + "\n")
		}
		fmt.Fprintf(w, body.String())
	})

	fmt.Printf("Webserver running on port %d\n", config.Port)
	http.ListenAndServe(":"+strconv.Itoa((int)(config.Port)), nil)
}
