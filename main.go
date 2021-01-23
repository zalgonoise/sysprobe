// Sysprobe is a simple, fast and modular binary to collect metadata
// from your devices and push it to a NoSQL database in your network
//
// == I am talking to you, old rooted smartphone from 2016
// which I no longer use but refuse to throw away ==
//
// The point of the project is to have a one-binary-for-all when it comes to
// collecting heartbeats from devices, regardless of the complexity or the
// amount of data you want to collect.
//
// Currently it collects boring metadata such as the device's battery status
// and its current IP address, but it's also packing a super powerful
// network scanner that will ping all hosts in a subnet, and figure out their
// open ports in just a few seconds. It puts nmap to shame. Boo, nmap, boo.
//
package main

import (
	"flag"
	"fmt"

	msg "github.com/ZalgoNoise/sysprobe/message"
)

func main() {

	request := &msg.Request{}

	if request.HelpOpt != false {
		flag.Usage()
	} else {

		response := &msg.Response{}

		request.New()

		response = response.New(*request)

		fmt.Println(string(response.JSON()))
	}
}
