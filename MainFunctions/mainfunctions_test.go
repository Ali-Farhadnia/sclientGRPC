package mainfunctions_test

import (
	"fmt"
	"testing"

	mainfunctions "github.com/Ali-Farhadnia/clientGRPC/MainFunctions"
	"github.com/Ali-Farhadnia/clientGRPC/connections"
	"github.com/Ali-Farhadnia/clientGRPC/models/book"
)

var id string

func TestInsertOneBook(t *testing.T) {
	cli, err := connections.GetGrpcClient("37.152.177.253", "8082")
	if err != nil {
		t.Error("some error:", err)
	}
	// correct input.
	test1 := `{name:test1,author:test1,pagecount:50,inventory:50}`
	res, err := mainfunctions.InsertOneBook(test1, cli)
	if err != nil {
		t.Error("some error:", err)
	} else {
		id = res
	}
	if res == "" {
		t.Error("InsertOneBook failed")
	}
	// wrong input.
	test2 := `aaa`
	res, err = mainfunctions.InsertOneBook(test2, cli)
	if err == nil || res != mainfunctions.Unvalid {
		t.Error("InsertOneBook empty input failed")
	}
	// empty input.
	test3 := ``
	res, err = mainfunctions.InsertOneBook(test3, cli)
	if err == nil || res != mainfunctions.Unvalid {
		t.Error("InsertOneBook empty input failed")
	}
	// wrong pagecount
	test4 := `{name:test1,author:test1,pagecount:aaa,inventory:50}`
	res, err = mainfunctions.InsertOneBook(test4, cli)
	if err == nil || res != "" {
		t.Error("InsertOneBook wrong pagecount input failed", err, res)
	}

	// wrong inventory
	test5 := `{name:test1,author:test1,pagecount:50,inventory:aaa}`
	res, err = mainfunctions.InsertOneBook(test5, cli)
	if err == nil || res != "" {
		t.Error("InsertOneBook wrong inventory input failed")
	}
	// null filed.
	test6 := `{name:test1,author:,pagecount:50,inventory:50}`
	res, err = mainfunctions.InsertOneBook(test6, cli)
	if err == nil || res != mainfunctions.Unvalid {
		t.Error("InsertOneBook null filed input failed")
	}
	/*
		// grpc error
		cli, err = connections.GetGrpcClient("48.8.2.1", "8086")
		if err != nil {
			t.Error("some error:", err)
		}

		res, err = mainfunctions.InsertOneBook(test1, cli)
		if err == nil {
			t.Error("InsertOneBook wrong grpc connection failed", err, res)
		}
	*/

}

func TestInsertManyBooks(t *testing.T) {
	cli, err := connections.GetGrpcClient("37.152.177.253", "8082")
	if err != nil {
		t.Error("some error:", err)
	}
	// correct input.
	test1 := `[{name:test1,author:test1,pagecount:50,inventory:50}-{name:test2,author:test1,pagecount:50,inventory:50}-{name:test3,author:test1,pagecount:50,inventory:50}]`
	res, err := mainfunctions.InsertManyBooks(test1, cli)
	if err != nil {
		t.Error("some error:", err)
	}
	if res == "" {
		t.Error("InsertManyBooks failed")
	}
	// wrong input.
	test2 := `aaa`
	res, err = mainfunctions.InsertManyBooks(test2, cli)
	if err == nil || res != mainfunctions.Unvalid {
		t.Error("InsertManyBooks empty input failed")
	}
	// empty input.
	test3 := ``
	res, err = mainfunctions.InsertManyBooks(test3, cli)
	if err == nil || res != mainfunctions.Unvalid {
		t.Error("InsertManyBooks empty input failed")
	}
	// wrong pagecount
	test4 := `[{name:test1,author:test1,pagecount:aaa,inventory:50}-{name:test2,author:test1,pagecount:50,inventory:50}]`
	res, err = mainfunctions.InsertManyBooks(test4, cli)
	if err == nil || res != "" {
		t.Error("InsertManyBooks wrong pagecount input failed", err, res)
	}

	// wrong inventory
	test5 := `[{name:test1,author:test1,pagecount:50,inventory:aaa}-{name:test2,author:test1,pagecount:50,inventory:50}]`
	res, err = mainfunctions.InsertManyBooks(test5, cli)
	if err == nil || res != "" {
		t.Error("InsertManyBooks wrong inventory input failed")
	}
	// null filed.
	test6 := `[{name:test1,author:,pagecount:50,inventory:50}-{name:test2,author:test1,pagecount:50,inventory:50}]`
	res, err = mainfunctions.InsertManyBooks(test6, cli)
	if err == nil || res != mainfunctions.Unvalid {
		t.Error("InsertManyBooks null filed input failed", res, err)
	}
	/*
		// grpc error
		cli, err = connections.GetGrpcClient("48.8.2.1", "8086")
		if err != nil {
			t.Error("some error:", err)
		}

		res, err = mainfunctions.InsertManyBooks(test1, cli)
		if err == nil {
			t.Error("InsertManyBooks wrong grpc connection failed", err, res)
		}
	*/
}

func TestUpdateBook(t *testing.T) {
	cli, err := connections.GetGrpcClient("37.152.177.253", "8082")
	if err != nil {
		t.Error("some error:", err)
	}
	// correct input.
	test1 := fmt.Sprintf(`{id:%s,name:test1,author:test1,pagecount:50,inventory:50}`, id)
	res, err := mainfunctions.UpdateBook(test1, cli)
	if err != nil {
		t.Error("some error:", err)
	}
	if res == "" {
		t.Error("UpdateBook failed")
	}
	// wrong input.
	test2 := `aaa`
	res, err = mainfunctions.UpdateBook(test2, cli)
	if err == nil || res != mainfunctions.Unvalid {
		t.Error("UpdateBook empty input failed")
	}
	// empty input.
	test3 := ``
	res, err = mainfunctions.UpdateBook(test3, cli)
	if err == nil || res != mainfunctions.Unvalid {
		t.Error("UpdateBook empty input failed")
	}
	// wrong pagecount
	test4 := fmt.Sprintf(`{id:%s,name:test1,author:test1,pagecount:aaa,inventory:50}`, id)
	res, err = mainfunctions.UpdateBook(test4, cli)
	if err == nil || res != "" {
		t.Error("UpdateBook wrong pagecount input failed", err, res)
	}

	// wrong inventory
	test5 := fmt.Sprintf(`{id:%s,name:test1,author:test1,pagecount:50,inventory:aaa}`, id)
	res, err = mainfunctions.UpdateBook(test5, cli)
	if err == nil || res != "" {
		t.Error("UpdateBook wrong inventory input failed")
	}
	// null filed.
	test6 := fmt.Sprintf(`{id:%s,name:test1,author:,pagecount:50,inventory:50}`, id)
	res, err = mainfunctions.UpdateBook(test6, cli)
	if err == nil || res != mainfunctions.Unvalid {
		t.Error("UpdateBook null filed input failed")
	}
	/*
		// grpc error
		cli, err = connections.GetGrpcClient("48.8.2.1", "8086")
		if err != nil {
			t.Error("some error:", err)
		}

		res, err = mainfunctions.UpdateBook(test1, cli)
		if err == nil {
			t.Error("UpdateBook wrong grpc connection failed", err, res)
		}
	*/
}
func TestDeleteBook(t *testing.T) {
	cli, err := connections.GetGrpcClient("37.152.177.253", "8082")
	if err != nil {
		t.Error("some error:", err)
	}
	res, err := mainfunctions.DeleteBook(id, cli)
	if err != nil || res == "" {
		t.Error("DeleteBook failed", id, err)
	}
	// wrong input.
	test2 := "1234"
	res, err = mainfunctions.UpdateBook(test2, cli)
	if err == nil || res != mainfunctions.Unvalid {
		t.Error("UpdateBook wrong input failed")
	}
	// empty input.
	test3 := ``
	res, err = mainfunctions.UpdateBook(test3, cli)
	if err == nil || res != mainfunctions.Unvalid {
		t.Error("UpdateBook empty input failed")
	}
}
func TestFindBookByID(t *testing.T) {
	cli, err := connections.GetGrpcClient("37.152.177.253", "8082")
	if err != nil {
		t.Error("some error:", err)
	}
	b := book.Book{ID: id, Name: "test1", Author: "test1", Pagecount: 50, Inventory: 50}

	res, err := mainfunctions.FindBookByID(b.ID, cli)
	if err != nil {
		t.Errorf("FindBookByID failed. id:%s , err:%e", id, err)
	}
	sb, err := b.ToString()
	if err != nil {
		t.Errorf("FindBookByID failed. id:%s , err:%e", id, err)
	}
	if res != sb {
		t.Errorf("FindBookByID failed. id:%s ,result:%s , err:%e", id, res, err)
	}

	// wrong input.
	test2 := "1234"
	res, err = mainfunctions.UpdateBook(test2, cli)
	if err == nil || res != mainfunctions.Unvalid {
		t.Error("FindBookByID wrong input failed")
	}
	// empty input.
	test3 := ``
	res, err = mainfunctions.UpdateBook(test3, cli)
	if err == nil || res != mainfunctions.Unvalid {
		t.Error("FindBookByID empty input failed")
	}
}

func TestUnarshalStringOne(t *testing.T) {
	test1 := `{name:test1,author:test1,pagecount:50,inventory:50}`
	m, err := mainfunctions.UnarshalStringOne(test1)
	if err != nil {
		t.Error("some error:", err)
	}
	if m["name"] != "test1" || m["author"] != "test1" || m["pagecount"] != "50" || m["inventory"] != "50" || len(m) != 4 {
		t.Error("UnarshalStringOne failed", m)
	}
	// wrong input.
	test2 := "name:test1,author:test1,pagecount:50,inventory:50"
	m, err = mainfunctions.UnarshalStringOne(test2)
	if err == nil || m != nil {
		t.Error("UnarshalStringOne wrong input failed")
	}
}

func TestUnarshalStringMany(t *testing.T) {
	test1 := `[{name:test1,author:test1,pagecount:50,inventory:50}-{name:test2,author:test1,pagecount:50,inventory:50}-{name:test3,author:test1,pagecount:50,inventory:50}]`
	m, err := mainfunctions.UnarshalStringMany(test1)
	if err != nil {
		t.Error("some error:", err)
	}

	if m[0]["name"] != "test1" || m[0]["author"] != "test1" || m[0]["pagecount"] != "50" || m[0]["inventory"] != "50" || len(m[0]) != 4 {
		t.Error("UnarshalStringMany failed", m[0])
	}
	if m[1]["name"] != "test2" || m[1]["author"] != "test1" || m[1]["pagecount"] != "50" || m[1]["inventory"] != "50" || len(m[1]) != 4 {
		t.Error("UnarshalStringMany failed", m[0])
	}
	if m[2]["name"] != "test3" || m[2]["author"] != "test1" || m[2]["pagecount"] != "50" || m[2]["inventory"] != "50" || len(m[2]) != 4 {
		t.Error("UnarshalStringMany failed", m[0])
	}
	// wrong input.
	test2 := `[{name:test1,author:test1,pagecount:50,inventory:50}-name:test2,author:test1,pagecount:50,inventory:50-{name:test3,author:test1,pagecount:50,inventory:50}]`
	m, err = mainfunctions.UnarshalStringMany(test2)
	if err == nil || m != nil {
		t.Error("UnarshalStringMany wrong input failed")
	}
}
func TestChecknull(t *testing.T) {
	m := make(map[string]string)
	m["name"] = "test"
	m["author"] = ""
	m["inventory"] = "1"
	m["Pagecount"] = "2"
	//null input
	if mainfunctions.Checknull(m) {
		t.Error("Checknull null input failed")
	}
	m["author"] = "test"
	//correct input
	if !mainfunctions.Checknull(m) {
		t.Error("Checknull correct input failed")
	}
}

func TestChecknulls(t *testing.T) {
	var m []map[string]string
	m1 := make(map[string]string)
	m2 := make(map[string]string)
	m1["name"] = "test1"
	m1["author"] = ""
	m1["inventory"] = "1"
	m1["Pagecount"] = "2"
	m = append(m, m1)
	m2["name"] = "test2"
	m2["author"] = "test"
	m2["inventory"] = "1"
	m2["Pagecount"] = "2"
	m = append(m, m2)
	//null input
	if mainfunctions.Checknulls(m) {
		t.Error("Checknull null input failed")
	}
	m[0]["author"] = "test"
	//correct input
	if !mainfunctions.Checknulls(m) {
		t.Error("Checknull correct input failed")
	}
}
func TestHelp(t *testing.T) {
	help := `
	Functions:

		insert_one:		e.g. value={name:string,author:string,pagecount: integer,Inventory:integer}

		insert_many:        e.g.value=[{name:string,author:string,pagecount: integer,Inventory:integer},
								{name:string,author:string,pagecount: integer,Inventory:integer}]

		update:		e.g.value={id:string,name:string,author:string,pagecount: integer,Inventory:integer}

		delete:       	e.g.value=string

		find_by_id:           e.g.value=string
	`
	res, _ := mainfunctions.Help()
	if help != res {
		t.Error("Help failed")

	}
}
