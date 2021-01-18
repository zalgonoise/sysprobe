package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/ZalgoNoise/sysprobe/battery"
	"github.com/ZalgoNoise/sysprobe/internet"
	"github.com/ZalgoNoise/sysprobe/message"
	"github.com/ZalgoNoise/sysprobe/utils"
)

func main() {

	BatteryLoc := flag.String("bat", "battery", "the default location for the battery uevent file, in /sys/class/power_supply/")
	IPDevice := flag.String("net", "wlan0", "the default network device to retrieve IP-related information, with the `ip` command")
	flag.Parse()

	b := battery.GetBattery(*BatteryLoc)
	i := internet.GetIP(*IPDevice)

	m := message.MakeMsg(i, b)

	json, err := json.Marshal(m)
	utils.Check(err)
	fmt.Println(string(json))

}
