package message

import (
	"encoding/json"
	"time"

	bat "github.com/ZalgoNoise/sysprobe/battery"
	net "github.com/ZalgoNoise/sysprobe/internet"
	"github.com/ZalgoNoise/sysprobe/utils"
)

// Message type will contain all structs
// and a timestamp for when the messsage was sent
type Message struct {
	Internet  net.Internet `json:"net"`
	Battery   bat.Battery  `json:"power"`
	Timestamp int32        `json:"timestamp"`
}

// New function - it builds a new Message struct,
// containing the data in Internet, Battery, as well
// as the current Unix timestamp
func (m *Message) New(batRef, netRef string) *Message {

	b := &bat.Battery{}
	i := &net.Internet{}

	b = b.Get(batRef)
	i = i.Get(netRef)

	m.Battery = *b
	m.Internet = *i
	m.Timestamp = int32(time.Now().Unix())

	return m
}

// JSON method - converts the Message struct into
// a JSON-encoded byte array
func (m *Message) JSON() []byte {
	json, err := json.Marshal(m)
	utils.Check(err)
	return json
}
