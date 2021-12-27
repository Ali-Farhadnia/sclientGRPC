package main

import (
	"github.com/Ali-Farhadnia/clientGRPC/cmd"
	"github.com/Ali-Farhadnia/clientGRPC/models/input"
	"github.com/jessevdk/go-flags"
)

var myinput *input.Input

func main() {
	err := cmd.SetConfig()
	if err != nil {
		panic(err)
	}
	myinput = input.NewInput()
	parser := flags.NewParser(myinput, flags.Default)
	_, err = parser.Parse()
	if err != nil {
		panic(err)
	}
	(*myinput).Handel()
}
