package structures

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

// Node represents an edge that hosts an instance of an application
type Node struct {
	Host           string  `json:"host"`
	Ping           uint8   `json:"ping"`
	CPU            uint8   `json:"cpu"`
	Available      bool    `json:"available"`
	LastStatusCode int     `json:"lastStatusCode"`
	AccessToken    string  `json:"accessToken"`
	Mem            uint8   `json:"mem"`
	LastScore      float32 `json:"lastScore"`
}

// GetScore returns the score for a node (lower = better)
func (n *Node) GetScore(config Config) float32 {
	score := ((float32)(n.CPU) * config.Weights.CPU) +
		((float32)(n.Mem) * config.Weights.Mem) +
		((float32)(n.Ping) * config.Weights.Ping)
	n.LastScore = score
	return score
}

// PingNode pings the node and updates stats
func (n *Node) PingNode() {
	timestampBefore := time.Now().UnixNano()
	response, err := http.Get(n.Host)
	if err != nil {
		fmt.Printf("[%s] could not ping node: %v", n.Host, err)
		n.Available = false
		n.LastStatusCode = 500
		return
	}

	n.Available = true
	n.LastStatusCode = response.StatusCode
	n.Ping = (uint8)((time.Now().UnixNano() - timestampBefore) / 1000000)
}

// ToString returns a string representing the node
func (n *Node) ToString(config Config) string {
	var result bytes.Buffer
	var available string
	if n.Available == true {
		available = "yes"
	} else {
		available = "no"
	}
	result.WriteString(fmt.Sprintf("[%s] Ping: %dms, CPU: %d, RAM: %d, Score: %f, Available: %s", n.Host, n.Ping, n.CPU, n.Mem, n.GetScore(config), available))
	return result.String()
}

// PingAllNodes pings all nodes and updates stats
func PingAllNodes(n *[]Node) {
	for {
		for i := range *n {
			(*n)[i].PingNode()
		}

		time.Sleep(time.Second * 5)
	}
}

// GetOptimalNode gets a node that has the lowest score
func GetOptimalNode(n *[]Node, config Config) Node {
	var node Node

	for i := range *n {
		if i == 0 || node.LastScore > (*n)[i].GetScore(config) {
			node = (*n)[i]
		}
	}

	return node
}
