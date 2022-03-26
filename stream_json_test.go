package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)


type Customer struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Age int `json:"age"`
}

func TestStreamDecoder(t *testing.T) {
	reader, _ := os.Open("sample.json")

	decoder := json.NewDecoder(reader)

	customer := &Customer{}

	decoder.Decode(customer)

	fmt.Println(customer)
}

func TestStreamEncoder(t *testing.T){
	customer := &Customer{
		"fahmi",
		"aziz",
		20,
	}

	writer, _ := os.Create("sample_encode.json")
	encoder := json.NewEncoder(writer)

	encoder.Encode(customer)

}