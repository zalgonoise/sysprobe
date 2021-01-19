package network

// Network type will be converted to JSON
// containing important information for this module
type Network struct {
	System System   `json:"sys"`
	Ping   PingScan `json:"ping"`
}
