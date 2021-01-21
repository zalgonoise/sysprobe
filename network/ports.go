package network

import (
	"net"
	"strconv"
	"time"
)

// ScanResults struct will hold a list of HostScans.
// This is the placeholder for all host queries
type ScanResults struct {
	Results []HostScan `json:"results"`
}

// HostScan struct will contain the probe results for
// a single host. It will hold the target IP address,
// the protocol used, and the open ports in a slice of ints.
type HostScan struct {
	Target   string `json:"target"`
	Protocol string `json:"proto"`
	Ports    []int  `json:"ports"`
}

// PortScan struct will contain the results for a single
// port scan. The Port key will contain an int with the port
// while the Open key will contain a boolean for open / closed
// status
type PortScan struct {
	Port   int    `json:"port"`
	Status string `json:"status"`
}

// Scan method will dial in for the referred {host}:{port} with
// a 10-second timeout. The `return` statements will allow for
// continuously probe for open/closed ports without falling back
// on an error
func (p *PortScan) Scan(proto, addr string, port int) *PortScan {

	address := addr + ":" + strconv.Itoa(port)
	var timeout time.Duration = 100000000
	conn, err := net.DialTimeout(proto, address, timeout)

	if err != nil {
		p.Status = "Closed"
		return p
	}
	defer conn.Close()
	p.Status = "Open"

	return p

}

// New method (PortScan) will initiate a new set of PortScan.Scan()
// events, and register only the open ports in the HostScan.Ports[]
// slice of ints
func (p *PortScan) New(h *HostScan, port int) *PortScan {

	p.Scan(h.Protocol, h.Target, port)

	if p.Status != "Closed" {

		h.Ports = append(h.Ports, p.Port)
	}

	return p
}

// New method (HostScan) will initiate a new wave of PortScan.New()
// events for the range defined in the scanScope parameter
// a quick scan would probe 1024 ports while a wide scan would
// target 49152 ports
func (h *HostScan) New(scanScope int) *HostScan {

	for i := 1; i <= scanScope; i++ {

		scan := &PortScan{Port: i}
		scan.New(h, i)
	}

	return h

}

// Create method will initialize a set of HostScan.New() events,
// based on the provided hosts, which are being fed as a slice of
// strings.
func (s *ScanResults) Create(hosts []string, scanScope int) *ScanResults {
	//var wg sync.WaitGroup

	for _, e := range hosts {
		//	wg.Add(1)
		//	go func(wg *sync.WaitGroup) {
		//		defer wg.Done()

		scan := &HostScan{Target: e, Protocol: "tcp"}
		scan.New(scanScope)

		if scan.Ports != nil {
			s.Results = append(s.Results, *scan)
		}
		//	}(&wg)
		//	wg.Wait()
	}

	return s
}
