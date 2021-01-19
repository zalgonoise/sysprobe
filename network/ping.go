package network

import (
	"fmt"
	"math"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ZalgoNoise/sysprobe/utils"
	"github.com/go-ping/ping"
)

// PingScan struct will contain the ping results
// from the last scan
type PingScan struct {
	Target string  `json:"target"`
	Alive  []Alive `json:"alive"`
}

// Alive struct will contain each responding element
// returned in the ping scan
type Alive struct {
	Address string        `json:"addr"`
	Time    time.Duration `json:"rtt"`
}

// Scan interaface will list the different ping options
// possible when using the tool
type Scan interface {
	Burst(addr string) *PingScan
	Paced(addr string) *PingScan
}

// ExpandCIDR method - expands (simple) CIDR addresses
// currently supporting 0/24 addresses and above,
// listing the number of addresses starting at
// 255 and downwards.
// Work needs to be done with this method in the future
func (p *PingScan) ExpandCIDR(ip, cidr string) []string {
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

// New method - issues a new ping event to the provided
// address, while building the PingScan.Alive struct with
// its results
func (p *PingScan) New(wg *sync.WaitGroup, ct int, t time.Duration, h string) {
	defer wg.Done()
	host := h
	timeout := t * 100000
	count := ct
	interval := time.Second
	privileged := false

	pinger, err := ping.NewPinger(host)
	if err != nil {
		panic(err)
	}

	pinger.Timeout = timeout
	pinger.Count = count
	pinger.Interval = interval
	pinger.SetPrivileged(privileged)

	// Listen for Ctrl-C.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			pinger.Stop()
		}
	}()
	pinger.OnRecv = func(pkt *ping.Packet) {

		a := Alive{Address: pkt.IPAddr.String(), Time: pkt.Rtt}

		p.Alive = append(p.Alive, a)

		//fmt.Printf("%s: OK!\t%v\n", pkt.IPAddr, pkt.Rtt)
	}

	//fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
	err = pinger.Run()
	if err != nil {
		panic(err)
	}

}

// Burst method - will take advantage of goroutines to
// issue all ping requests concurrently
// Go will automatically manage this sequence, which is
// aimed for performance, not order
func (p *PingScan) Burst(addr string) *PingScan {
	var wg sync.WaitGroup

	p.Target = addr

	ip := utils.Splitter(addr, "/", 0)

	cidr := utils.Splitter(addr, "/", 1)

	ipList := p.ExpandCIDR(ip, cidr)

	for _, e := range ipList {
		wg.Add(1)
		go p.New(&wg, 1, 30, e)
	}
	wg.Wait()
	return p
}

// Paced method - similar to Burst method, but with no
// goroutines. This will result in a slower, one-by-one
// ping execution, where the results will come back
// indexed as sent
func (p *PingScan) Paced(addr string) *PingScan {
	var wg sync.WaitGroup

	p.Target = addr

	ip := utils.Splitter(addr, "/", 0)

	cidr := utils.Splitter(addr, "/", 1)

	ipList := p.ExpandCIDR(ip, cidr)

	for _, e := range ipList {
		wg.Add(1)
		fmt.Println("# Paced-ping: scanning " + e)
		p.New(&wg, 1, 100, e)
	}
	wg.Wait()
	return p
}
