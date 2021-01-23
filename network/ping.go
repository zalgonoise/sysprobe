package network

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"

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

func (p *PingScan) Get() []string {
	var allHosts []string

	for _, v := range p.Alive {
		allHosts = append(allHosts, v.Address)
	}
	return allHosts
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
func (p *PingScan) Burst(mwg *sync.WaitGroup, addr []string) *PingScan {
	defer mwg.Done()
	var wg sync.WaitGroup

	for _, e := range addr {
		wg.Add(1)
		go p.New(&wg, 1, 10000, e)
	}
	wg.Wait()
	return p
}

// Paced method - similar to Burst method, but with no
// goroutines. This will result in a slower, one-by-one
// ping execution, where the results will come back
// indexed as sent
func (p *PingScan) Paced(mwg *sync.WaitGroup, addr []string) *PingScan {
	defer mwg.Done()
	var wg sync.WaitGroup

	for _, e := range addr {
		fmt.Println("Scanned IPv4 " + e)
		wg.Add(1)
		p.New(&wg, 1, 10000, e)
	}
	wg.Wait()
	return p
}
