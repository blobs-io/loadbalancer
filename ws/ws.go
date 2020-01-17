package ws

const (
	// OpcodeHello is used for telling the client when to send keepalive pings
	OpcodeHello = 10
)

// WSHelloMessage is sent when the client connects to the server
type WSHelloMessage struct {
	Interval int32 `json:"interval"`
}

// WSMessage represents a message that is sent by one of the nodes to post statistics
type WSMessage struct {
	Op   uint8       `json:"op"`
	Type string      `json:"t"`
	Data interface{} `json:"d"`
}
