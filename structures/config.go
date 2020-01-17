package structures

// Config is the structure for the config file
type Config struct {
	Port         int32 `json:"port"`
	PingInterval int32 `json:"pingInterval"`
}
