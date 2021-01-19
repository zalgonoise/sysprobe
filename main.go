package main

import (
	"flag"
	"fmt"

	msg "github.com/ZalgoNoise/sysprobe/message"
)

func main() {

	BatteryLoc := flag.String("bat", "battery", "the default location for the battery uevent file, in /sys/class/power_supply/")
	IPDevice := flag.String("net", "wlan0", "the default network device to retrieve IP-related information, with the `ip` command")
	PingAddr := flag.String("ping", "192.168.0.0/24", "the network or subnet address to issue ping events, similar to the *nix `ping` command (but in Go)")

	flag.Parse()

	m := &msg.Message{}
	m = m.New(*BatteryLoc, *IPDevice, *PingAddr)

	output := string(m.JSON())

	fmt.Println(output)

}
