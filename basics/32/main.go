package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {

	c := http.Client{Timeout: 10 * time.Millisecond}

	res, err := c.Get("https://google.com")

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

}
