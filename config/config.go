package config

import (
	"encoding/json"
)

type GrpcConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}
type AppConfig struct {
	GrpcConfig GrpcConfig `json:"grpc_config"`
}

func (a *AppConfig) SetConfig(file []byte) error {
	err := json.Unmarshal(file, a)
	if err != nil {
		return err
	}

	return nil
}

func NewAppConfig() *AppConfig {
	return &AppConfig{GrpcConfig{}}
}
