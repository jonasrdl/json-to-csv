package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {
	jsonData, err := ioutil.ReadFile("people.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var people []Person
	err = json.Unmarshal(jsonData, &people)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(people)
}
