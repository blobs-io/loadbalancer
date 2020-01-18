package structures

// WeightData stores weights
type WeightData struct {
	CPU  float32 `json:"cpu"`
	Mem  float32 `json:"mem"`
	Ping float32 `json:"ping"`
}

// Config is the structure for the config file
type Config struct {
	Port         int32      `json:"port"`
	PingInterval int32      `json:"pingInterval"`
	Weights      WeightData `json:"weights"`
}
