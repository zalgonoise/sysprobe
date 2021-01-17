package main

import (
	"encoding/json"
	"fmt"

	"github.com/ZalgoNoise/sysprobe/battery"
	"github.com/ZalgoNoise/sysprobe/internet"
	"github.com/ZalgoNoise/sysprobe/message"
	"github.com/ZalgoNoise/sysprobe/utils"
)

func main() {
	b := battery.GetBattery()
	i := internet.GetIP()

	m := message.MakeMsg(i, b)

	json, err := json.Marshal(m)
	utils.Check(err)
	fmt.Println(string(json))

}
