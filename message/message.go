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
	BatteryOpt  bool
	PingOpt     bool
	PortScanOpt bool
	HelpOpt     bool
	BatteryPath string
	IPDevice    string
	PingAddr    string
	SlowPing    bool
}

// Response struct will contain all structs
// and a timestamp for when the messsage was sent
type Response struct {
	Network   net.Network `json:"net"`
	Battery   bat.Battery `json:"power"`
	Timestamp int32       `json:"timestamp"`
}

// New method will gather the input flags and build the
// instructions for the request. Serves as a handler for
// default values as well
func (r *Request) New() *Request {
	// Triggers and toggles
	batteryOpt := flag.Bool("b", false, "Option to run a battery scan")
	pingOpt := flag.Bool("p", false, "Option to run a ping scan. Can take in the parameters -net and -slow.")
	portScanOpt := flag.Bool("P", false, "Option to perform a port scan on the alive IP addresses")
	helpOpt := flag.Bool("help", false, "Displays the usage for this binary")

	batteryPath := flag.String("bat", "battery", "The default location for the battery uevent file, in /sys/class/power_supply/")
	ipDevice := flag.String("net", "wlan0", "The default network device to retrieve IP-related information, with the `ip` command")
	pingAddr := flag.String("ping", "192.168.0.0/24", "The network or subnet address to issue ping events, similar to the *nix `ping` command (but in Go)")
	slowPing := flag.Bool("slow", false, "Skip goroutines - perform single-threaded actions only")

	flag.Parse()

	r.HelpOpt = *helpOpt
	r.BatteryOpt = *batteryOpt
	r.PingOpt = *pingOpt
	r.PortScanOpt = *portScanOpt
	r.BatteryPath = *batteryPath
	r.IPDevice = *ipDevice
	r.PingAddr = *pingAddr
	r.SlowPing = *slowPing

	return r
}

// New method - it builds a new Response struct,
// containing the data in Internet, Battery, as well
// as the current Unix timestamp
//func (m *Response) New(batRef, netRef, pingRef string, slowPing, portScanOpt bool) *Message {
func (m *Response) New(r Request) *Response {

	b := &bat.Battery{}

	if r.BatteryOpt != false {
		b = b.Get(r.BatteryPath)
	}

	m.Battery = *b

	if r.PingOpt != true {
		m.Network = net.Network{}
	} else {
		m.Network = *m.Network.Build(r.IPDevice, r.PingAddr, r.SlowPing, r.PortScanOpt)
	}

	m.Timestamp = int32(time.Now().Unix())

	return m
}

// JSON method - converts the Response struct into
// a JSON-encoded byte array
func (m *Response) JSON() []byte {
	json, err := json.Marshal(m)
	utils.Check(err)
	return json
}
