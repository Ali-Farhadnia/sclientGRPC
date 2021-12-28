package main

import (
	"os"

	mainfunctions "github.com/Ali-Farhadnia/clientGRPC/MainFunctions"
	"github.com/Ali-Farhadnia/clientGRPC/config"
	"github.com/Ali-Farhadnia/clientGRPC/models/input"
	"github.com/jessevdk/go-flags"
)

func main() {
	// open config.json.
	file, err := os.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	// create and set config.
	var appconfig = config.NewAppConfig()
	err = appconfig.SetConfig(file)
	if err != nil {
		panic(err)
	}
	// set input parser and parse it.
	var myinput = input.NewInput()
	parser := flags.NewParser(myinput, flags.Default)
	_, err = parser.Parse()
	if err != nil {
		panic(err)
	}
	// set main functions to handel input.
	funcs := mainfunctions.GetMainFuncs()
	// handel input
	myinput.Handel(funcs, appconfig.GrpcConfig)
}
