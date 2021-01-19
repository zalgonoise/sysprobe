package main

import (
	"flag"
	"fmt"

	msg "github.com/ZalgoNoise/sysprobe/message"
)

func main() {

	BatteryLoc := flag.String("bat", "battery", "the default location for the battery uevent file, in /sys/class/power_supply/")
	IPDevice := flag.String("net", "wlan0", "the default network device to retrieve IP-related information, with the `ip` command")
	flag.Parse()

	m := &msg.Message{}
	m = m.New(*BatteryLoc, *IPDevice)

	output := string(m.JSON())

	fmt.Println(output)

}
