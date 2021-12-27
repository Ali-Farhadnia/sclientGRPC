package book

import (
	"encoding/json"
)

//convert book th the json string
func (book *Book) ToString() (string, error) {
	result, err := json.Marshal(book)
	if err != nil {
		return "", err
	}
	sResult := string(result)
	return sResult, nil

}
