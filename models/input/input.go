package input

import (
	"fmt"
	"log"

	mainfunctions "github.com/Ali-Farhadnia/clientGRPC/MainFunctions"
	"github.com/Ali-Farhadnia/clientGRPC/config"
	"github.com/Ali-Farhadnia/clientGRPC/connections"
	"github.com/Ali-Farhadnia/clientGRPC/models/modelpb"
)

type Input struct {
	Key   string `short:"k" long:"key" description:"use --keylist to see all keys" default:"keylist"`
	Value string `short:"v" long:"value" description:"value to send with key" default:""`
}

func NewInput() *Input {
	return &Input{}
}

func (i Input) Handel(fm map[string]func(string, modelpb.CRUDClient) (string, error), grpcconfig config.GrpcConfig) {
	if i.Key == "keylist" {
		fmt.Println(mainfunctions.Help())
		return
	}
	client, err := connections.GetGrpcClient(grpcconfig.Host, grpcconfig.Port)
	if err != nil {
		log.Println(err)
		return
	}
	result, err := fm[i.Key](i.Value, client)
	if err != nil {
		fmt.Println(err)
	}
	if result != "" {
		fmt.Println(result)
	}
}
