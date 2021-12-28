package connections_test

import (
	"testing"

	"github.com/Ali-Farhadnia/clientGRPC/connections"
)

func TestGetGrpcClient(t *testing.T) {
	cli, err := connections.GetGrpcClient("37.152.177.253", "8082")
	if err != nil {
		t.Error("some error:", err)
	}
	if cli == nil {
		t.Error("GetGrpcClient failed")
	}
	//testing error

}
