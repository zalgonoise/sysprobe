package message

import (
	"encoding/json"
	"flag"
	"time"

	bat "github.com/ZalgoNoise/sysprobe/battery"
	net "github.com/ZalgoNoise/sysprobe/network"
	"github.com/ZalgoNoise/sysprobe/utils"
)

// Request struct will contain the flags necessary to
// execute the needed jobs and to return a Message object
type Request struct {
	BatteryPath string
	IPDevice    string
	PingAddr    string
	SlowPing    bool
	PortScanOpt bool
}

// Response struct will contain all structs
// and a timestamp for when the messsage was sent
type Response struct {
	Network   net.Network `json:"net"`
	Battery   bat.Battery `json:"power"`
	Timestamp int32       `json:"timestamp"`
}

// Create method will gather the input flags and build the
// instructions for the request. Serves as a handler for
// default values as well
func (r *Request) Create() *Request {
	batteryPath := flag.String("bat", "battery", "the default location for the battery uevent file, in /sys/class/power_supply/")
	ipDevice := flag.String("net", "wlan0", "the default network device to retrieve IP-related information, with the `ip` command")
	pingAddr := flag.String("ping", "192.168.0.0/24", "the network or subnet address to issue ping events, similar to the *nix `ping` command (but in Go)")
	slowPing := flag.Bool("slow", false, "skip goroutines - perform single-threaded actions only")
	portScanOpt := flag.Bool("port", false, "perform a port scan on the alive IP addresses")

	flag.Parse()

	r.BatteryPath = *batteryPath
	r.IPDevice = *ipDevice
	r.PingAddr = *pingAddr
	r.SlowPing = *slowPing
	r.PortScanOpt = *portScanOpt

	return r
}

// New function - it builds a new Message struct,
// containing the data in Internet, Battery, as well
// as the current Unix timestamp
//func (m *Response) New(batRef, netRef, pingRef string, slowPing, portScanOpt bool) *Message {
func (m *Response) New(r Request) *Response {

	b := &bat.Battery{}
	b = b.Get(r.BatteryPath)
	m.Battery = *b

	m.Network = *m.Network.Build(r.IPDevice, r.PingAddr, r.SlowPing, r.PortScanOpt)

	m.Timestamp = int32(time.Now().Unix())

	return m
}

// JSON method - converts the Message struct into
// a JSON-encoded byte array
func (m *Response) JSON() []byte {
	json, err := json.Marshal(m)
	utils.Check(err)
	return json
}
