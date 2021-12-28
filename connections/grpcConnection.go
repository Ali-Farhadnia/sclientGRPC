package connections

import (
	"sync"

	"github.com/Ali-Farhadnia/clientGRPC/models/modelpb"
	"google.golang.org/grpc"
)

/* Used to create a singleton object of grpc client.
Initialized and exposed through GetGrpcClient().*/
var clientInstance modelpb.CRUDClient

//Used during creation of singleton client object in GetGrpcClient().
var clientInstanceError error

//Used to execute client creation procedure only once.
var GrpcOnce sync.Once

// GetGrpcClient() set  modelpb.CRUDClient with client.
func GetGrpcClient(host string, port string) (modelpb.CRUDClient, error) {
	GrpcOnce.Do(func() {
		target := host + ":" + port
		cc, err := grpc.Dial(target, grpc.WithInsecure())
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = modelpb.NewCRUDClient(cc)
	})

	return clientInstance, clientInstanceError
}
