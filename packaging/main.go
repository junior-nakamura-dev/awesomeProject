package main

import (
	"fmt"
	"packaging/client"
)

func main() {

	var client100 = client.NewClient(1, "Junior")

	fmt.Printf("Id: %d Name: %s \n", client100.GetId(), client100.GetName())
}
