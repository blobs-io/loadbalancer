package webserver

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"

	"github.com/y21/loadbalancer/node"
	"github.com/y21/loadbalancer/structures"

	"github.com/gorilla/mux"
)

// Run runs the webserver
func Run(router *mux.Router, config *structures.Config, nodes *[]node.Node) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		node := node.GetOptimalNode(nodes, *config)
		http.Redirect(w, r, node.Host, 302)
	})

	router.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		var body bytes.Buffer
		body.WriteString("Total Nodes: " + strconv.Itoa(len(*nodes)) + "\n-------------\n")
		for _, node := range *nodes {
			body.WriteString(node.ToString(*config) + "\n")
		}
		fmt.Fprintf(w, body.String())
	})
}
