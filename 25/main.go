package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Address struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	fmt.Println("Process Started!")

	url := "https://viacep.com.br/ws/%s/json/"

	//Get the CEP number from CMD input
	for _, cep := range os.Args[1:] {
		url = fmt.Sprintf(url, cep)
	}

	req, errRequest := http.Get(url)
	if errRequest != nil {
		fmt.Fprintf(os.Stderr, "Request error: %v\n", errRequest)
	}

	defer req.Body.Close()
	res, errReadAll := io.ReadAll(req.Body)
	if errReadAll != nil {
		fmt.Fprintf(os.Stderr, "Read file error: %v", errReadAll)
	}

	var data Address
	errUnmarshal := json.Unmarshal(res, &data)
	if errUnmarshal != nil {
		fmt.Fprintf(os.Stderr, "Unmarshal error: %v\n", errUnmarshal)
	}

	file, errCreateFile := os.Create("addresses.txt")
	if errCreateFile != nil {
		fmt.Fprintf(os.Stderr, "Create file error: %v\n", errCreateFile)
	}
	defer file.Close()

	_, errWriteFile := io.WriteString(file, fmt.Sprintf("CEP: %s , Logradouro: %s", data.Cep, data.Logradouro))
	if errWriteFile != nil {
		fmt.Fprintf(os.Stderr, "Write file error: %v\n", errWriteFile)
	}

	fmt.Println("Process Finished!")

	fmt.Println(data.Cep)

}

//go build -o cep main.go to create an executable with name "cep
