package config

import (
	"encoding/json"
	"os"

	"github.com/Ali-Farhadnia/clientGRPC/models/modelpb"
)

var App *Application

type GrpcConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}
type AppConfig struct {
	GrpcConfig GrpcConfig `json:"grpc_config"`
}

func (a *AppConfig) SetConfig() error {
	file, err := os.ReadFile("./config.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, a)
	if err != nil {
		return err
	}

	return nil
}

type Application struct {
	Config     AppConfig
	GrpcClient modelpb.CRUDClient
	Tasks      map[string]func(string) (string, error)
}

func NewApp() *Application {
	return &Application{AppConfig{GrpcConfig{"", ""}}, nil, make(map[string]func(string) (string, error))}
}
