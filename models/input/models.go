package input

import (
	"fmt"

	"github.com/Ali-Farhadnia/clientGRPC/config"
)

type Input struct {
	Key   string `short:"k" long:"key" description:"use --keylist to see all keys" default:"keylist"`
	Value string `short:"v" long:"value" description:"value to send with key" default:""`
}

func NewInput() *Input {
	return &Input{}
}

func (i Input) Handel() {
	result, err := config.App.Tasks[i.Key](i.Value)
	if err != nil {
		fmt.Println(err)
	}
	if result != "" {
		fmt.Println(result)
	}
}
