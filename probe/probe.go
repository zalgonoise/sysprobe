// Package probe is a module to handle this binary's request and response
// messages, and the output's JSON encoding (possibly just for debugging or
// testing, considering the implementation of rpc between client/server).
//
package probe

import (
	"encoding/json"
	"flag"
	"os"
	"time"

	"github.com/ZalgoNoise/sysprobe/bat"
	"github.com/ZalgoNoise/sysprobe/net"
)

// Prober struct will contain the whole process of execution,
// in one object.
type Prober struct {

	// Toggles will contain boolean values for the enabled
	// modules in the pinger. These are read and their respective
	// processes are executed as per this configuration
	Toggles Toggles `json:"mods_enabled"`

	// Net will contain network-related modules and settings,
	// similar to Toggles, but isn't limited to boolean values.
	// The address to your ping queries is defined here.
	Net Net `json:"net_mods"`

	// BatteryPath is the folder in /sys/class/power_supply/,
	// where the uevent file is located
	BatteryPath string `json:"battery_path"`

	// Help will set to true or false whether the -help flag is
	// passed
	Help bool `json:"help_option"`

	// Exec will contain a set of customizable functions which
	// can be defined upon different implementations. These are
	// not producing any output by default until defined.
	Exec Exec `json:"execute"`

	// Debug will set to true or false whether the -debug flag is
	// passed.
	Debug bool `json:"debug"`

	// Response will contain the object with the results from the
	// probe scan
	Response Response `json:"response"`

	// JSON will contain an array of bytes after Response is
	// converted to JSON
	JSON []byte `json:"json_output"`
}

// Response struct will contain all produced objects
// and a timestamp for when the object was created
type Response struct {
	Network   net.Network `json:"net"`
	Battery   bat.Battery `json:"power"`
	Timestamp int32       `json:"timestamp"`
}

// Toggles struct is a simple object containing boolean values
// for the modules activated upon runtime, as per flags enabled.
// Services are either executed or not depending on how these are set
type Toggles struct {
	BatteryOpt  bool `json:"battery"`
	PingOpt     bool `json:"ping"`
	PortScanOpt bool `json:"ports"`
}

// Net struct will contain options and parameters to be used on
// network-related operations
type Net struct {
	IPDevice string `json:"network_device"`
	PingAddr string `json:"target_address"`
	SlowPing bool   `json:"slow_ping"`
}

// Exec struct will hold OnStart(), OnRun() and OnDone() functions
// which can be defined by the user. This can be be used to collect
// different metadata as the tool expands
type Exec struct {
	OnStart func(*Toggles)
	OnRun   func(*Response)
	OnDone  func(*Prober)
}

// OnStart(), OnRun() and OnDone() placeholder functions
func (p *Prober) onStart() {
	hook := p.Exec.OnStart

	if hook != nil {
		hook(&p.Toggles)
	}
}

func (p *Prober) onRun() {
	hook := p.Exec.OnRun

	if hook != nil {
		hook(&p.Response)
	}
}

func (p *Prober) onDone() {
	hook := p.Exec.OnDone

	if hook != nil {
		hook(p)
	}
}

// New function will generate and return a new Prober
func New() *Prober {
	p := build()
	defer p.onStart()
	return p
}

// Build method will gather the input flags and build the
// instructions for the request. Serves as a handler for
// default values as well
func build() *Prober {
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

	if *helpOpt != false {
		flag.Usage()
		os.Exit(0)
	}

	t := Toggles{
		BatteryOpt:  *batteryOpt,
		PingOpt:     *pingOpt,
		PortScanOpt: *portScanOpt,
	}

	n := Net{
		IPDevice: *ipDevice,
		PingAddr: *pingAddr,
		SlowPing: *slowPing,
	}

	r := &Prober{
		Toggles:     t,
		Net:         n,
		BatteryPath: *batteryPath,
		Help:        *helpOpt,
		Debug:       false,
	}

	return r
}

// Run method runs an instance of Prober, fetching sources
// containing the data in Internet, Battery, as well
// as the current Unix timestamp
func (p *Prober) Run() *Prober {
	defer p.onRun()

	b := &bat.Battery{}
	n := &net.Network{}

	if p.Toggles.BatteryOpt != false {
		b = b.Get(p.BatteryPath)
	}

	if p.Toggles.PingOpt != false {
		n = n.Build(p.Net.IPDevice, p.Net.PingAddr, p.Net.SlowPing, p.Toggles.PortScanOpt)
	}

	//fmt.Println(b)
	//fmt.Println(n)

	r := &Response{
		Network:   *n,
		Battery:   *b,
		Timestamp: int32(time.Now().Unix()),
	}

	p.Response = *r

	p.Done()
	return p
}

// Done method converts the Response struct into
// a JSON-encoded byte array
func (p *Prober) Done() {
	defer p.onDone()
	json, err := json.Marshal(p.Response)
	if err != nil {
		panic(err)
	}
	p.JSON = json
}
