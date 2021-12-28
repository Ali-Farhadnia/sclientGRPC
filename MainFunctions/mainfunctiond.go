package mainfunctions

import (
	"context"
	"strconv"
	"strings"

	"github.com/Ali-Farhadnia/clientGRPC/models/book"
	"github.com/Ali-Farhadnia/clientGRPC/models/modelpb"
)

// all the functions get string as input then pars that string and send it to grpc server.
const unvalid = "unvalid input"

func InsertOneBook(input string, client modelpb.CRUDClient) (string, error) {
	book, err := UnarshalStringOne(input)
	if err != nil {
		return unvalid, err
	}
	if !checknull(book) {
		return unvalid, err
	}
	grpcbooklist := []*modelpb.Book{}
	Pagescount, err := strconv.ParseInt(book["pagecount"], 10, 32)
	if err != nil {
		return "", err
	}
	Inventory, err := strconv.ParseInt(book["inventory"], 10, 64)
	if err != nil {
		return "", err
	}
	grpcbooklist = append(grpcbooklist, &modelpb.Book{Name: book["name"],
		Author: book["author"], Pagescount: int32(Pagescount), Inventory: Inventory})
	var grpcbooks = modelpb.Books{
		Books: grpcbooklist,
	}
	Status, err := client.InsertBook(context.Background(), &grpcbooks)
	if err != nil {
		return "somthing whent wrong", err
	}

	if Status.Status == "no" {
		return "somthing went wrong:" + Status.Description, err
	}

	return Status.Description, nil
}
func InsertManyBooks(input string, client modelpb.CRUDClient) (string, error) {
	books, err := UnarshalStringMany(input)
	if err != nil {
		return unvalid, err
	}
	if !checknulls(books) {
		return unvalid, err
	}
	grpcbooklist := []*modelpb.Book{}
	for _, book := range books {
		Pagescount, err := strconv.ParseInt(book["pagecount"], 10, 32)
		if err != nil {
			return "", err
		}
		Inventory, err := strconv.ParseInt(book["inventory"], 10, 64)
		if err != nil {
			return "", err
		}
		grpcbooklist = append(grpcbooklist, &modelpb.Book{Name: book["name"],
			Author: book["author"], Pagescount: int32(Pagescount), Inventory: Inventory})
	}
	var grpcbooks = modelpb.Books{
		Books: grpcbooklist,
	}
	Status, err := client.InsertBook(context.Background(), &grpcbooks)
	if err != nil {
		return "somthing whent wrong", err
	}
	if Status.Status == "no" {
		return "somthing whent wrong:" + Status.Description, err
	}

	return Status.Description, nil
}
func UpdateBook(input string, client modelpb.CRUDClient) (string, error) {
	book, err := UnarshalStringOne(input)
	if err != nil {
		return unvalid, err
	}
	if !checknull(book) {
		return unvalid, err
	}
	Pagescount, err := strconv.ParseInt(book["pagecount"], 10, 32)
	if err != nil {
		return "", err
	}
	Inventory, err := strconv.ParseInt(book["inventory"], 10, 64)
	if err != nil {
		return "", err
	}

	rec := modelpb.UpdateRequest{

		Book: &modelpb.Book{Name: book["name"],
			Author:     book["author"],
			Pagescount: int32(Pagescount),
			Inventory:  Inventory}, Id: book["id"],
	}

	Status, err := client.UpdateBook(context.Background(), &rec)
	if err != nil {
		return "somthing whent wrong", err
	}
	if Status.Status == "no" {
		return "somthing whent wrong:" + Status.Description, err
	}

	return Status.Description, nil
}
func DeleteBook(input string, client modelpb.CRUDClient) (string, error) {
	rec := modelpb.BookID{Id: input}
	Status, err := client.DeleteBook(context.Background(), &rec)
	if err != nil {
		return "somthing whent wrong", err
	}
	if Status.Status == "no" {
		return "somthing whent wrong:" + Status.Description, err
	}

	return Status.Description, nil
}
func FindBookByID(input string, client modelpb.CRUDClient) (string, error) {
	rec := modelpb.BookID{Id: input}
	res, err := client.FindBookById(context.Background(), &rec)
	if err != nil {
		return "somthing whent wrong", err
	}
	if res.Status.Status == "no" {
		return "somthing whent wrong:" + res.Status.Description, err
	}
	resbook := book.Book{
		ID:        res.Book.Id,
		Name:      res.Book.Name,
		Author:    res.Book.Author,
		Pagecount: res.Book.Pagescount,
		Inventory: res.Book.Inventory,
	}

	sres, err := resbook.ToString()
	if err != nil {
		return "somthing whent wrong", err
	}

	return sres, nil
}

// each time  Help(string) called it returns help string.
func Help() (string, error) {
	help := `
	Functions:

		insert_one:		e.g. value={"name":string,"author":string,"pagecount": integer,"Inventory":integer}

		insert_many:        e.g.value=[{"name":string,"author":string,"pagecount": integer,"Inventory":integer},
								{"name":string,"author":string,"pagecount": integer,"Inventory":integer}]

		update:		e.g.value={"id":string,"name":string,"author":string,"pagecount": integer,"Inventory":integer}

		delete:       	e.g.value=string

		find_by_id:           e.g.value=string
	`

	return help, nil
}

// get json string and parse it to the book.
func UnarshalStringOne(input string) (map[string]string, error) {
	input = strings.ReplaceAll(input, "}", "")
	input = strings.ReplaceAll(input, "{", "")
	splited := strings.Split(input, ",")
	m := make(map[string]string, len(splited))
	for _, v := range splited {
		result := strings.Split(v, ":")
		m[result[0]] = result[1]
	}

	return m, nil
}

// get json string and parse it to the books.
func UnarshalStringMany(input string) ([]map[string]string, error) {
	input = strings.ReplaceAll(input, "[", "")
	input = strings.ReplaceAll(input, "]", "")
	splited := strings.Split(input, "-")
	var splits []map[string]string
	for _, v := range splited {
		result, err := UnarshalStringOne(v)
		if err != nil {
			return nil, err
		}
		splits = append(splits, result)
	}

	return splits, nil
}

func checknull(m map[string]string) bool {
	if m["name"] == "" || m["author"] == "" || m["inventory"] == "0" || m["Pagecount"] == "0" {
		return false
	}

	return true
}
func checknulls(ms []map[string]string) bool {
	for _, m := range ms {
		if m["name"] == "" || m["author"] == "" || m["inventory"] == "0" || m["Pagecount"] == "0" {
			return false
		}
	}

	return true
}

// return main functions as a map
func GetMainFuncs() map[string]func(string, modelpb.CRUDClient) (string, error) {
	funcs := make(map[string]func(string, modelpb.CRUDClient) (string, error))
	// set main funcs
	funcs["insert_one"] = InsertOneBook
	funcs["insert_many"] = InsertManyBooks
	funcs["update"] = UpdateBook
	funcs["delete"] = DeleteBook
	funcs["find_by_id"] = FindBookByID
	return funcs
}
