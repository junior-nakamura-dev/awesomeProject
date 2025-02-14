package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

	http.HandleFunc("/", fetchCEPHandler)
	http.ListenAndServe(":8080", nil)

}

func fetchCEPHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	cepParam := req.URL.Query().Get("cep")
	if cepParam == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	data, err := fetchCEP(cepParam)

	if err != nil {
		res.WriteHeader(http.StatusConflict)
		res.Write([]byte(err.Error()))
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(data)
}

func fetchCEP(cep string) (*Address, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)

	var data Address
	err = json.Unmarshal(body, &data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}
