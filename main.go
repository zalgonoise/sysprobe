package main

import (
	"fmt"

	msg "github.com/ZalgoNoise/sysprobe/message"
)

func main() {

	request := &msg.Request{}
	response := &msg.Response{}

	request.New()

	response = response.New(*request)

	fmt.Println(string(response.JSON()))

}
