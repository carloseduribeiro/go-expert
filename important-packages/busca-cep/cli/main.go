package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const fileName = "cities.txt"

func main() {
	for _, cep := range os.Args[1:] {
		fetchCep(cep)
	}
}

func fetchCep(cep string) {
	formatedUrl := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	req, err := http.Get(formatedUrl)
	if err != nil {
		_, err = fmt.Fprintf(os.Stderr, "http request error: %v\n", err)
		if err != nil {
			panic(err)
		}
		return
	}
	defer req.Body.Close()
	resp, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading response: %v\n", err)
		return
	}
	var data ViaCep
	err = json.Unmarshal(resp, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing response: %v\n", err)
		return
	}
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("Cep: %s\tLocalidade: %s\n", data.Cep, data.Localidade))
	if err != nil {
		panic(err)
	}
}

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}
