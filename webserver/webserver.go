package webserver

import (
	"encoding/json"
	"net/http"

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

	router.HandleFunc("/api/stats", func(w http.ResponseWriter, r *http.Request) {
		body := make([]node.Node, 0)
		for _, node := range *nodes {
			node.AccessToken = "" // don't expose access token
			body = append(body, node)
		}

		json.NewEncoder(w).Encode(body)
	})
}
