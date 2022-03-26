package main

import (
	"encoding/json"
	"fmt"
	"testing"
)


func logJson(data interface{}){
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

type Data struct {
	Name string
	Num int
	Arr []string
}
func TestEncode(t *testing.T) {
	logJson("Fahmi")
	logJson(1)
	logJson(true)
	logJson([]string{"f","a","z"})
	logJson(map[string]interface{}{
		"key":"value",
		"num": 1,
	})
	logJson(Data{
		"fa",
		1,
		[]string{
			"a",
			"b",
			"c",
		},
	})
}