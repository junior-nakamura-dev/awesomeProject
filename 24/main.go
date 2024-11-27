package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {

	person := Person{Id: 100, Name: "John"}

	res, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(res))

	json.NewEncoder(os.Stdout).Encode(person)

	rawJson := []byte(`{"id": 200, "name":"Smith"}`)
	var newPerson Person
	err1 := json.Unmarshal(rawJson, &newPerson)
	if err1 != nil {
		fmt.Println(err)
	}

	json.NewEncoder(os.Stdout).Encode(newPerson)

}
