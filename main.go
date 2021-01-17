package main

import (
	"encoding/json"
	"fmt"

	g "github.com/ZalgoNoise/sysprobe/getters"
	u "github.com/ZalgoNoise/sysprobe/utils"
)

func main() {
	b := g.GetBattery()
	i := g.GetIP()

	m := g.MakeMsg(i, b)

	json, err := json.Marshal(m)
	u.Check(err)
	fmt.Println(string(json))

}
