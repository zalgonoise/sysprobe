package main

import (
	"fmt"

	msg "github.com/ZalgoNoise/sysprobe/message"
)

func main() {

	request := &msg.Request{}
	request.Create()

	m := &msg.Response{}
	m = m.New(*request)

	fmt.Println(string(m.JSON()))

}
