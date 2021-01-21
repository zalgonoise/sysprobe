package network

import (
	"math"
	"strconv"
	"strings"
	"sync"

	"github.com/ZalgoNoise/sysprobe/utils"
)

// Network type will be converted to JSON
// containing important information for this module
type Network struct {
	System System     `json:"sys"`
	Ping   PingScan   `json:"ping"`
	Ports  []HostScan `json:"ports"`
}

// Build method - issues network-related microprocesses
// which builds up to the Network struct
func (n *Network) Build(netRef, pingRef string, slowPing, portScanOpt bool) *Network {
	var wg sync.WaitGroup

	ping := &PingScan{}
	sys := &System{}
	port := &ScanResults{}
	ipList := ping.ExpandCIDR(pingRef)

	if slowPing != true {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			ping.Burst(ipList)

		}(&wg)

	} else {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			ping.Paced(ipList)

		}(&wg)

	}
	wg.Wait()

	if portScanOpt != false {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			alive := ping.Get()

			port.Create(alive, 1024)
		}(&wg)
	}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		sys.Get(netRef)

	}(&wg)

	wg.Wait()

	if portScanOpt != false {
		n = &Network{System: *sys, Ping: *ping, Ports: port.Results}

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
