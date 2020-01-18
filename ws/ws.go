package ws

const (
	// OpcodeHello is used for telling the client when to send keepalive pings
	OpcodeHello = 10
)

// Message represents a message that is sent by one of the nodes to post statistics
type Message struct {
	Op   uint8  `json:"op"`
	Type string `json:"t"`
	Data string `json:"d"`
}
