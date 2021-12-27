package connections

import (
	"fmt"

	"github.com/Ali-Farhadnia/clientGRPC/config"
	"github.com/Ali-Farhadnia/clientGRPC/models/modelpb"
	"google.golang.org/grpc"
)

//GetCRUDClient() set  modelpb.CRUDClient with client
func GetCRUDClient() (modelpb.CRUDClient, error) {
	fmt.Println("Hello i am client")
	target := config.App.Config.GrpcConfig.Host + ":" + config.App.Config.GrpcConfig.Port
	cc, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	Client := modelpb.NewCRUDClient(cc)
	return Client, nil
}
