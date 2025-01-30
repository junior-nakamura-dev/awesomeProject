package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {

	c := http.Client{}

	jsonVar := bytes.NewBuffer([]byte(`{"name": "Junior""}`))

	res, err := c.Post("https://google.com", "application/json", jsonVar)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	io.CopyBuffer(os.Stdout, res.Body, nil)
}
