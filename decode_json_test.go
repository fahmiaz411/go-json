package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	jsonString := `{"Name":"fa","Num":1,"Arr":["a","b","c"]}`
	jsonString2 := `{"Name":"ega","Num":2,"Arr":["ss","bc","c"]}`
	jsonBytes := []byte(jsonString)
	jsonBytes2 := []byte(jsonString2)

	customer := &Data{}
	cus2 := &Data{}

	err := json.Unmarshal(jsonBytes, customer)
	json.Unmarshal(jsonBytes2, cus2)

	if err != nil {
		panic(err)
	}

	fmt.Println(customer.Name, cus2.Arr)
}