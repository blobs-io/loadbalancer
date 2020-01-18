package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/y21/loadbalancer/node"

	"github.com/y21/loadbalancer/structures"
)

// ParseConfig parses a config file
func ParseConfig(path string, config *structures.Config) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("could not open config file: %v", err)
		return
	}

	err = json.Unmarshal(file, config)
	if err != nil {
		fmt.Printf("could not parse config file: %v", err)
		return
	}
}

// ParseNodes parses a JSON file that holds all nodes
func ParseNodes(path string, nodes *[]node.Node) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("could not open nodes file: %v", err)
		return
	}

	err = json.Unmarshal(file, nodes)
	if err != nil {
		fmt.Printf("could not parse nodes: %v", err)
		return
	}
}
