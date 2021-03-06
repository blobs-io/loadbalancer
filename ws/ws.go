package ws

const (
	// OpcodeHello is used for telling the client when to send keepalive pings
	OpcodeHello = 10
)

// UpdateData is the data that is sent every N ms by every node
type UpdateData struct {
	CPU         uint8  `json:"cpu"`
	Mem         uint8  `json:"mem"`
	AccessToken string `json:"accessToken"`
}

// Message represents a message that is sent by one of the nodes to post statistics
type Message struct {
	Op   uint8  `json:"op"`
	Type string `json:"t"`
	Data string `json:"d"`
}
