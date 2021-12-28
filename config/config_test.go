package config_test

import (
	"encoding/json"
	"testing"

	"github.com/Ali-Farhadnia/clientGRPC/config"
)

func TestNewAppConfig(t *testing.T) {
	a := config.AppConfig{config.GrpcConfig{}}
	b := config.NewAppConfig()
	if a != *b {
		t.Error("NewAppConfig Failed")
	}

}
func TestSetConfig(t *testing.T) {
	app := config.NewAppConfig()
	s := config.AppConfig{config.GrpcConfig{Host: "12.12.12.12", Port: "1234"}}
	res1, _ := json.Marshal(s)
	err := app.SetConfig(res1)
	if err != nil {
		t.Error("some error:", err)
	}
	if app.GrpcConfig.Host != "12.12.12.12" || app.GrpcConfig.Port != "1234" {
		t.Error("SetConfig Failed")
	}
	test := []byte{1, 2}
	err = app.SetConfig(test)
	if err == nil {
		t.Error("SetConfig error test failed")
	}
}
