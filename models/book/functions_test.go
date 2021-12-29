package book_test

import (
	"testing"

	"github.com/Ali-Farhadnia/clientGRPC/models/book"
)

func TestToString(t *testing.T) {
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
