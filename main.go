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
