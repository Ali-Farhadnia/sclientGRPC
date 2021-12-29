package input

import (
	"errors"

	mainfunctions "github.com/Ali-Farhadnia/clientGRPC/MainFunctions"
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

func (i Input) Handel(fm map[string]func(string, modelpb.CRUDClient) (string, error), host string, port string) (string, error) {
	if i.Key == "keylist" {
		res, err := mainfunctions.Help()
		if err != nil {
			return "", err
		}
		return res, nil
	}
	client, err := connections.GetGrpcClient(host, port)
	if err != nil {
		return "", err
	}
	if fm[i.Key] == nil {
		return "", errors.New("unvalid key")
	}
	result, err := fm[i.Key](i.Value, client)
	if err != nil {
		return err.Error(), nil
	}
	return result, nil
}
