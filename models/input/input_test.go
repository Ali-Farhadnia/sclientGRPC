package input_test

import (
	"errors"
	"testing"

	mainfunctions "github.com/Ali-Farhadnia/clientGRPC/MainFunctions"
	"github.com/Ali-Farhadnia/clientGRPC/models/input"
	"github.com/Ali-Farhadnia/clientGRPC/models/modelpb"
)

const grpc_host = "37.152.177.253"
const grpc_port = "8083"

func TestNewInput(t *testing.T) {
	a := input.Input{}
	b := input.NewInput()
	if a != *b {
		t.Error("NewInput Failed")
	}

}

func TestHandel(t *testing.T) {
	m := make(map[string]func(string, modelpb.CRUDClient) (string, error))
	f1 := func(s string, mo modelpb.CRUDClient) (string, error) {
		if s == "error" {
			return "", errors.New("error test")
		}

		return "test1" + s, nil
	}
	f2 := func(s string, mo modelpb.CRUDClient) (string, error) {
		return "test2" + s, nil
	}
	m["f1"] = f1
	m["f2"] = f2

	i := input.NewInput()
	i.Key = "f2"
	i.Value = "1234"
	// f2.
	res, err := i.Handel(m, grpc_host, grpc_port)
	if err != nil {
		t.Error("some error:", err)
	}
	if res != "test21234" {
		t.Error("Handle failed", res)
	}
	// help func.
	h, err := mainfunctions.Help()
	if err != nil {
		t.Error("some error:", err)
	}
	i.Key = "keylist"
	res, err = i.Handel(m, grpc_host, grpc_port)
	if err != nil {
		t.Error("some error:", err)
	}
	if res != h {
		t.Error("Handle failed", res)
	}
	// return error.
	i.Key = "f1"
	i.Value = "error"
	res, err = i.Handel(m, grpc_host, grpc_port)
	if err != nil {
		t.Error("some error:", err)
	}
	if res != "error test" {
		t.Error("return error failed", res)
	}

	// unvalid key.
	i.Key = "11"
	res, err = i.Handel(m, grpc_host, grpc_port)
	if err.Error() != "unvalid key" || res != "" {
		t.Error("unvalid key failed", err)
	}

}

/*
func TestToSring(t *testing.T) {
	b := book.Book{ID: "1234", Name: "test1", Author: "test1", Pagecount: 50, Inventory: 50}
	bs := `{"ID":"1234","Name":"test1","Author":"test1","Pagecount":50,"Inventory":50}`
	res, err := b.ToString()
	if err != nil {
		t.Error("some error:", err)
	}
	if bs != res {
		t.Error(res)
	}

}
*/
