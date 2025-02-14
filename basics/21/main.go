package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//creating a file
	file, err := os.Create("input.txt")

	if err != nil {
		panic(err)
	}

	//write in a file
	size, err := file.WriteString("I'm new file")
	if err != nil {
		panic(err)
	}
	fmt.Printf("File succesfully created with %d bytes \n", size)
	file.Close()

	//reading a file
	readFile, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(readFile))

	//reading file in chunks
	newFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(newFile)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	newFile.Close()

}
