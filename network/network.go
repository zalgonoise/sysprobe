package network

import (
	"math"
	"strconv"
	"strings"

	"github.com/ZalgoNoise/sysprobe/utils"
)

// Network type will be converted to JSON
// containing important information for this module
type Network struct {
	System System      `json:"sys"`
	Ping   PingScan    `json:"ping"`
	Ports  ScanResults `json:"ports"`
}

// Build method - issues network-related microprocesses
// which builds up to the Network struct
func (n *Network) Build(netRef, pingRef string, slowPing, portScanOpt bool) *Network {

	sys := &System{}
	sys.Get(netRef)

	ping := &PingScan{}
	ipList := ping.ExpandCIDR(pingRef)

	if slowPing != true {
		ping.Burst(ipList)
	} else {
		ping.Paced(ipList)
	}

	if portScanOpt != false {
		alive := ping.Get()

		port := &ScanResults{}
		port.Create(alive, 1024)
		n = &Network{System: *sys, Ping: *ping, Ports: *port}

	} else {
		n = &Network{System: *sys, Ping: *ping}
	}

	return n
}

// ExpandCIDR method - expands (simple) CIDR addresses
// currently supporting 0/24 addresses and above,
// listing the number of addresses starting at
// 255 and downwards.
// Work needs to be done with this method in the future
func (p *PingScan) ExpandCIDR(target string) []string {
	p.Target = target

	ip := utils.Splitter(p.Target, "/", 0)

	cidr := utils.Splitter(p.Target, "/", 1)

	input, _ := strconv.Atoi(cidr)
	exp := 32 - input

	const base float64 = 2
	calc := math.Pow(base, float64(exp))
	output := int32(calc) - 2

	ip = strings.TrimRight(ip, "0")

	var res []string

	for i := 255; i > 255-int(output); i-- {

		curIP := ip + strconv.Itoa(i)
		res = append(res, curIP)
	}

	return res
}
