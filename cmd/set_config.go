package cmd

import (
	mainfunctions "github.com/Ali-Farhadnia/clientGRPC/MainFunctions"
	"github.com/Ali-Farhadnia/clientGRPC/config"
	"github.com/Ali-Farhadnia/clientGRPC/connections"
)

func SetConfig() error {
	config.App = config.NewApp()
	err := config.App.Config.SetConfig()
	if err != nil {
		return err
	}
	client, err := connections.GetCRUDClient()
	if err != nil {
		return err
	}
	config.App.GrpcClient = client

	//set main funcs
	config.App.Tasks["insert_one"] = mainfunctions.InsertOneBook
	config.App.Tasks["insert_many"] = mainfunctions.InsertManyBooks
	config.App.Tasks["update"] = mainfunctions.UpdateBook
	config.App.Tasks["delete"] = mainfunctions.DeleteBook
	config.App.Tasks["find_by_id"] = mainfunctions.FindBookByID
	config.App.Tasks["keylist"] = mainfunctions.Help

	return nil
}
