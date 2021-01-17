package main

import (
	"encoding/json"
	"fmt"

	"github.com/ZalgoNoise/sysprobe/getters"
	"github.com/ZalgoNoise/sysprobe/utils"
)

func main() {
	b := getters.GetBattery()
	i := getters.GetIP()

	m := getters.MakeMsg(i, b)

	json, err := json.Marshal(m)
	utils.Check(err)
	fmt.Println(string(json))

}
