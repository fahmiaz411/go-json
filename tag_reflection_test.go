package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJSONTagEncode(t *testing.T) {
	product := &Product{
		"1",
		"Apple",
		"http://localhost/",
	}

	bytes, _ := json.Marshal(product)

	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonBytes := []byte(`{"id":"1","name":"Apple","image_url":"http://localhost/"}`)
	
	product := &Product{}
	productMap := map[string]interface{}{}

	json.Unmarshal(jsonBytes, product)
	json.Unmarshal(jsonBytes, &productMap)


	fmt.Println(productMap["image_url"])
}