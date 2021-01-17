package types

// Message type will contain all structs
// and a timestamp for when the messsage was sent
type Message struct {
	Internet  Internet `json:"net"`
	Battery   Battery  `json:"power"`
	Timestamp int32    `json:"timestamp"`
}

// Battery type will be converted to JSON
// containing important information for this module
type Battery struct {
	Status   string `json:"status"`
	Health   string `json:"health"`
	Capacity int    `json:"capacity"`
	Temp     struct {
		Internal float32 `json:"int"`
		Ambient  float32 `json:"ext"`
	} `json:"temp"`
}

// Internet type will be converted to JSON
// containing important information for this module
type Internet struct {
	Device     string `json:"device"`
	ID         int    `json:"id"`
	IPAddress  string `json:"ipv4"`
	SubnetMask string `json:"mask"`
}
