// Package net is the main module to handle all network-related
// metadata collection actions.
//
// It is capable of running each module individually (except port scan,
// which requires a ping scan), where the absense of data will simply
// return an empty object corresponding to that module (PingScan, HostScan,
// or the whole net object).
//
// Its main focus is the Network.Build() method, responsible for queueing
// and orchestrating all the network-related tasks, but it also contains the
// PingScan.ExpandCIDR() method, which is responsible for breaking down
// subnets into slices of strings, for the respective hosts in it.
//
package net

import (
	"net"
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

	ipList, err := ping.ExpandCIDR(pingRef)
	utils.Check(err)

	if slowPing != true {
		wg.Add(1)
		go ping.Burst(&wg, ipList)
	} else {
		wg.Add(1)
		go ping.Paced(&wg, ipList)
	}
	wg.Wait()

	if portScanOpt != false {
		alive := ping.Get()
		wg.Add(1)

		go port.Create(&wg, alive, 9999)
		//go port.Create(&wg, alive, 1024)
		//go port.Create(&wg, alive, 49152)
	}

	wg.Add(1)
	go sys.Get(&wg, netRef)
	wg.Wait()

	if portScanOpt != false {
		n = &Network{System: *sys, Ping: *ping, Ports: port.Results}

	} else {
		n = &Network{System: *sys, Ping: *ping}
	}

	return n
}

// ExpandCIDR method - expands (simple) CIDR addresses
// taking the example from https://gist.github.com/kotakanbe/d3059af990252ba89a82
// Fixed issue with /32 CIDR addresses
func (p *PingScan) ExpandCIDR(target string) ([]string, error) {
	// set target from the object's Target key
	p.Target = target

	// parse the CIDR in the input address
	ip, ipnet, err := net.ParseCIDR(target)
	if err != nil {
		return nil, err
	}

	var ips []string

	// iterate through the network mask, confirming if the input IP is present
	// in the network, and incrementing the IP address by 1. If all checks
	// evaluate to true, then add the IP address to the slice of strings
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); addrIncrement(ip) {
		ips = append(ips, ip.String())
	}
	// remove network address (e.g. 192.168.0.0)
	// removing the broadcast address only if the slice contains 3 or more elements
	// this will correctly evaluate the addresses for /30, /31 and /32 CIDRs:
	// [192.168.0.1 192.168.0.2], [192.168.0.1] and [], respectively
	if len(ips) <= 3 {
		return ips[1:], nil
	} else {
		return ips[1 : len(ips)-1], nil
	}

}

func addrIncrement(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
