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
	http.HandleFunc("/", FetchCepHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func FetchCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	cepParam := r.URL.Query().Get("cep")
	if len(cepParam) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cep, err := fetchCep(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(cep)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func fetchCep(cep string) (*ViaCep, error) {
	formatedUrl := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	req, err := http.Get(formatedUrl)
	if err != nil {
		_, err = fmt.Fprintf(os.Stderr, "http request error: %v\n", err)
		return nil, err
	}
	defer req.Body.Close()
	resp, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading response: %v\n", err)
		return nil, err
	}
	var data ViaCep
	err = json.Unmarshal(resp, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing response: %v\n", err)
		return nil, err
	}
	return &data, nil
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
